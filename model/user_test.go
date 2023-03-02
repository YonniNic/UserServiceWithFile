package model

import (
	"fmt"
	"testing"
)

type addTestCreate struct {
	id int
	//createdUser UserDefinition
	expected UserDefinition
}

var addTestsCreate = []UserDefinition{
	UserDefinition{Username: "test1", Password: "test1", Role: "test1"},
	UserDefinition{Username: "test2", Password: "test2", Role: "test2"},
	UserDefinition{Username: "test3", Password: "test3", Role: "test3"},
}

var user UserDefinition

var coll map[int]UserDefinition

func TestCreate(t *testing.T) {
	var expected = []addTestCreate{
		addTestCreate{1, UserDefinition{Username: "test1", Password: "test1", Role: "test1"}},
		addTestCreate{2, UserDefinition{Username: "test2", Password: "test2", Role: "test2"}},
		addTestCreate{3, UserDefinition{Username: "test3", Password: "test3", Role: "test3"}},
	}

	for i := 0; i < len(addTestsCreate); i++ {
		output := user.Create(addTestsCreate[i])
		if output != expected[i].expected {
			t.Errorf("Output %q not equal to expected %q", output, user.Users()[i])

		}
	}

}

func TestAGetUserById(t *testing.T) {
	var expected = []addTestCreate{
		addTestCreate{0, UserDefinition{Username: "test1", Password: "test1", Role: "test1"}},
		addTestCreate{1, UserDefinition{Username: "test2", Password: "test2", Role: "test2"}},
		addTestCreate{2, UserDefinition{Username: "test3", Password: "test3", Role: "test3"}},
	}

	for i := 0; i < len(addTestsCreate); i++ {
		user.Create(addTestsCreate[i])
	}

	if output := user.GetUserById(3); output != expected[2].expected {
		t.Errorf("Output %q not equal to expected %q", output, expected[2].expected)
	}
	fmt.Println(user.Users())
}

func TestUpdate(t *testing.T) {
	var expected = []addTestCreate{
		addTestCreate{1, UserDefinition{Username: "test4", Password: "test4", Role: "test4"}},
		addTestCreate{2, UserDefinition{Username: "test5", Password: "test5", Role: "test5"}},
		addTestCreate{3, UserDefinition{Username: "test6", Password: "test6", Role: "test6"}},
	}

	var ToUpdateWith = []UserDefinition{
		UserDefinition{Username: "test4", Password: "test4", Role: "test4"},
		UserDefinition{Username: "test5", Password: "test5", Role: "test5"},
		UserDefinition{Username: "test6", Password: "test6", Role: "test6"},
	}

	for i := 0; i < len(addTestsCreate); i++ {
		user.Create(addTestsCreate[i])

	}

	if user.Update(0, ToUpdateWith[0]); expected[0].expected != ToUpdateWith[0] {
		t.Errorf("Output %q not equal to expected %q", ToUpdateWith[0], expected[0].expected)
	}

}

func TestDelete(t *testing.T) {

	var expected = []addTestCreate{
		addTestCreate{0, UserDefinition{Username: "test1", Password: "test1", Role: "test1"}},
		addTestCreate{1, UserDefinition{Username: "test2", Password: "test2", Role: "test2"}},
		addTestCreate{2, UserDefinition{Username: "test3", Password: "test3", Role: "test3"}},
	}

	for i := 0; i < len(addTestsCreate); i++ {
		user.Create(addTestsCreate[i])

	}

	if output := user.Delete(2); expected[1].expected != output {
		t.Errorf("Output %q exists", output)

	}

}

func TestGetAllUsers(t *testing.T) {
	var expected = []addTestCreate{
		addTestCreate{0, UserDefinition{Username: "test1", Password: "test1", Role: "test1"}},
		addTestCreate{1, UserDefinition{Username: "test2", Password: "test2", Role: "test2"}},
		addTestCreate{2, UserDefinition{Username: "test3", Password: "test3", Role: "test3"}},
	}

	for i := 0; i < len(addTestsCreate); i++ {
		user.Create(addTestsCreate[i])

	}

	for i := 1; i <= len(addTestsCreate); i++ {
		if output := user.Users(); output[i] != expected[i-1].expected {
			t.Errorf("Output %q not equal to expected %q", output, expected[i-1].expected)
		}
	}
}
