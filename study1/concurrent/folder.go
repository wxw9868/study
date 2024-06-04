package concurrent

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

var sema = make(chan struct{}, 50)
var done = make(chan struct{})

func dirents(dir string) []fs.DirEntry {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}

	defer func() { <-sema }()

	entries, err := os.ReadDir(dir)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries

}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()

	if cancelled() {
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fi, _ := entry.Info()
			fileSizes <- fi.Size()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes/1e9))
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func StartMain() {
	flag.Parse()

	roots := flag.Args()

	var tick <-chan time.Time

	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	var nfiles, nbytes int64

	var n sync.WaitGroup

	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

loop:
	for {
		select {
		case <-done:
			for range fileSizes {
				//
			}
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}

			nfiles++
			nbytes += size

		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
}
