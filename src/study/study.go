package study

import (
	"fmt"
	"math"
	"time"
	"errors"
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
const ss string = "constant";
func TestConstants(){
	println("--->ConstantsTest")

	println(ss)

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

func sum1(nums ...int){
	println("nums:",nums)

	sum:=0
	for _,num :=range nums{
		sum+=num
	}
	println("sum:",sum)
}
func TestVariadicFunctions(){
	println("--->TestVariadicFunctions")

	sum1(1,2)
	sum1(1,2,3)

	nums:=[]int{1,2,3,4,5,6,7,8,9}
	sum1(nums...)
}

func intSeq() func() int{
	i:=0
	return func() int{
		i+=1
		return i
	}
}
func TestClosures(){
	println("--->TestClosures")

	nextInt := intSeq()

	println(nextInt())
	println(nextInt())
	println(nextInt())

	nextNInt := intSeq()
	println(nextNInt())
}

func fact(n int)int{
	if n == 0{return 1}
	return n*fact(n-1)
}
func TestRecursion(){
	println("--->TestClosures")
	println(fact(7))
}


func zeroVal(ival int){
	ival = 0
}
func zeroPtr(iptr *int){
	*iptr = 0
}
func TestPointer(){
	println("--->TestPointer")

	i:=1

	println("initial", i)

	zeroVal(i)
	println("zeroVal", i)

	zeroPtr(&i)
	println("zeroPtr", i)

	println("pointer", &i)
}

type person struct {
	name string
	age int
}
func TestStructs(){
	println("--->TestStructs")

	println(person{"bob",20})
	println(person{name:"Alice", age:30})
	println(person{name:"Fred"})
	println(&person{name:"Ann", age:40})

	s:= person{"Sean", 50}
	println(s.name)

	sp:=&s
	println(sp.age)

	sp.age = 51
	println(sp.age)
	println(s.age)
}

type rect struct {
	width,height int
}
func (r *rect) area() int{
	return r.width*r.height
}
func (r rect)  prim() int{
	return 2*r.width + 2*r.height
}

func TestMethods(){
	println("--->TestMethods")

	r:=rect{width:10,height:5}

	println("area", r.area())
	println("prim", r.prim())

	sp := &r
	println("area", sp.area())
	println("prim", sp.prim())
}

type geometry interface {
	area() float64
	prim() float64
}
type square struct {
	width,height float64
}
type circle struct {
	radius float64
}
func (s square)area() float64{
	return s.width*s.height
}
func (s square)prim() float64{
	return 2*s.width+2*s.height
}
func (c circle)area() float64{
	return math.Pi*c.radius*c.radius
}
func (c circle)prim() float64{
	return 2*math.Pi*c.radius
}
func measure(g geometry){
	println(g)
	println(g.area())
	println(g.prim())
}

func TestInterfaces(){
	println("--->TestInterfaces")

	s:=square{width:3,height:4}
	c:=circle{radius:5}

	measure(s)
	measure(c)
}


func f1(arg int)(int, error){
	if arg == 42 {
		return -1, errors.New("con not work with 42!")
	}
	return arg+3, nil
}
type argError struct {
	arg int
	prob string
}
func (e *argError) Error() string{
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func f2(arg int)(int, error){
	if arg == 42 {
		return -1, &argError{arg, "con not work with 42!"}
	}
	return arg+3, nil
}
func TestError(){
	println("--->TestError")

	for _, i:= range []int{7, 42}{
		if r,e := f1(i); e != nil{
			println("f1 failed", e)
		}else{
			println("f1 worked", r)
		}
	}

	for _, i:= range []int{7, 42}{
		if r,e := f2(i); e != nil{
			println("f2 failed", e)
		}else{
			println("f2 worked", r)
		}
	}

	_,e:=f2(42)
	if ae,ok := e.(*argError); ok{
		println(ae.arg)
		println(ae.prob)
	}
}


func f(from string){
	for i:=0;i<3;i++{
		println(from,":",i)
	}
}
func TestGoroutines(){
	println("--->TestGoroutines")

	f("direct")

	go f("goroutine1")
	go func(msg string){
		println(msg)
	}("going")
	go f("goroutine2")

	var input string
	fmt.Scanln(&input)
	println("done")
}


func TestChannels(){
	println("--->TestChannels")

	messages := make(chan string)

	go func(){
		messages<-"ping"
	}()

	msg := <-messages
	println(msg)
}


func TestChannelBuffering(){
	println("--->TestChannelBuffering")

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	println(<-messages, <-messages)
}


func worker(done chan bool){
	println("woking ...")
	time.Sleep(time.Second)
	println("done")

	done<-true
}

func TestChannelSynchronization(){
	println("--->TestChannelSynchronization")

	done := make(chan bool, 1)
	go worker(done)

	<-done
	println("finish")
}
