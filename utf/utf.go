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

/* This module is the author's first experiment in the Go programming language.
 * As such, it may contain poorly or incorrectly written code. Until it has
 * been more carefully reviewed and refined, do not use it as an example of
 * how to write Go code. Polite suggestions for improvement are welcome.
 */

/*
Package utf decodes and encodes UTF-encoded data. Supporting various formats
including UTF-8, UTF-16BE, UTF-16LE, UTF-32BE, UTF-32LE, and MUTF-8.

UTF encodings are described in several specifications (e.g. RFC 3629, RFC 2781,
RFC 5198, ISO/IEC 10464, The Java Virtual Machine Specification (MUTF-8), etc.).
This implementation supports strict compliance with these common specifications,
as well as more flexible interpretations that appear in many real world
implementations.

Typically, decoding starts with a call to ParseBom(data []byte) (int, int) or
(if the data length is already calculated) ParseBomLen(data []byte, dataLen int)
(int, int). These check the first bytes for a Byte Order Mark (BOM) and return
the detected encoding and length of the BOM. If no BOM was detected, then they
make a fairly reliable guess as to the encoding and 0 ofr the length of the
BOM. As such, you can start parsing Unicode code points at the returned offset
and assuming the returned encoding.

Alternatively, if the encoding is already known, then
ParseBomLenEncoding(data []byte, dataLen, encoding int) int should be called
instead. This checks if the corresponding BOM is present and returns its
length, or 0 if it is not found.

After ParseBom/ParseBomLen/ParseBomLenEncoding, the encoding and length of
the BOM are known (0 if no BOM), so decoding can start at that offset within
the data. Each Unicode code point is parsed with NextRune(data []byte, offset,
encoding int) (rune, int) or (if the data length is already calculated)
NextRuneLen(data []byte, offset, dataLength, encoding int) (rune, int). These
parse one Unicode code point at the provided offset, return it as a rune, and
return the byte length of that encoded code point. If there are not enough
bytes remaining in the data to form another code point, then a rune of
NotDetected (U+00FFFD) is returned with a length of 0. If an invalid encoding
is encountered, a rune of NotDetected and a non-zero length corresponding to
the length of the error sequence is returned.

If a UTF-16 surrogate pair is detected, the pair is combined to return the
corresponding code point in the range [U+010000,U+10FFFF). This means that
for UTF-16, a length of 4 can be returned. If a flexible interpretation of
UTF-8 or UTF-32 is used, they can also detect and combine UTF-16 surrogate
pairs, meaning that a length of 6 to 8 may be returned for such a flexible
UTF-8 and a length of 8 may be returned for such a flexible UTF-32.

Encoding is done by walking through the sequence of Unicode code points and
calling Bytes(r rune, encoding int) []byte, passing in the rune, the intended
encoding, and getting back a []byte of the encoded data. Unicode code points
that cannot be encoded return a zero-length []byte.
*/
package utf

import "io"

// ReplacementCode is the code point used whenever an error is encountered.
// Typically, the Unicode Replacement Character (U+FFFD) is used for this purpose.
const ReplacementCode = '\ufffd'

// Encoding is a type that distinguishes Utf*/Mutf* constants.
type Encoding int

// Utf* constants identify a specific encoding format to use within this code.
const (
	Utf8        Encoding = iota // UTF-8
	Utf8Lax                     // UTF-8, but allows overlong encodings, paired/un-paired UTF-16 surrogates
	Utf8Wild                    // UTF-8, Utf8Lax but also allows encodings above U+10FFFF
	Mutf8                       // MUTF-8: UTF-8 except for \0x00 encoded as overlong form 0xc0 0x80
	Mutf8Sur                    // MUTF-8: Mutf8 except UTF-16 surrogates must be used for >= U+010000
	Utf16Be                     // UTF-16BE
	Utf16BeLax                  // UTF-16BE, but allows un-paired surrogates
	Utf16Le                     // UTF-16LE
	Utf16LeLax                  // UTF-16LE, but allows un-paired surrogates
	Utf32Be                     // UTF-32BE
	Utf32BeLax                  // UTF-32BE, but allows paired/un-paired UTF-16 surrogates
	Utf32BeWild                 // UTF-32BE, Utf32BeLax but also allows encodings above U+10FFFF
	Utf32Le                     // UTF-32LE
	Utf32LeLax                  // UTF-32LE, but allows paired/un-paired UTF-16 surrogates
	Utf32LeWild                 // UTF-32LE, Utf32LeLax but also allows encodings above U+10FFFF
	UtfNotUsed                  // "undefined encoding" (typically used as a placeholder)
)

