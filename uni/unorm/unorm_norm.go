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
  with any of your daughters. If you do not have any such daughters, then
  choose your nicest family member who is willing as an alternative. If you
  have no such family members, then choose the nicest person in your social
  circle who is willing as an alternative.
***************************************************************************/

package unorm

import "io"

//===========================================================================
// API: NFC:
//===========================================================================

// NextNFC is like NextNFCLen but with the dataLength determined by len(data).
func NextNFC(data []rune, offset int) ([]rune, int) {
	return NextNFCLen(data, offset, len(data), false)
}

// NextNFKDLen returns the next cluster of runes normalized to NFC, and the
// number of runes consumed from the input to produce them.
//
// The parameter, data, is a slice of code points to normalize.
//
// The parameter, offset, is the offset within data of where to normalize
// the next chunk of code points.
//
// The parameter, dataLength, is the total length, in code points, of data.
//
// The parameter, abandonAtEnd, is set to true if normalization should be
// quickly abandoned if the result may be affected by additional data. If
// this causes abandonment, then the output is an empty slice of runes with
// the length indicating to the end of the buffer. The caller might set this
// to true if it is buffering up chunks of data from a stream to be
// normalized, to indicate it should be reloaded before completing
// normalization at this step. If the caller is not buffering up data from
// a stream, then this parameter should be set false so that normalization
// is completed on the last code points in the buffer.
//
// The output is a slice of normalized runes, and the length of input data
// consumed to generate that slice.
//
// The output may be a slice from the same input array, or it may be from a
// separately allocated array. The caller should be aware of this and avoid
// modifying the data in the input buffer before the output data is consumed.
func NextNFCLen(data []rune, offset int, dataLength int, abandonAtEnd bool) ([]rune, int) {
	var r rune
	offsetPastQuickies := offset
	for {
		if offsetPastQuickies >= dataLength {
			if abandonAtEnd {
				if offsetPastQuickies > offset+1 {
					return data[offset : offsetPastQuickies-1], offsetPastQuickies - 1 - offset
				}
				return []rune{}, 1
			}
			return data[offset:offsetPastQuickies], offsetPastQuickies - offset
		}
		r = data[offsetPastQuickies]
		if r >= nfcGoodBelow &&
			((r >= hangulLBase && r < hangulLBaseCount) || (r >= hangulSBase && r < hangulSBaseCount) || !IsStarterNfc(r)) {
			break
		}
		offsetPastQuickies++
	}
	if offsetPastQuickies > offset+1 {
		return data[offset : offsetPastQuickies-1], offsetPastQuickies - 1 - offset
	}
	if dataLength-offset == 1 {
		if abandonAtEnd {
			return []rune{}, 1
		}
		if r >= hangulLBase && r < hangulLBaseCount {
			return data[offset : offset+1], 1
		}
	} else if r >= hangulLBase && r < hangulLBaseCount {
		return normComposeHangulLv(r, data, offset, dataLength, abandonAtEnd)
	} else if r >= hangulSBase && r < hangulSBaseCount {
		return normComposeHangulLvT(r, data, offset, dataLength, abandonAtEnd)
	}
	return normCommon(r, data, offset, dataLength, abandonAtEnd, true, false, HasNfcqcY)
}

// NFCReader is an io.RuneReader to io.RuneReader wrapper over NextNFCLen.
// It provides the convenience of the io.RuneReader interface. It also
// copies the input stream upon being read to eliminate possible changes to
// the output resulting from modifications to the input buffer. There is a
// potentially significant performance cost of using this interface over
// calling NextNFC/NextNFCLen directly.
//
// Due to the cache implementation, some sequences of code points that are
// technically legal but unlikely to occur in real text may not normalize
// properly. These consist of very long sequences of non-starters that
// significantly exceed the maximum length of 30 for Stream-Safe text
// described in https://unicode.org/reports/tr15/tr15-51.html.
func NFCReader(rr io.RuneReader) io.RuneReader {
	return newReader(rr, NextNFCLen)
}

//===========================================================================
// API: NFD:
//===========================================================================

// NextNFD is like NextNFDLen but with the dataLength determined by len(data).
func NextNFD(data []rune, offset int) ([]rune, int) {
	return NextNFDLen(data, offset, len(data), false)
}

