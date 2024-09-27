package main

import "testing"

/* Функция GetUTFLength(input []byte) (int, error) возвращает длину строки UTF8
и ошибку ErrInvalidUTF8 (в случае возникновения).
Напишите тест, который бы проверял возвращаемые функцией значения.*/

func TestGetUTFLength(t *testing.T) {
	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}
	n, err := GetUTFLength(valid)
	if n != 9 || err != nil {
		t.Errorf("Error%s", err)
	}
	n, err = GetUTFLength(invalid)
	if n != 0 || err == nil {
		t.Errorf("Error%s", err)
	}
}

/*Напишите тест для функции DeleteVowels(s string) string,
которая должна удалять все гласные из строки английского языка (y не считается гласной).
Используйте table driven testing.*/

type Test struct {
	old string
	new string
}

var tests = []Test{
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", "BCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz"},
	{"zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA", "zyxwvtsrqpnmlkjhgfdcbZYXWVTSRQPNMLKJHGFDCB"},
	{"Hello, world!", "Hll, wrld!"},
	{"1234567890ao", "1234567890"},
}

func TestDeleteVowels(t *testing.T) {
	for i, test := range tests {
		line := DeleteVowels(test.old)
		if line != test.new {
			t.Errorf("#%d: DeleteVowels(%s)=%s; want %s", i, test.old, line, test.new)
		}
	}
}
