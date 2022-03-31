package main

import (
	"fmt"
	"strconv"
	"time"
)

func charLen(s string) int {
	return len(s)
}

// this creates a slice
func makeMenuRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

type item struct {
	Category  int
	Quantity  int
	Unit_Cost float64
}

func findCategory(C []string, c string) (int, bool) {
	for i := range C {
		if C[i] == c {
			return i, true
		}
	}
	return -1, false
}

func findChoiceInSlice(m []int, c int) bool {
	for i := range m {
		if m[i] == c {
			return true
		}
	}
	return false
}

var items map[string]item

var mainMenu = []string{
	"1. View entire shopping list.",
	"2. Generate Shopping List Report",
	"3. Add Items",
	"4. Modify Items",
	"5. Delete Item",
	"6. Print Current Data",
	"7. Add New Category Name",
}

var reportMenu = []string{
	"1. Total Cost of each category.",
	"2. List of item by category.",
	"3. Main Menu.",
}

func init() {
	// populating with some initial values
	items["Fork"] = item{0, 4, 3.0}
	items["Plates"] = item{0, 4, 3.0}
	items["Cups"] = item{0, 5, 3.0}
	items["Bread"] = item{1, 2, 2.0}
	items["Cake"] = item{1, 3, 1.0}
	items["Coke"] = item{2, 5, 2.0}
	items["Sprite"] = item{2, 5, 2.0}
}

func main() {
	//insert code here
	var input1 int
	choices := makeMenuRange(1, len(mainMenu))
	Category := []string{"Household", "Food", "Drinks"}

	fmt.Println("Shopping List Application")
	fmt.Println("=========================")
	for i := range mainMenu {
		fmt.Println(mainMenu[i])
	}
	fmt.Println("Selete your choice:")
	fmt.Scanln(&input1)

	for !findChoiceInSlice(choices, input1) {
		fmt.Println("Shopping List Application")
		fmt.Println("=========================")
		for i := range mainMenu {
			fmt.Println(mainMenu[i])
		}
		fmt.Println("Selete your choice:")
		fmt.Scanln(&input1)
	}

	switch input1 {
	case 1:
		fmt.Println("Shopping List Contents:")
		for key, value := range items {
			fmt.Println(Category[value.Category] + ": " + key + " - Item: " + strconv.Itoa(value.Quantity) + " Unit Cost: " + fmt.Sprintf("%g", value.Unit_Cost))
		}
	case 2:
		var input2 int
		rChoices := makeMenuRange(1, len(reportMenu))
		fmt.Println("Generate Report")
		for j := range mainMenu {
			fmt.Println(mainMenu[j])
		}
		fmt.Scanln(&input2)
		for !findChoiceInSlice(rChoices, input2) {
			fmt.Println("Generate Report")
			for j := range mainMenu {
				fmt.Println(mainMenu[j])
			}
			fmt.Scanln(&input2)
			if input2 == 1 {
				fmt.Println("Total cost by Category.")
				categoryCost := make([]float64, 2, 3)
				for _, value := range items {
					for k := range Category {
						if value.Category == k {
							categoryCost[k] = categoryCost[k] + value.Unit_Cost
						}
					}
				}
				for l := range Category {
					fmt.Println(Category[l] + " cost : " + fmt.Sprintf("%g", categoryCost[l]))
				}
			} else if input2 == 2 {

			} else {
				break
			}
		}
	case 3:
		fmt.Println("3. Add Items")
	case 4:
		fmt.Println("4. Modify Items")
	case 5:
		fmt.Println("5. Delete Item")
	case 6:
		fmt.Println("6. Print Current Data")
	case 7:
		var newCategory string
		timeStarted := time.Now()
		fmt.Println("Add New Category Name")
		fmt.Println("What is the New Category Name to add?")
		fmt.Scanln(&newCategory)
		t := time.Now()
		elapsed := t.Sub(timeStarted)
		i, found := findCategory(Category, newCategory)
		if !found && elapsed < 5 {
			//Append new category to Category
			Category = append(Category, newCategory)
			fmt.Println("New category: " + newCategory + " added at index " + strconv.Itoa(len(Category)))
		} else if found && elapsed < 5 {
			fmt.Println("Category: " + newCategory + " already exists at index " + strconv.Itoa(i))
		} else {
			fmt.Println("No input found")
		}
	}
}