// NextNFKDLen returns the next cluster of runes normalized to NFD, and the
// number of runes consumed from the input to produce them.
//
// The parameter, data, is a slice of code points to normalize.
//
// The parameter, offset, is the offset within data of where to normalize
// the next chunk of code points.
//
// The parameter, dataLength, is the total length, in code points, of data.
//
// The parameter, abandonAtEnd, is set to true if normalization should be
// quickly abandoned if the result may be affected by additional data. If
// this causes abandonment, then the output is an empty slice of runes with
// the length indicating to the end of the buffer. The caller might set this
// to true if it is buffering up chunks of data from a stream to be
// normalized, to indicate it should be reloaded before completing
// normalization at this step. If the caller is not buffering up data from
// a stream, then this parameter should be set false so that normalization
// is completed on the last code points in the buffer.
//
// The output is a slice of normalized runes, and the length of input data
// consumed to generate that slice.
//
// The output may be a slice from the same input array, or it may be from a
// separately allocated array.
func NextNFDLen(data []rune, offset int, dataLength int, abandonAtEnd bool) ([]rune, int) {
	var r rune
	offsetPastQuickies := offset
	for {
		if offsetPastQuickies >= dataLength {
			if abandonAtEnd {
				if offsetPastQuickies > offset+1 {
					return data[offset : offsetPastQuickies-1], offsetPastQuickies - 1 - offset
				}
				return []rune{}, 1
			}
			return data[offset:offsetPastQuickies], offsetPastQuickies - offset
		}
		r = data[offsetPastQuickies]
		if r >= nfdGoodBelow &&
			((r >= hangulSBase && r < hangulSBaseCount) || !IsStarterNfd(r)) {
			break
		}
		offsetPastQuickies++
	}
	if offsetPastQuickies > offset+1 {
		return data[offset : offsetPastQuickies-1], offsetPastQuickies - 1 - offset
	}
	r = data[offset]
	if dataLength-offset == 1 {
		if abandonAtEnd {
			return []rune{}, 1
		}
		if r >= hangulSBase && r < hangulSBaseCount {
			return normDecomposeHangulLvt(r)
		}
	} else if r >= hangulSBase && r < hangulSBaseCount {
		return normDecomposeHangulLvt(r)
	}
	return normCommon(r, data, offset, dataLength, abandonAtEnd, false, false, HasNfdqcY)
}

// NFDReader is an io.RuneReader to io.RuneReader wrapper over NextNFDLen.
// It provides the convenience of the io.RuneReader interface. It also
// copies the input stream upon being read to eliminate possible changes to
// the output resulting from modifications to the input buffer. There is a
// potentially significant performance cost of using this interface over
// calling NextNFD/NextNFDLen directly.
//
// Due to the cache implementation, some sequences of code points that are
// technically legal but unlikely to occur in real text may not normalize
// properly. These consist of very long sequences of non-starters that
// significantly exceed the maximum length of 30 for Stream-Safe text
// described in https://unicode.org/reports/tr15/tr15-51.html.
func NFDReader(rr io.RuneReader) io.RuneReader {
	return newReader(rr, NextNFDLen)
}

//===========================================================================
// API: NFKC:
//===========================================================================

// NextNFKC is like NextNFKCLen but with the dataLength determined by len(data).
func NextNFKC(data []rune, offset int) ([]rune, int) {
	return NextNFKCLen(data, offset, len(data), false)
}

// NextNFKDLen returns the next cluster of runes normalized to NFKC, and the
// number of runes consumed from the input to produce them.
//
// The parameter, data, is a slice of code points to normalize.
//
// The parameter, offset, is the offset within data of where to normalize
// the next chunk of code points.
//
// The parameter, dataLength, is the total length, in code points, of data.
//
// The parameter, abandonAtEnd, is set to true if normalization should be
// quickly abandoned if the result may be affected by additional data. If
// this causes abandonment, then the output is an empty slice of runes with
// the length indicating to the end of the buffer. The caller might set this
// to true if it is buffering up chunks of data from a stream to be
// normalized, to indicate it should be reloaded before completing
// normalization at this step. If the caller is not buffering up data from
// a stream, then this parameter should be set false so that normalization
// is completed on the last code points in the buffer.
//
// The output is a slice of normalized runes, and the length of input data
// consumed to generate that slice.
//
// The output may be a slice from the same input array, or it may be from a
// separately allocated array.
func NextNFKCLen(data []rune, offset int, dataLength int, abandonAtEnd bool) ([]rune, int) {
	var r rune
	offsetPastQuickies := offset
	for {
		if offsetPastQuickies >= dataLength {
			if abandonAtEnd {
				if offsetPastQuickies > offset+1 {
					return data[offset : offsetPastQuickies-1], offsetPastQuickies - 1 - offset
				}
				return []rune{}, 1
			}
			return data[offset:offsetPastQuickies], offsetPastQuickies - offset
		}
		r = data[offsetPastQuickies]
		if r >= nfkcGoodBelow &&
			((r >= hangulLBase && r < hangulLBaseCount) || (r >= hangulSBase && r < hangulSBaseCount) || !IsStarterNfkc(r)) {
			break
		}
		offsetPastQuickies++
	}
	if offsetPastQuickies > offset+1 {
		return data[offset : offsetPastQuickies-1], offsetPastQuickies - 1 - offset
	}
	if dataLength-offset == 1 {
		if abandonAtEnd {
			return []rune{}, 1
		}
		if r >= hangulLBase && r < hangulLBaseCount {
			return data[offset : offset+1], 1
		}
	} else if r >= hangulLBase && r < hangulLBaseCount {
		return normComposeHangulLv(r, data, offset, dataLength, abandonAtEnd)
	} else if r >= hangulSBase && r < hangulSBaseCount {
		return normComposeHangulLvT(r, data, offset, dataLength, abandonAtEnd)
	}
	return normCommon(r, data, offset, dataLength, abandonAtEnd, true, true, HasNfkcqcY)
}

// NFKCReader is an io.RuneReader to io.RuneReader wrapper over NextNFKCLen.
// It provides the convenience of the io.RuneReader interface. It also
// copies the input stream upon being read to eliminate possible changes to
// the output resulting from modifications to the input buffer. There is a
// potentially significant performance cost of using this interface over
// calling NextNFKC/NextNFKCLen directly.
//
// Due to the cache implementation, some sequences of code points that are
// technically legal but unlikely to occur in real text may not normalize
// properly. These consist of very long sequences of non-starters that
// significantly exceed the maximum length of 30 for Stream-Safe text
// described in https://unicode.org/reports/tr15/tr15-51.html.
func NFKCReader(rr io.RuneReader) io.RuneReader {
	return newReader(rr, NextNFKCLen)
}

