// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	statusGetResponse, err := client.Status.Get(context.TODO())
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", statusGetResponse.Message)
}
