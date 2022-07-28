package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/integration"
	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIntegration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIntegrationRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"integration_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Integration ID",
				AtLeastOneOf: []string{"integration_id", "name"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Name of the integration",
				AtLeastOneOf: []string{"integration_id", "name"},
			},
			"integration_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Integration type",
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
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Created by",
			},
			"created_ts": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Created on",
			},
			"last_modified_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last modified by",
			},
			"last_modified_ts": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Last modified timestamp",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status",
			},
			"valid": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Valid",
			},
			"reason": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Model for the integration status details",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"last_updated": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Last updated",
						},
						"error_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Error type",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Message",
						},
						"details": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Model for message details",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status_code": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Status code",
									},
									"subject": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Subject",
									},
									"message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Internationalization key",
									},
								},
							},
						},
					},
				},
			},
			"integration_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Integration configuration, the values depend on the integration type",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"queue_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Queue URL you used when you configured Prisma Cloud in Amazon SQS or Azure Service Bus Queue",
						},
						"more_info": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "true = specific IAM credentials are specified for SQS queue access",
						},
						"login": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "(Qualys/ServiceNow) Login",
						},
						"base_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Qualys Security Operations Center server API URL (without \"http(s)\")",
						},
						"password": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "(Qualys/ServiceNow) Password",
						},
						"host_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ServiceNow/Demisto URL",
						},
						"user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Snow Flake Username",
						},
						"pipe_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Snow Flake Pipename",
						},
						"private_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Snow Flake private key",
						},
						"pass_phrase": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Snow Flake Pass phrase ",
						},
						"staging_integration_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Amazon s3 ID for snowflake integration",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Okta Domain",
						},
						"api_token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Okta API Token",
						},
						"api_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Demisto API key",
						},
						"access_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Tenable access key",
						},
						"secret_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Tenable Secret key",
						},
						"tables": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Key/value pairs that identify the ServiceNow module tables with which to integrate (e.g. - incident, sn_si_incident, or em_event)",
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ServiceNow release version",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Webhook URL or Splunk HTTP event collector URL",
						},
						"headers": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Webhook headers",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Header name",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Header value",
									},
									"secure": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Secure",
									},
									"read_only": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Read only",
									},
								},
							},
						},
						"auth_token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PagerDuty authentication token for the event collector",
						},
						"integration_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PagerDuty/Splunk integration key",
						},
						"webhook_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Webhook url for slack integration ",
						},
						"source_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GCP Source ID for Google CSCC integration",
						},
						"org_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GCP Organization ID for Google CSCC integration",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS/Azure account ID for AWS Security Hub/Azure Service Bus Queue integration",
						},
						"connection_string": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Connection string for azure service bus queue integration",
						},
						"roll_up_interval": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "File Roll Up Time in minutes for Snowflake integration and AWS S3 integration",
						},
						"regions": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "AWS regions for AWS Security Hub integration",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "AWS region name",
									},
									"api_identifier": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "AWS region code",
									},
									"cloud_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud Type",
									},
								},
							},
						},
						"s3_uri": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS S3 URI for Amazon S3 integration",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS region for Amazon S3 integration",
						},
						"role_arn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS role ARN for Amazon S3 integration",
						},
						"external_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS external ID for Amazon S3 integration",
						},
						"source_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source type for splunk integration",
						},
					},
				},
			},
		},
	}
}

func dataSourceIntegrationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	var err error
	id := d.Get("integration_id").(string)

	prismaIdRequired := true
	integrationType := d.Get("integration_type").(string)
	if stringInSlice(integrationType, integration.InboundIntegrations) {
		prismaIdRequired = false
	}

	if id == "" {
		name := d.Get("name").(string)
		id, err = integration.Identify(client, name, prismaIdRequired)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return diag.FromErr(err)
		}
	}

	o, err := integration.Get(client, id, prismaIdRequired)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.SetId(o.Id)
	saveIntegration(d, o)

	return nil
}
