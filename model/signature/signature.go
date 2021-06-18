package signature

import (
	"database/sql"
	"log"
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
		nil,
	)

	return err
}

// UpdateIsSigned will update student's isSigned.
func (db *DB) UpdateIsSigned(userAlias string, isSigned bool) error {
	log.Println("start")
	_, err := db.Exec(
		"UPDATE Signature set isSigned = ? WHERE userAlias = ?",
		isSigned,
		userAlias,
	)
	log.Println("err")
	return err
}

// UpdateSignatureAddress will update parents' signature
func (db *DB) UpdateAddress(userAlias string, address string) error {
	_, err := db.Exec(
		"UPDATE Student set address = ? WHERE userAlias = ?",
		address,
		userAlias,
	)
	return err
}
