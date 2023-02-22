package test

import (
	"douyin/dao"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertUser(t *testing.T) {
	dao.Init()
	for i := 1; i <= 300; i++ {
		_, err := dao.NewUserDaoInstance().InsertUser(
			fmt.Sprintf("testname%d", i),
			fmt.Sprintf("testpassword%d", i),
		)
		if err != nil {
			t.Error(err)
		}
	}

}

func TestQueryUserById(t *testing.T) {
	dao.Init()
	user, err := dao.NewUserDaoInstance().QueryUserById(303)
	fmt.Printf("%#v", user)
	assert.Nil(t, err)
}

func TestQueryUserByUsername(t *testing.T) {
	dao.Init()
	user, err := dao.NewUserDaoInstance().QueryUserByUsername("admin")
	assert.NotNil(t, err)
	fmt.Printf("%#v", user)
}

func TestUpdateUserFollowCountAddOne(t *testing.T) {
	dao.Init()
	err := dao.NewUserDaoInstance().UpdateUserFollowCountAddOne(1)
	assert.Nil(t, err)
}

func TestUpdateUserFollowCountMinusOne(t *testing.T) {
	dao.Init()
	err := dao.NewUserDaoInstance().UpdateUserFollowCountMinusOne(1)
	assert.Nil(t, err)
}

func TestUpdateUserFollowerCountAddOne(t *testing.T) {
	dao.Init()
	err := dao.NewUserDaoInstance().UpdateUserFollowerCountAddOne(1)
	assert.Nil(t, err)
}

func TestUpdateUserFollowerCountMinusOne(t *testing.T) {
	dao.Init()
	err := dao.NewUserDaoInstance().UpdateUserFollowerCountMinusOne(1)
	assert.Nil(t, err)
}