package main

import "fmt"

func main() {
	const BaseRate = 5.50
	const TaxRate = 0.12
	const DistanceRate = 2.0
	const FragileFee = 0.2

	var senderName string
	var cargoWeight float64
	var distanceKm float64
	var fragilePackages int

	fmt.Print("Введите Имя:")
	fmt.Scan(&senderName)

	fmt.Print("Введите вес в килограммах:")
	fmt.Scan(&cargoWeight)

	fmt.Print("Введите расстояние в километрах:")
	fmt.Scan(&distanceKm)

	fmt.Print("Укажите количество хрупких упаковок:")
	fmt.Scan(&fragilePackages)

	var baseCost float64

	baseCost = cargoWeight * BaseRate

	fragileMultiplier := 1 + FragileFee*float64(fragilePackages)

	baseCost = baseCost*fragileMultiplier + distanceKm*DistanceRate

	finalCost := baseCost * (1 + TaxRate)

	fmt.Println("\n--- Отчет о доставке ---")
	fmt.Printf("Отправитель: %s\n", senderName)
	fmt.Printf("Итоговая стоимость: %.2f\n", finalCost)
}
