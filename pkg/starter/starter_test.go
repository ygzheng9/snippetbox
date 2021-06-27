package starter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	greeting := SayHello("William")
	assert.Equal(t, "Hello William. Welcome!", greeting)

	another_greeting := SayHello("asdf ghjkl")
	assert.Equal(t, "Hello asdf ghjkl. Welcome!", another_greeting)
}

func TestOddOrEven(t *testing.T) {
	assert.Equal(t, "45 is an odd number", OddOrEven(45))
	assert.Equal(t, "42 is an even number", OddOrEven(42))
	assert.Equal(t, "0 is an even number", OddOrEven(0))
	assert.Equal(t, "-45 is an odd number", OddOrEven(-45))
}

func TestOddOrEven2(t *testing.T) {
	t.Run("Check Non Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "45 is an odd number", OddOrEven(45))
		assert.Equal(t, "42 is an even number", OddOrEven(42))
		assert.Equal(t, "0 is an even number", OddOrEven(0))
	})

	t.Run("Check Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", OddOrEven(-45))
		assert.Equal(t, "-42 is an even number", OddOrEven(-42))
	})
}

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assert.Equal(t, got, want)

		// if got != want {
		// 	t.Errorf("got %d want %d given, %v", got, want, numbers)
		// }
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
