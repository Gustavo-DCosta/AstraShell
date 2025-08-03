package pushlynx

import (
	"fmt"
	"runtime"
)

func RunGit() error {
	fmt.Print(`      *     .     *       *   .     *     .
    *   .    *    A S T R A S H E L L   *    .
       .       *     *   .    *    .      *  
            The Terminal from the Stars
`)
	if runtime.GOOS == "windows" {
		fmt.Println("You are on windows!")
	} else {
		fmt.Println("Why complicate life buddy?")
	}
	return nil
}
