package obj_cache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new obj")
			return 100
		},
	}
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(30)
	runtime.GC() //GC会清除sync.pool的缓存私有对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
	v2, _ := pool.Get().(int) //每次都被取出来了，私有对象没有了，需要新建
	fmt.Println(v2)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new obj")
			return 10
		},
	}

	pool.Put(101)
	pool.Put(102)
	pool.Put(103)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
