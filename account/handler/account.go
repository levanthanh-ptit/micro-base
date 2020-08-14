package handler

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"

	account "account/proto"
)

type Account struct {
	gorm.Model
	Email    string `gorm:"NOT NULL"`
	FullName string `gorm:"NOT NULL"`
	Role     string `gorm:"NOT NULL"`
}

func (e *Account) Create(ctx context.Context, req *account.CreateReq, rsp *account.CreateRes) error {
	rsp.Email = req.Email
	rsp.FullName = req.FullName
	rsp.CreatedAt = time.Now().String()
	rsp.UpdatedAt = rsp.CreatedAt
	rsp.DeletedAt = ""
	return nil
}
