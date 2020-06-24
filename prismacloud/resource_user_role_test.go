package prismacloud

import (
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/user/role"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccUserRole(t *testing.T) {
	var o role.Role
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccUserRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserRoleConfig(name, "first desc", "Account Group Read Only"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserRoleExists("prismacloud_user_role.test", &o),
					testAccCheckUserRoleAttributes(&o, name, "first desc", "Account Group Read Only"),
				),
			},
			{
				Config: testAccUserRoleConfig(name, "second desc", "Cloud Provisioning Admin"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserRoleExists("prismacloud_user_role.test", &o),
					testAccCheckUserRoleAttributes(&o, name, "second desc", "Cloud Provisioning Admin"),
				),
			},
		},
	})
}

func testAccCheckUserRoleExists(n string, o *role.Role) resource.TestCheckFunc {
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
		lo, err := role.Get(client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo

		return nil
	}
}

func testAccCheckUserRoleAttributes(o *role.Role, name, desc, rt string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("Name is %s, expected %s", o.Name, name)
		}

		if o.Description != desc {
			return fmt.Errorf("Description is %s, expected %s", o.Description, desc)
		}

		if o.RoleType != rt {
			return fmt.Errorf("Role type is %q, expected %q", o.RoleType, rt)
		}

		return nil
	}
}

func testAccUserRoleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_user_role" {
			continue
		}

		if rs.Primary.ID != "" {
			id := rs.Primary.ID
			if _, err := role.Get(client, id); err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}

func testAccUserRoleConfig(name, desc, rt string) string {
	return fmt.Sprintf(`
resource "prismacloud_user_role" "test" {
    name = %q
    description = %q
    role_type = %q
}
`, name, desc, rt)
}
