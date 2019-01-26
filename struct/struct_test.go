package gostruct

import "testing"

func TestHasField(t *testing.T) {
	type User struct {
		Name string
	}

	var u User

	t.Log(HasField(u, "Name"))  // true
	t.Log(HasField(&u, "Name")) // true

	t.Log(HasField(&u, "name"))       // false
	t.Log(HasField("string", "name")) // false
}
