package main

/*
Программист Арсений тратит много денег на разработку проекта своей мечты. Ему нужна программа, которая подсчитывает его расходы.

Реализуйте функции:

topUpBalance(amount float64) error, которая увеличивает баланс на amount
chargeFromBalance(amount float64) error, которая уменьшает баланс на amount
TopUpAndGetBalance(amount float64) (float64, error) и ChargeFromAndGetBalance(amount float64) (float64, error), которые вызывают соответствующие функции и возвращают итоговый баланс
Баланс должен быть равен нулю (в начале работы программы) и храниться в глобальной переменной Balance.

Во всех функциях нужно возвращать итоговый баланс и/или ошибку в случае её возникновения:

если ошибка в неправильном значении amount — ошибку amount is incorrect
если некорректен итоговый баланс при выводе из функций TopUpAndGetBalance или ChargeFromAndGetBalance — ошибку balance is incorrect
*/
import (
	"errors"
)

var (
	Balance                      = 0.0
	ChargeFromAndGetBalanceError = errors.New("balance is incorrect")
	TopUpAndGetBalanceError      = errors.New("balance is incorrect")
	AmountError                  = errors.New("amount is incorrect")
)

func topUpBalance(amount float64) error {
	if amount <= 0 {
		return AmountError
	} else {
		Balance += amount
		return nil
	}
}
func chargeFromBalance(amount float64) error {
	if amount <= 0 {
		return AmountError
	} else {
		Balance -= amount
		return nil
	}
}
func TopUpAndGetBalance(amount float64) (float64, error) {
	err := topUpBalance(amount)
	if err == nil {
		if Balance <= 0 {
			return 0, TopUpAndGetBalanceError
		} else {
			return Balance, nil
		}
	} else {
		return 0, err
	}
}
func ChargeFromAndGetBalance(amount float64) (float64, error) {
	err := chargeFromBalance(amount)
	if err == nil {
		if Balance <= 0 {
			return 0, TopUpAndGetBalanceError
		} else {
			return Balance, nil
		}
	} else {
		return 0, err
	}
}
