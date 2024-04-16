package main

import "fmt"

// Определение лимитов в зависимости от количества проживающих
var maxRooms = map[int]int{
	1: 1,
	2: 2,
	3: 3,
	4: 4,
	5: 5,
	6: 6,
	7: 7,
	8: 8,
}

var maxRent = map[int]int{
	1: 683,
	2: 790,
	3: 951,
	4: 1080,
	5: 1209,
	6: 1338,
	7: 1461,
	8: 1595,
}

const maxAdditional = 200 // Максимальные дополнительные расходы
const maxTotalDeposit = 3 // Максимальная сумма кауции в месяц

var personen int
var zimmer int
var miete int
var betriebskosten int
var heizUNDwarmwasserkosten int
var wasserUNDentwasserungskosten int
var kaution int

func main() {
	fmt.Println("Personen im Haushalt:")
	fmt.Scan(&personen)

	fmt.Println("Wie viele zimmer?:")
	fmt.Scan(&zimmer)

	fmt.Println("Wie viele miete?:")
	fmt.Scan(&miete)

	fmt.Println("Wie viele Betriebskosten?:")
	fmt.Scan(&betriebskosten)

	fmt.Println("Wie viele Heiz und Warmwasserkosten?:")
	fmt.Scan(&heizUNDwarmwasserkosten)

	fmt.Println("Wie viele Wasser und Entwasserungskosten?:")
	fmt.Scan(&wasserUNDentwasserungskosten)

	fmt.Printf("Wie viele kaution (3 monat = %v ):", maxTotalDeposit*miete)
	fmt.Scan(&kaution)

	// Проверка соответствия
	totalRent := miete + betriebskosten + heizUNDwarmwasserkosten + wasserUNDentwasserungskosten
	if personen >= 1 && personen <= 4 && miete <= maxRent[personen] && zimmer <= maxRooms[personen] && totalRent <= maxRent[personen]+maxAdditional && kaution <= maxTotalDeposit*miete {
		fmt.Println("Квартира соответствует лимитам.")
	} else {
		fmt.Println("Квартира не соответствует лимитам.")
		if totalRent > maxRent[personen] {
			fmt.Println("Превышение максимальной арендной платы на", totalRent-maxRent[personen])
		}
		if zimmer > maxRooms[personen] {
			fmt.Println("Превышение максимального количества комнат на", zimmer-maxRooms[personen])
		}
	}
}
