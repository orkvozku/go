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

/*
Package ucutls provides common utilities for other packages of this module.
*/
package ucutls

import "io"

//===========================================================================
// API: RunesReader Interface:
//===========================================================================

// Runesreader is the interface that wraps the ReadRunes method.
//
// ReadRunes reads a single extended grapheme cluster of runes as a []rune.
// If no runes are available, err will be set. The middle parameter should
// be ignored for now. It is a placeholder for future support that provides
// additional information about the break.
type RunesReader interface {
	ReadRunes() ([]rune, bool, error)
}

// IsEOF returns true only if the error is io.EOF indicating that the
// RunesReader is at the end of data.
func IsEOF(err error) bool {
	return err == io.EOF
}

// NewRunesReaderFromBreaker returns an implementation of the RunesReader
// interface wrapping one of the Next*BreakLen functions in the ubasic and
// useg packages.
//
// cacheFillSize is the size, in runes of the local cache. It must be set
// large enough to span breaks. If the entire span from one break to the
// next does not fit into the cache, then an arbirary break will be forced.
// Generally, sentence breaks require a much larger cacheFillSize than
// extended grapheme cluster breaks.
//
// rr is an io.RuneReader interface that provides incomding data to the
// Next*BreakLen function.
func NewRunesReaderFromBreaker(nbf func([]rune, int, int) (int, bool),
	cacheFillSize int,
	rr io.RuneReader) RunesReader {
	var rrb runesReaderBreaker
	var rrbi runesReaderBreakerImpl
	rrb.impl = &rrbi
	rrbi.runeReader = rr
	rrbi.cacheFillSize = cacheFillSize
	rrbi.nextBreakFunc = nbf
	var rsr RunesReader = rrb
	return rsr
}

//===========================================================================
// Private Helpers: RunesReader Interface:
//===========================================================================

// runesReaderBreaker implements a RunesReader interface for wrapping the
// segment break functions in the ubasic and usegs packages.
type runesReaderBreaker struct {
	impl *runesReaderBreakerImpl
}

// runesReaderBreakerImpl supports an implementation of the RunesReader interface
// for wrapping the segment break functions in the ubasic and usegs packages.
type runesReaderBreakerImpl struct {
	runeReader     io.RuneReader
	runesCache     []rune
	cacheFillSize  int
	runesCacheSize int
	offset         int
	atEnd          bool
	nextBreakFunc  func([]rune, int, int) (int, bool)
}

// ReadRunes provides the RunesReader interface to runesReaderBreaker.
func (rrb runesReaderBreaker) ReadRunes() ([]rune, bool, error) {
	var err error = nil
	var rrbi = rrb.impl
	if rrbi.runesCacheSize <= rrbi.offset {
		err = rrbi.reload()
		if err != nil {
			return []rune{}, true, err
		}
	}
	if rrbi.runesCacheSize <= rrbi.offset {
		return []rune{}, true, io.EOF
	}
	delta, isMandatory := rrbi.nextBreakFunc(rrbi.runesCache, rrbi.offset, rrbi.runesCacheSize)
	newOffset := rrbi.offset + delta
	if rrbi.runesCacheSize <= newOffset {
		err = rrbi.reload()
		if err != nil {
			return []rune{}, true, err
		}
		delta, isMandatory = rrbi.nextBreakFunc(rrbi.runesCache, rrbi.offset, rrbi.runesCacheSize)
		newOffset = rrbi.offset + delta
	}
	if newOffset <= rrbi.runesCacheSize {
		offsetStart := rrbi.offset
		rrbi.offset = newOffset
		return rrbi.runesCache[offsetStart:newOffset], isMandatory, nil
	}
	if newOffset > rrbi.offset {
		rrbi.offset = newOffset
		return rrbi.runesCache[rrbi.offset:], isMandatory, nil
	}
	return []rune{}, true, io.EOF
}

// reload is a private method to reload the cache.
func (rrbi *runesReaderBreakerImpl) reload() error {
	if !rrbi.atEnd {
		rrbi.runesCache = append(make([]rune, 0, rrbi.cacheFillSize), rrbi.runesCache[rrbi.offset:]...)
		rrbi.runesCacheSize -= rrbi.offset
		rrbi.offset = 0
		for rrbi.runesCacheSize < rrbi.cacheFillSize {
			r, _, err := rrbi.runeReader.ReadRune()
			if err == nil {
				rrbi.runesCache = append(rrbi.runesCache, r)
				rrbi.runesCacheSize++
				continue
			}
			rrbi.atEnd = true
			if err != io.EOF {
				return err
			}
			return nil
		}
	}
	return nil
}
