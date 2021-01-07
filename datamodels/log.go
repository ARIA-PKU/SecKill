package datamodels

type Log struct {
	Name       string `json:"Name" sql:"Name" homework:"Name"`
	ProductID  int64 `json:"ProductID" sql:"ProductID" homework:"ProductID"`
	IP         string  `json:"IP" sql:"IP" homework:"IP"`
	Operation  string `json:"Operation" sql:"Operation" homework:"Operation"`
	Time       string  `json:"Time" sql:"Time" homework:"Time"`
}