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

func ResourceComputeNetworkEndpointGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkEndpointGroupCreate,
		Read:   resourceComputeNetworkEndpointGroupRead,
		Delete: resourceComputeNetworkEndpointGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkEndpointGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateGCEName,
				Description: `Name of the resource; provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The network to which all network endpoints in the NEG belong.
Uses "default" project network if unspecified.`,
			},
			"default_port": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Description: `The default port used if the port number is not specified in the
network endpoint.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `An optional description of this resource. Provide this property when
you create the resource.`,
			},
			"network_endpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"GCE_VM_IP", "GCE_VM_IP_PORT", "NON_GCP_PRIVATE_IP_PORT", ""}),
				Description: `Type of network endpoints in this network endpoint group.
NON_GCP_PRIVATE_IP_PORT is used for hybrid connectivity network
endpoint groups (see https://cloud.google.com/load-balancing/docs/hybrid).
Note that NON_GCP_PRIVATE_IP_PORT can only be used with Backend Services
that 1) have the following load balancing schemes: EXTERNAL, EXTERNAL_MANAGED,
INTERNAL_MANAGED, and INTERNAL_SELF_MANAGED and 2) support the RATE or
CONNECTION balancing modes.

Possible values include: GCE_VM_IP, GCE_VM_IP_PORT, and NON_GCP_PRIVATE_IP_PORT. Default value: "GCE_VM_IP_PORT" Possible values: ["GCE_VM_IP", "GCE_VM_IP_PORT", "NON_GCP_PRIVATE_IP_PORT"]`,
				Default: "GCE_VM_IP_PORT",
			},
			"subnetwork": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareOptionalSubnet,
				Description:      `Optional subnetwork to which all network endpoints in the NEG belong.`,
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `Zone where the network endpoint group is located.`,
			},
			"size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Number of network endpoints in the network endpoint group.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeNetworkEndpointGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeNetworkEndpointGroupName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeNetworkEndpointGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkEndpointTypeProp, err := expandComputeNetworkEndpointGroupNetworkEndpointType(d.Get("network_endpoint_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_endpoint_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkEndpointTypeProp)) && (ok || !reflect.DeepEqual(v, networkEndpointTypeProp)) {
		obj["networkEndpointType"] = networkEndpointTypeProp
	}
	networkProp, err := expandComputeNetworkEndpointGroupNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	subnetworkProp, err := expandComputeNetworkEndpointGroupSubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("subnetwork"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	defaultPortProp, err := expandComputeNetworkEndpointGroupDefaultPort(d.Get("default_port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_port"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultPortProp)) && (ok || !reflect.DeepEqual(v, defaultPortProp)) {
		obj["defaultPort"] = defaultPortProp
	}
	zoneProp, err := expandComputeNetworkEndpointGroupZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NetworkEndpointGroup: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkEndpointGroup: %s", err)
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
		return fmt.Errorf("Error creating NetworkEndpointGroup: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating NetworkEndpointGroup", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create NetworkEndpointGroup: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NetworkEndpointGroup %q: %#v", d.Id(), res)

	return resourceComputeNetworkEndpointGroupRead(d, meta)
}

func resourceComputeNetworkEndpointGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkEndpointGroup: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeNetworkEndpointGroup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}

	if err := d.Set("name", flattenComputeNetworkEndpointGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("description", flattenComputeNetworkEndpointGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("network_endpoint_type", flattenComputeNetworkEndpointGroupNetworkEndpointType(res["networkEndpointType"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("size", flattenComputeNetworkEndpointGroupSize(res["size"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("network", flattenComputeNetworkEndpointGroupNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("subnetwork", flattenComputeNetworkEndpointGroupSubnetwork(res["subnetwork"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("default_port", flattenComputeNetworkEndpointGroupDefaultPort(res["defaultPort"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("zone", flattenComputeNetworkEndpointGroupZone(res["zone"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}

	return nil
}

func resourceComputeNetworkEndpointGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkEndpointGroup: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting NetworkEndpointGroup %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "NetworkEndpointGroup")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting NetworkEndpointGroup", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting NetworkEndpointGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeNetworkEndpointGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/networkEndpointGroups/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeNetworkEndpointGroupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkEndpointGroupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkEndpointGroupNetworkEndpointType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkEndpointGroupSize(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeNetworkEndpointGroupNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeNetworkEndpointGroupSubnetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeNetworkEndpointGroupDefaultPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeNetworkEndpointGroupZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func expandComputeNetworkEndpointGroupName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEndpointGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEndpointGroupNetworkEndpointType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEndpointGroupNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeNetworkEndpointGroupSubnetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("subnetworks", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for subnetwork: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeNetworkEndpointGroupDefaultPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEndpointGroupZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}
