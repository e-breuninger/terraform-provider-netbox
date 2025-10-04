package netbox

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/status"
	"github.com/stretchr/testify/assert"
)

func TestValidClientWithAllData(t *testing.T) {
	config := Config{
		APIToken:  "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL: "https://localhost:8080",
	}

	client, err := config.Client()
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestURLMissingSchemaShouldWork(t *testing.T) {
	config := Config{
		APIToken:  "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL: "localhost:8080",
	}

	client, err := config.Client()
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestURLMaleformedUrlShouldFail(t *testing.T) {
	config := Config{
		APIToken:  "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL: "xyz:/localhost:8080",
	}

	_, err := config.Client()
	assert.Error(t, err)
}

func TestURLMissingPortShouldWork(t *testing.T) {
	config := Config{
		APIToken:  "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL: "http://localhost",
	}

	client, err := config.Client()
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestURLMissingAccessKey(t *testing.T) {
	config := Config{
		APIToken:  "",
		ServerURL: "http://localhost",
	}

	_, err := config.Client()
	assert.Error(t, err)
}

func TestClientWithCustomHeaders(t *testing.T) {
	config := Config{
		APIToken:  "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL: "https://localhost:8080",
		Headers: map[string]interface{}{
			"X-Custom-Header": "test-value",
			"X-Another-Header": 123,
		},
	}

	client, err := config.Client()
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestClientWithInsecureHTTPS(t *testing.T) {
	config := Config{
		APIToken:          "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL:         "https://localhost:8080",
		AllowInsecureHTTPS: true,
	}

	client, err := config.Client()
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestClientWithRequestTimeout(t *testing.T) {
	config := Config{
		APIToken:      "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL:     "https://localhost:8080",
		RequestTimeout: 30,
	}

	client, err := config.Client()
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestClientWithStripTrailingSlashes(t *testing.T) {
	config := Config{
		APIToken:                    "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL:                   "https://localhost:8080/",
		StripTrailingSlashesFromURL: true,
	}

	client, err := config.Client()
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestAdditionalHeadersSet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vals, ok := r.Header["Hello"]

		assert.True(t, ok)
		assert.Len(t, vals, 1)
		assert.Equal(t, vals[0], "World!")
	}))
	defer ts.Close()

	config := Config{
		APIToken:  "07b12b765127747e4afd56cb531b7bf9c61f3c30",
		ServerURL: ts.URL,
		Headers: map[string]interface{}{
			"Hello": "World!",
		},
	}

	client, err := config.Client()
	assert.NoError(t, err)

	req := status.NewStatusListParams()
	client.Status.StatusList(req, nil)
}

func TestCustomHeaderTransportRoundTrip(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vals, ok := r.Header["Custom-Header"]

		assert.True(t, ok)
		assert.Len(t, vals, 1)
		assert.Equal(t, vals[0], "test-value")
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// Create a custom header transport
	transport := customHeaderTransport{
		original: http.DefaultTransport,
		headers: map[string]interface{}{
			"Custom-Header": "test-value",
		},
	}

	req, _ := http.NewRequest("GET", ts.URL, nil)
	resp, err := transport.RoundTrip(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

/* TODO
func TestInvalidHttpsCertificate(t *testing.T) {}
*/
