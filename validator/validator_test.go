package govalidator

import "testing"

func TestIsEmail(t *testing.T) {
	if !IsEmail("ikaiguang@uuff.com") {
		t.Error("something write wrong")
	}
}
