/***************************************************************************
  Copyright 2021 https://github.com user @orkvozku

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
***************************************************************************/

/***************************************************************************
  This comment block is informational only and does not modify the License
  described in other comment blocks in this source and/or elsewhere (e.g.
  such as in a LICENSE file).

  You are requested to use this software for good and not evil. If you are
  unsure if your plans for good and not evil meet this request then consult
  with any of your daughters. If you do not have any daughters old enough
  and able to communicate concepts such as approval and disapproval, then
  consult with another family member or another person in your community to
  review your plans to determine if they satisfy this request.
***************************************************************************/

/* This package is the author's first experiment in the Go programming language.
 * As such, it may contain poorly or incorrectly written code. Until it has
 * been more carefully reviewed and refined, do not use it as an example of
 * how to write Go code. Polite suggestions for improvement are welcome.
 */

/*
Package utf_test is a unit test package corresponding to package utf.

This package is intended to be used with the "go test" command. For more
information, consult https://pkg.go.dev/cmd/go#hdr-Test_packages
*/
package utf_test

import (
	"fmt"
	"github.com/orkvozku/go/utf"
	"testing"
)

//===========================================================================
// Tests: To*:
//===========================================================================

// testCaseTo defines the structure of testCasesTo entries.
type testCaseTo struct {
	standard int
	lax      int
	wild     int
	mutf8    int
	mutf8Sur int
}

// testCasesTo is a set of test cases for utf.To*(). These are simply expected
// BOM codes for the various encodings. utf.UtfNotUsed in one of the standard/
// lax/wild/mutf8/mutf8Sur fields means that the input value should pass
// through unchanged.
var testCasesTo = []testCaseTo{
	{utf.Utf8, utf.Utf8Lax, utf.Utf8Wild, utf.Mutf8, utf.Mutf8Sur},
	{utf.Utf16Be, utf.Utf16BeLax, utf.Utf16BeLax, utf.UtfNotUsed, utf.UtfNotUsed},
	{utf.Utf16Le, utf.Utf16LeLax, utf.Utf16LeLax, utf.UtfNotUsed, utf.UtfNotUsed},
	{utf.Utf32Be, utf.Utf32BeLax, utf.Utf32BeWild, utf.UtfNotUsed, utf.UtfNotUsed},
	{utf.Utf32Le, utf.Utf32LeLax, utf.Utf32LeWild, utf.UtfNotUsed, utf.UtfNotUsed},
	{utf.UtfNotUsed, utf.UtfNotUsed, utf.UtfNotUsed, utf.UtfNotUsed, utf.UtfNotUsed},
}

// TestTo walks through the test cases of testCasesTo and performs tests
// ensuring the conversions produce the correct results.
func TestTo(t *testing.T) {
	for i := 0; i < len(testCasesTo); i++ {
		tc := testCasesTo[i]
		for k := 0; k < len(encodingsAll); k++ {
			encoding := encodingsAll[k]
			txtEncoding := encToString(encoding)
			inGroup := false
			if encoding == tc.standard || encoding == tc.lax || encoding == tc.wild {
				inGroup = true
			} else if tc.mutf8 != utf.UtfNotUsed && encoding == tc.mutf8 {
				inGroup = true
			} else if tc.mutf8Sur != utf.UtfNotUsed && encoding == tc.mutf8Sur {
				inGroup = true
			}
			if inGroup {
				encodingTo := utf.ToStandard(encoding)
				if encodingTo != tc.standard {
					t.Errorf("ToStandard(%s) = %s, want %s",
						txtEncoding,
						encToString(encodingTo),
						encToString(tc.standard))
				}
				encodingTo = utf.ToLax(encoding)
				if encodingTo != tc.lax {
					t.Errorf("ToLax(%s) = %s, want %s",
						txtEncoding,
						encToString(encodingTo),
						encToString(tc.lax))
				}
				encodingTo = utf.ToWild(encoding)
				if encodingTo != tc.wild {
					t.Errorf("ToWild(%s) = %s, want %s",
						txtEncoding,
						encToString(encodingTo),
						encToString(tc.wild))
				}
				encodingTo = utf.ToMutf8(encoding)
				if tc.mutf8 != utf.UtfNotUsed {
					if encodingTo != tc.mutf8 {
						t.Errorf("ToMutf8(%s) = %s, want %s",
							txtEncoding,
							encToString(encodingTo),
							encToString(tc.mutf8))
					}
				} else {
					if encodingTo != encoding {
						t.Errorf("ToMutf8(%s) = %s, want %s",
							txtEncoding,
							encToString(encodingTo),
							txtEncoding)
					}
				}
				encodingTo = utf.ToMutf8Sur(encoding)
				if tc.mutf8Sur != utf.UtfNotUsed {
					if encodingTo != tc.mutf8Sur {
						t.Errorf("ToMutf8Sur(%s) = %s, want %s",
							txtEncoding,
							encToString(encodingTo),
							encToString(tc.mutf8Sur))
					}
				} else {
					if encodingTo != encoding {
						t.Errorf("ToMutf8Sur(%s) = %s, want %s",
							txtEncoding,
							encToString(encodingTo),
							txtEncoding)
					}
				}
			} else if encoding == utf.UtfNotUsed {
				encodingTo := utf.ToStandard(encoding)
				if encodingTo != encoding {
					t.Errorf("ToStandard(%s) = %s, want %s",
						txtEncoding,
						encToString(encodingTo),
						txtEncoding)
				}
				encodingTo = utf.ToLax(encoding)
				if encodingTo != encoding {
					t.Errorf("ToLax(%s) = %s, want %s",
						txtEncoding,
						encToString(encodingTo),
						txtEncoding)
				}
				encodingTo = utf.ToWild(encoding)
				if encodingTo != encoding {
					t.Errorf("ToWild(%s) = %s, want %s",
						txtEncoding,
						encToString(encodingTo),
						txtEncoding)
				}
				encodingTo = utf.ToMutf8(encoding)
				if encodingTo != encoding {
					t.Errorf("ToMutf8(%s) = %s, want %s",
						txtEncoding,
						encToString(encodingTo),
						txtEncoding)
				}
				encodingTo = utf.ToMutf8Sur(encoding)
				if encodingTo != encoding {
					t.Errorf("ToMutf8Sur(%s) = %s, want %s",
						txtEncoding,
						encToString(encodingTo),
						txtEncoding)
				}
			}
		}
	}
}

