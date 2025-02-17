package utils

import (
	"context"
	"strings"

	"github.com/oauth2-proxy/oauth2-proxy/v7/providers"
)

// query parameter for the default rule
// for requests generated by oauth2 proxy and meant for oauth2 proxy (self-redirects), we'll inject following query parameter containing providerId
const DefaultProviderIDQueryParam = "provider-id"

type contextKey string

const (
	providerKey   contextKey = "provider"
	providerIDKey contextKey = "providerId"
)

// extarcts providerId stored in a context
// returns empty string if providerId not found
func ProviderIDFromContext(ctx context.Context) string {
	t, ok := ctx.Value(providerIDKey).(string)
	if !ok {
		return ""
	}
	return t
}

// stores providerId in the context's key value pair
func AppendProviderIDToContext(ctx context.Context, providerID string) context.Context {
	return context.WithValue(ctx, providerIDKey, providerID)
}

// injects provider-id in a url
func InjectProviderID(pid string, uri string) string {
	// Replace `ProviderID` with the provider id for Redirect Url
	uri = strings.Replace(uri, "ProviderID", pid, 1)
	return uri
	// Disable injects provider-id completely

	// if empty provider-id, no need to inject it since empty provider-id is loaded when it's not found in the request
	// if pid == "" {
	// 	return uri
	// }
	// Replace `ProviderID` with the provider id for Redirect Url
	// uri = strings.Replace(uri, "ProviderID", pid, 1)
	// u, err := url.Parse(uri)
	// if err != nil {
	// 	return uri
	// }
	// q := u.Query()
	// q.Set(DefaultProviderIDQueryParam, pid)
	// u.RawQuery = q.Encode()
	// return u.String()
}

// extracts provider stored in a context
// returns nil if provider not found
func ProviderFromContext(ctx context.Context) providers.Provider {
	t, ok := ctx.Value(providerKey).(providers.Provider)
	if !ok {
		return nil
	}
	return t
}

// stores provider in the context's key value pair
func AppendProviderToContext(ctx context.Context, p providers.Provider) context.Context {
	return context.WithValue(ctx, providerKey, p)
}
