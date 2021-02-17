/*
来源: github.com/i-coder-robot/design-patterns-in-golang
     B站 https://www.bilibili.com/video/BV1GD4y1D7D3/?p=6
*/
package Command

import "fmt"

type Worker struct {
	id  int
	cmd Command
}

// workerDo 被执行的函数
func WorkerDo(n int) {
	fmt.Println("Commond mode =>", n)
}

type Command struct {
	worker   *Worker
	methon   func(n int)
	firstArg int
}

func NewCommand(w *Worker, method func(n int)) Command {
	return Command{
		worker: w,
		methon: method,
	}
}

// Execute 命令执行器
func (c *Command) Execute() {
	// 将 c.firstArg 作为函数的参数传入
	c.methon(c.firstArg)
}

// NewWorker 构建命令执行者，第二个参数作为函数参数传入
func NewWorker(id, firstArg int, cmd Command) Worker {
	cmd.firstArg = firstArg
	return Worker{
		id:  id,
		cmd: cmd,
	}
}

func (w *Worker) Do() {
	w.cmd.Execute()
}
