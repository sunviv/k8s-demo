package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64  `gorm:"primaryKey, autoincrement"`
	Email     string `gorm:"unique"`
	Password  string `gorm:"size:255"`
	CreatedAt int64
	UpdatedAt int64
}

func (u *User) TableName() string {
	return "sys_user"
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (dao *UserDao) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.CreatedAt = now
	u.UpdatedAt = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	var sqlErr *mysql.MySQLError
	if errors.As(err, &sqlErr) {
		if sqlErr.Number == 1062 {
			return gorm.ErrDuplicatedKey
		}
	}
	return err
}

func (dao *UserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).First(&u, "email = ?", email).Error
	return u, err
}
