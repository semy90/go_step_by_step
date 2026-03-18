package main

/*Программист Арсений, как обычно, должен спасти мир, а именно — банковскую систему России. Дирекция крупного национального банка попросила нашего героя переписать их основную платёжную систему. Чтобы, не отвлекаясь, подумать над общей архитектурой программы, Арсений делегирует черновую работу вам.

Напишите программу для управления банковским счётом. Создайте структуру Account с приватными полями balance (float64) и owner (string).

Реализуйте конструктор NewAccount(owner string) *Account, который устанавливает нулевой баланс по умолчанию. Также реализуйте методы:

SetBalance(value float64) error — устанавливает значение баланса
GetBalance() float64 — получает баланс
Deposit(value float64) error — увеличивает баланс (проверьте, чтоvalue — положительное число)
Withdraw(value float64) error — уменьшает баланс (проверьте, чтоvalue — положительное число)
Кроме того, нужно убедиться, что баланс всегда строго положительный.*/
import (
	"fmt"
)

type Account struct {
	balance float64
	owner   string
}

func NewAccount(owner string) *Account {
	return &Account{balance: 0, owner: owner}
}

func (a *Account) SetBalance(value float64) error {
	if value <= 0 {
		return fmt.Errorf("your account is negative or zero")
	} else {
		a.balance = value
		return nil
	}
}
func (a Account) GetBalance() float64 {
	return a.balance
}
func (a *Account) Deposit(value float64) error {
	if value <= 0 {
		return fmt.Errorf("you want to top up a negative value?")
	} else {
		a.balance += value
		return nil
	}
}
func (a *Account) Withdraw(value float64) error {
	if value <= 0 || a.balance < value {
		return fmt.Errorf("you want to get a negative balance?")
	} else {
		a.balance -= value
		return nil
	}
}
