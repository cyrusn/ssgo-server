# Subject Selection System (Golang)

- an subject-selection web api server

## Startup

`go build -o ssgo main.go`

There are 3 types of user in this system. i.e. **STUDENT**, **TEACHER** and **ADMIN**.  
To startup a new subject system event, please follow the following steps:

- prepare `config.yaml` and users JSON files, please see the session [Schemas](#schemas) below
- build frontend and place in `static` in `config.yaml`
- create a MySQL database, place database dsn (uri) in `config.yaml`
- [TODO] init project with `./ssgo init`, which does the followings:
  - create database by `./ssgo create`
  - import subjects by `./ssgo import subject`
  - import student users by `./ssgo import student`
  - import teacher users by `./ssgo import teacher`
- start server by `./ssgo serve`

### Build for linux (Digital Ocean)

`GOOS=linux GOARCH=amd64 go build -o ./dist/ssgo main.go && scp -r ./dist/ssgo root@calp:~/ssgo/`

## Schemas

### `config.yaml`

```yaml
# default value of config.yaml
key: "skill-vein-planet-neigh-envoi" // JWT signature secret
dsn: "root@/ssgoTestDB"
overwrite: false
teacher: "./data/teacher.json"
student: "./data/student.json"
subject: "./data/subject.json"
port: ":5000"
static: "./public"
# the life time of jwt-token in minutes
time: 30
```

### `teacher.json`

```json
// teacher.json
// system admin have to declare the role of teacher user (either TEACHER or STUDENT).
// other information of users should fetch in front end program
[
  {
    "userAlias": "string",
    "password": "string",
    "role": "ADMIN"
  },
  {
    "userAlias": "string",
    "password": "string",
    "role": "TEACHER"
  }
]
```

### `student.json`

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

### `subject.json`

```json
// subject.json
// only subject code are required, the program only store the subject's capacity.
[
  "bio",
  "bafs",
  "chist",
  "phy",
  "ths",
  "va",
  "chem",
  "bio2", // cscb => bio2
  "econ",
  "geog",
  "hist",
  "hmsc",
  "ict"
]
```
