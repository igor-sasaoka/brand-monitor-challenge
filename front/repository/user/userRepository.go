package userRepository

import (
	"errors"

	"github.com/igor-sasaoka/brand-monitor-front/database"
	"github.com/igor-sasaoka/brand-monitor-front/model"
)

func Create(u *model.User) {
    database.DB().Create(u)
}

func GetUserBySession(s string) (*model.User, error) {
    var u model.User
    database.DB().First(&u, "session = ?", s)
    if u.ID == 0 {
        return nil, errors.New("User not found")
    }

    return &u, nil
}
