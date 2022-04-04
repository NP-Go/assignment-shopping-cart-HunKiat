package main

// reports.go handles reports print functions

import (
	"fmt"
	"strconv"
)

// for CateogryName, ItemName, Qty and UnitCost
const printLineItem = "%s: %s - Item: %d Unit Cost: %g\n"

func printShoppingList() {
	fmt.Println("Shopping List Contents:")
	for key, value := range items {
		fmt.Printf(printLineItem, Category[value.Category], key, value.Quantity, value.Unit_Cost)
	}
}

func printShoppingRpts() {
	var input2 int
	rChoices := makeMenuRange(1, len(reportMenu))
	displayRptSubmenu() // see frontEnd.go
	fmt.Scanln(&input2)
	for !validSelection(rChoices, input2) {
		displayRptSubmenu()
		fmt.Scanln(&input2)
	}
	if input2 == 1 {
		categoryCost := make([]float64, len(Category))
		for k := range Category {
			for _, value := range items {
				if value.Category == k {
					categoryCost[k] = categoryCost[k] + value.Unit_Cost*float64(value.Quantity)
				}
			}
		}
		for l := range Category {
			fmt.Println(Category[l] + " cost : " + fmt.Sprintf("%g", categoryCost[l]))
		}
		pauseToRead()
		printShoppingRpts()
	} else if input2 == 2 {
		// rpt by category list sorted
		fmt.Println("List by Category")
		for k := range Category {
			for key, value := range items {
				if value.Category == k {
					fmt.Printf(printLineItem, Category[value.Category], key, value.Quantity, value.Unit_Cost)
				}
			}
		}
		pauseToRead()
		printShoppingRpts()
	} else if input2 == 3 {
		// fmt.Println("Returning back to Main Menu...")
		// pauseToRead()
		return
	}
}

func printDataInMem() {
	fmt.Println("Print Current Data.")
	if len(items) != 0 {
		for key, value := range items {
			fmt.Println(key + " - {" + strconv.Itoa(value.Category) + " " + strconv.Itoa(value.Quantity) + " " + fmt.Sprintf("%g", value.Unit_Cost) + "}")
		}
	} else {
		fmt.Println("No data found")
	}
}
