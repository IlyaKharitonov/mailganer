package smtp

type Config struct {
	Host  	 string `json:"host"`
	Port  	 string `json:"port"`
	Login 	 string `json:"login"`
	Password string `json:"password"`
}
