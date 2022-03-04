package gateways

import (
	"log"
	"math/rand"
	"sp/src/domains/entities"
	"sp/src/usecases/port"
	"time"

	"github.com/oklog/ulid"
	"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) port.UserRepository {
	return &userRepository{
		Conn: conn,
	}
}

func (ur *userRepository) FindByID(id string) (*entities.User, error) {
	var user = &entities.User{}
	user.ID = id
	err := ur.Conn.First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) Create(u *entities.User) (*entities.User, error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	u.ID = ulid.MustNew(ulid.Timestamp(t), entropy).String()
	log.Print(u)
	err := ur.Conn.Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *userRepository) Update(u *entities.User) (user *entities.User, err error) {
	err = ur.Conn.Save(u).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
