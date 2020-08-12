package main

func (e *TimeoutError) Error() string{
  return e.errorString
}

func (e *InterruptError) Error() string{
  return e.errorString
}

func(e *InvalidStateError) Error() string{
  return e.errorString
}

func(e *CustomError) Error() string{
  return e.errorString
}


// Custom Error to be thrown during a timeout
type TimeoutError struct{
  errorString string
}

// Error to be thrown during contexual cancellation
type InterruptError struct{
  errorString string
}

// Error to be thrown during setting an exception to a completed future
type InvalidStateError struct{
  errorString string
}

// Error to be added to stop a working future
type CustomError struct{
  errorString string
}
