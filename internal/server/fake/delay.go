package fake

import (
	"fmt"
	"self_initializing_fake/internal/model"
	"time"
)

type Scheduler interface {
	Delay(data model.TestDouble) chan<- model.TestDouble
}
type schedule struct {
}

func (d schedule) Delay(data model.TestDouble) <- chan  model.TestDouble {

	result := make(chan model.TestDouble, 1)

	if data.Response.Latency == 0 {
		result <- data
		return result
	}

	timer := time.NewTimer(time.Millisecond * time.Duration(data.Response.Latency))

	<-timer.C
	fmt.Println("timer ended")
	result <- data
	return result
}
