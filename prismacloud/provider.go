package prismacloud

import (
	"fmt"

	pc "github.com/paloaltonetworks/prisma-cloud-go"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The API URL without the leading protocol",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUD_URL", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Access key ID",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUD_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret key",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUD_PASSWORD", nil),
				Sensitive:   true,
			},
			"customer_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Customer name",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUD_CUSTOMER_NAME", nil),
			},
			"protocol": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The protocol (https or http)",
				DefaultFunc:  schema.EnvDefaultFunc("PRISMACLOUD_PROTOCOL", nil),
				ValidateFunc: validation.StringInSlice([]string{"", "https", "http"}, false),
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "If the port is non-standard for the protocol, the port number to use",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUD_PORT", nil),
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout in seconds for all communications with Prisma Cloud",
				Default:     90,
			},
			"skip_ssl_cert_verification": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Skip SSL certificate verification",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUD_SKIP_SSL_CERT_VERIFICATION", nil),
			},
			"logging": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
				Optional:    true,
				Description: "Logging options for the API connection",
			},
			"disable_reconnect": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Disable reconnecting on JWT expiration",
			},
			"json_web_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "JSON web token to use",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUD_JSON_WEB_TOKEN", nil),
				Sensitive:   true,
			},
			"json_config_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Retrieve the provider configuration from this JSON file",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUD_JSON_CONFIG_FILE", nil),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"prismacloud_account_group":                            dataSourceAccountGroup(),
			"prismacloud_account_groups":                           dataSourceAccountGroups(),
			"prismacloud_alert_rule":                               dataSourceAlertRule(),
			"prismacloud_alert_rules":                              dataSourceAlertRules(),
			"prismacloud_alerts":                                   dataSourceAlerts(),
			"prismacloud_cloud_account":                            dataSourceCloudAccount(),
			"prismacloud_cloud_accounts":                           dataSourceCloudAccounts(),
			"prismacloud_compliance_standard":                      dataSourceComplianceStandard(),
			"prismacloud_compliance_standard_requirement":          dataSourceComplianceStandardRequirement(),
			"prismacloud_compliance_standard_requirement_section":  dataSourceComplianceStandardRequirementSection(),
			"prismacloud_compliance_standard_requirement_sections": dataSourceComplianceStandardRequirementSections(),
			"prismacloud_compliance_standard_requirements":         dataSourceComplianceStandardRequirements(),
			"prismacloud_compliance_standards":                     dataSourceComplianceStandards(),
			"prismacloud_enterprise_settings":                      dataSourceEnterpriseSettings(),
			"prismacloud_integration":                              dataSourceIntegration(),
			"prismacloud_integrations":                             dataSourceIntegrations(),
			"prismacloud_policies":                                 dataSourcePolicies(),
			"prismacloud_policy":                                   dataSourcePolicy(),
			"prismacloud_rql_historic_search":                      dataSourceRqlHistoricSearch(),
			"prismacloud_rql_historic_searches":                    dataSourceRqlHistoricSearches(),
			"prismacloud_org_cloud_account":                        dataSourceOrgCloudAccount(),
			"prismacloud_org_cloud_accounts":                       dataSourceOrgCloudAccounts(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"prismacloud_account_group":                           resourceAccountGroup(),
			"prismacloud_alert_rule":                              resourceAlertRule(),
			"prismacloud_cloud_account":                           resourceCloudAccount(),
			"prismacloud_compliance_standard":                     resourceComplianceStandard(),
			"prismacloud_compliance_standard_requirement":         resourceComplianceStandardRequirement(),
			"prismacloud_compliance_standard_requirement_section": resourceComplianceStandardRequirementSection(),
			"prismacloud_enterprise_settings":                     resourceEnterpriseSettings(),
			"prismacloud_integration":                             resourceIntegration(),
			"prismacloud_policy":                                  resourcePolicy(),
			"prismacloud_rql_search":                              resourceRqlSearch(),
			"prismacloud_saved_search":                            resourceSavedSearch(),
			"prismacloud_user_role":                               resourceUserRole(),
			"prismacloud_org_cloud_account":                       resourceOrgCloudAccount(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	/*
	   An int in Terraform is a Go "int", which can be either 32 or 64bit
	   depending on what the underlying OS is.  A Terraform "schema.TypeInt" is
	   also a Go "int".

	   Timestamps returned Prisma Cloud are 64bit ints.  In addition to this,
	   a Go time.Duration is an int64.

	   Thus, require that the OS is 64bit or bail.

	   If this becomes a problem in the future, then the fix is to go through all
	   resources and anywhere where a timestamp is returned, that needs to be either
	   a float64 or a string, either of which will require lots of casting.
	*/
	is64Bit := uint64(^uintptr(0)) == ^uint64(0)
	if !is64Bit {
		return nil, fmt.Errorf("This provider requires a 64bit OS")
	}

	logSetting := make(map[string]bool)
	logConfig := d.Get("logging").(map[string]interface{})
	for key := range logConfig {
		logSetting[key] = logConfig[key].(bool)
	}

	con := &pc.Client{
		Url:                     d.Get("url").(string),
		Username:                d.Get("username").(string),
		Password:                d.Get("password").(string),
		CustomerName:            d.Get("customer_name").(string),
		Protocol:                d.Get("protocol").(string),
		Port:                    d.Get("port").(int),
		Timeout:                 d.Get("timeout").(int),
		SkipSslCertVerification: d.Get("skip_ssl_cert_verification").(bool),
		DisableReconnect:        d.Get("disable_reconnect").(bool),
		JsonWebToken:            d.Get("json_web_token").(string),
		Logging:                 logSetting,
	}

	err := con.Initialize(d.Get("json_config_file").(string))
	return con, err
}
