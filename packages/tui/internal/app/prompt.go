package app

import (
	"errors"
	"time"

	"github.com/sst/axoncode-sdk-go"
	"github.com/sst/axoncode/internal/attachment"
	"github.com/sst/axoncode/internal/id"
)

type Prompt struct {
	Text        string                   `toml:"text"`
	Attachments []*attachment.Attachment `toml:"attachments"`
}

func (p Prompt) ToMessage(
	messageID string,
	sessionID string,
) Message {
	message := axoncode.UserMessage{
		ID:        messageID,
		SessionID: sessionID,
		Role:      axoncode.UserMessageRoleUser,
		Time: axoncode.UserMessageTime{
			Created: float64(time.Now().UnixMilli()),
		},
	}

	text := p.Text
	textAttachments := []*attachment.Attachment{}
	for _, attachment := range p.Attachments {
		if attachment.Type == "text" {
			textAttachments = append(textAttachments, attachment)
		}
	}
	for i := 0; i < len(textAttachments)-1; i++ {
		for j := i + 1; j < len(textAttachments); j++ {
			if textAttachments[i].StartIndex < textAttachments[j].StartIndex {
				textAttachments[i], textAttachments[j] = textAttachments[j], textAttachments[i]
			}
		}
	}
	for _, att := range textAttachments {
		if source, ok := att.GetTextSource(); ok {
			if att.StartIndex > att.EndIndex || att.EndIndex > len(text) {
				continue
			}
			text = text[:att.StartIndex] + source.Value + text[att.EndIndex:]
		}
	}

	parts := []axoncode.PartUnion{axoncode.TextPart{
		ID:        id.Ascending(id.Part),
		MessageID: messageID,
		SessionID: sessionID,
		Type:      axoncode.TextPartTypeText,
		Text:      text,
	}}
	for _, attachment := range p.Attachments {
		if attachment.Type == "agent" {
			source, _ := attachment.GetAgentSource()
			parts = append(parts, axoncode.AgentPart{
				ID:        id.Ascending(id.Part),
				MessageID: messageID,
				SessionID: sessionID,
				Name:      source.Name,
				Source: axoncode.AgentPartSource{
					Value: attachment.Display,
					Start: int64(attachment.StartIndex),
					End:   int64(attachment.EndIndex),
				},
			})
			continue
		}

		text := axoncode.FilePartSourceText{
			Start: int64(attachment.StartIndex),
			End:   int64(attachment.EndIndex),
			Value: attachment.Display,
		}
		source := &axoncode.FilePartSource{}
		switch attachment.Type {
		case "text":
			continue
		case "file":
			if fileSource, ok := attachment.GetFileSource(); ok {
				source = &axoncode.FilePartSource{
					Text: text,
					Path: fileSource.Path,
					Type: axoncode.FilePartSourceTypeFile,
				}
			}
		case "symbol":
			if symbolSource, ok := attachment.GetSymbolSource(); ok {
				source = &axoncode.FilePartSource{
					Text: text,
					Path: symbolSource.Path,
					Type: axoncode.FilePartSourceTypeSymbol,
					Kind: int64(symbolSource.Kind),
					Name: symbolSource.Name,
					Range: axoncode.SymbolSourceRange{
						Start: axoncode.SymbolSourceRangeStart{
							Line:      float64(symbolSource.Range.Start.Line),
							Character: float64(symbolSource.Range.Start.Char),
						},
						End: axoncode.SymbolSourceRangeEnd{
							Line:      float64(symbolSource.Range.End.Line),
							Character: float64(symbolSource.Range.End.Char),
						},
					},
				}
			}
		}
		parts = append(parts, axoncode.FilePart{
			ID:        id.Ascending(id.Part),
			MessageID: messageID,
			SessionID: sessionID,
			Type:      axoncode.FilePartTypeFile,
			Filename:  attachment.Filename,
			Mime:      attachment.MediaType,
			URL:       attachment.URL,
			Source:    *source,
		})
	}
	return Message{
		Info:  message,
		Parts: parts,
	}
}

