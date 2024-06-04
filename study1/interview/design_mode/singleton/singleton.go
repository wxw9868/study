package singleton

import "sync"

//Singleton 是单例模式类
type Singleton struct{}

var singleton *Singleton
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}

//Once 单例模式
type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		if o.done == 0 {
			o.done = 1
			f()
		}
	}
}
