package test

import (
	"fmt"
	"testing"
)

var BDD = BDDType{}

type BDDType struct {
}

// RunFeatureSet is main function used for running BDD tests
// Usage example:
// func TestFeatures(t *testing.T) {
//  BDD.RunFeatureSet(t,
//      BDD.Feature("Feature 1",
//          BDD.Description("Description 1"),
//          BDD.When("When 1", func(t test.TI) {
//              t.Log("Log 1")
//          }),
//          BDD.Then("Then 1",
//              func(t test.TI) {
//              t.Log("Log 1")
//          })
//      )
//  )
// }
func (b *BDDType)RunFeatureSet(t *testing.T, features ...func(TI)) {
	ti := WrapT(t)
	defer ti.Defer()
	for _, f := range features {
		f(ti)
	}
}

// Feature is function that creates new feature to test
// It takes name and subelements (Like Description, Rule, Given, When etc.)
func (b *BDDType)Feature(name string, subelements ...func(TI)) func(t TI) {
	return func(t TI) {
		t.Helper()
		t.Run(
			fmt.Sprintf("Feature: %s", name),
			func(rt *testing.T) {
				rt.Helper()
				rti := WrapT(rt)
				for _, f := range subelements {
					f(rti)
				}
			},
		)
	}
}

// Description is function that is used just to note some additional information
// This function is meant to be used only to note some detail in tests but it does not do anything
func (b *BDDType)Description(name string) func(t TI) {
	return func(t TI) {}
}

func (b *BDDType)Rule(name string, body func(TI)) func(TI) {
	return body
}

func (b *BDDType)Scenario(name string, subelements ...func(TI)) func(TI) {
	return func(t TI) {
		t.Helper()
		t.Run(
			fmt.Sprintf("Scenario: %s", name),
			func(rt *testing.T) {
				rt.Helper()
				rti := WrapT(rt)
				for _, f := range subelements {
					f(rti)
				}
			},
		)
	}
}

func (b *BDDType)Given(name string, body func(TI)) func(TI) {
	return body
}

func (b *BDDType)When(name string, body func(TI)) func(TI) {
	return body
}

func (b *BDDType)Then(name string, body func(TI)) func(TI) {
	return body
}

func (b *BDDType)And(name string, body func(TI)) func(TI) {
	return body
}

func (b *BDDType)But(name string, body func(TI)) func(TI) {
	return body
}

func (b *BDDType)Background(name string, subelements ...func(TI)) func(TI) {
	return func(t TI) {
		t.Helper()
		t.Run(
			fmt.Sprintf("Background: %s", name),
			func(rt *testing.T) {
				rt.Helper()
				rti := WrapT(rt)
				for _, f := range subelements {
					f(rti)
				}
			},
		)
	}
}

// func (b *BDDType)ScenarioOutline(name string, body func(TI)) func(TI) {
// }
//
// func (b *BDDType)ScenarioTemplate(name string, body func(TI)) func(TI) {
// }
//
// func (b *BDDType)Examples(name string, body func(TI)) func(TI) {
// }
//
// func (b *BDDType)Scenarios(name string, body func(TI)) func(TI) {
// }
