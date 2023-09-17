/***************************************************************************
  Copyright 2021-2023 https://github.com user @orkvozku

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

/***************************************************************************
  This source code is not intended to be modified by biological entities.

  It is recommended that biological entities review the test cases described
  by the variable, autogenTestCases, to ensure they appear to describe the
  expected results for the corresponding ranges.
***************************************************************************/

package uemojis_test

import "testing"
import "github.com/orkvozku/go/uni/uemojis"

// autogenTestCaseFunc contains a function to test.
type autogenTestCaseFunc struct {
	label string
	f     func(rune) bool
}

// autogenTestCasesFunc... describes a function to test.
var (
	autogenTestCasesFuncEbaseN   = autogenTestCaseFunc{"HasEbaseN", uemojis.HasEbaseN}
	autogenTestCasesFuncEbaseY   = autogenTestCaseFunc{"HasEbaseY", uemojis.HasEbaseY}
	autogenTestCasesFuncEcompN   = autogenTestCaseFunc{"HasEcompN", uemojis.HasEcompN}
	autogenTestCasesFuncEcompY   = autogenTestCaseFunc{"HasEcompY", uemojis.HasEcompY}
	autogenTestCasesFuncEmodN    = autogenTestCaseFunc{"HasEmodN", uemojis.HasEmodN}
	autogenTestCasesFuncEmodY    = autogenTestCaseFunc{"HasEmodY", uemojis.HasEmodY}
	autogenTestCasesFuncEmojiN   = autogenTestCaseFunc{"HasEmojiN", uemojis.HasEmojiN}
	autogenTestCasesFuncEmojiY   = autogenTestCaseFunc{"HasEmojiY", uemojis.HasEmojiY}
	autogenTestCasesFuncEpresN   = autogenTestCaseFunc{"HasEpresN", uemojis.HasEpresN}
	autogenTestCasesFuncEpresY   = autogenTestCaseFunc{"HasEpresY", uemojis.HasEpresY}
	autogenTestCasesFuncExtpictN = autogenTestCaseFunc{"HasExtpictN", uemojis.HasExtpictN}
	autogenTestCasesFuncExtpictY = autogenTestCaseFunc{"HasExtpictY", uemojis.HasExtpictY}
	autogenTestCasesFuncRiN      = autogenTestCaseFunc{"HasRiN", uemojis.HasRiN}
	autogenTestCasesFuncRiY      = autogenTestCaseFunc{"HasRiY", uemojis.HasRiY}
)

// autogenTestCase describes an entry in autogenTestCases.
type autogenTestCase struct {
	rMin       rune                   // code point range: minimum
	rMax       rune                   // code point range: maximum
	funcsTrue  []*autogenTestCaseFunc // test functions: true
	funcsFalse []*autogenTestCaseFunc // test functions: false
}

