package prismacloud

import (
	"fmt"
	"testing"

	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account/org"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccOrgDsCloudAccountAws(t *testing.T) {
	grp := fmt.Sprintf("tf%s", acctest.RandString(6))
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	conf, err := cloudAccountOrgFromEnv(org.TypeAwsOrg, "x", name, []string{"prismacloud_account_group.x.group_id"})
	if err != nil {
		t.Skip(fmt.Sprintf("%s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOrgDsCloudAccountConfig(org.TypeAwsOrg, grp, conf),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", "account_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.account_id", org.TypeAwsOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.enabled", org.TypeAwsOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.external_id", org.TypeAwsOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.name", org.TypeAwsOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.role_arn", org.TypeAwsOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.member_role_name", org.TypeAwsOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.member_external_id", org.TypeAwsOrg)),
				),
			},
		},
	})
}

func TestAccOrgDsCloudAccountAzure(t *testing.T) {
	grp := fmt.Sprintf("tf%s", acctest.RandString(6))
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	conf, err := cloudAccountOrgFromEnv(org.TypeAzureOrg, "x", name, []string{"prismacloud_account_group.x.group_id"})
	if err != nil {
		t.Skip(fmt.Sprintf("%s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOrgDsCloudAccountConfig(org.TypeAzureOrg, grp, conf),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", "account_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.account_id", org.TypeAzureOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.enabled", org.TypeAzureOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.name", org.TypeAzureOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.client_id", org.TypeAzureOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.key", org.TypeAzureOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.monitor_flow_logs", org.TypeAzureOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.tenant_id", org.TypeAzureOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.service_principal_id", org.TypeAzureOrg)),
				),
			},
		},
	})
}

func TestAccOrgDsCloudAccountGcp(t *testing.T) {
	grp := fmt.Sprintf("tf%s", acctest.RandString(6))
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	conf, err := cloudAccountOrgFromEnv(org.TypeGcpOrg, "x", name, []string{"prismacloud_account_group.x.group_id"})
	if err != nil {
		t.Skip(fmt.Sprintf("%s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOrgDsCloudAccountConfig(org.TypeGcpOrg, grp, conf),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", "account_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.account_id", org.TypeGcpOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.enabled", org.TypeGcpOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.name", org.TypeGcpOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.compression_enabled", org.TypeGcpOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.flow_log_storage_bucket", org.TypeGcpOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.credentials_json", org.TypeGcpOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.organization_name", org.TypeGcpOrg)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.dataflow_enabled_project", org.TypeGcpOrg)),
				),
			},
		},
	})
}

func TestAccOrgDsCloudAccountOci(t *testing.T) {
	grp := fmt.Sprintf("tf%s", acctest.RandString(6))
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	conf, err := cloudAccountOrgFromEnv(org.TypeOci, "x", name, nil)
	if err != nil {
		t.Skip(fmt.Sprintf("%s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOrgDsCloudAccountConfig(org.TypeOci, grp, conf),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", "account_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.account_id", org.TypeOci)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.name", org.TypeOci)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.enabled", org.TypeOci)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.group_name", org.TypeOci)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.home_region", org.TypeOci)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.user_name", org.TypeOci)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.user_ocid", org.TypeOci)),
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_account.test", fmt.Sprintf("%s.0.policy_name", org.TypeOci)),
				),
			},
		},
	})
}

func testAccOrgDsCloudAccountConfig(style, grp, conf string) string {
	return fmt.Sprintf(`
data "prismacloud_org_cloud_account" "test" {
    cloud_type = %q
    name = prismacloud_org_cloud_account.x.%s.0.name
}

resource "prismacloud_account_group" "x" {
    name = %q
    description = "acctest for org cloud account data source"
}

%s
`, style, style, grp, conf)
}
