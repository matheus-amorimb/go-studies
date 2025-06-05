package main

import (
	"errors"
	"fmt"
)

// Error Interface
func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	mCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}
	mSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}

	return (mCustomer + mSpouse), nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

// Custom Error
type useError struct {
	name string
}

func (e useError) Error() string {
	return fmt.Sprintf("$v has a problem with their account", e.name)
}

func sendMessage(msg, userName string) error {
	canSendToUser := true
	if !canSendToUser {
		return useError{name: "matheus"}
	}
	// ...
	return nil
}

// The Errors Package
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("no dividing by 0")
	}
	return x / y, nil
}

func main() {

}
