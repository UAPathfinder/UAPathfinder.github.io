package main

import (
"testing"
)


func TestGetCourses(t *testing.T){
	output := GetCourses()
	value := output[1]
	if value != 2{
		t.Fatalf("Get Courses is not returning the testing output: output[1] = 2")
	} 
}

