package concurrency

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type dirPath struct {
	dir  string
	size int64
}

func walkdir(dir string, fileSize chan<- int64) {
	entries := direntries(dir)
	for _, entry := range entries {
		if entry.IsDir() {
			innerDir := filepath.Join(dir, entry.Name())
			walkdir(innerDir, fileSize)
			continue
		}

		fileInfo, err := entry.Info()
		if err != nil {
			fmt.Printf("reading %s , err: %s", entry.Name(), err)
			continue
		}
		fileSize <- fileInfo.Size()
	}
}
func WalkdirV2(dir string, wg *sync.WaitGroup, fileSize chan<- int64, semaphore chan struct{}) {
	defer wg.Done()
	entries := direntriesV2(dir, semaphore)
	for _, entry := range entries {
		if entry.IsDir() {
			wg.Add(1)
			innerDir := filepath.Join(dir, entry.Name())

			go WalkdirV2(innerDir, wg, fileSize, semaphore)
			continue
		}

		fileInfo, err := entry.Info()
		if err != nil {
			fmt.Printf("reading %s , err: %s", entry.Name(), err)
			continue
		}
		fileSize <- fileInfo.Size()
	}
}
func walkdirV3(root, dir string, wg *sync.WaitGroup, fileSize chan<- dirPath, semaphore chan struct{}) {
	defer wg.Done()
	entries := direntriesV2(dir, semaphore)
	for _, entry := range entries {
		if entry.IsDir() {
			wg.Add(1)
			innerDir := filepath.Join(dir, entry.Name())

			go walkdirV3(root, innerDir, wg, fileSize, semaphore)
			continue
		}

		fileInfo, err := entry.Info()
		if err != nil {
			fmt.Printf("reading %s , err: %s", entry.Name(), err)
			continue
		}
		fileSize <- dirPath{root, fileInfo.Size()}
	}
}
func direntries(dir string) []fs.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func direntriesV2(dir string, semaphore chan struct{}) []fs.DirEntry {
	semaphore <- struct{}{}
	defer func() { <-semaphore }()

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
func printDiskUsage(files, totalSize int64) {
	fmt.Printf("%d files, %.1f GB\n ", files, float64(totalSize)/1e9)
}
func printDiskUsageV2(roots map[string]int64) {
	for root, size := range roots {
		fmt.Printf("dir: %s\tsize:%.1f GB | ", root, float64(size)/1e9)
	}
	fmt.Println()
}
func DiskUsage() {
	filesizes := make(chan int64)
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		for _, root := range roots {
			walkdir(root, filesizes)
		}
		close(filesizes)
	}()

	var files, totalSize int64

	for size := range filesizes {
		files++
		totalSize += size
	}
	printDiskUsage(files, totalSize)
}

// this function prints the disk usage status
// periodically forever
func DiskUsageV2() {
	filesizes := make(chan int64)
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		for _, root := range roots {
			walkdir(root, filesizes)
		}
		close(filesizes)
	}()
	// set up
	v := flag.Bool("v", false, "show progress message")
	flag.Parse()
	var tick <-chan time.Time

	if *v {
		tick = time.NewTicker(500 * time.Microsecond).C
	}

	var files, totalSize int64
loop:
	for {
		select {
		case size, ok := <-filesizes:
			if !ok {
				break loop
			}
			files++
			totalSize += size
		case <-tick:
			printDiskUsage(files, totalSize)
		}
	}

	printDiskUsage(files, totalSize)
}

func DiskUsageV3() {

	filesizes := make(chan dirPath)
	semaphore := make(chan struct{}, 20)
	rootSize := make(map[string]int64)
	var wg sync.WaitGroup

	v := flag.Bool("v", false, "show progress message")
	flag.Parse()
	roots := flag.Args()
	fmt.Println(roots)
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		for _, root := range roots {
			wg.Add(1)
			go walkdirV3(root, root, &wg, filesizes, semaphore)
		}
	}()

	go func() {
		wg.Wait()
		close(filesizes)
	}()

	// set up

	var tick <-chan time.Time

	if *v {
		tick = time.NewTicker(500 * time.Millisecond).C
	}

loop:
	for {
		select {
		case dirPath, ok := <-filesizes:
			if !ok {
				break loop
			}
			rootSize[dirPath.dir] += dirPath.size
		case <-tick:
			printDiskUsageV2(rootSize)
		}
	}

	printDiskUsageV2(rootSize)
}
func DiskUsageMain() {
	DiskUsageV3()
}
