package student

import (
	"database/sql"
	"encoding/json"
)

func (s *Student) scanStudent(v interface{}) error {
	var priorities []byte
	var err error

	var args = []interface{}{
		&s.UserAlias,
		&s.ClassCode,
		&s.ClassNo,
		&priorities,
		&s.IsX3,
		&s.IsConfirmed,
		&s.Rank,
		&s.Timestamp,
		&s.Name,
		&s.Cname,
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

	return nil
}
