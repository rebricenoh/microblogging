package domain

// User representa a un usuario en el sistema.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
