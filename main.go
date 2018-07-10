package main

import (
	"fmt"
	"megaplan"
)

func main() {
	meg := megaplan.Megaplan{}
	err := meg.ConnectMegaplan("https://megaplan.ath.ru", "bob_edition@mail.ru", "zaq12wsx#")
	//err := meg.ConnectMegaplan("https://megaplan.ath.ru", "azhokhov@ath.ru", "Gfhjkm henf*")
	fmt.Println("err: ", err)
	//res := meg.ListEventCategory()
	//fmt.Println("res: ", res)
	//for i := 0; i <= 13; i++ {
	getparticipants := getParticipants()
	for ok := true; ok; {
		resList := meg.ListEvent(0, 100, 0, false, "")
		if len(resList) < 100 {
			ok = false
		}
		for _, data := range resList {
			participants := getparticipants(data.Participants)
			fmt.Println("participants: ", participants)
		}
	}
	//fmt.Println("res: ", e)
	//megaplan.Test()
}

func getParticipants() func(Participants []megaplan.Participant) int {
	db := megaplan.Participants{}
	function := func(Participants []megaplan.Participant) int {
		return len(Participants)
	}
	return function
}
