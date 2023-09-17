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

package ubasic

//===========================================================================
// API: Extended Grapheme Cluster Breaks:
//===========================================================================

// NextExtGraphemeClusterBreak is like NextExtGraphemeClusterBreakLen but with the dataLength
// determined by len(data).
//
// The second returned parameter is true if the break is mandatory here. This
// is always true for this function.
func NextExtGraphemeClusterBreak(data []rune, offset int) (int, bool) {
	return NextExtGraphemeClusterBreakLen(data, offset, len(data))
}

// NextExtGraphemeClusterBreakLen returns the number of runes from the current offset
// until the next extended grapheme cluster break.
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
	// For informmation on this algorithm, see:
	// https://www.unicode.org/reports/tr29/
	remainingLength := dataLength - offset
	previousLength := 1
	gbState := gbStateNew()
	gb9cState := gb9cStateNew()
	for {
		if offset >= dataLength {
			return previousLength - 1, true
		}
		if remainingLength == 1 {
			return previousLength, true
		}
		r := data[offset]
		rn := data[offset+1]
		gbState.updateR(r)
		gb9cState.updateR(r)
		mustContinue, mustBreak := gbState.rulesGb3throughGb9b(rn)
		if mustBreak {
			return previousLength, true
		}
		if mustContinue {
			offset++
			remainingLength--
			previousLength++
			continue
		}
		mustContinue = gb9cState.ruleGb9c(rn)
		if mustContinue {
			offset++
			remainingLength--
			previousLength++
			continue
		}
		_, mustBreak = gbState.rulesGb11throughGb999(rn)
		if mustBreak {
			return previousLength, true
		}
		offset++
		remainingLength--
		previousLength++
	}
}

//===========================================================================
// Private Helpers: rules GB3 through GB999, except GB9c
//===========================================================================

// Constants for keeping track of current state of these rules:
const (
	gbNone          = iota // none of the other states
	gbCr                   // GB3, GB4: CR
	gbLfCn                 // GB4: (LF | Control)
	gbL                    // GB6: L
	gbLvV                  // GB7: LV | V
	gbLvtT                 // GB8: LVT | T
	gbPrepend              // GB9b
	gbExtP                 // GB911: \p{Extended_Pictographic}
	gbExtPExtend           // GB911: \p{Extended_Pictographic} Extend*
	gbExtPExtendZwj        // GB911: \p{Extended_Pictographic} Extend* ZWJ
	gbRiOdd                // GB912, GB13: odd length of RI sequence
	gbRiEven               // GB912, GB13: even length of RI sequence
	gbUndefined            // used for undefined values
)

// gbState holds the current parsing state for >=GB3 except GB9c
type gbState struct {
	cur     int            // current state (see gb* constants)
	next    int            // next state (see gb* constants), or -1 if unknown yet
	nextGcb int            // ValueOfGcb(r) for next rune, or -1 if unknown yet
}

// gbStateNew returns a new gbState initialized to the starting state
func gbStateNew() gbState {
	var s gbState
	s.cur = gbNone
	s.next = gbUndefined
	s.nextGcb = -1
	return s
}

// updateR updates the current gbState for the current rune
func (s *gbState) updateR(r rune) {
	if s.next != gbUndefined {
		s.cur = s.next
		s.next = gbUndefined
		s.nextGcb = -1
	} else {
		nextGcb := s.nextGcb
		s.nextGcb = -1
		if nextGcb < 0 {
			nextGcb = ValueOfGcb(r, -1)
		}
		switch nextGcb {
		case GcbValueCr:
			s.cur = gbCr
		case GcbValueLf, GcbValueCn:
			s.cur = gbLfCn
		case GcbValueL:
			s.cur = gbL
		case GcbValueLv, GcbValueV:
			s.cur = gbLvV
		case GcbValueLvt, GcbValueT:
			s.cur = gbLvtT
		case GcbValuePp:
			s.cur = gbPrepend
		case GcbValueRi:
			// This code assumes that there are no collisions between
			//   - Grapheme_Cluster_Break == Regional_Indicator: ValueOfGcb(r,-1)==GcbValueRi
			//   = Extended Pictographic: HasExtpictY
			// If there were such a collision, then we must check for HasExtPict here first
			// to ensure that rule GB11 takes precedence over rules GB12 and GB13.
			if s.cur == gbRiOdd {
				s.cur = gbRiEven
			} else {
				s.cur = gbRiOdd
			}
		case GcbValueEx:
			switch s.cur {
			case gbExtP, gbExtPExtend:
				s.cur = gbExtPExtend
			default:
				s.cur = gbNone
			}
		case GcbValueZwj:
			switch s.cur {
			case gbExtP, gbExtPExtend:
				s.cur = gbExtPExtendZwj
			default:
				s.cur = gbNone
			}
		default:
			if HasExtpictY(r) {
				s.cur = gbExtP
			} else {
				s.cur = gbNone
			}
		}
	}
}

