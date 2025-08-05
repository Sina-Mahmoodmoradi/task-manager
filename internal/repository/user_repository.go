package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Sina-Mahmoodmoradi/task-manager/internal/domain"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/repository/gorm/models"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

// Create inserts a new user into the database
func (r *userRepository) Create(user *domain.User) error {
	userModel := models.FromDomainUser(*user)

	if err := r.db.Create(&userModel).Error; err != nil {
		return err
	}

	// map generated ID back to domain
	user.ID = userModel.ID

	return nil
}

// GetByUsername retrieves a user by username
func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	var userModel models.User

	err := r.db.Preload("Tasks").Where("username = ?", username).First(&userModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	// Convert model.User → domain.User
	domainUser := models.ToDomainUser(userModel)

	return &domainUser, nil
}


func (r *userRepository) GetByID(id uint)(*domain.User,error){
	var userModel models.User

	if err:=r.db.First(&userModel,id).Error; err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,nil
		}
		return nil,err
	}

	domainUser := models.ToDomainUser(userModel)

	return &domainUser,nil
}
