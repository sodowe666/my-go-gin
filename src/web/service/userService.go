package service

type test struct {
	a int
	b string
}

var Obj test

//var once sync.Once

func GetInstance() test {
	Obj.b = "3"
	Obj.a = 1
	return Obj
}
