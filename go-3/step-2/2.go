package main

/*Реализуйте функцию DeserializeStringMap(data string) (map[string]string, error),
которая принимает строку JSON и возвращает map[string]string.
Если входная строка не является корректным JSON-объектом, функция должна вернуть ошибку.*/
import (
	"encoding/json"
)

func DeserializeStringMap(data string) (map[string]string, error) {
	mp := map[string]string{}
	err := json.Unmarshal([]byte(data), &mp)
	if err != nil {
		return nil, err
	}
	return mp, nil
}
