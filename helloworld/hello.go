package main

const (
	ingles             = "Ingles"
	frances            = "Frances"
	prefijoHolaIngles  = "Hello, "
	prefijoHolaEspañol = "Hola, "
	prefijoHolaFrances = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "Mundo"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefijo string) {
	switch language {
	case ingles:
		prefijo = prefijoHolaIngles
	case frances:
		prefijo = prefijoHolaFrances
	default:
		prefijo = prefijoHolaEspañol
	}

	return
}
