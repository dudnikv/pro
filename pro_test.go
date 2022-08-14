//

package pro

import (
	"fmt"
	"testing"
)

func TestPro(t *testing.T) {
	fmt.Println("TestPRO:", Nul, Any, All)

	//	main1()

	var sc Scanner
	sc = &ByteScanner{true}

	src := []byte("abc")
	as := &src
	for {
		tok, tip, err := sc.Scan(as)
		fmt.Println(tok, tip, *as, err)
		if tip == LexEnd {
			break
		}
	}

}
