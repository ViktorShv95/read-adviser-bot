package storage

import (
	"crypto/sha1"
	"fmt"
	"io"

	"github.com/ViktorShv95/read-adviser-bot/lib/e"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("cannot calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("cannot calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}