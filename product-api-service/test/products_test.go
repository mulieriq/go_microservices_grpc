package test

import (
	"product-api-service/data"
	"testing"
)
//Test
func TestChecksValidation(t *testing.T) {
	pd := &data.Product{
		Name: "Eric",
		Price: 1,
		SKU: "abc-wer-eww",
	}
	err := pd.Validate()
	if err != nil {
		t.Fatal(err)
	}

}
