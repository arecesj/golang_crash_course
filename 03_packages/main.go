package main

import (
	"fmt"
	"math"

	"github.com/arecesj/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.3999999))
	fmt.Println(math.Ceil(2.3999999))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("Hello"))
}
