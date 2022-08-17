package prismacloud

import (
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account/org"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudAccountAwsorg(t *testing.T) {
	var o org.AwsOrg
	name := fmt.Sprintf("tf%s", acctest.RandString(6))
	g1 := fmt.Sprintf("tf%s", acctest.RandString(6))
	g2 := fmt.Sprintf("tf%s", acctest.RandString(6))
	if _, err := cloudAccountOrgFromEnv(org.TypeAwsOrg, "x", name, nil); err != nil {
		t.Skip(err.Error())
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudorgAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudorgAccountConfig(org.TypeAwsOrg, g1, g2, name, []string{"prismacloud_account_group.x.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAwsorgExists("prismacloud_org_cloud_account.test", &o),
					testAccCheckCloudAccountAwsorgAttributes(&o, name, 1),
				),
			},
			{
				Config: testAccCloudorgAccountConfig(org.TypeAwsOrg, g1, g2, name, []string{"prismacloud_account_group.x.group_id", "prismacloud_account_group.y.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAwsorgExists("prismacloud_org_cloud_account.test", &o),
					testAccCheckCloudAccountAwsorgAttributes(&o, name, 2),
				),
			},
		},
	})
}

func TestAccCloudAccountAzureorg(t *testing.T) {
	var o org.AzureOrg
	name := fmt.Sprintf("tf%s", acctest.RandString(6))
	g1 := fmt.Sprintf("tf%s", acctest.RandString(6))
	g2 := fmt.Sprintf("tf%s", acctest.RandString(6))

	if _, err := cloudAccountOrgFromEnv(org.TypeAzureOrg, "x", name, nil); err != nil {
		t.Skip(err.Error())
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudorgAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudorgAccountConfig(org.TypeAzureOrg, g1, g2, name, []string{"prismacloud_account_group.x.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAzureorgExists("prismacloud_org_cloud_account.test", &o),
					testAccCheckCloudAccountAzureorgAttributes(&o, name, 1),
				),
			},
			{
				Config: testAccCloudorgAccountConfig(org.TypeAzureOrg, g1, g2, name, []string{"prismacloud_account_group.x.group_id", "prismacloud_account_group.y.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAzureorgExists("prismacloud_org_cloud_account.test", &o),
					testAccCheckCloudAccountAzureorgAttributes(&o, name, 2),
				),
			},
		},
	})
}

func TestAccCloudAccountGcporg(t *testing.T) {
	var o org.GcpOrg
	name := fmt.Sprintf("tf%s", acctest.RandString(6))
	g1 := fmt.Sprintf("tf%s", acctest.RandString(6))
	g2 := fmt.Sprintf("tf%s", acctest.RandString(6))

	if _, err := cloudAccountOrgFromEnv(org.TypeGcpOrg, "x", name, nil); err != nil {
		t.Skip(err.Error())
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudorgAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudorgAccountConfig(org.TypeGcpOrg, g1, g2, name, []string{"prismacloud_account_group.x.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountGcporgExists("prismacloud_org_cloud_account.test", &o),
					testAccCheckCloudAccountGcporgAttributes(&o, name, 1),
				),
			},
			{
				Config: testAccCloudorgAccountConfig(org.TypeGcpOrg, g1, g2, name, []string{"prismacloud_account_group.x.group_id", "prismacloud_account_group.y.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountGcporgExists("prismacloud_org_cloud_account.test", &o),
					testAccCheckCloudAccountGcporgAttributes(&o, name, 2),
				),
			},
		},
	})
}

func TestAccCloudAccountOci(t *testing.T) {
	var o org.Oci
	name := fmt.Sprintf("tf%s", acctest.RandString(6))
	g1 := fmt.Sprintf("tf%s", acctest.RandString(6))
	g2 := fmt.Sprintf("tf%s", acctest.RandString(6))

	if _, err := cloudAccountOrgFromEnv(org.TypeOci, "x", name, nil); err != nil {
		t.Skip(err.Error())
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudorgAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudorgAccountConfig(org.TypeOci, g1, g2, name, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountOciExists("prismacloud_org_cloud_account.test", &o),
				),
			},
			{
				Config: testAccCloudorgAccountConfig(org.TypeOci, g1, g2, name, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountOciExists("prismacloud_org_cloud_account.test", &o),
				),
			},
		},
	})
}

func testAccCheckCloudAccountAwsorgExists(n string, o *org.AwsOrg) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label ID is not set")
		}

		client := testAccProvider.Meta().(*pc.Client)
		ct, id := IdToTwoStrings(rs.Primary.ID)
		lo, err := org.Get(client, ct, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo.(org.AwsOrg)

		return nil
	}
}

func testAccCheckCloudAccountAzureorgExists(n string, o *org.AzureOrg) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label ID is not set")
		}

		client := testAccProvider.Meta().(*pc.Client)
		ct, id := IdToTwoStrings(rs.Primary.ID)
		lo, err := org.Get(client, ct, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo.(org.AzureOrg)

		return nil
	}
}

func testAccCheckCloudAccountGcporgExists(n string, o *org.GcpOrg) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label ID is not set")
		}

		client := testAccProvider.Meta().(*pc.Client)
		ct, id := IdToTwoStrings(rs.Primary.ID)
		lo, err := org.Get(client, ct, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo.(org.GcpOrg)

		return nil
	}
}

func testAccCheckCloudAccountOciExists(n string, o *org.Oci) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label ID is not set")
		}

		client := testAccProvider.Meta().(*pc.Client)
		ct, id := IdToTwoStrings(rs.Primary.ID)
		lo, err := org.Get(client, ct, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo.(org.Oci)

		return nil
	}
}

func testAccCheckCloudAccountAwsorgAttributes(o *org.AwsOrg, name string, gl int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("Name is %s, expected %s", o.Name, name)
		}

		if len(o.GroupIds) != gl {
			return fmt.Errorf("Group IDs len is %d, not %d", len(o.GroupIds), gl)
		}

		return nil
	}
}

func testAccCheckCloudAccountAzureorgAttributes(o *org.AzureOrg, name string, gl int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Account.Name != name {
			return fmt.Errorf("Name is %s, expected %s", o.Account.Name, name)
		}

		if len(o.Account.GroupIds) != gl {
			return fmt.Errorf("Group IDs len is %d, not %d", len(o.Account.GroupIds), gl)
		}

		return nil
	}
}

func testAccCheckCloudAccountGcporgAttributes(o *org.GcpOrg, name string, gl int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Account.Name != name {
			return fmt.Errorf("Name is %s, expected %s", o.Account.Name, name)
		}

		if len(o.Account.GroupIds) != gl {
			return fmt.Errorf("Group IDs len is %d, not %d", len(o.Account.GroupIds), gl)
		}

		return nil
	}
}

func testAccCloudorgAccountDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_org_cloud_account" {
			continue
		}

		if rs.Primary.ID != "" {
			ct, id := IdToTwoStrings(rs.Primary.ID)
			if _, err := org.Get(client, ct, id); err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}

func testAccCloudorgAccountConfig(style, g1, g2, name string, groups []string) string {
	conf, _ := cloudAccountOrgFromEnv(style, "test", name, groups)
	return fmt.Sprintf(`
resource "prismacloud_account_group" "x" {
    name = %q
    description = "for %s cloud account acctest"
}

resource "prismacloud_account_group" "y" {
    name = %q
    description = "for %s cloud account acctest"
}

%s
`, g1, style, g2, style, conf)
}
