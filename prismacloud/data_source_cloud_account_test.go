package prismacloud

import (
	"fmt"
	"testing"

	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDsCloudAccountAws(t *testing.T) {
	grp := fmt.Sprintf("tf%s", acctest.RandString(6))
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	conf, err := cloudAccountFromEnv(account.TypeAws, "x", name, []string{"prismacloud_account_group.x.group_id"})
	if err != nil {
		t.Skip(fmt.Sprintf("%s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsCloudAccountConfig(account.TypeAws, grp, conf),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", "account_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.account_id", account.TypeAws)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.enabled", account.TypeAws)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.external_id", account.TypeAws)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.name", account.TypeAws)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.role_arn", account.TypeAws)),
				),
			},
		},
	})
}

func TestAccDsCloudAccountAzure(t *testing.T) {
	grp := fmt.Sprintf("tf%s", acctest.RandString(6))
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	conf, err := cloudAccountFromEnv(account.TypeAzure, "x", name, []string{"prismacloud_account_group.x.group_id"})
	if err != nil {
		t.Skip(fmt.Sprintf("%s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsCloudAccountConfig(account.TypeAzure, grp, conf),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", "account_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.account_id", account.TypeAzure)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.enabled", account.TypeAzure)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.name", account.TypeAzure)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.client_id", account.TypeAzure)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.key", account.TypeAzure)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.monitor_flow_logs", account.TypeAzure)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.tenant_id", account.TypeAzure)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.service_principal_id", account.TypeAzure)),
				),
			},
		},
	})
}

func TestAccDsCloudAccountGcp(t *testing.T) {
	grp := fmt.Sprintf("tf%s", acctest.RandString(6))
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	conf, err := cloudAccountFromEnv(account.TypeGcp, "x", name, []string{"prismacloud_account_group.x.group_id"})
	if err != nil {
		t.Skip(fmt.Sprintf("%s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsCloudAccountConfig(account.TypeGcp, grp, conf),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", "account_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.account_id", account.TypeGcp)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.enabled", account.TypeGcp)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.name", account.TypeGcp)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.compression_enabled", account.TypeGcp)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.flow_log_storage_bucket", account.TypeGcp)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.credentials_json", account.TypeGcp)),
				),
			},
		},
	})
}

func TestAccDsCloudAccountAlibaba(t *testing.T) {
	grp := fmt.Sprintf("tf%s", acctest.RandString(6))
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	conf, err := cloudAccountFromEnv(account.TypeAlibaba, "x", name, []string{"prismacloud_account_group.x.group_id"})
	if err != nil {
		t.Skip(fmt.Sprintf("%s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsCloudAccountConfig(account.TypeAlibaba, grp, conf),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", "account_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.account_id", account.TypeAlibaba)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.name", account.TypeGcp)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.ram_arn", account.TypeGcp)),
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_account.test", fmt.Sprintf("%s.0.enabled", account.TypeGcp)),
				),
			},
		},
	})
}

func testAccDsCloudAccountConfig(style, grp, conf string) string {
	return fmt.Sprintf(`
data "prismacloud_cloud_account" "test" {
    cloud_type = %q
    name = prismacloud_cloud_account.x.%s.0.name
}

resource "prismacloud_account_group" "x" {
    name = %q
    description = "acctest for cloud account data source"
}

%s
`, style, style, grp, conf)
}
