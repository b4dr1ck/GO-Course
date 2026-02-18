package main

import "fmt"

func main() {
	var productNames [4]string
	productNames = [4]string{"Book"}
	productNames[2] = "Carpet"
	prices := [4]float64{10.99, 39.99, 50.0, 12.99}

	// can also use with mixed data-types
	mixedData := [5]any{1, 5.0, "hello"}

	// dynamic array
	amounts := []int{1, 2, 5, 6, 3, 1}
	amounts = append(amounts, 12)

	// remove first element (with slices)
	amounts = amounts[1:]

	// slices (start-index is included, end-index is excluded)
	featuredPrices := prices[1:3]

	fmt.Printf("prices: %v\n", prices)
	fmt.Printf("featuredPrices (prices[1:3]): %v\n", featuredPrices)
	fmt.Printf("productNames: %v\n", productNames)
	fmt.Printf("The price of the %s is %.2f€ \n", productNames[2], prices[2])
	fmt.Printf("mixedData: %v\n", mixedData)
	fmt.Printf("amounts: %v\n", amounts)

	// Exercise - Lists
	// Time to practice what you learned!

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.

	fmt.Println()
	fmt.Println("*EXERCISE Lists")
	hobbies := [3]string{"drawing", "hiking", "music"}
	hobbies2and3 := hobbies[1:3]
	mainHobbies := hobbies[0:2]
	goals := []string{"learn basics", "learn in-depth", "databases", "api"}
	goals[1] = "learn more"
	goals = append(goals, "web-dev")

	type Products struct {
		title string
		id    int
		price float64
	}

	product1 := Products{
		title: "Book",
		id:    0,
		price: 12.99,
	}

	product2 := Products{
		title: "Apple",
		id:    1,
		price: 0.99,
	}

	product3 := Products{
		title: "TV",
		id:    2,
		price: 599.99,
	}

	productsList := []Products{product1, product2}
	productsList = append(productsList, product3)

	fmt.Printf("My hobbies are: %v\n", hobbies)
	fmt.Printf("First Hobby: %v\n", hobbies[0])
	fmt.Printf("2nd and 3rd hobby: %v\n", hobbies2and3)
	fmt.Printf("mainHobbies: %v\n", mainHobbies)
	fmt.Printf("%v\n", mainHobbies[1:3]) // re-slicing: must specify the explicit index
	fmt.Printf("goals: %v\n", goals)
	fmt.Printf("Products: %v\n", productsList)

	for i := 0; i < len(productsList); i++ {
		fmt.Printf("title: %v\n", productsList[i].title)
		fmt.Printf("id: %v\n", productsList[i].id)
		fmt.Printf("price: %v\n\n", productsList[i].price)
	}
}