//===========================================================================
// Tests: BOM:
//===========================================================================

// testCaseBom defines the structure of testCasesBom entries.
type testCaseBom struct {
	bom       []byte
	encodings []int
}

// testCasesBom is a set of BOM test cases for utf.Bom(). These are simply
// expected BOM codes for the various encodings.
var testCasesBom = []testCaseBom{
	{[]byte{0xef, 0xbb, 0xbf}, []int{utf.Utf8, utf.Utf8Lax, utf.Utf8Wild, utf.Mutf8, utf.Mutf8Sur}},
	{[]byte{0xfe, 0xff}, []int{utf.Utf16Be, utf.Utf16BeLax}},
	{[]byte{0xff, 0xfe}, []int{utf.Utf16Le, utf.Utf16LeLax}},
	{[]byte{0x00, 0x00, 0xfe, 0xff}, []int{utf.Utf32Be, utf.Utf32BeLax, utf.Utf32BeWild}},
	{[]byte{0xff, 0xfe, 0x00, 0x00}, []int{utf.Utf32Le, utf.Utf32LeLax, utf.Utf32LeWild}},
	{[]byte{}, []int{utf.UtfNotUsed}},
}

// TestUftDecode walks through the test cases of testCasesDecode and performs
// various tests related to them. This includes testing if the BOM detection
// yields the expected encoding, and decoding the bytes with the various
// encodings specified for the test case to make sure that the expected output
// is produced.
func TestUtfBom(t *testing.T) {
	for i := 0; i < len(testCasesBom); i++ {
		tc := testCasesBom[i]
		for k := 0; k < len(tc.encodings); k++ {
			encoding := tc.encodings[k]
			txtEncoding := encToString(encoding)
			data := utf.Bom(encoding)
			if len(data) != len(tc.bom) {
				t.Errorf("Bom(%s) len = %d, want %d",
					txtEncoding,
					len(data),
					len(tc.bom))
			} else {
				for n := 0; n < len(data); n++ {
					if data[n] != tc.bom[n] {
						t.Errorf("Bom(%s)[%d] = 0x%02x, want 0x%02x",
							txtEncoding,
							n,
							data[n],
							tc.bom[n])
						break
					}
				}
			}
		}
	}
}

//===========================================================================
// Tests: Decode:
//===========================================================================

// testCaseDecode defines the structure of testCasesDecode entries.
type testCaseDecode struct {
	data       []byte
	bomLength  int
	encoding   int
	variations int
	out        []rune
}

