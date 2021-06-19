package signature

import (
	"database/sql"
)

type DB struct {
	*sql.DB
}

// Student stores information for student user.
type Signature struct {
	UserAlias string `json:"userAlias"`
	IsSigned  bool   `json:"isSigned"`
	Address   string `json:"address"`
}

// Get query signature by userAlias.
func (db *DB) Get(userAlias string) (*Signature, error) {
	s := new(Signature)

	row := db.QueryRow("SELECT * FROM Signature where userAlias = ?", userAlias)
	err := s.scanSignature(row)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Insert add new student to database.
func (db *DB) Insert(userAlias string) error {

	_, err := db.Exec(
		`INSERT INTO Signature (
			userAlias,
			isSigned,
			address
		) values (?, ?, ?)`,
		userAlias,
		false,
		"",
	)

	return err
}

// UpdateIsSigned will update student's isSigned.
func (db *DB) UpdateIsSigned(userAlias string, isSigned bool) error {
	_, err := db.Exec(
		"UPDATE Signature set isSigned = ? WHERE userAlias = ?",
		isSigned,
		userAlias,
	)
	return err
}

// UpdateSignatureAddress will update parents' signature
func (db *DB) UpdateAddress(userAlias string, address string) error {
	_, err := db.Exec(
		"UPDATE Signature set address = ? WHERE userAlias = ?",
		address,
		userAlias,
	)
	return err
}

// List get all signatures.
func (db *DB) List() ([]*Signature, error) {
	var list []*Signature
	rows, err := db.Query("SELECT * FROM Signature")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		s := new(Signature)
		err := s.scanSignature(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, s)
	}

	return list, nil
}
