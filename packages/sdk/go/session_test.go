// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package axoncode_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/sst/axoncode-sdk-go"
	"github.com/sst/axoncode-sdk-go/internal/testutil"
	"github.com/sst/axoncode-sdk-go/option"
)

func TestSessionNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.New(context.TODO(), axoncode.SessionNewParams{
		Directory: axoncode.F("directory"),
		ParentID:  axoncode.F("parentID"),
		Title:     axoncode.F("title"),
	})
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Update(
		context.TODO(),
		"id",
		axoncode.SessionUpdateParams{
			Directory: axoncode.F("directory"),
			Title:     axoncode.F("title"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.List(context.TODO(), axoncode.SessionListParams{
		Directory: axoncode.F("directory"),
	})
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionDeleteWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Delete(
		context.TODO(),
		"id",
		axoncode.SessionDeleteParams{
			Directory: axoncode.F("directory"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionAbortWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Abort(
		context.TODO(),
		"id",
		axoncode.SessionAbortParams{
			Directory: axoncode.F("directory"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionChildrenWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Children(
		context.TODO(),
		"id",
		axoncode.SessionChildrenParams{
			Directory: axoncode.F("directory"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionCommandWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Command(
		context.TODO(),
		"id",
		axoncode.SessionCommandParams{
			Arguments: axoncode.F("arguments"),
			Command:   axoncode.F("command"),
			Directory: axoncode.F("directory"),
			Agent:     axoncode.F("agent"),
			MessageID: axoncode.F("msgJ!"),
			Model:     axoncode.F("model"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Get(
		context.TODO(),
		"id",
		axoncode.SessionGetParams{
			Directory: axoncode.F("directory"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionInitWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Init(
		context.TODO(),
		"id",
		axoncode.SessionInitParams{
			MessageID:  axoncode.F("messageID"),
			ModelID:    axoncode.F("modelID"),
			ProviderID: axoncode.F("providerID"),
			Directory:  axoncode.F("directory"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionMessageWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Message(
		context.TODO(),
		"id",
		"messageID",
		axoncode.SessionMessageParams{
			Directory: axoncode.F("directory"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionMessagesWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Messages(
		context.TODO(),
		"id",
		axoncode.SessionMessagesParams{
			Directory: axoncode.F("directory"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionPromptWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Prompt(
		context.TODO(),
		"id",
		axoncode.SessionPromptParams{
			Parts: axoncode.F([]axoncode.SessionPromptParamsPartUnion{axoncode.TextPartInputParam{
				Text:      axoncode.F("text"),
				Type:      axoncode.F(axoncode.TextPartInputTypeText),
				ID:        axoncode.F("id"),
				Synthetic: axoncode.F(true),
				Time: axoncode.F(axoncode.TextPartInputTimeParam{
					Start: axoncode.F(0.000000),
					End:   axoncode.F(0.000000),
				}),
			}}),
			Directory: axoncode.F("directory"),
			Agent:     axoncode.F("agent"),
			MessageID: axoncode.F("msgJ!"),
			Model: axoncode.F(axoncode.SessionPromptParamsModel{
				ModelID:    axoncode.F("modelID"),
				ProviderID: axoncode.F("providerID"),
			}),
			System: axoncode.F("system"),
			Tools: axoncode.F(map[string]bool{
				"foo": true,
			}),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionRevertWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Revert(
		context.TODO(),
		"id",
		axoncode.SessionRevertParams{
			MessageID: axoncode.F("msgJ!"),
			Directory: axoncode.F("directory"),
			PartID:    axoncode.F("prtJ!"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionShareWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Share(
		context.TODO(),
		"id",
		axoncode.SessionShareParams{
			Directory: axoncode.F("directory"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionShellWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Shell(
		context.TODO(),
		"id",
		axoncode.SessionShellParams{
			Agent:     axoncode.F("agent"),
			Command:   axoncode.F("command"),
			Directory: axoncode.F("directory"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionSummarizeWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Summarize(
		context.TODO(),
		"id",
		axoncode.SessionSummarizeParams{
			ModelID:    axoncode.F("modelID"),
			ProviderID: axoncode.F("providerID"),
			Directory:  axoncode.F("directory"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionUnrevertWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Unrevert(
		context.TODO(),
		"id",
		axoncode.SessionUnrevertParams{
			Directory: axoncode.F("directory"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionUnshareWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := axoncode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Session.Unshare(
		context.TODO(),
		"id",
		axoncode.SessionUnshareParams{
			Directory: axoncode.F("directory"),
		},
	)
	if err != nil {
		var apierr *axoncode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