// testCasesDecode is a set of decode test cases. Each test case consists
// of a a slice of bytes of data to be decoded, the length of the Byte
// Order Mark (BOM) or 0 if not present, the base test encoding, variations
// of the base test encoding, and the expected output (runes). The test
// case bytes are expected to decode using the base encoding to the sequence
// of runes.
//
// In addition to the specified base encoding, there may be alternative
// encodings that are also expected to decode to the same runes. These
// are specified by variations of the pattern, v[nY][nY][nY][nY]. The first
// [nY] is for the corresponding "Lax" encoding. The second [nY] is for
// the corresponding "Wild" encoding. The third is for the "Mutf8" encoding.
// The fourth is for the "Mutf8Sur" encoding. For example, if the base
// encoding is specified as utf.Utf8 with variations vYYnY, then this test
// case should work with encodings utf.Utf8, utf.Utf8Lax, utf.Utf8Wild, and
// utf.Mutf8Sur.
var testCasesDecode = []testCaseDecode{
	// general encoding/BOM detection: UTF-8:
	{[]byte{}, 0, utf.Utf8, vYYYY, []rune{}},
	{[]byte{0x48, 0x69}, 0, utf.Utf8, vYYYY, []rune{0x48, 0x69}},
	{[]byte{0xef, 0xbb, 0xbf}, 3, utf.Utf8, vYYYY, []rune{}},
	{[]byte{0xef, 0xbb, 0xbf, 0x48, 0x69}, 3, utf.Utf8, vYYYY, []rune{0x48, 0x69}},
	// general encoding/BOM detection: 1-2 bytes (incomplete BOM / unparsable):
	{[]byte{0xef}, 0, utf.Utf8, vYYYY, []rune{bad}},
	{[]byte{0xef, 0xbb}, 0, utf.Utf8, vYYYY, []rune{bad, bad}},
	// general encoding/BOM detection: >= 4 bytes:
	{[]byte{0x00, 0x00, 0xfe, 0xff}, 4, utf.Utf32Be, vYYnn, []rune{}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x00, 0x00, 0x00, 0x41}, 4, utf.Utf32Be, vYYnn, []rune{0x41}},
	{[]byte{0x00, 0x00, 0x00, 0x41}, 0, utf.Utf32Be, vYYnn, []rune{0x41}},
	{[]byte{0xff, 0xfe, 0x00, 0x00}, 4, utf.Utf32Le, vYYnn, []rune{}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0x41, 0x00, 0x00, 0x00}, 4, utf.Utf32Le, vYYnn, []rune{0x41}},
	{[]byte{0xff, 0xfe, 0x41, 0x00}, 2, utf.Utf16Le, vYnnn, []rune{0x41}},
	{[]byte{0xff, 0xfe, 0x64, 0x21}, 2, utf.Utf16Le, vYnnn, []rune{0x2164}},
	{[]byte{0x41, 0x00, 0x00, 0x00}, 0, utf.Utf32Le, vYYnn, []rune{0x41}},
	{[]byte{0x48, 0x00, 0x69, 0x00}, 0, utf.Utf16Le, vYnnn, []rune{0x48, 0x69}},
	{[]byte{0xfe, 0xff, 0x00, 0x41}, 2, utf.Utf16Be, vYnnn, []rune{0x41}},
	{[]byte{0xfe, 0xff, 0x21, 0x64}, 2, utf.Utf16Be, vYnnn, []rune{0x2164}},
	{[]byte{0x00, 0x0a, 0x21, 0x64}, 0, utf.Utf16Be, vYnnn, []rune{0x0a, 0x2164}},
	{[]byte{0x00, 0x01, 0xdb, 0xff}, 0, utf.Utf16Be, vnnnn, []rune{0x01, bad}},
	{[]byte{0x00, 0x01, 0xdc, 0x00}, 0, utf.Utf32Be, vYYnn, []rune{0x01dc00}},
	{[]byte{0x00, 0x01, 0xdf, 0xff}, 0, utf.Utf32Be, vYYnn, []rune{0x01dfff}},
	{[]byte{0x00, 0x01, 0xe0, 0x00}, 0, utf.Utf16Be, vYnnn, []rune{0x01, 0xe000}},
	// general encoding/BOM detection: 5-7 bytes (incomplete UTF-16/32):
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x00}, 4, utf.Utf32Be, vYYnn, []rune{bad}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x41}, 4, utf.Utf32Be, vYYnn, []rune{bad}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x00, 0x41}, 4, utf.Utf32Be, vYYnn, []rune{bad}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x00, 0x00, 0x41}, 4, utf.Utf32Be, vYYnn, []rune{bad}},
	{[]byte{0x00, 0x00, 0x00, 0x41, 0x00}, 0, utf.Utf32Be, vYYnn, []rune{0x41, bad}},
	{[]byte{0x00, 0x00, 0x00, 0x41, 0x41}, 0, utf.Utf32Be, vYYnn, []rune{0x41, bad}},
	{[]byte{0x00, 0x00, 0x00, 0x41, 0x00, 0x41}, 0, utf.Utf32Be, vYYnn, []rune{0x41, bad}},
	{[]byte{0x00, 0x00, 0x00, 0x41, 0x00, 0x00, 0x41}, 0, utf.Utf32Be, vYYnn, []rune{0x41, bad}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0x00}, 4, utf.Utf32Le, vYYnn, []rune{bad}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0x41}, 4, utf.Utf32Le, vYYnn, []rune{bad}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0x41, 0x00}, 4, utf.Utf32Le, vYYnn, []rune{bad}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0x41, 0x00, 0x00}, 4, utf.Utf32Le, vYYnn, []rune{bad}},
	{[]byte{0xff, 0xfe, 0x41, 0x00, 0x00}, 2, utf.Utf16Le, vYnnn, []rune{0x41, bad}},
	{[]byte{0xff, 0xfe, 0x41, 0x00, 0x41}, 2, utf.Utf16Le, vYnnn, []rune{0x41, bad}},
	{[]byte{0x41, 0x00, 0x00, 0x00, 0x00}, 0, utf.Utf32Le, vYYnn, []rune{0x41, bad}},
	{[]byte{0x41, 0x00, 0x00, 0x00, 0x41}, 0, utf.Utf32Le, vYYnn, []rune{0x41, bad}},
	{[]byte{0x41, 0x00, 0x00, 0x00, 0x41, 0x00}, 0, utf.Utf32Le, vYYnn, []rune{0x41, bad}},
	{[]byte{0x41, 0x00, 0x00, 0x00, 0x41, 0x00, 0x00}, 0, utf.Utf32Le, vYYnn, []rune{0x41, bad}},
	{[]byte{0x48, 0x00, 0x69, 0x00, 0x00}, 0, utf.Utf16Le, vYnnn, []rune{0x48, 0x69, bad}},
	{[]byte{0x48, 0x00, 0x69, 0x00, 0x41}, 0, utf.Utf16Le, vYnnn, []rune{0x48, 0x69, bad}},
	{[]byte{0xfe, 0xff, 0x00, 0x41, 0x00}, 2, utf.Utf16Be, vYnnn, []rune{0x41, bad}},
	{[]byte{0xfe, 0xff, 0x00, 0x41, 0x41}, 2, utf.Utf16Be, vYnnn, []rune{0x41, bad}},
	{[]byte{0x00, 0x0a, 0x21, 0x64, 0x00}, 0, utf.Utf16Be, vYnnn, []rune{0x0a, 0x2164, bad}},
	{[]byte{0x00, 0x0a, 0x21, 0x64, 0x41}, 0, utf.Utf16Be, vYnnn, []rune{0x0a, 0x2164, bad}},
	{[]byte{0x00, 0x01, 0xdc, 0x00, 0x00}, 0, utf.Utf32Be, vYYnn, []rune{0x01dc00, bad}},
	{[]byte{0x00, 0x01, 0xdc, 0x00, 0x41}, 0, utf.Utf32Be, vYYnn, []rune{0x01dc00, bad}},
	{[]byte{0x00, 0x01, 0xdc, 0x00, 0x00, 0x41}, 0, utf.Utf32Be, vYYnn, []rune{0x01dc00, bad}},
	{[]byte{0x00, 0x01, 0xdc, 0x00, 0x00, 0x00, 0x41}, 0, utf.Utf32Be, vYYnn, []rune{0x01dc00, bad}},
	{[]byte{0x00, 0x01, 0xe0, 0x00, 0x00}, 0, utf.Utf16Be, vYnnn, []rune{0x01, 0xe000, bad}},
	{[]byte{0x00, 0x01, 0xe0, 0x00, 0x41}, 0, utf.Utf16Be, vYnnn, []rune{0x01, 0xe000, bad}},
	// general encoding/BOM detection: 2 bytes:
	{[]byte{0xff, 0xfe}, 2, utf.Utf16Le, vYnnn, []rune{}},
	{[]byte{0x41, 0x00}, 0, utf.Utf16Le, vYnnn, []rune{0x41}},
	{[]byte{0xfe, 0xff}, 2, utf.Utf16Be, vYnnn, []rune{}},
	{[]byte{0x00, 0x41}, 0, utf.Utf16Be, vYnnn, []rune{0x41}},
	// general encoding/BOM detection: 3 bytes (incomplete UTF-16):
	{[]byte{0xff, 0xfe, 0x00}, 2, utf.Utf16Le, vYnnn, []rune{bad}},
	{[]byte{0x41, 0x00, 0x00}, 0, utf.Utf16Le, vYnnn, []rune{0x41, bad}},
	{[]byte{0xfe, 0xff, 0x00}, 2, utf.Utf16Be, vYnnn, []rune{bad}},
	{[]byte{0x00, 0x41, 0x00}, 0, utf.Utf16Be, vYnnn, []rune{0x41, bad}},
	{[]byte{0xff, 0xfe, 0x41}, 2, utf.Utf16Le, vYnnn, []rune{bad}},
	{[]byte{0x41, 0x00, 0x41}, 0, utf.Utf16Le, vYnnn, []rune{0x41, bad}},
	{[]byte{0xfe, 0xff, 0x41}, 2, utf.Utf16Be, vYnnn, []rune{bad}},
	{[]byte{0x00, 0x41, 0x41}, 0, utf.Utf16Be, vYnnn, []rune{0x41, bad}},
	// UTF-8 multibyte encodings:
	{[]byte{0x7f}, 0, utf.Utf8, vYYYY, []rune{0x007f}},
	{[]byte{0xef, 0xbb, 0xbf, 0xc2, 0x80, 0x41}, 3, utf.Utf8, vYYYY, []rune{0x0080, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xdf, 0xbf, 0x41}, 3, utf.Utf8, vYYYY, []rune{0x07ff, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xe0, 0xa0, 0x80, 0x41}, 3, utf.Utf8, vYYYY, []rune{0x0800, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xef, 0xbf, 0xbf, 0x41}, 3, utf.Utf8, vYYYY, []rune{0xffff, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xf0, 0x90, 0x80, 0x80, 0x41}, 3, utf.Utf8, vYYYn, []rune{0x10000, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xf4, 0x8f, 0xbf, 0xbf, 0x41}, 3, utf.Utf8, vYYYn, []rune{0x10ffff, 0x41}},
	// UTF-8 with UTF-16 surrogates:
	{[]byte{0xef, 0xbb, 0xbf, 0xed, 0xa0, 0x80, 0xed, 0xb0, 0x80, 0x41}, 3, utf.Utf8, vnnYn, []rune{bad, bad, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xed, 0xa0, 0x80, 0xed, 0xb0, 0x80, 0x41}, 3, utf.Utf8Lax, vnYnY, []rune{0x10000, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xed, 0xaf, 0xbf, 0xed, 0xbf, 0xbf, 0x41}, 3, utf.Utf8, vnnYn, []rune{bad, bad, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xed, 0xaf, 0xbf, 0xed, 0xbf, 0xbf, 0x41}, 3, utf.Utf8Lax, vnYnY, []rune{0x10ffff, 0x41}},
	// UTF-8 valid/invalid ranges:
	{[]byte{0xef, 0xbb, 0xbf, 0xed, 0x9f, 0xbf, 0x41}, 3, utf.Utf8, vYYYY, []rune{0xd7ff, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xed, 0xa0, 0x80, 0x41}, 3, utf.Utf8Lax, vnYnY, []rune{0xd800, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xed, 0xbf, 0xbf, 0x41}, 3, utf.Utf8Lax, vnYnY, []rune{0xdfff, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xee, 0x80, 0x80, 0x41}, 3, utf.Utf8, vYYYY, []rune{0xe000, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xf4, 0x8f, 0xbf, 0xbf, 0x41}, 3, utf.Utf8, vYYYn, []rune{0x10ffff, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xf4, 0x8f, 0xbf, 0xbf, 0x41}, 3, utf.Mutf8Sur, vnnnn, []rune{bad, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xf4, 0x90, 0x80, 0x80, 0x41}, 3, utf.Utf8, vYnYY, []rune{bad, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xf4, 0x90, 0x80, 0x80, 0x41}, 3, utf.Utf8Wild, vnnnn, []rune{0x110000, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xf7, 0xbf, 0xbf, 0xbf, 0x41}, 3, utf.Utf8, vYnYY, []rune{bad, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xf7, 0xbf, 0xbf, 0xbf, 0x41}, 3, utf.Utf8Wild, vnnnn, []rune{0x1fffff, 0x41}},
	// UTF-8 errors: overlong encodings (error skips entire encoded sequence):
	{[]byte{0xef, 0xbb, 0xbf, 0xc1, 0xbf, 0x41}, 3, utf.Utf8, vnnYY, []rune{bad, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xc1, 0xbf, 0x41}, 3, utf.Utf8Lax, vnYnn, []rune{0x7f, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xc2, 0x80, 0x61}, 3, utf.Utf8, vYYYY, []rune{0x80, 0x61}},
	{[]byte{0xef, 0xbb, 0xbf, 0xe0, 0x90, 0x90, 0xd0, 0x90, 0x41}, 3, utf.Utf8, vnnYY, []rune{bad, 0x410, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xe0, 0x90, 0x90, 0xd0, 0x90, 0x41}, 3, utf.Utf8Lax, vnYnn, []rune{0x410, 0x410, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xe0, 0x9f, 0xbf, 0xdf, 0xbf, 0x41}, 3, utf.Utf8, vnnYY, []rune{bad, 0x7ff, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xe0, 0x9f, 0xbf, 0xdf, 0xbf, 0x41}, 3, utf.Utf8Lax, vnYnn, []rune{0x7ff, 0x7ff, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xe0, 0xa0, 0xbf, 0x41}, 3, utf.Utf8, vYYYY, []rune{0x83f, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xe1, 0x9f, 0xbf, 0x41}, 3, utf.Utf8, vYYYY, []rune{0x17ff, 0x41}},
	// UTF-8 errors: encoding format (error skips one byte):
	{[]byte{0xef, 0xbb, 0xbf, 0xf8, 0x41}, 3, utf.Utf8, vYYYY, []rune{bad, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xc0, 0xc2, 0x80, 0x41}, 3, utf.Utf8, vYYYY, []rune{bad, 0x80, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xe0, 0xc2, 0x80, 0x41}, 3, utf.Utf8, vYYYY, []rune{bad, 0x80, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xef, 0x80, 0xc2, 0x80}, 3, utf.Utf8, vYYYY, []rune{bad, bad, 0x80}},
	{[]byte{0xef, 0xbb, 0xbf, 0xf0, 0xc2, 0x80, 0x41}, 3, utf.Utf8, vYYYY, []rune{bad, 0x80, 0x41}},
	{[]byte{0xef, 0xbb, 0xbf, 0xf7, 0x80, 0xc2, 0x80}, 3, utf.Utf8, vYYYY, []rune{bad, bad, 0x80}},
	{[]byte{0xef, 0xbb, 0xbf, 0xf7, 0x80, 0x80, 0xc2, 0x80}, 3, utf.Utf8, vYYYY, []rune{bad, bad, bad, 0x80}},
	// UTF-16 surrogates:
	{[]byte{0xfe, 0xff, 0x00, 0x41, 0xd8, 0x3d, 0xde, 0x42, 0x00, 0x61}, 2, utf.Utf16Be, vYnnn, []rune{0x41, 0x1f642, 0x61}},
	{[]byte{0xff, 0xfe, 0x41, 0x00, 0x3d, 0xd8, 0x42, 0xde, 0x61, 0x00}, 2, utf.Utf16Le, vYnnn, []rune{0x41, 0x1f642, 0x61}},
	// UTF-16 surrogates errors:
	{[]byte{0xfe, 0xff, 0x00, 0x41, 0xd8, 0x3d, 0x00, 0x61}, 2, utf.Utf16Be, vYnnn, []rune{0x41, bad, 0x61}},
	{[]byte{0xff, 0xfe, 0x41, 0x00, 0x3d, 0xd8, 0x61, 0x00}, 2, utf.Utf16Le, vYnnn, []rune{0x41, bad, 0x61}},
	{[]byte{0xfe, 0xff, 0x00, 0x41, 0xde, 0x42, 0x00, 0x61}, 2, utf.Utf16Be, vnnnn, []rune{0x41, bad, 0x61}},
	{[]byte{0xfe, 0xff, 0x00, 0x41, 0xde, 0x42, 0x00, 0x61}, 2, utf.Utf16BeLax, vnnnn, []rune{0x41, 0xde42, 0x61}},
	{[]byte{0xff, 0xfe, 0x41, 0x00, 0x42, 0xde, 0x61, 0x00}, 2, utf.Utf16Le, vnnnn, []rune{0x41, bad, 0x61}},
	{[]byte{0xff, 0xfe, 0x41, 0x00, 0x42, 0xde, 0x61, 0x00}, 2, utf.Utf16LeLax, vnnnn, []rune{0x41, 0xde42, 0x61}},
	// UTF-32BE valid/invalid ranges:
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x00, 0x10, 0xff, 0xff}, 4, utf.Utf32Be, vYYnn, []rune{0x10ffff}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x00, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x41}, 4, utf.Utf32Be, vYnnn, []rune{bad, 0x41}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x00, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x41}, 4, utf.Utf32BeWild, vnnnn, []rune{0x110000, 0x41}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x01, 0x00, 0x00, 0x41, 0x00, 0x00, 0x00, 0x41}, 4, utf.Utf32Be, vYnnn, []rune{bad, 0x41}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x01, 0x00, 0x00, 0x41, 0x00, 0x00, 0x00, 0x41}, 4, utf.Utf32BeWild, vnnnn, []rune{0x1000041, 0x41}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x00, 0x00, 0xd7, 0xff}, 4, utf.Utf32Be, vYYnn, []rune{0xd7ff}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x00, 0x00, 0xd8, 0x00, 0x00, 0x00, 0x00, 0x41}, 4, utf.Utf32Be, vnnnn, []rune{bad, 0x41}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x00, 0x00, 0xd8, 0x00, 0x00, 0x00, 0x00, 0x41}, 4, utf.Utf32BeLax, vnYnn, []rune{0xd800, 0x41}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x00, 0x00, 0xdf, 0xff, 0x00, 0x00, 0x00, 0x41}, 4, utf.Utf32Be, vnnnn, []rune{bad, 0x41}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x00, 0x00, 0xdf, 0xff, 0x00, 0x00, 0x00, 0x41}, 4, utf.Utf32BeLax, vnYnn, []rune{0xdfff, 0x41}},
	{[]byte{0x00, 0x00, 0xfe, 0xff, 0x00, 0x00, 0xe0, 0x00}, 4, utf.Utf32Be, vYYnn, []rune{0xe000}},
	// UTF-32LE valid/invalid ranges:
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0xff, 0xff, 0x10, 0x00}, 4, utf.Utf32Le, vYYnn, []rune{0x10ffff}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0x00, 0x00, 0x11, 0x00, 0x41, 0x00, 0x00, 0x00}, 4, utf.Utf32Le, vYnnn, []rune{bad, 0x41}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0x00, 0x00, 0x11, 0x00, 0x41, 0x00, 0x00, 0x00}, 4, utf.Utf32LeWild, vnnnn, []rune{0x110000, 0x41}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0x41, 0x00, 0x00, 0x01, 0x41, 0x00, 0x00, 0x00}, 4, utf.Utf32Le, vYnnn, []rune{bad, 0x41}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0x41, 0x00, 0x00, 0x01, 0x41, 0x00, 0x00, 0x00}, 4, utf.Utf32LeWild, vnnnn, []rune{0x1000041, 0x41}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0xff, 0xd7, 0x00, 0x00}, 4, utf.Utf32Le, vYYnn, []rune{0xd7ff}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0x00, 0xd8, 0x00, 0x00, 0x41, 0x00, 0x00, 0x00}, 4, utf.Utf32Le, vnnnn, []rune{bad, 0x41}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0x00, 0xd8, 0x00, 0x00, 0x41, 0x00, 0x00, 0x00}, 4, utf.Utf32LeLax, vnYnn, []rune{0xd800, 0x41}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0xff, 0xdf, 0x00, 0x00, 0x41, 0x00, 0x00, 0x00}, 4, utf.Utf32Le, vnnnn, []rune{bad, 0x41}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0xff, 0xdf, 0x00, 0x00, 0x41, 0x00, 0x00, 0x00}, 4, utf.Utf32LeLax, vnYnn, []rune{0xdfff, 0x41}},
	{[]byte{0xff, 0xfe, 0x00, 0x00, 0x00, 0xe0, 0x00, 0x00}, 4, utf.Utf32Le, vYYnn, []rune{0xe000}},
}

