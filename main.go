package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	var out [][]string
	start := time.Now()

	fmt.Println("starting...")

	startidx := 4
	endidx := 5

	charset := ""

	charset += "abcdefghijklmnopqrstuvwxyz"
	// charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// charset += "0123456789"

	for i := startidx; i <= endidx; i++ {
		out = append(out, generateString(i, charset, nil))
	}

	for i := range out {
		fmt.Printf("%d passwords of length %d for a charset of length %d\n", len(out[i]), i+startidx, len(charset))
	}

	fmt.Println("concuted in", time.Since(start))
	// fmt.Println(out)
}

func generateString(passwdLen int, charset string, wg *sync.WaitGroup) []string {
	var out []string

	if wg != nil {
		defer wg.Done()
	}

	possibilities := math.Pow(float64(len(charset)), float64(passwdLen))
	for i := 0; i < int(possibilities); i++ {
		tempout := ""
		val := i

		for j := 0; j < passwdLen; j++ {
			ch := val % len(charset)
			tempout = fmt.Sprintf("%c%s", charset[ch], tempout)
			val = val / len(charset)
		}
		out = append(out, tempout)
	}

	return out
}
