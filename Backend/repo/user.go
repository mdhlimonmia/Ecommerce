package repo

import (
	"ecommerce/domain"
	"ecommerce/user"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(u domain.User) (*domain.User, error) {
	// if u.ID != 0 {
	// 	return &u, nil
	// }
	// u.ID = len(r.users) + 1

	// r.users = append(r.users, &u)
	// return &u, nil

	query := `
		INSERT INTO users (
			first_name, 
			last_name, 
			email, 
			password, 
			is_shop_owner
		) 
        VALUES (
			:first_name, 
			:last_name, 
			:email, 
			:password, 
			:is_shop_owner
		)
		RETURNING id
	`
	var userId int
	rows, err := r.db.NamedQuery(query, u)
	if err != nil {
		// fmt.Println("Error creating user:", err)
		// println("Query:", q uery)
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&userId)
		if err != nil {
			return nil, err
		}
		u.ID = userId
	}
	return &u, nil
}

func (r *userRepo) UserList() ([]*domain.User, error) {
	// var userList []*User
	// for _, u := range r.users {
	// 	var us User
	// 	us.FirstName = u.FirstName
	// 	us.LastName = u.LastName
	// 	userList = append(userList, &us)
	// }
	// return userList, nil
	var userList []*domain.User
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner FROM users`
	err := r.db.Select(&userList, query)
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (r *userRepo) FindUser(email string) bool {
	var user domain.User
	query := `SELECT id FROM users WHERE email = $1 LIMIT 1`

	err := r.db.Get(&user, query, email)
	if err != nil {
		return false
	}
	return true
}

func (r *userRepo) AuthUser(email string, pass string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner FROM users WHERE email = $1 AND password = $2 LIMIT 1`

	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