//===========================================================================
// API: NFKD:
//===========================================================================

// NextNFKD is like NextNFKDLen but with the dataLength determined by len(data).
func NextNFKD(data []rune, offset int) ([]rune, int) {
	return NextNFKDLen(data, offset, len(data), false)
}

// NextNFKDLen returns the next cluster of runes normalized to NFKD, and the
// number of runes consumed from the input to produce them.
//
// The parameter, data, is a slice of code points to normalize.
//
// The parameter, offset, is the offset within data of where to normalize
// the next chunk of code points.
//
// The parameter, dataLength, is the total length, in code points, of data.
//
// The parameter, abandonAtEnd, is set to true if normalization should be
// quickly abandoned if the result may be affected by additional data. If
// this causes abandonment, then the output is an empty slice of runes with
// the length indicating to the end of the buffer. The caller might set this
// to true if it is buffering up chunks of data from a stream to be
// normalized, to indicate it should be reloaded before completing
// normalization at this step. If the caller is not buffering up data from
// a stream, then this parameter should be set false so that normalization
// is completed on the last code points in the buffer.
//
// The output is a slice of normalized runes, and the length of input data
// consumed to generate that slice.
//
// The output may be a slice from the same input array, or it may be from a
// separately allocated array.
func NextNFKDLen(data []rune, offset int, dataLength int, abandonAtEnd bool) ([]rune, int) {
	var r rune
	offsetPastQuickies := offset
	for {
		if offsetPastQuickies >= dataLength {
			if abandonAtEnd {
				if offsetPastQuickies > offset+1 {
					return data[offset : offsetPastQuickies-1], offsetPastQuickies - 1 - offset
				}
				return []rune{}, 1
			}
			return data[offset:offsetPastQuickies], offsetPastQuickies - offset
		}
		r = data[offsetPastQuickies]
		if r >= nfkdGoodBelow &&
			((r >= hangulSBase && r < hangulSBaseCount) || !IsStarterNfkd(r)) {
			break
		}
		offsetPastQuickies++
	}
	if offsetPastQuickies > offset+1 {
		return data[offset : offsetPastQuickies-1], offsetPastQuickies - 1 - offset
	}
	r = data[offset]
	if dataLength-offset == 1 {
		if abandonAtEnd {
			return []rune{}, 1
		}
		if r >= hangulSBase && r < hangulSBaseCount {
			return normDecomposeHangulLvt(r)
		}
	} else if r >= hangulSBase && r < hangulSBaseCount {
		return normDecomposeHangulLvt(r)
	}
	return normCommon(r, data, offset, dataLength, abandonAtEnd, false, true, HasNfkdqcY)
}

// NFKDReader is an io.RuneReader to io.RuneReader wrapper over NextNFKDLen.
// It provides the convenience of the io.RuneReader interface. It also
// copies the input stream upon being read to eliminate possible changes to
// the output resulting from modifications to the input buffer. There is a
// potentially significant performance cost of using this interface over
// calling NextNFKD/NextNFKDLen directly.
//
// Due to the cache implementation, some sequences of code points that are
// technically legal but unlikely to occur in real text may not normalize
// properly. These consist of very long sequences of non-starters that
// significantly exceed the maximum length of 30 for Stream-Safe text
// described in https://unicode.org/reports/tr15/tr15-51.html.
func NFKDReader(rr io.RuneReader) io.RuneReader {
	return newReader(rr, NextNFKDLen)
}

//===========================================================================
// Private Helpers:
//===========================================================================

// Constants for quickly filtering out common cases.
const (
	nfcGoodBelow  rune = 0x300 // < this are assumed NFC_QC==YES and CCC==0
	nfdGoodBelow       = 0xC0  // < this are assumed NFD_QC==YES and CCC==0
	nfkcGoodBelow      = 0xA0  // < this are assumed NFKC_QC==YES and CCC==0
	nfkdGoodBelow      = 0xA0  // < this are assumed NFKD_QC==YES and CCC==0
)

// cccsZ is a 0-filler for CCC space we skip over during Hangul composition.
// It just needs to be a slice of ints > largest Hangule decomposition.
// We don't care what it contains. It just keeps cccs in line with rs. The
// contents are discarded.
var cccsZ []int = []int{0, 0, 0, 0}

