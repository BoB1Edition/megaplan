package main

import (
	"fmt"
	"log"
	"megaplan"
	"time"
	"xlsWrite"

	"github.com/scorredoira/email"
	"github.com/tealeg/xlsx"
)

const (
	limit = 100
)

type AnswerStruct struct {
	Name        string
	Type        string
	Date        string
	Participant string
	Description string
	Clients     string
	LocalGlobal string
}

func main() {
	meg := megaplan.Megaplan{}
	err := meg.ConnectMegaplan("https://megaplan.ath.ru", "bob_edition@mail.ru", "zaq12wsx#")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("--------------------------------------------------------------")
	tmp := meg.ListEventCategory()
	for _, t := range tmp.Categories {
		fmt.Println(t)
	}
	fmt.Println("--------------------------------------------------------------")
	/*rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5000; i++ {
		err := meg.AddContractor(i, 1000125)
		if err != nil {
			fmt.Printf("err: %d\t EventID: %d\n", err, i)
		}
		if rand.Int()%2 == 0 {
			r := rand.Intn(3)
			fmt.Println("r: ", r)
			time.Sleep(time.Duration(r) * time.Second)
		} else {
			fmt.Println("Not odd")
		}
	}
	*/
	offset := 0
	employee := meg.ListEmployee(0, "", "", "")
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
	createXls(ev, &meg, &employee)
	m := email.NewMessage("Hi", "this is the body")
	m.From = "megaplan_report@ath.ru"
	m.To = []string{"azhokhov@ath.ru"}

	err := m.Attach("report.xlsx")
	if err != nil {
		log.Println(err)
	}

	err = email.Send("srvmail-exim1.ath.ru", nil, m)
	if err != nil {
		log.Println(err)
	}
}

func createXls(events megaplan.Events, meg *megaplan.Megaplan, employee *megaplan.Employees) {
	file := xlsx.NewFile()
	defer file.Save("report.xlsx")
	sales, err := file.AddSheet("Sales")
	if err != nil {
		fmt.Println(err)
	}
	cgm, err := file.AddSheet("CGM")
	if err != nil {
		fmt.Println(err)
	}
	call, err := file.AddSheet("Звонки")
	if err != nil {
		fmt.Println(err)
	}
	meeting, err := file.AddSheet("Встречи")
	if err != nil {
		fmt.Println(err)
	}
	for _, event := range events {
		owner := employee.GetOwnerInfo(event.Owner)
		if owner.Department.Name == "CGM" {
			dt := write(&event, meg, &owner)
			WriteSlise(cgm, dt)
			if event.EventCategory == "Встреча" {
				dt := write(&event, meg, &owner)
				fmt.Println("Встреча: ", dt)
				WriteSlise(meeting, dt)
			}
			if event.EventCategory == "Звонок" {
				dt := write(&event, meg, &owner)
				fmt.Println("Звонок: ", dt)
				WriteSlise(call, dt)
			}
			file.Save("report.xlsx")
		}
		if owner.Department.Name == "Sales" {
			dt := write(&event, meg, &owner)
			WriteSlise(sales, dt)
			if event.EventCategory == "Встреча" {
				dt := write(&event, meg, &owner)
				fmt.Println("Встреча: ", dt)
				WriteSlise(meeting, dt)
			}
			if event.EventCategory == "Звонок" {
				dt := write(&event, meg, &owner)
				fmt.Println("Звонок: ", dt)
				WriteSlise(call, dt)
			}
			file.Save("report.xlsx")
		}
		for _, part := range event.Participants {
			participant := employee.GetParticipantInfo(part.Id)
			if participant.Department.Name == "CGM" {
				if event.EventCategory == "Встреча" {
					dt := write(&event, meg, &participant)
					WriteSlise(meeting, dt)
				}
				if owner.Department.Name == "Звонок" {
					dt := write(&event, meg, &participant)
					WriteSlise(call, dt)
				}
				dt := write(&event, meg, &participant)
				WriteSlise(cgm, dt)
			}
			if participant.Department.Name == "Sales" {
				if event.EventCategory == "Встреча" {
					dt := write(&event, meg, &participant)
					WriteSlise(meeting, dt)
				}
				if owner.Department.Name == "Звонок" {
					dt := write(&event, meg, &participant)
					WriteSlise(call, dt)
				}
				dt := write(&event, meg, &participant)
				WriteSlise(sales, dt)
			}
			file.Save("report.xlsx")
		}
	}
	xlsWrite.All()
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

func write(event *megaplan.Event, meg *megaplan.Megaplan, owner *megaplan.Employee) map[int]AnswerStruct {
	ans := make(map[int]AnswerStruct)
	//a := ans[0].Clients
	//fmt.Println(a)
	var i int
	i = 0
	for _, contactor := range event.Contractors {
		a := ans[i]
		a.Date = event.TimeCreated
		a.Name = event.Name
		a.Description = event.Description
		a.Type = event.EventCategory
		a.Participant = owner.FirstName + " " + owner.LastName

		cl := meg.GetCardContactor(contactor.Id)
		a.Clients = cl.Contractor.Name
		a.LocalGlobal = cl.Contractor.Category183CustomFieldLokalniyGlobalniy
		ans[i] = a
	}
	return ans
}

func WriteSlise(Sheet *xlsx.Sheet, dt map[int]AnswerStruct) {
	for _, data := range dt {
		row := Sheet.AddRow()
		d := &data
		//fmt.Println(d)
		row.WriteStruct(d, -1)
	}
}
