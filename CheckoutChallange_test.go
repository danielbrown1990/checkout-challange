package main

import (
	"checkout-challange/catalogue"
	"checkout-challange/checkout"
	"testing"
)

func Test_SingleItems(t *testing.T) {
	//Arrange
	cat := catalogue.NewTestCatalogue()
	tests := []struct {
		SKU  string
		want int
	}{
		{"A", 50},
		{"B", 30},
		{"C", 20},
		{"D", 15},
	}

	for _, test := range tests {
		//Act
		sut := checkout.NewCheckout(cat)
		err := sut.Scan(test.SKU)
		if err != nil {
			t.Fatalf("failed to scan item: %v", err)
		}

		got := sut.Total()

		//Assert
		if test.want != got {
			t.Fatalf("incorrect total want(%v) have (%v)", test.want, got)
		}
	}
}

func Test_MultipleItems(t *testing.T) {
	//Arrange
	cat := catalogue.NewTestCatalogue()
	tests := []struct {
		SKUs []string
		want int
	}{
		{[]string{"A", "B", "C", "D"}, 115},
		{[]string{"A", "A", "A"}, 130},
		{[]string{"B", "B"}, 45},
		{[]string{"A", "B", "C", "A", "B", "D", "A"}, 210},
		{[]string{"A", "A", "A", "A"}, 180},
	}

	for _, test := range tests {
		//Act
		sut := checkout.NewCheckout(cat)
		for _, SKU := range test.SKUs {
			err := sut.Scan(SKU)
			if err != nil {
				t.Fatalf("failed to scan item: %v", err)
			}
		}

		got := sut.Total()

		//Assert
		if test.want != got {
			t.Fatalf("incorrect total want(%v) have (%v)", test.want, got)
		}
	}
}

func Test_ScanUnknownSKU_ReturnsError(t *testing.T) {
	//Arrange
	cat := catalogue.NewTestCatalogue()
	sut := checkout.NewCheckout(cat)

	//Act
	err := sut.Scan("X")

	//Assert
	if err == nil {
		t.Fatal("no error returned for scan")
	}
}
