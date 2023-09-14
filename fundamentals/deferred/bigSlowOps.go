package deferred

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

func BigSlowOperation() {
	defer trace("BigSlowOperation")()
	time.Sleep(10 * time.Second)

}
func trace(msg string) func() {
	start := time.Now()
	log.Printf(" enter %s", msg)
	return func() {
		log.Printf("exit %s (%ss)", msg, time.Since(start))
	}
}

func Double(x int) int {
	return x + x
}

func Tripple(x int) (res int) {
	defer func() { res += x }()
	return Double(x)
}

// defer statement in loop

func ProcessFiles() error {
	fileNames := []string{}
	for _, filename := range fileNames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		// it stacks n deferred opearions and
		// no file will be closed until the last file gets processed
		// which might run out of file descriptors to some f.close() operations
		//process file ...
	}
	return nil
}

func ProcessFiles2() error {
	fileNames := []string{}
	for _, filename := range fileNames {
		doSomethingWithFile(filename)
	}
	return nil
}

// this might be a solution
// since this is function the deferred operarion will run
// at the very end of this function even beign inside a loop
// it will do every iteration
func doSomethingWithFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	//do somethink with file
	return nil
}

func FetchUrlAndWriteOnFile() (filename string, n int64, err error) {
	const url = "https://golang.org"
	response, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer response.Body.Close()

	local := path.Base(response.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	//dont't defer function which report erros after its return
	// wrong: defer f.Close()
	n, err = io.Copy(f, response.Body)
	if closeErr := f.Close(); closeErr != nil {
		err = closeErr
	}
	return local, n, err
}
