// cron_test.go

package cron

import (
	"fmt"
	"testing"
	"time"
	)

var times_run int

func mytask(t time.Time) {
	fmt.Printf("time = %s\n",t.String())
	times_run++
}

func Test_0001(t *testing.T) {
	fmt.Printf("Test_0001\n")
	//now := time.Now()
	NewDailyJob(ANY, ANY, 5, mytask)	// every minute at 5 second mark, print time
	time.Sleep(200 * time.Second)	// 200 means it will print at least 3 times
//	if times_run != 3 {
//		t.Errorf("print fail, but keep testing")
//	}
//	fmt.Printf("PASS\n")
}

