package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/status"
	"github.com/stretchr/testify/assert"
)

func TestValidClientWithAllData(t *testing.T) {
	config := &Config{
		APIToken:  "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL: "https://localhost:8080",
	}

	legacyClient, err := NewLegacyClient(config)
	assert.NotNil(t, legacyClient)
	assert.NoError(t, err)

	client, err := NewClient(config)
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestURLMissingSchemaShouldWork(t *testing.T) {
	config := &Config{
		APIToken:  "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL: "localhost:8080",
	}

	legacyClient, err := NewLegacyClient(config)
	assert.NotNil(t, legacyClient)
	assert.NoError(t, err)

	client, err := NewClient(config)
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestURLMalformedUrlShouldFail(t *testing.T) {
	config := &Config{
		APIToken:  "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL: "xyz:/localhost:8080",
	}

	_, err := NewLegacyClient(config)
	assert.Error(t, err)
}

func TestURLMissingPortShouldWork(t *testing.T) {
	config := &Config{
		APIToken:  "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL: "http://localhost",
	}

	legacyClient, err := NewLegacyClient(config)
	assert.NotNil(t, legacyClient)
	assert.NoError(t, err)

	client, err := NewClient(config)
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestURLMissingAccessKey(t *testing.T) {
	config := &Config{
		APIToken:  "",
		ServerURL: "http://localhost",
	}

	_, err := NewLegacyClient(config)
	assert.Error(t, err)

	_, err = NewClient(config)
	assert.Error(t, err)
}

func TestAdditionalHeadersSet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vals, ok := r.Header["Hello"]

		assert.True(t, ok)
		assert.Len(t, vals, 1)
		assert.Equal(t, vals[0], "World!")
	}))
	defer ts.Close()

	config := &Config{
		APIToken:  "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL: ts.URL,
		Headers: map[string]interface{}{
			"Hello": "World!",
		},
	}

	legacyClient, err := NewLegacyClient(config)
	assert.NoError(t, err)

	req := status.NewStatusListParams()
	legacyClient.Status.StatusList(req, nil)

	client, err := NewClient(config)
	assert.NoError(t, err)

	ctx := context.Background()
	client.StatusAPI.StatusRetrieve(ctx).Execute()
}

/* TODO
func TestInvalidHttpsCertificate(t *testing.T) {}
*/
