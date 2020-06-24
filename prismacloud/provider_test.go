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
	"github.com/paloaltonetworks/prisma-cloud-go/settings/enterprise"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	PrismacloudJsonConfigFileEnvVar = "PRISMACLOUD_JSON_CONFIG_FILE"
)

var (
	testAccProviders                   map[string]terraform.ResourceProvider
	testAccProvider                    *schema.Provider
	originalEnterpriseSettings         *enterprise.Config
	cloudAccounts                      map[string][]string
	sessionTimeoutOrig, sessionTimeout int
)

func init() {
	var err error

	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
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

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv(PrismacloudJsonConfigFileEnvVar) == "" {
		t.Fatalf("%s must be set for acceptance tests", PrismacloudJsonConfigFileEnvVar)
	}
}
