package ntp

import (
	"time"

	"github.com/beevik/ntp"
)

type FTP interface {
	GetTime() (time.Time, error)
}

type usecaseFTP struct {
	servAdd string
}

func New(servAdd string) FTP {
	return &usecaseFTP{
		servAdd: servAdd,
	}
}

func (u *usecaseFTP) GetTime() (time.Time, error) {
	time, err := ntp.Time(u.servAdd)
	return time, err
}
