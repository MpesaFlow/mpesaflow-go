// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mpesaflow_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/mpesaflow-go"
	"github.com/stainless-sdks/mpesaflow-go/internal/testutil"
	"github.com/stainless-sdks/mpesaflow-go/option"
)

func TestUsage(t *testing.T) {
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
	app, err := client.Apps.New(context.TODO(), mpesaflow.AppNewParams{})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", app.ApplicationID)
}
