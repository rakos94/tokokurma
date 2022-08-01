package helper_test

import (
	"testing"
	"tokokurma/helper"
)

func TestHashPassword(t *testing.T) {
	pwd, err := helper.HashPassword("1")
	if err != nil {
		t.Error(err)
	}

	t.Log(helper.CheckPasswordHash("1", pwd))
}
