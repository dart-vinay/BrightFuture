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

// Future task that implements the future interface
type FutureTask struct{
  success bool
  done bool
  error error
  result Result
	interfaceChannel <-chan Result
}

//FutureTask implementation of get method for future interface
func (futureTask *FutureTask) get() Result{
  if(futureTask.done){
    return futureTask.result
  }
  ctx := context.Background()
  return futureTask.getWithContext(ctx)
}

//FutureTask implementation of getWithTimeout method for future interface
func (futureTask *FutureTask) getWithTimeout(timeout time.Duration) Result{
  if(futureTask.done){
    return futureTask.result
  }
  ctx, cancel := context.WithTimeout(context.Background(), timeout)
  defer cancel()
  return futureTask.getWithContext(ctx)
}

// Using Context to support cancelling and closing of channel on timeout and cancel
func (futureTask *FutureTask) getWithContext(ctx context.Context) Result{
  select {
  case <-ctx.Done():
    futureTask.done = true
    futureTask.success = false
    futureTask.error = ctx.Err()
    futureTask.result = Result{resultValue:nil,error:ctx.Err()}
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

func (futureTask *FutureTask) isComplete() bool{
  if(futureTask.done){
    return true
  }else{
    return false;
  }
}

func (futureTask *FutureTask) isCancelled() bool{
  if(futureTask.done){
    if(futureTask.error!=nil && futureTask.error.Error()=="Cancelled Manually"){
      return true;
    }
  }
  return false;
}
func (futureTask *FutureTask) cancel(){
  if(futureTask.isComplete() || futureTask.isCancelled()){
    return;
  }
  interruptionError := &InterruptError{errorString:"Cancelled Manually"}
  futureTask.done = true
  futureTask.success = false
  futureTask.error = interruptionError
  futureTask.result = Result{resultValue:nil,error:interruptionError}
  // close(futureTask.interfaceChannel)
}
