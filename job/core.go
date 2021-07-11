package job

import "time"

// Run runs the handler with a given interval
func Run(interval int64, handler func()) {
	t := time.Tick(time.Duration(interval) * time.Second)

	for range t {
		handler()
	}
}
