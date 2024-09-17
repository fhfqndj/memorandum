package entities

import (
    "gorm.io/gorm"
    "time"
)

type User struct {
    gorm.Model
    ID        uint      `gorm:"primaryKey"`
    FirstName string    `gorm:"size:100;not null"`
    LastName  string    `gorm:"size:100;not null"`
    Email     string    `gorm:"size:100;unique;not null"`
    Password  string    `gorm:"size:255;not null"` // Строка для хэшированного пароля
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserRepository interface {
    Create(user *User) error
    GetByEmail(email string) (*User, error)
    GetByID(id uint) (*User, error)
    Update(user *User) error
    Delete(id uint) error
}

type TokenService interface {
    GenerateToken(user *User) (string, error)
}