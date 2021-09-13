package main

import (
	"fmt"
	myUtils "myProject/myUtils"
	"time"
)

func main() {
	myUtils.InsertIntoCityWithTx()
	//fmt.Println("--------------Getting data from Internat-------------------")
	//myUtils.GetJasonFromInternet()
	//fmt.Println("---------------Got the data from Internet------------------")
	var beforeInv, afterInv int64
	beforeInv = changeToMillisecond()
	fmt.Println("--------------Getting data from Cities Table-------------------")
	myUtils.GetDataFromMySQL()
	afterInv = changeToMillisecond()
	fmt.Println("--------------------------------------------------------------", (afterInv - beforeInv))
	//myUtils.DeleteDataFromMySQL()

}
func changeToMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
