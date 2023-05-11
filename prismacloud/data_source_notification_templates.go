package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	notification_template "github.com/paloaltonetworks/prisma-cloud-go/notification-template"
	"golang.org/x/net/context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

func dataSourceNotificationTemplates() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNotificationTemplatesRead,

		Schema: map[string]*schema.Schema{
			//Output
			"total": totalSchema("notification_templates_list"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Notification Templates",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "id",
						},
						"integration_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "integrationId",
						},
						"created_ts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "createdTs",
						},
						"last_modified_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "lastModifiedBy",
						},
						"last_modified_ts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "lastModifiedTs",
						},
						"integration_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "integrationType",
						},
						"created_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "createdBy",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "name",
						},
						"integration_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "integration Name",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"customer_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "customerId",
						},
						"module": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "module",
						},
						"template_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "templateType",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "enabled",
						},
					},
				},
			},
		},
	}
}

func dataSourceNotificationTemplatesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	items, err := notification_template.List(client)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("notification_templates_list")
	d.Set("total", len(items))
	list := make([]interface{}, 0, len(items))
	for _, i := range items {
		list = append(list, map[string]interface{}{
			"id":               i.Id,
			"name":             i.Name,
			"integration_id":   i.IntegrationId,
			"integration_type": i.IntegrationType,
			"integration_name": i.IntegrationName,
			"customer_id":      i.CustomerId,
			"module":           i.Module,
			"template_type":    i.TemplateType,
			"last_modified_by": i.LastModifiedBy,
			"enabled":          i.Enabled,
			"created_by":       i.CreatedBy,
			"created_ts":       i.CreatedTs,
		})
	}
	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}
	return nil
}
