package common

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func GitClone(gitUrl string, targetDirectory string) {

	_, err := git.PlainClone(targetDirectory, false, &git.CloneOptions{
		URL:      gitUrl,
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Println(err)
	}
}
