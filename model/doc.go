/*
Package model provide an interface to manage the table in database for ssgo

TEACHER Table

TEACHER store the information for all users in `ssgo`, including Student, Teacher and Admin

  - username strign
  - password string (encrypted string)
  - name string
  - cname string

STUDENT Table

STUDENT Table store further information for student user

  - username strign
  - password string (encrypted string)
  - name string
  - cname string
  - classCode string
  - classNo int
  - priority []int (BLOB in sqlite3)
  - isConfirmed bool (int in sqlite3)

SUBJECT Table

SUBJECT Table store the capacity (number of student can be enrolled to this subject)
for each subject on each Submission

  - code string
  - group int
  - name string
  - cname string
  - capacity int
*/
package model
