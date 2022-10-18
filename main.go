package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	path, err := os.Getwd()
	check(err)
	fmt.Println(path)
	files, e := ioutil.ReadDir("./")
	check(e)

	finish := make(chan string)

	for _, f := range files {
		f := f
		go func() {
			f := f
			if f.IsDir() {
				goThere(f.Name())
				finish <- f.Name()
			}
		}()
	}

	for i := 0; i < len(files)-1; i++ {
		fmt.Println(<-finish)
	}

}

func check(e error) {
	if e != nil {
		fmt.Printf("Error: [%s]\n", e)
	}
}

func pull(name string) {
	cmd := exec.Command("git", "pull")
	cmd.Dir = name
	log, errs := cmd.Output()
	check(errs)
	fmt.Println(string(log))
	cmd.Start()
}
func goThere(name string) {
	cmd := exec.Command("bash", "-c", "cd "+name+"/")
	log, errs := cmd.Output()
	check(errs)
	fmt.Println(string(log))
	cmd.Start()
	pull(name)
}
