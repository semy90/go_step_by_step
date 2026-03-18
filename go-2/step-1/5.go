package main

/*Напишите функцию

func Contains(r io.Reader, seq []byte) (bool, error)
которая должна найти в данных первое вхождение байт seq, которые доступны через Reader r.

Если последовательность найдена, программа возвращает true, nil, иначе false, nil. Если возникает ошибка, функция должна возвращать false и ошибку.*/
import (
	"io"
)

func Equal(tmp, seq *[]byte) bool {
	for i := 0; i < len(*tmp); i++ {
		if (*tmp)[i] != (*seq)[i] {
			return false
		}
	}
	return true
}

func Contains(r io.Reader, seq []byte) (bool, error) {

	buffer := make([]byte, 1)
	tmp := make([]byte, len(seq))
	_, err := r.Read(tmp)
	if err != nil && err != io.EOF {
		return false, err
	}

	if Equal(&tmp, &seq) {
		return true, nil
	}

	for {
		_, err := r.Read(buffer)
		if err != nil && err != io.EOF {
			return false, err
		}
		if err == io.EOF {
			break
		}
		copy(tmp, tmp[1:])
		tmp[len(tmp)-1] = buffer[0]
		if Equal(&tmp, &seq) {
			return true, nil
		}

	}
	return false, nil

}
