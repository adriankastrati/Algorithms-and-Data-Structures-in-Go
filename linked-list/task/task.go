package task

import (
	"fmt"
	"linked-list/list"
	"math"
	"time"
)

var (
	loops   int = 250
	maxSize int = 20000
)

func TimeListAOnB() {
	var (
		listASize int = 0
		listBSize int = 1000
		tDelta    float64
		t0        time.Time
	)

	//list A appends to list B, list A should vary in size while list B is fixed

	//increase length of linked list
	for mult := 0; listASize < maxSize; mult++ {
		//restart time delta every iteration of loop

		tDelta = 0
		//size of list A
		listASize = int(math.Pow(2, float64(mult)))

		for i := 0; i < loops; i++ {

			//create inside of loop to minimize the effect on time complexity, should always be the same
			listB := list.MakeLinkedList()
			for k := 0; k < listBSize; k++ {
				listB.AppendNode(k)
			}

			//append nodes to linked list through appendation
			listA := list.MakeLinkedList()
			for k := 0; k < listASize; k++ {
				listA.AppendNode(k)
			}

			t0 = time.Now()
			listB.AppendList(&listA)
			tDelta += float64(time.Since(t0))
		}

		tDelta /= float64(loops)
		fmt.Printf("%d %f\n", listASize, tDelta/1000)
	}

}

func TimeListBOnA() {
	var (
		listASize int = 0
		listBSize int = 1000
		tDelta    float64
		t0        time.Time
	)

	//list A appends to list B, list A is fixed while list B increases in size

	//increase length of linked list
	for mult := 0; listASize < maxSize; mult++ {
		//restart time delta every iteration of loop

		tDelta = 0
		//size of list A
		listASize = int(math.Pow(2, float64(mult)))

		for i := 0; i < loops; i++ {

			//create inside of loop to minimize the effect on time complexity, should always be the same
			listB := list.MakeLinkedList()
			for k := 0; k < listBSize; k++ {
				listB.AppendNode(k)
			}

			//append nodes to linked list through appendation
			listA := list.MakeLinkedList()
			for k := 0; k < listASize; k++ {
				listA.AppendNode(k)
			}

			t0 = time.Now()
			listA.AppendList(&listB)
			tDelta += float64(time.Since(t0))
		}

		tDelta /= float64(loops)
		fmt.Printf("%d %f\n", listASize, tDelta/1000)
	}

}

func TimeSliceBOnA() {
	var (
		sliceB    []int
		sliceA    []int
		sliceSize int = 1
		tDelta    float64
	)

	sliceB = make([]int, 1000)

	for i, _ := range sliceB {
		sliceB[i] = i
	}

	for mult := 5; sliceSize < maxSize; mult++ {
		tDelta = 0

		for i := 0; i < loops; i++ {
			sliceSize = int(math.Pow(2, float64(mult)))
			sliceA = make([]int, sliceSize)

			for i, _ := range sliceA {
				sliceA[i] = i
			}

			t0 := time.Now()

			list.AppendSlices(sliceA, sliceB)

			tDelta += float64(time.Since(t0))

		}

		tDelta /= float64(loops)

		fmt.Printf("%d %f\n", sliceSize, tDelta/1000)

	}
}

func TimeAllocateSlice() {
	var (
		slice     []int
		sliceSize int = 1
		tDelta    float64
		t0        time.Time
	)

	for mult := 5; sliceSize < maxSize; mult++ {
		tDelta = 0

		for i := 0; i < loops; i++ {
			t0 = time.Now()

			sliceSize = int(math.Pow(2, float64(mult)))
			slice = make([]int, sliceSize)

			tDelta += float64(time.Since(t0))
		}

		tDelta /= float64(loops)

		fmt.Printf("%d %f\n", sliceSize, tDelta/1000)

	}
	slice[0] = 1
}

func TimeAllocateList() {
	var (
		tDelta   float64
		t0       time.Time
		listSize int
	)

	//list A appends to list B, list A is fixed while list B increases in size

	//increase length of linked list
	for mult := 0; listSize < maxSize; mult++ {
		//restart time delta every iteration of loop
		tDelta = 0

		//size of list
		listSize = int(math.Pow(2, float64(mult)))

		for i := 0; i < loops; i++ {
			//append nodes to linked list through appendation
			t0 = time.Now()

			listA := list.MakeLinkedList()
			for k := 0; k < listSize; k++ {
				listA.AppendNode(k)
			}

			tDelta += float64(time.Since(t0))
		}

		tDelta /= float64(loops)
		fmt.Printf("%d %f\n", listSize, tDelta/1000)
	}

}

func LinkedStack() {
	stack := list.MakeStack()

	el1 := list.MakeElement(5)
	el2 := list.MakeElement(6)

	stack.Push(el1)

	println(stack.Top.GetVal())

	stack.Push(el2)

	println(stack.Top.GetVal())

	println(stack.Top.Under.GetVal())

	n := stack.Pop()
	println(n.GetVal())
	println(stack.Top.GetVal())
}
