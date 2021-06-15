//line scanner.rl:1
package expressions

import "strconv"

//line scanner.go:9
var _expression_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 10,
	1, 11, 1, 12, 1, 13, 1, 14,
	1, 15, 1, 16, 1, 17, 1, 18,
	1, 19, 1, 20, 1, 21, 1, 22,
	1, 23, 1, 24, 1, 25, 1, 26,
	1, 27, 1, 28, 1, 29, 1, 30,
	1, 31, 2, 2, 3, 2, 2, 4,
	2, 2, 5, 2, 2, 6, 2, 2,
	7, 2, 2, 8, 2, 2, 9,
}

var _expression_key_offsets []int16 = []int16{
	0, 2, 2, 3, 4, 5, 6, 7,
	8, 9, 10, 11, 12, 13, 15, 17,
	18, 19, 20, 21, 22, 23, 24, 25,
	26, 53, 56, 57, 59, 61, 62, 64,
	67, 69, 75, 84, 85, 86, 87, 97,
	98, 109, 120, 131, 142, 153, 164, 175,
	186, 197, 208, 219, 230, 241, 252, 263,
	274, 285, 296, 307,
}

var _expression_trans_keys []byte = []byte{
	34, 92, 115, 115, 105, 103, 110, 32,
	111, 111, 112, 32, 39, 48, 57, 99,
	119, 121, 99, 108, 101, 32, 104, 101,
	110, 32, 32, 33, 34, 37, 39, 45,
	46, 60, 61, 62, 95, 97, 99, 102,
	105, 110, 111, 116, 123, 9, 13, 48,
	57, 65, 90, 98, 122, 32, 9, 13,
	61, 34, 92, 97, 108, 39, 48, 57,
	46, 48, 57, 48, 57, 46, 95, 65,
	90, 97, 122, 45, 63, 95, 48, 57,
	65, 90, 97, 122, 61, 61, 61, 45,
	58, 63, 95, 48, 57, 65, 90, 97,
	122, 58, 45, 58, 63, 95, 110, 48,
	57, 65, 90, 97, 122, 45, 58, 63,
	95, 100, 48, 57, 65, 90, 97, 122,
	45, 58, 63, 95, 111, 48, 57, 65,
	90, 97, 122, 45, 58, 63, 95, 110,
	48, 57, 65, 90, 97, 122, 45, 58,
	63, 95, 116, 48, 57, 65, 90, 97,
	122, 45, 58, 63, 95, 97, 48, 57,
	65, 90, 98, 122, 45, 58, 63, 95,
	105, 48, 57, 65, 90, 97, 122, 45,
	58, 63, 95, 110, 48, 57, 65, 90,
	97, 122, 45, 58, 63, 95, 115, 48,
	57, 65, 90, 97, 122, 45, 58, 63,
	95, 97, 48, 57, 65, 90, 98, 122,
	45, 58, 63, 95, 108, 48, 57, 65,
	90, 97, 122, 45, 58, 63, 95, 115,
	48, 57, 65, 90, 97, 122, 45, 58,
	63, 95, 101, 48, 57, 65, 90, 97,
	122, 45, 58, 63, 95, 110, 48, 57,
	65, 90, 97, 122, 45, 58, 63, 95,
	105, 48, 57, 65, 90, 97, 122, 45,
	58, 63, 95, 108, 48, 57, 65, 90,
	97, 122, 45, 58, 63, 95, 114, 48,
	57, 65, 90, 97, 122, 45, 58, 63,
	95, 114, 48, 57, 65, 90, 97, 122,
	45, 58, 63, 95, 117, 48, 57, 65,
	90, 97, 122, 37,
}

var _expression_single_lengths []byte = []byte{
	2, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 0, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	19, 1, 1, 2, 2, 1, 0, 1,
	0, 2, 3, 1, 1, 1, 4, 1,
	5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 1,
}

var _expression_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	4, 1, 0, 0, 0, 0, 1, 1,
	1, 2, 3, 0, 0, 0, 3, 0,
	3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 0,
}

var _expression_index_offsets []int16 = []int16{
	0, 3, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 31,
	33, 35, 37, 39, 41, 43, 45, 47,
	49, 73, 76, 78, 81, 84, 86, 88,
	91, 93, 98, 105, 107, 109, 111, 119,
	121, 130, 139, 148, 157, 166, 175, 184,
	193, 202, 211, 220, 229, 238, 247, 256,
	265, 274, 283, 292,
}

