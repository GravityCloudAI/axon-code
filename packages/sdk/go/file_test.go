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

func TestFileListWithOptionalParams(t *testing.T) {
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
	_, err := client.File.List(context.TODO(), axoncode.FileListParams{
		Path:      axoncode.F("path"),
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

func TestFileReadWithOptionalParams(t *testing.T) {
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
	_, err := client.File.Read(context.TODO(), axoncode.FileReadParams{
		Path:      axoncode.F("path"),
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

func TestFileStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.File.Status(context.TODO(), axoncode.FileStatusParams{
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