// normCommon performs normalization on a sequence of runes starting at the
// current position up to, but not including, the next rune that is a starter
// (i.e. CCC==0).
//
// For more details about normalization, see chapter 3 of the Unicode
// Specification and also https://www.unicode.org/reports/tr15/.
//
// The paramter, r, is the current code point at data[offset].
//
// The parameter, data, is a slice of code points of input data.
//
// The paramter, dataLength, is count of code points at data.
//
// The parameter, abandonAtEnd, is set to true if normalization should be
// quickly abandoned if the result may be affected by additional data. If
// this causes abandonment, then the output is an empty slice of runes with
// the length indicating to the end of the buffer. The caller can use this
// to refill the buffer, if it is doing buffered input. This parameter
// should be set to false if the caller will not be loading more data into
// the input buffer at the end, to ensure that the last data gets normalized.
//
// The parameter, doComposition, is true for NFC and NFKC forms. It causes
// the decomposed data to be recomposed.
//
// The parameter, doCompatibility, is true for the NFKC and NFKD forms. It
// causes decomposition to include compatibility decompositions instead of
// just canonical decompositions.
//
// The paramter, qcFunc, is set the the quick-check function for the
// normalization form. That is, HasNfcqcY, HasNfdqcY, HasnfkcqcY, or
// HasNfkdqcY.
//
// The output is a slice of normalized runes, and the length of input data
// consumed to generate that slice.
//
// The output may be a slice from the same input array, or it may be from a
// separately allocated array.
func normCommon(r rune,
	data []rune,
	offset,
	dataLength int,
	abandonAtEnd,
	doComposition,
	doCompatibility bool,
	qcFunc func(rune) bool) ([]rune, int) {
	remainingLength := dataLength - offset
	rsActive := false
	rs := make([]rune, 0, 4)
	cccs := make([]int, 0, 4)
	cccPrev := -1
	i := 0
	for ; i < remainingLength; i++ {
		r := data[offset+i]
		ccc := ValueOfCcc(r, 0)
		if i > 0 && ccc == 0 {
			break
		}
		if (ccc < cccPrev || !qcFunc(r)) && !rsActive {
			rsActive = true
			if doComposition {
				cccsN := make([]int, 0, 4)
				rsK := []rune{}
				cccsK := []int{}
				for k := 0; k < i; k++ {
					if doCompatibility {
						rsK, cccsK = CompatibleDecomposition(data[offset+k])
					} else {
						rsK, cccsK = CanonicalDecomposition(data[offset+k])
					}
					if len(rsK) == 0 {
						rs = append(rs, data[offset+k])
						cccsN = append(cccsN, cccs[k])
					} else {
						rs = append(rs, rsK...)
						cccsN = append(cccsN, cccsK...)
					}
				}
				cccs = cccsN
			} else {
				rs = append(rs, data[offset:offset+i]...)
			}
		}
		cccPrev = ccc
		if rsActive {
			if doCompatibility {
				rsI, cccsI := CompatibleDecomposition(r)
				if len(rsI) == 0 {
					rs = append(rs, r)
					cccs = append(cccs, ccc)
				} else {
					rs = append(rs, rsI...)
					cccs = append(cccs, cccsI...)
				}
			} else {
				rsI, cccsI := CanonicalDecomposition(r)
				if len(rsI) == 0 {
					rs = append(rs, r)
					cccs = append(cccs, ccc)
				} else {
					rs = append(rs, rsI...)
					cccs = append(cccs, cccsI...)
				}
			}
		} else {
			cccs = append(cccs, ccc)
		}
	}
	if abandonAtEnd && i == remainingLength {
		// we got to the end of data so uncertain; and abandonAtEnd requested
		return []rune{}, i
	}
	if !rsActive {
		// if !rsActive, then quick check passed, so return data unchanged
		return data[offset : offset+i], i
	}
	// sort decomposition by CCC:
	//
	// Note that in most cases, the sequence is short: 1, 2, 3, and on some
	// occasions 4. It very rarely exceeds 4. Therefore, we try a fast but
	// cumbersome sort of those shorter lengths, and revert to insertion sorting
	// as a fallback. This requires extensive unit testing because each order
	// lands on a different path through the conditions.
	//
	// Also note that once we are finished sorting, we don't use the CCC's
	// (in cccs[] except during composition, to find boundaries (cccs[*]==0) and
	// blockages (cccs[i-1]==cccs[i]). So we don't need to sort cccs[] beyond
	// ensuring we keep those patterns in their respective indices. This means
	// we can skip cccs[] sorting in some cases; for example, when the length is
	// 2. In the code below, we have commented out the full cccs[] sorting lines
	// when we can, and sometimes replaced them with partial sorting if some is
	// still needed.
	if i > 1 {
		switch len(cccs) {
		case 0: // does not happen
		case 1: // does not happen
		case 2:
			if cccs[0] > cccs[1] && cccs[1] > 0 {
				//				cccs[0], cccs[1] = cccs[1], cccs[0]
				rs[0], rs[1] = rs[1], rs[0]
			}
		case 3:
			if cccs[0] == 0 {
				if cccs[1] > cccs[2] && cccs[2] > 0 {
					//					cccs[1], cccs[2] = cccs[2], cccs[1]
					rs[1], rs[2] = rs[2], rs[1]
				}
			} else if cccs[2] == 0 {
				if cccs[0] > cccs[1] && cccs[1] > 0 {
					//					cccs[0], cccs[1] = cccs[1], cccs[0]
					rs[0], rs[1] = rs[1], rs[0]
				}
			} else if cccs[1] > 0 {
				if cccs[0] > cccs[1] {
					if cccs[1] > cccs[2] {
						//						cccs[0], cccs[2] = cccs[2], cccs[0]
						rs[0], rs[2] = rs[2], rs[0]
					} else if cccs[0] > cccs[2] {
						cccs[0], cccs[1], cccs[2] = cccs[1], cccs[2], cccs[0]
						rs[0], rs[1], rs[2] = rs[1], rs[2], rs[0]
					} else {
						cccs[0], cccs[1] = cccs[1], cccs[0]
						rs[0], rs[1] = rs[1], rs[0]
					}
				} else if cccs[1] > cccs[2] {
					if cccs[0] > cccs[2] {
						cccs[0], cccs[1], cccs[2] = cccs[2], cccs[0], cccs[1]
						rs[0], rs[1], rs[2] = rs[2], rs[0], rs[1]
					} else {
						cccs[1], cccs[2] = cccs[2], cccs[1]
						rs[1], rs[2] = rs[2], rs[1]
					}
				}
			}
		case 4:
			if cccs[0] == 0 {
				// 3-way: cccs[1], cccs[2], cccs[3]
				if cccs[1] == 0 {
					if cccs[2] > cccs[3] && cccs[3] > 0 {
						//						cccs[2], cccs[3] = cccs[3], cccs[2]
						rs[2], rs[3] = rs[3], rs[2]
					}
				} else if cccs[3] == 0 {
					if cccs[1] > cccs[2] && cccs[2] > 0 {
						//						cccs[1], cccs[2] = cccs[2], cccs[1]
						rs[1], rs[2] = rs[2], rs[1]
					}
				} else if cccs[2] > 0 {

					if cccs[1] > cccs[2] {
						if cccs[2] > cccs[3] {
							//							cccs[1], cccs[3] = cccs[3], cccs[1]
							rs[1], rs[3] = rs[3], rs[1]
						} else if cccs[1] > cccs[3] {
							cccs[1], cccs[2], cccs[3] = cccs[2], cccs[3], cccs[1]
							rs[1], rs[2], rs[3] = rs[2], rs[3], rs[1]
						} else {
							cccs[1], cccs[2] = cccs[2], cccs[1]
							rs[1], rs[2] = rs[2], rs[1]
						}
					} else if cccs[2] > cccs[3] {
						if cccs[1] > cccs[3] {
							cccs[1], cccs[2], cccs[3] = cccs[3], cccs[1], cccs[2]
							rs[1], rs[2], rs[3] = rs[3], rs[1], rs[2]
						} else {
							cccs[2], cccs[3] = cccs[3], cccs[2]
							rs[2], rs[3] = rs[3], rs[2]
						}
					}
				}
			} else if cccs[3] == 0 {
				// 3-way: cccs[0], cccs[1], cccs[2]
				if cccs[2] == 0 {
					if cccs[0] > cccs[1] && cccs[1] > 0 {
						//						cccs[0], cccs[1] = cccs[1], cccs[0]
						rs[0], rs[1] = rs[1], rs[0]
					}
				} else if cccs[1] > 0 {
					if cccs[0] > cccs[1] {
						if cccs[1] > cccs[2] {
							//							cccs[0], cccs[2] = cccs[2], cccs[0]
							rs[0], rs[2] = rs[2], rs[0]
						} else if cccs[0] > cccs[2] {
							cccs[0], cccs[1], cccs[2] = cccs[1], cccs[2], cccs[0]
							rs[0], rs[1], rs[2] = rs[1], rs[2], rs[0]
						} else {
							cccs[0], cccs[1] = cccs[1], cccs[0]
							rs[0], rs[1] = rs[1], rs[0]
						}
					} else if cccs[1] > cccs[2] {
						if cccs[0] > cccs[2] {
							cccs[0], cccs[1], cccs[2] = cccs[2], cccs[0], cccs[1]
							rs[0], rs[1], rs[2] = rs[2], rs[0], rs[1]
						} else {
							cccs[1], cccs[2] = cccs[2], cccs[1]
							rs[1], rs[2] = rs[2], rs[1]
						}
					}
				}
			} else if cccs[2] == 0 {
				if cccs[0] > cccs[1] && cccs[1] > 0 {
					//					cccs[0], cccs[1] = cccs[1], cccs[0]
					rs[0], rs[1] = rs[1], rs[0]
				}
			} else if cccs[1] == 0 {
				if cccs[2] > cccs[3] {
					//					cccs[2], cccs[3] = cccs[3], cccs[2]
					rs[2], rs[3] = rs[3], rs[2]
				}
			} else {
				// 4-way: cccs[0], cccs[1], cccs[2], cccs[3]
				if cccs[0] <= cccs[1] {
					if cccs[1] <= cccs[2] {
						if cccs[2] > cccs[3] {
							if cccs[0] > cccs[3] { // 3 5 7 1
								cccs[0], cccs[1], cccs[2], cccs[3] = cccs[3], cccs[0], cccs[1], cccs[2]
								rs[0], rs[1], rs[2], rs[3] = rs[3], rs[0], rs[1], rs[2]
							} else if cccs[1] > cccs[3] { // 3 5 7 4
								cccs[1], cccs[2], cccs[3] = cccs[3], cccs[1], cccs[2]
								rs[1], rs[2], rs[3] = rs[3], rs[1], rs[2]
							} else { // 3 5 7 6
								cccs[2], cccs[3] = cccs[3], cccs[2]
								rs[2], rs[3] = rs[3], rs[2]
							}
						}
					} else if cccs[0] <= cccs[2] {
						if cccs[1] <= cccs[3] { // 3 7 5 8
							cccs[1], cccs[2] = cccs[2], cccs[1]
							rs[1], rs[2] = rs[2], rs[1]
						} else if cccs[2] <= cccs[3] { // 3 7 5 6
							cccs[1], cccs[2], cccs[3] = cccs[2], cccs[3], cccs[1]
							rs[1], rs[2], rs[3] = rs[2], rs[3], rs[1]
						} else if cccs[0] <= cccs[3] { // 3 7 5 4
							cccs[1], cccs[2], cccs[3] = cccs[3], cccs[2], cccs[1]
							rs[1], rs[2], rs[3] = rs[3], rs[2], rs[1]
						} else { // 3 7 5 1
							cccs[0], cccs[1], cccs[2], cccs[3] = cccs[3], cccs[0], cccs[2], cccs[1]
							rs[0], rs[1], rs[2], rs[3] = rs[3], rs[0], rs[2], rs[1]
						}
					} else if cccs[1] <= cccs[3] { // 5 7 1 9
						cccs[0], cccs[1], cccs[2] = cccs[2], cccs[0], cccs[1]
						rs[0], rs[1], rs[2] = rs[2], rs[0], rs[1]
					} else if cccs[0] <= cccs[3] { // 5 7 1 6
						cccs[0], cccs[1], cccs[2], cccs[3] = cccs[2], cccs[0], cccs[3], cccs[1]
						rs[0], rs[1], rs[2], rs[3] = rs[2], rs[0], rs[3], rs[1]
					} else if cccs[2] <= cccs[3] { // 5 7 1 3
						cccs[0], cccs[1], cccs[2], cccs[3] = cccs[2], cccs[3], cccs[0], cccs[1]
						rs[0], rs[1], rs[2], rs[3] = rs[2], rs[3], rs[0], rs[1]
					} else { // 5 7 3 1
						cccs[0], cccs[1], cccs[2], cccs[3] = cccs[3], cccs[2], cccs[0], cccs[1]
						rs[0], rs[1], rs[2], rs[3] = rs[3], rs[2], rs[0], rs[1]
					}
				} else if cccs[0] <= cccs[2] { // 7 4 8 ?
					if cccs[2] <= cccs[3] { // 7 4 8 9
						cccs[0], cccs[1] = cccs[1], cccs[0]
						rs[0], rs[1] = rs[1], rs[0]
					} else if cccs[0] <= cccs[3] { // 7 4 9 8
						cccs[0], cccs[1], cccs[2], cccs[3] = cccs[1], cccs[0], cccs[3], cccs[2]
						rs[0], rs[1], rs[2], rs[3] = rs[1], rs[0], rs[3], rs[2]
					} else if cccs[1] <= cccs[3] { // 7 4 8 5
						cccs[0], cccs[1], cccs[2], cccs[3] = cccs[1], cccs[3], cccs[0], cccs[2]
						rs[0], rs[1], rs[2], rs[3] = rs[1], rs[3], rs[0], rs[2]
					} else { // 7 4 8 1
						cccs[0], cccs[2], cccs[3] = cccs[3], cccs[0], cccs[2]
						rs[0], rs[2], rs[3] = rs[3], rs[0], rs[2]
					}
				} else if cccs[1] <= cccs[2] { // 7 4 5 ?
					if cccs[0] <= cccs[3] { // 7 4 5 9
						cccs[0], cccs[1], cccs[2] = cccs[1], cccs[2], cccs[0]
						rs[0], rs[1], rs[2] = rs[1], rs[2], rs[0]
					} else if cccs[2] <= cccs[3] { // 7 4 5 6
						cccs[0], cccs[1], cccs[2], cccs[3] = cccs[1], cccs[2], cccs[3], cccs[0]
						rs[0], rs[1], rs[2], rs[3] = rs[1], rs[2], rs[3], rs[0]
					} else if cccs[1] <= cccs[3] { // 7 4 6 5
						cccs[0], cccs[1], cccs[3] = cccs[1], cccs[3], cccs[0]
						rs[0], rs[1], rs[3] = rs[1], rs[3], rs[0]
					} else { // 7 4 5 1
						cccs[0], cccs[3] = cccs[3], cccs[0]
						rs[0], rs[3] = rs[3], rs[0]
					}
				} else if cccs[0] <= cccs[3] { // 7 4 1 8
					cccs[0], cccs[2] = cccs[2], cccs[0]
					rs[0], rs[2] = rs[2], rs[0]
				} else if cccs[1] <= cccs[3] { // 7 4 1 5
					cccs[0], cccs[2], cccs[3] = cccs[2], cccs[3], cccs[0]
					rs[0], rs[2], rs[3] = rs[2], rs[3], rs[0]
				} else if cccs[2] <= cccs[3] { // 7 4 1 3
					cccs[0], cccs[1], cccs[2], cccs[3] = cccs[2], cccs[3], cccs[1], cccs[0]
					rs[0], rs[1], rs[2], rs[3] = rs[2], rs[3], rs[1], rs[0]
				} else { // 7 4 2 1
					//					cccs[0], cccs[1], cccs[2], cccs[3] = cccs[3], cccs[2], cccs[1], cccs[0]
					rs[0], rs[1], rs[2], rs[3] = rs[3], rs[2], rs[1], rs[0]
				}
			}
		default:
			rsS := make([]rune, 1, len(rs))
			cccsS := make([]int, 1, len(cccs))
			rsS[0] = rs[0]
			cccsS[0] = cccs[0]
			for k := 1; k < len(cccs); k++ {
				m := len(cccsS) - 1
				for ; m >= 0; m-- {
					if cccsS[m] <= cccs[k] {
						if m >= len(cccsS)-1 {
							rsS = append(rsS, rs[k])
							cccsS = append(cccsS, cccs[k])
							break
						}
						rsS = append(rsS[:m+1], rsS[m:]...)
						cccsS = append(cccsS[:m+1], cccsS[m:]...)
						rsS[m+1] = rs[k]
						cccsS[m+1] = cccs[k]
						break
					}
				}
				if m < 0 {
					rsS = append([]rune{rs[k]}, rsS...)
					cccsS = append([]int{cccs[k]}, cccsS...)
				}
			}
			rs = rsS
			cccs = cccsS
		}
	}
	if !doComposition {
		// NFD/NFDK: we're done
		return rs, i
	}
	// canonical composition (NFC/NFCK only):
	n := 0
	kPastEnd := len(cccs)
	for k := 0; k < kPastEnd; {
		if cccs[k] == 0 {
			if doCompatibility && k < kPastEnd-1 {
				r := rs[k]
				if r >= hangulLBase && r < hangulLBaseCount {
					rsc, delta := normComposeHangulLv(r, rs, k, kPastEnd, false)
					if k+delta >= kPastEnd {
						rs = append(rs[:k], rsc...)
						break
					}
					rscLen := len(rsc)
					rsc = append(rsc, rs[k+delta:]...)
					rs = append(rs[:k], rsc...)
					cccsN := append(cccsZ[:rscLen], cccs[k+delta:]...)
					cccs = append(cccs[:k], cccsN...)
					kPastEnd += rscLen - delta
					k += rscLen
					continue
				} else if r >= hangulSBase && r < hangulSBaseCount {
					rsc, delta := normComposeHangulLvT(r, rs, k, kPastEnd, false)
					if k+delta >= kPastEnd {
						rs = append(rs[:k], rsc...)
						break
					}
					rscLen := len(rsc)
					rsc = append(rsc, rs[k+delta:]...)
					rs = append(rs[:k], rsc...)
					cccsN := append(cccsZ[:rscLen], cccs[k+delta:]...)
					cccs = append(cccs[:k], cccsN...)
					kPastEnd += rscLen - delta
					k += rscLen
					continue
				}
			}
			n = k
			k++
			continue
		}
		// Because our decomposition sorting may not have put all the cccs[] in
		// the proper order so long as it maintained cccs[*]==0 and
		// cccs[*-1]==cccs[*], and the only blockage is when they are equal, we
		// can only use cccs[k] == cccs[k-1] and not cccs[k] <= cccs[k-1].
		if k == 0 || cccs[k] == cccs[k-1] {
			k++
			continue
		}
		rc, isComposed := CanonicalComposition(rs[n], rs[k])
		if isComposed {
			rs[n] = rc
			if k < len(cccs)-1 {
				rs = append(rs[:k], rs[k+1:]...)
				cccs = append(cccs[:k], cccs[k+1:]...)
			} else {
				rs = rs[:k]
				cccs = cccs[:k]
			}
			kPastEnd--
		} else {
			k++
		}
	}
	return rs, i
}