//===========================================================================
// API: Utilities:
//===========================================================================

// ToStandard takes an encoding and returns the corresponding standard version.
// For example, Utf8Lax, Utf8Wild, Mutf8, and Mutf8Sur all convert to Utf8.
// Unrecognized encodings are returned unchanged.
func ToStandard(encoding Encoding) Encoding {
	switch encoding {
	case Utf8Lax, Utf8Wild, Mutf8, Mutf8Sur:
		return Utf8
	case Utf16BeLax:
		return Utf16Be
	case Utf16LeLax:
		return Utf16Le
	case Utf32BeLax, Utf32BeWild:
		return Utf32Be
	case Utf32LeLax, Utf32LeWild:
		return Utf32Le
	default:
		return encoding
	}
}

// ToStandard is method alternative of ToStandard() that is bound to the encoding.
func (encoding Encoding) ToStandard() Encoding {
	return ToStandard(encoding)
}

// ToLax takes an encoding and returns the corresponding "Lax" version.
// If there is no corresponding "Lax" version, the input encoding
// is returned.
//
// Typical usage is to take the encoding returned from ParseBom or
// ParseBomLen and filter it through this function.
func ToLax(encoding Encoding) Encoding {
	switch encoding {
	case Utf8, Utf8Wild, Mutf8, Mutf8Sur:
		return Utf8Lax
	case Utf16Be:
		return Utf16BeLax
	case Utf16Le:
		return Utf16LeLax
	case Utf32Be, Utf32BeWild:
		return Utf32BeLax
	case Utf32Le, Utf32LeWild:
		return Utf32LeLax
	default:
		return encoding
	}
}

// ToLax is method alternative of ToLax() that is bound to the encoding.
func (encoding Encoding) ToLax() Encoding {
	return ToLax(encoding)
}

// ToWild takes an encoding and returns the corresponding "Wild" version.
// If there is no corresponding "Wild" version, the input encoding
// is returned.
//
// Typical usage is to take the encoding returned from ParseBom or
// ParseBomLen and filter it through this function.
func ToWild(encoding Encoding) Encoding {
	switch encoding {
	case Utf8, Utf8Lax, Mutf8, Mutf8Sur:
		return Utf8Wild
	case Utf16Be:
		return Utf16BeLax
	case Utf16Le:
		return Utf16LeLax
	case Utf32Be, Utf32BeLax:
		return Utf32BeWild
	case Utf32Le, Utf32LeLax:
		return Utf32LeWild
	default:
		return encoding
	}
}

// ToWild is method alternative of ToWild() that is bound to the encoding.
func (encoding Encoding) ToWild() Encoding {
	return ToWild(encoding)
}

// ToMutf8 takes an encoding and returns Mutf8 if it is one of the
// other UTF-8 encodings. If it is not one of the other UTF-8
// encodings, the input encoding is returned.
//
// Typical usage is to take the encoding returned from ParseBom or
// ParseBomLen and filter it through this function.
func ToMutf8(encoding Encoding) Encoding {
	switch encoding {
	case Utf8, Utf8Lax, Utf8Wild, Mutf8Sur:
		return Mutf8
	default:
		return encoding
	}
}

// ToMutf8 is method alternative of ToMutf8() that is bound to the encoding.
func (encoding Encoding) ToMutf8() Encoding {
	return ToMutf8(encoding)
}

// ToMutf8Sur takes an encoding and returns Mutf8Sur if it is one
// of the other UTF-8 encodings. If it is not one of the other
// UTF-8 encodings, the input encoding is returned.
//
// Typical usage is to take the encoding returned from ParseBom or
// ParseBomLen and filter it through this function.
func ToMutf8Sur(encoding Encoding) Encoding {
	switch encoding {
	case Utf8, Utf8Lax, Utf8Wild, Mutf8:
		return Mutf8Sur
	default:
		return encoding
	}
}

// ToMutf8Sur is method alternative of ToMutf8Sur() that is bound to the encoding.
func (encoding Encoding) ToMutf8Sur() Encoding {
	return ToMutf8Sur(encoding)
}

//===========================================================================
// API: Byte Order Mark (BOM):
//===========================================================================

