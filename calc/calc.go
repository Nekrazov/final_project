package calc

import (
 "container/list"
 "reflect"
 "strconv"
)

func Calc(expression string) (float64, error) {
 if len(expression) == 0 {
  return 0, ErrPystV
 }

 var containsValidChars bool
 for _, char := range expression {
  if ('0' <= char && char <= '9') ||
   char == '+' ||
   char == '-' ||
   char == '*' ||
   char == '/' {
   containsValidChars = true
   break
  }
 }
 if !containsValidChars {
  return 0, ErrNecorV
 }

 var err error
 var res float64
 var buf string
 l := list.New()

 for _, x := range expression {
  i := string(x)
  switch i {
  case "(", ")", "-", "+", "*", "/":
   res, err = strconv.ParseFloat(buf, 64)
   if err == nil {
    l.PushBack(res)
   }
   l.PushBack(i)
   buf = ""
  default:
   buf += i
  }
 }

 if buf != "" {
  res, err = strconv.ParseFloat(buf, 64)
  if err == nil {
   l.PushBack(res)
  }
 }

 ans := list.New()
 stack := list.New()

 for e := l.Front(); e != nil; e = e.Next() {
  xt := reflect.TypeOf(e.Value).Kind()
  if xt == reflect.Float64 {
   ans.PushBack(e.Value)
  } else {
   switch e.Value {
   case "(":
    stack.PushBack(e.Value)
   case ")":
    for stack.Back() != nil && stack.Back().Value != "(" {
     ans.PushBack(stack.Back().Value)
     stack.Remove(stack.Back())
    }
    if stack.Back() == nil {
     return 0, ErrNecorV
    }
    stack.Remove(stack.Back())
   case "*", "/":
    for stack.Back() != nil &&
     (stack.Back().Value == "*" || stack.Back().Value == "/") {
     ans.PushBack(stack.Back().Value)
     stack.Remove(stack.Back())
    }
    stack.PushBack(e.Value)
   case "+", "-":
    for stack.Back() != nil &&
     (stack.Back().Value == "*" || stack.Back().Value == "/" ||
      stack.Back().Value == "+" || stack.Back().Value == "-") {
     ans.PushBack(stack.Back().Value)
     stack.Remove(stack.Back())
    }
    stack.PushBack(e.Value)
   }
  }
 }

 for stack.Back() != nil {
  ans.PushBack(stack.Back().Value)
  stack.Remove(stack.Back())
 }

 result := list.New()

 for e := ans.Front(); e != nil; e = e.Next() {
  xt := reflect.TypeOf(e.Value).Kind()
  if xt == reflect.Float64 {
   result.PushBack(e.Value)
  } else {
   if result.Len() < 2 {
    return 0, ErrNecorV
   }
   x2 := result.Back().Value.(float64)
   result.Remove(result.Back())
   x1 := result.Back().Value.(float64)
   result.Remove(result.Back())
   var x float64

   switch e.Value {
   case "*":
    x = x1 * x2
   case "/":
    if x2 == 0 {
     return 0, ErrDelZero
    }
    x = x1 / x2
   case "+":
    x = x1 + x2
   case "-":
    x = x1 - x2
   }
   result.PushBack(x)
  }
 }

 return result.Back().Value.(float64), nil
}
