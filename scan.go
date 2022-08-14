//

package pro

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main1() {
	// An artificial input source.
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}

type Scanner interface {
	Scan(src *[]byte) (token []byte, lexType Ent, err error)
}

//type scanner struct
type ByteScanner struct {
	EndError bool
}

func (s *ByteScanner) scanEnd() (token []byte, lexType Ent, err error) {
	token = nil
	lexType = LexEnd
	if s.EndError {
		err = fmt.Errorf("Scan: EndStream")
	}
	return
}

func (s *ByteScanner) Scan(src *[]byte) (token []byte, lexType Ent, err error) {
	if len(*src) == 0 {
		return s.scanEnd()
	}
	token = (*src)[0:1]
	*src = (*src)[1:]
	lexType = LexByte
	return
}
