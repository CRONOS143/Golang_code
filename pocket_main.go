package main

import "fmt"

var desktop string = "Card of worker\n"

const Age = 18

func main() {
	var info = " Name: Mikle\n" +
		"  Age: 18\n" +
		"  Job: IT\n" +
		"Skill: Student\n" +
		"  EXP: low\n" +
		"Mob.Num: +7(978)788-53-06\n"
	fmt.Println(desktop, info)

	fmt.Println("Varification: Yes\t"+"Passport: RU\t"+"Age:", Age)
	var a int = 59
	var b int = 10
	var c int = 2
	var d float32 = ((float32(a) + float32(b) + float32(c)) / 3)
	fmt.Println("Average: ", d)

}
