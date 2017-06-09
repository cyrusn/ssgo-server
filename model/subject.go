package model

// Subject store subject's information, where subject.Capacity is the capacity
// for subject allocation (Quota that student can enrol to this subject.)
type Subject struct {
	Code     string
	Group    int
	Name     string
	Cname    string
	Capacity int
}

const subjectTableSchema = `
CREATE TABLE IF NOT EXISTS subject (
  code TEXT NOT NULL,
  gp INTEGER NOT NULL,
  name TEXT NOT NULL,
  cname TEXT NOT NULL,
  capacity INTEGER NOT NULL
);`

// CreateSubjectTable create subject table
func (db *DB) CreateSubjectTable() error {
	return db.createTable(subjectTableSchema)
}

// InsertSubject insert subject subject information to subject database
func (db *DB) InsertSubject(s Subject) error {
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

// GetSubject return Subject by subject code
func (db *DB) GetSubject(subjectCode string) (*Subject, error) {
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

// AllSubjects return Subject by subject code
func (db *DB) AllSubjects() ([]*Subject, error) {
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

// UpdateSubjectCapacity update subject Capacity by subject Code
func (db *DB) UpdateSubjectCapacity(subjectCode string, capacity int) error {
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
