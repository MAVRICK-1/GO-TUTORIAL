package main

import (
	"fmt"
	"sync"
	
)

type post struct{
	views int
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup,){
	defer func (){ 
		wg.Done()
	}()

	p.mu.Lock() //locked that no other thread can access it 

	p.views+=1

	p.mu.Unlock() //unlocked
}

func main () {

	var wg  sync.WaitGroup
	p := post{
		views: 1,
	}

	for i:=1 ; i<100 ; i++ {
		wg.Add(1)
		go p.inc(&wg)
	}
	wg.Wait()

	fmt.Println(p.views)


}