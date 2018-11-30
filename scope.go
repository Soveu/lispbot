package main

import(
	"math"
	"errors"
	"time"

	"github.com/spy16/parens"
)

func initScope() (scope *parens.Scope){
	scope = parens.NewScope(nil)
	//basicScope(s)
	//mathScope(s)

	scope.Bind("ping", func() string{
		return "pong"
	})
	scope.Bind("pong", func() string{
		return "ping"
	})
	scope.Bind("sleep", func(t float64) error{
		if(t > 30){
			return errors.New("I cannot sleep THAT long!")
		}
		time.Sleep(time.Duration(t) * time.Second)
		return nil
	})

	scope.Bind("+", Addition)
	scope.Bind("-", Subtraction)
	scope.Bind("*", Multiplication)
	scope.Bind("/", Division)
	scope.Bind("pow", math.Pow)
	scope.Bind("sqrt", math.Sqrt)
	scope.Bind("cbrt", math.Cbrt)
	scope.Bind("sin", math.Sin)
	scope.Bind("cos", math.Cos)
	scope.Bind("abs", math.Abs)

	scope.Bind("nan", math.NaN())

	scope.Bind("e", math.E)
	scope.Bind("pi", math.Pi)
	scope.Bind("phi", math.Phi)

	return
}

//func basicScope(scope *parser.Scope){
//	scope.Bind("ping", func() string{
//		return "pong"
//	})
//	scope.Bind("pong", func() string{
//		return "ping"
//	})
//	scope.Bind("say", func(args ...interface{}) string{
//		return fmt.Sprint(args...)
//	})
//	scope.Bind("sleep", func(t float64) error{
//		if(t > 30){
//			return errors.New("I cannot sleep THAT long!")
//		}
//		time.Sleep(time.Duration(t) * time.Second)
//		return nil
//	})
//}
//
//func mathScope(scope *parens.Scope){
//	scope.Bind("+", Addition)
//	scope.Bind("-", Subtraction)
//	scope.Bind("*", Multiplication)
//	scope.Bind("/", Division)
//	scope.Bind("pow", math.Pow)
//	scope.Bind("sqrt", math.Sqrt)
//	scope.Bind("cbrt", math.Cbrt)
//	scope.Bind("sin", math.Sin)
//	scope.Bind("cos", math.Cos)
//
//	scope.Bind("e", math.E)
//	scope.Bind("pi", math.Pi)
//}

func Addition(nums... float64) (res float64){
	for _, num := range nums{
		res += num
	}
	return
}

func Subtraction(nums... float64) (res float64){
	res = nums[0]
	for _, num := range nums[1:]{
		res -= num
	}
	return
}

func Multiplication(nums... float64) (res float64){
	res = 1.0;
	for _, num := range nums{
		res = res * num
	}
	return
}

func Division(nums... float64) (res float64){
	res = nums[0]
	for _, num := range nums[1:]{
		res = res / num
	}
	return
}


