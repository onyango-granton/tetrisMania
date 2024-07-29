package utils

import (
	"errors"
	"os"
	"strings"
)

/*The ErrorHandling function is designed to validate the command-line arguments provided when running a Go program. Specifically, it checks 
if the correct number of arguments is provided and if the filename and its extension are correctly formatted. 
This function returns an error if any of these checks fail, providing informative error messages to guide the user on how to correctly use the program.*/
func ErrorHandling() error {
	if len(os.Args) != 2 {
		errorMsg := "invalid number of arguments\nUsage:\ngo run . <filename.txt>"
		return errors.New(errorMsg)
	}
	filenameList := strings.Split(os.Args[1], ".")
	if len(filenameList) != 2 {
		errorMsg := "invalid filename and/or extension\nUsage:\ngo run . <filename.txt>"
		return errors.New(errorMsg)
	}
	return nil
}
