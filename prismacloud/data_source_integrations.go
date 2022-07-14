package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/integration"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIntegrations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIntegrationsRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("all integrations"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of all integrations",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"integration_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Integration ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Integration name",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description",
						},
						"integration_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Integration type",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enabled",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status",
						},
						"valid": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Valid",
						},
					},
				},
			},
		},
	}
}

func dataSourceIntegrationsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	outboundIntegrations, err := integration.List(client, "", true)
	if err != nil {
		return err
	}

	inboundIntegrations, err := integration.List(client, "", false)
	if err != nil {
		return err
	}

	allIntegrations := append(outboundIntegrations, inboundIntegrations...)
	d.SetId("integrations")
	d.Set("total", len(allIntegrations))

	listing := make([]interface{}, 0, len(allIntegrations))
	for _, o := range allIntegrations {
		listing = append(listing, map[string]interface{}{
			"integration_id":   o.Id,
			"name":             o.Name,
			"description":      o.Description,
			"integration_type": o.IntegrationType,
			"enabled":          o.Enabled,
			"status":           o.Status,
			"valid":            o.Valid,
		})
	}

	if err := d.Set("listing", listing); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
