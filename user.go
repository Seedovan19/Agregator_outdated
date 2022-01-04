package agregator

type User struct {
	ID       int    `json: "-"; gorm: "primaryKey"`
	Name     string `json: "name" binding: "required"`
	Username string `json: "username" binding: "required"`
	Password string `json: "password" binding: "required"`
}
