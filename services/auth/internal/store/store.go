package store

import (
	"errors"

	"gorm.io/gorm"
)

// Временные затычки

var (
	ErrNotFound  = errors.New("not found")
	ErrNoChanges = errors.New("no changes")
)

type Store struct{ db *gorm.DB }

func New(db *gorm.DB) *Store { return &Store{db: db} }
