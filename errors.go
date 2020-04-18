package main

func (e *TimeoutError) Error() string{
  return e.errorString
}

func (e *InterruptError) Error() string{
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
