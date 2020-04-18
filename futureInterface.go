// Java Version of a Future Interface in GO

package main

import(
  "time"
)

type Future interface {
	cancel() 
	get() Result
	getWithTimeout(duration time.Duration) Result
	isComplete() bool
	isCancelled() bool
}
