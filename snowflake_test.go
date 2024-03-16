package snowflake

import (
	"sync"
	"testing"
)

func Test_idWorker_NextID(t *testing.T) {
	idGene := NewIDGenerator()
	id := idGene.NextID()
	for i := 0; i < 100000; i++ {
		if id == idGene.NextID() {
			t.Fatalf("id: %d is not unique", id)
		}
		id = idGene.NextID()
	}
	t.Logf("id: %d", id)
}

func Test_Unique(t *testing.T) {
	delta := 5000

	idGene := NewIDGenerator()
	idMap := make(map[int64]struct{})
	wg := sync.WaitGroup{}
	ch := make(chan int64, delta*4096)
	wg.Add(delta)
	for i := 0; i < delta; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 4096; i++ {
				id := idGene.NextID()
				ch <- id
			}
		}()
	}
	wg.Wait()
	close(ch)
	for id := range ch {
		if _, ok := idMap[id]; ok {
			t.Fatalf("id: %d is not unique", id)
		}
		idMap[id] = struct{}{}
	}
}

func Test_Rent_Unique(t *testing.T) {
	delta := 5000

	idGene := NewIDGenerator(WithRent())
	idMap := make(map[int64]struct{})
	wg := sync.WaitGroup{}
	ch := make(chan int64, delta*4096)
	wg.Add(delta)
	for i := 0; i < delta; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 4096; i++ {
				id := idGene.NextID()
				ch <- id
			}
		}()
	}
	wg.Wait()
	close(ch)
	for id := range ch {
		if _, ok := idMap[id]; ok {
			t.Fatalf("id: %d is not unique", id)
		}
		idMap[id] = struct{}{}
	}
}
