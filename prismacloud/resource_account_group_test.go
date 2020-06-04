package prismacloud

import (
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account/group"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAccountGroup(t *testing.T) {
	var o group.Group
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccAccountGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAccountGroupConfig(name, "first desc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountGroupExists("prismacloud_account_group.test", &o),
					testAccCheckAccountGroupAttributes(&o, name, "first desc"),
				),
			},
			{
				Config: testAccAccountGroupConfig(name, "second desc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountGroupExists("prismacloud_account_group.test", &o),
					testAccCheckAccountGroupAttributes(&o, name, "second desc"),
				),
			},
		},
	})
}

func testAccCheckAccountGroupExists(n string, o *group.Group) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label ID is not set")
		}

		client := testAccProvider.Meta().(*pc.Client)
		id := rs.Primary.ID
		lo, err := group.Get(client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo

		return nil
	}
}

func testAccCheckAccountGroupAttributes(o *group.Group, name, desc string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("Name is %s, expected %s", o.Name, name)
		}

		if o.Description != desc {
			return fmt.Errorf("Description is %s, expected %s", o.Description, desc)
		}

		return nil
	}
}

func testAccAccountGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_account_group" {
			continue
		}

		if rs.Primary.ID != "" {
			id := rs.Primary.ID
			if _, err := group.Get(client, id); err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}

func testAccAccountGroupConfig(name, desc string) string {
	return fmt.Sprintf(`
resource "prismacloud_account_group" "test" {
    name = %q
    description = %q
}
`, name, desc)
}
