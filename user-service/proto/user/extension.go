package user

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func (m *User) BeforeCreate(scope *gorm.DB) error {
	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	m.Id = uid.String()
	return nil
}
