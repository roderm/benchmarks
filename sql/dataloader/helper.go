package dataloader

import (
	"fmt"
	"strings"
)

func joinN(n int, value string, sep string) string {
	arr := make([]string, n)
	for i, _ := range arr {
		arr[i] = fmt.Sprintf("$%d", i+1)
	}
	return strings.Join(arr, sep)
}