// ParseBom is like ParseBomLen, but the length of data is not included.
//
// It returns the encoding (e.g. utf.Utf8) and the byte length of the BOM (0 if none).
func ParseBom(data []byte) (Encoding, int) {
	return ParseBomLen(data, len(data))
}

// UtfParseBomLen detects the encoding and returns it along with how many leading bytes
// are used by the Byte Order Mark (BOM) (0 if no BOM found).
//
// None of the "Lax" or "Wild" encodings are returned. Also, Mutf8 and Mutf8Sur are not
// returned. To use one of these encodings, take the returned value and translate it.
// For convenience, one of the "To" functions, ToLax, ToWild, ToMutf8, and ToMutf8Sur,
// can be used for this.
//
// This function assumes that the data is either encoded in UTF-8, has a BOM, or the
// first character is in the ASCII range 01-0x7f. Also, it is assumed that the first
// 4 bytes do not include NUL (U+0000) characters. If these assumptions are false then
// the encoding is ambiguous. For example, the starting sequence, 0x00 0x0a 0x00 0x00,
// could be any of the 5 possible UTF encodings without these assumptions.
//
// It returns the encoding (e.g. utf.Utf8) and the byte length of the BOM (0 if none).
func ParseBomLen(data []byte,
	dataLen int) (Encoding, int) {
	// encoding checklist:
	//   3 bytes:
	//     ef bb bf     UTF-8 BOM
	//   4 bytes:
	//     00 00
	//           fe ff   UTF-32BE BOM
	//           00 xx   UTF-32BE
	//     ff fe
	//           00 00   UTF-32LE BOM
	//           xx xx   UTF-16LE BOM
	//     xx 00
	//           00 00   UTF-32LE
	//           xx xx   UTF-16LE (only if first byte non-zero)
	//     fe ff xx xx   UTF-16BE BOM
	//     00 xx yy yy   UTF-16BE (or UTF-32BE if valid UTF-32BE but not UTF-16BE)
	//   2 bytes:
	//     ff fe         UTF-16LE BOM
	//     xx 00         UTF-16LE
	//     fe ff         UTF-16BE BOM
	//     00 xx         UTF-16BE
	//  If the above checks fall through, UTF-8 is assumed.
	if dataLen >= 2 {
		if dataLen >= 3 {
			if data[0] == 0xef && data[1] == 0xbb && data[2] == 0xbf {
				return Utf8, 3 // 3: ef bb bf    UTF-8 BOM
			}
			if dataLen >= 4 {
				if data[0] == 0 && data[1] == 0 {
					if data[2] == 0xfe && data[3] == 0xff {
						return Utf32Be, 4 // 4: 00 00 fe ff UTF-32BE BOM
					}
					if data[2] == 0 {
						return Utf32Be, 0 // 4: 00 00 00 xx UTF-32BE
					}
				} else if data[0] == 0xff && data[1] == 0xfe {
					if data[2] == 0 && data[3] == 0 {
						return Utf32Le, 4 // 4: ff fe 00 00 UTF-32LE BOM
					}
					return Utf16Le, 2 // 4: ff fe xx xx UTF-16LE BOM
				}
				if data[1] == 0 {
					if data[2] == 0 && data[3] == 0 {
						return Utf32Le, 0 // 4: xx 00 00 00 UTF-32LE
					}
					if data[0] != 0 {
						return Utf16Le, 0 // 4: xx 00 xx xx UTF-16LE
					}
				}
				if data[0] == 0xfe && data[1] == 0xff {
					return Utf16Be, 2 // 4: fe ff xx xx UTF-16BE BOM
				}
				if data[0] == 0 {
					// 00 xx yy yy is a little more tricky; could still be UTF-32BE
					if data[1] <= 0x10 {
						if data[2] >= 0xdc && data[2] <= 0xdf ||
							data[2] == 0xff && data[3] >= 0xfe {
							return Utf32Be, 0 // 4: 00 xx xx xx UTF-32BE
						}
						return Utf16Be, 0 // 4: 00 xx xx xx UTF-16BE
					}
					return Utf16Be, 0 // 4: 00 xx xx xx UTF-16BE
				}
			}
		}
		if dataLen < 4 {
			if data[0] == 0xff && data[1] == 0xfe {
				return Utf16Le, 2 // 2: ff fe UTF-16LE BOM
			}
			if data[0] == 0xfe && data[1] == 0xff {
				return Utf16Be, 2 // 2: ff fe UTF-16BE BOM
			}
			if data[0] == 0 {
				return Utf16Be, 0 // 2: 00 xx UTF-16BE
			}
			if data[1] == 0 {
				return Utf16Le, 0 // 2: xx 00 UTF-16LE
			}
		}
	}
	return Utf8, 0
}

