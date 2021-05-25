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
			"pyversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "v2",
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
	pyversion := d.Get("pyversion").(string)
	//ensure that args and scrript is not null
	d.SetId(args)
	d.SetId(script)
	d.SetId(pyversion)
	if pyversion == "v2" {
		scriptExec := "python " + script + " " + args
		cmd := exec.Command("sh", "-c", scriptExec)
		b, err2 := cmd.Output()
		if err2 != nil {
			return err2
		}
		d.Set("exec_py", string(b))
		return resourcePythonRead(d, m)
	} else {
		scriptExec := "python3 " + script + " " + args
		cmd := exec.Command("sh", "-c", scriptExec)
		b, err2 := cmd.Output()
		if err2 != nil {
			return err2
		}
		d.Set("exec_py", string(b))
		return resourcePythonRead(d, m)
	}
	return nil
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
