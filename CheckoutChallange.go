package main

import (
	"bufio"
	"checkout-challange/catalogue"
	"checkout-challange/checkout"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	catalogueLocation = "cmd/testdata/catalogue.json"
)

func main() {
	flag.StringVar(&catalogueLocation, "catalogue_location", catalogueLocation, "flag to override default catalogue location")
	flag.Parse()

	cat, err := catalogue.NewCatalogue(catalogueLocation)
	if err != nil {
		panic(fmt.Errorf("Error reading catalogue configuration at %v", catalogueLocation))
	}

	check := checkout.NewCheckout(cat)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Enter SKU or type R to reset\n current total: %v\n>", check.Total())
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(fmt.Errorf("Error reading user input %v", err))
		}
		SKU := strings.ReplaceAll(input, "\n", "")

		if SKU == "R" {
			check = checkout.NewCheckout(cat)
			continue
		}

		err = check.Scan(SKU)
		if err != nil {
			fmt.Printf("Unknown SKU entered: %v\nenter SKU type R to reset\n current total: %v\n>", SKU, check.Total())
		}

	}
}
