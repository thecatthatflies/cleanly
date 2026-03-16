package main

import (
	"cleanly/tools"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	if len(os.Args) > 1 && (os.Args[1] == "help" || os.Args[1] == "-h") {
		fmt.Println(`cleanly — sort any folder by file type

usage:
	cleanly .                   sort current directory
	cleanly ~/Downloads         sort a specific folder
  	cleanly . -c "images audio" only sort specific categories
  	cleanly --no-clean          skip cleanup verification
  	cleanly --undo              reverse the last sort

categories:
  	images, audio, video, documents, archives, apps, data, code, other`)
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "update" {
		fmt.Print("updating cleanly")
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					fmt.Print(".")
					time.Sleep(500 * time.Millisecond)
				}
			}
		}()

		cmd := exec.Command("go", "install", "github.com/thecatthatflies/cleanly@latest")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		done <- true

		fmt.Println()
		if err != nil {
			fmt.Println("update failed:", err)
			os.Exit(1)
		}
		fmt.Println("updated.")
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "update" {
		fmt.Println("updating cleanly...")
		cmd := exec.Command("go", "install", "github.com/thecatthatflies/cleanly@latest")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("update failed:", err)
			os.Exit(1)
		}
		fmt.Println("updated.")
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "remove" {
		tools.Remove(os.Args[2:])
		return
	}
	input := tools.ParseInput()

	if input.Undo {
		err := tools.Undo()
		if err != nil {
			fmt.Println("error undoing last sort:", err)
			os.Exit(1)
		}
		fmt.Println("undone.")
		return
	}

	files := tools.Scan(input.Path)
	identified := tools.Identify(files, input.Categories)

	err := tools.MakeFolders(input.Path, identified)
	if err != nil {
		fmt.Println("error creating folders:", err)
		os.Exit(1)
	}

	err = tools.Sort(input.Path, identified)
	if err != nil {
		fmt.Println("error sorting files:", err)
		os.Exit(1)
	}

	if !input.NoClean {
		err = tools.Clean(input.Path, identified)
		if err != nil {
			fmt.Println("error during cleanup:", err)
			os.Exit(1)
		}
	}

	fmt.Println("done.")
}
