package main

// frontEnd.go is created for code orgainisation and ease of maintenance
// This handles all the display menu and report

import "fmt"

// this creates a slice of integer
// to be used with menu selection
func makeMenuRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// returns true if integer exists in slice
func validSelection(m []int, c int) bool {
	for i := range m {
		if m[i] == c {
			return true
		}
	}
	return false
}

var mainMenu = []string{
	"1. View entire shopping list.",
	"2. Generate Shopping List Report",
	"3. Add Items",
	"4. Modify Items",
	"5. Delete Item",
	"6. Print Current Data",
	"7. Add New Category Name",
	"0. Quit Program",
}

func displayMainMenu() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Shopping List Application")
	fmt.Println("=========================")
	for i := range mainMenu {
		fmt.Println(mainMenu[i])
	}
	fmt.Println("Select your choice:")
}

var reportMenu = []string{
	"1. Total Cost of each category.",
	"2. List of item by category.",
	"3. Main Menu.",
}

func displayRptSubmenu() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Generate Report")
	for j := range reportMenu {
		fmt.Println(reportMenu[j])
	}
}

func pauseToRead(input *int) {
	fmt.Println("")
	fmt.Println("Pausing... Press enter to continue.")
	*input = -1
	fmt.Scanln()
}
