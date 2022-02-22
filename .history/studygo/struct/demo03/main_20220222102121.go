package main

type person struct {
	name, gender string
	age          int
	number       int
	guns         []string
	holy         map[int]string
}

func main() {
	var p person

	p.name = "holyshit"
	p.age = 19
	p.guns = []string{"AK-47", ""}
	p.holy = make(map[int]string, 10)
	p.holy[1] = "shit"
	p.holy[2] = "fuck"
	p.holy[3] = "smoke"

}
