package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/notification-template"
	"golang.org/x/net/context"
)

func dataSourceNotificationTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNotificationTemplateRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
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
			"template_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Template Config",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"basic_config": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: getConfigSchema(),
							},
						},
						"open": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: getConfigSchema(),
							},
						},
						"resolved": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: getConfigSchema(),
							},
						},
						"dismissed": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: getConfigSchema(),
							},
						},
						"snoozed": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: getConfigSchema(),
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceNotificationTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var err error
	client := meta.(*pc.Client)
	id := d.Get("id").(string)
	if id == "" {
		return nil
	}
	obj, err := notification_template.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	d.SetId(id)
	saveNotificationTemplate(d, obj)
	return nil
}
