package main

import(
	"fmt"
	"time"
)

// Function to return a reference to future object that returns the result of the task provided in future.
func ReturnAFuture(task func() (Result)) *FutureTask{

	channelForExecution := make(chan Result,1)

	futureObject := FutureTask{
		success :          false,
		done    :          false,
		error   :          nil,
		result  : 				 Result{},
		interfaceChannel : channelForExecution,
	}

	go func(){
		defer close(channelForExecution)
		resultObject := task()
		channelForExecution <- resultObject
	}()
	return &futureObject
}

//Main method
func main(){

	futureInstance1:= ReturnAFuture(func() (Result){
			var res interface{}
			res="30"
			time.Sleep(4*time.Second)
			return Result{resultValue:res}
	})
	futureInstance2:= ReturnAFuture(func() (Result){
			var res interface{}
			res="40"
			time.Sleep(3*time.Second)
			return Result{resultValue:res}
	})

	fmt.Println(futureInstance2.getWithTimeout(3*time.Second))
	fmt.Println(futureInstance2.result.error)
	fmt.Println(futureInstance1.get())
}
