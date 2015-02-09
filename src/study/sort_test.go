package study

import (
	"fmt"
	"time"
	"sort"
	"math/rand"
	"testing"
)

func TestSort(t *testing.T){
//	testQuickSort()
//	testQuickSort2()
//	testQuickSort3()
//	testHeapSort()
//	testMergeSort()
}

const size int = 10000000

type IntSortable int

func (a IntSortable) Great(b interface{}) bool{
	return a > b.(IntSortable)
}

func testQuickSort(){

	data := make([]Sortable, size, size)
	for index:=0; index < size; index++{
		data[index] = IntSortable(rand.Intn(size))
	}
	now := time.Now()
	fmt.Println("start = ",now)
	QuicksortT(data, false)
	end := time.Now()
	fmt.Println("end = ",end ,": dx = ", (end.Unix()-now.Unix()))
}

func testQuickSort2(){
	data := make([]int, size, size)
	for index:=0; index < size; index++{
		data[index] = rand.Intn(size)
	}
	now := time.Now()
	fmt.Println("start = ",now)
	sort.Ints(data)
	end := time.Now()
	fmt.Println("end = ",end ,": dx = ", (end.Unix()-now.Unix()))
}

func testQuickSort3(){
	data := make([]int, size, size)
	for index:=0; index < size; index++{
		data[index] = rand.Intn(size)
	}
	now := time.Now()
	fmt.Println("start = ",now)
	Quicksort(data, true)
	end := time.Now()
	fmt.Println("end = ",end ,": dx = ", (end.Unix()-now.Unix()))
}

func testHeapSort(){
	data := make([]int, size, size)
	for index:=0; index < size; index++{
		data[index] = rand.Intn(size)
	}
	now := time.Now()
	fmt.Println("start = ",now)
	MinHeapSort(data)
	end := time.Now()
	fmt.Println("end = ",end ,": dx = ", (end.Unix()-now.Unix()))

	if size < 100 {
		fmt.Println(data)
	}
}

func testMergeSort(){
	data := make([]int, size, size)
	for index:=0; index < size; index++{
		data[index] = rand.Intn(size)
	}
	now := time.Now()
	fmt.Println("start = ",now)
	MergeSort(data)
	end := time.Now()
	fmt.Println("end = ",end ,": dx = ", (end.Unix()-now.Unix()))

	if size < 100 {
		fmt.Println(data)
	}
}

