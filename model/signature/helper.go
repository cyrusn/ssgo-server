package signature

import (
	"database/sql"
)

func (s *Signature) scanSignature(v interface{}) error {
	var err error

	var args = []interface{}{
		&s.UserAlias,
		&s.IsSigned,
		&s.Address,
	}

	switch t := v.(type) {
	case *sql.Row:
		err = t.Scan(args...)
	case *sql.Rows:
		err = t.Scan(args...)
	}

	return err
}
