package db

type Config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Name     string `json:"name"`
}
