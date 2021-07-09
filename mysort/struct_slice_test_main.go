package main

import (
	"fmt"

	"github.com/user/mysort"
)

type coin struct {
	Currency string
	Diameter int
	Value    float32
}

func main() {
	tenRubles := coin{Currency: "rubles", Diameter: 10, Value: 10.0}
	halfEuro := coin{Currency: "euro", Diameter: 5, Value: 0.5}
	fiveEuro := coin{Currency: "euro", Diameter: 40, Value: 5.0}
	wallet := []coin{tenRubles, halfEuro, fiveEuro}

	backup := wallet
	fmt.Printf("wallet: %+v\n%v - orig\n", wallet[0], wallet)
	mysort.SortStructs(&wallet, "Currency", true)
	fmt.Printf("%v - asc by \"Currency\"\n", wallet)
	wallet = backup
	mysort.SortStructs(&wallet, "Currency", false)
	fmt.Printf("%v - desc by \"Currency\"\n\n", wallet)
	wallet = backup
	mysort.SortStructs(&wallet, "Diameter", true)
	fmt.Printf("%v - asc by \"Diameter\"\n", wallet)
	wallet = backup
	mysort.SortStructs(&wallet, "Diameter", false)
	fmt.Printf("%v - desc by \"Diameter\"\n\n", wallet)
	wallet = backup
	mysort.SortStructs(&wallet, "Value", true)
	fmt.Printf("%v - asc by \"Value\"\n", wallet)
	wallet = backup
	mysort.SortStructs(&wallet, "Value", false)
	fmt.Printf("%v - desc by \"Value\"\n\n", wallet)
}
