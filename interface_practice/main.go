package main

import (
	"fmt"
	"sort"
	"math/rand"
)

type Student struct {
	Name string
	Age int
	Score float64
}

type StuSlice []Student

func (s StuSlice) Len() int {
	return len(s)
}
// 判斷 i and j index 誰要擺在前面 如果return true代表 i 擺在 j 前面
// 若要由大排到小 i 作為index其value就要比較大 這樣才能回true
func (s StuSlice) Less(i,j int) bool{
	return s[i].Score > s[j].Score
}

func (s StuSlice) Swap(i,j int) {
	s[i] , s[j] = s[j] , s[i]
}

func main() {
	var stuArr StuSlice
	for i := 0; i < 10; i++ {
		stu := Student{
			Name : fmt.Sprintf("No. %d ~", rand.Intn(150)),
			Age : rand.Intn(30),
			Score : rand.Float64(),
		}
		stuArr = append(stuArr,stu)
	}

	for _,v := range stuArr {
		fmt.Println(v)
	}

	fmt.Println("------------------")

	sort.Sort(stuArr)

	for _,v := range stuArr {
		fmt.Println(v)
	}

}