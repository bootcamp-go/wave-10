package calculator_test

import (
	"app/calculator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDivide(t *testing.T) {
	t.Run("case 1: success - division by non-zero number", func(t *testing.T) {
		// arrange
		limit := 1
		cl := calculator.NewDefault(limit)

		// act
		n1, n2 := 10.0, 2.0
		result, err := cl.Divide(n1, n2)

		// assert
		expectedResult := 5.0
		require.NoError(t, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("case 2: failure - no more operations available", func(t *testing.T) {
		// arrange
		limit := 0
		cl := calculator.NewDefault(limit)
		
		// act
		n1, n2 := 10.0, 2.0
		result, err := cl.Divide(n1, n2)

		// assert
		expectedError := calculator.ErrNoMoreOperationsAvailable
		expectedResult := 0.0
		require.Error(t, err)
		require.ErrorIs(t, err, expectedError)
		require.EqualError(t, err, expectedError.Error() + ": 0")
		require.Equal(t, expectedResult, result)	
	})

	t.Run("case 3: failure - division by zero", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("case 4: failure - both numbers are zero", func(t *testing.T) {
		// arrange

		// act

		// assert

	})


}