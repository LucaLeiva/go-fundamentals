package pointerserrors

import "errors"

// var hace que sea una variable global al paquete
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// en Go si un simbolo esta en minusculas, entonces es privado para los otros paquetes
type Wallet struct {
	balance Bitcoin
}

// en Go por defecto el paso de argumentos es por COPIA, NO por referencia, tanto los
// parametros de la función como el receiver
// poniendo el & adelante de una variable nos imprime la dirección de memoria
func (w *Wallet) Deposit(ammount Bitcoin) {
	w.balance += ammount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