func (m Message) ToPrompt() (*Prompt, error) {
	switch m.Info.(type) {
	case axoncode.UserMessage:
		text := ""
		attachments := []*attachment.Attachment{}
		for _, part := range m.Parts {
			switch p := part.(type) {
			case axoncode.TextPart:
				if p.Synthetic {
					continue
				}
				text += p.Text + " "
			case axoncode.AgentPart:
				attachments = append(attachments, &attachment.Attachment{
					ID:         p.ID,
					Type:       "agent",
					Display:    p.Source.Value,
					StartIndex: int(p.Source.Start),
					EndIndex:   int(p.Source.End),
					Source: &attachment.AgentSource{
						Name: p.Name,
					},
				})
			case axoncode.FilePart:
				switch p.Source.Type {
				case "file":
					attachments = append(attachments, &attachment.Attachment{
						ID:         p.ID,
						Type:       "file",
						Display:    p.Source.Text.Value,
						URL:        p.URL,
						Filename:   p.Filename,
						MediaType:  p.Mime,
						StartIndex: int(p.Source.Text.Start),
						EndIndex:   int(p.Source.Text.End),
						Source: &attachment.FileSource{
							Path: p.Source.Path,
							Mime: p.Mime,
						},
					})
				case "symbol":
					r := p.Source.Range.(axoncode.SymbolSourceRange)
					attachments = append(attachments, &attachment.Attachment{
						ID:         p.ID,
						Type:       "symbol",
						Display:    p.Source.Text.Value,
						URL:        p.URL,
						Filename:   p.Filename,
						MediaType:  p.Mime,
						StartIndex: int(p.Source.Text.Start),
						EndIndex:   int(p.Source.Text.End),
						Source: &attachment.SymbolSource{
							Path: p.Source.Path,
							Name: p.Source.Name,
							Kind: int(p.Source.Kind),
							Range: attachment.SymbolRange{
								Start: attachment.Position{
									Line: int(r.Start.Line),
									Char: int(r.Start.Character),
								},
								End: attachment.Position{
									Line: int(r.End.Line),
									Char: int(r.End.Character),
								},
							},
						},
					})
				}
			}
		}
		return &Prompt{
			Text:        text,
			Attachments: attachments,
		}, nil
	}
	return nil, errors.New("unknown message type")
}

func (m Message) ToSessionChatParams() []axoncode.SessionPromptParamsPartUnion {
	parts := []axoncode.SessionPromptParamsPartUnion{}
	for _, part := range m.Parts {
		switch p := part.(type) {
		case axoncode.TextPart:
			parts = append(parts, axoncode.TextPartInputParam{
				ID:        axoncode.F(p.ID),
				Type:      axoncode.F(axoncode.TextPartInputTypeText),
				Text:      axoncode.F(p.Text),
				Synthetic: axoncode.F(p.Synthetic),
				Time: axoncode.F(axoncode.TextPartInputTimeParam{
					Start: axoncode.F(p.Time.Start),
					End:   axoncode.F(p.Time.End),
				}),
			})
		case axoncode.FilePart:
			var source axoncode.FilePartSourceUnionParam
			switch p.Source.Type {
			case "file":
				source = axoncode.FileSourceParam{
					Type: axoncode.F(axoncode.FileSourceTypeFile),
					Path: axoncode.F(p.Source.Path),
					Text: axoncode.F(axoncode.FilePartSourceTextParam{
						Start: axoncode.F(int64(p.Source.Text.Start)),
						End:   axoncode.F(int64(p.Source.Text.End)),
						Value: axoncode.F(p.Source.Text.Value),
					}),
				}
			case "symbol":
				source = axoncode.SymbolSourceParam{
					Type: axoncode.F(axoncode.SymbolSourceTypeSymbol),
					Path: axoncode.F(p.Source.Path),
					Name: axoncode.F(p.Source.Name),
					Kind: axoncode.F(p.Source.Kind),
					Range: axoncode.F(axoncode.SymbolSourceRangeParam{
						Start: axoncode.F(axoncode.SymbolSourceRangeStartParam{
							Line:      axoncode.F(float64(p.Source.Range.(axoncode.SymbolSourceRange).Start.Line)),
							Character: axoncode.F(float64(p.Source.Range.(axoncode.SymbolSourceRange).Start.Character)),
						}),
						End: axoncode.F(axoncode.SymbolSourceRangeEndParam{
							Line:      axoncode.F(float64(p.Source.Range.(axoncode.SymbolSourceRange).End.Line)),
							Character: axoncode.F(float64(p.Source.Range.(axoncode.SymbolSourceRange).End.Character)),
						}),
					}),
					Text: axoncode.F(axoncode.FilePartSourceTextParam{
						Value: axoncode.F(p.Source.Text.Value),
						Start: axoncode.F(p.Source.Text.Start),
						End:   axoncode.F(p.Source.Text.End),
					}),
				}
			}
			parts = append(parts, axoncode.FilePartInputParam{
				ID:       axoncode.F(p.ID),
				Type:     axoncode.F(axoncode.FilePartInputTypeFile),
				Mime:     axoncode.F(p.Mime),
				URL:      axoncode.F(p.URL),
				Filename: axoncode.F(p.Filename),
				Source:   axoncode.F(source),
			})
		case axoncode.AgentPart:
			parts = append(parts, axoncode.AgentPartInputParam{
				ID:   axoncode.F(p.ID),
				Type: axoncode.F(axoncode.AgentPartInputTypeAgent),
				Name: axoncode.F(p.Name),
				Source: axoncode.F(axoncode.AgentPartInputSourceParam{
					Value: axoncode.F(p.Source.Value),
					Start: axoncode.F(p.Source.Start),
					End:   axoncode.F(p.Source.End),
				}),
			})
		}
	}
	return parts
}
