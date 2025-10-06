package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/jianyuan/terraform-provider-porkbun/internal/acctest"
)

func TestAccDnssecRecordsDataSource(t *testing.T) {
	rn := "data.porkbun_dnssec_record.test"
	config := testAccDnssecRecordsDataSourceConfig()
	t.Logf("Test config:\n%s", config)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnssecRecordsDataSourceConfig(),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(rn, tfjsonpath.New("domain"), knownvalue.StringExact(acctest.TestDomain)),
					statecheck.ExpectKnownValue(rn, tfjsonpath.New("records"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func testAccDnssecRecordsDataSourceConfig() string {
	return fmt.Sprintf(`
data "porkbun_dnssec_record" "test" {
    domain = "%[1]s"
}
`, acctest.TestDomain)
}
