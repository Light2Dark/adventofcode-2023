package utils

import (
	"time"
	"fmt"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %v time\n", name, elapsed)
}