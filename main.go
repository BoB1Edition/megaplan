package main

import (
	"fmt"
	"megaplan"
)

func main() {
	meg := megaplan.Megaplan{}
	err := meg.ConnectMegaplan("https://megaplan.ath.ru", "bob_edition@mail.ru", "zaq12wsx#")
	fmt.Println("err: ", err)
	//res := meg.ListEventCategory()
	//fmt.Println("res: ", res)
	//for i := 0; i <= 13; i++ {
	resList := meg.ListEvent(0, 100, 0, false, "")
	r := resList.Data.([]interface{})
	//e := r[10]
	for key, data := range r {
		fmt.Println("key: ", key)
		fmt.Println("---------------------------------------------------------------")
		fmt.Println("data: ", data)
		fmt.Println("---------------------------------------------------------------")
	}
	//}
	//fmt.Println("res: ", e)
	//megaplan.Test()
}
