package contracts

import (
	"fmt"
	"testing"
)

func Test3233221t(t *testing.T) {
	//c, _, err := NewContractsApi()
	//if err != nil {
	//	panic(err)
	//}
	//i := 0
	//for {
	//	i++
	//	list, err := GetAddressLoansList("0xFFcf8FDEE72ac1105c542428B35EEF5769C409f0")
	//	fmt.Println(i, err, len(list))
	//}

	list, err := GetAddressLoansList("0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0")
	fmt.Println(err, len(list))

	for _, v := range list {
		fmt.Println(v.ReceivedTime.Int64())
	}
	//list, err := GetAddressLoansList("0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0")
	//fmt.Println(err)

	//for k, v := range list {
	//	fmt.Println(k, v.ReceivedTime)
	//}
}
