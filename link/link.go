package link

import (
	"encoding/json"
	"errors"

	"github.com/adelowo/mapped/utils"
)

var (
	errLinkNotFound = errors.New("link does not exists")
)

type (
	// Link is an abstraction
	Link string

	// Mapped is a representation of a transformed url
	Mapped struct {
		Original  Link            `json:"original"`
		TO        Link            `json:"to"`
		CreatedAt utils.Timestamp `json:"created_at"`
	}

	Repository interface {
		Find(Link) (Mapped, error)
		Create(Mapped) error
	}
)

func IsLinkNotFoundError(err error) bool {
	return err == errLinkNotFound
}

func (m Mapped) Bytes() []byte {
	buf, _ := json.Marshal(m)
	return buf
}

func BytesToMapped(buf []byte) (Mapped, error) {
	var m Mapped
	return m, json.Unmarshal(buf, &m)
}
