package python

import (
	"os/exec"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePython() *schema.Resource {
	return &schema.Resource{
		Create: resourcePythonCreate,
		Read:   resourcePythonRead,
		Update: resourcePythonUpdate,
		Delete: resourcePythonDelete,
		Schema: map[string]*schema.Schema{
			"args": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"exec_py": {
				Type:     schema.TypeString,
				Computed: true,
				//Sensitive: true,
			},
		},
	}
}

func resourcePythonCreate(d *schema.ResourceData, m interface{}) error {

	args := d.Get("args").(string)
	script := d.Get("script").(string)
	//ensure that args and scrript is not null
	d.SetId(args)
	d.SetId(script)
	scriptExec := "python " + script + " " + args
	cmd := exec.Command("bash", "-c", scriptExec)
	b, err2 := cmd.Output()
	if err2 != nil {
		return err2
	}
	d.Set("exec_py", string(b))
	return resourcePythonRead(d, m)
}

func resourcePythonRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePythonUpdate(d *schema.ResourceData, m interface{}) error {

	return resourcePythonRead(d, m)
}

func resourcePythonDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
