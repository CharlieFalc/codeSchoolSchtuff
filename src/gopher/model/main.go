package model

type gopher struct {
	name string
	age int
	isAdult bool
}

func (g gopher) Jump() string{
	if g.age>65{
		return g.name + " can still jump"
	} else {
		return g.name + " can jump HIGH"
	}
}

type horse struct {
	name string
	weight int
}

func (h horse) Jump() string {
	if h.weight > 2500 {
		return "I'm too heavy, can't jump..."
	} 
	return "I will jump, neigh!!!"
}

// convention for go is to add a 'er' to the interface name of the method that it is implementing
type jumper interface {
	Jump() string
}

func GetList() []jumper{
	phil := &gopher{name: "Phil", age: 30}
	noodles := &gopher{name: "Noddles", age: 90}
	barbaro := &horse {name: "Barbaro", weight: 2000}

	list := []jumper{phil, noodles, barbaro}
	return list
}