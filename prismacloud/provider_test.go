package prismacloud

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account/org"
	"github.com/paloaltonetworks/prisma-cloud-go/settings/enterprise"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	PrismacloudJsonConfigFileEnvVar = "PRISMACLOUD_JSON_CONFIG_FILE"
)

var (
	testAccProviders                   map[string]*schema.Provider
	testAccProvider                    *schema.Provider
	originalEnterpriseSettings         *enterprise.Config
	cloudAccounts                      map[string][]string
	sessionTimeoutOrig, sessionTimeout int
)

func init() {
	var err error

	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"prismacloud": testAccProvider,
	}

	client := &pc.Client{}
	if err = client.Initialize(os.Getenv(PrismacloudJsonConfigFileEnvVar)); err == nil {
		if o, err := enterprise.Get(client); err == nil {
			originalEnterpriseSettings = &o
		}
	}
}

func cloudAccountFromEnv(style, label, name string, groups []string) (string, error) {
	ctDesc := map[string]string{
		account.TypeAws:     "AWS",
		account.TypeAzure:   "Azure",
		account.TypeGcp:     "GCP",
		account.TypeAlibaba: "Alibaba",
	}

	evMap := map[string][]string{
		account.TypeAws: []string{
			"PRISMACLOUD_AWS_ACCOUNT_ID",
			"PRISMACLOUD_AWS_EXTERNAL_ID",
			"PRISMACLOUD_AWS_ROLE_ARN",
		},
		account.TypeAzure: []string{
			"PRISMACLOUD_AZURE_ACCOUNT_ID",
			"PRISMACLOUD_AZURE_CLIENT_ID",
			"PRISMACLOUD_AZURE_KEY",
			"PRISMACLOUD_AZURE_MONITOR_FLOW_LOGS",
			"PRISMACLOUD_AZURE_TENANT_ID",
			"PRISMACLOUD_AZURE_SERVICE_PRINCIPAL_ID",
		},
		account.TypeGcp: []string{
			"PRISMACLOUD_GCP_ACCOUNT_ID",
			"PRISMACLOUD_GCP_COMPRESSION_ENABLED",
			"PRISMACLOUD_GCP_DATAFLOW_ENABLED_PROJECT",
			"PRISMACLOUD_GCP_FLOW_LOG_STORAGE_BUCKET",
			"PRISMACLOUD_GCP_CREDENTIALS_FILE",
		},
		account.TypeAlibaba: []string{
			"PRISMACLOUD_ALIBABA_ACCOUNT_ID",
			"PRISMACLOUD_ALIBABA_RAM_ARN",
		},
		org.TypeOci: []string{
			"PRISMACLOUD_OCI_ACCOUNT_ID",
			"PRISMACLOUD_OCI_GROUP_NAME",
			"PRISMACLOUD_OCI_HOME_REGION",
			"PRISMACLOUD_OCI_USER_NAME",
			"PRISMACLOUD_OCI_USER_OCID",
			"PRISMACLOUD_OCI_POLICY_NAME",
		},
	}

	vlist, ok := evMap[style]
	if !ok {
		return "", fmt.Errorf("unknown style %q", style)
	}

	missing := make([]string, 0, len(vlist))
	for _, key := range vlist {
		if _, ok := os.LookupEnv(key); !ok {
			missing = append(missing, key)
		}
	}

	if len(missing) != 0 {
		msg := strings.Join(missing, ", ")
		return "", fmt.Errorf("%s test requires these environment variables: %s", ctDesc[style], msg)
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`
resource "prismacloud_cloud_account" %q {
    %s {`, label, style))
	switch style {
	case account.TypeAws:
		buf.WriteString(fmt.Sprintf(`
        account_id = %q
        enabled = false
        external_id = %q
        name = %q
        role_arn = %q`,
			os.Getenv("PRISMACLOUD_AWS_ACCOUNT_ID"),
			os.Getenv("PRISMACLOUD_AWS_EXTERNAL_ID"),
			name,
			os.Getenv("PRISMACLOUD_AWS_ROLE_ARN"),
		))
	case account.TypeAzure:
		mfl, _ := strconv.ParseBool(os.Getenv("PRISMACLOUD_AZURE_MONITOR_FLOW_LOGS"))

		buf.WriteString(fmt.Sprintf(`
        account_id = %q
        enabled = false
        name = %q
        client_id = %q
        key = %q
        monitor_flow_logs = %t
        tenant_id = %q
        service_principal_id = %q`,
			os.Getenv("PRISMACLOUD_AZURE_ACCOUNT_ID"),
			name,
			os.Getenv("PRISMACLOUD_AZURE_CLIENT_ID"),
			os.Getenv("PRISMACLOUD_AZURE_KEY"),
			mfl,
			os.Getenv("PRISMACLOUD_AZURE_TENANT_ID"),
			os.Getenv("PRISMACLOUD_AZURE_SERVICE_PRINCIPAL_ID"),
		))
	case account.TypeGcp:
		ce, _ := strconv.ParseBool(os.Getenv("PRISMACLOUD_GCP_COMPRESSION_ENABLED"))

		buf.WriteString(fmt.Sprintf(`
        account_id = %q
        enabled = false
        name = %q
        compression_enabled = %t
        dataflow_enabled_project = %q
        flow_log_storage_bucket = %q
        credentials_json = file(%q)`,
			os.Getenv("PRISMACLOUD_GCP_ACCOUNT_ID"),
			name,
			ce,
			os.Getenv("PRISMACLOUD_GCP_DATAFLOW_ENABLED_PROJECT"),
			os.Getenv("PRISMACLOUD_GCP_FLOW_LOG_STORAGE_BUCKET"),
			os.Getenv("PRISMACLOUD_GCP_CREDENTIALS_FILE"),
		))
	case account.TypeAlibaba:
		buf.WriteString(fmt.Sprintf(`
        account_id = %q
        name = %q
        ram_arn = %q
        enabled = false`,
			os.Getenv("PRISMACLOUD_ALIBABA_ACCOUNT_ID"),
			name,
			os.Getenv("PRISMACLOUD_ALIBABA_RAM_ARN"),
		))
	case org.TypeOci:

	}

	buf.WriteString(`
        group_ids = [`)
	for _, g := range groups {
		buf.WriteString(fmt.Sprintf(`
            %s,`, g))
	}
	buf.WriteString(`
        ]
    }
}`)

	return buf.String(), nil
}
func cloudAccountOrgFromEnv(style, label, name string, groups []string) (string, error) {
	ctDescorg := map[string]string{
		org.TypeAwsOrg:   "AWSORG",
		org.TypeAzureOrg: "AZUREORG",
		org.TypeGcpOrg:   "GCPORg",
		org.TypeOci:      "OCI",
	}

	evMaporg := map[string][]string{
		org.TypeAwsOrg: []string{
			"PRISMACLOUD_AWSORG_ACCOUNT_ID",
			"PRISMACLOUD_AWSORG_EXTERNAL_ID",
			"PRISMACLOUD_AWSORG_ROLE_ARN",
			"PRISMACLOUD_AWSORG_MEMBER_ROLE_NAME",
			"PRISMACLOUD_AWSORG_MEMBER_EXTERNAL_ID",
		},
		org.TypeAzureOrg: []string{
			"PRISMACLOUD_AZUREORG_ACCOUNT_ID",
			"PRISMACLOUD_AZUREORG_CLIENT_ID",
			"PRISMACLOUD_AZUREORG_KEY",
			"PRISMACLOUD_AZUREORG_MONITOR_FLOW_LOGS",
			"PRISMACLOUD_AZUREORG_TENANT_ID",
			"PRISMACLOUD_AZUREORG_SERVICE_PRINCIPAL_ID",
		},
		org.TypeGcpOrg: []string{
			"PRISMACLOUD_GCPORG_ACCOUNT_ID",
			"PRISMACLOUD_GCPORG_COMPRESSION_ENABLED",
			"PRISMACLOUD_GCPORG_DATAFLOW_ENABLED_PROJECT",
			"PRISMACLOUD_GCPORG_FLOW_LOG_STORAGE_BUCKET",
			"PRISMACLOUD_GCPORG_CREDENTIALS_FILE",
			"PRISMACLOUD_GCPORG_ORGANIZATION_NAME",
		},
		org.TypeOci: []string{
			"PRISMACLOUD_OCI_ACCOUNT_ID",
			"PRISMACLOUD_OCI_GROUP_NAME",
			"PRISMACLOUD_OCI_HOME_REGION",
			"PRISMACLOUD_OCI_USER_NAME",
			"PRISMACLOUD_OCI_USER_OCID",
			"PRISMACLOUD_OCI_POLICY_NAME",
		},
	}

	vlistorg, ok := evMaporg[style]
	if !ok {
		return "", fmt.Errorf("unknown style %q", style)
	}

	missingorg := make([]string, 0, len(vlistorg))
	for _, key := range vlistorg {
		if _, ok := os.LookupEnv(key); !ok {
			missingorg = append(missingorg, key)
		}
	}

	if len(missingorg) != 0 {
		msg := strings.Join(missingorg, ", ")
		return "", fmt.Errorf("%s test requires these environment variables: %s", ctDescorg[style], msg)
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`
resource "prismacloud_org_cloud_account" %q {
    %s {`, label, style))
	switch style {
	case org.TypeAwsOrg:
		buf.WriteString(fmt.Sprintf(`
        account_id = %q
        enabled = false
        external_id = %q
        name = %q
        role_arn = %q
		member_role_name = %q
		member_external_id = %q`,
			os.Getenv("PRISMACLOUD_AWSORG_ACCOUNT_ID"),
			os.Getenv("PRISMACLOUD_AWSORG_EXTERNAL_ID"),
			name,
			os.Getenv("PRISMACLOUD_AWSORG_ROLE_ARN"),
			os.Getenv("PRISMACLOUD_AWSORG_MEMBER_ROLE_NAME"),
			os.Getenv("PRISMACLOUD_AWSORG_MEMBER_EXTERNAL_ID"),
		))

	case org.TypeAzureOrg:
		mfl, _ := strconv.ParseBool(os.Getenv("PRISMACLOUD_AZUREORG_MONITOR_FLOW_LOGS"))
		buf.WriteString(fmt.Sprintf(`
        account_id = %q
        enabled = false
        name = %q
        client_id = %q
        key = %q
        monitor_flow_logs = %t
        tenant_id = %q
        service_principal_id = %q
		monitor_flow_logs = %q`,
			os.Getenv("PRISMACLOUD_AZUREORG_ACCOUNT_ID"),
			name,
			os.Getenv("PRISMACLOUD_AZUREORG_CLIENT_ID"),
			os.Getenv("PRISMACLOUD_AZUREORG_KEY"),
			mfl,
			os.Getenv("PRISMACLOUD_AZUREORG_TENANT_ID"),
			os.Getenv("PRISMACLOUD_AZUREORG_SERVICE_PRINCIPAL_ID"),
			os.Getenv("PRISMACLOUD_AZUREORG_MONITOR_FLOW_LOGS"),
		))
	case org.TypeGcpOrg:
		ce, _ := strconv.ParseBool(os.Getenv("PRISMACLOUD_GCPORG_COMPRESSION_ENABLED"))

		buf.WriteString(fmt.Sprintf(`
        account_id = %q
        enabled = false
        name = %q
        compression_enabled = %t
        dataflow_enabled_project = %q
        flow_log_storage_bucket = %q
        credentials_json = file(%q)
		organization_name = %q `,
			os.Getenv("PRISMACLOUD_GCPORG_ACCOUNT_ID"),
			name,
			ce,
			os.Getenv("PRISMACLOUD_GCPORG_DATAFLOW_ENABLED_PROJECT"),
			os.Getenv("PRISMACLOUD_GCPORG_FLOW_LOG_STORAGE_BUCKET"),
			os.Getenv("PRISMACLOUD_GCPORG_CREDENTIALS_FILE"),
			os.Getenv("PRISMACLOUD_GCPORG_ORGANIZATION_NAME"),
		))
	case org.TypeOci:
		buf.WriteString(fmt.Sprintf(`
        account_id = %q
		enabled = false
		group_name = %q
		home_region = %q
		policy_name = %q
		user_name = %q
		user_ocid = %q
        name = %q
		enabled = false`,
			os.Getenv("PRISMACLOUD_OCI_ACCOUNT_ID"),
			os.Getenv("PRISMACLOUD_OCI_GROUP_NAME"),
			os.Getenv("PRISMACLOUD_OCI_HOME_REGION"),
			os.Getenv("PRISMACLOUD_OCI_POLICY_NAME"),
			os.Getenv("PRISMACLOUD_OCI_USER_NAME"),
			os.Getenv("PRISMACLOUD_OCI_USER_OCID"),
			name,
		))

	}

	buf.WriteString(`
        group_ids = [`)
	if style != org.TypeOci {
		for _, g := range groups {
			buf.WriteString(fmt.Sprintf(`
            %s,`, g))
		}
	}
	buf.WriteString(`
        ]
    }
}`)

	return buf.String(), nil
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv(PrismacloudJsonConfigFileEnvVar) == "" {
		t.Fatalf("%s must be set for acceptance tests", PrismacloudJsonConfigFileEnvVar)
	}
}
