// internal/repository/interfaces/user_repository.go
package interfaces

import "github.com/rezalaal/coral/internal/user/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id int64) (*models.User, error)
	GetByMobile(mobile string) (*models.User, error)
	List() ([]*models.User, error)
	Delete(id int64) error
}
