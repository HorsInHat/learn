package types

import (
	"fmt"
	"strconv"
	"reflect"
)

type Shop struct {
	ProductList []Product
}

type Product struct {
	Code int
	Name string
	Weight float32
	Price float64
	WeightType string
}

func ShopArray() {
	productList := []string{
		"Apple",
		"Pineapple",
		"Banana",
		"Water",
	}

	// Print array element
	//fmt.Printf("Products: %v \n", productList[2])

	// count
	//for i:=0;i<5;i++{
	//	fmt.Println(i)
	//}

	//lenght := len(productList)
	//fmt.Println(lenght)

	// Normal cikl
	//lenght := len(productList)
	//for i:=0; i<lenght; i++{
	//	fmt.Println(productList[i])
	//}


	// search elem in array
	for index, value := range productList{
		if value == "Banana"{
			fmt.Println(index)
		}
	}
}

func ShopMap() {
	productList := make(map[string]string)

	productList["Banana"] = "1"

	var count string
	sum := 15
	count = strconv.Itoa(sum)
	productList["Apple"] = count
	var weight string
	weightOld := 3.14
	// Float to string
	weight = strconv.FormatFloat(weightOld, 'f', 2, 64)
	productList["Water"] = weight
	productList["Cola"] = "black"
	productList["Milk"] = "5"

	for index , value := range productList{
		fmt.Println("--------------------")
		fmt.Println(index,value)
		fmt.Println(reflect.TypeOf(value))
	}
}

func ShopStruct() {
	apple := Product{
		Code:0001,
		Name: "Apple green",
		Price: 5.95,
		Weight: 1.2,
		WeightType: "kg",
	}
	milk := Product{
		Code:0002,
		Name: "Milk green",
		Price: 5.95,
		Weight: 1.2,
		WeightType: "kg",
	}
	water := Product{
		Code:0003,
		Name: "Water green",
		Price: 5.95,
		Weight: 1.2,
		WeightType: "kg",
	}

	productList := []Product{}
	productList = append(productList, apple)
	productList = append(productList, milk)
	productList = append(productList, water)

	// Print product list with price
	//fmt.Printf("Product: %v", productList[1].Name )
	//for _ , value := range productList {
	//	fmt.Println("--------------------")
	//	fmt.Printf("%v cost %s $", value.Name, value.Price)
	//}

	for _, value := range productList{
		if value.Name == "Milk green"{
			fmt.Println(value.Price)
		}
	}
}