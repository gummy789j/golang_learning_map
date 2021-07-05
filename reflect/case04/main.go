package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float64
	Sex   string
}

func (s Student) Print() {
	fmt.Println("-----------start----------")
	fmt.Println(s)
	fmt.Println("-----------end----------")
}

func (s Student) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Student) Set(name string, age int, score float64, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	// for type的(像是Name)，把一般的type轉成 *reflect type (也就是使用reflect操作的進入口)
	typ := reflect.TypeOf(a)

	// for value的(像是Steven)， 把一般的value轉成 reflect value (也就是使用reflect操作的進入口)
	val := reflect.ValueOf(a)

	// 做一個filter 確保進來的是struct
	kd := val.Kind()
	// 他不是真正的struct 而是 reflect版本的struct
	if kd != reflect.Struct {
		return
	}

	// 計算struct fields 數
	num := val.NumField()

	fmt.Printf("struct has %d fields\n", num)

	// 取得此stud裡存放的值
	for i := 0; i < num; i++ {
		// print出 Steven, 24, 95.6, male
		fmt.Printf("Field %d: value is = %v\n", i, val.Field(i))

		nameVal := typ.Field(i).Name            // 顯示 Name, Age, Score, Sex
		tagVal := typ.Field(i).Tag              // 顯示 json:name , json:age
		jsonVal := typ.Field(i).Tag.Get("json") // 顯示 name, age (**這就是在做json marshall時真正做的事)
		fmt.Println(nameVal, tagVal, jsonVal)
	}

	// 取得這個struct有多少個method
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %v method\n", numOfMethod)

	// 取得特定方法並調用他
	val.Method(1).Call(nil) // 他會調用Print()，因為他是用asciicode排序method後，Print()為1(也就是排在第2個 )

	// 假設要調用GetSum() method也就是 Method(0)
	// 他的arguments也要轉乘reflect版本的value
	// 並且要包裹在一個slice裡
	// 呼叫完的返回值也是一個[]reflect.Value的 slice
	// 像是 GetSum只有返回一個值，所以用res[0]去讀
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := val.Method(0).Call(params)
	fmt.Println("res =", res[0]) //(裡面的值的type也還是reflect.Value，要使用要做轉換)
	fmt.Printf("%T\n", res[0])
	fmt.Printf("%T\n", res[0].Interface())       // 轉成了interface type(在runtime時)，在編譯氣會顯示int
	fmt.Printf("%T\n", res[0].Interface().(int)) //故要做斷言
}

func main() {
	stud := Student{
		Name:  "Steven",
		Age:   24,
		Score: 95.6,
		Sex:   "male",
	}

	TestStruct(stud)
}
