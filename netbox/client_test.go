package netbox

import (
	"testing"

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

/* TODO
func TestInvalidHttpsCertificate(t *testing.T) {}
*/
