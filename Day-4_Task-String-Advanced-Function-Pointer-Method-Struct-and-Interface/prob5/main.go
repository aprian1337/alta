package main

import "fmt"

type Student struct {
	name []string

	score []int
}

func (s Student) Avarage() float64 {
	sum := 0.0
	for _, v := range s.score {
		sum += float64(v)
	}
	return sum / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	score := map[string]int{}
	for i := range s.name {
		if i == 0 {
			min = s.score[i]
			name = s.name[i]
			continue
		}
		score[s.name[i]] = s.score[i]
	}
	for k, v := range score {
		if v < min {
			min = v
			name = k
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	score := map[string]int{}
	for i := range s.name {
		if i == 0 {
			max = s.score[i]
			name = s.name[i]
			continue
		}
		score[s.name[i]] = s.score[i]
	}
	for k, v := range score {
		if v > max {
			max = v
			name = k
		}
	}
	return max, name
}

func main() {

	var a = Student{}

	for i := 0; i < 5; i++ {

		var name string

		fmt.Print("Input " + string(i) + " Student’s Name : ")

		fmt.Scan(&name)

		a.name = append(a.name, name)

		var score int

		fmt.Print("Input " + name + " Score : ")

		fmt.Scan(&score)

		a.score = append(a.score, score)

	}

	fmt.Println("\n\nAvarage Score Students is", a.Avarage())

	scoreMax, nameMax := a.Max()

	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")

	scoreMin, nameMin := a.Min()

	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")

}
