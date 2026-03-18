package main

/*Реализуйте функцию SerializeIntSlice(nums []int) ([]byte, error), которая принимает на вход слайс целых чисел и возвращает его сериализацию в формате JSON.
Функция должна возвращать срез байт с JSON и ошибку, если сериализация не удалась*/
import (
	"encoding/json"
)

func SerializeIntSlice(nums []int) ([]byte, error) {
	jsonBytes, err := json.Marshal(nums)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}
