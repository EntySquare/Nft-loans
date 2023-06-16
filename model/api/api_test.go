package api

import (
	"nft-loans/database"
	"testing"
)

func TestRunP(t *testing.T) {
	database.ConnectDB()

	IncomeRunP(database.DB)
	CovenantCycle(database.DB)

}
