/**
 * @author trung
 * @created 08/09/2022
 * @project go-play-around
 */
package main

import (
	"runtime/debug"
	"sync"
)

var (
	v     string
	vOnce sync.Once
)

func version() string {
	vOnce.Do(func() {
		bi, ok := debug.ReadBuildInfo()
		if ok {
			v = bi.Main.Version
		} else {
			v = "dev"
		}
	})
	return v
}
