package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/trusted-alert-ip"
	"golang.org/x/net/context"
	"log"
)

func resourceTrustedAlertIp() *schema.Resource {
	return &schema.Resource{
		CreateContext: createTrustedAlertIp,
		ReadContext:   readTrustedAlertIp,
		UpdateContext: updateTrustedAlertIp,
		DeleteContext: deleteTrustedAlertIp,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Trusted alert ip ID",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Trusted alert ip name",
			},
			"cidrs": {
				Type:        schema.TypeSet, //data source of both working fine with list
				Optional:    true,
				Description: "CIDRs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "CIDR",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description",
						},
						"uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "UUID",
						},
						"created_on": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Created on",
						},
					},
				},
			},
			"cidr_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Associated cloud account type",
			},
		},
	}
}

func parseTrustedAlertIp(d *schema.ResourceData, id string) trustedalertip.TrustedAlertIP {
	ans := trustedalertip.TrustedAlertIP{
		UUID:      id,
		Name:      d.Get("name").(string),
		CidrCount: d.Get("cidr_count").(int),
	}
	cidrs := d.Get("cidrs").(*schema.Set).List()
	ans.CIDRS = make([]trustedalertip.CIDRS, 0, len(cidrs))
	for _, tlex := range cidrs {
		tle := tlex.(map[string]interface{})
		ans.CIDRS = append(ans.CIDRS, trustedalertip.CIDRS{
			CIDR:        tle["cidr"].(string),
			UUID:        tle["uuid"].(string),
			Description: tle["description"].(string),
			CreatedOn:   tle["created_on"].(int),
		})
	}
	return ans
}

func saveTrustedAlertIp(d *schema.ResourceData, obj trustedalertip.TrustedAlertIP) {
	d.Set("uuid", obj.UUID)
	d.Set("name", obj.Name)
	d.Set("cidr_count", obj.CidrCount)

	cidrs := make([]interface{}, 0, len(obj.CIDRS))
	if len(obj.CIDRS) != 0 {
		for _, cidr := range obj.CIDRS {
			cidrs = append(cidrs, map[string]interface{}{
				"cidr":        cidr.CIDR,
				"created_on":  cidr.CreatedOn,
				"description": cidr.Description,
				"uuid":        cidr.UUID,
			})
		}
	}
	if err := d.Set("cidrs", cidrs); err != nil {
		log.Printf("[WARN] Error setting 'cidrs' for %q: %s", d.Id(), err)
	}
}

func createTrustedAlertIp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	obj := parseTrustedAlertIp(d, "")

	var id string
	var id1 string
	id1, _ = trustedalertip.Identify(client, obj.Name)
	if id1 == "" {
		if _, err := trustedalertip.Create(client, obj); err != nil {
			return diag.FromErr(err)
		}
	}

	PollApiUntilSuccess(func() error {
		id2, err := trustedalertip.Identify(client, obj.Name)
		id = id2
		return err
	})
	for _, o := range obj.CIDRS {
		_, err := trustedalertip.CreateCIDR(client, o, id)
		if err == pc.OverlappingCIDRError { //change here
			log.Printf("OverlappingCIDRError : %s", id)
		}
	}

	PollApiUntilSuccess(func() error {
		_, err := trustedalertip.Get(client, id)
		return err
	})

	d.SetId(id)
	return readTrustedAlertIp(ctx, d, meta)
}

func readTrustedAlertIp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()

	obj, err := trustedalertip.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	saveTrustedAlertIp(d, obj)

	return nil
}

func updateTrustedAlertIp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	obj := trustedalertip.TrustedAlertIP{
		UUID:      d.Get("uuid").(string),
		Name:      d.Get("name").(string),
		CidrCount: d.Get("cidr_count").(int),
	}
	cidrs := d.Get("cidrs").(*schema.Set).List()
	obj.CIDRS = make([]trustedalertip.CIDRS, 0, len(cidrs))
	for _, tlex := range cidrs {
		tle := tlex.(map[string]interface{})
		obj.CIDRS = append(obj.CIDRS, trustedalertip.CIDRS{
			CIDR:        tle["cidr"].(string),
			UUID:        tle["uuid"].(string),
			Description: tle["description"].(string),
			CreatedOn:   tle["created_on"].(int),
		})
	}
	client := meta.(*pc.Client)

	var id string
	var id1 string
	id1, _ = trustedalertip.Identify(client, obj.Name)
	if id1 == "" {
		if _, err := trustedalertip.Create(client, obj); err != nil {
			return diag.FromErr(err)
		}
	}
	PollApiUntilSuccess(func() error {
		id2, err := trustedalertip.Identify(client, obj.Name)
		id = id2
		return err
	})

	listing, _ := trustedalertip.Get(client, id)
	get_api_all_ips := listing.CIDRS
	var ips_to_delete = Intersection(get_api_all_ips, obj.CIDRS)

	for _, ip_uuid := range ips_to_delete {
		if _, err := trustedalertip.DeleteCIDRFromTrustedAlertIp(client, id, ip_uuid); err != nil {
			return diag.FromErr(err)
		}
	}

	for _, o := range obj.CIDRS {
		_, err := trustedalertip.CreateCIDR(client, o, id)
		if err == pc.OverlappingCIDRError {
			if _, err := trustedalertip.UpdateCIDR(client, o, id, o.UUID); err != nil {
				return diag.FromErr(err)
			}
		}
	}
	return readTrustedAlertIp(ctx, d, meta)
}

func deleteTrustedAlertIp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()

	obj := parseTrustedAlertIp(d, d.Id())

	for _, o := range obj.CIDRS {
		if err := trustedalertip.Delete(client, id, o.UUID); err != nil {
			if err != pc.ObjectNotFoundError {
				return diag.FromErr(err)
			}
		}
	}

	d.SetId("")
	return nil
}

func Intersection(a, b []trustedalertip.CIDRS) (c []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item.UUID] = true
	}

	for _, item := range a {
		if _, ok := m[item.UUID]; !ok {
			c = append(c, item.UUID)
		}
	}
	return c
}
