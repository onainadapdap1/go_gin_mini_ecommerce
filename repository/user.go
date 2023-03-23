package repository

import (
	"fmt"
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
	GetByEmail(string) (models.User, error)
	GetAllUser() ([]models.User, error)
	UpdateUser(uint, models.User) (models.User, error)
	DeleteUser(uint) (string, error)
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

func (r *userRepository) GetByEmail(email string) (user models.User, err error) {
	return user, r.db.First(&user, "email = ?",email).Error
}

func (r *userRepository) GetUser(id int) (user models.User, err error) {
	return user, r.db.First(&user, id).Error
}

func (r *userRepository) GetAllUser() (users []models.User, err error)  {
	return users, r.db.Find(&users).Error
}

func (r *userRepository) UpdateUser(id uint, user models.User) (models.User, error) {
	// get data from db by id
	userDB := models.User{}
	if err := r.db.First(&userDB, "id = ?", id).Error; err != nil {
		return user, err
	}

	user.ID = userDB.ID
	user.Email = userDB.Email
	user.Password = ""
	return user, r.db.Model(&user).Where("id = ?", id).Updates(&user).Error
}

func (r *userRepository) DeleteUser(id uint) (string, error) {
	user := models.User{}
	if err := r.db.Find(&user, "id = ?", id).Error; err != nil {
		return "error fetching user data", err
	}
	r.db.Delete(&user)
	return fmt.Sprintf("Data id %v Successfull deleted", id), nil

}