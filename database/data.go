package database

type People struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
	City    string `json:"city"`
}

var Peoples = []People{
	{ID: "1", Name: "Sonia", Contact: "sonia@gmail.com", City: "Bangalore"},
	{ID: "2", Name: "Shubhi", Contact: "shubhi@gmail.com", City: "Raipur"},
	{ID: "3", Name: "Lalit", Contact: "lalit@gmail.com", City: "Delhi"},
	{ID: "4", Name: "Udit", Contact: "udit@gmail.com", City: "Singapore"},
}
