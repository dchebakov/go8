package auth

type Info interface {
	Username() string
	ID() int64
}

type DefaultUser struct {
	name string
	id   int64
}

func (d *DefaultUser) Username() string {
	return d.name
}

func (d *DefaultUser) ID() int64 {
	return d.id
}

func NewDefaultUser(name string, id int64) Info {
	return &DefaultUser{
		name: name,
		id:   id,
	}
}
