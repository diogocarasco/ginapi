package repository

import "github.com/diogocarasco/ginapi/internal/model"

type MoviesRepository interface {
	GetByID(id int) (*model.Movie, error)
	GetAll() ([]*model.Movie, error)
	Create(movie *model.Movie) error
	Update(movie *model.Movie) error
	Delete(id int) error
}
