package student

import (
	"database/sql"
	"encoding/json"
)

func (s *Student) scanStudent(v interface{}) error {
	var priorities []byte
	var isConfirmed int
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

	s.IsConfirmed = convertInt2Bool(isConfirmed)

	return nil
}

func convertBool2Int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func convertInt2Bool(i int) bool {
	if i == 0 {
		return false
	}
	return true
}
