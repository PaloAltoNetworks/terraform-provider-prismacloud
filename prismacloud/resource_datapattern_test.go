package prismacloud

import (
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/data-security/datapattern"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestDataPattern(t *testing.T) {
	var o datapattern.Pattern
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccDataPatternDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataPatternConfig(name, "desc made by terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataPatternExists("prismacloud_datapattern.test", &o),
					testAccCheckDataPatternAttributes(&o, name, "desc made by terraform"),
				),
			},
		},
	})
}

func testAccCheckDataPatternExists(n string, o *datapattern.Pattern) resource.TestCheckFunc {
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
		lo, err := datapattern.Get(client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo

		return nil
	}
}

func testAccCheckDataPatternAttributes(o *datapattern.Pattern, name, desc string) resource.TestCheckFunc {
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

func testAccDataPatternDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_datapattern" {
			continue
		}

		if rs.Primary.ID != "" {
			id := rs.Primary.ID
			if _, err := datapattern.Get(client, id); err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}

func testAccDataPatternConfig(name, desc string) string {
	return fmt.Sprintf(`
resource "prismacloud_datapattern" "test" {
  name = %q
  description = %q
  proximity_keywords = ["terraform", "prisma"]
  regexes{
    regex = "prisma"
    weight = 4
  }
}
`, name, desc)
}
