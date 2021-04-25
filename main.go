package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Seed(time.Now().UnixNano())

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	if text == "" {
		fmt.Println("")
	} else {
		fmt.Println(roll(text, r))
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func roll(in string, r *rand.Rand) int {
	var out int

	re := regexp.MustCompile(`\d+`)
	re2 := regexp.MustCompile(`(\d+[d]\d+$|\d+[d]\d+\+\d+$)`)
	if !re2.MatchString(in) {
		log.Fatal("bad input")
		return -1
	}
	arr := re.FindAllString(in, 3)
	num1, num2 := arr[0], arr[1]

	amt, err := strconv.Atoi(num1)
	if err != nil {
		log.Fatal("error processing roll", err.Error())
		return -1
	}
	die, err := strconv.Atoi(num2)
	if err != nil {
		log.Fatal("error processing roll", err.Error())
		return -1
	}

	for i := 1; i <= amt; i++ {
		out += r.Intn(die) + 1
	}

	if len(arr) == 3 {
		bonus, err := strconv.Atoi(arr[2])
		if err != nil {
			log.Fatal("error processing roll", err.Error())
			return -1
		}
		out += bonus
	}

	return out
}
