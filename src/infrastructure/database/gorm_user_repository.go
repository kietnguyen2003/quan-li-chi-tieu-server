package database

import (
	"quan-li-chi-tieu/src/application/auth"
	domainAuth "quan-li-chi-tieu/src/domain/auth"

	"gorm.io/gorm"
)

type GormUser struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	FullName  string `gorm:"not null"`
	CreatedAt int64
	UpdatedAt int64
}

func (GormUser) TableName() string {
	return "users"
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

var _ auth.UserRepository = (*GormUserRepository)(nil)

func (r *GormUserRepository) GetUserByEmail(email string) (*domainAuth.User, error) {
	var gormUser GormUser

	if err := r.db.Where("email = ?", email).First(&gormUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &domainAuth.User{
		ID:        gormUser.ID,
		Username:  gormUser.Username,
		Email:     gormUser.Email,
		Password:  gormUser.Password,
		FullName:  gormUser.FullName,
		CreatedAt: timeFromUnix(gormUser.CreatedAt),
		UpdatedAt: timeFromUnix(gormUser.UpdatedAt),
	}, nil
}

func (r *GormUserRepository) toDomainUser(gormUser *GormUser) *domainAuth.User {
	return &domainAuth.User{
		ID:        gormUser.ID,
		Username:  gormUser.Username,
		Email:     gormUser.Email,
		Password:  gormUser.Password,
		CreatedAt: timeFromUnix(gormUser.CreatedAt),
		UpdatedAt: timeFromUnix(gormUser.UpdatedAt),
	}
}

func (r *GormUserRepository) CreateUser(user *domainAuth.User) (uint, error) {
	gormUser := GormUser{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		FullName: user.FullName,
	}
	err := r.db.Create(&gormUser).Error
	return gormUser.ID, err
}
