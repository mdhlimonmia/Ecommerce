package repo

import "fmt"

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(id int, p Product) (*Product, error)
	Delete(id int) error
}

type productRepo struct {
	productList []*Product
}

// constructor function that crate a new instance of the product repository
func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	addInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	fmt.Println("Storing product: ", p)
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	return r.productList[id], nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Update(id int, p Product) (*Product, error) {
	p.ID = id
	for i, prd := range r.productList {
		if prd.ID == id {
			r.productList[i] = &p
			fmt.Println(p)
			return &p, nil
		}
	}
	return &p, nil
}

func (r *productRepo) Delete(id int) error {
	for i, prd := range r.productList {
		if prd.ID == id {
			r.productList = append(r.productList[:i], r.productList[i+1:]...)
			return nil
		}
	}
	return nil
}

func addInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Banana",
		Description: "Banana color is Yellow",
		Price:       20,
		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
	}
	prd2 := &Product{
		ID:          2,
		Title:       "Banana",
		Description: "Banana color is Yellow",
		Price:       20,
		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
	}
	prd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana color is Yellow",
		Price:       20,
		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
	}
	prd4 := &Product{
		ID:          4,
		Title:       "Banana",
		Description: "Banana color is Yellow",
		Price:       20,
		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
	}
	prd5 := &Product{
		ID:          5,
		Title:       "Banana",
		Description: "Banana color is Yellow",
		Price:       20,
		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
	}

	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd3)
	r.productList = append(r.productList, prd4)
	r.productList = append(r.productList, prd5)
}
