package xchannels

import (
	"testing"
	"time"
)

func TestRobotsWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "this is a pointer of signal test"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RobotsWork()
		})
	}
}

func TestRobotsRunner_RobotsGo(t *testing.T) {
	type fields struct {
		RobotTimer1      <-chan time.Time
		RobotTimer2      <-chan time.Time
		RobotTimer3      <-chan time.Time
		RobotTimer4      <-chan time.Time
		RobotActiveLevel []int
		Uid2Timer        map[int64]*<-chan time.Time
		signChan         chan int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{name: "测试time.AfterFunc()"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RobotsRunner{
				RobotTimer1:      tt.fields.RobotTimer1,
				RobotTimer2:      tt.fields.RobotTimer2,
				RobotTimer3:      tt.fields.RobotTimer3,
				RobotTimer4:      tt.fields.RobotTimer4,
				RobotActiveLevel: tt.fields.RobotActiveLevel,
				Uid2Timer:        tt.fields.Uid2Timer,
				signChan:         tt.fields.signChan,
			}
			r.RobotsGo()
		})
	}
}
