package study

import (
	"fmt"
	"math"
	"time"
	"errors"
	"sync/atomic"
	"runtime"
	"sync"
	"math/rand"
	"sort"
	"os"
	str "strings"
	"regexp"
	"bytes"
	"encoding/json"
	"flag"
	"os/exec"
	"io/ioutil"
	"syscall"
	"os/signal"
)

var println = fmt.Println

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


func worker1(done chan bool){
	println("woking ...")
	time.Sleep(time.Second)
	println("done")

	done<-true
}

func TestChannelSynchronization(){
	println("--->TestChannelSynchronization")

	done := make(chan bool, 1)
	go worker1(done)

	<-done
	println("finish")
}


func ping(pings chan<-string, msg string){
	pings<-msg
}

func pong(pings <-chan string, pongs chan<-string){
	msg:=<-pings
	pongs<-msg
}
func TestChannelDirections(){
	println("--->TestChannelDirections")

	pings:=make(chan string,1)
	pongs:=make(chan string,1)
	ping((chan<-string)(pings), "pressed message")
	pong((<-chan string)(pings), (chan<-string)(pongs))

	println("pongs", <-pongs)
}

func TestSelect(){
	println("--->TestSelect")

	c1:=make(chan string,1)
	c2:=make(chan string,1)

	go func (){
		time.Sleep(time.Second)
		c1<-"one"
	}()
	go func(){
		time.Sleep(time.Second*2)
		c2<-"two"
	}()

	for i:=0; i<2; i++{
		select {
		case msg1:=<-c1:
			println("received",msg1)
		case msg2:=<-c2:
			println("received",msg2)
		}
	}
}

func TestTimeouts(){
	println("--->TestTimeouts")

	c1:=make(chan string,1)
	go func(){
		time.Sleep(time.Second*2)
		c1<-"result1"
	}()

	select {
	case res:=<-c1:
		println(res)
	case <-time.After(time.Second*1):
		println("tiemout 1")
	}

	c2:=make(chan string,1)
	go func(){
		time.Sleep(time.Second*2)
		c2<-"result2"
	}()

	select {
	case res:=<-c2:
		println(res)
	case <-time.After(time.Second*3):
		println("timeout2")
	}
}

func TestNonBlockingChannel(){
	println("--->TestNonBlockingChannel")

	messages:=make(chan string)
	signals:=make(chan bool)

	select {
		case msg:=<-messages:
			println("received",msg)
	 	default:
			println("no messages received")
	}

	msg:="hi"
	select {
	case messages<-msg:
		println("msg sent", msg)
	default:
		println("no msg sent")
	}

	select {
	case msg:=<-messages:
		println("received msg",msg)
	case sig:=<-signals:
		println("received signals", sig)
	default:
		println("no activity")
	}
}

func TestClosingChannels(){
	println("--->TestClosingChannels")

	jobs:=make(chan int, 5)
	done:=make(chan bool)

	go func(){
		for {
			j,more:=<-jobs
			if more{
				println("received jobs", j)
			}else{
				println("receoved all jobs")
				done<-true
				return
			}
		}
	}()

	for i:=0; i<3; i++{
		jobs<-i
		println("sent job", i)
	}
	close(jobs)
	println("sent all jobs")
	<-done
}

func TestRangeOverChannels(){
	println("--->TestRangeOverChannels")

	queue:=make(chan string,2)
	queue<-"one"
	queue<-"two"
	close(queue)

	for elem := range queue{
		println(elem)
	}
}

func TestTimers(){
	println("--->TestTimers")

	timer1:=time.NewTimer(time.Second*2)

	<-timer1.C
	println("Timer1 expired")

	timer2:=time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		println("Timer2 expired")
	}()
	stop2:=timer2.Stop()
	if stop2 {
		println("Timer2 stopped")
	}
}

func TestTickers(){
	println("--->TestTickers")

	ticker := time.NewTicker(time.Millisecond*500)
	go func() {
		for t := range ticker.C{
			println(t)
		}
	}()
	time.Sleep(time.Second*2)
	ticker.Stop()
	println("ticker stopped")
}

