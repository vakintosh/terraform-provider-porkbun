package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DnssecResourceModel struct {
	ID      types.String `tfsdk:"id"`
	Domain  types.String `tfsdk:"domain"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

func NewDnssecResource() resource.Resource {
	return &DnssecResource{}
}

var _ resource.Resource = &DnssecResource{}

type DnssecResource struct {
	baseResource
}

func (r *DnssecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnssec_record"
}

func (r *DnssecResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage a DNSSEC record for your domain",

		Attributes: map[string]schema.Attribute{
			"domain": schema.StringAttribute{
				Required: true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Wheter or not to activate dnssec record",
				Required:            true,
			},
		},
	}
}
