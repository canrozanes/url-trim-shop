package utils

import (
	"reflect"
	"testing"
)

func TestHashing(t *testing.T) {
	t.Run("properly encodes to strings and then returns to original on decoding", func(t *testing.T) {
		tesNums := []uint64{
			0, 1, 100, 10000, 1000000, 123412378,
		}
		hashes := []string{}
		for _, num := range tesNums {
			hashes = append(hashes, ToBase62(num))
		}

		decodedValues := []uint64{}
		for _, hash := range hashes {
			decodedValue, _ := FromBase62(hash)
			decodedValues = append(decodedValues, decodedValue)
		}
		if !reflect.DeepEqual(tesNums, decodedValues) {
			t.Errorf("got %v, want %v", decodedValues, tesNums)
		}
	})
}
