package requests

import "github.com/eduhub/util"

type LoginRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type ClaimsResponse struct {
	User     *User      `json:"user"`
	Roles    []*Roles   `json:"roles"`
	Colleges []*College `json:"colleges"`
	Success  bool       `json:"success"`
	Message  string     `json:"message"`
}

type User struct {
	FirstName  string        `json:"first_name"`
	MiddleName string        `json:"middle_name"`
	LastName   string        `json:"last_name"`
	Email      string        `json:"email"`
	Address    string        `json:"address"`
	Phone      string        `json:"phone"`
	UserType   util.UserType `json:"user_type"`
}

type Roles struct {
	RoleName    string        `json:"role_name"`
	Permissions []*Permission `json:"permissions"`
}

type Permission struct {
	Name string `json:"permission_name"`
}

type College struct {
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Course  []Course `json:"course"`
}

type ClaimsRequest struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	Message string `json:"message"`
}

type BasicResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SignUpRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserName  string `json:"username"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Password  string `json:"password"`
}

type UpdateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

type Course struct {
	Name          string `json:"name"`
	Code          string `json:"code"`
	Description   string `json:"description"`
	Prerequisites string `json:"prerequisites"`
}
