package model_test

import (
	"log"
	"os"
	"testing"
)

const (
	DBPath = "../database/test.db"
)

func init() {
	cleanUpDB(DBPath)
}

func Test(t *testing.T) {
	t.Run("[Panic] CreateTables", PanicTestCreateTables)
	t.Run("[Panic] InitDB", PanicTestInitDB)
	t.Run("InitDB", TestInitDB(DBPath))
	t.Run("CreateTables", TestCreateTable)

	// Test Teacher
	t.Run("Insert teacher user", TestTeacher_Insert)
	t.Run("Insert teacher user", TestTeacher_Get)

	// Test Students
	t.Run("Insert all students", TestStudent_Insert)
	t.Run("Insert duplicated student", TestStudent_Insert_Errors)
	t.Run("List all students", TestStudent_All)
	t.Run("Get student", TestStudent_Get)
	t.Run("Update student1 priority", TestStudent_UpdatePriority)
	t.Run("Update student1 isConfirmed", TestStudent_UpdateIsConfirmed)
	t.Run("Get student", TestStudent_Get)

	// Test Subject
	t.Run("Insert subjects", TestInsert)
	t.Run("Get all subjects", TestAll)
	t.Run("Get subject by subject code (bio)", TestGet(0))
	t.Run("Update bafs's Capacity to 20'", TestUpdateCapacity(1, 20))
	t.Run("check if bafs's capacity updated", TestGet(1))
}

// cleanup remove DB if it is exist and create a new empty database
func cleanup(DBPath string) {
	if _, err := os.Stat(DBPath); os.IsExist(err) {
		if err := os.Remove(DBPath); err != nil {
			log.Fatal(err)
		}
	}
	os.Create(DBPath)
}
