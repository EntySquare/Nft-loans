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
		Token:         "test123123123",
		RecommendCode: "test",
		RecommendId:   3,
		Password:      "test",
		Phone:         "13000000004",
		Flag:          "1",
		Area:          "86",
		GoogleFlag:    "0",
		GoogleKey:     "123123123",
	}).Error
	fmt.Println(err)
}
