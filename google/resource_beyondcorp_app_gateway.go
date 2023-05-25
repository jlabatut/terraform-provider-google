// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceBeyondcorpAppGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceBeyondcorpAppGatewayCreate,
		Read:   resourceBeyondcorpAppGatewayRead,
		Delete: resourceBeyondcorpAppGatewayDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBeyondcorpAppGatewayImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID of the AppGateway.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An arbitrary user-provided name for the AppGateway.`,
			},
			"host_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"HOST_TYPE_UNSPECIFIED", "GCP_REGIONAL_MIG", ""}),
				Description:  `The type of hosting used by the AppGateway. Default value: "HOST_TYPE_UNSPECIFIED" Possible values: ["HOST_TYPE_UNSPECIFIED", "GCP_REGIONAL_MIG"]`,
				Default:      "HOST_TYPE_UNSPECIFIED",
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: `Resource labels to represent user provided metadata.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the AppGateway.`,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"TYPE_UNSPECIFIED", "TCP_PROXY", ""}),
				Description:  `The type of network connectivity used by the AppGateway. Default value: "TYPE_UNSPECIFIED" Possible values: ["TYPE_UNSPECIFIED", "TCP_PROXY"]`,
				Default:      "TYPE_UNSPECIFIED",
			},
			"allocated_connections": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `A list of connections allocated for the Gateway.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ingress_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: `The ingress port of an allocated connection.`,
						},
						"psc_uri": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The PSC uri of an allocated connection.`,
						},
					},
				},
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Represents the different states of a AppGateway.`,
			},
			"uri": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URI for this resource.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceBeyondcorpAppGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	typeProp, err := expandBeyondcorpAppGatewayType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	hostTypeProp, err := expandBeyondcorpAppGatewayHostType(d.Get("host_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("host_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(hostTypeProp)) && (ok || !reflect.DeepEqual(v, hostTypeProp)) {
		obj["hostType"] = hostTypeProp
	}
	displayNameProp, err := expandBeyondcorpAppGatewayDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandBeyondcorpAppGatewayLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{region}}/appGateways?app_gateway_id={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AppGateway: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppGateway: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating AppGateway: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/appGateways/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = BeyondcorpOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating AppGateway", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create AppGateway: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/appGateways/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating AppGateway %q: %#v", d.Id(), res)

	return resourceBeyondcorpAppGatewayRead(d, meta)
}

func resourceBeyondcorpAppGatewayRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{region}}/appGateways/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppGateway: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BeyondcorpAppGateway %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading AppGateway: %s", err)
	}

	if err := d.Set("type", flattenBeyondcorpAppGatewayType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppGateway: %s", err)
	}
	if err := d.Set("host_type", flattenBeyondcorpAppGatewayHostType(res["hostType"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppGateway: %s", err)
	}
	if err := d.Set("display_name", flattenBeyondcorpAppGatewayDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppGateway: %s", err)
	}
	if err := d.Set("labels", flattenBeyondcorpAppGatewayLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppGateway: %s", err)
	}
	if err := d.Set("state", flattenBeyondcorpAppGatewayState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppGateway: %s", err)
	}
	if err := d.Set("uri", flattenBeyondcorpAppGatewayUri(res["uri"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppGateway: %s", err)
	}
	if err := d.Set("allocated_connections", flattenBeyondcorpAppGatewayAllocatedConnections(res["allocatedConnections"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppGateway: %s", err)
	}

	return nil
}

func resourceBeyondcorpAppGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppGateway: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{region}}/appGateways/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting AppGateway %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "AppGateway")
	}

	err = BeyondcorpOperationWaitTime(
		config, res, project, "Deleting AppGateway", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting AppGateway %q: %#v", d.Id(), res)
	return nil
}

func resourceBeyondcorpAppGatewayImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/appGateways/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/appGateways/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBeyondcorpAppGatewayType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppGatewayHostType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppGatewayDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppGatewayLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppGatewayState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppGatewayUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppGatewayAllocatedConnections(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["psc_uri"] =
		flattenBeyondcorpAppGatewayAllocatedConnectionsPscUri(original["pscUri"], d, config)
	transformed["ingress_port"] =
		flattenBeyondcorpAppGatewayAllocatedConnectionsIngressPort(original["ingressPort"], d, config)
	return []interface{}{transformed}
}
func flattenBeyondcorpAppGatewayAllocatedConnectionsPscUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppGatewayAllocatedConnectionsIngressPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandBeyondcorpAppGatewayType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppGatewayHostType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppGatewayDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppGatewayLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
