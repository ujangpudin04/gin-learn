package people

type Person struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Alamat string `json:"alamat"`
}
type PersonBook struct {
	ID       int `json:"id"`
	IdPerson int `json:"id_person"`
	IdBook   int `json:"id_book"`
}
