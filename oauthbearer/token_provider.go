package oauthbearer

import (
	"context"

	"github.com/Shopify/sarama"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// TokenProvider is a simple struct that implements sarama.AccessTokenProvider.
// It encapsulates an oauth2.TokenSource which is leveraged for AccessToken retrieval through the
// oauth2 client credentials flow, the token will auto-refresh as necessary.
type TokenProvider struct {
	tokenSource oauth2.TokenSource
}

// NewTokenProvider creates a new sarama.AccessTokenProvider with the provided clientID and clientSecret.
// The provided tokenURL is used to perform the 2 legged client credentials flow.
func NewTokenProvider(clientID, clientSecret, tokenURL string) sarama.AccessTokenProvider {
	cfg := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
	}

	return &TokenProvider{
		tokenSource: cfg.TokenSource(context.Background()),
	}
}

// Token returns a new *sarama.AccessToken or an error as appropriate.
func (t *TokenProvider) Token() (*sarama.AccessToken, error) {
	token, err := t.tokenSource.Token()
	if err != nil {
		return nil, err
	}

	return &sarama.AccessToken{Token: token.AccessToken}, nil
}