//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Hangul:
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// Hangul syllable constants. See Section 3.12 of the Unicode Specification
// for a description of these.
const (
	hangulSBase      rune = 0xAC00
	hangulLBase           = 0x1100
	hangulVBase           = 0x1161
	hangulTBase           = 0x11A7
	hangulLCount          = 19
	hangulVCount          = 21
	hangulTCount          = 28
	hangulNCount          = hangulVCount * hangulTCount
	hangulSCount          = hangulLCount * hangulNCount
	hangulSBaseCount      = hangulSBase + hangulSCount
	hangulLBaseCount      = hangulLBase + hangulLCount
	hangulVBaseCount      = hangulVBase + hangulVCount
	hangulTBaseCount      = hangulTBase + hangulTCount
)

// normDecomposeHangulLvt handles Hangul decomposition of the forms,
// LV=>L+V and LVT=>L+V+T.
//
// It is assumed that the caller has provided parameters as follows.
//
//	r >= hangulSBase && r < hangulSBaseCount
func normDecomposeHangulLvt(r rune) ([]rune, int) {
	sIndex := r - hangulSBase
	rL := hangulLBase + sIndex/hangulNCount
	rV := hangulVBase + (sIndex%hangulNCount)/hangulTCount
	rT := hangulTBase + sIndex%hangulTCount
	if rT == hangulTBase {
		return []rune{rL, rV}, 1
	}
	return []rune{rL, rV, rT}, 1
}

