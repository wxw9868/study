package decorator

import (
	"fmt"
	"time"
)

type Decorator interface {
	Set(name string) string
}

type People struct{}

func (p *People) Set(name string) string {
	return fmt.Sprintf("姓名：%s", name)
}

type Middleware func(Decorator) Decorator

type Log struct {
	Decorator
}

func (l *Log) Set(name string) (result string) {
	result = l.Decorator.Set(name)
	fmt.Println(fmt.Sprintf("时间：%s；%s", time.Now().String(), result))
	return
}

func NewMiddleware() Middleware {
	return func(decorator Decorator) Decorator {
		return &Log{decorator}
	}
}

func Run() {
	var decorator Decorator
	decorator = &People{}

	m := NewMiddleware()
	m(decorator).Set("卫雄伟 ")
}
