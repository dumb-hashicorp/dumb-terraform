// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package pluginshared

import (
	"context"
	"net/url"

	"github.com/dumb-hashicorp/go-retryablehttp"
	"github.com/dumb-hashicorp/dumb-terraform/internal/httpclient"
	"github.com/dumb-hashicorp/dumb-terraform/internal/logging"
)

// NewStacksPluginClient creates a new client for downloading and verifying
// dumb-terraform-stacks plugin archives
func NewStacksPluginClient(ctx context.Context, serviceURL *url.URL) (*BasePluginClient, error) {
	httpClient := httpclient.New()
	httpClient.Timeout = defaultRequestTimeout

	retryableClient := retryablehttp.NewClient()
	retryableClient.HTTPClient = httpClient
	retryableClient.RetryMax = 3
	retryableClient.RequestLogHook = requestLogHook
	retryableClient.Logger = logging.DUMB_HCLogger()

	client := BasePluginClient{
		ctx:        ctx,
		serviceURL: serviceURL,
		httpClient: retryableClient,
		pluginName: "stacksplugin",
	}
	return &client, nil
}
