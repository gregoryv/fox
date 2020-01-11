package doc

import "testing"

func Test_generate_diagrams(t *testing.T) {
	DesignOverview().SaveAs("design_overview.svg")
}
