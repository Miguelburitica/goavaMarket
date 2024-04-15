package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserToUpdate struct {
	User
}

type GetUsersRequest struct {
	Page    string `json:"page"`
	Offset  string `json:"offset"`
	Pattern string `json:"pattern"`
}
