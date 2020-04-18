# BrightFuture

## About

Bright Future contains a JAVA 8 implementation of future interface and FutureTask in Golang. The project contains four files:

* **main.go** This contains the main method from where out focus will start
* **errors.go** This contains custom errors that we might need during execution of our future task
* **futureInterface.go** This contains the definition of futureInterface(similar to JAVA future)
* **futureTask.go** This contains the implementation of futureInterface defined in the above file.

## Installation

Install Golang

## Usage

Clone the repository and from the root directory run:
```bash
go run **.go
```
To test the behaviour of out FutureTask we can create a new FutureTask using `ReturnAFuture` function in the `main.go` file which takes a task function to be performed and returns a `FutureTask` object. We can monitor out task and related details using this object.

Go ahead and test it and let me know your thoughts over this.


## Updates

I have recently started with Golang and finding it really fascinating. Will try to keep this updated as I explore more about the power of the language. Cheers!
