package main

// START OMIT
import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano()) // Seed the rng
}

func trackTime(s time.Time, msg string) {
	fmt.Println(msg, ":", time.Since(s))
}

func some_costly_fn() {
	defer trackTime(time.Now(), "some_costly_fn") // HL

	random_wait := rand.Intn(1000)
	time.Sleep(time.Duration(random_wait) * time.Millisecond) // Some complicated work
}

func main() {
	some_costly_fn()
}

// END OMIT
