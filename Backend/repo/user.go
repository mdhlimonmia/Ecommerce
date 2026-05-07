package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type userRepo struct {
	users []*User
}

type UserRepo interface {
	Create(u User) (*User, error)
	UserList() ([]*User, error)
	FindUser(email string) bool
	AuthUser(email string, pass string) (*User, error)
}

func NewUserRepo() UserRepo {
	repo := &userRepo{}
	return repo
}

func (r *userRepo) Create(u User) (*User, error) {
	if u.ID != 0 {
		return &u, nil
	}
	u.ID = len(r.users) + 1

	r.users = append(r.users, &u)
	return &u, nil
}

func (r *userRepo) UserList() ([]*User, error) {
	var userList []*User
	for _, u := range r.users {
		var us User
		us.FirstName = u.FirstName
		us.LastName = u.LastName
		userList = append(userList, &us)
	}
	return userList, nil
}

func (r *userRepo) FindUser(email string) bool {
	for _, u := range r.users {
		if u.Email == email {
			return true
		}
	}
	return false
}

func (r *userRepo) AuthUser(email string, pass string) (*User, error) {
	for _, u := range r.users {
		if u.Email == email && u.Password == pass {
			return u, nil
		}
	}
	return nil, nil
}
