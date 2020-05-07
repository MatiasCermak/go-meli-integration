package main

import (
	"fmt"
	"Calc"
)

func main(){
	Calc.SecretSwitch = true
	var num1 float64
	var num2 float64
	var op, usr string
	fmt.Printf("Hola! Bienvenido a GoCalc, por favor ingresa tu nombre: ")
	fmt.Scanln(&usr)
	fmt.Printf("Por favor ingresa un número: ")
	fmt.Scanln(&num1)
	fmt.Printf("Ingresa otro número: ")
	fmt.Scanln(&num2)
	fmt.Printf("Ahora ingresa la operación deseada (+, -, /, *): ")
	fmt.Scanln(&op)
	log := Calc.Operate(num1, num2, usr, op)
	fmt.Println(log)

}
//Ejemplos



