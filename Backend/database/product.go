package database

import "fmt"

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var productList []Product

func StoreProduct(p Product) int {
	p.ID = len(productList) + 1
	fmt.Println("Storing product: ", p)
	productList = append(productList, p)
	return p.ID
}

func GetProductsList() []Product {
	return productList
}

func GetProductById(id int) (Product, bool) {
	for _, product := range productList {
		if product.ID == id {
			return product, true
		}
	}
	return Product{}, false
}

func Update(id int, p Product) (Product, bool) {
	p.ID = id
	for i, prd := range productList {
		if prd.ID == id {
			productList[i] = p
			fmt.Println(p)
			return p, true
		}
	}
	return Product{}, false
}

func DeleteProduct(id int) bool {
	for i, prd := range productList {
		if prd.ID == id {
			productList = append(productList[:i], productList[i+1:]...)
			return true
		}
	}
	return false
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Banana",
		Description: "Banana color is Yellow",
		Price:       20,
		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Banana",
		Description: "Banana color is Yellow",
		Price:       20,
		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana color is Yellow",
		Price:       20,
		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Banana",
		Description: "Banana color is Yellow",
		Price:       20,
		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Banana",
		Description: "Banana color is Yellow",
		Price:       20,
		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
}
