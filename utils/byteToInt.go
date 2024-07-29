package utils

import (
	"errors"
	"fmt"
	"os"
)
/*byteToInt function converts a given byte character to its corresponding integer representation. 
It is designed to handle specific characters and returns an error for any character that it cannot convert.*/
func byteToInt(b byte) (int, error) {
	if b == '.' {
		return 0, nil
	} else if b == '#' {
		return 1, nil
	} else {
		fmt.Println(os.Getwd())
		return 0, errors.New("error in sample file")
	}
}