// normComposeHangulLv handles Hangul composition of the forms,
// L+V=>LV and L+V+T=>LVT.
//
// It is assumed that the caller has provided parameters as follows.
//
//	r = data[offset]
//	r >= hangulLBase && r < hangulLBaseCount
//	dataLength > offset+1
func normComposeHangulLv(r rune,
	data []rune,
	offset,
	dataLength int,
	abandonAtEnd bool) ([]rune, int) {
	rv := data[offset+1]
	if rv >= hangulVBase && rv < hangulVBaseCount {
		lIndex := r - hangulLBase
		vIndex := rv - hangulVBase
		rn := hangulSBase + (lIndex*hangulVCount+vIndex)*hangulTCount
		if dataLength-offset == 2 {
			if abandonAtEnd {
				return []rune{}, 2
			}
			return []rune{rn}, 2
		}
		rt := data[offset+2]
		if rt <= hangulTBase || rt >= hangulTBaseCount {
			return []rune{rn}, 2
		}
		return []rune{rn + rt - hangulTBase}, 3
	}
	return data[offset : offset+1], 1
}

// normComposeHangulLvTLvt handles Hangul composition of the forms,
// LV+T=>LVT. It also skips over pre-composed forms (LV+^T and LVT).
//
// It is assumed that the caller has provided parameters as follows.
//
//	r = data[offset]
//	r >= hangulSBase && r < hangulSBaseCount
//	dataLength > offset+1
func normComposeHangulLvT(r rune,
	data []rune,
	offset,
	dataLength int,
	abandonAtEnd bool) ([]rune, int) {
	sIndex := r - hangulSBase
	if sIndex%hangulTCount == 0 {
		// LV; see if LV, T => LVT applies
		rt := data[offset+1]
		if rt < hangulTBase+1 || rt >= hangulTBase+hangulTCount {
			return data[offset : offset+1], 1
		}
		tIndex := rt - hangulTBase
		return []rune{r + tIndex}, 2
	}
	return data[offset : offset+1], 1
}

