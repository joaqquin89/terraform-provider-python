package python

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"python_exec": resourcePython(), // Add this line
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
