package subject

import "github.com/cyrusn/ssgo/model/ssdb"

// Subject store subject's information. The capacity of subject is the
// quota that student can be enrolled in.
type Subject struct {
	Code     string
	Group    int
	Name     string
	Cname    string
	Capacity int
}

// Insert insert subject information to database
func Insert(db *ssdb.DB, s Subject) error {
	_, err := db.Exec(`
    INSERT INTO subject (
      code, gp, name, cname, capacity
    ) values (
      ?, ?, ?, ?, ?
    )
    `,
		s.Code, s.Group, s.Name, s.Cname, s.Capacity)
	return err
}

// Get return Subject by subject code
func Get(db *ssdb.DB, subjectCode string) (*Subject, error) {
	row := db.QueryRow(
		"SELECT * FROM subject where code = ?",
		subjectCode,
	)
	s := new(Subject)
	if err := row.Scan(&s.Code, &s.Group, &s.Name, &s.Cname, &s.Capacity); err != nil {
		return nil, err
	}

	return s, nil
}

// All return all subjects
func All(db *ssdb.DB) ([]*Subject, error) {
	rows, err := db.Query("SELECT * FROM subject")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var subjects []*Subject

	for rows.Next() {
		s := new(Subject)
		if err := rows.Scan(
			&s.Code, &s.Group, &s.Name, &s.Cname, &s.Capacity,
		); err != nil {
			return nil, err
		}

		subjects = append(subjects, s)
	}
	return subjects, nil
}

// UpdateCapacity update subject Capacity by subject Code
func UpdateCapacity(db *ssdb.DB, subjectCode string, capacity int) error {
	_, err := db.Exec(`
		UPDATE subject set
			capacity = ?
		where code = ?;
		`,
		capacity,
		subjectCode,
	)

	return err
}
