package generics

import (
	"math"
)

func Add[T int | float64 | string](a, b T) T {
	return a + b
}

func Square(a interface{}) float64 {
	switch a.(type) {
	case int:
		return math.Pow(float64(parseInt(a)), 2)
	case float64:
		return math.Pow(parseFloat(a), 2)
	}
	// instead, we can return 2nd result as error and add error here to handle it better
	return 0
}

func parseInt(a interface{}) int {
	intVal, ok := a.(int)
	if !ok {
		// since it is practice returning 0 instead of showing an error
		return 0
	}
	return intVal
}
func parseFloat(a interface{}) float64 {
	floatVal, ok := a.(float64)
	if !ok {
		// since it is practice returning 0 instead of showing an error
		return 0.0
	}
	return floatVal
}
