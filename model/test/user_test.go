package test

import (
	"nft-loans/database"
	"nft-loans/model/api"
	"testing"
)

func TestInsertUser(t *testing.T) {
	database.ConnectDB()

	api.SelectChainData(database.DB, 6)
}
