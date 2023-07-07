package lei_test

import (
	"errors"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trisacrypto/lei"
)

func TestRAIndex(t *testing.T) {
	testCases := []struct {
		input string
		err   error
	}{
		{"RA000001", nil},
		{"RA888888", nil},
		{"RA999999", nil},
		{"RA100001", errors.New("unknown registration authority: RA100001")},
		{"AA000001", errors.New("unknown registration authority: AA000001")},
		{"RA00009", errors.New("unknown registration authority: RA00009")},
		{"RA0000945", errors.New("unknown registration authority: RA0000945")},
		{"RA100094", errors.New("unknown registration authority: RA100094")},
	}

	for i, tc := range testCases {
		ra, err := lei.NewRA(tc.input)
		if tc.err != nil {
			require.ErrorIs(t, err, lei.ErrUnknownRA, "test case %d expected an unknown ra error")
			require.EqualError(t, err, tc.err.Error(), "test case %d expected an error", i)
			require.Equal(t, "UNKNOWN", ra.String(), "test case %d expected invalid RA", i)
		} else {
			require.Equal(t, tc.input, ra.String(), "test case %d expected a valid RA")
		}
	}
}

func TestRegistrationAuthoritiesAreSorted(t *testing.T) {
	authorities := lei.RegistrationAuthorities()
	require.True(t, sort.StringsAreSorted(authorities), "the registration authories must be sorted")
}
