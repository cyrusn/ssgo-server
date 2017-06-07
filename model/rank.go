package model

// Rank store the ranking of student
type Rank struct {
	Username string
	Ranking  int
}

const rankTableSchema = `
CREATE TABLE IF NOT EXISTS rank (
	username TEXT UNIQUE NOT NULL,
  ranking INT NOT NULL
);
`

// CreateRankTable create rank table
func (db *DB) CreateRankTable() error {
	return db.createTable(rankTableSchema)
}

// TruncateRankTable will remove all data in the rank table, run this command
// before import a list of student rank. Update student rank is not recommended,
// admin should truncate the rank table and import all student rank in whole.
func (db *DB) TruncateRankTable() error {
	_, err := db.Exec("TRUNCATE Table rank")
	return err
}

// GetStudentRank query student's rank by username
func (db *DB) GetStudentRank(username string) (*Rank, error) {
	row := db.QueryRow("SELECT * FROM rank where username = ?;", username)

	r := new(Rank)
	err := row.Scan(&r.Username, &r.Ranking)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// InsertStudentRank insert username and ranking to rank table
func (db *DB) InsertStudentRank(r Rank) error {
	_, err := db.Exec(`INSERT INTO rank (
    username,
    ranking
  ) values (?, ?)`,
		r.Username,
		r.Ranking,
	)
	return err
}
