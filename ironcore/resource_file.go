package ironcore

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/ironcore864/terraform-provider-ironcore/helper"
)

func resourceFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFileCreate,
		Read:   resourceFileRead,
		Update: resourceFileUpdate,
		Delete: resourceFileDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFileCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	if helper.FileExists(name) {
		log.Printf("[WARN] File already exists: %s", name)
		return fmt.Errorf("File already exists: %s", name)
	}
	file, err := os.Create(name)
	if err != nil {
		log.Printf("[WARN] No file created: %s", name)
		return fmt.Errorf("[WARN] No file created: %s", name)
	}
	d.SetId(file.Name())
	return resourceFileRead(d, m)
}

func resourceFileRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	if !helper.FileExists(name) {
		log.Printf("[WARN] No file found: %s", name)
		d.SetId("")
		return nil
	}
	return nil
}

func resourceFileUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	if d.HasChange("name") {
		name := d.Get("name").(string)
		if err := helper.FileRename(d.Id(), name); err != nil {
			log.Printf("[WARN] File rename failed: %s", d.Id())
			return err
		}
		d.SetPartial("name")
		d.SetId(name)
	}
	d.Partial(false)
	return resourceFileRead(d, m)
}

func resourceFileDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	err := os.Remove(name)
	if err != nil {
		log.Printf("[WARN] File not deleted: %s", d.Id())
		return err
	}
	d.SetId("")
	return nil
}
