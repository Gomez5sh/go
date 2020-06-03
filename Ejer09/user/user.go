package user

import "time"

type Usuario struct {
	Id		int
	Name	string
	Date	time.Time
	Status	bool
}

func (this *Usuario) AltaUsuario(id int, name string, date time.Time, status bool) {
	this.Id = id
	this.Name = name
	this.Date = date
	this.Status = status
}