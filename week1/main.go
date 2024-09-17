// Hands-on Week 1
// Muhammad Grandiv Lava Putra
// 22/493242/TK/54023
// Teknologi Informasi

package main

import (
	"fmt"
)

func main() {
	var celsius float64
	var choice int

	// Input suhu dalam Celcius
	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scanln(&celsius)

	// Pilihan satuan konversi
	fmt.Println("Pilih konversi:")
	fmt.Println("1. Fahrenheit")
	fmt.Println("2. Kelvin")
	fmt.Println("3. Reamur")
	fmt.Print("Masukkan pilihan (1/2/3): ")
	fmt.Scanln(&choice)

	// Switch case untuk melakukan konversi
	switch choice {
	case 1:
		// Konversi ke Fahrenheit
		fahrenheit := (celsius * 9 / 5) + 32
		fmt.Printf("Suhu dalam Fahrenheit: %.2f°F\n", fahrenheit)
	case 2:
		// Konversi ke Kelvin
		kelvin := celsius + 273.15
		fmt.Printf("Suhu dalam Kelvin: %.2fK\n", kelvin)
	case 3:
		// Konversi ke Reamur
		reamur := celsius * 4 / 5
		fmt.Printf("Suhu dalam Reamur: %.2f°Re\n", reamur)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
