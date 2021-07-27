package services

import "github.com/luispfcanales/online-shop/internal/models"

type Repository interface {
	Save(user *models.User) *models.User
}
