package tui

import (
	"bufio"
	"fmt"
	"github.com/gzipchrist/dont_at_me/pkg/components"
	"github.com/gzipchrist/dont_at_me/pkg/cursor"
	"github.com/gzipchrist/dont_at_me/pkg/style"
	"github.com/gzipchrist/dont_at_me/pkg/username"
	"os"
	"strings"
	"time"
)

func Run() {
	fmt.Printf("%s\n%s\n", style.Cyan.Colorize(components.Header), components.Prompt)

	for {
		fmt.Printf("\n%s", components.TextInput)

		userInput := getUserInput()
		fmt.Print(cursor.Hide)
		fmt.Print(cursor.Up)
		fmt.Print(cursor.ClearLine)
		fmt.Print("  @ ") // Rerender cleared prompt
		fmt.Print(cursor.Down)
		fmt.Printf(cursor.ClearLine)

		if userInput == "q" {
			fmt.Printf(cursor.ClearAfter)
			return
		}

		fmt.Printf("Results for \"%s\"\n\n", userInput)

		t := time.Now()

		username.CheckAvailabilityConcurrent(userInput)

		fmt.Printf("\n    completed search in %v\n\n", time.Since(t))
		fmt.Print(strings.Repeat(cursor.Up, 15))
		fmt.Print(cursor.ClearLine)
		fmt.Print(cursor.Show)
	}

}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
