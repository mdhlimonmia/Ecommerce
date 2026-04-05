package database

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

var ProductList []Product

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

	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)
	ProductList = append(ProductList, prd4)
	ProductList = append(ProductList, prd5)
}
