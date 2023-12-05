package connector

import (
	"log"

	"github.com/flashkoef/go-ct-rest-api/config"
	"github.com/labd/commercetools-go-sdk/platform"
	"golang.org/x/oauth2/clientcredentials"
)

type Connector struct {
	conf *config.Config
}

func New() *Connector {
	return &Connector{conf: config.New()}
}

func (c *Connector) GetProjectClient() *platform.ByProjectKeyRequestBuilder {
	client := c.createClient()

	return client.WithProjectKey(c.conf.Ctp.ProjectKey)
}

func (c *Connector) createClient() *platform.Client {
	client, err := platform.NewClient(&platform.ClientConfig{
		URL: "https://api.europe-west1.gcp.commercetools.com",
		Credentials: &clientcredentials.Config{
			TokenURL:     "https://auth.europe-west1.gcp.commercetools.com/oauth/token",
			ClientID:     c.conf.Ctp.ClientID,
			ClientSecret: c.conf.Ctp.ClientSecret,
			Scopes:       []string{c.conf.Ctp.Scopes},
		},
	})

	if err != nil {
		log.Fatalf("error while creating ctp client %s", err)
	}

	return client
}
