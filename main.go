package main

import (
	"fmt"
	"sort"
	"strconv"
)

type item struct {
	Category  int
	Quantity  int
	Unit_Cost float64
}

// returns the index of needle in the slice haystack, and true
// returns -1 and false if otherwise
func findCategory(C []string, c string) (int, bool) {
	for i := range C {
		if C[i] == c {
			return i, true
		}
	}
	return -1, false
}

var items map[string]item

func init() {
	items = make(map[string]item)
	items["Fork"] = item{0, 4, 3.0}
	items["Plates"] = item{0, 4, 3.0}
	items["Cups"] = item{0, 5, 3.0}
	items["Bread"] = item{1, 2, 2.0}
	items["Cake"] = item{1, 3, 1.0}
	items["Coke"] = item{2, 5, 2.0}
	items["Sprite"] = item{2, 5, 2.0}
}

func main() {
	var input1 int
	choices := makeMenuRange(1, len(mainMenu))
	Category := []string{"Household", "Food", "Drinks"}

	displayMainMenu()
	fmt.Scanln(&input1)
	for !validSelection(choices, input1) {
		displayMainMenu()
		fmt.Scanln(&input1)
	}

	for validSelection(choices, input1) {
		switch input1 {
		case 1:
			fmt.Println("Shopping List Contents:")
			for key, value := range items {
				fmt.Println(Category[value.Category] + ": " + key + " - Item: " + strconv.Itoa(value.Quantity) + " Unit Cost: " + fmt.Sprintf("%g", value.Unit_Cost))
			}
			pauseToRead(&input1)
		case 2:
			var input2 int
			rChoices := makeMenuRange(1, len(reportMenu))
			displayRptSubmenu()
			fmt.Scanln(&input2)
			for !validSelection(rChoices, input2) {
				displayRptSubmenu()
				fmt.Scanln(&input2)
			}
			for validSelection(rChoices, input2) {
				if input2 == 1 {
					categoryCost := make([]float64, len(Category))
					// cycle thru all the items...
					for _, value := range items {
						// and for each value.Category...
						// check to see if it matches with on
						for k := range Category {
							if value.Category == k {
								categoryCost[k] = categoryCost[k] + value.Unit_Cost
								break // break out of k loop
							}
						}
					}
					for l := range Category {
						fmt.Println(Category[l] + " cost : " + fmt.Sprintf("%g", categoryCost[l]))
					}
					pauseToRead(&input2)
					displayRptSubmenu()
					fmt.Scanln(&input2)
				} else if input2 == 2 {
					// rpt by category list sorted
					fmt.Println("List by Category")
					sort.Strings(Category)
					for k := range Category {
						for key, value := range items {
							if value.Category == k {
								fmt.Println(Category[value.Category] + ": " + key + " - Item: " + strconv.Itoa(value.Quantity) + " Unit Cost: " + fmt.Sprintf("%g", value.Unit_Cost))
							}
						}
					}
					pauseToRead(&input2)
					displayRptSubmenu()
					fmt.Scanln(&input2)
				} else if input2 == 3 {
					break
				}
			}
		case 3:
			var input31, input32, input35 string
			var input33 int
			var input34 float64
			fmt.Println("Name the item you wish to add...?")
			fmt.Scanln(&input31)
			fmt.Println("What category does it belong to?")
			fmt.Scanln(&input32)
			fmt.Println("How many units are there?")
			fmt.Scanln(&input33)
			fmt.Println("How much does it cost per unit?")
			fmt.Scanln(&input34)
			// first search if the item's category exists in the map

			j, categoryFound := findCategory(Category, input32)
			if !categoryFound && j == -1 {
				fmt.Println("You item does not belong to any existing category...")
				fmt.Println("Please create a new category for it and come back here to add your item.")
				return
			}
			for key, value := range items {
				if key == input31 {
					value.Quantity = value.Quantity + input33
					if value.Unit_Cost != input34 {
						fmt.Println("The existing price is " + fmt.Sprintf("%g", value.Unit_Cost) + ", and different from your new price of " + fmt.Sprintf("%g", input34) + ".")
						fmt.Println("Do you wish to change the price to the new one, y(es)/n(o)?")
						fmt.Scanln(&input35)
						if input35 == "y" {
							fmt.Println("Noted, the item, " + key + " price is now changed to ...:" + fmt.Sprintf("%g", input34))
							value.Unit_Cost = input34
						} else if input35 == "n" {
							fmt.Println("Noted, price is left unchanged.")
						} else {
							fmt.Println("I don't know what you mean. Exiting...")
							break
						}
					} else {
						value.Unit_Cost = value.Unit_Cost + input34
					}
				} else { // item is not in existing mapping, do add it in
					items[input31] = item{j, input33, input34}
				}
			}
			pauseToRead(&input1)
		case 4:
			fmt.Println("Modify Items.")
			var input41, input42, input43, input44, input45 string
			// var input45 float64
			fmt.Println("Which item do you wish to modify?")
			fmt.Scanln(&input41)
			for key, value := range items {
				if key == input41 {
					fmt.Println("Current item name is " + key + " - Category is " + Category[value.Category] + " - Quantity is " + strconv.Itoa(value.Quantity) + " - Unit Cost " + fmt.Sprintf("%g", value.Unit_Cost))
					fmt.Println("Enter new name. Enter for no change.")
					fmt.Scanln(&input42)
					if input42 != "" && input42 != key {
						items[input42] = item{value.Category, value.Quantity, value.Unit_Cost}
						delete(items, key)
					}
					fmt.Println("Enter new Category. Enter for no change.")
					fmt.Scanln(&input43)
					if input43 != "" {
						j, categoryFound := findCategory(Category, input43)
						if !categoryFound && j == -1 {
							fmt.Println("You item does not belong to any existing category...")
							fmt.Println("Please create a new category for it and come back here to add your item.")
						} else {
							items[key] = item{j, value.Quantity, value.Unit_Cost}
						}
					}

					fmt.Println("Enter new Quantity. Enter for no change.")
					fmt.Scanln(&input44)
					qty, err := strconv.Atoi(input44)
					for err != nil {
						fmt.Println("Enter new Quantity. Enter for no change.")
						fmt.Scanln(&input44)
						qty, err = strconv.Atoi(input44)
					}
					items[key] = item{value.Category, qty, value.Unit_Cost}

					fmt.Println("Enter new Unit cost. Enter for no change.")
					fmt.Scanln(&input45)
					cost, err := strconv.ParseFloat(input45, 64)
					for err != nil {
						fmt.Println("Enter new Unit cost. Enter for no change.")
						fmt.Scanln(&input45)
						cost, err = strconv.ParseFloat(input45, 64)
					}
					items[key] = item{value.Category, value.Quantity, cost}
				}
			}
			pauseToRead(&input1)
		case 5:
			fmt.Println("Delete Items")
			var input51 string
			fmt.Println("Enter the item to delete")
			fmt.Scanln(&input51)
			_, found := items[input51]
			if found {
				delete(items, input51)
				fmt.Println("Deleted" + input51)
			} else {
				fmt.Println("Item not found. Nothing to delete")
			}
			pauseToRead(&input1)
		case 6:
			fmt.Println("Print Current Data.")
			if len(items) != 0 {
				for key, value := range items {
					fmt.Println(key + " - {" + strconv.Itoa(value.Category) + " " + strconv.Itoa(value.Quantity) + " " + fmt.Sprintf("%g", value.Unit_Cost) + "}")
				}
			} else {
				fmt.Println("No data found")
			}
			pauseToRead(&input1)
		case 7:
			var newCategory string
			fmt.Println("Add New Category Name")
			fmt.Println("What is the New Category Name to add?")
			fmt.Scanln(&newCategory)
			i, found := findCategory(Category, newCategory)
			if !found {
				//Append new category to Category
				Category = append(Category, newCategory)
				fmt.Println("New category: " + newCategory + " added at index " + strconv.Itoa(len(Category)))
			} else if found {
				fmt.Println("Category: " + newCategory + " already exists at index " + strconv.Itoa(i))
			} else {
				fmt.Println("No input found")
			}
			pauseToRead(&input1)
		case 0:
			fmt.Println("Quiting Program. Bye bye!")
			return
		default:
			displayMainMenu()
			fmt.Scanln(&input1)
		}
		displayMainMenu()
		fmt.Scanln(&input1)
	}
}
