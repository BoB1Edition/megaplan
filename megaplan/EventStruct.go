package megaplan

import "fmt"

type Participants map[int]Participant
type Events map[int]Event

type Event struct {
	Id                int
	Description       string
	Name              string
	TimeCreated       string
	StartTime         string
	Duration          int
	IsPersonal        bool
	EventCategory     string
	OriginalStartTime string
	Participants      []Participant
	Contractors       []Contractor
	Reminders         []Reminder
	HasTodo           bool
	HasCommunication  bool
	TodoLisId         int
	Position          int
	Owner             int
	IsFinished        int
	Place             string
	IsFavorite        bool
	TimeUpdated       string
	CanEdit           bool
	IsOverdue         bool
	DateFrom          string
}

type Contractor struct {
	Id   int
	Name string
}

type Reminder struct {
	Transport  string
	TimeBefore int
}

type Participant struct {
	Id   int
	Name string
}

func (e *Events) Filter(f func(Event) bool) Events {
	fmt.Println("Filter: ", len(*e))
	events := Events{}
	for _, data := range *e {
		if f(data) {
			//fmt.Println("key: ", key, "\tdata.Id: ", data.Id)
			events[data.Id] = data
		}
	}
	fmt.Println("Filtered: ", len(events))
	return events
}
