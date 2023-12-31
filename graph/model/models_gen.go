// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Company struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Job struct {
	ID        string `json:"id"`
	CompanyID string `json:"CompanyId"`
	Role      string `json:"role"`
	Salary    string `json:"salary"`
}

type NewCompany struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type NewJob struct {
	CompanyID string `json:"CompanyId"`
	Role      string `json:"role"`
	Salary    string `json:"salary"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
