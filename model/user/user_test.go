package user_test

// var userList = []model.User{
// 	model.User{"student1", "password1", "Alice", "愛麗絲", false},
// 	model.User{"student2", "password2", "Bob", "鮑伯", false},
// 	model.User{"student3", "password3", "Charlie", "查利", false},
// 	model.User{"teacher1", "password4", "Dave", "戴夫", true},
// 	model.User{"teacher2", "password5", "Eve", "伊夫", true},
// 	model.User{"teacher3", "password6", "Frank", "佛蘭克", true},
// }

// var TestUserTable = func(t *testing.T) {
// t.Run("Add users", TestInsertUser)
// t.Run("List All user", TestAllUsers)
// t.Run("Get user info", TestGetUser(1))
// t.Run("Validate user password", TestUser_Validate(1))
// }

// var TestInsertUser = func(t *testing.T) {
// 	for _, u := range userList {
// 		if err := db.InsertUser(u); err != nil {
// 			t.Fatal(err)
// 		}
// 	}
// }

// var TestAllUsers = func(t *testing.T) {
// 	users, err := db.AllUsers()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	for i, got := range users {
// 		diffUserTest(got, &userList[i], t)
// 	}
// }

// var TestGetUser = func(index int) func(*testing.T) {
// 	return func(t *testing.T) {
// 		username := userList[index].Username
// 		got, err := db.GetUser(username)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		diffUserTest(got, &userList[index], t)
// 	}
// }

// var TestUser_Validate = func(index int) func(*testing.T) {
// 	return func(t *testing.T) {
// 		user := userList[index]
// 		got, err := db.GetUser(user.Username)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		if err := got.Validate(user.Password); err != nil {
// 			t.Fatal(err)
// 		}
// 	}
// }

// func diffUserTest(got, want *model.User, t *testing.T) {
// 	hash := []byte(got.Password)
// 	password := []byte(want.Password)
// 	err := bcrypt.CompareHashAndPassword(hash, password)
// 	if err != nil {
// 		fmt.Println(got, want)
// 		t.Fatal(err)
// 	}

// 	got.Password = want.Password
// 	diffTest(want, got, t)
// }