// TestUftDecode walks through the test cases of testCasesDecode and performs
// various tests related to them. This includes testing if the BOM detection
// yields the expected encoding, and decoding the bytes with the various
// encodings specified for the test case to make sure that the expected output
// is produced.
func TestUtfDecode(t *testing.T) {
	for i := 0; i < len(testCasesDecode); i++ {
		tc := testCasesDecode[i]
		tcStdEncoding := utf.ToStandard(tc.encoding)
		hexData := bytesToString(tc.data)
		encoding, ofs := utf.ParseBom(tc.data)
		encodingAlt, ofsAlt := utf.ParseBomLen(tc.data, len(tc.data))
		if encodingAlt != encoding || ofsAlt != ofs {
			t.Errorf("ParseBomLen(%s)/ParseBom = %s,%d, want %s,%d",
				hexData,
				encToString(encodingAlt),
				ofsAlt,
				encToString(encoding),
				ofs)
		}
		if ofs > 0 {
			bom := utf.Bom(tc.encoding)
			if len(bom) != ofs {
				t.Errorf("ParseBom(%s),Bom(%s) len = %d, want %d",
					hexData,
					encToString(tc.encoding),
					len(bom),
					ofs)
			} else {
				for k := 0; k < len(bom); k++ {
					if bom[k] != tc.data[k] {
						t.Errorf("ParseBom(%s),Bom(%s)[%d] = 0x%02x, want 0x%02x",
							hexData,
							encToString(tc.encoding),
							ofs,
							bom[k],
							tc.data[k])
						break
					}
				}
			}
		}
		if encoding != tcStdEncoding || ofs != tc.bomLength {
			t.Errorf("ParseBom(%s) = %s,%d, want %s,%d",
				hexData,
				encToString(encoding),
				ofs,
				encToString(tcStdEncoding),
				tc.bomLength)
		}
		for k := 0; k < len(encodingsAll); k++ {
			ofsExp := 0
			stdEncodingK := utf.ToStandard(encodingsAll[k])
			if ofs > 0 {
				if stdEncodingK == tcStdEncoding {
					ofsExp = ofs
				} else if stdEncodingK == utf.Utf16Le && encoding == utf.Utf32Le {
					ofsExp = 2
				}
			}
			ofsAlt := utf.ParseBomLenEncoding(tc.data,
				len(tc.data),
				encodingsAll[k])
			if ofsAlt != ofsExp {
				t.Errorf("ParseBomLenEncoding(%s,%d,%s) = %d, want %d",
					hexData,
					len(tc.data),
					encToString(encodingsAll[k]),
					ofsAlt,
					ofsExp)
			}
		}
		decodeTestEncoding(t, &tc, &hexData, ofs, tc.encoding)
		switch tc.variations {
		case vYnnn, vYnnY, vYnYn, vYnYY, vYYnn, vYYnY, vYYYn, vYYYY:
			decodeTestEncoding(t, &tc, &hexData, ofs, utf.ToLax(tc.encoding))
		}
		switch tc.variations {
		case vnYnn, vnYnY, vnYYn, vnYYY, vYYnn, vYYnY, vYYYn, vYYYY:
			decodeTestEncoding(t, &tc, &hexData, ofs, utf.ToWild(tc.encoding))
		}
		switch tc.variations {
		case vnnYn, vnnYY, vnYYn, vnYYY, vYnYn, vYnYY, vYYYn, vYYYY:
			decodeTestEncoding(t, &tc, &hexData, ofs, utf.ToMutf8(tc.encoding))
		}
		switch tc.variations {
		case vnnnY, vnnYY, vnYnY, vnYYY, vYnnY, vYnYY, vYYnY, vYYYY:
			decodeTestEncoding(t, &tc, &hexData, ofs, utf.ToMutf8Sur(tc.encoding))
		}
	}
}

