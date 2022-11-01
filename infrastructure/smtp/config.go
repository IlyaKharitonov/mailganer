package smtp

type Config struct {
	Host  	 string `json:"host"`
	Port  	 int  `json:"port"`
	Login 	 string `json:"login"`
	Password string `json:"password"`
}
