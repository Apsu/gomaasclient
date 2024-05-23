package entity

import (
	"testing"

	"github.com/apsu/gomaasclient/test/helper"
)

func TestDomaint(t *testing.T) {
	domains := new([]Domain)
	if err := helper.TestdataFromJSON("maas/domains.json", domains); err != nil {
		t.Fatal(err)
	}
}
