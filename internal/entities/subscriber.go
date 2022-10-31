package entities

type Subscriber struct {
	ID 			int		`json:"id"`
	FirstName  	string  `json:"first_name"`
	LastName 	string  `json:"last_name"`
	Birthday 	string  `json:"birthday"`
	Email  		string  `json:"email"`
	//Subject     string  `json:"subject"`
}

