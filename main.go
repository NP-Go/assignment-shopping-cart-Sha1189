package main

import (
	"fmt"
	"strconv"
)

func main() {

	var userInput int

	displayMenu()
	fmt.Printf("Select your choice: ")
	fmt.Scanln(&userInput)

	switch userInput {
	case 1:

		for name, value := range itemInfo {

			fmt.Printf("Category: %s - Item: %s Quantity: %d Unit Cost: %g\n", category[value.category], name, value.quantity, value.unitCost)
		}
		delay()
	case 2:

		var reportInput int
		fmt.Println("Generate Report")
		for i := range reportMenu {
			fmt.Println(reportMenu[i])
		}
		fmt.Printf("Choose your report: ")
		fmt.Scanln(&reportInput)

		if reportInput == 1 {

			fmt.Println("Total Cost by Category.")
			totalCostbyCat := make([]float64, len(category))
			for _, value := range itemInfo {
				for name := range category {
					if value.category == name {
						totalCostbyCat[name] += value.unitCost
						break
					}
				}
			}
			for s := range category {
				fmt.Printf(category[s]+" cost : %g\n", totalCostbyCat[s])
			}

		}
		if reportInput == 2 {
			fmt.Println("List by Category.")
			// sort list
			for slice, name := range category {
				for key, value := range itemInfo {
					if slice == int(value.category) {
						fmt.Printf("Category: %v - Item: %v  Quantity: %v  Unit Cost: %g\n", name, key, value.quantity, value.unitCost)
					}
				}
			}

		}
		if reportInput == 3 {
			displayMenu()
		}
		delay()
	case 3:

		var inputItem, inputCat string
		var inputQty int
		var inputCost float64
		fmt.Println("What is the name of your item?")
		fmt.Scan(&inputItem)
		fmt.Println("What Category does it belong to?")
		fmt.Scan(&inputCat)
		cName, cExists := checkCategory(category, inputCat)
		if !cExists && cName == 0 {

			fmt.Printf("%s is not in 'List of Category', please select '7' from Main Menu to Add New Category Name", inputCat)
			return
		}
		fmt.Println("How many units are there?")
		fmt.Scan(&inputQty)
		fmt.Println("How much does it cost per unit?")
		fmt.Scan(&inputCost)

		for name, value := range itemInfo {
			if name == inputItem {

				fmt.Printf("%s is in inventory under %s, please select '4' from Main Menu to Modify Items", inputItem, category[value.category])

			} else {

				itemInfo[inputItem] = item{cName, inputQty, inputCost}
				break
			}
		}
		delay()
	case 4:

		var inputItemMod, newItemMod, newCatMod string
		var newQtyMod int
		var newCostMod float64

		fmt.Println("Modify Items.")
		fmt.Println("Which item would you wish to modify?")
		fmt.Scan(&inputItemMod)

		fmt.Printf("Current item name is %s - Category is %s - Quantity is %d - Unit Cost is %g\n", inputItemMod, category[itemInfo[inputItemMod].category], itemInfo[inputItemMod].quantity, itemInfo[inputItemMod].unitCost)

		fmt.Println("Enter new name. Enter for no change.")
		fmt.Scan(&newItemMod)
		fmt.Println("Enter new Category. Enter for no change.")
		fmt.Scan(&newCatMod)
		fmt.Println("Enter new Quantity. Enter for no change.")
		fmt.Scan(&newQtyMod)

		if newItemMod != "" && newItemMod != inputItemMod {
			itemInfo[newItemMod] = item{itemInfo[inputItemMod].category, itemInfo[inputItemMod].quantity, itemInfo[inputItemMod].unitCost}
			delete(itemInfo, inputItemMod)
		}
		if newCatMod != "" {
			cName, cExists := checkCategory(category, newCatMod)
			if !cExists && cName == 0 {
				fmt.Printf("%s is not in 'List of Category', please select '7' from Main Menu to Add New Category Name", newCatMod)

			} else {
				fmt.Println("No changes to category made")
			}
			itemInfo[inputItemMod] = item{itemInfo[inputItemMod].category, itemInfo[inputItemMod].quantity, itemInfo[inputItemMod].unitCost}
		}

		if newQtyMod != 0 {
			itemInfo[newItemMod] = item{itemInfo[inputItemMod].category, itemInfo[inputItemMod].quantity, itemInfo[inputItemMod].unitCost}

		} else {
			fmt.Println("No changes to quantity made")
		}
		fmt.Println("Enter new Cost. Enter for no change.")
		fmt.Scan(&newCostMod)

		if newCostMod != 0.0 {
			itemInfo[newItemMod] = item{itemInfo[inputItemMod].category, itemInfo[inputItemMod].quantity, itemInfo[inputItemMod].unitCost}
		} else {
			fmt.Println("No changes to cost made.")
		}
		delay()
	case 5:

		var delItem string
		fmt.Println("Delete Item.")
		fmt.Println("Enter item name to delete:")
		fmt.Scan(&delItem)
		_, ok := itemInfo[delItem]
		if ok {
			delete(itemInfo, delItem)
			fmt.Printf("Deleted %s", delItem)
		} else {
			fmt.Printf("Item: %s not found in inventory. Nothing to delete!", delItem)
		}
		delay()
	case 6:

		fmt.Println("Print Current Data.")
		if len(itemInfo) != 0 {
			for name, value := range itemInfo {
				fmt.Printf("%s - {%s %s %g}  \n", name, strconv.Itoa(value.category), strconv.Itoa(value.quantity), value.unitCost)
			}
		} else {
			fmt.Println("No data found!")
		}
		delay()
	case 7:

		var newCat string
		fmt.Println("Add New Category Name")
		fmt.Println("What is the New Category Name to add?")
		fmt.Scan(&newCat)

		c, ok := checkCategory(category, newCat)
		if !ok {
			category = append(category, newCat)
			fmt.Printf("New category: %s added at index %s\n", newCat, strconv.Itoa(len(category)-1))
		} else if ok {
			fmt.Printf("Category: %s already exists at index %s\n", newCat, strconv.Itoa(c))

		} else {
			fmt.Println("No input found")

		}
		delay()
	}
	main()
}
