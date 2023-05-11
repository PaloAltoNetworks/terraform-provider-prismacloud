package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/ip-address"
	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoginIpStatus() *schema.Resource {
	return &schema.Resource{
		CreateContext: createLoginIpStatus,
		UpdateContext: updateLoginIpStatus,
		ReadContext:   readLoginIpStatus,
		DeleteContext: deleteLoginIpStatus,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Set to true when the ip login is enabled",
			},
		},
	}
}

func createLoginIpStatus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	req := ip_address.LoginIpAllowStatus{
		Enabled: d.Get("enabled").(bool),
	}

	if err := ip_address.LoginIpStatusUpdate(client, req); err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := ip_address.GetLoginIpStatus(client)
		return err
	})

	d.SetId("login ip status")

	return readLoginIpStatus(ctx, d, meta)
}

func updateLoginIpStatus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	req := ip_address.LoginIpAllowStatus{
		Enabled: d.Get("enabled").(bool),
	}

	if err := ip_address.LoginIpStatusUpdate(client, req); err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := ip_address.GetLoginIpStatus(client)
		return err
	})

	d.SetId("login ip status")

	return readLoginIpStatus(ctx, d, meta)
}

//done
func readLoginIpStatus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	info, err := ip_address.GetLoginIpStatus(client)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("enabled", info.Enabled)

	return nil
}

//done
func deleteLoginIpStatus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
