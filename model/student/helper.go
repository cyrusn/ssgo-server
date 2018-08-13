package student

import (
	"database/sql"
	"encoding/json"
)

func (s *Student) scanStudent(v interface{}) error {
	var priorities []byte
	var isConfirmed bool
	var err error

	var args = []interface{}{
		&s.UserAlias,
		&priorities,
		&isConfirmed,
		&s.Rank,
	}

	switch t := v.(type) {
	case *sql.Row:
		err = t.Scan(args...)
	case *sql.Rows:
		err = t.Scan(args...)
	}

	if err != nil {
		return err
	}

	if err := json.Unmarshal(priorities, &s.Priorities); err != nil {
		return err
	}

	s.IsConfirmed = isConfirmed

	return nil
}
