package main

import (
	"crypto/rand"
	"fmt"
	"log"
)

// Created for https://devtidbits.com/2020/08/10/go-unicode-random-string-generator/

const random string = "🥺😂🥰😊😍😊😝🤗abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321 .!?"

func main() {
	s, err := randString(20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}

func randString(n int) (string, error) {
	s, r := make([]rune, n), []rune(random)
	for i := range s {
		p, err := rand.Prime(rand.Reader, len(r))
		if err != nil {
			return "", fmt.Errorf("random string n %d: %w", n, err)
		}
		x, y := p.Uint64(), uint64(len(r))
		// fmt.Printf("x: %d y: %d\tx %% y = %d\trandom[%d] = %q\n", x, y, x%y, x%y, string(r[x%y]))
		s[i] = r[x%y]
	}
	return string(s), nil
}
