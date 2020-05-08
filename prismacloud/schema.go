package prismacloud

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func totalSchema(desc string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeInt,
		Computed:    true,
		Description: fmt.Sprintf("Total number of %s", desc),
	}
}
