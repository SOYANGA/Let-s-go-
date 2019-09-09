package loop

import (
	"fmt"
	"testing"
)

func TestWhileLoop(t *testing.T) {
	n := 0
	/*while(n < 5) */
	for n < 5 {
		t.Log(n)
		n++
	}
}

/*
goto 跳转语句
*/
func TestBreakLoop(t *testing.T) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println("enter Loop")
			if i < 5 {
				//break Hear
				goto Hear
			}
		}
	}
Hear:
	fmt.Println("goto Hear")
}
