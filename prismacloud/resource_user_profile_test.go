package prismacloud

import (
	"fmt"
	"github.com/paloaltonetworks/prisma-cloud-go/user/profile"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

func TestAccUser(t *testing.T) {
	var o profile.Profile
	firstname := fmt.Sprintf("tf%s", acctest.RandString(6))
	secondname := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccUserProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserProfileConfig(firstname, secondname, "test@paloaltonetworks.com", "Test user Profile", "Asia/Calcutta"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserProfileExists("˳", &o),
					testAccCheckUserProfileAttributes(&o, firstname, secondname, "test@paloaltonetworks.com", "Test user Profile", "Asia/Calcutta"),
				),
			},
			{
				Config: testAccUserProfileConfig(firstname, secondname, "test@paloaltonetworks.com", "Test user Profile", "Asia/Calcutta"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserProfileExists("prismacloud_user_profile.test", &o),
					testAccCheckUserProfileAttributes(&o, firstname, secondname, "test@paloaltonetworks.com", "Test user Profile", "Asia/Calcutta"),
				),
			},
		},
	})
}

func testAccCheckUserProfileExists(n string, o *profile.Profile) resource.TestCheckFunc {
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
		lo, err := profile.Get(client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo

		return nil
	}
}

func testAccCheckUserProfileAttributes(o *profile.Profile, firstname, lastname, email, username, timezone string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.FirstName != firstname {
			return fmt.Errorf("FirstName is %s, expected %s", o.FirstName, firstname)
		}

		if o.LastName != lastname {
			return fmt.Errorf("LastName is %s, expected %s", o.LastName, lastname)
		}

		if o.Email != email {
			return fmt.Errorf("Email is %q, expected %q", o.Email, email)
		}

		if o.Username != username {
			return fmt.Errorf("Username is %q, expected %q", o.DisplayName, username)
		}

		if o.TimeZone != timezone {
			return fmt.Errorf("TimeZone is %q, expected %q", o.TimeZone, timezone)
		}

		return nil
	}
}

func testAccUserProfileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_user_profile" {
			continue
		}

		if rs.Primary.ID != "" {
			id := rs.Primary.ID
			if _, err := profile.Get(client, id); err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}

func testAccUserProfileConfig(firstname, lastname, email, username, timezone string) string {
	return fmt.Sprintf(`
data "prismacloud_user_roles" "x" {}

locals {
    id_listing = [
        for inst in data.prismacloud_user_roles.x.listing :
        inst.search_id
    ]
}
resource "prismacloud_user_profile" "test" {
    first_name = %q
  	last_name = %q
  	email = %q
  	username = %q
  	role_ids = [
    	local.id_listing[0]
  	]
  	time_zone = %q
  	default_role_id = local.id_listing[0]
}
`, firstname, lastname, email, username, timezone)
}
