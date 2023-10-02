package concurrency

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

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

func direntries(dir string) []fs.DirEntry {
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
func DiskUsageMain() {
	DiskUsageV2()
}
