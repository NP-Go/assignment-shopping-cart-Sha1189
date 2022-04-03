package main

import (
	"fmt"
	"time"
)

func displayMenu() {

	fmt.Println("Shopping List Application")
	fmt.Println("=========================")
	fmt.Println("1. View entire shopping list.")
	fmt.Println("2. Generate Shopping List Report")
	fmt.Println("3. Add Items")
	fmt.Println("4. Modify Items")
	fmt.Println("5. Delete Item")
	fmt.Println("6. Print Current Data")
	fmt.Println("7. Add New Category Name")
}

func checkCategory(categoryName []string, cInput string) (int, bool) {
	for n := range categoryName {
		if categoryName[n] == cInput {
			return n, true
		}
	}
	return 0, false
}

func delay() {

	time.Sleep(3 * time.Second)
}

func init() {

	itemInfo["Forks"] = item{0, 4, 3.0}
	itemInfo["Plates"] = item{0, 4, 3.0}
	itemInfo["Cups"] = item{0, 5, 3.0}
	itemInfo["Bread"] = item{1, 2, 2.0}
	itemInfo["Cake"] = item{1, 3, 1.0}
	itemInfo["Coke"] = item{2, 5, 2.0}
	itemInfo["Sprite"] = item{2, 5, 2.0}
}

var reportMenu = []string{

	"1. Total Cost of each Category.",
	"2. List of item by category.",
	"3. Main Menu",
}

type item struct {
	category int
	quantity int
	unitCost float64
}

var itemInfo = make(map[string]item)

var category = []string{"Household", "Food", "Drinks"}
