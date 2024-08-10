package structs

import (
	"errors"
	"unicode/utf8"
)

type News struct {
	Id         *int64   `json:"Id"`
	Title      *string  `json:"Title"`
	Content    *string  `json:"Content"`
	Categories *[]int64 `json:"Categories"`
}

func (n *News) CheckId() error {
	if n.Id == nil {
		return errors.New("Id field is not set")
	}
	return nil
}

func (n *News) CheckTitle() error {
	if n.Title == nil {
		return errors.New("Title field is not set")
	}
	l := utf8.RuneCountInString(*n.Title)
	if l > 256 {
		return errors.New("Title size must be no more than 256 characters")
	}
	if l == 0 {
		return errors.New("Title is empty")
	}

	return nil
}

func (n *News) CheckContent() error {
	if n.Content == nil {
		return errors.New("Content field is not set")
	}
	if len(*n.Content) == 0 {
		return errors.New("Content is empty")
	}
	return nil
}
