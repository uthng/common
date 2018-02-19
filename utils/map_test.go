package utils

import (
    "fmt"
    "testing"
    //"reflect"

    "github.com/stretchr/testify/assert"
)

func TestGetMapSortedKeys(t *testing.T) {
    map1 := map[string]int {
        "bj": 3,
        "ba": 4,
        "ac": 6,
        "cc": 7,
    }

    res1 := []string{"ac", "ba", "bj", "cc"}
    res2 := []string{"cc", "bj", "ba", "ac"}

    sortedKeys := GetMapSortedKeys(map1, false)
    fmt.Println(sortedKeys)
    assert.Equal(t, res1, sortedKeys)

    sortedKeys = GetMapSortedKeys(map1, true)
    fmt.Println(sortedKeys)
    assert.Equal(t, res2, sortedKeys)

}
