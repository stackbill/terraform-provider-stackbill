package snapshot

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func VmSnapshotProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Instance Snapshot Api's",
		CreateContext: vmSnapshotResourceObj.Create,
		ReadContext:   vmSnapshotResourceObj.Read,
		UpdateContext: vmSnapshotResourceObj.Update,
		DeleteContext: vmSnapshotResourceObj.Delete,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_machine_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"zone_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"snapshot_memory": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}
