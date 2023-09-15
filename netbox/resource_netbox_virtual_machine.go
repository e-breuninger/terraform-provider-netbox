package netbox

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxVirtualMachineStatusOptions = []string{"offline", "active", "planned", "staged", "failed", "decommissioning"}

func resourceNetboxVirtualMachine() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxVirtualMachineCreate,
		ReadContext:   resourceNetboxVirtualMachineRead,
		UpdateContext: resourceNetboxVirtualMachineUpdate,
		DeleteContext: resourceNetboxVirtualMachineDelete,

		Description: `:meta:subcategory:Virtualization:From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#virtual-machines):

> A virtual machine is a virtualized compute instance. These behave in NetBox very similarly to device objects, but without any physical attributes. For example, a VM may have interfaces assigned to it with IP addresses and VLANs, however its interfaces cannot be connected via cables (because they are virtual). Each VM may also define its compute, memory, and storage resources as well.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				AtLeastOneOf: []string{"site_id", "cluster_id"},
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"device_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"platform_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"site_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				AtLeastOneOf: []string{"site_id", "cluster_id"},
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"memory_mb": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vcpus": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"disk_size_gb": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxVirtualMachineStatusOptions, false),
				Default:      "active",
				Description:  buildValidValueDescription(resourceNetboxVirtualMachineStatusOptions),
			},
			tagsKey: tagsSchema,
			"primary_ipv4": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"primary_ipv6": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"local_context_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This is best managed through the use of `jsonencode` and a map of settings.",
			},
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceNetboxVirtualMachineResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceNetboxVirtualMachineStateUpgradeV0,
				Version: 0,
			},
		},
	}
}

func resourceNetboxVirtualMachineCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)

	data := models.WritableVirtualMachineWithConfigContext{
		Name: &name,
	}

	clusterIDValue, ok := d.GetOk("cluster_id")
	if ok {
		clusterID := int64(clusterIDValue.(int))
		data.Cluster = &clusterID
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	data.Comments = d.Get("comments").(string)
	data.Description = d.Get("description").(string)

	vcpusValue, ok := d.GetOk("vcpus")
	if ok {
		vcpus := vcpusValue.(float64)
		data.Vcpus = &vcpus
	}

	memoryMbValue, ok := d.GetOk("memory_mb")
	if ok {
		memoryMb := int64(memoryMbValue.(int))
		data.Memory = &memoryMb
	}

	diskSizeValue, ok := d.GetOk("disk_size_gb")
	if ok {
		diskSize := int64(diskSizeValue.(int))
		data.Disk = &diskSize
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	deviceIDValue, ok := d.GetOk("device_id")
	if ok {
		deviceID := int64(deviceIDValue.(int))
		data.Device = &deviceID
	}

	platformIDValue, ok := d.GetOk("platform_id")
	if ok {
		platformID := int64(platformIDValue.(int))
		data.Platform = &platformID
	}

	roleIDValue, ok := d.GetOk("role_id")
	if ok {
		roleID := int64(roleIDValue.(int))
		data.Role = &roleID
	}

	localContextValue, ok := d.GetOk("local_context_data")
	if ok {
		var jsonObj any
		localContextBA := []byte(localContextValue.(string))
		if err := json.Unmarshal(localContextBA, &jsonObj); err == nil {
			data.LocalContextData = jsonObj
		}
	}

	data.Status = d.Get("status").(string)

	tags, diags := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	data.Tags = tags
	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := virtualization.NewVirtualizationVirtualMachinesCreateParams().WithData(&data)

	res, err := api.Virtualization.VirtualizationVirtualMachinesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return append(resourceNetboxVirtualMachineRead(ctx, d, m), diags...)
}

func resourceNetboxVirtualMachineRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := virtualization.NewVirtualizationVirtualMachinesReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationVirtualMachinesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*virtualization.VirtualizationVirtualMachinesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	vm := res.GetPayload()

	d.Set("name", vm.Name)

	if vm.Cluster != nil {
		d.Set("cluster_id", vm.Cluster.ID)
	} else {
		d.Set("cluster_id", nil)
	}

	if vm.PrimaryIp4 != nil {
		d.Set("primary_ipv4", vm.PrimaryIp4.ID)
	} else {
		d.Set("primary_ipv4", nil)
	}

	if vm.PrimaryIp6 != nil {
		d.Set("primary_ipv6", vm.PrimaryIp6.ID)
	} else {
		d.Set("primary_ipv6", nil)
	}

	if vm.Tenant != nil {
		d.Set("tenant_id", vm.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if vm.Device != nil {
		d.Set("device_id", vm.Device.ID)
	} else {
		d.Set("device_id", nil)
	}

	if vm.Role != nil {
		d.Set("role_id", vm.Role.ID)
	} else {
		d.Set("role_id", nil)
	}

	if vm.Platform != nil {
		d.Set("platform_id", vm.Platform.ID)
	} else {
		d.Set("platform_id", nil)
	}

	if vm.Role != nil {
		d.Set("role_id", vm.Role.ID)
	} else {
		d.Set("role_id", nil)
	}

	if vm.Site != nil {
		d.Set("site_id", vm.Site.ID)
	} else {
		d.Set("site_id", nil)
	}

	if vm.LocalContextData != nil {
		if jsonArr, err := json.Marshal(vm.LocalContextData); err == nil {
			d.Set("local_context_data", string(jsonArr))
		}
	} else {
		d.Set("local_context_data", nil)
	}

	d.Set("comments", vm.Comments)
	d.Set("description", vm.Description)
	vcpus := vm.Vcpus
	if vcpus != nil {
		d.Set("vcpus", vm.Vcpus)
	} else {
		d.Set("vcpus", nil)
	}
	d.Set("memory_mb", vm.Memory)
	d.Set("disk_size_gb", vm.Disk)
	if vm.Status != nil {
		d.Set("status", vm.Status.Value)
	} else {
		d.Set("status", nil)
	}
	d.Set(tagsKey, getTagListFromNestedTagList(vm.Tags))

	cf := getCustomFields(vm.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	return diags
}

func resourceNetboxVirtualMachineUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableVirtualMachineWithConfigContext{}

	name := d.Get("name").(string)
	data.Name = &name

	clusterIDValue, ok := d.GetOk("cluster_id")
	if ok {
		clusterID := int64(clusterIDValue.(int))
		data.Cluster = &clusterID
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	deviceIDValue, ok := d.GetOk("device_id")
	if ok {
		deviceID := int64(deviceIDValue.(int))
		data.Device = &deviceID
	}

	platformIDValue, ok := d.GetOk("platform_id")
	if ok {
		platformID := int64(platformIDValue.(int))
		data.Platform = &platformID
	}

	roleIDValue, ok := d.GetOk("role_id")
	if ok {
		roleID := int64(roleIDValue.(int))
		data.Role = &roleID
	}

	memoryMbValue, ok := d.GetOk("memory_mb")
	if ok {
		memoryMb := int64(memoryMbValue.(int))
		data.Memory = &memoryMb
	}

	vcpusValue, ok := d.GetOk("vcpus")
	if ok {
		vcpus := vcpusValue.(float64)
		data.Vcpus = &vcpus
	}

	diskSizeValue, ok := d.GetOk("disk_size_gb")
	if ok {
		diskSize := int64(diskSizeValue.(int))
		data.Disk = &diskSize
	}

	primaryIP4Value, ok := d.GetOk("primary_ipv4")
	if ok {
		primaryIP4 := int64(primaryIP4Value.(int))
		data.PrimaryIp4 = &primaryIP4
	}

	primaryIP6Value, ok := d.GetOk("primary_ipv6")
	if ok {
		primaryIP6 := int64(primaryIP6Value.(int))
		data.PrimaryIp6 = &primaryIP6
	}

	localContextValue, ok := d.GetOk("local_context_data")
	if ok {
		var jsonObj any
		localContextBA := []byte(localContextValue.(string))
		if err := json.Unmarshal(localContextBA, &jsonObj); err == nil {
			data.LocalContextData = jsonObj
		}
	}

	tags, diags := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	data.Tags = tags
	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	if d.HasChanges("comments") {
		// check if comment is set
		if commentsValue, ok := d.GetOk("comments"); ok {
			data.Comments = commentsValue.(string)
		} else {
			data.Comments = " "
		}
	}
	if d.HasChanges("description") {
		// check if description is set
		if descriptionValue, ok := d.GetOk("description"); ok {
			data.Description = descriptionValue.(string)
		} else {
			data.Description = " "
		}
	}

	// if d.HasChanges("status") {
	if status, ok := d.GetOk("status"); ok {
		data.Status = status.(string)
	}
	//}

	params := virtualization.NewVirtualizationVirtualMachinesUpdateParams().WithID(id).WithData(&data)

	_, err := api.Virtualization.VirtualizationVirtualMachinesUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return append(resourceNetboxVirtualMachineRead(ctx, d, m), diags...)
}

func resourceNetboxVirtualMachineDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationVirtualMachinesDeleteParams().WithID(id)

	_, err := api.Virtualization.VirtualizationVirtualMachinesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*virtualization.VirtualizationVirtualMachinesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	return diags
}
