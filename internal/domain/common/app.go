package common

import "gorm.io/gorm"

type App struct {
	DB  map[string]*gorm.DB
	Env Env
}

type Env struct {
	HTTP_ADDR  string
	JWT_SECRET string
}
