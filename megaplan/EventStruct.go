package megaplan

type Participants map[int]Participant

type Event struct {
	Id               int
	Description      string
	Name             string
	TimeCreated      string
	StartTime        string
	Duration         int
	IsPersonal       bool
	EventCategory    string
	Participants     []Participant
	Contractors      []Contractor
	Reminders        []Reminder
	HasTodo          bool
	HasCommunication bool
	TodoLisId        int
	Position         int
	Owner            int
	IsFinished       int
	Place            string
	IsFavorite       bool
	TimeUpdated      string
	CanEdit          bool
	IsOverdue        bool
}

type Contractor struct {
	Id   int
	Name string
}

type Events struct {
	Events []Event
}

type Reminder struct {
	Transport  string
	TimeBefore int
}

type Participant struct {
	Id   int
	Name string
}
