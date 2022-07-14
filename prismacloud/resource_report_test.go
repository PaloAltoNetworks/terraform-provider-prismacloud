package prismacloud

import (
	"fmt"
	"github.com/paloaltonetworks/prisma-cloud-go/report"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccReportRIS(t *testing.T) {
	var o report.Report
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccReportDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccReportConfigRIS(name, "RIS", "aws"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReportExists("prismacloud_report.test", &o),
					testAccCheckReportAttributes(&o, name, "RIS", "aws", false, false, ""),
				),
			},
		},
	})
}

func TestAccReportInventory(t *testing.T) {
	var o report.Report
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccReportDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccReportConfigInventory(name, "INVENTORY_DETAIL", "aws", true, true, "DTSTART;TZID=Asia/Calcutta\nBYHOUR=8;BYMINUTE=0;BYSECOND=0;FREQ=WEEKLY;INTERVAL=1;BYDAY=WE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReportExists("prismacloud_report.test", &o),
					testAccCheckReportAttributes(&o, name, "INVENTORY_DETAIL", "aws", true, true, "DTSTART;TZID=Asia/Calcutta\nBYHOUR=8;BYMINUTE=0;BYSECOND=0;FREQ=WEEKLY;INTERVAL=1;BYDAY=WE"),
				),
			},
		},
	})
}

func testAccCheckReportExists(n string, o *report.Report) resource.TestCheckFunc {
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
		lo, err := report.Get(client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo

		return nil
	}
}

func testAccCheckReportAttributes(o *report.Report, name, rt, ct string, compression_enabled, sch_enabled bool, sch string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("Name is %s, expected %s", o.Name, name)
		}

		if o.Type != rt {
			return fmt.Errorf("Report type is %q, expected %q", o.Type, rt)
		}

		if o.CloudType != ct {
			return fmt.Errorf("Cloud type is %q, expected %q", o.CloudType, ct)
		}

		if o.Target.CompressionEnabled != compression_enabled {
			return fmt.Errorf("Compression enabled is %t, expected %t", o.Target.CompressionEnabled, compression_enabled)
		}

		if o.Target.ScheduleEnabled != sch_enabled {
			return fmt.Errorf("Schedule enabled is %t, expected %t", o.Target.ScheduleEnabled, sch_enabled)
		}

		if o.Target.Schedule != sch {
			return fmt.Errorf("Schedule is %q, expected %q", o.Target.Schedule, sch)
		}

		return nil
	}
}

func testAccReportDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_report" {
			continue
		}

		if rs.Primary.ID != "" {
			id := rs.Primary.ID
			if _, err := report.Get(client, id); err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}

func testAccReportConfigRIS(name, rt, ct string) string {
	return fmt.Sprintf(`
resource "prismacloud_report" "test" {
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
}`, name, rt, ct)
}

func testAccReportConfigInventory(name, rt, ct string, compression_enabled, sch_enabled bool, sch string) string {
	return fmt.Sprintf(`
resource "prismacloud_report" "test" {
    name = %q
    report_type = %q
	cloud_type = %q
    target{
      compression_enabled = %t
      schedule_enabled = %t
      schedule = %q
      time_range  {
          to_now {
            unit = "epoch"
          }
      }
    }    
}`, name, rt, ct, compression_enabled, sch_enabled, sch)
}
