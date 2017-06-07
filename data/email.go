package data

type Email struct {
	FromName    string `json:"FromName"`
	FromAddress string `json:"FromAddress"`
	Subject     string `json:"Subject"`
	ToName      string `json:"ToName"`
	ToAddress   string `json:"ToAddress"`
	Body        string `json:"Body"`
}

func (email *Email) Save() error {
	//TODO
	return nil
}
