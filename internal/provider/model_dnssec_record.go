package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jianyuan/terraform-provider-porkbun/internal/apiclient"
)

type DnssecRecordModel struct {
	Domain          types.String `tfsdk:"domain"`
	KeyTag          types.String `tfsdk:"key_tag"`
	Alg             types.String `tfsdk:"alg"`
	DigestType      types.String `tfsdk:"digest_type"`
	Digest          types.String `tfsdk:"digest"`
	MaxSigLife      types.String `tfsdk:"max_sig_life"`
	KeyDataFlags    types.String `tfsdk:"key_data_flags"`
	KeyDataProtocol types.String `tfsdk:"key_data_protocol"`
	KeyDataAlgo     types.String `tfsdk:"key_data_algo"`
	KeyDataPubKey   types.String `tfsdk:"key_data_pub_key"`
	ID              types.String `tfsdk:"id"` // computed: domain:keyTag
}

func (m *DnssecRecordModel) Fill(ctx context.Context, record apiclient.DnssecGetRecord) (diags diag.Diagnostics) {
	m.Domain = types.StringPointerValue(record.Domain)
	m.KeyTag = types.StringValue(record.KeyTag)
	m.Alg = types.StringValue(record.Alg)
	m.DigestType = types.StringValue(record.DigestType)
	m.Digest = types.StringValue(record.Digest)
	m.MaxSigLife = types.StringPointerValue(record.MaxSigLife)
	m.KeyDataFlags = types.StringPointerValue(record.KeyDataFlags)
	m.KeyDataProtocol = types.StringPointerValue(record.KeyDataProtocol)
	m.KeyDataAlgo = types.StringPointerValue(record.KeyDataAlgo)
	m.KeyDataPubKey = types.StringPointerValue(record.KeyDataPubKey)
	m.ID = types.StringValue(fmt.Sprintf("%s", record.KeyTag))
	return
}
