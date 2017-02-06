package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	//"log"
	"testing"
)

func TestCanGetCourseData(t *testing.T) {
	db, err := gorm.Open("sqlite3", "data/test")
	if err != nil {
	}
	accessor := &DatabaseAccessor{db}

	result := accessor.GetCourse("3001 200")

	assert.NotNil(t, result, "blarg")
	//log.Println(result)
}

func TestCanGetClassData(t *testing.T) {
	db, err := gorm.Open("sqlite3", "data/test")
	if err != nil {
	}
	accessor := &DatabaseAccessor{db}

	result := accessor.GetClasses("3001 200")

	assert.NotNil(t, result, "blarg")
	//log.Println(result)
}

func TestGetClassesGetsDays(t *testing.T) {
	db, err := gorm.Open("sqlite3", "data/test")
	if err != nil {
	}
	accessor := &DatabaseAccessor{db}

	result := accessor.GetClasses("3460 210")

	assert.True(t, result[0].Times.Monday, "3460 210 Monday should be true")
	//log.Println(result) //[0].Monday)

}

func TestGetClassesGetsTime(t *testing.T) {
	db, err := gorm.Open("sqlite3", "data/test")
	if err != nil {
	}
	accessor := &DatabaseAccessor{db}

	_ = accessor.GetClasses("3460 210")

	//assert.True(t, result[0].Monday, "3460 210 Monday should be true")
	//log.Println(result[0].Monday)

}
