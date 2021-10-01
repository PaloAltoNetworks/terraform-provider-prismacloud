package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/integration"
	"log"
	"strings"
)

func resourceIntegration() *schema.Resource {
	return &schema.Resource{
		Create: createIntegration,
		Read:   readIntegration,
		Update: updateIntegration,
		Delete: deleteIntegration,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"integration_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Integration ID",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the integration",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description",
			},
			"integration_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Integration type",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
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
			"jira_password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Jira account password",
			},
			"jira_username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Jira account Username",
			},

			"oauth_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Jira account password",
			},
			"approve": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Jira account Username",
			},

			"oauth_callback": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Jira account password",
			},

			"integration_config": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "Integration configuration, the values depend on the integration type",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"queue_url": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Queue URL you used when you configured Prisma Cloud in Amazon SQS",
						},
						"login": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "(Qualys/ServiceNow) Login",
						},
						"base_url": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Qualys Security Operations Center server API URL (without \"http(s)\")",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "(Qualys/ServiceNow) Password",
						},
						"host_url": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ServiceNow URL/Jira Url",
						},
						"secret_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Jira Secret Key",
						},
						"oauth_token": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Jira Auth token",
						},
						"consumer_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Jira consumer key",
						},
						"tables": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Key/value pairs that identify the ServiceNow module tables with which to integrate (e.g. - incident, sn_si_incident, or em_event)",
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ServiceNow release version",
						},
						"url": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Webhook URL",
						},
						"headers": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Webhook headers",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Header name",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Header value",
									},
									"secure": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Secure",
									},
									"read_only": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Read only",
									},
								},
							},
						},
						"auth_token": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "PagerDuty authentication token for the event collector",
						},
						"integration_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "PagerDuty integration key",
						},
						"source_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "GCP Source ID for Google CSCC integration",
						},
						"org_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "GCP Organization ID for Google CSCC integration",
						},
					},
				},
			},
		},
	}
}

func parseIntegration(d *schema.ResourceData, id string, c pc.PrismaCloudClient) integration.Integration {
	ic := ResourceDataInterfaceMap(d, "integration_config")
	var secretKey string
	var oauthToken string
	log.Printf("%d secretkey", secretKey)
	if d.Get("integration_type") == "jira"{
		var authjiraurl integration.AuthUrl
		authjiraurl.HostUrl = ic["host_url"].(string)
		authjiraurl.ConsumerKey = ic["consumer_key"].(string)
		authurlresponse, _ := integration.JiraAuthurl(c, authjiraurl)

		var seckeyjira  integration.SecretKeyJira
		tokenfromUrl := strings.Split(authurlresponse, "=")[1]
		token := tokenfromUrl[:len(tokenfromUrl)-1]
		seckeyjira.OauthToken = token
		seckeyjira.JiraUserName = d.Get("jira_username").(string)
		seckeyjira.JiraPassword = d.Get("jira_password").(string)
		secretKey, _ = integration.JiraSecretKey(c, seckeyjira, ic["host_url"].(string) )

		var oauthtoken  integration.OauthTokenJira
		oauthtoken.AuthenticationUrl = authurlresponse[1: len(authurlresponse) -1]
		oauthtoken.HostUrl = ic["host_url"].(string)
		oauthtoken.ConsumerKey = ic["consumer_key"].(string)
		oauthtoken.SecretKey = secretKey
		oauthtoken.TmpToken = token
		tokenresponse, _ := integration.JiraOauthToken(c, oauthtoken)
		oauthToken = tokenresponse[1:len(tokenresponse)-1]
	}

	var tables []map[string]bool
	var headers []integration.Header

	if ic["tables"] != nil && len(ic["tables"].(map[string]interface{})) > 0 {
		tlist := ic["tables"].(map[string]interface{})
		tables = make([]map[string]bool, 0, len(tlist))
		for key, value := range tlist {
			tables = append(tables, map[string]bool{key: value.(bool)})
		}
	}

	if ic["headers"] != nil && len(ic["headers"].([]interface{})) > 0 {
		hlist := ic["headers"].([]interface{})
		headers = make([]integration.Header, 0, len(hlist))
		for i := range hlist {
			hdr := hlist[i].(map[string]interface{})
			headers = append(headers, integration.Header{
				Key:      hdr["key"].(string),
				Value:    hdr["value"].(string),
				Secure:   hdr["secure"].(bool),
				ReadOnly: hdr["read_only"].(bool),
			})
		}
	}

	return integration.Integration{
		Id:              id,
		Name:            d.Get("name").(string),
		Description:     d.Get("description").(string),
		IntegrationType: d.Get("integration_type").(string),
		IntegrationConfig: integration.IntegrationConfig{
			QueueUrl:       ic["queue_url"].(string),
			Login:          ic["login"].(string),
			BaseUrl:        ic["base_url"].(string),
			Password:       ic["password"].(string),
			HostUrl:        ic["host_url"].(string),
			Tables:         tables,
			Version:        ic["version"].(string),
			Url:            ic["url"].(string),
			Headers:        headers,
			AuthToken:      ic["auth_token"].(string),
			IntegrationKey: ic["integration_key"].(string),
			SourceId:       ic["source_id"].(string),
			OrgId:          ic["org_id"].(string),
			ConsumerKey:    ic["consumer_key"].(string),
			SecretKey:      secretKey,
			OauthToken:    	oauthToken,
		},
		Enabled: d.Get("enabled").(bool),
	}
}