// decodeTestEncoding is a utility function used by TestUftDecode to perform a
// decode test for a given encoding.
func decodeTestEncoding(t *testing.T, tc *testCaseDecode, hexData *string, ofs, encoding int) {
	for i := 0; ; i++ {
		r, delta := utf.NextRune(tc.data, ofs, encoding)
		rAlt, deltaAlt := utf.NextRuneLen(tc.data, ofs, len(tc.data), encoding)
		if rAlt != r || deltaAlt != delta {
			t.Errorf("NextRuneLen(%s,%d,%s)/NextRune = 0x%0x,%d, want 0x%0x,%d",
				*hexData,
				ofs,
				encToString(encoding),
				rAlt,
				deltaAlt,
				r,
				delta)
		}
		if r == utf.ReplacementCode && delta == 0 && i == len(tc.out) {
			break
		}
		if i >= len(tc.out) {
			t.Errorf("NextRune(%s,%d,%s) = 0x%0x,%d, want 0x%0x,0",
				*hexData,
				ofs,
				encToString(encoding),
				r,
				delta,
				utf.ReplacementCode)
			break
		}
		if r != tc.out[i] {
			t.Errorf("NextRune(%s,%d,%s)[%d] = 0x%0x,%d, want 0x%0x",
				*hexData,
				ofs,
				encToString(encoding),
				i,
				r,
				delta,
				tc.out[i])
			break
		}
		ofs += delta
	}
}