// autogenTestCases describes all of our test cases.
var autogenTestCases = [][]autogenTestCase{
	{ // Ebase
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
			&autogenTestCasesFuncEbaseY,
		}},
		{0x00, 0x261C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x261D, 0x261D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x261E, 0x26F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x26F9, 0x26F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x26FA, 0x2709, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x270A, 0x270D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x270E, 0x1F384, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F385, 0x1F385, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F386, 0x1F3C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F3C2, 0x1F3C4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F3C5, 0x1F3C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F3C7, 0x1F3C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F3C8, 0x1F3C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F3CA, 0x1F3CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F3CD, 0x1F441, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F442, 0x1F443, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F444, 0x1F445, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F446, 0x1F450, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F451, 0x1F465, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F466, 0x1F478, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F479, 0x1F47B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F47C, 0x1F47C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F47D, 0x1F480, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F481, 0x1F483, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F484, 0x1F484, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F485, 0x1F487, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F488, 0x1F48E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F48F, 0x1F48F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F490, 0x1F490, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F491, 0x1F491, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F492, 0x1F4A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F4AA, 0x1F4AA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F4AB, 0x1F573, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F574, 0x1F575, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F576, 0x1F579, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F57A, 0x1F57A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F57B, 0x1F58F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F590, 0x1F590, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F591, 0x1F594, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F595, 0x1F596, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F597, 0x1F644, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F645, 0x1F647, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F648, 0x1F64A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F64B, 0x1F64F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F650, 0x1F6A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6A3, 0x1F6A3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6A4, 0x1F6B3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6B4, 0x1F6B6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6B7, 0x1F6BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6C0, 0x1F6C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6C1, 0x1F6CB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6CC, 0x1F6CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6CD, 0x1F90B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F90C, 0x1F90C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F90D, 0x1F90E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F90F, 0x1F90F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F910, 0x1F917, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F918, 0x1F91F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F920, 0x1F925, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F926, 0x1F926, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F927, 0x1F92F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F930, 0x1F939, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F93A, 0x1F93B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F93C, 0x1F93E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F93F, 0x1F976, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F977, 0x1F977, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F978, 0x1F9B4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F9B5, 0x1F9B6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F9B7, 0x1F9B7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F9B8, 0x1F9B9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F9BA, 0x1F9BA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F9BB, 0x1F9BB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F9BC, 0x1F9CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F9CD, 0x1F9CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F9D0, 0x1F9D0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1F9D1, 0x1F9DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1F9DE, 0x1FAC2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1FAC3, 0x1FAC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1FAC6, 0x1FAEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x1FAF0, 0x1FAF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseY,
		}, []*autogenTestCaseFunc{}},
		{0x1FAF9, 0x10FFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEbaseN,
			&autogenTestCasesFuncEbaseY,
		}},
	},
	{ // Ecomp
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompN,
			&autogenTestCasesFuncEcompY,
		}},
		{0x00, 0x22, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompN,
		}, []*autogenTestCaseFunc{}},
		{0x23, 0x23, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompY,
		}, []*autogenTestCaseFunc{}},
		{0x24, 0x29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompN,
		}, []*autogenTestCaseFunc{}},
		{0x2A, 0x2A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompY,
		}, []*autogenTestCaseFunc{}},
		{0x2B, 0x2F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompN,
		}, []*autogenTestCaseFunc{}},
		{0x30, 0x39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompY,
		}, []*autogenTestCaseFunc{}},
		{0x3A, 0x200C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompN,
		}, []*autogenTestCaseFunc{}},
		{0x200D, 0x200D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompY,
		}, []*autogenTestCaseFunc{}},
		{0x200E, 0x20E2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompN,
		}, []*autogenTestCaseFunc{}},
		{0x20E3, 0x20E3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompY,
		}, []*autogenTestCaseFunc{}},
		{0x20E4, 0xFE0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompN,
		}, []*autogenTestCaseFunc{}},
		{0xFE0F, 0xFE0F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompY,
		}, []*autogenTestCaseFunc{}},
		{0xFE10, 0x1F1E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompN,
		}, []*autogenTestCaseFunc{}},
		{0x1F1E6, 0x1F1FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompY,
		}, []*autogenTestCaseFunc{}},
		{0x1F200, 0x1F3FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompN,
		}, []*autogenTestCaseFunc{}},
		{0x1F3FB, 0x1F3FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompY,
		}, []*autogenTestCaseFunc{}},
		{0x1F400, 0x1F9AF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompN,
		}, []*autogenTestCaseFunc{}},
		{0x1F9B0, 0x1F9B3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompY,
		}, []*autogenTestCaseFunc{}},
		{0x1F9B4, 0xE001F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompN,
		}, []*autogenTestCaseFunc{}},
		{0xE0020, 0xE007F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompY,
		}, []*autogenTestCaseFunc{}},
		{0xE0080, 0x10FFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompN,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEcompN,
			&autogenTestCasesFuncEcompY,
		}},
	},
	{ // Emod
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmodN,
			&autogenTestCasesFuncEmodY,
		}},
		{0x00, 0x1F3FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmodN,
		}, []*autogenTestCaseFunc{}},
		{0x1F3FB, 0x1F3FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmodY,
		}, []*autogenTestCaseFunc{}},
		{0x1F400, 0x10FFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmodN,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmodN,
			&autogenTestCasesFuncEmodY,
		}},
	},
	{ // Emoji
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
			&autogenTestCasesFuncEmojiY,
		}},
		{0x00, 0x22, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x23, 0x23, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x24, 0x29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2A, 0x2A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2B, 0x2F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x30, 0x39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x3A, 0xA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0xA9, 0xA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0xAA, 0xAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0xAE, 0xAE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0xAF, 0x203B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x203C, 0x203C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x203D, 0x2048, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2049, 0x2049, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x204A, 0x2121, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2122, 0x2122, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2123, 0x2138, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2139, 0x2139, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x213A, 0x2193, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2194, 0x2199, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x219A, 0x21A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x21A9, 0x21AA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x21AB, 0x2319, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x231A, 0x231B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x231C, 0x2327, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2328, 0x2328, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2329, 0x23CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x23CF, 0x23CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x23D0, 0x23E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x23E9, 0x23F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x23F4, 0x23F7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x23F8, 0x23FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x23FB, 0x24C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x24C2, 0x24C2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x24C3, 0x25A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x25AA, 0x25AB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x25AC, 0x25B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x25B6, 0x25B6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x25B7, 0x25BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x25C0, 0x25C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x25C1, 0x25FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x25FB, 0x25FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x25FF, 0x25FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2600, 0x2604, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2605, 0x260D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x260E, 0x260E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x260F, 0x2610, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2611, 0x2611, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2612, 0x2613, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2614, 0x2615, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2616, 0x2617, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2618, 0x2618, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2619, 0x261C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x261D, 0x261D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x261E, 0x261F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2620, 0x2620, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2621, 0x2621, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2622, 0x2623, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2624, 0x2625, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2626, 0x2626, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2627, 0x2629, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x262A, 0x262A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x262B, 0x262D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x262E, 0x262F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2630, 0x2637, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2638, 0x263A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x263B, 0x263F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2640, 0x2640, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2641, 0x2641, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2642, 0x2642, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2643, 0x2647, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2648, 0x2653, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2654, 0x265E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x265F, 0x2660, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2661, 0x2662, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2663, 0x2663, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2664, 0x2664, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2665, 0x2666, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2667, 0x2667, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2668, 0x2668, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2669, 0x267A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x267B, 0x267B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x267C, 0x267D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x267E, 0x267F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2680, 0x2691, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2692, 0x2697, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2698, 0x2698, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2699, 0x2699, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x269A, 0x269A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x269B, 0x269C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x269D, 0x269F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26A0, 0x26A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26A2, 0x26A6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26A7, 0x26A7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26A8, 0x26A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26AA, 0x26AB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26AC, 0x26AF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26B0, 0x26B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26B2, 0x26BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26BD, 0x26BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26BF, 0x26C3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26C4, 0x26C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26C6, 0x26C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26C8, 0x26C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26C9, 0x26CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26CE, 0x26CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26D0, 0x26D0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26D1, 0x26D1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26D2, 0x26D2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26D3, 0x26D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26D5, 0x26E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26E9, 0x26EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26EB, 0x26EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26F0, 0x26F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26F6, 0x26F6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26F7, 0x26FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26FB, 0x26FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x26FD, 0x26FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x26FE, 0x2701, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2702, 0x2702, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2703, 0x2704, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2705, 0x2705, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2706, 0x2707, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2708, 0x270D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x270E, 0x270E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x270F, 0x270F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2710, 0x2711, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2712, 0x2712, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2713, 0x2713, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2714, 0x2714, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2715, 0x2715, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2716, 0x2716, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2717, 0x271C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x271D, 0x271D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x271E, 0x2720, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2721, 0x2721, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2722, 0x2727, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2728, 0x2728, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2729, 0x2732, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2733, 0x2734, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2735, 0x2743, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2744, 0x2744, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2745, 0x2746, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2747, 0x2747, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2748, 0x274B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x274C, 0x274C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x274D, 0x274D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x274E, 0x274E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x274F, 0x2752, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2753, 0x2755, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2756, 0x2756, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2757, 0x2757, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2758, 0x2762, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2763, 0x2764, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2765, 0x2794, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2795, 0x2797, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2798, 0x27A0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x27A1, 0x27A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x27A2, 0x27AF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x27B0, 0x27B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x27B1, 0x27BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x27BF, 0x27BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x27C0, 0x2933, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2934, 0x2935, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2936, 0x2B04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2B05, 0x2B07, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2B08, 0x2B1A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2B1B, 0x2B1C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2B1D, 0x2B4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2B50, 0x2B50, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2B51, 0x2B54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x2B55, 0x2B55, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x2B56, 0x302F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x3030, 0x3030, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x3031, 0x303C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x303D, 0x303D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x303E, 0x3296, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x3297, 0x3297, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x3298, 0x3298, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x3299, 0x3299, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x329A, 0x1F003, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F004, 0x1F004, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F005, 0x1F0CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F0CF, 0x1F0CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F0D0, 0x1F16F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F170, 0x1F171, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F172, 0x1F17D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F17E, 0x1F17F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F180, 0x1F18D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F18E, 0x1F18E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F18F, 0x1F190, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F191, 0x1F19A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F19B, 0x1F1E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F1E6, 0x1F1FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F200, 0x1F200, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F201, 0x1F202, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F203, 0x1F219, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F21A, 0x1F21A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F21B, 0x1F22E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F22F, 0x1F22F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F230, 0x1F231, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F232, 0x1F23A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F23B, 0x1F24F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F250, 0x1F251, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F252, 0x1F2FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F300, 0x1F321, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F322, 0x1F323, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F324, 0x1F393, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F394, 0x1F395, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F396, 0x1F397, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F398, 0x1F398, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F399, 0x1F39B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F39C, 0x1F39D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F39E, 0x1F3F0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F3F1, 0x1F3F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F3F3, 0x1F3F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F3F6, 0x1F3F6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F3F7, 0x1F4FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F4FE, 0x1F4FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F4FF, 0x1F53D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F53E, 0x1F548, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F549, 0x1F54E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F54F, 0x1F54F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F550, 0x1F567, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F568, 0x1F56E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F56F, 0x1F570, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F571, 0x1F572, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F573, 0x1F57A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F57B, 0x1F586, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F587, 0x1F587, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F588, 0x1F589, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F58A, 0x1F58D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F58E, 0x1F58F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F590, 0x1F590, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F591, 0x1F594, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F595, 0x1F596, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F597, 0x1F5A3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5A4, 0x1F5A5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5A6, 0x1F5A7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5A8, 0x1F5A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5A9, 0x1F5B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5B1, 0x1F5B2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5B3, 0x1F5BB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5BC, 0x1F5BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5BD, 0x1F5C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5C2, 0x1F5C4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5C5, 0x1F5D0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5D1, 0x1F5D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5D4, 0x1F5DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5DC, 0x1F5DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5DF, 0x1F5E0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5E1, 0x1F5E1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5E2, 0x1F5E2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5E3, 0x1F5E3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5E4, 0x1F5E7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5E8, 0x1F5E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5E9, 0x1F5EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5EF, 0x1F5EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5F0, 0x1F5F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5F3, 0x1F5F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5F4, 0x1F5F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5FA, 0x1F64F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F650, 0x1F67F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F680, 0x1F6C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6C6, 0x1F6CA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6CB, 0x1F6D2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6D3, 0x1F6D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6D5, 0x1F6D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6D8, 0x1F6DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6DC, 0x1F6E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6E6, 0x1F6E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6E9, 0x1F6E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6EA, 0x1F6EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6EB, 0x1F6EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6ED, 0x1F6EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6F0, 0x1F6F0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6F1, 0x1F6F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6F3, 0x1F6FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6FD, 0x1F7DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F7E0, 0x1F7EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F7EC, 0x1F7EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F7F0, 0x1F7F0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F7F1, 0x1F90B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F90C, 0x1F93A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F93B, 0x1F93B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F93C, 0x1F945, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F946, 0x1F946, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F947, 0x1F9FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1FA00, 0x1FA6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1FA70, 0x1FA7C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1FA7D, 0x1FA7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1FA80, 0x1FA88, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1FA89, 0x1FA8F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1FA90, 0x1FABD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1FABE, 0x1FABE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1FABF, 0x1FAC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1FAC6, 0x1FACD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1FACE, 0x1FADB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1FADC, 0x1FADF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1FAE0, 0x1FAE8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1FAE9, 0x1FAEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x1FAF0, 0x1FAF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiY,
		}, []*autogenTestCaseFunc{}},
		{0x1FAF9, 0x10FFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEmojiN,
			&autogenTestCasesFuncEmojiY,
		}},
	},
	{ // Epres
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
			&autogenTestCasesFuncEpresY,
		}},
		{0x00, 0x2319, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x231A, 0x231B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x231C, 0x23E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x23E9, 0x23EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x23ED, 0x23EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x23F0, 0x23F0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x23F1, 0x23F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x23F3, 0x23F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x23F4, 0x25FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x25FD, 0x25FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x25FF, 0x2613, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x2614, 0x2615, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x2616, 0x2647, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x2648, 0x2653, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x2654, 0x267E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x267F, 0x267F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x2680, 0x2692, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x2693, 0x2693, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x2694, 0x26A0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x26A1, 0x26A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x26A2, 0x26A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x26AA, 0x26AB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x26AC, 0x26BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x26BD, 0x26BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x26BF, 0x26C3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x26C4, 0x26C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x26C6, 0x26CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x26CE, 0x26CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x26CF, 0x26D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x26D4, 0x26D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x26D5, 0x26E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x26EA, 0x26EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x26EB, 0x26F1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x26F2, 0x26F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x26F4, 0x26F4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x26F5, 0x26F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x26F6, 0x26F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x26FA, 0x26FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x26FB, 0x26FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x26FD, 0x26FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x26FE, 0x2704, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x2705, 0x2705, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x2706, 0x2709, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x270A, 0x270B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x270C, 0x2727, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x2728, 0x2728, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x2729, 0x274B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x274C, 0x274C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x274D, 0x274D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x274E, 0x274E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x274F, 0x2752, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x2753, 0x2755, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x2756, 0x2756, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x2757, 0x2757, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x2758, 0x2794, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x2795, 0x2797, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x2798, 0x27AF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x27B0, 0x27B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x27B1, 0x27BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x27BF, 0x27BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x27C0, 0x2B1A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x2B1B, 0x2B1C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x2B1D, 0x2B4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x2B50, 0x2B50, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x2B51, 0x2B54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x2B55, 0x2B55, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x2B56, 0x1F003, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F004, 0x1F004, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F005, 0x1F0CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F0CF, 0x1F0CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F0D0, 0x1F18D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F18E, 0x1F18E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F18F, 0x1F190, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F191, 0x1F19A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F19B, 0x1F1E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F1E6, 0x1F1FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F200, 0x1F200, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F201, 0x1F201, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F202, 0x1F219, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F21A, 0x1F21A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F21B, 0x1F22E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F22F, 0x1F22F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F230, 0x1F231, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F232, 0x1F236, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F237, 0x1F237, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F238, 0x1F23A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F23B, 0x1F24F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F250, 0x1F251, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F252, 0x1F2FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F300, 0x1F320, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F321, 0x1F32C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F32D, 0x1F335, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F336, 0x1F336, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F337, 0x1F37C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F37D, 0x1F37D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F37E, 0x1F393, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F394, 0x1F39F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F3A0, 0x1F3CA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F3CB, 0x1F3CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F3CF, 0x1F3D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F3D4, 0x1F3DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F3E0, 0x1F3F0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F3F1, 0x1F3F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F3F4, 0x1F3F4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F3F5, 0x1F3F7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F3F8, 0x1F43E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F43F, 0x1F43F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F440, 0x1F440, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F441, 0x1F441, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F442, 0x1F4FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F4FD, 0x1F4FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F4FF, 0x1F53D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F53E, 0x1F54A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F54B, 0x1F54E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F54F, 0x1F54F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F550, 0x1F567, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F568, 0x1F579, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F57A, 0x1F57A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F57B, 0x1F594, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F595, 0x1F596, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F597, 0x1F5A3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5A4, 0x1F5A4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5A5, 0x1F5FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5FB, 0x1F64F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F650, 0x1F67F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F680, 0x1F6C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6C6, 0x1F6CB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6CC, 0x1F6CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6CD, 0x1F6CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6D0, 0x1F6D2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6D3, 0x1F6D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6D5, 0x1F6D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6D8, 0x1F6DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6DC, 0x1F6DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6E0, 0x1F6EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6EB, 0x1F6EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6ED, 0x1F6F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F6F4, 0x1F6FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F6FD, 0x1F7DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F7E0, 0x1F7EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F7EC, 0x1F7EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F7F0, 0x1F7F0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F7F1, 0x1F90B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F90C, 0x1F93A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F93B, 0x1F93B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F93C, 0x1F945, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1F946, 0x1F946, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1F947, 0x1F9FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1FA00, 0x1FA6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1FA70, 0x1FA7C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1FA7D, 0x1FA7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1FA80, 0x1FA88, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1FA89, 0x1FA8F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1FA90, 0x1FABD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1FABE, 0x1FABE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1FABF, 0x1FAC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1FAC6, 0x1FACD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1FACE, 0x1FADB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1FADC, 0x1FADF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1FAE0, 0x1FAE8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1FAE9, 0x1FAEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x1FAF0, 0x1FAF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresY,
		}, []*autogenTestCaseFunc{}},
		{0x1FAF9, 0x10FFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncEpresN,
			&autogenTestCasesFuncEpresY,
		}},
	},
	{ // Extpict
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
			&autogenTestCasesFuncExtpictY,
		}},
		{0x00, 0xA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0xA9, 0xA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0xAA, 0xAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0xAE, 0xAE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0xAF, 0x203B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x203C, 0x203C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x203D, 0x2048, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x2049, 0x2049, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x204A, 0x2121, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x2122, 0x2122, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x2123, 0x2138, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x2139, 0x2139, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x213A, 0x2193, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x2194, 0x2199, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x219A, 0x21A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x21A9, 0x21AA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x21AB, 0x2319, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x231A, 0x231B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x231C, 0x2327, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x2328, 0x2328, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x2329, 0x2387, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x2388, 0x2388, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x2389, 0x23CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x23CF, 0x23CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x23D0, 0x23E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x23E9, 0x23F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x23F4, 0x23F7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x23F8, 0x23FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x23FB, 0x24C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x24C2, 0x24C2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x24C3, 0x25A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x25AA, 0x25AB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x25AC, 0x25B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x25B6, 0x25B6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x25B7, 0x25BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x25C0, 0x25C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x25C1, 0x25FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x25FB, 0x25FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x25FF, 0x25FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x2600, 0x2605, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x2606, 0x2606, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x2607, 0x2612, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x2613, 0x2613, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x2614, 0x2685, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x2686, 0x268F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x2690, 0x2705, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x2706, 0x2707, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x2708, 0x2712, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x2713, 0x2713, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x2714, 0x2714, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictY,
		}, []*autogenTestCaseFunc{}},
		{0x2715, 0x2715, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x2B1D, 0x2B4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x1F240, 0x1F248, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncExtpictN,
			&autogenTestCasesFuncExtpictY,
		}},
	},
	{ // Ri
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncRiN,
			&autogenTestCasesFuncRiY,
		}},
		{0x00, 0x1F1E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncRiN,
		}, []*autogenTestCaseFunc{}},
		{0x1F1E6, 0x1F1FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncRiY,
		}, []*autogenTestCaseFunc{}},
		{0x1F200, 0x10FFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncRiN,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncRiN,
			&autogenTestCasesFuncRiY,
		}},
	},
}

// TestAutogenTestCases walks through autogenTestCases and tests the various
// test functions for all code points within the range.
func TestAutogenTestCases(t *testing.T) {
	for i := range autogenTestCases {
		for k := range autogenTestCases[i] {
			testCase := &autogenTestCases[i][k]
			for r := testCase.rMin; r <= testCase.rMax; r++ {
				for n := range testCase.funcsTrue {
					testFunc := testCase.funcsTrue[n]
					if !testFunc.f(r) {
						t.Errorf("AutogenTestCases_test: %s(0x%02X) = false, want true", testFunc.label, r)
					}
				}
				for n := range testCase.funcsFalse {
					testFunc := testCase.funcsFalse[n]
					if testFunc.f(r) {
						t.Errorf("AutogenTestCases_test: %s(0x%02X) = true, want false", testFunc.label, r)
					}
				}
			}
		}
	}
}
