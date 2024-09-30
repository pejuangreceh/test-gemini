package main

import "gorm.io/gorm"

type Server struct {
	DB *gorm.DB
}