//===========================================================================
// Private Helpers: RuneReader Interface:
//===========================================================================

// runeReaderNormCacheSize is the maximum cache size for our io.RuneReader
// implementation. It should be >= the maximum number of runes consumed via
// normalization. In realistic cases, this should not be more than a few
// code points, but, theoretically, it can be unlimited. UAX#15 defines a
// Stream-Safe Text Format of which the maximum buffer size needed is 32
// code points. Therefore, anything >= 32 should work in all realistic
// scenerios.
const runeReaderNormCacheSize = 256

// newReader provides an io.RuneReader to io.RuneReader wrapper over one of
// our normalization functions of the form: Next*Len.
func newReader(rr io.RuneReader, normFunc func([]rune, int, int, bool) ([]rune, int)) io.RuneReader {
	var rrn runeReaderNorm
	var impl runeReaderNormImpl
	impl.runeReader = rr
	impl.normFunc = normFunc
	rrn.impl = &impl
	return rrn
}

// runeReaderNorm implements an io.RuneReader interface for wrapping the
// normalization functions.
type runeReaderNorm struct {
	impl *runeReaderNormImpl
}

// runesREaderBreakerImpl supports an implementation of the RunesReader interface
// for wrapping the segment break functions in the ubasic and usegs packages.
type runeReaderNormImpl struct {
	runeReader        io.RuneReader
	runesCache        []rune
	runesCacheSize    int
	offset            int
	atEnd             bool
	runesOutCache     []rune
	runesOutCacheSize int
	runesOutOffset    int
	normFunc          func([]rune, int, int, bool) ([]rune, int)
}

