package crypto

import (
	"testing"
	"fmt"
)

func TestGetSHA256(t *testing.T) {
	
	var hashs []string
	for i := 0; i < 10; i++ {
		hashs = append(hashs, GetSHA256())
	}

	l := len(hashs)

	for i := 0; i < l; i++ {
		a := hashs[i]
		fmt.Printf("-- Start %d = %s --\n", i, a)
		for j := i + 1; j < l; j++ {
			b := hashs[j]
			fmt.Printf("-- Compare to %d = %s\n", j, b)
			if a == b {
				t.Fatal(i,"=",a,"\n",j,"=",b,"\n")
			}
		}
	}
}