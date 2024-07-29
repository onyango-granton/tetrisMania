package utils

import (
	"errors"
	"fmt"
)

/*The stringToIntSlice function converts a string of exactly four characters into a slice of integers. 
Each character in the string is converted to its corresponding integer value using the byteToInt function. 
If the input string is not exactly four characters long,
 or if any character cannot be converted to an integer, an error is returned.*/
func stringToIntSlice(s string) ([]int, error) {
	// file,_ := os.ReadFile("sample.txt")
	// fmt.Println(file)
	res := []int{}
	if len(s) != 4 {
		fmt.Println(s)
		return nil, errors.New("invalid length entry in file")
	}
	for _, b := range s {
		num, err := byteToInt(byte(b))
		if err != nil {
			return nil, err
		}
		res = append(res, num)
	}
	return res, nil
}
