package lei_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trisacrypto/lei"
)

func TestMod97(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
		err      error
	}{}

	for i, tc := range testCases {
		actual, err := lei.Mod97(tc.input)
		if tc.err != nil {
			require.ErrorIs(t, err, lei.ErrInvalidChar, "expected invalid char error for test case %d", i)
			require.EqualError(t, err, tc.err.Error(), "expected specified error for test case %d", i)
		} else {
			require.Equal(t, tc.expected, actual, "test case %d failed", i)
		}
	}
}

func TestHappyParse(t *testing.T) {
	testCases := []string{
		"2594007XIACKNMUAW223",
		"54930084UKLVMY22DS16",
		"213800WSGIIZCXF1P572",
		"5493000IBP32UQZ0KL24",
		"RILFO74KP1CM8P6PCT96",
	}

	for _, tc := range testCases {
		_, err := lei.Parse(tc)
		require.NoError(t, err, "expected valid lei to be returned")
	}

}

func TestMalformed(t *testing.T) {
	testCases := []struct {
		input string
		err   error
	}{
		{"2594007XIACKNUAW223", lei.ErrInvalidLength},
		{"2594007XIACKNUAW22334", lei.ErrInvalidLength},
		{"2594007XIACKNMUAW224", lei.ErrInvalidChecksum},
	}

	for i, tc := range testCases {
		_, err := lei.Parse(tc.input)
		require.ErrorIs(t, err, tc.err, "expected malformed error for test case %d", i)
	}
}

func TestRandom(t *testing.T) {
	made := make(map[lei.LEI]struct{})
	for i := 0; i < 100; i++ {
		made[lei.Random()] = struct{}{}
	}
	require.Len(t, made, 100, "expected 100 different, random LEIs to be generated")
}
