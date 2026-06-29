package helper

import "errors"

// devide

func Pembagian(nilai float32, pembagi float32) (float32, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi error")
	} else {
		return nilai / pembagi, nil
	}
}
