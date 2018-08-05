package store;

import (
	"fmt"
)

type Repository struct {};
const SERVER = "http://something";
const DBNAME ="coolname";
const COLLECTION="anothercoolname";

type Product struct {
	name string
	cost uint32
}
func(r Repository)GetProducts() []Product {
	results:= []Product{
		{
			name: "iron",
			cost: 680;
		},
		{
			name: "fan",
			cost: 333;
		}
	}
	return results;
}

func(r Repository) AddProduct(product Product) bool {
	fmt.Print("The product was added");
	return true;
}