package entities

type Letter struct {
	Message 		string  `json:"message"`
	RecipientsID 	[]uint  `json:"recipients_id"`
	DispatchTime 	uint  	`json:"dispatch_time"`
	TemplateID 		uint 	`json:"template_id"`
}

