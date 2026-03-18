package main

/*Программист Арсений снова взял непростой заказ: сервис каршеринга попросил его разработать систему тарификации для автомобилей и мотоциклов. У каждого транспортного средства должен быть метод для расчёта времени в пути до пункта назначения, чтобы на основе этих данных вычислять стоимость проката.

Запрограммируйте интерфейс Vehicle, который будет представлять транспортное средство и рассчитывать время в пути с помощью метода CalculateTravelTime(distance float64) float64.

Добавьте две структуры, Car (автомобиль) и Motorcycle (мотоцикл). Обе должны реализовывать интерфейс Vehicle и у обеих должны быть поля для хранения данных о транспортных средствах: Type (string), Speed (float64), FuelType (string).

Создайте функцию EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64. Она должна принимать список транспортных средств и расстояние до пункта назначения, а затем возвращать мапу. Ключи в ней — это типы транспортных средств, а значения — время в пути для обоих (оно рассчитывается через отношение пути к скорости).*/
import (
	"reflect"
)

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
}

type Car struct {
	Type     string
	Speed    float64
	FuelType string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

type Motorcycle struct {
	Type     string
	Speed    float64
	FuelType string
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	mp := make(map[string]float64)
	for _, veh := range vehicles {
		tp := reflect.TypeOf(veh).Name()
		tp = "main." + tp
		time := veh.CalculateTravelTime(distance)
		mp[tp] = time
	}
	return mp
}
