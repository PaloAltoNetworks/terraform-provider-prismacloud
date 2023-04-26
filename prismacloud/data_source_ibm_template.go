package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account-v2/ibmTemplate"
	"golang.org/x/net/context"
)

func dataSourceIbmTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmTemplateRead,

		Schema: map[string]*schema.Schema{
			"account_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IBM account type",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"account",
					},
					false,
				),
			},
			"file_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "File name to store ibm template",
			},
		},
	}
}

func dataSourceIbmTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	req := ibmTemplate.IbmTemplateReq{
		AccountType: d.Get("account_type").(string),
		FileName:    d.Get("file_name").(string),
	}

	err := ibmTemplate.GetIbmTemplate(client, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get("account_type").(string))
	d.Set("account_type", d.Get("account_type").(string))
	d.Set("file_name", d.Get("file_name").(string))

	return nil
}
