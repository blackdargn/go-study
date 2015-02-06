package study

import (
	"fmt"
	"math"
	"time"
)

func println(a ...interface{})(n int, err error){
	return fmt.Println(a);
}

func HelloWorld(){
	println("Hello Go World!")
}

func TestValue(){
	println("--->ValueTest")

	println("go" + "lang")
	println("1+1 =", 1+1)
	println("7.0/3.0 =", 7.0/3.0)
	println(true && false)
	println(true || false)
	println(!true)
}

func TestVariables(){
	println("--->VariablesTest")

	var a string = "initial"
	println(a)

	var b,c int = 1,2
	println(b,c)

	var d = true
	println(d)

	var e int
	println(e)

	f:="short"
	println(f)
}
const s string = "constant";
func TestConstants(){
	println("--->ConstantsTest")

	println(s)

	const n = 50000000
	const d = 3e20/n
	println(d)

	println(int64(d))

	println(math.Sin(n))
}

func TestFor(){
	println("--->ForTest")

	i:=1
	for i<=3 {
		println(i)
		i=i+1
	}

	for j:=7; j<=9; j++ {
		println(j)
	}

	for {
		println("loop");
		break
	}
}

func TestIf(){
	println("--->IfTest")

	if 7%2 == 0 {
		println("7 is even")
	}else{
		println("7 is odd")
	}

	if 8%4 == 0 {
		println("8 is divisible by 4")
	}

	if num:=9;num<0 {
		println(num, " is negative")
	}else if num<10 {
		println(num, " has 1 digit")
	}else{
		println(num, " has mutiple digits")
	}
}

func TestSwitch(){
	println("--->SwitchTest")

	i:=2
	println("write ",i," as")
	switch i{
	case 1:
		println("one")
	case 2:
		println("two")
	case 3:
		println("three")
	}

	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		println("it is weekend")
	default:
		println("it is weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("it's before noon")
	default:
		println("it's after noon")
	}
}

func TestArrays(){
	println("--->ArraysTest")

	var a[5] int
	println("emp:", a)

	a[4] = 100
	println("set a ", a)
	println("get a[4]", a[4])

	println("len(a)", len(a))

	b:=[5]int{1,2,3,4,5}
	println("dcl:",b)

	var twoD[2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j] = i+j
		}
	}
	println("2d:", twoD)
}

func TestSlices(){
	println("--->SlicesTest")

	s:=make([]string, 3)
	println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	println("set s", s)
	println("get s", s[2])

	println("len(s)", len(s))

	s=append(s,"d")
	s=append(s,"e","f")
	println("append", s)
	println("len(s)", len(s))

	c:=make([]string, len(s))
	copy(c, s)
	println("copy", c)

	l:=s[2:5]
	println("sl1",l)

	l=s[:5]
	println("sl2",l)

	l=s[2:]
	println("sl3",l)

	t:=[]string{"g","h","i"}
	println("del:", t)

	twoD :=make([][]int, 3)
	for i:=0;i<3;i++{
		iLen := i+1
		twoD[i] = make([]int,iLen)
		for j :=0; j<iLen; j++{
			twoD[i][j] = i+j
		}
	}
	println("2d sl", twoD)
}

func TestMaps(){
	println("--->MapsTest")

	m:=make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	println("map", m)

	v1:=m["k1"]
	println("get v1",v1)
	println("map len", len(m))

	delete(m,"k2")
	println("map", m)

	_,prs :=m["k2"]
	println("prs",prs)

	n:=map[string]int{"foo":1,"bar":2}
	println("dcl", n)
}

func TestRange(){
	println("--->RangeTest")

	nums:=[]int{2,3,4}
	sum:=0
	for _,num:=range nums{
		sum+=num
	}
	println("sums",sum)

	for i,num :=range nums{
		if num == 3{
			println("index", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k,v := range kvs{
		println("%s->%s",k,v)
	}

	for i,c:=range "go"{
		println(i, c)
	}
}

func plus(a int, b int) int{
	return a+b;
}
func TestFunctions(){
	println("--->FunctionsTest")

	res:=plus(1,2)
	println("1+2=", res)
}

func vals()(int ,int){
	return 3,7
}

func TestMultipleReturn(){
	println("--->MultipleReturnTest")

	a,b := vals()
	println("a,b = ",a,b)

	_,c := vals()
	println(c)
}

func sum(nums ...int){
	println("nums:",nums)

	sum:=0
	for _,num :=range nums{
		sum+=num
	}
	println("sum:",sum)
}
func TestVariadicFunctions(){
	sum(1,2)
	sum(1,2,3)

	nums:=[]int{1,2,3,4,5,6,7,8,9}
	sum(nums...)
}


func TestClosures(){

}

