package repository

import "github.com/rezalaal/coral/internal/models"

// RestaurantRepository defines methods to interact with restaurants in the database.
type RestaurantRepository interface {
    Create(restaurant *models.Restaurant) error
    GetByID(id int) (*models.Restaurant, error)
    List() ([]*models.Restaurant, error)
    Delete(id int) error
}
