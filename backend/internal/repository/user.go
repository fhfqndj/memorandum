package repository

import (
    "memorandum-backend/internal/entities"
    "gorm.io/gorm"
)


type GormUserRepository struct {
    db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) entities.UserRepository {
    return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(user *entities.User) error {
    return r.db.Create(user).Error
}

func (r *GormUserRepository) GetByEmail(email string) (*entities.User, error) {
    var user entities.User
    err := r.db.Where("email = ?", email).First(&user).Error
    return &user, err
}

func (r *GormUserRepository) GetByID(id uint) (*entities.User, error) {
    var user entities.User
    err := r.db.First(&user, id).Error
    return &user, err
}

func (r *GormUserRepository) Update(user *entities.User) error {
    return r.db.Save(user).Error
}

func (r *GormUserRepository) Delete(id uint) error {
    return r.db.Delete(&entities.User{}, id).Error
}