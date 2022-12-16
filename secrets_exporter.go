// Implement /secrets_export/*

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"log"
	"net/http"
        "github.com/asaskevich/govalidator"
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
		return err
	}

	config := meta.(*Configuration)
        _, config_error := govalidator.ValidateStruct(config)
        if config_error != nil {
            log.Printf("error: " + config_error.Error())
            return config_error
        }


	client := &http.Client{}
	url := config.url + "secrets_export/export/" + file_path
	log.Printf("secrets_exporter.go: creating new secret %s\n", url)

	req, err := http.NewRequest(
		http.MethodPut,
		url,
		bytes.NewBuffer(secrets),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

        req.SetBasicAuth(config.username, config.password)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("secrets_exporter.go: HTTP status code = %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	id := string(data)
	log.Printf("secrets_exporter.go: success, setting ID to '%s'\n", id)
	d.SetId(id)

	return secretRead(d, meta)
}

func secretRead(d *schema.ResourceData, meta interface{}) error {
	file_path := d.Get("file_path").(string)

	config := meta.(*Configuration)
        _, config_error := govalidator.ValidateStruct(config)
        if config_error != nil {
            println("error: " + config_error.Error())
            return config_error
        }

	client := &http.Client{}
	url := config.url + "secrets_export/export/" + file_path
	log.Printf("secrets_exporter.go: reading %s\n", url)

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return err
	}

	if config.username != "" && config.password != "" {
		req.SetBasicAuth(config.username, config.password)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("secrets_exporter.go: HTTP status code = %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	id := string(data)
	log.Printf("secrets_exporter.go: success, setting ID to '%s'\n", id)
	d.SetId(id)

	return nil
}

func secretDelete(d *schema.ResourceData, meta interface{}) error {
	// Dummy unset, as config-updater does not implment DELETE verb for
	// secrets_export yet.
	d.SetId("")
	return nil
}
