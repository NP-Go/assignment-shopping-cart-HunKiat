package main

// back.go is intended for processing functions

import (
	"fmt"
	"strconv"
	"strings"
)

// returns the string, and true if exists
// returns -1 and false if otherwise
func findCategory(C []string, c string) (int, bool) {
	for i := range C {
		if strings.ToLower(C[i]) == strings.ToLower(c) {
			return i, true
		}
	}
	return -1, false
}

// add an item to map items
func addItems(itemName, itemCategory string, quantity int, unit_cost float64) {
	var input1 string
	_, categoryFound := findCategory(Category, itemCategory)
	if !categoryFound && itemCategory != "" {
		fmt.Printf("\nCategory, %s does not exists. These are the existing categories...:\n", itemCategory)
		for i := range Category {
			fmt.Println(" - " + strconv.Itoa(i) + ". " + Category[i])
		}
		fmt.Println("Enter 'a' to add " + itemCategory + " in, or 'enter' to ignore this.")
		fmt.Scanln(&input1)
		if input1 == "a" {
			_ = addNewCategory(itemCategory)
		}
	}

	_, found := items[itemName]
	if found {
		_ = addItemQty(itemName, quantity)
		_ = updateItemUnitCost(itemName, unit_cost)
	} else {
		if itemCategory != "" {
			// check if user added the category earlier or not
			index, categoryFound := findCategory(Category, itemCategory)
			if categoryFound {
				items[itemName] = item{index, quantity, unit_cost}
			} else {
				// itemCategory is blank and item is new
				fmt.Println("Unable to proceed... item is new, but category wasn't added earlier.")
			}
		} else {
			// itemCategory is blank and item is new
			fmt.Println("Unable to proceed, item is new, but category of it is unknown.")
		}
	}
}

func modifyItem(itemToMod string) (itemNameUpdated, categoryUpdated, qtyUpdated, ucUpdated bool) {
	var itemName, itemCategory, quantity, unit_cost string
	value, _ := items[itemToMod]
	fmt.Println("Current item name is " + itemToMod + " - Category is " + Category[value.Category] + " - Quantity is " + strconv.Itoa(value.Quantity) + " - Unit Cost " + fmt.Sprintf("%g", value.Unit_Cost))
	fmt.Println("Enter new name. Enter for no change.")
	fmt.Scanln(&itemName)
	itemNameUpdated = updateItemName(itemToMod, itemName)
	if itemNameUpdated {
		itemToMod = itemName
	}

	fmt.Println("Enter new Category. Enter for no change.")
	fmt.Scanln(&itemCategory)
	categoryUpdated = updateCategoryName(itemToMod, itemCategory)

	fmt.Println("Enter new Quantity. Enter for no change.")
	fmt.Scanln(&quantity)
	if quantity != "" {
		qty, err := strconv.Atoi(quantity)
		for err != nil {
			fmt.Println("Enter new Quantity. Enter for no change.")
			fmt.Scanln(&quantity)
			qty, err = strconv.Atoi(quantity)
		}
		qtyUpdated = updateItemQty(itemToMod, qty)
	}

	fmt.Println("Enter new Unit cost. Enter for no change.")
	fmt.Scanln(&unit_cost)
	cost, err := strconv.ParseFloat(unit_cost, 64)
	if unit_cost != "" {
		for err != nil {
			fmt.Println("Enter new Unit cost. Enter for no change.")
			fmt.Scanln(&unit_cost)
			cost, err = strconv.ParseFloat(unit_cost, 64)
		}
		ucUpdated = updateItemUnitCost(itemToMod, cost)
	}
	return itemNameUpdated, categoryUpdated, qtyUpdated, ucUpdated
}

func addNewCategory(newCategory string) int {
	i, found := findCategory(Category, newCategory)
	if !found {
		Category = append(Category, newCategory)
		i, _ = findCategory(Category, newCategory)
		fmt.Println("New category: " + newCategory + " added at index " + strconv.Itoa(i))
		return i // return the index position of the new category
	} else {
		fmt.Println("Category: " + newCategory + " already exists at index " + strconv.Itoa(i))
		return i
	}
}

// returns true if s is an integer
func isInt(s string) (int, bool) {
	s = strings.TrimSpace(s)
	intValue, err := strconv.Atoi(s)
	if err != nil {
		return -1, false
	}
	return intValue, true
}

// returns true if s is a string
func isString(s string) (string, bool) {
	s = strings.TrimSpace(s)
	_, err := strconv.Atoi(s)
	if err != nil {
		return s, true
	}
	return "", false
}

// add Item Qty adds the new qty to the item's qty
func addItemQty(itemName string, qty int) bool {
	for key, value := range items {
		if key == itemName {
			value.Quantity = value.Quantity + qty
			items[key] = item{value.Category, value.Quantity, value.Unit_Cost}
			return true
		}
	}
	return false
}

func updateItemName(oldName, newName string) bool {
	value, _ := items[oldName]
	if newName != "" {
		items[newName] = item{value.Category, value.Quantity, value.Unit_Cost}
		delete(items, oldName)
		return true
	}
	return false
}

func updateCategoryName(itemName, newCat string) bool {
	if newCat != "" {
		i, found := findCategory(Category, newCat)
		if found {
			value, ok := items[itemName]
			if ok {
				items[itemName] = item{i, value.Quantity, value.Unit_Cost}
				return true
			}
			return false
		}
		return false
	}
	return false
}

// update Item Qty replaces the item qty with new value
func updateItemQty(itemName string, qty int) bool {
	value, ok := items[itemName]
	if ok {
		items[itemName] = item{value.Category, qty, value.Unit_Cost}
		return true
	}
	return false
}

// update Item Unit-Cost replaces the item Unit-Cost with new value
func updateItemUnitCost(itemName string, uc float64) bool {
	value, ok := items[itemName]
	if ok {
		items[itemName] = item{value.Category, value.Quantity, uc}
		return true
	}
	return false
}
