/**
 * @author trung
 * @created 08/09/2022
 * @project go-play-around
 */
package main

import "sync"

type User struct {
	lock *sync.RWMutex
	Name string
}

func doSomething(u User) {
	u.lock.RLock()
	defer u.lock.RUnlock()
	// do something with `u`
}