// ReadRunes provides the RunesReader interface to runesReaderBreaker.
func (rrn runeReaderNorm) ReadRune() (rune, int, error) {
	var err error = nil
	var impl = rrn.impl
	if impl.runesOutOffset < impl.runesOutCacheSize {
		r := impl.runesOutCache[impl.runesOutOffset]
		impl.runesOutOffset++
		return runesOut(r)
	}
	if impl.runesCacheSize <= impl.offset {
		err = impl.reload()
		if err != nil {
			return 0, 0, err
		}
	}
	if impl.runesCacheSize <= impl.offset {
		return 0, 0, io.EOF
	}
	rs, d := impl.normFunc(impl.runesCache, impl.offset, impl.runesCacheSize, true)
	if len(rs) == 0 && d > 0 {
		impl.reload() // we don't care about the err here
		rs, d = impl.normFunc(impl.runesCache, impl.offset, impl.runesCacheSize, false)
	}
	impl.offset += d
	impl.runesOutCache = rs
	impl.runesOutCacheSize = len(rs)
	impl.runesOutOffset = 0
	if impl.runesOutOffset < impl.runesOutCacheSize {
		r := impl.runesOutCache[impl.runesOutOffset]
		impl.runesOutOffset++
		return runesOut(r)
	}
	return 0, 0, io.EOF
}

// reload is a private method to reload the cache.
func (impl *runeReaderNormImpl) reload() error {
	if !impl.atEnd {
		impl.runesCache = append(make([]rune, 0, runeReaderNormCacheSize), impl.runesCache[impl.offset:]...)
		impl.runesCacheSize -= impl.offset
		impl.offset = 0
		for impl.runesCacheSize < runeReaderNormCacheSize {
			r, _, err := impl.runeReader.ReadRune()
			if err == nil {
				impl.runesCache = append(impl.runesCache, r)
				impl.runesCacheSize++
				continue
			}
			impl.atEnd = true
			if err != io.EOF {
				return err
			}
			return nil
		}
	}
	return nil
}

// runesOut takes a rune and translates it into the output of ReadRune()
func runesOut(r rune) (rune, int, error) {
	if r <= 0x7f {
		return r, 1, nil
	}
	if r <= 0x7ff {
		return r, 2, nil
	}
	if r <= 0xffff {
		return r, 3, nil
	}
	return r, 4, nil
}
