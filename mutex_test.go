package main

//https://github.com/Dima3108/mygorep1
/*
https://habr.com/ru/articles/268585/
https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
*/
import (
	"sync"
	"testing"

	"github.com/Dima3108/mygorep1"
)

/*
	func mytest() {
		mym := mygorep1.NewMyMutex()
		mym.Lock()
		mym.Unlock()
		mym.Wait()
	}
*/

/*func BenchmarkSample(b *testing.B) {
	b.SetBytes(2)
	for i := 0; i < b.N; i++ {
		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
}*/

const IterCount = 1000

func mumain() bool {
	mym := mygorep1.NewMyMutex()
	var wg sync.WaitGroup
	wg.Add(3)
	l1 := []int{}
	l2 := []int{}
	l3 := []int{}
	go func() {
		mym.Lock()
		//fmt.Println("T1")
		i := 0
		offset := 0
		for true {
			i++
			l1 = append(l1, (offset*IterCount)+i)
			if i > IterCount {
				break
			}
		}
		//fmt.Println("!T1")
		mym.Unlock()
	}()
	go func() {
		mym.Lock()
		//fmt.Println("T2")
		i := 0
		offset := 1
		for true {
			i++
			l2 = append(l2, (offset*IterCount)+i)
			if i > IterCount {
				break
			}
		}
		//fmt.Println("!T2")
		mym.Unlock()
	}()
	go func() {
		mym.Lock()
		//fmt.Println("T3")
		i := 0
		offset := 2
		for true {
			i++
			l3 = append(*&l3, (offset*IterCount)+i)
			if i > IterCount {
				break
			}
		}
		//fmt.Println("!T3")
		mym.Unlock()
	}()
	mym.Wait()
	if len(l1) != len(l2) || (len(l2) != len(l3)) {
		return false
	}
	return false
}
func TestMutex(t *testing.T) {
	t.Logf("init test!\n")
	{
		testID := 0
		//mumain()
		t.Logf("ID:{%d}", testID)
		{
			if mumain() {
				t.Errorf("error mutex test!\n")
			}
		}
	}

}
func BenchmarkMutex(b *testing.B) {
	//for i := 0; i < b.N; i++ {
	//mytest()

	/*mym := mygorep1.NewMyMutex()

	mym.Lock()
	mym.Unlock()
	mym.Wait()*/
	//fmt.Println("start iter!")
	mumain()
	//fmt.Println("stop iter")
	//}
}

/*
go test -bench=Mutex -count 5 | tee old.txt
https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
*/
