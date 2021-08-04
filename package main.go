package main

import (
	"fmt"
	"time"
)

const way int64 = 1600

func data() {
	tTime := time.Date(2021, time.August, 2, 18, 24, 5, 8, time.Local)
	fmt.Printf("Дата:  %s \n", tTime.Format("01/02/2006"))
	fmt.Printf("Время: %s \n", tTime.Format("15:04"))
}
func main() {
	data()
	WayCalculated(45)

	PrintMessage(80)
}
func WayCalculated(Speed int64) {
	var tTime = int64(way) / int64(Speed)
	fmt.Println("Время отбытия через: ", tTime, "минут\n")
}
func PrintMessage(Speed int) float64 {
	fmt.Println("Время: ", ("19:00"))
	fmt.Printf("Ваша скорость: %d км/ч\n", Speed)
	fmt.Println("Ваш путь:", way, "км")
	fmt.Println("Время в пути ", float64(way)/float64(Speed), "часов")
	return float64(way) / float64(Speed)
}
