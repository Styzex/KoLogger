package main

import (
	"fmt"
	"strings"

	"github.com/MarinX/keylogger"
	"github.com/sirupsen/logrus"
)

func main() {
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
	}
	defer k.Close()

	events := k.Read()

	for e := range events {
		if e.KeyPress() {
			logrus.Println("Registered key press event: ", e.KeyString())
		}
	}
}