//===========================================================================
// Tests: Encode:
//===========================================================================

// testCaseEncode defines the structure of testCasesEncode entries.
type testCaseEncode struct {
	runes       []rune
	utf8        []byte
	utf8Wild    []byte
	mutf8       []byte
	mutf8Sur    []byte
	utf16Be     []byte
	utf16Le     []byte
	utf32Be     []byte
	utf32BeWild []byte
	utf32Le     []byte
	utf32LeWild []byte
}

// testCasesEncode is a set of encode test cases. Each test case consists
// of a a slice of runes of data to be encoded, along with bytes corresponding
// to the various encodings.
var testCasesEncode = []testCaseEncode{
	{[]rune{0x48, 0x69, 0x00, 0x1f642},
		[]byte{0x48, 0x69, 0x00, 0xf0, 0x9f, 0x99, 0x82},
		[]byte{0x48, 0x69, 0x00, 0xf0, 0x9f, 0x99, 0x82},
		[]byte{0x48, 0x69, 0xc0, 0x80, 0xf0, 0x9f, 0x99, 0x82},
		[]byte{0x48, 0x69, 0xc0, 0x80, 0xed, 0xa0, 0xbd, 0xed, 0xb9, 0x82},
		[]byte{0x00, 0x48, 0x00, 0x69, 0x00, 0x00, 0xd8, 0x3d, 0xde, 0x42},
		[]byte{0x48, 0x00, 0x69, 0x00, 0x00, 0x00, 0x3d, 0xd8, 0x42, 0xde},
		[]byte{0x00, 0x00, 0x00, 0x48, 0x00, 0x00, 0x00, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xf6, 0x42},
		[]byte{0x00, 0x00, 0x00, 0x48, 0x00, 0x00, 0x00, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xf6, 0x42},
		[]byte{0x48, 0x00, 0x00, 0x00, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x42, 0xf6, 0x01, 0x00},
		[]byte{0x48, 0x00, 0x00, 0x00, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x42, 0xf6, 0x01, 0x00}},
	{[]rune{0x61, 0x10ffff, 0x62, 0x110000, 0x63, 0x1fffff, 0x64, -1, 0x65},
		[]byte{0x61, 0xf4, 0x8f, 0xbf, 0xbf, 0x62, 0xef, 0xbf, 0xbd, 0x63, 0xef, 0xbf, 0xbd, 0x64, 0xef, 0xbf, 0xbd, 0x65},
		[]byte{0x61, 0xf4, 0x8f, 0xbf, 0xbf, 0x62, 0xf4, 0x90, 0x80, 0x80, 0x63, 0xf7, 0xbf, 0xbf, 0xbf, 0x64, 0xef, 0xbf, 0xbd, 0x65},
		[]byte{0x61, 0xf4, 0x8f, 0xbf, 0xbf, 0x62, 0xef, 0xbf, 0xbd, 0x63, 0xef, 0xbf, 0xbd, 0x64, 0xef, 0xbf, 0xbd, 0x65},
		[]byte{0x61, 0xed, 0xaf, 0xbf, 0xed, 0xbf, 0xbf, 0x62, 0xef, 0xbf, 0xbd, 0x63, 0xef, 0xbf, 0xbd, 0x64, 0xef, 0xbf, 0xbd, 0x65},
		[]byte{0x00, 0x61, 0xdb, 0xff, 0xdf, 0xff, 0x00, 0x62, 0xff, 0xfd, 0x00, 0x63, 0xff, 0xfd, 0x00, 0x64, 0xff, 0xfd, 0x00, 0x65},
		[]byte{0x61, 0x00, 0xff, 0xdb, 0xff, 0xdf, 0x62, 0x00, 0xfd, 0xff, 0x63, 0x00, 0xfd, 0xff, 0x64, 0x00, 0xfd, 0xff, 0x65, 0x00},
		[]byte{0x00, 0x00, 0x00, 0x61, 0x00, 0x10, 0xff, 0xff, 0x00, 0x00, 0x00, 0x62, 0x00, 0x00, 0xff, 0xfd, 0x00, 0x00,
			0x00, 0x63, 0x00, 0x00, 0xff, 0xfd, 0x00, 0x00, 0x00, 0x64, 0x00, 0x00, 0xff, 0xfd, 0x00, 0x00, 0x00, 0x65},
		[]byte{0x00, 0x00, 0x00, 0x61, 0x00, 0x10, 0xff, 0xff, 0x00, 0x00, 0x00, 0x62, 0x00, 0x11, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x63, 0x00, 0x1f, 0xff, 0xff, 0x00, 0x00, 0x00, 0x64, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x65},
		[]byte{0x61, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x00, 0x62, 0x00, 0x00, 0x00, 0xfd, 0xff, 0x00, 0x00, 0x63, 0x00,
			0x00, 0x00, 0xfd, 0xff, 0x00, 0x00, 0x64, 0x00, 0x00, 0x00, 0xfd, 0xff, 0x00, 0x00, 0x65, 0x00, 0x00, 0x00},
		[]byte{0x61, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x00, 0x62, 0x00, 0x00, 0x00, 0x00, 0x00, 0x11, 0x00, 0x63, 0x00,
			0x00, 0x00, 0xff, 0xff, 0x1f, 0x00, 0x64, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0x65, 0x00, 0x00, 0x00}},
}

