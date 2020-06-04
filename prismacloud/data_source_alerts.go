package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/alert"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceAlerts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlertsRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"time_range": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The time range spec",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"absolute": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "An absolute time range",
							MaxItems:    1,
							ConflictsWith: []string{
								"time_range.relative",
								"time_range.to_now",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Start time",
									},
									"end": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "End time",
									},
								},
							},
						},
						"relative": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Relative time range",
							MaxItems:    1,
							ConflictsWith: []string{
								"time_range.absolute",
								"time_range.to_now",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"amount": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The time number",
									},
									"unit": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The time unit",
										ValidateFunc: validation.StringInSlice(
											[]string{
												alert.TimeHour,
												alert.TimeDay,
												alert.TimeWeek,
												alert.TimeMonth,
												alert.TimeYear,
											},
											false,
										),
									},
								},
							},
						},
						"to_now": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "From some time in the past to now",
							MaxItems:    1,
							ConflictsWith: []string{
								"time_range.absolute",
								"time_range.relative",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"unit": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The time unit",
										ValidateFunc: validation.StringInSlice(
											[]string{
												alert.TimeLogin,
												alert.TimeEpoch,
												alert.TimeDay,
												alert.TimeWeek,
												alert.TimeMonth,
												alert.TimeYear,
											},
											false,
										),
									},
								},
							},
						},
					},
				},
			},
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
			"data": {
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
		Limit:  d.Get("limit").(int),
		SortBy: ListToStringSlice(d.Get("sort_by").([]interface{})),
	}

	tr := (d.Get("time_range").([]interface{})[0]).(map[string]interface{})
	if atr := ToInterfaceMap(tr, "absolute"); len(atr) != 0 {
		ans.TimeRange.Value = alert.Absolute{
			Start: atr["start"].(int),
			End:   atr["end"].(int),
		}
	} else if rtr := ToInterfaceMap(tr, "relative"); len(rtr) != 0 {
		ans.TimeRange.Value = alert.Relative{
			Amount: rtr["amount"].(int),
			Unit:   rtr["unit"].(string),
		}
	} else if tntr := ToInterfaceMap(tr, "to_now"); len(tntr) != 0 {
		ans.TimeRange.Value = alert.ToNow{
			Unit: tntr["unit"].(string),
		}
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
	for _, info := range ans.Data {
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
	if err := d.Set("data", data); err != nil {
		log.Printf("[WARN] Error setting 'data' for %q: %s", d.Id(), err)
	}

	return nil
}
