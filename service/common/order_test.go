package common

import (
	"sort"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
)

func TestOrderValid(t *testing.T) {
	var order Order
	gofakeit.Struct(&order)
	err := validator.New().Struct(order)
	if err != nil {
		t.Fatalf("Order is valid, but validation fails: %v", err)
	}
}

func TestOrderInvalid(t *testing.T) {
	var order Order = Order{}
	err := validator.New().Struct(&order)
	if err == nil {
		t.Fatalf("Order is invalid, but validation succeeds")
	}
}

func TestKeys(t *testing.T) {
	sampleMap := map[string]int{
		"a": 2, "b": 3, "d": 0,
	}
	sampleKeys := []string{"a", "b", "d"}
	keys := keys(sampleMap)
	sort.Strings(keys)
	if !slices.Equal(keys, sampleKeys) {
		t.Fatal()
	}
}
