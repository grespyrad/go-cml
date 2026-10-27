package cml_test

import (
	"testing"

	"github.com/grespyrad/go-cml"
)

// TestEscapedDatatypeRoundTrip сохраняет escaped сегменты JavaIdentifier.
func TestEscapedDatatypeRoundTrip(t *testing.T) {
	t.Parallel()

	source := []byte(
		"BoundedContext ServiceIntegration { Module m { basePackage = foo.^type " +
			"Aggregate Items { Entity E { foo.^type field } } } }",
	)

	var model cml.ContextMappingModel

	if err := cml.Unmarshal(source, &model); err != nil {
		t.Fatal(err)
	}

	encoded, err := cml.Marshal(&model)
	if err != nil {
		t.Fatal(err)
	}

	if result := cml.Validate("test.cml", encoded); !result.Valid {
		t.Fatal(result.Diagnostics)
	}
}
