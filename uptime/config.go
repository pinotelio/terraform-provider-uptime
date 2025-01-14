package uptime

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/uptime-com/uptime-client-go"
)

// Config defines configuration options for the Uptime.com client
type Config struct {
	// Uptime.com API token
	Token            string
	RateMilliseconds int
}

const badCredentials = `

No credentials found for Uptime.com provider.
Please provide an API token in the provider block or as an environment
variable.
`

func (c *Config) Client() (*uptime.Client, error) {
	if c.Token == "" {
		return nil, fmt.Errorf(badCredentials)
	}

	var httpClient *http.Client
	httpClient = http.DefaultClient
	httpClient.Transport = logging.NewTransport("Uptime.com", http.DefaultTransport)

	config := &uptime.Config{
		HTTPClient:       httpClient,
		Token:            c.Token,
		UserAgent:        fmt.Sprintf("(%s %s) Terraform", runtime.GOOS, runtime.GOARCH),
		RateMilliseconds: c.RateMilliseconds,
	}

	client, err := uptime.NewClient(config)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] Uptime.com client configured")

	return client, nil
}
