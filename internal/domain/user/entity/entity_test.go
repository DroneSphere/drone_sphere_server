package entity

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestUser_HashPassword(t *testing.T) {
	u := &User{
		Password: "123456",
	}
	if err := u.HashPassword(); err != nil {
		t.Errorf("expect nil, got %v", err)
	}
	expected := "hashed:123456"
	assert.Equal(t, u.Password, expected)
}

func TestUser_Authenticate(t *testing.T) {
	u := &User{
		Password: "hashed:123456",
	}

	assert.Equal(t, u.Authenticate("123456"), true)
	assert.Equal(t, u.Authenticate("1234567"), false)
	assert.Equal(t, u.Authenticate(""), false)
}
