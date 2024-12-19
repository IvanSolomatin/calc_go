package calc

import (
	"container/list"
	"errors"
	"reflect"
	"strconv"
)

func Calc(expression string) (float64, error) {
    if len(expression) == 0{
		return 0, errors.New("")
	}
	var err error
	var res float64
	var buf string
	l := list.New()
	var i string
	for _, x := range expression {
        i = string(x)
		if i == "(" || i == ")" || i == "-" || i == "+" || i == "*" || i == "/" {
			res, err = strconv.ParseFloat(buf, 64)
			if err == nil {
				l.PushBack(res)
			}
			l.PushBack(i)
			buf = ""
		} else {
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
		} else if e.Value == "(" {
			stack.PushBack(e.Value)
		} else if e.Value == ")" {
			for stack.Back() != nil && stack.Back().Value != "(" {
				ans.PushBack(stack.Back().Value)
				stack.Remove(stack.Back())
			}
			if stack.Back() == nil {
				return 0, errors.New("need (")
			}
			stack.Remove(stack.Back())
		} else if e.Value == "*" || e.Value == "/" {
			for stack.Back() != nil && (stack.Back().Value == "*" || stack.Back().Value == "/") {
				ans.PushBack(stack.Back().Value)
				stack.Remove(stack.Back())
			}
			stack.PushBack(e.Value)
		} else if e.Value == "-" || e.Value == "+" {
			for stack.Back() != nil && (stack.Back().Value == "*" || stack.Back().Value == "/" || stack.Back().Value == "+" || stack.Back().Value == "-") {
				ans.PushBack(stack.Back().Value)
				stack.Remove(stack.Back())
			}
			stack.PushBack(e.Value)
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
				return 0, errors.New("wrong expression")
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
					return 0, errors.New("division by zero")
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


