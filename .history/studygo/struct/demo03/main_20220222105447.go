package main

type person struct {
	name, gender string
}

func changeGender(x *person) {
	(*x).gender = "BBBBBBBBBBBBBBBBBBBBBBoy"
}
func main() {
	var p person

	p.name = "holyshit"
	p.gender = "GGGGGGGGGGGGGGGGGirl"

}