// ParseBomLenRules1 is like ParseBomLen, but some adjustments to the selection
// methodology that make it compliant with Version 1.2.2 of the YAML Specification.
// These changes are as follows.
//
// All patterns of the form, 0x00.0x?? but are not any of 0x00.0x00.0xFE.0xFF,
// 0x00.0x00.0x00, or 0x00.0x00.0x00.0x00, resolve to UTF16-BE. This means that some
// patterns of 0x00.0xXX.0xYY.0xZZ of which 0xXXYYZZ is a valid code point but 0x00YYZZ
// is not will still resolve to UTF-16BE in spite of being invalid as UTF-16BE but not
// as UTF-32BE. Currently, the cases where these are valid code points for UTF-32BE but
// not UTF-16BE are in the Supplementary Private Use Area-A [U+0F0000,U+0FFFFF] and
// Supplimentary Private Use Area-B [U+100000,U+10FFFF].
//
// It returns the encoding (e.g. utf.Utf8) and the byte length of the BOM (0 if none).
func ParseBomLenRules1(data []byte,
	dataLen int) (Encoding, int) {
	if dataLen >= 2 && data[0] == 0 &&
		(dataLen < 3 || data[1] != 0 || data[2] != 0) &&
		(dataLen < 4 || data[1] != 0 || data[2] != 0xfe || data[3] != 0xff) &&
		(dataLen < 4 || data[1] != 0 || data[2] != 0 || data[3] != 0) {
		return Utf16Be, 0
	}
	return ParseBomLen(data, dataLen)
}

// UtfParseBomLenEncoding returns how many leading bytes are used by the Byte Order Mark
// (BOM) for the provided encoding. If no BOM is found or the encoding is unrecognized,
// 0 is returned.
func ParseBomLenEncoding(data []byte,
	dataLen int,
	encoding Encoding) int {
	switch encoding {
	case Utf8, Utf8Lax, Utf8Wild, Mutf8, Mutf8Sur:
		if dataLen >= 3 && data[0] == 0xef && data[1] == 0xbb && data[2] == 0xbf {
			return 3
		}
	case Utf16Be, Utf16BeLax:
		if dataLen >= 2 && data[0] == 0xfe && data[1] == 0xff {
			return 2
		}
	case Utf16Le, Utf16LeLax:
		if dataLen >= 2 && data[0] == 0xff && data[1] == 0xfe {
			return 2
		}
	case Utf32Be, Utf32BeLax, Utf32BeWild:
		if dataLen >= 4 && data[0] == 0 && data[1] == 0 && data[2] == 0xfe && data[3] == 0xff {
			return 4
		}
	case Utf32Le, Utf32LeLax, Utf32LeWild:
		if dataLen >= 4 && data[0] == 0xff && data[1] == 0xfe && data[2] == 0 && data[3] == 0 {
			return 4
		}
	}
	return 0
}

// ParseBomLenEncoding is method alternative of ParseBomLenEncoding() that
// is bound to the encoding.
func (encoding Encoding) ParseBomLenEncoding(data []byte,
	dataLen int) int {
	return ParseBomLenEncoding(data, dataLen, encoding)
}

// Bom returns a Byte Order Mark (BOM) corresponding with the specified encoding.
// An unrecognized encoding returns a byte slice of 0 bytes.
func Bom(encoding Encoding) []byte {
	switch encoding {
	case Utf8, Utf8Lax, Utf8Wild, Mutf8, Mutf8Sur:
		return []byte{0xef, 0xbb, 0xbf}
	case Utf16Be, Utf16BeLax:
		return []byte{0xfe, 0xff}
	case Utf16Le, Utf16LeLax:
		return []byte{0xff, 0xfe}
	case Utf32Be, Utf32BeLax, Utf32BeWild:
		return []byte{0x00, 0x00, 0xfe, 0xff}
	case Utf32Le, Utf32LeLax, Utf32LeWild:
		return []byte{0xff, 0xfe, 0x00, 0x00}
	}
	return []byte{}
}

// Bom is method alternative of Bom() that is bound to the encoding.
func (encoding Encoding) Bom() []byte {
	return Bom(encoding)
}

