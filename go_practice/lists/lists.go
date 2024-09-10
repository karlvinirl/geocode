package main

import "fmt"

func main() {

	//create an array containing three hobbies and output to screen

	hobbies := [3]string{"reading", "music", "running"}
	fmt.Println(hobbies)

	elementone := hobbies[0]
	sliceone := hobbies[1:]
	slicetwo := hobbies[0:2]
	slicethree := slicetwo[1:3]
	dynamicArray := []string{"fast", "remember"}
	dynamicArray = append(dynamicArray, "easy")

	type product struct {
		title string
		id    float64
		price float64
	}

	product1 := product{
		title: "prod1",
		id:    1,
		price: 10.00,
	}

	product2 := product{
		title: "prod2",
		id:    2,
		price: 20.00,
	}

	product3 := product{
		title: "prod3",
		id:    3,
		price: 30.00,
	}

	listOfProducts := []product{product1, product2}
	// alterantive
	listOfProducts2 := []product{
		product{
			"new prod", 4, 40.00,
		},
		{"new prod2", 5, 50.00},
	}
	fmt.Println(listOfProducts)
	fmt.Println(listOfProducts2)
	listOfProducts = append(listOfProducts, product3)
	fmt.Println(listOfProducts)

	fmt.Printf("fist element is %v\nNew slice with 2nd and 3rd element is: %v\n", elementone, sliceone)
	fmt.Printf("fist and second element: %v\n", slicetwo)
	fmt.Printf("second and last  element: %v\n", slicethree)
	fmt.Printf("goals: %v\n", dynamicArray)

}
