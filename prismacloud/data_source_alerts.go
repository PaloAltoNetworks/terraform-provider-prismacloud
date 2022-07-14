package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/alert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAlerts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlertsRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"time_range": timeRangeSchema("data_source_alerts"),
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Max number of alerts to return.  This uses the v2 version of the API, where the default and max is 10,000.",
				Default:     10000,
			},
			"filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filtering parameters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Param name to filter on",
						},
						"operator": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Operator between the name and value params",
							Default:     "=",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Param value for the filter",
						},
					},
				},
			},
			"sort_by": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Array of sort properties. Append :asc or :desc to the key to sort by ascending or descending order respectively.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Attributes.
			"page_token": {
				Type:        schema.TypeString,
				Description: "The next page token returned",
				Computed:    true,
			},
			"total": totalSchema("alerts"),
			"listing": {
				Type:        schema.TypeList,
				Description: "Alert listing",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert_id": {
							Type:        schema.TypeString,
							Description: "Alert ID",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Alert status",
							Computed:    true,
						},
						"first_seen": {
							Type:        schema.TypeInt,
							Description: "First seen",
							Computed:    true,
						},
						"last_seen": {
							Type:        schema.TypeInt,
							Description: "Last seen",
							Computed:    true,
						},
						"alert_time": {
							Type:        schema.TypeInt,
							Description: "Alert time",
							Computed:    true,
						},
						"event_occurred": {
							Type:        schema.TypeInt,
							Description: "Event occurred",
							Computed:    true,
						},
						"triggered_by": {
							Type:        schema.TypeString,
							Description: "Triggered by",
							Computed:    true,
						},
						"alert_count": {
							Type:        schema.TypeInt,
							Description: "Alert count",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func parseAlertsRequest(d *schema.ResourceData) *alert.Request {
	ans := alert.Request{
		Limit:     d.Get("limit").(int),
		SortBy:    ListToStringSlice(d.Get("sort_by").([]interface{})),
		TimeRange: ParseTimeRange(ResourceDataInterfaceMap(d, "time_range")),
	}

	filters := d.Get("filters").([]interface{})
	ans.Filters = make([]alert.Filter, 0, len(filters))
	for i := range filters {
		f := filters[i].(map[string]interface{})
		ans.Filters = append(ans.Filters, alert.Filter{
			Name:     f["name"].(string),
			Operator: f["operator"].(string),
			Value:    f["value"].(string),
		})
	}

	return &ans
}

func dataSourceAlertsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	req := parseAlertsRequest(d)
	ans, err := alert.List(client, *req)
	if err != nil {
		return err
	}

	d.SetId(client.Url)
	d.Set("page_token", ans.PageToken)
	d.Set("total", ans.Total)

	data := make([]interface{}, 0, len(ans.Data))
	for num, info := range ans.Data {
		// TODO(shinmog) - Remove this workaround when Prisma Cloud fixes their bug.
		//
		// WORKAROUND: Prisma Cloud does not honor the limit for to_now queries, so
		// enforce it here to prevent resource size overruns in Terraform:
		//
		// Error: rpc error: code = ResourceExhausted desc = grpc: received message larger than max (5685945 vs. 4194304)
		//
		// The `total` value is being intentionally left as-is so later on it will be
		// easier to see when they've fixed this on their end.
		if num >= req.Limit {
			break
		}
		item := map[string]interface{}{
			"alert_id":       info.Id,
			"status":         info.Status,
			"first_seen":     info.FirstSeen,
			"last_seen":      info.LastSeen,
			"alert_time":     info.AlertTime,
			"event_occurred": info.EventOccurred,
			"triggered_by":   info.TriggeredBy,
			"alert_count":    info.AlertCount,
		}
		data = append(data, item)
	}
	if err := d.Set("listing", data); err != nil {
		log.Printf("[WARN] Error setting 'data' for %q: %s", d.Id(), err)
	}

	return nil
}