//===========================================================================
// API: Decode: Encoded Bytes:
//===========================================================================

// NextRune is like NextRuneLen, but the length of data is not included.
//
// It returns the rune and the corresponding byte length from the encoded bytes.
func NextRune(data []byte,
	offset int,
	encoding Encoding) (rune, int) {
	return NextRuneLen(data, offset, len(data), encoding)
}

// NextRune is method alternative of NextRune() that is bound to the encoding.
func (encoding Encoding) NextRune(data []byte,
	offset int) (rune, int) {
	return NextRuneLen(data, offset, len(data), encoding)
}

// NextRuneLen decodes the next code point in the byte sequence at the
// provided offset.
//
// If the returned byte length is 0, then the code point ReplacementCode
// (U+00FFFD) is also returned. A 0 length means that there is insufficient
// data remaining in the buffer to decode another code point. This can be used
// as an end of stream indication. Also, if the data is being piped in as
// chunks and there is more data left, it can be used as an indicator to
// fetch more data and append it to the end of the existing data.
//
// If the returned code point is ReplacementCode (U+00FFFD), and the returned
// length is non 0, then there is an invalid sequence of bytes. The returned
// length is the length of the invalid sequence. This length depends on the
// encoding and the type of error discovered. For example, with UTF-8, a
// data format error will typically return a length of 1 byte, while errors
// that involve a correct UTF-8 "package" (e.g. overlong errors, or code
// points above U+10FFFF) will return the length of the code point.
//
// If the encoding supports UTF-16 surrogate pairs and a correct pair was
// discovered, then they are combined and returned as one code point. The
// returned length will cover both the high and low surrogates. Therefore,
// the client software must not make assumptions about the returned length
// (e.g. UTF-16 length of 2 bytes or UTF-32 length of 4 bytes) and use the
// value returned by this function instead.
//
// It returns the rune and the corresponding byte length from the encoded bytes.
func NextRuneLen(data []byte,
	offset,
	dataLength int,
	encoding Encoding) (rune, int) {
	remainingLength := dataLength - offset
	switch encoding {
	case Utf8, Utf8Lax, Utf8Wild, Mutf8, Mutf8Sur:
		for {
			r, delta, failed := utf8Decode(data,
				offset,
				remainingLength,
				encoding == Utf8Lax || encoding == Utf8Wild,
				encoding == Mutf8 || encoding == Mutf8Sur,
				encoding != Utf8 && encoding != Mutf8)
			if failed {
				return r, delta
			}
			switch encoding {
			case Utf8, Mutf8:
				return returnIfNotInvalid(r, delta, delta, false, false)
			case Utf8Lax, Utf8Wild, Mutf8Sur:
				if r >= 0x10000 && encoding == Mutf8Sur {
					return ReplacementCode, delta
				}
				if r >= 0xd800 && r < 0xdc00 {
					rAlt, deltaAlt, failed := utf8Decode(data,
						offset+delta,
						remainingLength-delta,
						encoding == Utf8Lax || encoding == Utf8Wild,
						false,
						true)
					if failed || rAlt < 0xdc00 || rAlt > 0xdfff {
						return r, delta
					}
					return (r&0x3ff)<<10 + (rAlt & 0x3ff) + 0x10000, delta + deltaAlt
				}
				if r >= 0xdc00 && r < 0xe000 {
					return r, delta
				}
				return returnIfNotInvalid(r,
					delta, delta,
					encoding == Utf8Lax,
					encoding == Utf8Wild)
			}
		}
	case Utf16Be, Utf16BeLax:
		return utf16NextRuneLen(data, offset, remainingLength, 0, 1, encoding == Utf16BeLax)
	case Utf16Le, Utf16LeLax:
		return utf16NextRuneLen(data, offset, remainingLength, 1, 0, encoding == Utf16LeLax)
	case Utf32Be, Utf32BeLax, Utf32BeWild:
		return utf32NextRuneLen(data, offset, remainingLength,
			0, 1, 2, 3,
			encoding == Utf32BeLax,
			encoding == Utf32BeWild)
	case Utf32Le, Utf32LeLax, Utf32LeWild:
		return utf32NextRuneLen(data, offset, remainingLength,
			3, 2, 1, 0,
			encoding == Utf32LeLax,
			encoding == Utf32LeWild)
	}
	return ReplacementCode, 0
}

