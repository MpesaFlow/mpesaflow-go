// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mpesaflow_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/mpesaflow-go"
	"github.com/stainless-sdks/mpesaflow-go/internal/testutil"
	"github.com/stainless-sdks/mpesaflow-go/option"
)

func TestAppAPIKeyNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mpesaflow.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppAPIKey("My App API Key"),
		option.WithRootAPIKey("My Root API Key"),
	)
	_, err := client.Apps.APIKeys.New(
		context.TODO(),
		"appId",
		mpesaflow.AppAPIKeyNewParams{
			KeyName: mpesaflow.F("keyName"),
		},
	)
	if err != nil {
		var apierr *mpesaflow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppAPIKeyListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mpesaflow.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppAPIKey("My App API Key"),
		option.WithRootAPIKey("My Root API Key"),
	)
	_, err := client.Apps.APIKeys.List(
		context.TODO(),
		"appId",
		mpesaflow.AppAPIKeyListParams{
			EndingBefore:  mpesaflow.F("ending_before"),
			Limit:         mpesaflow.F(int64(0)),
			StartingAfter: mpesaflow.F("starting_after"),
		},
	)
	if err != nil {
		var apierr *mpesaflow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppAPIKeyDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mpesaflow.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppAPIKey("My App API Key"),
		option.WithRootAPIKey("My Root API Key"),
	)
	_, err := client.Apps.APIKeys.Delete(
		context.TODO(),
		"appId",
		"apiKeyId",
	)
	if err != nil {
		var apierr *mpesaflow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