// TestUtfEncode walks through the test cases of testCasesEncode and performs
// various tests related to them. This includes encoding the runes with all
// of the various encodings and ensuring that the encoded data matches the
// expected result. It also performs a corresponding decode test of the same
// data to make sure it decodes back to the original runes (with some
// exceptions for out-of-range values).
func TestUtfEncode(t *testing.T) {
	for i := 0; i < len(testCasesEncode); i++ {
		thisTestCase := testCasesEncode[i]
		txtRunes := runesToString(thisTestCase.runes)
		for k := 0; k < len(encodingsAll); k++ {
			thisEncoding := encodingsAll[k]
			if thisEncoding == utf.UtfNotUsed {
				continue
			}
			thisBytes := thisTestCase.utf8
			switch thisEncoding {
			case utf.Utf8, utf.Utf8Lax:
				thisBytes = thisTestCase.utf8
			case utf.Utf8Wild:
				thisBytes = thisTestCase.utf8Wild
			case utf.Mutf8:
				thisBytes = thisTestCase.mutf8
			case utf.Mutf8Sur:
				thisBytes = thisTestCase.mutf8Sur
			case utf.Utf16Be, utf.Utf16BeLax:
				thisBytes = thisTestCase.utf16Be
			case utf.Utf16Le, utf.Utf16LeLax:
				thisBytes = thisTestCase.utf16Le
			case utf.Utf32Be, utf.Utf32BeLax:
				thisBytes = thisTestCase.utf32Be
			case utf.Utf32BeWild:
				thisBytes = thisTestCase.utf32BeWild
			case utf.Utf32Le, utf.Utf32LeLax:
				thisBytes = thisTestCase.utf32Le
			case utf.Utf32LeWild:
				thisBytes = thisTestCase.utf32LeWild
			}
			var bytes []byte = []byte{}
			for n := 0; n < len(thisTestCase.runes); n++ {
				newBytes := utf.Bytes(thisTestCase.runes[n], thisEncoding)
				for m := 0; m < len(newBytes); m++ {
					bytes = append(bytes, newBytes[m])
				}
			}
			if len(bytes) != len(thisBytes) {
				t.Errorf("encoding{%s,%s} len = %d, want %d",
					txtRunes,
					encToString(thisEncoding),
					len(bytes),
					len(thisBytes))
				continue
			}
			for m := 0; m < len(bytes); m++ {
				if bytes[m] != thisBytes[m] {
					t.Errorf("encoding{%s,%s}[%d] = 0x%02x, want 0x%02x",
						txtRunes,
						encToString(thisEncoding),
						m,
						bytes[m],
						thisBytes[m])
					break
				}
			}
			ofs := 0
			for n := 0; ; n++ {
				r, delta := utf.NextRune(bytes, ofs, thisEncoding)
				if r == utf.ReplacementCode && delta == 0 && n == len(thisTestCase.runes) {
					break
				}
				ofs += delta
				if n >= len(thisTestCase.runes) {
					t.Errorf("encoding{%s,%s} decode  len = %d, want %d",
						txtRunes,
						encToString(thisEncoding),
						len(bytes),
						len(thisTestCase.runes))
					break
				}
				if r == utf.ReplacementCode &&
					(thisTestCase.runes[n] < 0 || thisTestCase.runes[n] > 0x10ffff) {
					if thisEncoding == utf.Utf8Wild &&
						(thisTestCase.runes[n] < 0 || thisTestCase.runes[n] > 0x1fffff) {
						// For utf.Utf8Wild, we still have to convert values above
						// 0x1fffff to utf.ReplacementCode.
						continue
					}
					if thisEncoding != utf.Utf8Wild && thisEncoding != utf.Utf32BeWild && thisEncoding != utf.Utf32LeWild {
						// We processed an out-of-range code point and we are not "Wild".
						// We correctly converted the code point to utf.ReplacementCode.
						// This is correct behavior so we can forgive this discrepency.
						continue
					}
				}
				if r != thisTestCase.runes[n] {
					t.Errorf("encoding{%s,%s} decode = 0x%0x,%d, want 0x%0x {n=%d}",
						txtRunes,
						encToString(thisEncoding),
						r,
						delta,
						thisTestCase.runes[n],
						n)
					break
				}
			}
		}
	}
}

