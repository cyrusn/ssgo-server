# Subject Selection System (Golang)

- an subject-selection web api server

## Documentation
- run `godoc -play -http:5050`
- access the link [src/github.com/cyrusn/ssgo/route/route.go - The Go Programming Language](http://localhost:5050/src/github.com/cyrusn/ssgo/route/route.go?s=545:577#L18)

## Startup
There are 3 roles of user in this system *STUDENT*, TEAHCER* and *ADMIN* user.
To startup a new subject system event, please follow the following steps.
  - create new database by using `create` command.
  - import subjects by using `import` command with `subject` as subcommand.
  - import student users by using `import` command with `student` as subcommand.
  - import teacher users by using `import` command with `teacher` as subcommand.
  - the schema of JSON files for the import commands, please see be session *Schema* below.
  - start server by using `serve` command

## Schemas

```json
// teacher.json
// system admin have to declare the role of teacher user (either TEACHER or STUDENT).
// other information of users should fetch in front end program
[{
  "userAlias": "string",
  "password": "string",
  "role": "ADMIN"
 }, {
  "userAlias": "string",
  "password": "string",
  "role": "TEACHER"
  }
]
```

```json
// student.json
// other information of users should fetch in front end program
[
  {
    "userAlias": "string",
    "password": "string"
  }
]
```

```json
// subject.json
// only subject code are required, the program only store the subject's capacity.
[
  "bio", "bafs", "chist", "phy", "ths", "va",
  "chem", "cscb", "econ", "geog", "hist", "ict"
]
```