func worker(id int, jobs <-chan int, result chan<-int, done chan int){
	for {
		j,more:= <-jobs
		if(more) {
			println("worker ", id, "processing job", j)
			time.Sleep(time.Second)
			result<- 2*j
		}else{
			done<-1
			break
		}
	}
	println("worker finish", id)
}
func collect(result chan int, done chan int){
	for {
		a,more:= <-result
		if(more) {
			println(a)
		}else{
			println("result end")
			done<-1
			break
		}
	}
}
func TestWorkerPool(){
	println("--->TestWorkerPool")

	jobs:=make(chan int, 100)
	result:=make(chan int, 100)
	done:=make(chan int)

	for i:=0;  i< 3; i++{
		go worker(i, (<-chan int)(jobs), (chan<-int)(result), done)
	}

	for j:=0; j< 9; j++{
		jobs<-j
	}
	close(jobs)

	go collect(result, done)

	for d:=0; d<4; d++{
		<-done
		if d == 2{
			close(result)
		}
	}
}


func TestRateLimiting(){
	println("--->TestRateLimiting")

	requests := make(chan int, 5)
	for i:=1; i<=5; i++{
		requests<-i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond*200)
	for req:=range requests{
		<-limiter
		println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i:=0; i<3; i++{
		burstyLimiter<-time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond*200){
			burstyLimiter<-t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i:=1; i<=5; i++{
		burstyRequests<-i
	}
	close(burstyRequests)

	for req:=range burstyRequests{
		t:= <-burstyLimiter
		println("request2", req, t)
	}
}


func TestAtomicCounters(){
	println("--->TestAtomicCounters")

	var ops uint64 = 0
	for i:=0; i < 500; i++{
		go func() {
			for i:=0; i <30; i++{
				atomic.AddUint64(&ops, 1)
				//ops++
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
	finalOps := atomic.LoadUint64(&ops)
	//finalOps := ops
	println("final ops", finalOps)
}

func TestMutexes(){
	println("--->TestMutexes")

	state := make(map[int]int)
	mutex := &sync.Mutex{}
	var opsRead,opsWrite uint64 = 0,0

	for r:=0 ; r < 100; r++{
		go func() {
			total := 0
			for {
				key:=rand.Intn(5)

				mutex.Lock()
				total+=state[key]
				mutex.Unlock()

				atomic.AddUint64(&opsRead, 1)
				runtime.Gosched()
			}
		}()
	}

	for w:=0; w<10; w++{
		go func() {
			for {
				key:=rand.Intn(5)
				val:=rand.Intn(100)

				mutex.Lock()
				state[key] = val
				mutex.Unlock()

				atomic.AddUint64(&opsWrite, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsReadF := atomic.LoadUint64(&opsRead)
	println("final opsRead", opsReadF)
	opsWriteF := atomic.LoadUint64(&opsWrite)
	println("final opsWrite", opsWriteF)

	mutex.Lock()
	println(state)
	mutex.Unlock()
}

type readOp struct {
	key int
	resp chan int
}
type writeOp struct {
	key int
	val int
	resp chan bool
}
func TestStatefulGoroutines(){
	println("--->TestStatefulGoroutines")

	var opsRead,opsWrite uint64 = 0,0

	reads := make(chan *readOp)
	writes:= make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read:=<-reads:
				read.resp<-state[read.key]
			case write:=<-writes:
				state[write.key] = write.val
				write.resp<-true
			}
		}
	}()

	for r :=0; r< 100; r++{
		go func() {
			read := &readOp{
			key:rand.Intn(5),
			resp:make(chan int)}
			for {
				read.key = rand.Intn(5)
				reads<-read
				<-read.resp
				atomic.AddUint64(&opsRead, 1)
			}
		}()
	}

	for w:=0; w<10; w++{
		go func() {
			write:= &writeOp{
			key:rand.Intn(5),
			val:rand.Intn(10),
			resp:make(chan bool)}
			for {
				write.key = rand.Intn(5)
				write.val = rand.Intn(10)
				writes<-write
				<-write.resp
				atomic.AddUint64(&opsWrite, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	opsReadF := atomic.LoadUint64(&opsRead)
	println("final opsRead", opsReadF)
	opsWriteF := atomic.LoadUint64(&opsWrite)
	println("final opsWrite", opsWriteF)
}


func TestSorting(){
	println("--->TestSorting")

	strs := []string{"d", "c", "a"}
	sort.Strings(strs)
	println("strings", strs)

	ints := []int{7,2,4}
	sort.Ints(ints)
	println("ints", ints)

	s:=sort.IntsAreSorted(ints)
	println("sorted", s)
}


type ByLength []string
func (s ByLength) Len()int{
	return len(s)
}
func (s ByLength) Swap(i,j int){
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i,j int) bool{
	return len(s[i]) < len(s[j])
}
func TestFunctionsSort(){
	println("--->TestFunctionsSort")

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	println(fruits)
}


func TestPanic(){
	println("--->TestPanic")

	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}


func TestDefer(){
	println("--->TestDefer")

	file := createFile("/tmp/defer.txt")
	defer closeFile(file)
	writeFile(file)
}
func createFile(p string) *os.File{
	println("creating")
	f,err:=os.Create(p)
	if err!=nil {
		panic(err)
	}
	return f
}
func writeFile(file *os.File){
	println("writing")
	fmt.Fprintln(file, "data")
}
func closeFile(file *os.File){
	println("closing")
	file.Close()
}


func TestCollectionFunctions(){
	println("--->TestCollectionFunctions")

	var strs = []string{"peach", "apple", "pear", "plum"}

	println("index", Index(strs,"apple"))
	println("include", Include(strs,"grape"))
	println("any", Any(strs, func(v string)bool{
				return str.HasPrefix(v, "p")
			}))
	println("all", All(strs, func(v string)bool{
				return str.HasPrefix(v, "p")
			}))
	println("filter", Filter(strs, func(v string)bool{
				return str.Contains(v, "e")
			}))
	println("map", Map(strs, str.ToUpper))
}
func Index(vs []string, t string) int{
	for i,s := range vs{
		if s == t {
			return i
		}
	}
	return -1
}
func Include(vs []string, t string) bool{
	return Index(vs, t) > 0
}
func Any(vs []string, f func(string) bool) bool{
	for _,s := range vs{
		if f(s) {
			return true
		}
	}
	return false
}
func All(vs []string, f func(string) bool) bool{
	for _,s := range vs{
		if !f(s) {
			return false
		}
	}
	return true
}
func Filter(vs []string, f func(string) bool)[]string{
	vsf := make([]string, 0)
	for _,s := range vs{
		if f(s) {
			vsf = append(vsf, s)
		}
	}
	return vsf
}
func Map(vs []string, f func(string) string)[]string{
	vsm := make([]string, len(vs))
	for i, s := range vs{
		vsm[i] = f(s)
	}
	return vsm
}


func TestStringFunctions(){
	println("--->TestStringFunctions")

	println("Contains", str.Contains("test","es"))
	println("Count", str.Count("test","t"))
	println("HasPrefix", str.HasPrefix("test", "t"))
	println("HasSuffix", str.HasSuffix("test", "st"))
	println("Index",str.Index("test", "e"))
	println("Join",[]string{"a","b"}, "-")
	println("Repeat",str.Repeat("a", 5))
	println("Replace", str.Replace("foo","o","i", -1))
	println("Replace", str.Replace("foo","o","i", 1))
	println("Split", str.Split("a-b-c-d-e", "-"))
	println("ToLower", str.ToLower("TEST"))
	println("ToUpper", str.ToUpper("test"))

	println()
	println("len", len("hello"))
	println("char", "hello"[0])
}

type point struct {
	x,y int
}
func TestStringFormatting(){
	println("--->TestStringFormatting")

	p := point{1,2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 123)
	fmt.Printf("%c\n", 123)
	fmt.Printf("%x\n", 123)
	fmt.Printf("%f\n", 12.3)
	fmt.Printf("%e\n", 12000000000000000.3)
	fmt.Printf("%E\n", 12000000000000000.3)
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s:=fmt.Sprintf("a %s", "string")
	println(s)

	fmt.Fprintf(os.Stderr,"an %s\n", "error")
}


func TestRegularExpressions(){
	println("--->TestStringFormatting")

	match, str := regexp.MatchString("p([a-z]+)ch", "peach")
	println(match, str)

	r,_ := regexp.Compile("p([a-z]+)ch")

	println("match", r.MatchString("peach"))
	println("find",  r.FindString("peach punch"))
	println("find index",  r.FindStringIndex("peach punch"))
	println("find sub",  r.FindStringSubmatch("peach punch"))
	println("find subindex",  r.FindStringSubmatchIndex("peach punch"))
	println("find all", r.FindAllString("peach punch pinch", -1))
	println("find all subindex", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	println("find all 2", r.FindAllString("peach punch pinch", 2))
	println("match", r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	println(r)
	println("replace all", r.ReplaceAllString("a peach", "<fruit>"))

	in:=[]byte(" a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	println(string(out))
}


type Response1 struct {
	Page int
	Fruits []string
}
type Response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}
func TestJSON(){
	println("--->TestJSON")

	bolB,_ := json.Marshal(true)
	println(string(bolB))

	intB,_ := json.Marshal(1)
	println(string(intB))

	fltB, _ := json.Marshal(2.34)
	println(string(fltB))

	strB, _ := json.Marshal("gopher")
	println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	println(string(mapB))

	res1D := &Response1{
		Page:1,
		Fruits:[]string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	println(string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var data map[string]interface {}

	if err := json.Unmarshal(byt, &data); err != nil{
		panic(err)
	}
	println(data)

	num:=data["num"].(float64)
	println(num)
	strs:=data["strs"].([]interface{})
	str1:=strs[0].(string)
	println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res2 := &Response2{}
	json.Unmarshal([]byte(str), res2)
	println(res2)
	println(res2.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}


func TestCommandLineArguments(){
	println("--->TestCommandLineArguments")

	argsWithProg := os.Args
	argsWithNoProg := os.Args[1:]
	arg := os.Args[3]

	println(argsWithProg)
	println(argsWithNoProg)
	println(arg)
}


func TestCommandLineFlag(){
	println("--->TestCommandLineFlag")

	wordPtr := flag.String("word","foo","a string")
	intPrt := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	println("word:",*wordPtr)
	println("numb:",*intPrt)
	println("fork:",*boolPtr)
	println("svar:",svar)
	println("tail:",flag.Args())
}

func TestEnvironmentVariables(){
	println("--->TestEnvironmentVariables")

	os.Setenv("FOO", "1")
	println("FOO:", os.Getenv("FOO"))
	println("BAR:", os.Getenv("BAR"))

	for _,e := range os.Environ(){
		pair := str.Split(e, "=")
		println(pair[0],"=", pair[1])
	}
}

func TestSpawningProcesses(){
	println("--->TestSpawningProcesses")

	dateCmd := exec.Command("date")
	dateOut,err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	println(">date", string(dateOut))

	grepCmd := exec.Command("grep", "hello")
	grepIn,_ := grepCmd.StdinPipe()
	grepOut,_ := grepCmd.StdoutPipe()
	grepCmd.Start()

	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()

	grepBytes,_ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	println(">grep hello")
	println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}


func TestExecingProcesses(){
	println("--->TestExecingProcesses")

	binary,lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}


func TestSignals(){
	println("--->TestSignals")

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		println(sig)
		done<-true
	}()

	println("awaiting signal")
	<-done
	println("exiting")
}

