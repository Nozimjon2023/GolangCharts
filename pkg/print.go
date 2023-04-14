package pkg

import (
	"fmt"
	"time"
)
func A () {
	fmt.Println(time.Now().Second())
}
func B () {
	fmt.Println("Минут")
}
func C () {
	fmt.Println(time.Now().String() + " пол час")
}
func D () {
	fmt.Println(time.Now().String() + " час")
}
