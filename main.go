package main

import (
    "fmt"
    "strings"

    "path/filepath"

    "os"
    "os/exec"

    "github.com/yookoala/realpath"
)

func pull() {
    cmd := exec.Command("git", "pull")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
		fmt.Println("Run failed - ", err)
	}
}

func fetchDir(fp string, fi os.FileInfo, err error) error {
    // can't walk here,
    // but continue walking elsewhere
    if err != nil {
        return nil
    }
    if !!fi.IsDir() && fp != "." {
        oldPath, _ := realpath.Realpath(".")
        matched, _ := filepath.Match("*/.git", fp)
        if matched {
            dirs := strings.Split(fp, ".")
            fmt.Println("Enter in :", dirs[0])
            fmt.Println("Trying to pull...")
            os.Chdir(dirs[0])
            pull()
            os.Chdir(oldPath)
            fmt.Println()
        }
    }
    return nil
}

func main() {
    fmt.Println()
    fmt.Println("### Running go-pull script - Fetching directories...")
    filepath.Walk(".", fetchDir)
}
