package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var users []User

func (u User) Store() User {
	if u.ID != 0 {
		return u
	}
	u.ID = len(users) + 1

	users = append(users, u)
	return u
}

func UserList() []User {
	var userList []User
	for _, u := range users {
		var us User
		us.FirstName = u.FirstName
		us.LastName = u.LastName
		userList = append(userList, us)
	}
	return userList
}

func FindUser(email string) bool {
	for _, u := range users {
		if u.Email == email {
			return true
		}
	}
	return false
}

func AuthUser(email string, pass string) *User {
	for _, u := range users {
		if u.Email == email && u.Password == pass {
			return &u
		}
	}
	return nil
}
