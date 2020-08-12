// Implements the Future Interface

package main

import(
  "time"
  "golang.org/x/net/context"
  "fmt"
)

// Result structure returned by the future task
type Result struct{
  resultValue interface{}
  error error
}

type callbackReturn interface{}

// Future task that implements the future interface
type FutureTask struct{
  success bool         // returns true on success of async task
  done bool            // returns true if the async task execution is complete
  error error          // returns errors(if any) during the execution
  result Result        // returns the result object
	interfaceChannel <-chan Result // returns a channel to wait for the result
  callbackMethod func() //contains the callback method if any
}

//FutureTask implementation of get method for future interface
func (futureTask *FutureTask) get() Result{
  if(futureTask.done){
    return futureTask.result
  }
  if(futureTask.callbackMethod!=nil){
    defer futureTask.callbackMethod()
  }
  ctx := context.Background()
  return futureTask.getWithContext(ctx)
}

//FutureTask implementation of getWithTimeout method for future interface
func (futureTask *FutureTask) getWithTimeout(timeout time.Duration) Result{
  if(futureTask.done){
    return futureTask.result
  }
  if(futureTask.callbackMethod!=nil){
    defer futureTask.callbackMethod()
  }
  ctx, cancel := context.WithTimeout(context.Background(), timeout)
  defer cancel()
  return futureTask.getWithContext(ctx)
}

// Using Context to support cancelling and closing of channel on timeout and cancel
func (futureTask *FutureTask) getWithContext(ctx context.Context) Result{
  fmt.Println("Happening....")
  select {
  case <-ctx.Done():
    futureTask.done = true
    futureTask.success = false
    futureTask.error = &TimeoutError{errorString:"Request Timeout!"}
    futureTask.result = Result{resultValue:nil,error:futureTask.error}
    return futureTask.result
  case futureTask.result = <-futureTask.interfaceChannel:
    if(futureTask.result.error!=nil){
      futureTask.done = true
      futureTask.success = false
      futureTask.error = futureTask.result.error
    }else{
      futureTask.success = true
      futureTask.done = true
      futureTask.error = nil
    }
    return futureTask.result
  }
}

// FutureTask implementation of isComplete() method for future interface
func (futureTask *FutureTask) isComplete() bool{
  if(futureTask.done){
    return true
  }else{
    return false;
  }
}

// FutureTask implementation of isCancelled() method for future interface
func (futureTask *FutureTask) isCancelled() bool{
  if(futureTask.done){
    if(futureTask.error!=nil && futureTask.error.Error()=="Cancelled Manually"){
      return true;
    }
  }
  return false;
}

// FutureTask implementation of cancel() method for future interface
func (futureTask *FutureTask) cancel(){
  if(futureTask.isComplete() || futureTask.isCancelled()){
    return;
  }
  if(futureTask.callbackMethod!=nil){
    defer futureTask.callbackMethod()
  }
  interruptionError := &InterruptError{errorString:"Cancelled Manually"}
  futureTask.done = true
  futureTask.success = false
  futureTask.error = interruptionError
  futureTask.result = Result{resultValue:nil,error:interruptionError}
  // close(futureTask.interfaceChannel)
}

// FutureTask implementation of addDoneCallback() method for future interface
func (futureTask *FutureTask) addDoneCallback(callbackMethod func()){
  futureTask.callbackMethod = callbackMethod
}

//FutureTask implementation of setException() method for future interface
func (futureTask *FutureTask) setException(error error) error{
  if(futureTask.isComplete()){
    throwError := &InvalidStateError{errorString:"The future is already complete"}
    futureTask.error = throwError
  }else{
    futureTask.done = true
    futureTask.success = false
    futureTask.error = error
  }
  return futureTask.error
}
