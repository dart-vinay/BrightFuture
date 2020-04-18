package main

// func (e *TimeoutError) Error() string{
//   return e.errorString
// }

func (e *InterruptError) Error() string{
  return e.errorString
}

// type TimeoutError struct{
//   errorString string
// }

type InterruptError struct{
  errorString string
}
