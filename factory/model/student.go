package model

type student struct {
	name string
	age int
}

// 由於student字首為小寫，只能在model包能用
// 使用factory Mode 解決


// 就像java得建構函式(必為public)
func NewStudent(n string, a int) (*student) {
	
	return &student{n,a}
	
}

// 就像是get 與 set操作private variable
func (s *student) GetAge() int {
	return s.age
}

func (s *student) GetName() string {
	return s.name
}