func saveIntegration(d *schema.ResourceData, o integration.Integration) {
	var err error

	d.Set("integration_id", o.Id)
	d.Set("name", o.Name)
	d.Set("description", o.Description)
	d.Set("integration_type", o.IntegrationType)
	d.Set("enabled", o.Enabled)
	d.Set("created_by", o.CreatedBy)
	d.Set("created_ts", o.CreatedTs)
	d.Set("last_modified_by", o.LastModifiedBy)
	d.Set("last_modified_ts", o.LastModifiedTs)
	d.Set("status", o.Status)
	d.Set("valid", o.Valid)

	reason := map[string]interface{}{
		"last_updated": o.Reason.LastUpdated,
		"error_type":   o.Reason.ErrorType,
		"message":      o.Reason.Message,
		"details":      nil,
	}
	if o.Reason.Details != nil {
		reason["details"] = []interface{}{map[string]interface{}{
			"status_code": o.Reason.Details.StatusCode,
			"subject":     o.Reason.Details.Subject,
			"message":     o.Reason.Details.Message,
		}}
	}
	if err = d.Set("reason", []interface{}{reason}); err != nil {
		log.Printf("[WARN] Error setting 'reason' for %s: %s", d.Id(), err)
	}

	ic := map[string]interface{}{
		"queue_url":       o.IntegrationConfig.QueueUrl,
		"login":           o.IntegrationConfig.Login,
		"base_url":        o.IntegrationConfig.BaseUrl,
		"password":        o.IntegrationConfig.Password,
		"host_url":        o.IntegrationConfig.HostUrl,
		"tables":          nil,
		"version":         o.IntegrationConfig.Version,
		"url":             o.IntegrationConfig.Url,
		"headers":         nil,
		"auth_token":      o.IntegrationConfig.AuthToken,
		"integration_key": o.IntegrationConfig.IntegrationKey,
		"source_id":       o.IntegrationConfig.SourceId,
		"org_id":          o.IntegrationConfig.OrgId,
		"consumer_key":    o.IntegrationConfig.ConsumerKey,
		"secret_key":      o.IntegrationConfig.SecretKey,
		"oauth_token":     o.IntegrationConfig.OauthToken,
	}

	if len(o.IntegrationConfig.Tables) != 0 {
		tables := make(map[string]interface{})
		for _, t := range o.IntegrationConfig.Tables {
			for key, value := range t {
				tables[key] = value
			}
		}
		ic["tables"] = tables
	}
	if len(o.IntegrationConfig.Headers) != 0 {
		headers := make([]interface{}, 0, len(o.IntegrationConfig.Headers))
		for _, h := range o.IntegrationConfig.Headers {
			headers = append(headers, map[string]interface{}{
				"key":       h.Key,
				"value":     h.Value,
				"secure":    h.Secure,
				"read_only": h.ReadOnly,
			})
		}
		ic["headers"] = headers
	}
	if err = d.Set("integration_config", []interface{}{ic}); err != nil {
		log.Printf("[WARN] Error setting 'integration_config' for %s: %s", d.Id(), err)
	}
}

func createIntegration(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	o := parseIntegration(d, "", client)
	if err := integration.Create(client, o); err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := integration.Identify(client, o.Name)
		return err
	})

	id, err := integration.Identify(client, o.Name)
	if err != nil {
		return err
	}
	PollApiUntilSuccess(func() error {
		_, err := integration.Get(client, id)
		return err
	})

	d.SetId(id)
	return readIntegration(d, meta)
}

func readIntegration(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	o, err := integration.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	saveIntegration(d, o)

	return nil
}

func updateIntegration(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()
	o := parseIntegration(d, id, client)
	if err := integration.Update(client, o); err != nil {
		return err
	}

	return readIntegration(d, meta)
}

func deleteIntegration(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	err := integration.Delete(client, id)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return err
		}
	}

	d.SetId("")
	return nil
}