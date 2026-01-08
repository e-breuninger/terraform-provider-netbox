package netbox

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// deviceRenderConfigResponse represents the response from the render-config API
type deviceRenderConfigResponse struct {
	ConfigTemplate *struct {
		ID          int64  `json:"id"`
		URL         string `json:"url"`
		Display     string `json:"display"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"configtemplate"`
	Content string `json:"content"`
}

func dataSourceNetboxDeviceRenderConfig() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxDeviceRenderConfigRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):Render the configuration template assigned to a device using the device's config context.`,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID of the device to render configuration for.",
			},
			"content": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The rendered configuration content.",
			},
			"config_template_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the config template that was used for rendering.",
			},
			"config_template_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the config template that was used for rendering.",
			},
		},
	}
}

func dataSourceNetboxDeviceRenderConfigRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	deviceID := d.Get("device_id").(int)

	// Build the URL for the render-config endpoint
	url := fmt.Sprintf("%s/api/dcim/devices/%d/render-config/", api.httpClient.BaseURL, deviceID)

	// Create a POST request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", api.httpClient.APIToken))

	// Execute the request
	resp, err := api.httpClient.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("error calling render-config API: %w", err)
	}
	defer resp.Body.Close()

	// Check for errors
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("render-config API returned status %d: %s", resp.StatusCode, string(body))
	}

	// Parse the response
	var result deviceRenderConfigResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("error decoding render-config response: %w", err)
	}

	d.SetId(strconv.Itoa(deviceID))
	d.Set("content", result.Content)

	if result.ConfigTemplate != nil {
		d.Set("config_template_id", result.ConfigTemplate.ID)
		d.Set("config_template_name", result.ConfigTemplate.Name)
	}

	return nil
}
