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

func TestFindFilesWithOptionalParams(t *testing.T) {
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
	_, err := client.Find.Files(context.TODO(), axoncode.FindFilesParams{
		Query:     axoncode.F("query"),
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

func TestFindSymbolsWithOptionalParams(t *testing.T) {
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
	_, err := client.Find.Symbols(context.TODO(), axoncode.FindSymbolsParams{
		Query:     axoncode.F("query"),
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

func TestFindTextWithOptionalParams(t *testing.T) {
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
	_, err := client.Find.Text(context.TODO(), axoncode.FindTextParams{
		Pattern:   axoncode.F("pattern"),
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
