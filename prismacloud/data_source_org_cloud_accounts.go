package prismacloud

import (
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account/org"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

func dataSourceOrgCloudAccounts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOrgCloudAccountsRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("cloud accounts"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of accounts",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account name",
						},
						"cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud type",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account ID",
						},
					},
				},
			},
		},
	}
}

func dataSourceOrgCloudAccountsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	items, err := org.Names(client)
	if err != nil {
		return err
	}

	d.SetId("cloud_accounts")
	d.Set("total", len(items))

	list := make([]interface{}, 0, len(items))
	for _, i := range items {
		list = append(list, map[string]interface{}{
			"name":       i.Name,
			"cloud_type": i.CloudType,
			"account_id": i.AccountId,
		})
	}

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'accounts' field for %q: %s", d.Id(), err)
	}

	return nil
}
