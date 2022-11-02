package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	path, err := os.Getwd()
	check(err, "")
	fmt.Println(path)
	files, e := ioutil.ReadDir("./")
	check(e, "")

	finish := make(chan string)

	for _, f := range files {
		f := f
		go func() {
			f := f
			if f.IsDir() {
				pull(f.Name())
				finish <- f.Name()
			}
		}()
	}

	for i := 0; i < len(files)-1; i++ {
		<-finish
	}

}

func check(e error, str string) {
	if e != nil {
		fmt.Printf("Error: [%s] [%s]\n", str, e)
	}
}

func pull(name string) {
	cmd := exec.Command("bash", "-c", "cd "+name+"/")
	_, errs := cmd.Output()
	check(errs, "'bash', '-c', 'cd '"+name+"'/'")
	cmd.Start()
	cmd = exec.Command("git", "pull")
	cmd.Dir = name
	log, errs := cmd.Output()
	fmt.Println("pull "+name, string(log))
	check(errs, name)
	cmd.Start()
}
