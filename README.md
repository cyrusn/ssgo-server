# Rewrite Subject-Selection in golang

- an subject-selection web api



# Feature
## startup
There are only teacher and student user, *admin* are the one who do the setup of the programme. The setup setup step include the following:
  - create database for certain subject-selection event.
  - prepare student list for importing to database
  - basic setting of config file

Once the programme server started, student and teacher user can login to system, 2 different web interface will be launched by teacher and student
- student
  - fill up their priority
  - confirm submission
  - print out reply slip for sign back by parents
- teacher
  - can upload students rank for subject allocation
  - can adjust subject capacity for subject allocation
  - can allocate students preference if the above 2 information are provided
  - view student's application status (sort by isConfirmed, classCode, classNo ...)

# Reference

- [go-sql-driver/mysql: Go MySQL Driver is a lightweight and fast MySQL driver for Go's (golang) database/sql package](https://github.com/go-sql-driver/mysql)
- [gorilla/mux: A powerful URL router and dispatcher for golang.](https://github.com/gorilla/mux)
- [spf13/viper: Go configuration with fangs](https://github.com/spf13/viper)
- [spf13/cobra: A Commander for modern Go CLI interactions](https://github.com/spf13/cobra/)
- [Practical Persistence in Go: Organising Database Access](http://www.alexedwards.net/blog/organising-database-access)
- [Practical Persistence in Go: SQL Databases](http://www.alexedwards.net/blog/practical-persistence-sql)



# How to use this application
- Create new Database or use existing database
- prepare config file

``` toml
[server]
port=":5000"

# https://godoc.org/github.com/go-sql-driver/mysql#Config
[mysql]
user="root"
net="tcp"
dbname="ssapi"

[event]
name="test"
school-year="2016-17"

[json]
user="./data/user.json"
rank="./data/rank.json"
```

# Learn Test
- [leesei example](https://github.com/leesei/openslide-prop2json)

# using `create`

- prepare `user.json` with following KEY
  + In the `role` key in user.json, 1 for student, 2 for teacher, 3 for admin (enum in mysql)

```json
[{
  "username": "lpxxxxxxx",
  "password": "vosine76",
  "name": "testing one",
  "cname": "測試一",
  "role": 1,
  "classCode": "3A",
  "classNo": 1
}]
```


# setup event
- update all field in `config.toml`
- run the following command in order
  + `create`
  + `import`
  + `setup`
