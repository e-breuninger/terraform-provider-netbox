package netbox

import (
	"context"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxVirtualMachine() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxVirtualMachineCreate,
		ReadContext:   resourceNetboxVirtualMachineRead,
		UpdateContext: resourceNetboxVirtualMachineUpdate,
		DeleteContext: resourceNetboxVirtualMachineDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"platform_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"memory_mb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"disk_size_gb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Set:      schema.HashString,
			},
			"primary_ipv4": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxVirtualMachineCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	clusterID := int64(d.Get("cluster_id").(int))

	data := models.WritableVirtualMachineWithConfigContext{
		Name:    &name,
		Cluster: &clusterID,
	}

	comments := d.Get("comments").(string)
	data.Comments = comments

	vcpusValue, ok := d.GetOk("vcpus")
	if ok {
		vcpus := int64(vcpusValue.(int))
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

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	params := virtualization.NewVirtualizationVirtualMachinesCreateParams().WithData(&data)

	res, err := api.Virtualization.VirtualizationVirtualMachinesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVirtualMachineUpdate(ctx, d, m)
}

func resourceNetboxVirtualMachineRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := virtualization.NewVirtualizationVirtualMachinesReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationVirtualMachinesRead(params, nil)
	if err != nil {
		errorcode := err.(*virtualization.VirtualizationVirtualMachinesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("cluster_id", res.GetPayload().Cluster.ID)

	if res.GetPayload().PrimaryIp4 != nil {
		d.Set("primary_ipv4", res.GetPayload().PrimaryIp4.ID)
	} else {
		d.Set("primary_ipv4", nil)
	}

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if res.GetPayload().Platform != nil {
		d.Set("platform_id", res.GetPayload().Platform.ID)
	} else {
		d.Set("platform_id", nil)
	}

	if res.GetPayload().Role != nil {
		d.Set("role_id", res.GetPayload().Role.ID)
	} else {
		d.Set("role_id", nil)
	}

	d.Set("comments", res.GetPayload().Comments)
	d.Set("vcpus", res.GetPayload().Vcpus)
	d.Set("memory_mb", res.GetPayload().Memory)
	d.Set("disk_size_gb", res.GetPayload().Disk)
	d.Set("tags", getTagListFromNestedTagList(res.GetPayload().Tags))
	return diags
}

func resourceNetboxVirtualMachineUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableVirtualMachineWithConfigContext{}

	name := d.Get("name").(string)
	data.Name = &name

	clusterID := int64(d.Get("cluster_id").(int))
	data.Cluster = &clusterID

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
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
		//	} else {
		//		memorymb := int64(0)
		//		data.memory = &memorymb
	}

	vcpusValue, ok := d.GetOk("vcpus")
	if ok {
		vcpus := int64(vcpusValue.(int))
		data.Vcpus = &vcpus
		//	} else {
		//		vcpus := int64(0)
		//		data.Vcpus = &vcpus
	}

	diskSizeValue, ok := d.GetOk("disk_size_gb")
	if ok {
		diskSize := int64(diskSizeValue.(int))
		data.Disk = &diskSize
	}

	commentsValue, ok := d.GetOk("comments")
	if ok {
		comments := commentsValue.(string)
		data.Comments = comments
	} else {
		comments := " "
		data.Comments = comments
	}

	primaryIPValue, ok := d.GetOk("primary_ipv4")
	if ok {
		primaryIP := int64(primaryIPValue.(int))
		data.PrimaryIp4 = &primaryIP
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	if d.HasChanges("comments") {
		// check if comment is set
		commentsValue, ok := d.GetOk("comments")
		comments := ""
		if !ok {
			// Setting an space string deletes the comment
			comments = " "
		} else {
			comments = commentsValue.(string)
		}
		data.Comments = comments
	}

	params := virtualization.NewVirtualizationVirtualMachinesUpdateParams().WithID(id).WithData(&data)

	_, err := api.Virtualization.VirtualizationVirtualMachinesUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceNetboxVirtualMachineRead(ctx, d, m)
}

func resourceNetboxVirtualMachineDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationVirtualMachinesDeleteParams().WithID(id)

	_, err := api.Virtualization.VirtualizationVirtualMachinesDelete(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
