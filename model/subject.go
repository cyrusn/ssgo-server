package model

// Subject store subject's information. The capacity of subject is the
// quota that student can be enrolled in.
type Subject struct {
	Code     string
	Group    int
	Name     string
	Cname    string
	Capacity int
}

type SubjestList []*Subject

// Insert insert subject information to database
func (s *Subject) Insert() error {
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
func (s *Subject) Get(subjectCode string) (*Subject, error) {
	row := db.QueryRow(
		"SELECT * FROM subject where code = ?",
		subjectCode,
	)
	if err := row.Scan(&s.Code, &s.Group, &s.Name, &s.Cname, &s.Capacity); err != nil {
		return nil, err
	}

	return s, nil
}

// AllSubjects return all subjects
func (list SubjestList) Get() (SubjestList, error) {
	rows, err := db.Query("SELECT * FROM subject")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := new(Subject)
		if err := rows.Scan(
			&s.Code, &s.Group, &s.Name, &s.Cname, &s.Capacity,
		); err != nil {
			return nil, err
		}

		list = append(list, s)
	}
	return list, nil
}

// UpdateCapacity update subject Capacity by subject Code
func (s *Subject) UpdateCapacity(subjectCode string, capacity int) error {
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
