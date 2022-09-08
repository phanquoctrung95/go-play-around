/**
 * @author trung
 * @created 08/09/2022
 * @project go-play-around
 */
package main

import "time"

func example() {
	var timerCh <-chan time.Time
	if timeout > 0 {
		timer := time.NewTimer(timeout)
		defer timer.Stop()
		timerCh = timer.C
	}

	// Perform the restore.
	restore := &userRestoreFuture{
		meta:   meta,
		reader: reader,
	}
	restore.init()
	select {
	case <-timer:
		return ErrEnqueueTimeout
	case <-r.shutdownCh:
		return ErrRaftShutdown
	case r.userRestoreCh <- restore:
		// If the restore is ingested then wait for it to complete.
		if err := restore.Error(); err != nil {
			return err
		}
	}
}
