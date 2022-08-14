//

package pro

type Ent int32

const (
	Nul = iota
	Any
	All

	EnumLex

	LexByte
	LexChar
	LexSpace
	LexId
	LexNum
	LexOp
	LexReal
	LexStr

	LexEnd

	LexErr
)
