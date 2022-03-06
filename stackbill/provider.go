package stackbill

import (
	"context"
	"terraform-provider-stackbill/api"
	"terraform-provider-stackbill/auth"
	"terraform-provider-stackbill/computeoffering"
	"terraform-provider-stackbill/instance"
	"terraform-provider-stackbill/iso"
	"terraform-provider-stackbill/network"
	"terraform-provider-stackbill/securitygroup"
	"terraform-provider-stackbill/snapshot"
	"terraform-provider-stackbill/sshkey"
	"terraform-provider-stackbill/storageoffering"
	"terraform-provider-stackbill/template"
	"terraform-provider-stackbill/volume"
	"terraform-provider-stackbill/zone"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	authClient auth.AuthClient = auth.NewAuthClient()
)

// Provider Interface
type Provider interface {
	Provider() *schema.Provider
	configure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics)
}

// Provider struct for implementation
type provider struct{}

// Provider Object
func NewProvider() Provider {
	return &provider{}
}

// Provider
func (p *provider) Provider() *schema.Provider {
	pro := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"stackbill_instance":              instance.InstanceProvider(),
			"stackbill_instance_actions":      instance.InstanceActionsProvider(),
			"stackbill_instance_iso_actions":  instance.InstanceIsoActionsProvider(),
			"stackbill_instance_update_name":  instance.InstanceUpdateNameProvider(),
			"stackbill_instance_resize":       instance.InstanceResizeProvider(),
			"stackbill_instance_reset_sshkey": instance.InstanceResetSshKeyProvider(),
			"stackbill_instance_snapshot":     snapshot.VmSnapshotProvider(),
			"stackbill_volume_actions":        volume.VolumeActionsProvider(),
			"stackbill_volume":                volume.VolumeProvider(),
			"stackbill_network":               network.NetworkProvider(),
			"stackbill_network_actions":       network.NetworkActionProvider(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"stackbill_compute_offering_list":     computeoffering.ComputeOfferingListProvider(),
			"stackbill_instance_list":             instance.InstanceListProvider(),
			"stackbill_network_list":              network.NetworkListProvider(),
			"stackbill_network_offering_list":     network.NetworkOfferingListProvider(),
			"stackbill_vpc_network_offering_list": network.VpcNetworkOfferingListProvider(),
			"stackbill_storage_offering_list":     storageoffering.StorageOfferingListProvider(),
			"stackbill_zone_list":                 zone.ZoneListProvider(),
			"stackbill_template_list":             template.TemplateListProvider(),
			"stackbill_sshkey_list":               sshkey.SshkeyListProvider(),
			"stackbill_iso_list":                  iso.IsoListProvider(),
			"stackbill_vm_snapshot_list":          snapshot.VmSnapshotListProvider(),
			"stackbill_volume_list":               volume.VolumeListProvider(),
			"stackbill_security_group_list":       securitygroup.SecurityGroupListProvider(),
		},
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Stackbill URL",
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Stackbill Api key",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Stackbill Secret Key",
			},
		},
	}
	pro.ConfigureContextFunc = p.configure
	return pro
}

// Configure function
func (p *provider) configure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	url := d.Get("url").(string)
	apiKey := d.Get("api_key").(string)
	secretKey := d.Get("secret_key").(string)
	api.END_POINT = d.Get("url").(string)
	api.SECRET_KEY = secretKey
	api.API_KEY = apiKey
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := authClient.New(&apiKey, &secretKey, &url)
	return client, diags
}
