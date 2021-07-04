package main

//表示，hello.go文件所在的包是main
//每一個go文件都要屬於一個包
import (
	"fmt"
	"go_code/project01/include"
	"io"
	"os"
)

var Name string
var Age string
var City string
var School string
var Phone string
var ptr *int
var str = "hello world"

var (
	Func1 = func(n1 int, n2 int) int {
		return n1 + n2
	} 
)
//表示，引入一個包，包名為fmt，引入該包後，可使用裡面的函數



//init function in charge of initialize, it will be called before main function
func init() {
	fmt.Println("hi i am init ")
}

func main() {
	
	res1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10,35)

	func (n1 int, n2 int) {
		fmt.Println(n1+n2)
	}(10,20)

	fmt.Println(res1)
	
	res2 := Func1(1,2)

	fmt.Println(res2)

	fmt.Println("hi i am main")

	for index, val := range str {
		fmt.Printf("第%d %c\n",index,val)
	}
	
	// fmt.Println("Please enter your name: ")
	// fmt.Scanln(&Name)
	// fmt.Println("Please enter your age: ")
	// fmt.Scanln(&Age)
	// fmt.Println("Please enter your city: ")
	// fmt.Scanln(&City)
	// fmt.Println("Please enter your school: ")
	// fmt.Scanln(&School)
	// fmt.Println("Please enter your phone: ")
	// fmt.Scanln(&Phone)
	include.FindMyBirthday()
	fmt.Printf("%v/%d/%d\n",86,6,26)

	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)

	io.WriteString(os.Stdout, s)
	
	fmt.Println(myfunc1(getSum,5,5))
	fmt.Println(myfunc2(getSum,10,10))

	
	// Ignoring error for simplicity.
	// fmt.Printf("\n%s %s %s %s %s",Name, Age, City, School, Phone)
	// sum, sub := cal(5,6)
	// fmt.Printf("sum is %d, sub is %d\n", sum, sub)
	n1, sum2 := cal2(5,5,5,5,5,5,5)
	fmt.Println("Sum is",sum2, "n1 is", n1)
	n1 = 5
	n2 := 8
	swap(&n1, &n2)
	fmt.Printf("n1 = %d, n2 = %d\n",n1, n2)

	printPyramid(10)

	printMultiply(15)
}

type myFuncType func(int,int) int

func getSum(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func myfunc1(funcvar func(int,int)int, num1 int, num2 int) int {
	return funcvar(num1,num2)
}

func myfunc2(funcvar myFuncType, num1 int, num2 int) int {
	return funcvar(num1,num2)
}

func cal(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

func cal2(n int ,args... int) (n1 int, sum2 int) {
	n1 = n
	//var temp int
	for i := 0; i < len(args); i++ {
		sum2 += args[i]
	}
	return
}

func swap(n1 *int, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

var val int = test()

func test() int {
	fmt.Println("hello i am test")
	return 2
}

func printStar(n int) {
	for i := 0; i < 2*n+1; i++ {
		fmt.Print("*")
	}
}

func printSpace(n int) {
	for i := 0; i < n-1; i++ {
		fmt.Print(" ")
	} 
}

func printPyramid(h int) {
	for i := 0; i < h; i++ {
		printSpace(h-i)
		printStar(i)
		fmt.Println()
	} 
}

func printMultiply(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dX%d ",i,j)
		}
		fmt.Println()
	}
}