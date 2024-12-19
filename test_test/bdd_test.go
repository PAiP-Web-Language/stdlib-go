package test_test

import (
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/test"
)

func TestBdd1(t *testing.T) {
	test.BDD.RunFeatureSet(
		t,
		test.BDD.Feature(
			"BDD Testing",
			test.BDD.Description("Description should do nothing"),
			test.BDD.When(
				"When 'When' function works",
				func(ti test.TI) {
					test.Assert(t, true, "When should work")
				},
			),
		),
	)
}
