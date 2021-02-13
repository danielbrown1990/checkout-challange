package main

import (
	"checkout-challange/checkout"
	"testing"
)

func Test_SingleItems(t *testing.T) {
	//Arrange
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
		sut := checkout.NewCheckout()
		sut.Scan(test.SKU)
		got := sut.Total()

		//Assert
		if test.want != got {
			t.Fatalf("incorrect total want(%v) have (%v)", test.want, got)
		}
	}
}

func Test_MultipleItems(t *testing.T) {
	//Arrange
	tests := []struct {
		SKUs []string
		want int
	}{
		{[]string{"A", "B", "C", "D"}, 115},
	}

	for _, test := range tests {
		//Act
		sut := checkout.NewCheckout()
		for _, SKU := range test.SKUs {
			sut.Scan(SKU)
		}

		got := sut.Total()

		//Assert
		if test.want != got {
			t.Fatalf("incorrect total want(%v) have (%v)", test.want, got)
		}
	}
}
