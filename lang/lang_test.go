package lang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseStringListToTimeList(t *testing.T) {
	sList := []string{"2025-06-30T16:00:00.000Z", "2025-07-31T15:59:59.000Z"}
	tList, err := ParseStringListToTimeList(sList)

	require.Equal(t, 2, len(tList))
	require.Nil(t, err)

	sList = []string{"2025-06-30T16:00:00.000Z"}
	tList, err = ParseStringListToTimeList(sList)
	require.Equal(t, 1, len(tList))
	require.Nil(t, err)

	sList = []string{}
	tList, err = ParseStringListToTimeList(sList)
	require.Equal(t, 0, len(tList))
	require.Nil(t, err)

	sList = []string{"2025-06-30T16:00:00"}
	tList, err = ParseStringListToTimeList(sList)
	require.Equal(t, 0, len(tList))
	require.Error(t, err)
}
