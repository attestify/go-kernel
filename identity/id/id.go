package id

import "strconv"

type Id struct {
	id int64
}

func New(id int64) Id {
	return Id{id: id}
}

func (id *Id) Value() string {
	return strconv.Itoa(int(id.id))
}

func (id *Id) AsInteger() int64 {
	return id.id
}
