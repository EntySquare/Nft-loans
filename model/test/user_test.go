package test

import (
	"fmt"
	"nft-loans/database"
	"nft-loans/model"
	"testing"
)

func TestInsertUser(t *testing.T) {
	database.ConnectDB()

	err := database.DB.Create(&model.User{
		Token:       "test123123123",
		RecommendId: 3,
		Flag:        "1",
	}).Error
	fmt.Println(err)
}
