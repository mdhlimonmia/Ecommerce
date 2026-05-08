package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imgUrl" db:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(id int, p Product) (*Product, error)
	Delete(id int) error
}

type productRepo struct {
	db *sqlx.DB
}

// constructor function that crate a new instance of the product repository
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p Product) (*Product, error) {
	// p.ID = len(r.productList) + 1
	// fmt.Println("Storing product: ", p)
	// r.productList = append(r.productList, &p)
	// return &p, nil
	query := `
		insert into products(
			title,
			description,
			price,
			img_url
		)
		values(
			$1,
			$2,
			$3,
			$4
		)
		returning id
	`
	product := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := product.Scan(&p.ID)
	if err != nil {
		fmt.Println("Error creating product:", err)
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	var p Product
	query := `
		select
		  id,
		  title,
		  description,
		  price,
		  img_url
		from products
		where id = $1
	`
	err := r.db.Get(&p, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	p.ID = id
	return &p, nil

}

func (r *productRepo) List() ([]*Product, error) {
	var products []*Product

	query := `
		select
		  id,
		  title,
		  description,
		  price,
		  img_url
		from products
	`
	err := r.db.Select(&products, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return products, nil
}

func (r *productRepo) Update(id int, p Product) (*Product, error) {
	// p.ID = id
	// for i, prd := range r.productList {
	// 	if prd.ID == id {
	// 		r.productList[i] = &p
	// 		fmt.Println(p)
	// 		return &p, nil
	// 	}
	// }
	// return &p, nil
	query := `
		update products
		set 
		title = $1,
		description = $2,
		price = $3,
		img_url = $4
		where id = $5
	`
	_, err := r.db.Exec(query, p.Title, p.Description, p.Price, p.ImgUrl, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	p.ID = id
	return &p, nil
}

func (r *productRepo) Delete(id int) error {
	// for i, prd := range r.productList {
	// 	if prd.ID == id {
	// 		r.productList = append(r.productList[:i], r.productList[i+1:]...)
	// 		return nil
	// 	}
	// }
	// return nil
	query := `
		delete from products
		where id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	return nil
}

// func addInitialProducts(r *productRepo) {
// 	prd1 := &Product{
// 		ID:          1,
// 		Title:       "Banana",
// 		Description: "Banana color is Yellow",
// 		Price:       20,
// 		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
// 	}
// 	prd2 := &Product{
// 		ID:          2,
// 		Title:       "Banana",
// 		Description: "Banana color is Yellow",
// 		Price:       20,
// 		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
// 	}
// 	prd3 := &Product{
// 		ID:          3,
// 		Title:       "Banana",
// 		Description: "Banana color is Yellow",
// 		Price:       20,
// 		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
// 	}
// 	prd4 := &Product{
// 		ID:          4,
// 		Title:       "Banana",
// 		Description: "Banana color is Yellow",
// 		Price:       20,
// 		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
// 	}
// 	prd5 := &Product{
// 		ID:          5,
// 		Title:       "Banana",
// 		Description: "Banana color is Yellow",
// 		Price:       20,
// 		ImgUrl:      "https://www.veggiesbasket.co.za/wp-content/uploads/2021/06/42E9as7NaTaAi4A6JcuFwG.jpg",
// 	}

// 	r.productList = append(r.productList, prd1)
// 	r.productList = append(r.productList, prd2)
// 	r.productList = append(r.productList, prd3)
// 	r.productList = append(r.productList, prd4)
// 	r.productList = append(r.productList, prd5)
// }
