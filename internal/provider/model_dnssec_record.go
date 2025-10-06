package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jianyuan/terraform-provider-porkbun/internal/apiclient"
)

type DnssecRecordModel struct {
	Id              types.String `tfsdk:"id"`
	Alg             types.String `tfsdk:"alg"`
	DigestType      types.String `tfsdk:"digest_type"`
	Digest          types.String `tfsdk:"digest"`
	MaxSigLife      types.String `tfsdk:"max_sig_life"`
	KeyDataFlags    types.String `tfsdk:"key_data_flags"`
	KeyDataProtocol types.String `tfsdk:"key_data_protocol"`
	KeyDataAlgo     types.String `tfsdk:"key_data_algo"`
	KeyDataPubKey   types.String `tfsdk:"key_data_pub_key"`
}

func coalesceString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (m *DnssecRecordModel) Fill(ctx context.Context, record apiclient.DnssecGetRecord) (diags diag.Diagnostics) {
	m.Id = types.StringValue(record.Id)
	m.Alg = types.StringValue(record.Alg)
	m.DigestType = types.StringValue(record.DigestType)
	m.Digest = types.StringValue(record.Digest)
	m.MaxSigLife = types.StringValue(coalesceString(record.MaxSigLife))
	m.KeyDataFlags = types.StringValue(coalesceString(record.KeyDataFlags))
	m.KeyDataProtocol = types.StringValue(coalesceString(record.KeyDataProtocol))
	m.KeyDataAlgo = types.StringValue(coalesceString(record.KeyDataAlgo))
	m.KeyDataPubKey = types.StringValue(coalesceString(record.KeyDataPubKey))
	return
}
