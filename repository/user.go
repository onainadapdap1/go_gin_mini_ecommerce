package repository

import (
	"go_gin_mini_ecommerce/models"

	"github.com/jinzhu/gorm"
)

// UserRepository -> User CRUD
// digunakan sebagai tipe data pada variabel,
// dimana variable ini menampung objek dari struct yg kita buat
// UserRepository = userRepository(db: Db())
type UserRepository interface {
	GetUser(int) (models.User, error)
	AddUser(models.User) (models.User, error)
	// GetByEmail(string) (models.User, error)
	// GetAllUser() ([]models.User, error)
	// UpdateUser(models.User) (models.User, error)
	// DeleteUser(models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// buat objek UserRepository berisi objek userRepository
func NewUserRepository() UserRepository {
	return &userRepository{db: DB()}
}

func (r *userRepository) AddUser(user models.User) (models.User, error) {
	return user, r.db.Create(&user).Error
}

func (r *userRepository) GetUser(id int) (user models.User, err error) {
	return user, r.db.First(&user, id).Error
}