var _expression_indicies []byte = []byte{
	2, 3, 1, 1, 4, 0, 5, 0,
	6, 0, 7, 0, 8, 0, 9, 0,
	10, 0, 11, 0, 12, 0, 13, 0,
	2, 14, 16, 15, 17, 18, 0, 19,
	0, 20, 0, 21, 0, 22, 0, 23,
	0, 24, 0, 25, 0, 26, 0, 27,
	0, 29, 30, 31, 32, 33, 34, 35,
	37, 38, 39, 40, 41, 42, 43, 44,
	45, 46, 47, 48, 29, 36, 40, 40,
	28, 29, 29, 49, 51, 50, 2, 3,
	1, 52, 53, 50, 2, 14, 36, 50,
	55, 36, 54, 16, 56, 57, 58, 58,
	58, 50, 58, 60, 58, 58, 58, 58,
	59, 61, 50, 62, 50, 63, 50, 40,
	65, 66, 40, 40, 40, 40, 64, 65,
	67, 40, 65, 66, 40, 68, 40, 40,
	40, 67, 40, 65, 66, 40, 69, 40,
	40, 40, 67, 40, 65, 66, 40, 70,
	40, 40, 40, 67, 40, 65, 66, 40,
	71, 40, 40, 40, 67, 40, 65, 66,
	40, 72, 40, 40, 40, 67, 40, 65,
	66, 40, 73, 40, 40, 40, 67, 40,
	65, 66, 40, 74, 40, 40, 40, 67,
	40, 65, 66, 40, 75, 40, 40, 40,
	67, 40, 65, 66, 40, 76, 40, 40,
	40, 67, 40, 65, 66, 40, 77, 40,
	40, 40, 67, 40, 65, 66, 40, 78,
	40, 40, 40, 67, 40, 65, 66, 40,
	79, 40, 40, 40, 67, 40, 65, 66,
	40, 80, 40, 40, 40, 67, 40, 65,
	66, 40, 81, 40, 40, 40, 67, 40,
	65, 66, 40, 82, 40, 40, 40, 67,
	40, 65, 66, 40, 83, 40, 40, 40,
	67, 40, 65, 66, 40, 84, 40, 40,
	40, 67, 40, 65, 66, 40, 85, 40,
	40, 40, 67, 40, 65, 66, 40, 79,
	40, 40, 40, 67, 86, 50,
}

var _expression_trans_targs []byte = []byte{
	24, 0, 24, 1, 3, 4, 5, 6,
	7, 24, 9, 10, 11, 24, 12, 24,
	32, 15, 20, 16, 17, 18, 19, 24,
	21, 22, 23, 24, 24, 25, 26, 27,
	28, 29, 30, 33, 31, 35, 36, 37,
	38, 40, 42, 49, 53, 54, 56, 57,
	59, 24, 24, 24, 2, 8, 24, 13,
	24, 24, 34, 24, 24, 24, 24, 24,
	24, 24, 39, 24, 41, 38, 43, 44,
	45, 46, 47, 48, 38, 50, 51, 52,
	38, 38, 55, 38, 38, 58, 14,
}

var _expression_trans_actions []byte = []byte{
	47, 0, 15, 0, 0, 0, 0, 0,
	0, 7, 0, 0, 0, 11, 0, 45,
	0, 0, 0, 0, 0, 0, 0, 9,
	0, 0, 0, 13, 31, 0, 0, 5,
	5, 5, 0, 0, 5, 0, 0, 0,
	69, 0, 0, 0, 0, 0, 0, 0,
	5, 41, 43, 19, 0, 0, 33, 0,
	35, 25, 0, 39, 29, 23, 17, 21,
	49, 27, 0, 37, 0, 57, 0, 0,
	0, 0, 0, 0, 63, 0, 0, 0,
	51, 66, 0, 54, 60, 0, 0,
}

var _expression_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0,
}

var _expression_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	3, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0,
}

var _expression_eof_trans []int16 = []int16{
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 16, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	0, 50, 51, 51, 51, 51, 51, 55,
	57, 51, 60, 51, 51, 51, 65, 68,
	68, 68, 68, 68, 68, 68, 68, 68,
	68, 68, 68, 68, 68, 68, 68, 68,
	68, 68, 68, 51,
}

const expression_start int = 24
const expression_first_final int = 24
const expression_error int = -1

const expression_en_main int = 24

//line scanner.rl:11

type lexer struct {
	parseValue
	data        []byte
	p, pe, cs   int
	ts, te, act int
}

func (l *lexer) token() string {
	return string(l.data[l.ts:l.te])
}

