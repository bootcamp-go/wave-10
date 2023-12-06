package switcher_test

import (
	"app/switcher"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwitcher(t *testing.T) {
	// arrange
	// ...

	// act
	fn := switcher.Switcher("A")

	// assert
	expectedFn := switcher.A
	require.Equal(t, reflect.ValueOf(expectedFn).Pointer(), reflect.ValueOf(fn).Pointer())
}


func TestSwitcherWithHandler(t *testing.T) {
	// arrange
	fn := switcher.Handler()

	// act
	result := fn()

	// assert
	expectedResult := 5
	require.Equal(t, expectedResult, result)
}