package main

import (
	"cleanly/tools"
	"fmt"
	"os"
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
