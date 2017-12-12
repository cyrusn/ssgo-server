package model

import (
	"database/sql"
)

type StudentDB struct{ *sql.DB }
type TeacherDB struct{ *sql.DB }
type SubjectDB struct{ *sql.DB }

type Repository struct {
	DB *sql.DB
	StudentDB
	TeacherDB
	SubjectDB
}

// NewRepository startup DB for model package
func NewRepository(dbPath string) *Repository {
	var err error
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}

	// Ping also establish a connection if necessary
	if err = db.Ping(); err != nil {
		panic(err)
	}

	return &Repository{
		db,
		StudentDB{db},
		TeacherDB{db},
		SubjectDB{db},
	}
}
