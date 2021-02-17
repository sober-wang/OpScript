package Command

import "testing"

func TestCommmand(t *testing.T) {
	workerSober := NewWorker(9527, 9331, NewCommand(nil, WorkerDo))

	workerSober.Do()
}
