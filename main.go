package main

import (
	"fmt"
	"strconv"
)

type item struct {
	Category  int
	Quantity  int
	Unit_Cost float64
}

var Category []string

var items map[string]item

func init() {
	Category = []string{"Household", "Food", "Drinks"}

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
	var input string
	var input1 int
	choices := makeMenuRange(0, len(mainMenu))

	displayMainMenu()
	fmt.Scanln(&input)
	input1, _ = isInt(input)
	for !validSelection(choices, input1) {
		displayMainMenu()
		fmt.Scanln(&input)
		input1, _ = isInt(input)
	}
	switch input1 {
	case 1:
		printShoppingList()
		pauseToRead()
	case 2:
		printShoppingRpts()
	case 3:
		var input31, input32 string
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
		addItems(input31, input32, input33, input34)
		pauseToRead()
	case 4:
		fmt.Println("Modify Items.")
		var input41 string
		var modItem, modCat, modQty, modUC bool
		fmt.Println("Which item do you wish to modify?")
		fmt.Scanln(&input41)
		_, found := items[input41]
		if found {
			modItem, modCat, modQty, modUC = modifyItem(input41)
			fmt.Println("Item Name is changed: " + strconv.FormatBool(modItem))
			fmt.Println("Category is changed: " + strconv.FormatBool(modCat))
			fmt.Println("Quantity is changed: " + strconv.FormatBool(modQty))
			fmt.Println("Unit-cost is changed: " + strconv.FormatBool(modUC))
		} else {
			fmt.Println("Item is not found. Nothing to modify!")
		}
		pauseToRead()
	case 5:
		fmt.Println("Delete Items")
		var input51 string
		fmt.Println("Enter the item to delete")
		fmt.Scanln(&input51)
		_, found := items[input51]
		if found {
			delete(items, input51)
			fmt.Println("Deleted " + input51)
		} else {
			fmt.Println("Item not found. Nothing to delete")
		}
		pauseToRead()
	case 6:
		printDataInMem()
		pauseToRead()
	case 7:
		var newCategory string
		fmt.Println("Add New Category Name")
		fmt.Println("What is the New Category Name to add?")
		fmt.Scanln(&newCategory)
		_ = addNewCategory(newCategory)
		pauseToRead()
	case 0:
		fmt.Println("Quiting Program. Bye bye!")
		return
	}
	main()
}