func newLexer(data []byte) *lexer {
	lex := &lexer{
		data: data,
		pe:   len(data),
	}

//line scanner.go:236
	{
		lex.cs = expression_start
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

//line scanner.rl:30
	return lex
}

func (lex *lexer) Lex(out *yySymType) int {
	eof := lex.pe
	tok := 0

//line scanner.go:253
	{
		var _klen int
		var _trans int
		var _acts int
		var _nacts uint
		var _keys int
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
	_resume:
		_acts = int(_expression_from_state_actions[lex.cs])
		_nacts = uint(_expression_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _expression_actions[_acts-1] {
			case 1:
//line NONE:1
				lex.ts = (lex.p)

//line scanner.go:273
			}
		}

		_keys = int(_expression_key_offsets[lex.cs])
		_trans = int(_expression_index_offsets[lex.cs])

		_klen = int(_expression_single_lengths[lex.cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + _klen - 1)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + ((_upper - _lower) >> 1)
				switch {
				case lex.data[(lex.p)] < _expression_trans_keys[_mid]:
					_upper = _mid - 1
				case lex.data[(lex.p)] > _expression_trans_keys[_mid]:
					_lower = _mid + 1
				default:
					_trans += int(_mid - int(_keys))
					goto _match
				}
			}
			_keys += _klen
			_trans += _klen
		}

		_klen = int(_expression_range_lengths[lex.cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + (_klen << 1) - 2)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + (((_upper - _lower) >> 1) & ^1)
				switch {
				case lex.data[(lex.p)] < _expression_trans_keys[_mid]:
					_upper = _mid - 2
				case lex.data[(lex.p)] > _expression_trans_keys[_mid+1]:
					_lower = _mid + 2
				default:
					_trans += int((_mid - int(_keys)) >> 1)
					goto _match
				}
			}
			_trans += _klen
		}

	_match:
		_trans = int(_expression_indicies[_trans])
	_eof_trans:
		lex.cs = int(_expression_trans_targs[_trans])

		if _expression_trans_actions[_trans] == 0 {
			goto _again
		}

		_acts = int(_expression_trans_actions[_trans])
		_nacts = uint(_expression_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _expression_actions[_acts-1] {
			case 2:
//line NONE:1
				lex.te = (lex.p) + 1

			case 3:
//line scanner.rl:38
				lex.act = 8
			case 4:
//line scanner.rl:109
				lex.act = 9
			case 5:
//line scanner.rl:116
				lex.act = 14
			case 6:
//line scanner.rl:117
				lex.act = 15
			case 7:
//line scanner.rl:118
				lex.act = 16
			case 8:
//line scanner.rl:121
				lex.act = 17
			case 9:
//line scanner.rl:43
				lex.act = 20
			case 10:
//line scanner.rl:97
				lex.te = (lex.p) + 1
				{
					tok = ASSIGN
					(lex.p)++
					goto _out
				}
			case 11:
//line scanner.rl:98
				lex.te = (lex.p) + 1
				{
					tok = CYCLE
					(lex.p)++
					goto _out
				}
			case 12:
//line scanner.rl:99
				lex.te = (lex.p) + 1
				{
					tok = LOOP
					(lex.p)++
					goto _out
				}
			case 13:
//line scanner.rl:100
				lex.te = (lex.p) + 1
				{
					tok = WHEN
					(lex.p)++
					goto _out
				}
			case 14:
//line scanner.rl:66
				lex.te = (lex.p) + 1
				{
					tok = LITERAL

					var s string
					// TODO: single-quoted cannot be escaped
					if lex.data[lex.ts] == '\'' {
						s = string(lex.data[lex.ts+1 : lex.te-1])
					} else {
						unquotedStr, err := strconv.Unquote(string(lex.data[lex.ts:lex.te]))
						if err != nil {
							panic(err)
						}
						s = unquotedStr
					}

					out.val = s
					(lex.p)++
					goto _out

				}
			case 15:
//line scanner.rl:112
				lex.te = (lex.p) + 1
				{
					tok = EQ
					(lex.p)++
					goto _out
				}
			case 16:
//line scanner.rl:113
				lex.te = (lex.p) + 1
				{
					tok = NEQ
					(lex.p)++
					goto _out
				}
			case 17:
//line scanner.rl:114
				lex.te = (lex.p) + 1
				{
					tok = GE
					(lex.p)++
					goto _out
				}
			case 18:
//line scanner.rl:115
				lex.te = (lex.p) + 1
				{
					tok = LE
					(lex.p)++
					goto _out
				}
			case 19:
//line scanner.rl:122
				lex.te = (lex.p) + 1
				{
					tok = DOTDOT
					(lex.p)++
					goto _out
				}
			case 20:
//line scanner.rl:124
				lex.te = (lex.p) + 1
				{
					tok = KEYWORD
					out.name = string(lex.data[lex.ts : lex.te-1])
					(lex.p)++
					goto _out
				}
			case 21:
//line scanner.rl:126
				lex.te = (lex.p) + 1
				{
					tok = PROPERTY
					out.name = string(lex.data[lex.ts+1 : lex.te])
					(lex.p)++
					goto _out
				}
			case 22:
//line scanner.rl:129
				lex.te = (lex.p) + 1
				{
					tok = int(lex.data[lex.ts])
					(lex.p)++
					goto _out
				}
			case 23:
//line scanner.rl:48
				lex.te = (lex.p)
				(lex.p)--
				{
					tok = LITERAL
					n, err := strconv.ParseInt(lex.token(), 10, 64)
					if err != nil {
						panic(err)
					}
					out.val = int(n)
					(lex.p)++
					goto _out

				}
			case 24:
//line scanner.rl:57
				lex.te = (lex.p)
				(lex.p)--
				{
					tok = LITERAL
					n, err := strconv.ParseFloat(lex.token(), 64)
					if err != nil {
						panic(err)
					}
					out.val = n
					(lex.p)++
					goto _out

				}
			case 25:
//line scanner.rl:43
				lex.te = (lex.p)
				(lex.p)--
				{
					tok = IDENTIFIER
					out.name = lex.token()
					(lex.p)++
					goto _out

				}
			case 26:
//line scanner.rl:126
				lex.te = (lex.p)
				(lex.p)--
				{
					tok = PROPERTY
					out.name = string(lex.data[lex.ts+1 : lex.te])
					(lex.p)++
					goto _out
				}
			case 27:
//line scanner.rl:128
				lex.te = (lex.p)
				(lex.p)--

			case 28:
//line scanner.rl:129
				lex.te = (lex.p)
				(lex.p)--
				{
					tok = int(lex.data[lex.ts])
					(lex.p)++
					goto _out
				}
			case 29:
//line scanner.rl:48
				(lex.p) = (lex.te) - 1
				{
					tok = LITERAL
					n, err := strconv.ParseInt(lex.token(), 10, 64)
					if err != nil {
						panic(err)
					}
					out.val = int(n)
					(lex.p)++
					goto _out

				}
			case 30:
//line scanner.rl:129
				(lex.p) = (lex.te) - 1
				{
					tok = int(lex.data[lex.ts])
					(lex.p)++
					goto _out
				}
			case 31:
//line NONE:1
				switch lex.act {
				case 8:
					{
						(lex.p) = (lex.te) - 1

						tok = LITERAL
						out.val = lex.token() == "true"
						(lex.p)++
						goto _out

					}
				case 9:
					{
						(lex.p) = (lex.te) - 1
						tok = LITERAL
						out.val = nil
						(lex.p)++
						goto _out
					}
				case 14:
					{
						(lex.p) = (lex.te) - 1
						tok = AND
						(lex.p)++
						goto _out
					}
				case 15:
					{
						(lex.p) = (lex.te) - 1
						tok = OR
						(lex.p)++
						goto _out
					}
				case 16:
					{
						(lex.p) = (lex.te) - 1
						tok = CONTAINS
						(lex.p)++
						goto _out
					}
				case 17:
					{
						(lex.p) = (lex.te) - 1
						tok = IN
						(lex.p)++
						goto _out
					}
				case 20:
					{
						(lex.p) = (lex.te) - 1

						tok = IDENTIFIER
						out.name = lex.token()
						(lex.p)++
						goto _out

					}
				}

//line scanner.go:564
			}
		}

	_again:
		_acts = int(_expression_to_state_actions[lex.cs])
		_nacts = uint(_expression_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _expression_actions[_acts-1] {
			case 0:
//line NONE:1
				lex.ts = 0

//line scanner.go:578
			}
		}

		(lex.p)++
		if (lex.p) != (lex.pe) {
			goto _resume
		}
	_test_eof:
		{
		}
		if (lex.p) == eof {
			if _expression_eof_trans[lex.cs] > 0 {
				_trans = int(_expression_eof_trans[lex.cs] - 1)
				goto _eof_trans
			}
		}

	_out:
		{
		}
	}

//line scanner.rl:133

	return tok
}

func (lex *lexer) Error(e string) {
	// fmt.Println("scan error:", e)
}
