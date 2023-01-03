package gopkgdemo

import (
	"math/rand"
	"time"
	"fmt"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func Lotto() [6]int {
	var l [6]int
	for i := 0; i < len(l); i++ {
		l[i] = rand.Intn(48) + 1
		for n := 0; n < i; n++ {
			if l[i] == l[n] {
				i--
			}
		}
	}
	return l
} 

func Ausgabe(a [6]int) {
	for i, l := range a {
		fmt.Printf("%d: %d\n", (i + 1), l)
	}
}