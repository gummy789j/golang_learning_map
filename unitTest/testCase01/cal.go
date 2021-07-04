
// go test -v (這樣無論成功失敗都會log)
// go test (只有報錯才會log)
// go test -v cal_test.go cal.go (單獨測試一個測試檔)
// go test -v sub_test.go cal.go (單獨測試一個測試檔)
// go test -v -run TestAddUpper (單獨測試一個function)
package test

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n ; i++ {
		res += i 
	}
	return res
}

func getSub(n1 int, n2 int) int{
	return n1 - n2
}