package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jianyuan/terraform-provider-porkbun/internal/apiclient"
)

type DnssecRecordDataSourceModel struct {
	Domain  types.String        `tfsdk:"domain"`
	Records []DnssecRecordModel `tfsdk:"records"`
}

var _ datasource.DataSource = &DnssecRecordDataSource{}

type DnssecRecordDataSource struct {
	baseDataSource
}

func NewDnssecRecordDataSource() datasource.DataSource {
	return &DnssecRecordDataSource{}
}

func (d *DnssecRecordDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnssec_record"
}

func (d *DnssecRecordDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a list of DNSSEC records for a given domain.",
		Attributes: map[string]schema.Attribute{
			"domain": schema.StringAttribute{
				Required: true,
			},
			"records": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"key_tag": schema.StringAttribute{
							Computed: true,
						},
						"alg": schema.StringAttribute{
							Computed: true,
						},
						"digest_type": schema.StringAttribute{
							Computed: true,
						},
						"digest": schema.StringAttribute{
							Computed: true,
						},
						"max_sig_life": schema.StringAttribute{
							Computed: true,
						},
						"key_data_flags": schema.StringAttribute{
							Computed: true,
						},
						"key_data_protocol": schema.StringAttribute{
							Computed: true,
						},
						"key_data_algo": schema.StringAttribute{
							Computed: true,
						},
						"key_data_pub_key": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *DnssecRecordDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DnssecRecordDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := d.client.DnssecGetRecordsWithResponse(
		ctx,
		data.Domain.ValueString(),
		apiclient.DnssecGetRecordsJSONRequestBody{
			Apikey:       d.apiKey,
			Secretapikey: d.secretKey,
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read, got error: %s", err))
		return
	} else if httpResp.StatusCode() != http.StatusOK || httpResp.JSON200 == nil || httpResp.JSON200.Status != "SUCCESS" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read, got status code %d: %s", httpResp.StatusCode(), string(httpResp.Body)))
		return
	}

	records := make([]DnssecRecordModel, 0, len(httpResp.JSON200.Records))
	for id, rec := range httpResp.JSON200.Records {
		var model DnssecRecordModel
		model.Fill(ctx, rec)

		model.KeyTag = types.StringValue(id)
		model.Id = types.StringValue(id)

		records = append(records, model)
	}
	data.Records = records

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
