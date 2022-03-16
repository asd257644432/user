package dal

import (
	"context"
	"gorm.io/gorm"
	"log"
	"user/model"
)

type UserDal struct {
	ctx context.Context
	db  *gorm.DB
}

func NewUserDal(ctx context.Context) *UserDal {
	return &UserDal{ctx: ctx}
}

func NewUserDalWithTx(ctx context.Context, tx *gorm.DB) *UserDal {
	return &UserDal{ctx: ctx, db: tx}
}

func (d *UserDal) Create(user *model.User) (int64, error) {
	if err := d.db.Model(&model.User{}).Create(user).Error; err != nil {
		log.Printf("failed to create user, err = %+v", err)
		return 0, err
	}
	return user.Id, nil
}

func (d *UserDal) GetUser(id int64) (*model.User, error) {
	var user model.User
	if err := d.db.Model(&model.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		log.Printf("failed to get user by id, err = %+v", err)
		return nil, err
	}
	return &user, nil
}
