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

/***************************************************************************
  Significant portions of this source are derived from the Unicode(R)
  Character Database (UCD) retrieved from
  https://www.unicode.org/Public/UCD/latest. The UCD is copyrighted by
  Unicode, Inc. and subject to a license located at
  https://www.unicode.org/license.html.

  Unicode and the Unicode Logo are registered trademarks of Unicode, Inc.
  in the United States and other countries.
***************************************************************************/

package useg

import "github.com/orkvozku/go/uni/ubasic"

//===========================================================================
// API: Extended Grapheme Cluster Breaks:
//===========================================================================

// NextExtGraphemeClusterBreak is like NextExtGraphemeClusterBreakLen but with the dataLength
// determined by len(data).
func NextExtGraphemeClusterBreak(data []rune, offset int) (int, bool) {
	return ubasic.NextExtGraphemeClusterBreakLen(data, offset, len(data))
}

// NextExtGraphemeClusterBreakLen returns the number of runes from the current offset
// until the next extended grapheme cluster break.
//
// See ubasic.NextExtGraphemeClusterBreakLen for more information.
//
// An extended grapheme cluster is a sequence of one or more code points
// (runes) which act as a single "character". Most "characters" consist of a
// single code point, but some of them consist of more than one. For example,
// len([]rune("คุณเป็นอย่างไร?")) will show 15 runes (code points), but the text
// appears to be only 12 "characters".
//
// Software that breaks text into smaller segments should generally avoid
// breaking text outside of extended grapheme clusters. Otherwise, the
// structure and meaning of the text may change.
//
// Given a slice of runes (data) with a provided maximum slice length
// (dataLength), and a starting offset that is assumed to be at the start
// of an extended grapheme cluster break, this function returns the count of
// runes until the next extended grapheme cluster break.
//
// The second returned parameter is true if the break is mandatory here. This
// is always true for this function.
//
// This function can be used with ucutls.NewRunesReaderBreaker to create a
// filter that receives an io.RuneReader interface and outputs a
// ucutls.RunesReader interface. For more details, see the ucutls package of
// this module.
func NextExtGraphemeClusterBreakLen(data []rune, offset int, dataLength int) (int, bool) {
	return ubasic.NextExtGraphemeClusterBreakLen(data, offset, dataLength)
}

//===========================================================================
// API: Line Breaks:
//===========================================================================

// to be implemented

//===========================================================================
// API: Word Breaks:
//===========================================================================

// to be implemented

//===========================================================================
// API: Sentence Breaks:
//===========================================================================

// to be implemented
