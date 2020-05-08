package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account/group"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAccountGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAccountGroupsRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("account groups"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of accounts",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account group ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account group name",
						},
						"accounts": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Associated cloud accounts",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Associated cloud account ID",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Associated cloud account name",
									},
									"account_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Associated cloud account type",
									},
								},
							},
						},
						"alert_rules": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Singly associated alert rules which cannot exist in the system without the account group",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alert_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The alert ID",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Alert name",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceAccountGroupsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	items, err := group.List(client)
	if err != nil {
		return err
	}

	d.SetId("account_groups")
	d.Set("total", len(items))

	list := make([]interface{}, 0, len(items))
	for _, i := range items {
		acts := make([]interface{}, 0, len(i.Accounts))
		for _, j := range i.Accounts {
			acts = append(acts, map[string]interface{}{
				"account_id":   j.Id,
				"name":         j.Name,
				"account_type": j.Type,
			})
		}

		rules := make([]interface{}, 0, len(i.AlertRules))
		for _, k := range i.AlertRules {
			rules = append(rules, map[string]interface{}{
				"alert_id": k.Id,
				"name":     k.Name,
			})
		}

		list = append(list, map[string]interface{}{
			"group_id":    i.Id,
			"name":        i.Name,
			"accounts":    acts,
			"alert_rules": rules,
		})
	}

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