// NextRuneLen is method alternative of NextRuneLen() that is bound to the encoding.
func (encoding Encoding) NextRuneLen(data []byte,
	offset,
	dataLength int) (rune, int) {
	return NextRuneLen(data, offset, dataLength, encoding)
}

// NewReader returns an object implementing the io.RuneReader interface that
// receives data from an io.ByteReader interface that performs UTF decodng.
//
// If a BOM is present, it will be parsed at the start. encoding may be set to
// utf.UtfNotUsed to leave it up to the BOM or default (UTF8) encoding.
func NewReader(br io.ByteReader,
	encoding Encoding) io.RuneReader {
	var rsw readerStateWrapper
	var rs readerState
	rs.byteReader = br
	rs.encoding = encoding
	rs.atStart = true
	rsw.impl = &rs
	var rr io.RuneReader = rsw
	return rr
}

// NewReader is method alternative of NewReader() that is bound to the encoding.
func (encoding Encoding) NewReader(br io.ByteReader) io.RuneReader {
	return NewReader(br, encoding)
}

//===========================================================================
// API: Encode:
//===========================================================================

// Bytes takes a code point and encoding and returns the []byte representation.
//
// The encoding will be the standard format whenever possible. Therefore,
// using a "Lax" encoding, such as Utf8Lax, will never produce a result
// different than the corresponding standard encoding, such as Utf8.
//
// If the code point is out of range (i.e. above U+10FFFF for encodings other
// than the "Wild" encodings, or above U+1FFFFF for Utf8Wild), then the
// returned []byte encodes ReplacementCode (U+00FFFD).
func Bytes(r rune,
	encoding Encoding) []byte {
	u := uint32(r) // avoids dealing with negative numbers
	switch encoding {
	case Mutf8, Mutf8Sur:
		if u == 0x00 {
			return []byte{0xc0, 0x80}
		}
		fallthrough
	case Utf8, Utf8Lax, Utf8Wild:
		if u <= 0x7f {
			return []byte{byte(u)}
		}
		if u <= 0x7ff {
			return []byte{0xc0 + byte(u>>6), 0x80 + byte(u&0x3f)}
		}
		if u <= 0xffff {
			return []byte{0xe0 + byte(u>>12), 0x80 + byte((u>>6)&0x3f), 0x80 + byte(u&0x3f)}
		}
		if u <= 0x10ffff || (u <= 0x1fffff && encoding == Utf8Wild) {
			if encoding != Mutf8Sur {
				return []byte{0xf0 + byte(u>>18), 0x80 + byte((u>>12)&0x3f), 0x80 + byte((u>>6)&0x3f), 0x80 + byte(u&0x3f)}
			}
			s0, s1 := utf16SurrogatePair(u)
			return []byte{
				0xe0 + byte(s0>>12), 0x80 + byte((s0>>6)&0x3f), 0x80 + byte(s0&0x3f),
				0xe0 + byte(s1>>12), 0x80 + byte((s1>>6)&0x3f), 0x80 + byte(s1&0x3f),
			}
		}
	case Utf16Be, Utf16BeLax:
		if u <= 0xffff {
			return []byte{byte((u >> 8) & 0xff), byte(u & 0xff)}
		}
		if u <= 0x10ffff {
			s0, s1 := utf16SurrogatePair(u)
			return []byte{byte((s0 >> 8) & 0xff), byte(s0 & 0xff), byte((s1 >> 8) & 0xff), byte(s1 & 0xff)}
		}
	case Utf16Le, Utf16LeLax:
		if u <= 0xffff {
			return []byte{byte(u & 0xff), byte((u >> 8) & 0xff)}
		}
		if u <= 0x10ffff {
			s0, s1 := utf16SurrogatePair(u)
			return []byte{byte(s0 & 0xff), byte((s0 >> 8) & 0xff), byte(s1 & 0xff), byte((s1 >> 8) & 0xff)}
		}
	case Utf32Be, Utf32BeLax, Utf32BeWild:
		if u <= 0x10ffff || encoding == Utf32BeWild {
			return []byte{byte((u >> 24) & 0xff), byte((u >> 16) & 0xff), byte((u >> 8) & 0xff), byte(u & 0xff)}
		}
	case Utf32Le, Utf32LeLax, Utf32LeWild:
		if u <= 0x10ffff || encoding == Utf32LeWild {
			return []byte{byte(u & 0xff), byte((u >> 8) & 0xff), byte((u >> 16) & 0xff), byte((u >> 24) & 0xff)}
		}
	}
	return Bytes(ReplacementCode, encoding)
}

