package service

import (
	"errors"
	"sync"

	"github.com/igor-sasaoka/brand-monitor-job/model"
)

var r *accountRegistry

type accountRegistry struct {
    mu sync.Mutex
    list map[model.Account]interface{}
    SendMail chan string 
}

func (r *accountRegistry) Add(a model.Account) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    if _, ok := r.list[a]; ok {
        return errors.New("Account already exists")
    }
    r.list[a] = struct{}{}

    go r.dispatch(a.Email)
    return nil
}

func (r *accountRegistry) dispatch(email string) {
    r.SendMail <- email
}

func GetAccountRegistry() *accountRegistry {
    if r == nil {
        r = &accountRegistry{
        	list:      make(map[model.Account]interface{}),
        	SendMail: make(chan string, 10),
        }
    }

    return r
}
