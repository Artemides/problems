package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

type cake int

type Bakery struct {
	Verbose        bool
	Cakes          int
	BakeTime       time.Duration
	BakeStdDev     time.Duration
	BakeBuf        int
	NumIcers       int
	IceTime        time.Duration
	IceStdDev      time.Duration
	IceBuf         int
	InscribeTime   time.Duration
	InscribeStdDev time.Duration
}

func (b *Bakery) baker(bakes chan<- cake) {
	for i := 0; i < b.Cakes; i++ {
		c := cake(i)
		if b.Verbose {
			fmt.Printf("backing %v\n", c)
		}
		work(b.BakeTime, b.BakeStdDev)
		bakes <- c
	}
	close(bakes)
}

func (b *Bakery) icer(ices chan<- cake, bakes <-chan cake) {
	for bake := range bakes {
		if b.Verbose {
			fmt.Printf("\ticing %v\n", bake)
		}
		work(b.IceTime, b.IceStdDev)
		ices <- bake
	}
	close(ices)
}

func (b *Bakery) inscriber(iced <-chan cake) {
	for i := 0; i < b.Cakes; i++ {
		icedCake := <-iced
		if b.Verbose {
			fmt.Printf("\t\tinscribing %v\n", icedCake)
		}
		work(b.InscribeTime, b.InscribeStdDev)
		if b.Verbose {
			fmt.Printf("\t\tCake (%v)Finished\n", icedCake)
		}

	}

}

func (b *Bakery) Work(runs int) {

	for run := 0; run < runs; run++ {
		baked := make(chan cake, b.BakeBuf)
		iced := make(chan cake, b.IceBuf)
		go b.baker(baked)
		for i := 0; i < b.NumIcers; i++ {
			go b.icer(iced, baked)
		}
		b.inscriber(iced)

	}
}

func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))

	time.Sleep(delay)
}

func BakeryMain() {
	backery := Bakery{
		Verbose:        true,
		Cakes:          20,
		BakeTime:       1 * time.Second,
		BakeStdDev:     1 * time.Second,
		BakeBuf:        1,
		NumIcers:       4,
		IceTime:        3 * time.Second,
		IceStdDev:      1 * time.Second,
		IceBuf:         1,
		InscribeTime:   2 * time.Second,
		InscribeStdDev: 1 * time.Second,
	}
	backery.Work(1)
}
