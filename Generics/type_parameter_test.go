package generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	return param
}

func TestLength(t *testing.T) {
	param := Length[string]("Hubla")

	fmt.Println(param)
}

func MultipleParams[T1 any, T2 any](param1 T1, param2 T2) (T1, T2) {
	return param1, param2
}

func TestMultipleParams(t *testing.T) {
	name, age := MultipleParams[string, int]("andhika", 24)

	assert.Equal(t, "andhika", name)
	assert.Equal(t, 24, age)
}

func ComparableParam[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	}

	return false
}

func TestComparableParam(t *testing.T) {

	result := ComparableParam[int](1, 1)
	result2 := ComparableParam[int](1, 2)

	assert.True(t, result)
	assert.True(t, result2)
}

type Age int

type Number interface {
	~int | int16 | int32 | int64 | float32 | float64
}

func Substract[T Number](num1, num2 T) T {
	return num2 - num1
}

func TestTypeSet(t *testing.T) {
	result := Substract[int](5, 10)
	result2 := Substract[float32](5, 10)
	result3 := Substract[Age](5, 10)

	assert.Equal(t, int(5), result)
	assert.Equal(t, float32(5), result2)
	// assert.Equal(t, int(5), result2)
	assert.Equal(t, Age(5), result3)

}

type Words[T any] []T

func PrintWords[T any](words Words[T]) {
	for _, value := range words {
		fmt.Println(value)
	}
}

func TestGenericType(t *testing.T) {
	words := Words[string]{"andhika", "hubla", "rido"}

	PrintWords[string](words)

}
