package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/ip-address"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrustedLoginIp() *schema.Resource {
	return &schema.Resource{
		CreateContext: createTrustedLoginIp,
		ReadContext:   readTrustedLoginIp,
		UpdateContext: updateTrustedLoginIp,
		DeleteContext: deleteTrustedLoginIp,

		Schema: map[string]*schema.Schema{
			"trusted_login_ip_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Login IP Allow List ID",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique name for CIDR (IP addresses) allow list",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of CIDR (IP addresses) allow list",
			},
			"cidr": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of CIDRs to Allow List for login access. You can include from 1 to 10 CIDRs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"last_modified_ts": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Timestamp for last modification of CIDR block list",
			},
		},
	}
}

func createTrustedLoginIp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	obj := parseTrustedLoginIp(d, "")

	if err := ip_address.Create(client, obj); err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := ip_address.Identify(client, obj.Name)
		return err
	})

	id, err := ip_address.Identify(client, obj.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := ip_address.Get(client, id)
		return err
	})

	d.SetId(id)
	return readTrustedLoginIp(ctx, d, meta)
}

func readTrustedLoginIp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	print("Step1: Read")
	id := d.Id()

	obj, err := ip_address.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	saveTrustedLoginIpList(d, obj)

	return nil
}

func updateTrustedLoginIp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	print("Step1: Update")
	id := d.Get("trusted_login_ip_id").(string)
	obj := parseTrustedLoginIp(d, id)

	if err := ip_address.Update(client, obj); err != nil {
		return diag.FromErr(err)
	}

	return readTrustedLoginIp(ctx, d, meta)
}

func saveTrustedLoginIpList(d *schema.ResourceData, obj ip_address.LoginIpAllow) {
	var err error

	d.Set("trusted_login_ip_id", obj.Id)
	d.Set("name", obj.Name)
	d.Set("description", obj.Description)
	d.Set("last_modified_ts", obj.LastModifiedTs)
	if err = d.Set("cidr", obj.Cidr); err != nil {
		log.Printf("[WARN] Error setting 'cidr' field for %q: %s", d.Id(), err)
	}
}

func parseTrustedLoginIp(d *schema.ResourceData, id string) ip_address.LoginIpAllow {
	cidr := d.Get("cidr")
	return ip_address.LoginIpAllow{
		Id:             id,
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		Cidr:           SetToStringSlice(cidr.(*schema.Set)),
		LastModifiedTs: d.Get("last_modified_ts").(int),
	}
}

func deleteTrustedLoginIp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()

	if err := ip_address.Delete(client, id); err != nil {
		if err != pc.ObjectNotFoundError {
			return diag.FromErr(err)
		}
	}

	d.SetId("")
	return nil
}
