package prismacloud

import (
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/data-security/dataprofile"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestDataProfile(t *testing.T) {
	var o dataprofile.Profile
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccDataProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataProfileConfig(name, "desc made by terraform", "low", "include", "Source Code - javascript", "any"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataProfileExists("prismacloud_dataprofile.test", &o),
					testAccCheckDataProfileAttributes(&o, name, "desc made by terraform", "low", "include", "Source Code - javascript", "any"),
				),
			},
		},
	})
}

func testAccCheckDataProfileExists(n string, o *dataprofile.Profile) resource.TestCheckFunc {
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
		lo, err := dataprofile.Get(client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo

		return nil
	}
}

func testAccCheckDataProfileAttributes(o *dataprofile.Profile, name, desc, cl, mt, rn, opt string) resource.TestCheckFunc {
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

func testAccDataProfileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_datapolicy" {
			continue
		}

		if rs.Primary.ID != "" {
			id := rs.Primary.ID
			if _, err := dataprofile.Get(client, id); err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}

func testAccDataProfileConfig(name, desc, cl, mt, rn, opt string) string {
	return fmt.Sprintf(`
resource "prismacloud_datapolicy" "test" {
  name = %q
  description = %q
  profile_type = "custom"
  status = "non_hidden"
  profile_status = "disabled"
  types = "basic"
  data_patterns_rule_1{
    operator_type = "or"
	data_pattern_rules {
        confidence_level = %q
        match_type = %q
        name = %q
        occurrence_operator_type = %q
    } 
  }
}
`, name, desc, cl, mt, rn, opt)
}
