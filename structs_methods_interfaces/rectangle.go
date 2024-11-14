package structsmethodsinterfaces

type Rectangle struct {
	Width  float64
	Height float64
}

// en go no existe la sobrecarga de metodos
// lo que viene antes del nombre de la funcion y los parametros es un "recibidor", donde se "vincula" el struct con el método
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}
