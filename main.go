package main

import(
	"fmt"
	"time"
)

// Function to return a reference to future object that returns the result of the task provided in future.
func ReturnAFuture(task func() (Result)) *FutureTask{

	channelForExecution := make(chan Result)

	futureObject := FutureTask{
		success :          false,
		done    :          false,
		error   :          nil,
		result  : 		   Result{},
		interfaceChannel : channelForExecution,
		callbackMethod : nil,
	}

	go func(){
		fmt.Println("Heck")
		defer close(channelForExecution)
		resultObject := task()
		channelForExecution <- resultObject
		fmt.Println("Heck3")
	}()
	return &futureObject
}

//Main method
func main(){

	futureInstance1:= ReturnAFuture(func() (Result){
			var res interface{}
			res=30+23
			time.Sleep(4*time.Second)
			return Result{resultValue:res}
	})
	futureInstance2:= ReturnAFuture(func() (Result){
			var res interface{}
			res="40"
			time.Sleep(1*time.Second)
			return Result{resultValue:res}
	})

	futureInstance1.addDoneCallback(func(){
		fmt.Println("CallBack Added")
	})
	futureInstance2.addDoneCallback(func(){
		fmt.Println("Another CallBack Added")
	})

	exception1 := futureInstance1.setException(&CustomError{errorString:"Mark the task as complete"})
	fmt.Println(exception1)

	fmt.Println(futureInstance2.getWithTimeout(4*time.Second))
	fmt.Println(futureInstance1.get())

	exception2 := futureInstance1.setException(&CustomError{errorString:"Mark the task as complete"})
	fmt.Println(exception2)
}
