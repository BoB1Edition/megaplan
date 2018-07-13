package main

import (
	"fmt"
	"megaplan"
	"time"

	"github.com/tealeg/xlsx"
)

const (
	limit = 100
)

type AnswerStruct struct {
	Name        string
	Date        string
	Description string
	Clients     string
}

func main() {
	meg := megaplan.Megaplan{}
	err := meg.ConnectMegaplan("https://megaplan.ath.ru", "bob_edition@mail.ru", "zaq12wsx#")
	fmt.Println("err: ", err)
	offset := 0
	e := megaplan.Events{}
	for ok := true; ok; {
		resList := meg.ListEvent(0, limit, offset, false, "")
		if len(resList) < limit {
			ok = false
		}
		for _, data := range resList {
			e[data.Id] = data
		}
		offset += limit
	}
	ev := e.Filter(filtered)
	createXls(ev, &meg)
}

func createXls(events megaplan.Events, meg *megaplan.Megaplan) {
	file := xlsx.NewFile()
	defer file.Save("report.xlsx")
	/*sales, err := file.AddSheet("Sales")
	if err != nil {
		fmt.Println(err)
	}
	cgm, err := file.AddSheet("CGM")
	if err != nil {
		fmt.Println(err)
	}*/
	/*call, err := file.AddSheet("Звонки")
	if err != nil {
		fmt.Println(err)
	}
	meeting, err := file.AddSheet("Встречи")
	if err != nil {
		fmt.Println(err)
	}*/
	getParticipantInfo := getfParticipantInfo(meg)
	for key, event := range events {
		fmt.Println("key: ", key)
		getParticipantInfo(event.Owner)
		for _, part := range event.Participants {
			fmt.Println("event: ", event.Name, "part: ", part.Name)
			getParticipantInfo(part.Id)
		}
		/*fmt.Println("event.Name: ", event.Name)
		fmt.Println("event.Description: ", event.Description)
		fmt.Println("event.Place: ", event.Place)
		fmt.Println("event.EventCategory: ", event.EventCategory)
		switch event.EventCategory {
		case "Встреча":
			row := meeting.AddRow()

			//row.AddCell().SetString(event.Name)
			//row.AddCell().SetString(event.Description)
			row.WriteStruct(&answers, -1)
		case "Звонок":
			row := call.AddRow()
			row.AddCell().SetString(event.Name)
			row.AddCell().SetString(event.Description)
		}*/

	}
}

func filtered(e megaplan.Event) bool {
	t := time.Time{}
	now := time.Now()
	var err error
	if len(e.StartTime) < 26 {
		t, err = time.Parse("2006-01-02 15:04:05", e.StartTime)
	} else {
		t, err = time.Parse("2006-01-02 15:04:05 -07:00", e.StartTime)
	}
	if err != nil {
		fmt.Println("err: ", err)
		return false
	}
	if int(t.Month()) == int(now.Month())-1 {
		return true
	} else {
		return false
	}
}

func getfParticipantInfo(Connection *megaplan.Megaplan) func(id int) megaplan.Employee {
	employees := megaplan.Employees{}
	f := func(id int) megaplan.Employee {
		fmt.Println("employees[", id, "]: ", employees[id])
		return megaplan.Employee{}
	}
	return f
}
