package solution

import (
	"fmt"
	"time"
)

func Track(start time.Time, msg string) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}
