package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/alert/rule"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAlertRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlertRuleRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"policy_scan_config_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Policy scan config ID",
				AtLeastOneOf: []string{"name", "policy_scan_config_id"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Rule/Scan name",
				AtLeastOneOf: []string{"name", "policy_scan_config_id"},
			},

			// Output.
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enabled",
			},
			"scan_all": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Scan all policies",
			},
			"policies": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of specific policies to scan",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"policy_labels": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Policy labels",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"excluded_policies": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of policies to exclude from scan",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"last_modified_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Last modified on",
			},
			"last_modified_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last modified by",
			},
			"allow_auto_remediate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Allow auto-remediation",
			},
			"delay_notification_ms": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Delay notifications by the specified milliseconds",
			},
			"notify_on_open": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Include open alerts in notification",
			},
			"notify_on_snoozed": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Include snoozed alerts in notification",
			},
			"notify_on_dismissed": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Include dismissed alerts in notification",
			},
			"notify_on_resolved": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Include resolved alerts in notification",
			},
			"owner": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Owner",
			},
			"notification_channels": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of notification channels",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"open_alerts_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Open alerts count",
			},
			"read_only": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Read only",
			},
			"target": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Model for the target filter",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_groups": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of account groups",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"excluded_accounts": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of excluded accounts",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"regions": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of regions",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tags": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of TargetTag objects",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resource tag target",
									},
									"values": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "List of values for resource tag key",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"notification_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of data for notifications to third-party tools",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alert rule notification config ID",
						},
						"frequency": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Frequency",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Scan enabled",
						},
						"recipients": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of unique email addresses to notify",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"detailed_report": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Provide CSV detailed report",
						},
						"with_compression": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Compress detailed report",
						},
						"include_remediation": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Include remediation in detailed report",
						},
						"last_updated": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Last updated",
						},
						"last_sent_ts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Time of last notification in milliseconds",
						},
						"config_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Config type",
						},
						"template_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template ID",
						},
						"timezone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timezone ID",
						},
						"day_of_month": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Day of month",
						},
						"r_rule_schedule": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "R rule schedule",
						},
						"frequency_from_r_rule": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Frequency from R rule",
						},
						"hour_of_day": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Hour of day",
						},
						"days_of_week": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Days of week",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"day": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Day",
									},
									"offset": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Offset",
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

func dataSourceAlertRuleRead(d *schema.ResourceData, meta interface{}) error {
	var err error
	client := meta.(*pc.Client)

	id := d.Get("policy_scan_config_id").(string)

	if id == "" {
		name := d.Get("name").(string)
		id, err = rule.Identify(client, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return err
		}
	}

	obj, err := rule.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(id)
	saveAlertRule(d, obj)

	return nil
}
