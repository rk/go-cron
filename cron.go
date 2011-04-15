/*
  This package implements a simplistic Cron Job system that calls functions and
  closures on given days/times. If you need events that fire more often than
  once a second, use time.Ticker or time.After instead. The event must be able
  to accept a *Time object (though it doesn't have to use it).

  ---------------------------------------------------------------------------------------

  Copyright 2011 Robert Kosek <thewickedflea@gmail.com>. All rights reserved.

  Redistribution and use in source and binary forms, with or without modification, are
  permitted provided that the following conditions are met:

     1. Redistributions of source code must retain the above copyright notice, this list of
        conditions and the following disclaimer.

     2. Redistributions in binary form must reproduce the above copyright notice, this list
        of conditions and the following disclaimer in the documentation and/or other materials
        provided with the distribution.

  THIS SOFTWARE IS PROVIDED BY ROBERT KOSEK ``AS IS'' AND ANY EXPRESS OR IMPLIED
  WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
  FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL ROBERT KOSEK OR
  CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
  CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
  SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
  ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
  NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
  ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

  The views and conclusions contained in the software and documentation are those of the
  authors and should not be interpreted as representing official policies, either expressed
  or implied, of Robert Kosek.
*/

package cron

import (
  . "time"
  "syscall"
)

type job struct {
  Month, Day, Weekday  int8
  Hour, Minute, Second int8
  Task                 func(*Time)
}

var jobs []job

// This function creates a new job that occurs at the given day and the given
// 24hour time. Any of the values may be -1 as an "any" match, so passing in
// a day of -1, the event occurs every day; passing in a second value of -1, the
// event will fire every second that the other parameters match.
func NewCronJob(month, day, weekday, hour, minute, second int8, task func(*Time)) {
  cj := job{month, day, weekday, hour, minute, second, task}
  jobs = append(jobs, cj)
}

// This creates a job that fires monthly at a given time on a given day.
func NewMonthlyJob(day, hour, minute, second int8, task func(*Time)) {
  NewCronJob(-1, day, -1, hour, minute, second, task)
}

// This creates a job that fires on the given day of the week and time.
func NewWeeklyJob(weekday, hour, minute, second int8, task func(*Time)) {
  NewCronJob(-1, -1, weekday, hour, minute, second, task)
}

// This creates a job that fires daily at a specified time.
func NewDailyJob(hour, minute, second int8, task func(*Time)) {
  NewCronJob(-1, -1, -1, hour, minute, second, task)
}

func (cj job) Matches(t *Time) (ok bool) {
  ok = (cj.Month == -1 || cj.Month == int8(t.Month)) &&
    (cj.Day == -1 || cj.Day == int8(t.Day)) &&
    (cj.Weekday == -1 || cj.Weekday == int8(t.Weekday)) &&
    (cj.Hour == -1 || cj.Hour == int8(t.Hour)) &&
    (cj.Minute == -1 || cj.Minute == int8(t.Minute)) &&
    (cj.Second == -1 || cj.Second == int8(t.Second))

  return
}

func processJobs() {
  for {
    now := LocalTime()
    for _, j := range jobs {
      // execute all our cron tasks asynchronously
      if j.Matches(now) {
        go j.Task(now)
      }
    }
    syscall.Sleep(1e9) // 1 second
  }
}

func init() {
  go processJobs()
}
