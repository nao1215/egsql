package cache

import (
	"container/list"
	"sync"
	"testing"
)

func TestLRU_Insert(t *testing.T) {
	t.Run("[Success] Insert data and return nil", func(t *testing.T) {
		lru := NewLRU(1)
		result := lru.Insert("key", 100)
		if result != nil {
			t.Errorf("Insert result(=%v) is not nil", result)
		}
	})

	t.Run("[Success] Insert one data, remove one victim data", func(t *testing.T) {
		lru := NewLRU(1)
		result := lru.Insert("key1", 100)
		if result != nil {
			t.Errorf("Insert result(=%v) is not nil", result)
		}

		want := 100
		wantKey := "key1"
		got := lru.Insert("key2", 1000)
		if got.(*list.Element).Value.(*entry).value.(int) != want {
			t.Errorf("mismatch want:%d, got:%d", want, got.(*list.Element).Value.(*entry).value.(int))
		}
		if got.(*list.Element).Value.(*entry).key.(string) != wantKey {
			t.Errorf("mismatch want:%s, got:%s", wantKey, got.(*list.Element).Value.(*entry).key.(string))
		}
	})
}

func TestLRU_Get(t *testing.T) {
	t.Run("[Success] Get data", func(t *testing.T) {
		lru := NewLRU(3)
		lru.Insert("key1", 100)
		lru.Insert("key2", 200)
		lru.Insert("key3", 300)

		want := 200
		got := lru.Get("key2")
		if want != got {
			t.Errorf("mismatch want:%d, got:%d", want, got)
		}
	})

	t.Run("[Error] Not get non-existent data", func(t *testing.T) {
		lru := NewLRU(3)
		lru.Insert("key1", 100)
		lru.Insert("key2", 200)
		lru.Insert("key3", 300)

		got := lru.Get("non-exist-key")
		if got != nil {
			t.Errorf("mismatch want:nil, got:%d", got)
		}
	})
}

func TestLRU_Concurrency(t *testing.T) {
	loopNum := 1000000
	lru := NewLRU(loopNum)

	var wg sync.WaitGroup

	for i := 0; i < loopNum; i++ {
		wg.Add(1)
		num := i
		go func() {
			lru.Insert(num, num)
			wg.Done()
		}()
	}
	wg.Wait()

	if lru.Len() != loopNum {
		t.Errorf("Some data could not be inserted")
	}
}
