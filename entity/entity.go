package entity

type Account struct {
	EmplId   string `json:"EmplId"`
	EmplName string `json:"EmplName"`
	Alfa     string `json:"Alfa"`
	Gazprom  string `json:"Gazprom"`
	Databank string `json:"Databank"`
	Rosbank  string `json:"Rosbank"`
	Rusbank  string `json:"Rusbank"`
	Sber     string `json:"Sber"`
	Hlyn     string `json:"Hlyn"`
}

type Recoupment struct {
	EmplId   string `json:"EmplId"`
	EmplName string `json:"EmplName"`
	RecType  string `json:"RecType"`
	Alfa     string `json:"Alfa"`
	Gazprom  string `json:"Gazprom"`
	Databank string `json:"Databank"`
	Rosbank  string `json:"Rosbank"`
	Rusbank  string `json:"Rusbank"`
	Sber     string `json:"Sber"`
	Hlyn     string `json:"Hlyn"`
}