// rulesGb3throughGb9b processes rules GB3 through GB9b.
//
// The first returned value is true if we should not break here, while
// the second returned value is true if we should break here.
func (s *gbState) rulesGb3throughGb9b(rn rune) (bool, bool) {
	s.nextGcb = ValueOfGcb(rn, -1)
	switch s.cur {
	case gbCr:
		if s.nextGcb == GcbValueLf {
			s.next = gbLfCn
			return true, false
		}
		return false, true
	case gbLfCn:
		return false, true
	}
	switch s.nextGcb {
	case GcbValueCn, GcbValueCr, GcbValueLf:
		return false, true
	}
	switch s.cur {
	case gbL:
		switch s.nextGcb {
		case GcbValueL:
			s.next = gbL
			return true, false
		case GcbValueLv, GcbValueV:
			s.next = gbLvV
			return true, false
		case GcbValueLvt:
			s.next = gbLvtT
			return true, false
		}
	case gbLvV:
		switch s.nextGcb {
		case GcbValueV:
			s.next = gbLvV
			return true, false
		case GcbValueT:
			s.next = gbLvtT
			return true, false
		}
	case gbLvtT:
		switch s.nextGcb {
		case GcbValueT:
			s.next = gbLvtT
			return true, false
		}
	case gbExtP, gbExtPExtend:
		// We are processing rule GB9, but we must still keep track of state for
		// rule GB11, in case we use it later.
		switch s.nextGcb {
		case GcbValueEx:
			s.next = gbExtPExtend
			return true, false
		case GcbValueZwj:
			s.next = gbExtPExtendZwj
			return true, false
		}
	}
	switch s.nextGcb {
	case GcbValueEx, GcbValueZwj, GcbValueSm:
		s.next = gbNone
		return true, false
	}
	switch s.cur {
	case gbPrepend:
		return true, false
	}
	return false, false
}

// rulesGb3throughGb9b processes rules GB11 through GB999.
//
// The first returned value is true if we should not break here, while
// the second returned value is true if we should break here.
func (s *gbState) rulesGb11throughGb999(rn rune) (bool, bool) {
	switch s.cur {
	case gbExtP:
		switch s.nextGcb {
		case GcbValueEx:
			s.next = gbExtPExtend
			return false, false
		}
	case gbExtPExtend:
		switch s.nextGcb {
		case GcbValueEx:
			s.next = gbExtPExtend
			return false, false
		case GcbValueZwj:
			s.next = gbExtPExtendZwj
			return false, false
		}
	case gbExtPExtendZwj:
		if HasExtpictY(rn) {
			s.next = gbExtP
			return true, false
		}
	case gbRiOdd:
		switch s.nextGcb {
		case GcbValueRi:
			s.next = gbRiEven
			return true, false
		}
	case gbRiEven:
		return false, true
	}
	return false, true
}

//===========================================================================
// Private Helpers: rule GB9c - Indic_Conjunct_Break:
//===========================================================================

// Constants for keeping track of current state of rule GB9c:
const (
	gb9cNone          = iota // not in GB9c rule pattern
	gb9cConsonant            // \p{InCB=Consonant}
	gb9cExtendExtends        // \p{InCB=Consonant} [ \p{InCB=Extend} ]+
	gb9cExtendLinker         // \p{InCB=Consonant} [ \p{InCB=Extend} \p{InCB=Linker} ]* (>=1 Linker)
)

// gb9cState holds the current parsing state for GB9c
type gb9cState struct {
	val      int
	nextInCb int
}

// gbStateNew returns a new gb9cState initialized to the starting state
func gb9cStateNew() gb9cState {
	var s gb9cState
	s.val = gb9cNone
	s.nextInCb = -1
	return s
}

// updateR updates the current gb9cState for the current rune
func (s *gb9cState) updateR(r rune) {
	nextInCb := s.nextInCb
	s.nextInCb = -1
	if nextInCb < 0 {
		nextInCb = ValueOfIncb(r, IncbValueNone)
	}
	switch nextInCb {
	case IncbValueNone:
		s.val = gb9cNone
	case IncbValueConsonant:
		s.val = gb9cConsonant
	case IncbValueExtend:
		switch s.val {
		case gb9cConsonant:
			s.val = gb9cExtendExtends
		}
	case IncbValueLinker:
		switch s.val {
		case gb9cConsonant, gb9cExtendExtends:
			s.val = gb9cExtendLinker
		}
	default:
		s.val = gb9cNone
	}
}

// rulesGb3throughGb9b processes rule GB9c.
//
// The value, true, is returned if we should break here.
func (s *gb9cState) ruleGb9c(rn rune) bool {
	if s.val == gb9cExtendLinker {
		s.nextInCb = ValueOfIncb(rn, IncbValueNone)
		if s.nextInCb == IncbValueConsonant {
			return true
		}
	}
	return false
}
