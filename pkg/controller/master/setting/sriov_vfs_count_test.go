package setting

import (
	"testing"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
)

func TestSRIOVVirtualFunctionCount(t *testing.T) {
	h := Handler{}
	sriovNumVfs := harvesterv1.Setting{
		Value: "4",
	}

	err := h.setSRIOVVirtualFunctionCount(&sriovNumVfs)
	if err != nil {
		t.Fatalf("%v", err)
	}

	sriovNumVfs = harvesterv1.Setting{
		Value: "not a number",
	}

	err = h.setSRIOVVirtualFunctionCount(&sriovNumVfs)
	if err == nil {
		t.Fatalf("%v", err)
	}
}
