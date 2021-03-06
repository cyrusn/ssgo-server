package subject

import "database/sql"

type DB struct {
	*sql.DB
}

// Subject store subject's information. The capacity of subject is the
// quota that student can be enrolled in.
type Subject struct {
	Code     string `json:"code"`
	Capacity int    `json:"capacity"`
}

// Insert insert subject information to database
func (db *DB) Insert(s *Subject) error {
	_, err := db.Exec(`
    INSERT INTO Subject (
      code, capacity
    ) values (
      ?, ?
    )
    `,
		s.Code, s.Capacity,
	)
	return err
}

// Get return Subject by subject code
func (db *DB) Get(subjectCode string) (*Subject, error) {
	row := db.QueryRow(
		"SELECT * FROM Subject where code = ?",
		subjectCode,
	)
	s := new(Subject)
	if err := row.Scan(&s.Code, &s.Capacity); err != nil {
		return nil, err
	}

	return s, nil
}

// List return all subjects
func (db *DB) List() ([]*Subject, error) {
	rows, err := db.Query("SELECT * FROM Subject")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*Subject

	for rows.Next() {
		s := new(Subject)
		if err := rows.Scan(
			&s.Code, &s.Capacity,
		); err != nil {
			return nil, err
		}

		list = append(list, s)
	}
	return list, nil
}

// UpdateCapacity update subject Capacity by subject Code
func (db *DB) UpdateCapacity(subjectCode string, capacity int) error {
	_, err := db.Exec(`
		UPDATE Subject set
			capacity = ?
		where code = ?;
		`,
		capacity,
		subjectCode,
	)

	return err
}
