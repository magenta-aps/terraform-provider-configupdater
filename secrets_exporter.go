// Implement /secrets_export/*

package main

import (
	"bytes"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"net/http"
)

func secretsExport() *schema.Resource {
	return &schema.Resource{
		Create: secretCreate,
		Read:   secretRead,
		Update: secretCreate,
		Delete: secretDelete,

		Schema: map[string]*schema.Schema{
			"file_path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"secret": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}

func secretCreate(d *schema.ResourceData, meta interface{}) error {
	file_path := d.Get("file_path").(string)
	secrets, err := json.Marshal(d.Get("secret"))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		http.MethodPut,
		meta.(string)+"secrets_export/export/"+file_path,
		bytes.NewBuffer(secrets),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode > 299 {
		panic("Status code over > 299")
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	d.SetId(string(data))

	return secretRead(d, meta)
}

func secretRead(d *schema.ResourceData, meta interface{}) error {
	file_path := d.Get("file_path").(string)

	resp, err := http.Get(meta.(string) + "secrets_export/export/" + file_path)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	d.SetId(string(data))

	return nil
}

func secretDelete(d *schema.ResourceData, meta interface{}) error {
	// Dummy unset, as config-updater does not implment DELETE verb for
	// secrets_export yet.
	d.SetId("")
	return nil
}
