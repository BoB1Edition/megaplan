package megaplan

type Employees struct {
	Employees []Employee `json:"employees"`
}

type Position struct {
	Id   int
	Name string
}

type Department struct {
	Id   int
	Name string
}

type Employee struct {
	Id             int
	Name           string
	LastName       string
	FirstName      string
	Description    string
	MiddleName     string
	Gender         string
	Position       Position
	Department     Department
	Birthday       string
	HideMyBirthday bool
	Age            int
	Phones         interface{}
	Email          string
	Icq            string
	Skype          string
	Jabber         string
	//Address	object (Id, City, Street, House)	Адрес
	Behaviour    string
	Inn          string
	PassportData string
	AboutMe      string
	User         User
	//ChiefsWithoutMe	array<object> (Id, Name)
	//SubordinatesWithoutMe	array<object> (Id, Name)
	//Coordinators	array<object> (Id, Name)
	//Status	object (Id, Name)
	AppearanceDay       string
	FireDay             string
	TimeCreated         string
	TimeUpdated         string
	Avatar              string
	Photo               string
	Login               string
	LastOnline          string
	IsOnline            bool
	UnreadCommentsCount int
}

type User struct {
	Id   int
	Name string
}

func (es *Employees) GetOwnerInfo(id int) Employee {
	for _, e := range es.Employees {
		if e.User.Id == id {
			return e
		}
	}
	return Employee{}
}

func (es *Employees) GetParticipantInfo(id int) Employee {
	for _, e := range es.Employees {
		if e.Id == id {
			return e
		}
	}
	return Employee{}
}
