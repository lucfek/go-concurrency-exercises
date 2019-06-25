//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

func timeCounter(u *User, procDone *bool, killed chan<- bool) {
	for {
		time.Sleep(time.Second)
		u.TimeUsed++
		if *procDone {
			if u.TimeUsed > 10 && !u.IsPremium {
				killed <- true
			}
			killed <- false
			break
		}
	}
}
func HandleRequest(process func(), u *User) bool {
	procDone := false
	killed := make(chan bool, 1)
	go timeCounter(u, &procDone, killed)
	if u.TimeUsed > 10 && !u.IsPremium {
		return false
	}
	process()
	procDone = true
	if <-killed {
		return false
	}
	return true
}

func main() {
	RunMockServer()
}
