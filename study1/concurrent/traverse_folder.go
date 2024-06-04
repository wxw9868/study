package concurrent

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
)

type GoTask struct {
	path string
	f    func(path string)
}

func NewGoTask(f func(path string), path string) *GoTask {
	return &GoTask{
		path: path,
		f:    f,
	}
}

func (t *GoTask) execute() {
	t.f(t.path)
}

type GoPool struct {
	entryChannel chan *GoTask
	num          int
	jobChannel   chan *GoTask
}

func NewGoPool(cap int) *GoPool {
	return &GoPool{
		entryChannel: make(chan *GoTask),
		num:          cap,
		jobChannel:   make(chan *GoTask),
	}
}

func (p *GoPool) worker() {
	for task := range p.jobChannel {
		task.execute()
	}
}

func (p *GoPool) Run() {
	for i := 0; i < p.num; i++ {
		go p.worker()
	}

	for task := range p.entryChannel {
		p.jobChannel <- task
	}

	close(p.jobChannel)

	close(p.entryChannel)
}

var folderChan = make(chan string)

func folders(folder string) {
	fis, _ := ioutil.ReadDir(folder)

	for _, v := range fis {
		path := folder + "/" + v.Name()
		if v.IsDir() {
			folderChan <- path
		}
	}
}

var pool *GoPool

func loadForWait() {
	for {
		select {
		case path, ok := <-folderChan:
			if ok {
				task := NewGoTask(folders, path)
				pool.entryChannel <- task
			}
		}
	}
}

func RunFolder() {
	pool = NewGoPool(runtime.NumCPU())

	go func() {
		task := NewGoTask(folders, "C:\\WeGameApps")
		pool.entryChannel <- task
	}()

	pool.Run()

	loadForWait()
}

func TraverseFolder(folder string) {

}

var paths []string

func Folder(folder string) {
	fis, err := ioutil.ReadDir(folder)
	if err != nil {
		panic(err)
	}

	for _, v := range fis {
		path := filepath.Join(folder, v.Name())
		if v.IsDir() {
			paths = append(paths, path)
			Folder(path)
		}
		//fmt.Println(path)
	}
}