// Bytes is method alternative of Bytes() that is bound to the encoding.
func (encoding Encoding) Bytes(r rune) []byte {
	return Bytes(r, encoding)
}

//===========================================================================
// Private: Decode Helpers:
//===========================================================================

// utf8Decode is a private helper function to help NextRuneLen decode UTF-8.
// NextRuneLen might call this twice to decode a code point if the encoding
// supports UTF-16 surrogates (i.e. Utf8Lax, Utf8Wild, Mutf8Sur) in order to
// extract both surrogates.
//
// It returns the rune and the corresponding byte length from the encoded bytes,
// and true if it failed to decode a rune.
func utf8Decode(data []byte,
	offset,
	remainingLength int,
	laxOrWild,
	lax0,
	comb bool) (rune, int, bool) {
	if remainingLength >= 1 {
		b0 := data[offset]
		if b0 <= 0x7f {
			if b0 > 0 || laxOrWild || !lax0 {
				return rune(b0), 1, false
			}
			return ReplacementCode, 1, true
		} else if remainingLength >= 2 {
			b1 := data[offset+1]
			if b0 >= 0xc0 && b0 <= 0xdf && b1 <= 0xbf {
				if b0 > 0xc1 || laxOrWild {
					return rune(b0&0x1f)<<6 + rune(b1&0x3f), 2, false
				} else if b0 == 0xc0 && b1 == 0x80 && lax0 {
					return rune(0), 2, false
				}
				return ReplacementCode, 2, true
			} else if remainingLength >= 3 {
				b2 := data[offset+2]
				if b0 <= 0xef && b1 <= 0xbf && b2 <= 0xbf {
					if b0 > 0xe0 || b1 > 0x9f || laxOrWild {
						r := rune(b0&0x0f)<<12 +
							rune(b1&0x3f)<<6 +
							rune(b2&0x3f)
						if r < 0xd800 || r > 0xdbff || laxOrWild || comb {
							return r, 3, false
						}
						return ReplacementCode, 3, true
					}
					return ReplacementCode, 3, true
				} else if remainingLength >= 4 {
					b3 := data[offset+3]
					if b0 <= 0xf7 && b1 <= 0xbf && b2 <= 0xbf && b3 <= 0xbf {
						if b0 > 0xf0 || b1 > 0x8f || laxOrWild {
							r := rune(b0&0x07)<<18 +
								rune(b1&0x3f)<<12 +
								rune(b2&0x3f)<<6 +
								rune(b3&0x3f)
							if r < 0xd800 || r > 0xdbff || laxOrWild {
								return r, 4, false
							}
							return ReplacementCode, 4, true
						}
						return ReplacementCode, 4, true
					}
				}
			}
			return ReplacementCode, 1, true
		}
		return ReplacementCode, 1, true
	}
	return ReplacementCode, 0, true
}

// utf16NextRuneLen is a private helper function to help NextRuneLen decode
// both UTF-16BE and UTF-16LE.
//
// It returns the rune and the corresponding byte length from the encoded bytes.
func utf16NextRuneLen(data []byte,
	offset,
	remainingLength,
	ofsHi,
	ofsLo int,
	lax bool) (rune, int) {
	if remainingLength >= 2 {
		r0 := rune(data[offset+ofsHi])<<8 + rune(data[offset+ofsLo])
		if r0 < 0xd800 || r0 > 0xdfff {
			return r0, 2
		}
		if r0 < 0xdc00 && remainingLength >= 4 {
			r1 := rune(data[offset+ofsHi+2])<<8 + rune(data[offset+ofsLo+2])
			if r1 >= 0xdc00 && r1 <= 0xdfff {
				return (r0&0x3ff)<<10 + (r1 & 0x3ff) + 0x10000, 4
			}
			return ReplacementCode, 2 // invalid low surrogate
		}
		if lax {
			return r0, 2
		}
		return ReplacementCode, 2 // unexpected low surrogate (expecting high surrogate)
	}
	return ReplacementCode, 0 // less than 2 bytes
}

