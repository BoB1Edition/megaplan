package megaplan

type Clients struct {
	Clients []Client `json:"clients"`
}

type Client struct {
	Birthday      string `json:"Birthday"`
	ChildContacts []struct {
		ID   int    `json:"Id"`
		Name string `json:"Name"`
	} `json:"ChildContacts"`
	Category183CustomFieldLokalniyGlobalniy string
	CompanyName                             string        `json:"CompanyName"`
	Description                             string        `json:"Description"`
	Email                                   string        `json:"Email"`
	Facebook                                string        `json:"Facebook"`
	GUID                                    string        `json:"GUID"`
	Icq                                     string        `json:"Icq"`
	Id                                      int           `json:"Id"`
	Jabber                                  string        `json:"Jabber"`
	Locations                               []interface{} `json:"Locations"`
	Name                                    string        `json:"Name"`
	Payers                                  []struct {
		ID   int    `json:"Id"`
		Name string `json:"Name"`
	} `json:"Payers"`
	PersonType      string `json:"PersonType"`
	PreferTransport string `json:"PreferTransport"`
	PromisingRate   string `json:"PromisingRate"`
	Responsibles    []struct {
		ID   int    `json:"Id"`
		Name string `json:"Name"`
	} `json:"Responsibles"`
	Site   string `json:"Site"`
	Skype  string `json:"Skype"`
	Status struct {
		ID   int    `json:"Id"`
		Name string `json:"Name"`
	} `json:"Status"`
	TimeCreated string `json:"TimeCreated"`
	TimeUpdated string `json:"TimeUpdated"`
	Twitter     string `json:"Twitter"`
	Type        struct {
		ID   int    `json:"Id"`
		Name string `json:"Name"`
	} `json:"Type"`
}

type ClientType struct {
	Id   int
	Name string
}

type ClientFields struct {
	Fields []struct {
		Name        string `json:"Name"`
		Translation string `json:"Translation"`
	} `json:"Fields"`
}

type ClientContractor struct {
	Contractor struct {
		Category183CustomFieldLokalniyGlobalniy string        `json:"Category183CustomFieldLokalniyGlobalniy"`
		Id                                      int           `json:"Id"`
		Name                                    string        `json:"Name"`
		ResponsibleContractors                  []interface{} `json:"ResponsibleContractors"`
	} `json:"contractor"`
}
