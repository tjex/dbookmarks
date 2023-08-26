package main

import (
	"bufio"
	"fmt"
	"mvdan.cc/xurls/v2"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// the first arg of Args is the command itself
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// check if dmenu-mac is in PATH
	dmenuMacPath, err := exec.LookPath("dmenu-mac")
	if err != nil {
		fmt.Println(err)
	}

	// url extract object
	xu := xurls.Relaxed()
	bookmarks := []string{}
	for scanner.Scan() {
		line := xu.FindString(scanner.Text())
		bookmarks = append(bookmarks, line)
	}

	// create dmenu-mac command (executes when called)
	dmenu := exec.Command(dmenuMacPath)

	// Join the bookmarks into a single string separated by newline characters
	// could we just read the output of scanner.Text() instead?
	input := strings.Join(bookmarks, "\n")

	// Set up a pipe for the command's stdin and write the input to it
	dmenu.Stdin = strings.NewReader(input)

	output, err := dmenu.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}

	// trim any leading/trailing whitespaces from output
	url := strings.TrimSpace(string(output))

	open := exec.Command("open", url)
	err = open.Run()
	if err != nil {
		fmt.Println(err)
	}
}