// utf32NextRuneLen is a private helper function to help NextRuneLen decode
// both UTF-32BE and UTF-32LE.
//
// It returns the rune and the corresponding byte length from the encoded bytes.
func utf32NextRuneLen(data []byte,
	offset,
	remainingLength,
	o3,
	o2,
	o1,
	o0 int,
	lax,
	wild bool) (rune, int) {
	if remainingLength >= 4 {
		r0 := rune(data[offset+o3])<<24 +
			rune(data[offset+o2])<<16 +
			rune(data[offset+o1])<<8 +
			rune(data[offset+o0])
		if r0 < 0xd800 || r0 > 0xdbff || (!lax && !wild) || remainingLength < 8 {
			return returnIfNotInvalid(r0, 4, 4, lax, wild)
		}
		r1 := rune(data[offset+4+o3])<<24 +
			rune(data[offset+4+o2])<<16 +
			rune(data[offset+4+o1])<<8 +
			rune(data[offset+4+o0])
		if r1 >= 0xdc00 && r1 <= 0xdfff {
			return (r0&0x3ff)<<10 + (r1 & 0x3ff) + 0x10000, 8
		}
		return returnIfNotInvalid(r0, 4, 4, lax, wild)
	}
	return ReplacementCode, 0 // less than 4 bytes
}

// returnIfNotInvalid is a private helper function for NextRuneLen to
// check the final decoded code point to ensure it is in the valid range
// before returning it. It returns the rune and corresponding byte length.
func returnIfNotInvalid(r rune,
	byteLength,
	byteLengthInvalid int,
	lax,
	wild bool) (rune, int) {
	if r < 0xd800 || (r > 0xdfff && r <= 0x10ffff) {
		return r, byteLength
	}
	if lax && r <= 0x10ffff {
		return r, byteLength
	}
	if wild {
		return r, byteLength
	}
	return ReplacementCode, byteLengthInvalid
}

//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Private: Encode Helpers: NewReader:
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

const readerStateCacheSize = 64

type readerStateWrapper struct {
	impl *readerState
}

type readerState struct {
	byteReader io.ByteReader
	encoding   Encoding
	data       []byte
	dataSize   int
	offset     int
	atStart    bool
	atEnd      bool
}

// ReadRune provides the io.RuneReader interface. See documentation for the
// io package for more details about ReadRune.
func (rsw readerStateWrapper) ReadRune() (rune, int, error) {
	var rs = rsw.impl
	if rs.dataSize-rs.offset < 4 && !rs.atEnd {
		rs.data = append(make([]byte, 0, readerStateCacheSize),
			rs.data[rs.offset:]...)
		rs.dataSize -= rs.offset
		rs.offset = 0
		for rs.dataSize < readerStateCacheSize {
			b, err := rs.byteReader.ReadByte()
			if err == nil {
				rs.data = append(rs.data, b)
				rs.dataSize++
				continue
			}
			rs.atEnd = true
			if err != io.EOF {
				return 0, rs.dataSize - rs.offset, err
			}
			break
		}
	}
	if rs.atStart {
		if rs.encoding == UtfNotUsed {
			encoding, bomLen := ParseBomLen(rs.data, rs.dataSize)
			rs.encoding = encoding
			rs.offset += bomLen
		} else {
			bomLen := ParseBomLenEncoding(rs.data, rs.dataSize, rs.encoding)
			rs.offset += bomLen
		}
		rs.atStart = false
		// When rs.atStart, we have just filled the cache to readerStateCacheSize
		// or max input data. Since the BOM is vastly smaller than
		// readerStateCacheSize (it only needs to be at least 4 bytes smaller),
		// there is no need to re-try loading rs.data if insufficient data at
		// this moment.
	}
	if rs.dataSize <= rs.offset && rs.atEnd {
		return 0, 0, io.EOF
	}
	r, delta := NextRuneLen(rs.data, rs.offset, rs.dataSize, rs.encoding)
	rs.offset += delta
	if r <= 0x7f {
		delta = 1
	} else if r <= 0x7ff {
		delta = 2
	} else if r <= 0xffff {
		delta = 3
	} else {
		delta = 4
	}
	return r, delta, nil
}

//===========================================================================
// Private: Encode Helpers:
//===========================================================================

// returnIfNotInvalid is a private helper function for Bytes to split a
// code point into UTF-16 surrogates. It assumes that the incoming code
// point is in the correct range [0x10000, 0x10FFFF] and returns the
// corresponding surrogate pair.
func utf16SurrogatePair(u uint32) (uint32, uint32) {
	u -= 0x10000
	return u>>10 + 0xd800, u&0x3ff + 0xdc00
}
