package prismacloud

import (
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudAccountAws(t *testing.T) {
	var o account.Aws
	name := fmt.Sprintf("tf%s", acctest.RandString(6))
	g1 := fmt.Sprintf("tf%s", acctest.RandString(6))
	g2 := fmt.Sprintf("tf%s", acctest.RandString(6))

	if _, err := cloudAccountFromEnv(account.TypeAws, "x", name, nil); err != nil {
		t.Skip(err.Error())
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudAccountConfig(account.TypeAws, g1, g2, name, []string{"prismacloud_account_group.x.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAwsExists("prismacloud_cloud_account.test", &o),
					testAccCheckCloudAccountAwsAttributes(&o, name, 1),
				),
			},
			{
				Config: testAccCloudAccountConfig(account.TypeAws, g1, g2, name, []string{"prismacloud_account_group.x.group_id", "prismacloud_account_group.y.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAwsExists("prismacloud_cloud_account.test", &o),
					testAccCheckCloudAccountAwsAttributes(&o, name, 2),
				),
			},
		},
	})
}

func TestAccCloudAccountAzure(t *testing.T) {
	var o account.Azure
	name := fmt.Sprintf("tf%s", acctest.RandString(6))
	g1 := fmt.Sprintf("tf%s", acctest.RandString(6))
	g2 := fmt.Sprintf("tf%s", acctest.RandString(6))

	if _, err := cloudAccountFromEnv(account.TypeAzure, "x", name, nil); err != nil {
		t.Skip(err.Error())
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudAccountConfig(account.TypeAzure, g1, g2, name, []string{"prismacloud_account_group.x.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAzureExists("prismacloud_cloud_account.test", &o),
					testAccCheckCloudAccountAzureAttributes(&o, name, 1),
				),
			},
			{
				Config: testAccCloudAccountConfig(account.TypeAzure, g1, g2, name, []string{"prismacloud_account_group.x.group_id", "prismacloud_account_group.y.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAzureExists("prismacloud_cloud_account.test", &o),
					testAccCheckCloudAccountAzureAttributes(&o, name, 2),
				),
			},
		},
	})
}

func TestAccCloudAccountGcp(t *testing.T) {
	var o account.Gcp
	name := fmt.Sprintf("tf%s", acctest.RandString(6))
	g1 := fmt.Sprintf("tf%s", acctest.RandString(6))
	g2 := fmt.Sprintf("tf%s", acctest.RandString(6))

	if _, err := cloudAccountFromEnv(account.TypeGcp, "x", name, nil); err != nil {
		t.Skip(err.Error())
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudAccountConfig(account.TypeGcp, g1, g2, name, []string{"prismacloud_account_group.x.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountGcpExists("prismacloud_cloud_account.test", &o),
					testAccCheckCloudAccountGcpAttributes(&o, name, 1),
				),
			},
			{
				Config: testAccCloudAccountConfig(account.TypeGcp, g1, g2, name, []string{"prismacloud_account_group.x.group_id", "prismacloud_account_group.y.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountGcpExists("prismacloud_cloud_account.test", &o),
					testAccCheckCloudAccountGcpAttributes(&o, name, 2),
				),
			},
		},
	})
}

func TestAccCloudAccountAlibaba(t *testing.T) {
	var o account.Alibaba
	name := fmt.Sprintf("tf%s", acctest.RandString(6))
	g1 := fmt.Sprintf("tf%s", acctest.RandString(6))
	g2 := fmt.Sprintf("tf%s", acctest.RandString(6))

	if _, err := cloudAccountFromEnv(account.TypeAlibaba, "x", name, nil); err != nil {
		t.Skip(err.Error())
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudAccountConfig(account.TypeAlibaba, g1, g2, name, []string{"prismacloud_account_group.x.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAlibabaExists("prismacloud_cloud_account.test", &o),
					testAccCheckCloudAccountAlibabaAttributes(&o, name, 1),
				),
			},
			{
				Config: testAccCloudAccountConfig(account.TypeAlibaba, g1, g2, name, []string{"prismacloud_account_group.x.group_id", "prismacloud_account_group.y.group_id"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAlibabaExists("prismacloud_cloud_account.test", &o),
					testAccCheckCloudAccountAlibabaAttributes(&o, name, 2),
				),
			},
		},
	})
}

func testAccCheckCloudAccountAwsExists(n string, o *account.Aws) resource.TestCheckFunc {
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
		lo, err := account.Get(client, ct, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo.(account.Aws)

		return nil
	}
}

func testAccCheckCloudAccountAzureExists(n string, o *account.Azure) resource.TestCheckFunc {
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
		lo, err := account.Get(client, ct, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo.(account.Azure)

		return nil
	}
}

func testAccCheckCloudAccountGcpExists(n string, o *account.Gcp) resource.TestCheckFunc {
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
		lo, err := account.Get(client, ct, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo.(account.Gcp)

		return nil
	}
}

func testAccCheckCloudAccountAlibabaExists(n string, o *account.Alibaba) resource.TestCheckFunc {
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
		lo, err := account.Get(client, ct, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo.(account.Alibaba)

		return nil
	}
}

func testAccCheckCloudAccountAwsAttributes(o *account.Aws, name string, gl int) resource.TestCheckFunc {
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

func testAccCheckCloudAccountAzureAttributes(o *account.Azure, name string, gl int) resource.TestCheckFunc {
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

func testAccCheckCloudAccountGcpAttributes(o *account.Gcp, name string, gl int) resource.TestCheckFunc {
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

func testAccCheckCloudAccountAlibabaAttributes(o *account.Alibaba, name string, gl int) resource.TestCheckFunc {
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

func testAccCloudAccountDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_cloud_account" {
			continue
		}

		if rs.Primary.ID != "" {
			ct, id := IdToTwoStrings(rs.Primary.ID)
			if _, err := account.Get(client, ct, id); err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}

func testAccCloudAccountConfig(style, g1, g2, name string, groups []string) string {
	conf, _ := cloudAccountFromEnv(style, "test", name, groups)
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
