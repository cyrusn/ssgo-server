/*
Package model provide an interface to manage the table in database for ssgo

for creating corresponding database and table, please refer to `package db`

all exported functions in package model only handle "query", "update" and "delete"
actions in database

- USER store the information for all users in `ssgo`, including Student, Teacher and Admin
  - username strign
  - password string (encrypted string)
  - ename string
  - cname string
  - isTeacher bool

- STUDENT
  - username string
  - classCode string
  - classNo int
  - priority []int (BLOB in sqlite3)
  - isConfirmed bool (int in sqlite3)

- RANK store the academic ranking for each Registry for certain Submission
  - username
  - ranking int

- CAPACITY store the capacity (number of student can be enrolled to this subject)
  for each subject on each Submission
  - subjectID string
  - capacity int
*/
package model
