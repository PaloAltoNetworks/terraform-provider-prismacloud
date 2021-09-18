package prismacloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsReport(t *testing.T) {
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsReportConfig(name, "RIS", "aws"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_report.test", "report_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_report.test", "name"),
					resource.TestCheckResourceAttrSet("data.prismacloud_report.test", "last_modified_on"),
				),
			},
		},
	})
}

func testAccDsReportConfig(name, rt, ct string) string {
	return fmt.Sprintf(`
data "prismacloud_report" "test" {
    report_id = prismacloud_report.x.report_id
}

resource "prismacloud_report" "x" {
    name = %q
    report_type = %q
    cloud_type = %q
    target{
		time_range {
			relative {
				unit = "hour"
				amount = 24		
			}		
		 }
		  
    }    
}
`, name, rt, ct)
}
