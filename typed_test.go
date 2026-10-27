package cml_test

import (
	"os"
	"testing"

	"github.com/grespyrad/go-cml"
)

// TestTypedModel проверяет контракт go-cml.
func TestTypedModel(t *testing.T) {
	t.Parallel()

	source, err := os.ReadFile("testdata/reference/service.cml")
	if err != nil {
		t.Fatal(err)
	}

	var model cml.ContextMappingModel

	if decodeErr := cml.Unmarshal(source, &model); decodeErr != nil {
		t.Fatal(decodeErr)
	}

	if len(model.BoundedContexts) != 2 || model.BoundedContexts[0].Name != "ExternalService" {
		t.Fatalf("%+v", model.BoundedContexts)
	}

	if model.BoundedContexts[0].Aggregates[0].DomainObjects[0].Entity.Name != "ExternalRecord" {
		t.Fatal("typed entity missing")
	}

	encoded, err := cml.Marshal(&model)
	if err != nil {
		t.Fatal(err)
	}

	var again cml.ContextMappingModel

	if err := cml.Unmarshal(encoded, &again); err != nil {
		t.Fatalf("%v\n%s", err, encoded)
	}

	if len(again.BoundedContexts) != len(model.BoundedContexts) {
		t.Fatal("round trip mismatch")
	}
}

// TestCreateTypedModel проверяет контракт go-cml.
//
//nolint:gosmopolitan // Unicode контракт требует Han и emoji в fixture.
func TestCreateTypedModel(t *testing.T) {
	t.Parallel()

	model := &cml.ContextMappingModel{
		BoundedContexts: []*cml.BoundedContext{
			{
				Name:                  "ServiceIntegration",
				DomainVisionStatement: "Сервис 文 🔧",
				Aggregates: []*cml.Aggregate{
					{
						Name: "A",
						DomainObjects: []*cml.SimpleDomainObject{
							{Entity: &cml.Entity{Name: "ServiceIntegration", AggregateRoot: true}},
						},
					},
				},
			},
		},
	}

	data, err := cml.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	if r := cml.Validate("created.cml", data); !r.Valid {
		t.Fatalf("%+v\n%s", r.Diagnostics, data)
	}
}