//===========================================================================
// Private Helpers:
//===========================================================================

// bad is defined as a shortcut for utf.ReplacementCode. It is used to reduce
// syntactic clutter in our test data.
const bad = utf.ReplacementCode

// encodingsAll is a list of all valid encodings that might test.
var encodingsAll = []int{
	utf.Utf8,
	utf.Utf8Lax,
	utf.Utf8Wild,
	utf.Mutf8,
	utf.Mutf8Sur,
	utf.Utf16Be,
	utf.Utf16BeLax,
	utf.Utf16Le,
	utf.Utf16LeLax,
	utf.Utf32Be,
	utf.Utf32BeLax,
	utf.Utf32BeWild,
	utf.Utf32Le,
	utf.Utf32LeLax,
	utf.Utf32LeWild,
	utf.UtfNotUsed,
}

// v[nY][nY][nY][nY] symbols are used in test data to indicate if
// the test data applies also to other related encodings.
// The four [nY] patterns correspond to "Lax", "Wild", "Mutf8",
// and "Mutf8Sur". For example, vYYnY for a utf.Utf8 test case
// means that the same results are expected for utf.Utf8Lax,
// utf.Utf8Wild, and utf.Mutf8Sur, but not utf.Mutf8.
const (
	vnnnn = iota
	vnnnY
	vnnYn
	vnnYY
	vnYnn
	vnYnY
	vnYYn
	vnYYY
	vYnnn
	vYnnY
	vYnYn
	vYnYY
	vYYnn
	vYYnY
	vYYYn
	vYYYY
)

// bytesToString returns a compact hexadecimal string representation
// of a []byte.
func bytesToString(data []byte) string {
	var hexData string = ""
	var hexSep string = ""
	for i := 0; i < len(data); i++ {
		hexData += fmt.Sprintf("%s0x%02x", hexSep, data[i])
		hexSep = ","
	}
	return hexData
}

// runesToString returns a compact hexadecimal string representation
// of a []rune.
func runesToString(data []rune) string {
	var hexData string = ""
	var hexSep string = ""
	for i := 0; i < len(data); i++ {
		hexData += fmt.Sprintf("%s%02x", hexSep, uint32(data[i]))
		hexSep = ","
	}
	return hexData
}

// encToString returns a string representation of an encoding
// (utf.Utf*).
func encToString(encoding int) string {
	switch encoding {
	case utf.Utf8:
		return "Utf8"
	case utf.Utf8Lax:
		return "Utf8Lax"
	case utf.Utf8Wild:
		return "Utf8Wild"
	case utf.Mutf8:
		return "Mutf8"
	case utf.Mutf8Sur:
		return "Mutf8Sur"
	case utf.Utf16Be:
		return "Utf16Be"
	case utf.Utf16BeLax:
		return "Utf16BeLax"
	case utf.Utf16Le:
		return "Utf16Le"
	case utf.Utf16LeLax:
		return "Utf16LeLax"
	case utf.Utf32Be:
		return "Utf32Be"
	case utf.Utf32BeLax:
		return "Utf32BeLax"
	case utf.Utf32BeWild:
		return "Utf32BeWild"
	case utf.Utf32Le:
		return "Utf32Le"
	case utf.Utf32LeLax:
		return "Utf32LeLax"
	case utf.Utf32LeWild:
		return "Utf32LeWild"
	case utf.UtfNotUsed:
		return "UtfNotUsed"
	default:
		return "Utf????"
	}
}
