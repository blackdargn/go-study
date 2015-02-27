package study
import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
	"time"
	)
	
var x,y,z int = 1, 2, 3
const wrold = "世界"
const(
	Big = 1<<100
	Small = Big >> 99
)
var(
	MaxInt uint64 = 1<<64 -1
	zz complex128 = cmplx.Sqrt( -5 + 12i)
)

type Vertex struct{
	x float64
	y float64
}

type Location struct{
	lat float32
	lon float32
}

var locmap map[string]Location

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // has type *Vertex
	r = Vertex{x: 1}  // Y:0 is implicit
	s = Vertex{}      // X:0 and Y:0
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func TestStudy2(){
	c,python,java := true, false , "no"
	fmt.Println(x,y,z,c,python,java,wrold)
	fmt.Println(swap("hello","world"))
	fmt.Println(spilt(30))
	fmt.Println(needInt(10))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(sum(10))
	fmt.Println(sum2(100))
	fmt.Println(sqrt(99))
	fmt.Println(power(10,12, 100))
	
	const f = "%T(%v)\n"
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, zz, zz)
	
	fmt.Println(Vertex{1, 2})
	
	v := Vertex{1, 2}
	v.x = 4
	fmt.Println(v.x)
	
	fmt.Println(p, q, r, s)
	
	vv := new(Vertex)
	fmt.Println(vv)
	vv.x, vv.y = 11, 9
	fmt.Println(vv)
	
	locmap = make(map[string]Location)
	locmap["pos1"] = Location{113.45323, 36.34222}
	fmt.Println(locmap["pos1"])
	
	testMap()
	
	testSlice()
	
	testSlice2()
	
	testNil()
	
	testFun()
	
	testFun2()
	
	testRange()
	
	testRange2()
	
	testSwitch()
	
	testSwitch2()
	
	testStructFun()
	
	testDeFun()
	
	testDeFun2()
	
	testInterface()
	
	testError()
	
	fmt.Println(sqrt2(-2))
	
	testGoRuntine()
	
	testChannel()
	
	testChannle2()
	
	testChannel3()
	
	testSelelct()
	
	testSelect2()
}

func testSelect2(){
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	
	for {
		select {
			case <- tick :
				fmt.Println("tick!")
			case <- boom:
				fmt.Println("BOOM!")
                return
			default:
				fmt.Println("    .")
                time.Sleep(5e7)
		}
	}
}

func fibonacci2(c, quit chan int){
	x, y := 0, 1
	for {
		select {
			case c <- x:
				x,y = y,x+y
			case <- quit:
				fmt.Println("quit")
                return
		}
	}
}

func testSelelct(){
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func fibonacci(n int, c chan int){
	x, y := 0 ,1
	for i := 0; i < n; i++ {
		c <- x
		x,y = y,x+y
	}
	close(c)
}

func testChannel3(){
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func testChannle2(){
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func sumch(a []int, c chan int){
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func testChannel(){
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sumch(a[:len(a)/2], c)
	go sumch(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y , x+y)
}

func testGoRuntine(){
	go say("world")
	say("hello")
}

func say(s string){
	for i:=0; i <5 ; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func testError() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type Abser interface {
	Abs() float64
}

func testInterface(){
	var a ,b Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	
	a = f  // a MyFloat implements Abser
	b = &v // a *Vertex implements Abser
//	a = v  // a Vertex, does NOT implement Abser

	fmt.Println(a.Abs(),b.Abs())
}

func testDeFun2(){
	v := &Vertex{3, 4}
	v.Scale2(5)
	fmt.Println(v, v.Abs())
}

func (v Vertex) Scale2(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func testDeFun(){
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


func testStructFun(){
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func testSwitch2(){
	t := time.Now()
	switch {
	case t.Hour() < 12:
	    fmt.Println("Good morning!")
	case t.Hour() < 17:
	    fmt.Println("Good afternoon.")
	default:
	    fmt.Println("Good evening.")
	}
}

func testSwitch(){
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:		
			fmt.Printf("%s.", os)
	}
}

func testRange2(){
	pow := make([]int, 10)
	for i:= range pow {
		pow[i] = 1<<uint(i)
	}
	for _,value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func testRange(){
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func testFun2() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func testFun(){
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))
}

func testNil(){
	var z []int
		fmt.Println(z, len(z), cap(z))
		if z == nil {
			fmt.Println("nil!")
		}
}

func testSlice(){
	p := []int {1,2,3,4,5,6,7,8}
	fmt.Println("p ==", p)
	
	fmt.Println("p[1:4] ==", p[1:4])
	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])
	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:])
}

func testSlice2(){
	a := make([]int, 5)
	printSlice("a", a)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	b := make([]int, 0, 5)
	b = append(b,6,7,8,9,10)
	printSlice("b", b)
	c := a[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func testMap(){
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v1, ok := m["Answer"]
	fmt.Println("The value:", v1, "Present?", ok)
}

func power(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim{
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

type ErrorngSqrt float64

func (h ErrorngSqrt) Error() string{
	return fmt.Sprint("cannot Sqrt negative number: ", float64(h))
}

func sqrt2(x float64) (float64, error){
	if x <0 {
		return 0, ErrorngSqrt(x)
	}else{
		return math.Sqrt(x), nil
	}
}

func sum(x int) int{
	sum := 0
	for i := 0; i < x; i++{
		sum += i
	}
	return sum
}

func sum2(x int) int{
	sum := 1
	for sum < x {
		sum += sum
	}
	return sum
}

func needInt(x int) int{
	return x*10 +1
}

func needFloat(x float64) float64{
	return x*0.1
}

func swap(x,y string)(string,string){
	return y,x
}

func spilt(sum int)(x, y int){
	x = sum* 4/9
	y = sum - x
	return
}
