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

	list, err := GetAddressLoansList("0x7a6c19A76Ac5866cC8a82fb5F1E57e09aaF2E416")
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

func Test3233221t121(t *testing.T) {
	hash, s, s2, f, err := CheckHash("0x30d763fbc14d4a601c15aa1ef0645eb6daa5e2b37a45fb1f522fccffa1145b6e")
	if err != nil {
		panic(err)
	}
	fmt.Println(hash, s, s2, f, err)
}
