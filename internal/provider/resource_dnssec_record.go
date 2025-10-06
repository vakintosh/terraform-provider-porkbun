package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jianyuan/terraform-provider-porkbun/internal/apiclient"
)

type DnssecRecordResourceModel struct {
	Domain types.String `tfsdk:"domain"`
	KeyTag types.String `tfsdk:"key_tag"`
	DnssecRecordModel
}

func NewDnssecRecordResource() resource.Resource {
	return &DnssecRecordResource{}
}

var _ resource.Resource = &DnssecRecordResource{}

type DnssecRecordResource struct {
	baseResource
}

func (r *DnssecRecordResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnssec_record"
}

func (r *DnssecRecordResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DNSSEC record for a domain in Porkbun.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The domain name this DNSSEC record belongs to.",
			},
			"key_tag": schema.StringAttribute{
				Required:    true,
				Description: "The DNSSEC Key Tag.",
			},
			"alg": schema.StringAttribute{
				Required:    true,
				Description: "The DS Data Algorithm used for the DNSSEC record.",
			},
			"digest_type": schema.StringAttribute{
				Required:    true,
				Description: "The digest type for the DNSSEC record.",
			},
			"digest": schema.StringAttribute{
				Required:    true,
				Description: "The digest value of the DNSSEC record.",
			},
			"max_sig_life": schema.StringAttribute{
				Optional:    true,
				Description: "Maximum signature lifetime (optional, often unused).",
			},
			"key_data_flags": schema.StringAttribute{
				Optional:    true,
				Description: "Key data flags (optional).",
			},
			"key_data_protocol": schema.StringAttribute{
				Optional:    true,
				Description: "Key data protocol (optional).",
			},
			"key_data_algo": schema.StringAttribute{
				Optional:    true,
				Description: "Key data algorithm (optional).",
			},
			"key_data_pub_key": schema.StringAttribute{
				Optional:    true,
				Description: "Key data public key (optional).",
			},
		},
	}
}

func (r *DnssecRecordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnssecRecordResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createHttpResp, err := r.client.DnssecCreateRecordWithResponse(
		ctx,
		data.Domain.ValueString(),
		apiclient.DnssecCreateRecordJSONRequestBody{
			Apikey:          r.apiKey,
			Secretapikey:    r.secretKey,
			KeyTag:          data.KeyTag.ValueString(),
			Alg:             data.Alg.ValueString(),
			DigestType:      data.DigestType.ValueString(),
			Digest:          data.Digest.ValueString(),
			MaxSigLife:      data.MaxSigLife.ValueString(),
			KeyDataFlags:    data.KeyDataFlags.ValueString(),
			KeyDataProtocol: data.KeyDataProtocol.ValueString(),
			KeyDataAlgo:     data.KeyDataAlgo.ValueString(),
			KeyDataPubKey:   data.KeyDataPubKey.ValueString(),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create DNSSEC record: %s", err))
		return
	} else if createHttpResp.StatusCode() != http.StatusOK || createHttpResp.JSON200 == nil || createHttpResp.JSON200.Status != "SUCCESS" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create DNSSEC record, got status code %d: %s", createHttpResp.StatusCode(), string(createHttpResp.Body)))
		return
	}
	data.Id = types.StringValue(data.Domain.ValueString() + "-" + data.KeyTag.ValueString())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssecRecordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnssecRecordResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.DnssecGetRecordsWithResponse(
		ctx,
		data.Domain.ValueString(),
		apiclient.DnssecGetRecordsJSONRequestBody{
			Apikey:       r.apiKey,
			Secretapikey: r.secretKey,
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read DNSSEC records: %s", err))
		return
	} else if httpResp.StatusCode() != http.StatusOK || httpResp.JSON200 == nil || httpResp.JSON200.Status != "SUCCESS" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read DNSSEC records, got status code %d: %s", httpResp.StatusCode(), string(httpResp.Body)))
		return
	}

	resp.Diagnostics.Append(data.Fill(ctx, httpResp.JSON200.Records[data.KeyTag.ValueString()])...)
	if resp.Diagnostics.HasError() {
		return
	}

	// var parsed struct {
	// 	Status  string                            `json:"status"`
	// 	Records map[string]apiclient.DnssecRecord `json:"records"`
	// }
	// if err := json.Unmarshal(httpResp.Body, &parsed); err != nil {
	// 	resp.Diagnostics.AddError("Unmarshal Error", fmt.Sprintf("Failed to decode response: %s", err))
	// 	return
	// }

	// record, ok := parsed.Records[data.KeyTag.ValueString()]
	// if !ok {
	// 	resp.State.RemoveResource(ctx)
	// 	return
	// }

	// data.Alg = types.StringValue(record.Alg)
	// data.DigestType = types.StringValue(record.DigestType)
	// data.Digest = types.StringValue(record.Digest)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssecRecordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Update Not Supported",
		"Updating DNSSEC records is not supported by the Porkbun API. Please delete and recreate the resource if changes are needed.",
	)
}

func (r *DnssecRecordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnssecRecordResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	deleteResp, err := r.client.DnssecDeleteRecordByKeyTagWithResponse(
		ctx,
		data.Domain.ValueString(),
		data.KeyTag.ValueString(),
		apiclient.DnssecDeleteRecordJSONRequestBody{
			Apikey:       r.apiKey,
			Secretapikey: r.secretKey,
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete DNSSEC record: %s", err))
		return
	} else if deleteResp.StatusCode() != http.StatusOK || deleteResp.JSON200 == nil || deleteResp.JSON200.Status != "SUCCESS" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete DNSSEC record, got status code %d: %s", deleteResp.StatusCode(), string(deleteResp.Body)))
		return
	}
}
