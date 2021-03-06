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

/***************************************************************************
  This source code is not intended to be modified by biological entities.

  It is recommended that biological entities review the test cases described
  by the variable, autogenTestCases, to ensure they appear to describe the
  expected results for the corresponding ranges.
***************************************************************************/

package useg_test

import "testing"
import "github.com/orkvozku/go/uni/useg"

// autogenTestCaseFunc contains a function to test.
type autogenTestCaseFunc struct {
	label string
	f     func(rune) bool
}

// autogenTestCasesFunc... describes a function to test.
var (
	autogenTestCasesFuncExtpictN    = autogenTestCaseFunc{"HasExtpictN", useg.HasExtpictN}
	autogenTestCasesFuncExtpictY    = autogenTestCaseFunc{"HasExtpictY", useg.HasExtpictY}
	autogenTestCasesFuncGcbCn       = autogenTestCaseFunc{"HasGcbCn", useg.HasGcbCn}
	autogenTestCasesFuncGcbCr       = autogenTestCaseFunc{"HasGcbCr", useg.HasGcbCr}
	autogenTestCasesFuncGcbEx       = autogenTestCaseFunc{"HasGcbEx", useg.HasGcbEx}
	autogenTestCasesFuncGcbL        = autogenTestCaseFunc{"HasGcbL", useg.HasGcbL}
	autogenTestCasesFuncGcbLf       = autogenTestCaseFunc{"HasGcbLf", useg.HasGcbLf}
	autogenTestCasesFuncGcbLv       = autogenTestCaseFunc{"HasGcbLv", useg.HasGcbLv}
	autogenTestCasesFuncGcbLvt      = autogenTestCaseFunc{"HasGcbLvt", useg.HasGcbLvt}
	autogenTestCasesFuncGcbPp       = autogenTestCaseFunc{"HasGcbPp", useg.HasGcbPp}
	autogenTestCasesFuncGcbRi       = autogenTestCaseFunc{"HasGcbRi", useg.HasGcbRi}
	autogenTestCasesFuncGcbSm       = autogenTestCaseFunc{"HasGcbSm", useg.HasGcbSm}
	autogenTestCasesFuncGcbT        = autogenTestCaseFunc{"HasGcbT", useg.HasGcbT}
	autogenTestCasesFuncGcbV        = autogenTestCaseFunc{"HasGcbV", useg.HasGcbV}
	autogenTestCasesFuncGcbXx       = autogenTestCaseFunc{"HasGcbXx", useg.HasGcbXx}
	autogenTestCasesFuncGcbZwj      = autogenTestCaseFunc{"HasGcbZwj", useg.HasGcbZwj}
	autogenTestCasesFuncLbAi        = autogenTestCaseFunc{"HasLbAi", useg.HasLbAi}
	autogenTestCasesFuncLbAl        = autogenTestCaseFunc{"HasLbAl", useg.HasLbAl}
	autogenTestCasesFuncLbB2        = autogenTestCaseFunc{"HasLbB2", useg.HasLbB2}
	autogenTestCasesFuncLbBa        = autogenTestCaseFunc{"HasLbBa", useg.HasLbBa}
	autogenTestCasesFuncLbBb        = autogenTestCaseFunc{"HasLbBb", useg.HasLbBb}
	autogenTestCasesFuncLbBk        = autogenTestCaseFunc{"HasLbBk", useg.HasLbBk}
	autogenTestCasesFuncLbCb        = autogenTestCaseFunc{"HasLbCb", useg.HasLbCb}
	autogenTestCasesFuncLbCj        = autogenTestCaseFunc{"HasLbCj", useg.HasLbCj}
	autogenTestCasesFuncLbCl        = autogenTestCaseFunc{"HasLbCl", useg.HasLbCl}
	autogenTestCasesFuncLbCm        = autogenTestCaseFunc{"HasLbCm", useg.HasLbCm}
	autogenTestCasesFuncLbCp        = autogenTestCaseFunc{"HasLbCp", useg.HasLbCp}
	autogenTestCasesFuncLbCr        = autogenTestCaseFunc{"HasLbCr", useg.HasLbCr}
	autogenTestCasesFuncLbEb        = autogenTestCaseFunc{"HasLbEb", useg.HasLbEb}
	autogenTestCasesFuncLbEm        = autogenTestCaseFunc{"HasLbEm", useg.HasLbEm}
	autogenTestCasesFuncLbEx        = autogenTestCaseFunc{"HasLbEx", useg.HasLbEx}
	autogenTestCasesFuncLbGl        = autogenTestCaseFunc{"HasLbGl", useg.HasLbGl}
	autogenTestCasesFuncLbH2        = autogenTestCaseFunc{"HasLbH2", useg.HasLbH2}
	autogenTestCasesFuncLbH3        = autogenTestCaseFunc{"HasLbH3", useg.HasLbH3}
	autogenTestCasesFuncLbHl        = autogenTestCaseFunc{"HasLbHl", useg.HasLbHl}
	autogenTestCasesFuncLbHy        = autogenTestCaseFunc{"HasLbHy", useg.HasLbHy}
	autogenTestCasesFuncLbId        = autogenTestCaseFunc{"HasLbId", useg.HasLbId}
	autogenTestCasesFuncLbIn        = autogenTestCaseFunc{"HasLbIn", useg.HasLbIn}
	autogenTestCasesFuncLbIs        = autogenTestCaseFunc{"HasLbIs", useg.HasLbIs}
	autogenTestCasesFuncLbJl        = autogenTestCaseFunc{"HasLbJl", useg.HasLbJl}
	autogenTestCasesFuncLbJt        = autogenTestCaseFunc{"HasLbJt", useg.HasLbJt}
	autogenTestCasesFuncLbJv        = autogenTestCaseFunc{"HasLbJv", useg.HasLbJv}
	autogenTestCasesFuncLbLf        = autogenTestCaseFunc{"HasLbLf", useg.HasLbLf}
	autogenTestCasesFuncLbNl        = autogenTestCaseFunc{"HasLbNl", useg.HasLbNl}
	autogenTestCasesFuncLbNs        = autogenTestCaseFunc{"HasLbNs", useg.HasLbNs}
	autogenTestCasesFuncLbNu        = autogenTestCaseFunc{"HasLbNu", useg.HasLbNu}
	autogenTestCasesFuncLbOp        = autogenTestCaseFunc{"HasLbOp", useg.HasLbOp}
	autogenTestCasesFuncLbPo        = autogenTestCaseFunc{"HasLbPo", useg.HasLbPo}
	autogenTestCasesFuncLbPr        = autogenTestCaseFunc{"HasLbPr", useg.HasLbPr}
	autogenTestCasesFuncLbQu        = autogenTestCaseFunc{"HasLbQu", useg.HasLbQu}
	autogenTestCasesFuncLbRi        = autogenTestCaseFunc{"HasLbRi", useg.HasLbRi}
	autogenTestCasesFuncLbSa        = autogenTestCaseFunc{"HasLbSa", useg.HasLbSa}
	autogenTestCasesFuncLbSg        = autogenTestCaseFunc{"HasLbSg", useg.HasLbSg}
	autogenTestCasesFuncLbSp        = autogenTestCaseFunc{"HasLbSp", useg.HasLbSp}
	autogenTestCasesFuncLbSy        = autogenTestCaseFunc{"HasLbSy", useg.HasLbSy}
	autogenTestCasesFuncLbWj        = autogenTestCaseFunc{"HasLbWj", useg.HasLbWj}
	autogenTestCasesFuncLbXx        = autogenTestCaseFunc{"HasLbXx", useg.HasLbXx}
	autogenTestCasesFuncLbZw        = autogenTestCaseFunc{"HasLbZw", useg.HasLbZw}
	autogenTestCasesFuncLbZwj       = autogenTestCaseFunc{"HasLbZwj", useg.HasLbZwj}
	autogenTestCasesFuncSbAt        = autogenTestCaseFunc{"HasSbAt", useg.HasSbAt}
	autogenTestCasesFuncSbCl        = autogenTestCaseFunc{"HasSbCl", useg.HasSbCl}
	autogenTestCasesFuncSbCr        = autogenTestCaseFunc{"HasSbCr", useg.HasSbCr}
	autogenTestCasesFuncSbEx        = autogenTestCaseFunc{"HasSbEx", useg.HasSbEx}
	autogenTestCasesFuncSbFo        = autogenTestCaseFunc{"HasSbFo", useg.HasSbFo}
	autogenTestCasesFuncSbLe        = autogenTestCaseFunc{"HasSbLe", useg.HasSbLe}
	autogenTestCasesFuncSbLf        = autogenTestCaseFunc{"HasSbLf", useg.HasSbLf}
	autogenTestCasesFuncSbLo        = autogenTestCaseFunc{"HasSbLo", useg.HasSbLo}
	autogenTestCasesFuncSbNu        = autogenTestCaseFunc{"HasSbNu", useg.HasSbNu}
	autogenTestCasesFuncSbSc        = autogenTestCaseFunc{"HasSbSc", useg.HasSbSc}
	autogenTestCasesFuncSbSe        = autogenTestCaseFunc{"HasSbSe", useg.HasSbSe}
	autogenTestCasesFuncSbSp        = autogenTestCaseFunc{"HasSbSp", useg.HasSbSp}
	autogenTestCasesFuncSbSt        = autogenTestCaseFunc{"HasSbSt", useg.HasSbSt}
	autogenTestCasesFuncSbUp        = autogenTestCaseFunc{"HasSbUp", useg.HasSbUp}
	autogenTestCasesFuncSbXx        = autogenTestCaseFunc{"HasSbXx", useg.HasSbXx}
	autogenTestCasesFuncWbCr        = autogenTestCaseFunc{"HasWbCr", useg.HasWbCr}
	autogenTestCasesFuncWbDq        = autogenTestCaseFunc{"HasWbDq", useg.HasWbDq}
	autogenTestCasesFuncWbEx        = autogenTestCaseFunc{"HasWbEx", useg.HasWbEx}
	autogenTestCasesFuncWbExtend    = autogenTestCaseFunc{"HasWbExtend", useg.HasWbExtend}
	autogenTestCasesFuncWbFo        = autogenTestCaseFunc{"HasWbFo", useg.HasWbFo}
	autogenTestCasesFuncWbHl        = autogenTestCaseFunc{"HasWbHl", useg.HasWbHl}
	autogenTestCasesFuncWbKa        = autogenTestCaseFunc{"HasWbKa", useg.HasWbKa}
	autogenTestCasesFuncWbLe        = autogenTestCaseFunc{"HasWbLe", useg.HasWbLe}
	autogenTestCasesFuncWbLf        = autogenTestCaseFunc{"HasWbLf", useg.HasWbLf}
	autogenTestCasesFuncWbMb        = autogenTestCaseFunc{"HasWbMb", useg.HasWbMb}
	autogenTestCasesFuncWbMl        = autogenTestCaseFunc{"HasWbMl", useg.HasWbMl}
	autogenTestCasesFuncWbMn        = autogenTestCaseFunc{"HasWbMn", useg.HasWbMn}
	autogenTestCasesFuncWbNl        = autogenTestCaseFunc{"HasWbNl", useg.HasWbNl}
	autogenTestCasesFuncWbNu        = autogenTestCaseFunc{"HasWbNu", useg.HasWbNu}
	autogenTestCasesFuncWbRi        = autogenTestCaseFunc{"HasWbRi", useg.HasWbRi}
	autogenTestCasesFuncWbSq        = autogenTestCaseFunc{"HasWbSq", useg.HasWbSq}
	autogenTestCasesFuncWbWsegspace = autogenTestCaseFunc{"HasWbWsegspace", useg.HasWbWsegspace}
	autogenTestCasesFuncWbXx        = autogenTestCaseFunc{"HasWbXx", useg.HasWbXx}
	autogenTestCasesFuncWbZwj       = autogenTestCaseFunc{"HasWbZwj", useg.HasWbZwj}
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
	{ // Gcb
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbCn,
			&autogenTestCasesFuncGcbCr,
			&autogenTestCasesFuncGcbEx,
			&autogenTestCasesFuncGcbL,
			&autogenTestCasesFuncGcbLf,
			&autogenTestCasesFuncGcbLv,
			&autogenTestCasesFuncGcbLvt,
			&autogenTestCasesFuncGcbPp,
			&autogenTestCasesFuncGcbRi,
			&autogenTestCasesFuncGcbSm,
			&autogenTestCasesFuncGcbT,
			&autogenTestCasesFuncGcbV,
			&autogenTestCasesFuncGcbXx,
			&autogenTestCasesFuncGcbZwj,
		}},
		{0x00, 0x09, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbCn,
		}, []*autogenTestCaseFunc{}},
		{0x0A, 0x0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLf,
		}, []*autogenTestCaseFunc{}},
		{0x0B, 0x0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbCn,
		}, []*autogenTestCaseFunc{}},
		{0x0D, 0x0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbCr,
		}, []*autogenTestCaseFunc{}},
		{0x0E, 0x1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbCn,
		}, []*autogenTestCaseFunc{}},
		{0x20, 0x7E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x7F, 0x9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbCn,
		}, []*autogenTestCaseFunc{}},
		{0xA0, 0xAC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAD, 0xAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbCn,
		}, []*autogenTestCaseFunc{}},
		{0xAE, 0x2FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x300, 0x36F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x370, 0x482, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x483, 0x489, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x48A, 0x590, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x591, 0x5BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x5BE, 0x5BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5BF, 0x5BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x5C0, 0x5C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5C1, 0x5C2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x5C3, 0x5C3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5C4, 0x5C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x5C6, 0x5C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5C7, 0x5C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x5C8, 0x5FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x600, 0x605, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbPp,
		}, []*autogenTestCaseFunc{}},
		{0x606, 0x60F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x610, 0x61A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x61B, 0x61B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x61C, 0x61C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbCn,
		}, []*autogenTestCaseFunc{}},
		{0x61D, 0x64A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x64B, 0x65F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x660, 0x66F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x670, 0x670, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x671, 0x6D5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x6D6, 0x6DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x6DD, 0x6DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbPp,
		}, []*autogenTestCaseFunc{}},
		{0x6DE, 0x6DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x6DF, 0x6E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x6E5, 0x6E6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x6E7, 0x6E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x6E9, 0x6E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x6EA, 0x6ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x6EE, 0x70E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x70F, 0x70F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbPp,
		}, []*autogenTestCaseFunc{}},
		{0x710, 0x710, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x711, 0x711, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x712, 0x72F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x730, 0x74A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x74B, 0x7A5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x7A6, 0x7B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x7B1, 0x7EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x981, 0x981, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0xA81, 0xA82, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0xBBE, 0xBBE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0xCC3, 0xCC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbSm,
		}, []*autogenTestCaseFunc{}},
		{0xDD2, 0xDD4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0xF88, 0xF8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1715, 0x1715, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbSm,
		}, []*autogenTestCaseFunc{}},
		{0x1A19, 0x1A1A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbSm,
		}, []*autogenTestCaseFunc{}},
		{0x1BA8, 0x1BA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x2028, 0x202E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbCn,
		}, []*autogenTestCaseFunc{}},
		{0xA8B4, 0xA8C3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbSm,
		}, []*autogenTestCaseFunc{}},
		{0xAA7D, 0xAAAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAC8D, 0xACA7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xAEA1, 0xAEBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xB0B5, 0xB0CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xB2C9, 0xB2E3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xB4DD, 0xB4F7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xB6F1, 0xB70B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xB905, 0xB91F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xBB19, 0xBB33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xBD2D, 0xBD47, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xBF41, 0xBF5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xC155, 0xC16F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xC369, 0xC383, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xC57D, 0xC597, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xC791, 0xC7AB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xC9A5, 0xC9BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xCBB9, 0xCBD3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xCDCD, 0xCDE7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xCFE1, 0xCFFB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xD1F5, 0xD20F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xD409, 0xD423, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xD61D, 0xD637, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbLvt,
		}, []*autogenTestCaseFunc{}},
		{0xFEFF, 0xFEFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbCn,
		}, []*autogenTestCaseFunc{}},
		{0x11070, 0x11070, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x111CD, 0x111CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11375, 0x11434, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1163F, 0x11640, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x119DA, 0x119DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11CB5, 0x11CB6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbEx,
		}, []*autogenTestCaseFunc{}},
		{0x16FF0, 0x16FF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbSm,
		}, []*autogenTestCaseFunc{}},
		{0x1E007, 0x1E007, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbXx,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncGcbCn,
			&autogenTestCasesFuncGcbCr,
			&autogenTestCasesFuncGcbEx,
			&autogenTestCasesFuncGcbL,
			&autogenTestCasesFuncGcbLf,
			&autogenTestCasesFuncGcbLv,
			&autogenTestCasesFuncGcbLvt,
			&autogenTestCasesFuncGcbPp,
			&autogenTestCasesFuncGcbRi,
			&autogenTestCasesFuncGcbSm,
			&autogenTestCasesFuncGcbT,
			&autogenTestCasesFuncGcbV,
			&autogenTestCasesFuncGcbXx,
			&autogenTestCasesFuncGcbZwj,
		}},
	},
	{ // Lb
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
			&autogenTestCasesFuncLbAl,
			&autogenTestCasesFuncLbB2,
			&autogenTestCasesFuncLbBa,
			&autogenTestCasesFuncLbBb,
			&autogenTestCasesFuncLbBk,
			&autogenTestCasesFuncLbCb,
			&autogenTestCasesFuncLbCj,
			&autogenTestCasesFuncLbCl,
			&autogenTestCasesFuncLbCm,
			&autogenTestCasesFuncLbCp,
			&autogenTestCasesFuncLbCr,
			&autogenTestCasesFuncLbEb,
			&autogenTestCasesFuncLbEm,
			&autogenTestCasesFuncLbEx,
			&autogenTestCasesFuncLbGl,
			&autogenTestCasesFuncLbH2,
			&autogenTestCasesFuncLbH3,
			&autogenTestCasesFuncLbHl,
			&autogenTestCasesFuncLbHy,
			&autogenTestCasesFuncLbId,
			&autogenTestCasesFuncLbIn,
			&autogenTestCasesFuncLbIs,
			&autogenTestCasesFuncLbJl,
			&autogenTestCasesFuncLbJt,
			&autogenTestCasesFuncLbJv,
			&autogenTestCasesFuncLbLf,
			&autogenTestCasesFuncLbNl,
			&autogenTestCasesFuncLbNs,
			&autogenTestCasesFuncLbNu,
			&autogenTestCasesFuncLbOp,
			&autogenTestCasesFuncLbPo,
			&autogenTestCasesFuncLbPr,
			&autogenTestCasesFuncLbQu,
			&autogenTestCasesFuncLbRi,
			&autogenTestCasesFuncLbSa,
			&autogenTestCasesFuncLbSg,
			&autogenTestCasesFuncLbSp,
			&autogenTestCasesFuncLbSy,
			&autogenTestCasesFuncLbWj,
			&autogenTestCasesFuncLbXx,
			&autogenTestCasesFuncLbZw,
			&autogenTestCasesFuncLbZwj,
		}},
		{0x00, 0x08, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x09, 0x09, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x0A, 0x0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbLf,
		}, []*autogenTestCaseFunc{}},
		{0x0B, 0x0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBk,
		}, []*autogenTestCaseFunc{}},
		{0x0D, 0x0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCr,
		}, []*autogenTestCaseFunc{}},
		{0x0E, 0x1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x20, 0x20, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSp,
		}, []*autogenTestCaseFunc{}},
		{0x21, 0x21, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEx,
		}, []*autogenTestCaseFunc{}},
		{0x22, 0x22, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbQu,
		}, []*autogenTestCaseFunc{}},
		{0x23, 0x23, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x24, 0x24, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0x25, 0x25, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPo,
		}, []*autogenTestCaseFunc{}},
		{0x26, 0x26, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x27, 0x27, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbQu,
		}, []*autogenTestCaseFunc{}},
		{0x28, 0x28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x29, 0x29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCp,
		}, []*autogenTestCaseFunc{}},
		{0x2A, 0x2A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2B, 0x2B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0x2C, 0x2C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbIs,
		}, []*autogenTestCaseFunc{}},
		{0x2D, 0x2D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbHy,
		}, []*autogenTestCaseFunc{}},
		{0x2E, 0x2E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbIs,
		}, []*autogenTestCaseFunc{}},
		{0x2F, 0x2F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSy,
		}, []*autogenTestCaseFunc{}},
		{0x30, 0x39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x3A, 0x3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbIs,
		}, []*autogenTestCaseFunc{}},
		{0x3C, 0x3E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x3F, 0x3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEx,
		}, []*autogenTestCaseFunc{}},
		{0x40, 0x5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x5B, 0x5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x5C, 0x5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0x5D, 0x5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCp,
		}, []*autogenTestCaseFunc{}},
		{0x5E, 0x7A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x7B, 0x7B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x7C, 0x7C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x7D, 0x7D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0x7E, 0x7E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x7F, 0x84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x85, 0x85, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNl,
		}, []*autogenTestCaseFunc{}},
		{0x86, 0x9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA0, 0xA0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbGl,
		}, []*autogenTestCaseFunc{}},
		{0xA1, 0xA1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0xA2, 0xA2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPo,
		}, []*autogenTestCaseFunc{}},
		{0xA3, 0xA5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0xA6, 0xA6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA7, 0xA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0xA9, 0xA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xAA, 0xAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0xAB, 0xAB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbQu,
		}, []*autogenTestCaseFunc{}},
		{0xAC, 0xAC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xAD, 0xAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0xAE, 0xAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB0, 0xB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPo,
		}, []*autogenTestCaseFunc{}},
		{0xB1, 0xB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0xB2, 0xB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0xB4, 0xB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBb,
		}, []*autogenTestCaseFunc{}},
		{0xB5, 0xB5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB6, 0xBA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0xBB, 0xBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbQu,
		}, []*autogenTestCaseFunc{}},
		{0xBC, 0xBE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0xBF, 0xBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0xC0, 0xD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xD7, 0xD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0xD8, 0xF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xF7, 0xF7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0xF8, 0x2C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2C7, 0x2C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2C8, 0x2C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBb,
		}, []*autogenTestCaseFunc{}},
		{0x2C9, 0x2CB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2CC, 0x2CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBb,
		}, []*autogenTestCaseFunc{}},
		{0x2CD, 0x2CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2CE, 0x2CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2D0, 0x2D0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2D1, 0x2D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2D8, 0x2DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2DC, 0x2DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2DD, 0x2DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2DE, 0x2DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2DF, 0x2DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBb,
		}, []*autogenTestCaseFunc{}},
		{0x2E0, 0x2FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x300, 0x34E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x34F, 0x34F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbGl,
		}, []*autogenTestCaseFunc{}},
		{0x350, 0x35B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x35C, 0x362, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbGl,
		}, []*autogenTestCaseFunc{}},
		{0x363, 0x36F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x370, 0x377, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x378, 0x379, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x37A, 0x37D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x37E, 0x37E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbIs,
		}, []*autogenTestCaseFunc{}},
		{0x37F, 0x37F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x380, 0x383, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x384, 0x38A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x38B, 0x38B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x38C, 0x38C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x38D, 0x38D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x38E, 0x3A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x3A2, 0x3A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x3A3, 0x482, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x483, 0x489, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x48A, 0x52F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x530, 0x530, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x531, 0x556, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x557, 0x558, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x559, 0x588, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x589, 0x589, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbIs,
		}, []*autogenTestCaseFunc{}},
		{0x58A, 0x58A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x58B, 0x58C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x58D, 0x58E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x58F, 0x58F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0x590, 0x590, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x591, 0x5BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x5BE, 0x5BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x5BF, 0x5BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x5C0, 0x5C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x5C1, 0x5C2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x5C3, 0x5C3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x5C4, 0x5C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x5C6, 0x5C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEx,
		}, []*autogenTestCaseFunc{}},
		{0x5C7, 0x5C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x5C8, 0x5CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5D0, 0x5EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbHl,
		}, []*autogenTestCaseFunc{}},
		{0x5EB, 0x5EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5EF, 0x5F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbHl,
		}, []*autogenTestCaseFunc{}},
		{0x5F3, 0x5F4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x5F5, 0x5FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x600, 0x608, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x609, 0x60B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPo,
		}, []*autogenTestCaseFunc{}},
		{0x60C, 0x60D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbIs,
		}, []*autogenTestCaseFunc{}},
		{0x60E, 0x60F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x610, 0x61A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x61B, 0x61B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEx,
		}, []*autogenTestCaseFunc{}},
		{0x61C, 0x61C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x61D, 0x61F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEx,
		}, []*autogenTestCaseFunc{}},
		{0x620, 0x64A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x64B, 0x65F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x660, 0x669, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x66A, 0x66A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPo,
		}, []*autogenTestCaseFunc{}},
		{0x66B, 0x66C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x66D, 0x66F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x670, 0x670, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x671, 0x6D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x6D4, 0x6D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEx,
		}, []*autogenTestCaseFunc{}},
		{0x6D5, 0x6D5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x6D6, 0x6DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x6DD, 0x6DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x6DF, 0x6E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x6E5, 0x6E6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x6E7, 0x6E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x6E9, 0x6E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x6EA, 0x6ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x6EE, 0x6EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x6F0, 0x6F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x6FA, 0x70D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x70E, 0x70E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x70F, 0x710, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x711, 0x711, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x712, 0x72F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x730, 0x74A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x74B, 0x74C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x74D, 0x7A5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x7A6, 0x7B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x7B1, 0x7B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x7B2, 0x7BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x7C0, 0x7C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x7CA, 0x7EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x7EB, 0x7F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x7F4, 0x7F7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x7F8, 0x7F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbIs,
		}, []*autogenTestCaseFunc{}},
		{0x7F9, 0x7F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEx,
		}, []*autogenTestCaseFunc{}},
		{0x7FA, 0x7FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x7FB, 0x7FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x7FD, 0x7FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x7FE, 0x7FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0x800, 0x815, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x816, 0x819, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x81A, 0x81A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x81B, 0x823, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x824, 0x824, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x825, 0x827, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x828, 0x828, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x829, 0x82D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x82E, 0x82F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x830, 0x83E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x83F, 0x83F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x840, 0x858, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x859, 0x85B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x85C, 0x85D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x85E, 0x85E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x85F, 0x85F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x860, 0x86A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x86B, 0x86F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x870, 0x88E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x88F, 0x88F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x890, 0x891, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x892, 0x897, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x898, 0x89F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x8A0, 0x8C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x8CA, 0x8E1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x8E2, 0x8E2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x8E3, 0x903, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x904, 0x939, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x93A, 0x93C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x93D, 0x93D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x93E, 0x94F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x950, 0x950, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x951, 0x957, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x958, 0x961, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x962, 0x963, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x964, 0x965, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x966, 0x96F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x970, 0x980, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x981, 0x983, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x984, 0x984, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x985, 0x98C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x98D, 0x98E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x98F, 0x990, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x991, 0x992, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x993, 0x9A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x9A9, 0x9A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9AA, 0x9B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x9B1, 0x9B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9B2, 0x9B2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x9B3, 0x9B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9B6, 0x9B9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x9BA, 0x9BB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9BC, 0x9BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x9BD, 0x9BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x9BE, 0x9C4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x9C5, 0x9C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9C7, 0x9C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x9C9, 0x9CA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9CB, 0x9CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x9CE, 0x9CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x9CF, 0x9D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9D7, 0x9D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x9D8, 0x9DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9DC, 0x9DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x9DE, 0x9DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9DF, 0x9E1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x9E2, 0x9E3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x9E4, 0x9E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9E6, 0x9EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x9F0, 0x9F1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x9F2, 0x9F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPo,
		}, []*autogenTestCaseFunc{}},
		{0x9F4, 0x9F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x9F9, 0x9F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPo,
		}, []*autogenTestCaseFunc{}},
		{0x9FA, 0x9FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x9FB, 0x9FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0x9FC, 0x9FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x9FE, 0x9FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x9FF, 0xA00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA01, 0xA03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA04, 0xA04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA05, 0xA0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA0B, 0xA0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA0F, 0xA10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA11, 0xA12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA13, 0xA28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA29, 0xA29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA2A, 0xA30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA31, 0xA31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA32, 0xA33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA34, 0xA34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA35, 0xA36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA37, 0xA37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA38, 0xA39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA3A, 0xA3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA3C, 0xA3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA3D, 0xA3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA3E, 0xA42, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA43, 0xA46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA47, 0xA48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA49, 0xA4A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA4B, 0xA4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA4E, 0xA50, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA51, 0xA51, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA52, 0xA58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA59, 0xA5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA5D, 0xA5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA5E, 0xA5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA5F, 0xA65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA66, 0xA6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0xA70, 0xA71, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA72, 0xA74, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA75, 0xA75, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA76, 0xA76, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA77, 0xA80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA81, 0xA83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA84, 0xA84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA85, 0xA8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA8E, 0xA8E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA8F, 0xA91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA92, 0xA92, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA93, 0xAA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xAA9, 0xAA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAAA, 0xAB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xAB1, 0xAB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAB2, 0xAB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xAB4, 0xAB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAB5, 0xAB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xABA, 0xABB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xABC, 0xABC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xABD, 0xABD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xABE, 0xAC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xAC6, 0xAC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAC7, 0xAC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xACA, 0xACA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xACB, 0xACD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xACE, 0xACF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAD0, 0xAD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xAD1, 0xADF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAE0, 0xAE1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xAE2, 0xAE3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xAE4, 0xAE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAE6, 0xAEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0xAF0, 0xAF0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xAF1, 0xAF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0xAF2, 0xAF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAF9, 0xAF9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xAFA, 0xAFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xB00, 0xB00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB01, 0xB03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xB04, 0xB04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB05, 0xB0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB0D, 0xB0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB0F, 0xB10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB11, 0xB12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB13, 0xB28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB29, 0xB29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB2A, 0xB30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB31, 0xB31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB32, 0xB33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB34, 0xB34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB35, 0xB39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB3A, 0xB3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB3C, 0xB3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xB3D, 0xB3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB3E, 0xB44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xB45, 0xB46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB47, 0xB48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xB49, 0xB4A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB4B, 0xB4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xB4E, 0xB54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB55, 0xB57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xB58, 0xB5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB5C, 0xB5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB5E, 0xB5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB5F, 0xB61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB62, 0xB63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xB64, 0xB65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB66, 0xB6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0xB70, 0xB77, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB78, 0xB81, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB82, 0xB82, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xB83, 0xB83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB84, 0xB84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB85, 0xB8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB8B, 0xB8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB8E, 0xB90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB91, 0xB91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB92, 0xB95, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB96, 0xB98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB99, 0xB9A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB9B, 0xB9B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB9C, 0xB9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xB9D, 0xB9D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB9E, 0xB9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xBA0, 0xBA2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBA3, 0xBA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xBA5, 0xBA7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBA8, 0xBAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xBAB, 0xBAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBAE, 0xBB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xBBA, 0xBBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBBE, 0xBC2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xBC3, 0xBC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBC6, 0xBC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xBC9, 0xBC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBCA, 0xBCD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xBCE, 0xBCF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBD0, 0xBD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xBD1, 0xBD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBD7, 0xBD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xBD8, 0xBE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBE6, 0xBEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0xBF0, 0xBF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xBF9, 0xBF9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0xBFA, 0xBFA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xBFB, 0xBFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC00, 0xC04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xC05, 0xC0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xC0D, 0xC0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC0E, 0xC10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xC11, 0xC11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC12, 0xC28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xC29, 0xC29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC2A, 0xC39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xC3A, 0xC3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC3C, 0xC3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xC3D, 0xC3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xC3E, 0xC44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xC45, 0xC45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC46, 0xC48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xC49, 0xC49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC4A, 0xC4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xC4E, 0xC54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC55, 0xC56, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xC57, 0xC57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC58, 0xC5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xC5B, 0xC5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC5D, 0xC5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xC5E, 0xC5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC60, 0xC61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xC62, 0xC63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xC64, 0xC65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC66, 0xC6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0xC70, 0xC76, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC77, 0xC77, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBb,
		}, []*autogenTestCaseFunc{}},
		{0xC78, 0xC80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xC81, 0xC83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xC84, 0xC84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBb,
		}, []*autogenTestCaseFunc{}},
		{0xC85, 0xC8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xC8D, 0xC8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC8E, 0xC90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xC91, 0xC91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC92, 0xCA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xCA9, 0xCA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCAA, 0xCB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xCB4, 0xCB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCB5, 0xCB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xCBA, 0xCBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCBC, 0xCBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xCBD, 0xCBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xCBE, 0xCC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xCC5, 0xCC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCC6, 0xCC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xCC9, 0xCC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCCA, 0xCCD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xCCE, 0xCD4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCD5, 0xCD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xCD7, 0xCDC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCDD, 0xCDE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xCDF, 0xCDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCE0, 0xCE1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xCE2, 0xCE3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xCE4, 0xCE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCE6, 0xCEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0xCF0, 0xCF0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCF1, 0xCF2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xCF3, 0xCFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD00, 0xD03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xD04, 0xD0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xD0D, 0xD0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD0E, 0xD10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xD11, 0xD11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD12, 0xD3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xD3B, 0xD3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xD3D, 0xD3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xD3E, 0xD44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xD45, 0xD45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD46, 0xD48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xD49, 0xD49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD4A, 0xD4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xD4E, 0xD4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xD50, 0xD53, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD54, 0xD56, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xD57, 0xD57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xD58, 0xD61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xD62, 0xD63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xD64, 0xD65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD66, 0xD6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0xD70, 0xD78, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xD79, 0xD79, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPo,
		}, []*autogenTestCaseFunc{}},
		{0xD7A, 0xD7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xD80, 0xD80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD81, 0xD83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xD84, 0xD84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD85, 0xD96, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xD97, 0xD99, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD9A, 0xDB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xDB2, 0xDB2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDB3, 0xDBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xDBC, 0xDBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDBD, 0xDBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xDBE, 0xDBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDC0, 0xDC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xDC7, 0xDC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDCA, 0xDCA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xDCB, 0xDCE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDCF, 0xDD4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xDD5, 0xDD5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDD6, 0xDD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xDD7, 0xDD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDD8, 0xDDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xDE0, 0xDE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDE6, 0xDEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0xDF0, 0xDF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDF2, 0xDF3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xDF4, 0xDF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xDF5, 0xE00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xE01, 0xE3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSa,
		}, []*autogenTestCaseFunc{}},
		{0xE3B, 0xE3E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xE3F, 0xE3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0xE84, 0xE84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSa,
		}, []*autogenTestCaseFunc{}},
		{0xEA7, 0xEBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSa,
		}, []*autogenTestCaseFunc{}},
		{0xED0, 0xED9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0xF08, 0xF08, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbGl,
		}, []*autogenTestCaseFunc{}},
		{0xF15, 0xF17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xF37, 0xF37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xF40, 0xF47, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xF86, 0xF87, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xFC6, 0xFC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xFD9, 0xFDA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbGl,
		}, []*autogenTestCaseFunc{}},
		{0x109A, 0x109F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSa,
		}, []*autogenTestCaseFunc{}},
		{0x1100, 0x115F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbJl,
		}, []*autogenTestCaseFunc{}},
		{0x1257, 0x1257, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x128E, 0x128F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x12C1, 0x12C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1316, 0x1317, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1380, 0x1399, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1680, 0x1680, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x16F9, 0x16FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1740, 0x1751, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1774, 0x177F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x17DB, 0x17DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0x1802, 0x1803, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEx,
		}, []*autogenTestCaseFunc{}},
		{0x180F, 0x180F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x18A9, 0x18A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x192C, 0x192F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x196E, 0x196F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x19DA, 0x19DA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSa,
		}, []*autogenTestCaseFunc{}},
		{0x1A5F, 0x1A5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1AA0, 0x1AAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSa,
		}, []*autogenTestCaseFunc{}},
		{0x1B4D, 0x1B4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1B7D, 0x1B7E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x1BE6, 0x1BF3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x1C4D, 0x1C4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1CBD, 0x1CC7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1CF4, 0x1CF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x1F16, 0x1F17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1F58, 0x1F58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1F7E, 0x1F7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1FDC, 0x1FDC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1FFF, 0x1FFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2010, 0x2010, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x201B, 0x201D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbQu,
		}, []*autogenTestCaseFunc{}},
		{0x202A, 0x202E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x2044, 0x2044, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbIs,
		}, []*autogenTestCaseFunc{}},
		{0x205C, 0x205C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2074, 0x2074, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x208D, 0x208D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x20B6, 0x20B6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPo,
		}, []*autogenTestCaseFunc{}},
		{0x20D0, 0x20F0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x210A, 0x2112, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x212C, 0x2153, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x216C, 0x216F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x21D2, 0x21D2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2207, 0x2208, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2214, 0x2214, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2224, 0x2224, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2238, 0x223B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2253, 0x225F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2270, 0x2281, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x229A, 0x22A4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2309, 0x2309, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0x2329, 0x2329, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x2460, 0x24FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2592, 0x2595, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x25B6, 0x25B7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x25CB, 0x25CB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2600, 0x2603, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x2614, 0x2615, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x2639, 0x263B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x2663, 0x2665, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x266F, 0x266F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x26CD, 0x26CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x26DC, 0x26DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x26EB, 0x26F0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2705, 0x2707, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2762, 0x2763, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEx,
		}, []*autogenTestCaseFunc{}},
		{0x276D, 0x276D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0x2775, 0x2775, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0x27E8, 0x27E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x27F0, 0x2982, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x298A, 0x298A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0x2992, 0x2992, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0x29D8, 0x29D8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x2B55, 0x2B59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2CF4, 0x2CF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2D27, 0x2D27, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2D71, 0x2D7E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2DB0, 0x2DB6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2DD0, 0x2DD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2E17, 0x2E17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x2E23, 0x2E23, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0x2E2E, 0x2E2E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEx,
		}, []*autogenTestCaseFunc{}},
		{0x2E3F, 0x2E3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2E50, 0x2E52, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2E5B, 0x2E5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x2F00, 0x2FD5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x3006, 0x3007, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x300F, 0x300F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0x3018, 0x3018, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x302A, 0x302F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x3042, 0x3042, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x304A, 0x3062, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x3088, 0x308D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x30A0, 0x30A0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNs,
		}, []*autogenTestCaseFunc{}},
		{0x30A8, 0x30A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x30E6, 0x30E6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x30FC, 0x30FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCj,
		}, []*autogenTestCaseFunc{}},
		{0x3190, 0x31E3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x4DC0, 0x4DFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA4FE, 0xA4FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0xA62C, 0xA63F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA6F0, 0xA6F1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA7D3, 0xA7D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA807, 0xA80A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA838, 0xA838, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPo,
		}, []*autogenTestCaseFunc{}},
		{0xA882, 0xA8B3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA8FC, 0xA8FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBb,
		}, []*autogenTestCaseFunc{}},
		{0xA947, 0xA953, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA9C1, 0xA9C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA9E0, 0xA9EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSa,
		}, []*autogenTestCaseFunc{}},
		{0xAA43, 0xAA43, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xAA60, 0xAAC2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSa,
		}, []*autogenTestCaseFunc{}},
		{0xAAF7, 0xAB00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAB27, 0xAB27, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xABEC, 0xABED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xAC38, 0xAC38, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xACA8, 0xACA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xAD18, 0xAD18, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xAD88, 0xAD88, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xADF8, 0xADF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xAE68, 0xAE68, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xAED8, 0xAED8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xAF48, 0xAF48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xAFB8, 0xAFB8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB028, 0xB028, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB098, 0xB098, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB108, 0xB108, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB178, 0xB178, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB1E8, 0xB1E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB258, 0xB258, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB2C8, 0xB2C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB338, 0xB338, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB3A8, 0xB3A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB418, 0xB418, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB488, 0xB488, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB4F8, 0xB4F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB568, 0xB568, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB5D8, 0xB5D8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB648, 0xB648, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB6B8, 0xB6B8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB728, 0xB728, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB798, 0xB798, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB808, 0xB808, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB878, 0xB878, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB8E8, 0xB8E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB958, 0xB958, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB9C8, 0xB9C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBA38, 0xBA38, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBAA8, 0xBAA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBB18, 0xBB18, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBB88, 0xBB88, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBBF8, 0xBBF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBC68, 0xBC68, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBCD8, 0xBCD8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBD48, 0xBD48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBDB8, 0xBDB8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBE28, 0xBE28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBE98, 0xBE98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBF08, 0xBF08, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBF78, 0xBF78, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBFE8, 0xBFE8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC058, 0xC058, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC0C8, 0xC0C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC138, 0xC138, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC1A8, 0xC1A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC218, 0xC218, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC288, 0xC288, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC2F8, 0xC2F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC368, 0xC368, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC3D8, 0xC3D8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC448, 0xC448, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC4B8, 0xC4B8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC528, 0xC528, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC598, 0xC598, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC608, 0xC608, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC678, 0xC678, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC6E8, 0xC6E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC758, 0xC758, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC7C8, 0xC7C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC838, 0xC838, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC8A8, 0xC8A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC918, 0xC918, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC988, 0xC988, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC9F8, 0xC9F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCA68, 0xCA68, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCAD8, 0xCAD8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCB48, 0xCB48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCBB8, 0xCBB8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCC28, 0xCC28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCC98, 0xCC98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCD08, 0xCD08, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCD78, 0xCD78, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCDE8, 0xCDE8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCE58, 0xCE58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCEC8, 0xCEC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCF38, 0xCF38, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCFA8, 0xCFA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD018, 0xD018, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD088, 0xD088, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD0F8, 0xD0F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD168, 0xD168, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD1D8, 0xD1D8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD248, 0xD248, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD2B8, 0xD2B8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD328, 0xD328, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD398, 0xD398, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD408, 0xD408, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD478, 0xD478, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD4E8, 0xD4E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD558, 0xD558, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD5C8, 0xD5C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD638, 0xD638, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD6A8, 0xD6A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD718, 0xD718, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD788, 0xD788, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xE000, 0xF8FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFB1F, 0xFB28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbHl,
		}, []*autogenTestCaseFunc{}},
		{0xFB40, 0xFB41, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbHl,
		}, []*autogenTestCaseFunc{}},
		{0xFD3E, 0xFD3E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0xFDF0, 0xFDFB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xFE17, 0xFE17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0xFE37, 0xFE37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0xFE3F, 0xFE3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0xFE48, 0xFE48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0xFE58, 0xFE58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0xFE67, 0xFE67, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFE76, 0xFEFC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xFF06, 0xFF07, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0xFF1A, 0xFF1B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNs,
		}, []*autogenTestCaseFunc{}},
		{0xFF5B, 0xFF5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0xFF65, 0xFF65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNs,
		}, []*autogenTestCaseFunc{}},
		{0xFFC8, 0xFFC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFFE1, 0xFFE1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0xFFFD, 0xFFFD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x1003C, 0x1003D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10100, 0x10102, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x101A0, 0x101A0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x102D1, 0x102DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10350, 0x10375, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x103C8, 0x103CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x104B0, 0x104D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1056F, 0x1057A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10597, 0x105A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10600, 0x10736, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10787, 0x107B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1080A, 0x10835, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10857, 0x10857, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x108F6, 0x108FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10980, 0x109B7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10A07, 0x10A0B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10A38, 0x10A3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x10A60, 0x10A9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10AF7, 0x10AFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10B78, 0x10B91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10C80, 0x10CB2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10D3A, 0x10E5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10EB0, 0x10EB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10F70, 0x10F81, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11000, 0x11002, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x11070, 0x11070, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x110BB, 0x110BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x110F0, 0x110F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x11144, 0x11144, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11176, 0x11176, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x111C8, 0x111C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x111DD, 0x111DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x11238, 0x11239, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x11288, 0x11288, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x112AA, 0x112AF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11305, 0x1130C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11332, 0x11333, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11347, 0x11348, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x1135D, 0x11361, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11435, 0x11446, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x1145E, 0x1145E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x114DA, 0x1157F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x115C6, 0x115C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11643, 0x11644, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x116B8, 0x116B9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11730, 0x11739, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x1183C, 0x1189F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1190A, 0x1190B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11937, 0x11938, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x11947, 0x1194F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x119DA, 0x119E0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x11A0B, 0x11A32, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11A46, 0x11A46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11A9D, 0x11A9D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11C0A, 0x11C2E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11C5A, 0x11C6C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11CA9, 0x11CB6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x11D37, 0x11D39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11D48, 0x11D4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11D8A, 0x11D8E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x11DAA, 0x11EDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11FDD, 0x11FE0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPo,
		}, []*autogenTestCaseFunc{}},
		{0x12470, 0x12474, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x1325B, 0x1325D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0x1328A, 0x13378, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x13439, 0x143FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16A40, 0x16A5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x16ACA, 0x16ACF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16B37, 0x16B39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x16B62, 0x16B62, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16E9B, 0x16EFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16F93, 0x16F9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x187F8, 0x187FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1AFF5, 0x1AFFB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1B164, 0x1B167, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCj,
		}, []*autogenTestCaseFunc{}},
		{0x1BC80, 0x1BC88, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1BCA4, 0x1CEFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D0F6, 0x1D0FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D185, 0x1D18B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x1D246, 0x1D2DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D455, 0x1D455, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D4A7, 0x1D4A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D4C4, 0x1D4C4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D51D, 0x1D51D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D547, 0x1D549, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D800, 0x1D9FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1DA85, 0x1DA86, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1DF00, 0x1DF1E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1E023, 0x1E024, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x1E13E, 0x1E13F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E2C0, 0x1E2EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1E7E8, 0x1E7EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1E8C7, 0x1E8CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1E95A, 0x1E95D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1ECB5, 0x1ED00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE23, 0x1EE23, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE38, 0x1EE38, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE48, 0x1EE48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE53, 0x1EE53, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE5C, 0x1EE5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE65, 0x1EE66, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE7D, 0x1EE7D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EEA4, 0x1EEA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1F100, 0x1F10C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x1F1AD, 0x1F1E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F3B7, 0x1F3BB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F3CD, 0x1F3FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F479, 0x1F47B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F490, 0x1F490, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F4A5, 0x1F4A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F507, 0x1F516, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F57B, 0x1F58F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F5FA, 0x1F644, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F680, 0x1F6A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F6CD, 0x1F6FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F848, 0x1F84F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F90C, 0x1F90C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEb,
		}, []*autogenTestCaseFunc{}},
		{0x1F930, 0x1F939, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEb,
		}, []*autogenTestCaseFunc{}},
		{0x1F9B8, 0x1F9B9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEb,
		}, []*autogenTestCaseFunc{}},
		{0x1FA00, 0x1FA53, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1FB94, 0x1FBCA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x30000, 0x3FFFD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x100000, 0x10FFFD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
			&autogenTestCasesFuncLbAl,
			&autogenTestCasesFuncLbB2,
			&autogenTestCasesFuncLbBa,
			&autogenTestCasesFuncLbBb,
			&autogenTestCasesFuncLbBk,
			&autogenTestCasesFuncLbCb,
			&autogenTestCasesFuncLbCj,
			&autogenTestCasesFuncLbCl,
			&autogenTestCasesFuncLbCm,
			&autogenTestCasesFuncLbCp,
			&autogenTestCasesFuncLbCr,
			&autogenTestCasesFuncLbEb,
			&autogenTestCasesFuncLbEm,
			&autogenTestCasesFuncLbEx,
			&autogenTestCasesFuncLbGl,
			&autogenTestCasesFuncLbH2,
			&autogenTestCasesFuncLbH3,
			&autogenTestCasesFuncLbHl,
			&autogenTestCasesFuncLbHy,
			&autogenTestCasesFuncLbId,
			&autogenTestCasesFuncLbIn,
			&autogenTestCasesFuncLbIs,
			&autogenTestCasesFuncLbJl,
			&autogenTestCasesFuncLbJt,
			&autogenTestCasesFuncLbJv,
			&autogenTestCasesFuncLbLf,
			&autogenTestCasesFuncLbNl,
			&autogenTestCasesFuncLbNs,
			&autogenTestCasesFuncLbNu,
			&autogenTestCasesFuncLbOp,
			&autogenTestCasesFuncLbPo,
			&autogenTestCasesFuncLbPr,
			&autogenTestCasesFuncLbQu,
			&autogenTestCasesFuncLbRi,
			&autogenTestCasesFuncLbSa,
			&autogenTestCasesFuncLbSg,
			&autogenTestCasesFuncLbSp,
			&autogenTestCasesFuncLbSy,
			&autogenTestCasesFuncLbWj,
			&autogenTestCasesFuncLbXx,
			&autogenTestCasesFuncLbZw,
			&autogenTestCasesFuncLbZwj,
		}},
	},
	{ // Sb
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbAt,
			&autogenTestCasesFuncSbCl,
			&autogenTestCasesFuncSbCr,
			&autogenTestCasesFuncSbEx,
			&autogenTestCasesFuncSbFo,
			&autogenTestCasesFuncSbLe,
			&autogenTestCasesFuncSbLf,
			&autogenTestCasesFuncSbLo,
			&autogenTestCasesFuncSbNu,
			&autogenTestCasesFuncSbSc,
			&autogenTestCasesFuncSbSe,
			&autogenTestCasesFuncSbSp,
			&autogenTestCasesFuncSbSt,
			&autogenTestCasesFuncSbUp,
			&autogenTestCasesFuncSbXx,
		}},
		{0x00, 0x08, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x09, 0x09, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSp,
		}, []*autogenTestCaseFunc{}},
		{0x0A, 0x0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLf,
		}, []*autogenTestCaseFunc{}},
		{0x0B, 0x0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSp,
		}, []*autogenTestCaseFunc{}},
		{0x0D, 0x0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCr,
		}, []*autogenTestCaseFunc{}},
		{0x0E, 0x1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x20, 0x20, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSp,
		}, []*autogenTestCaseFunc{}},
		{0x21, 0x21, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0x22, 0x22, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0x23, 0x26, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x27, 0x29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0x2A, 0x2B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2C, 0x2D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSc,
		}, []*autogenTestCaseFunc{}},
		{0x2E, 0x2E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbAt,
		}, []*autogenTestCaseFunc{}},
		{0x2F, 0x2F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x30, 0x39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0x3A, 0x3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSc,
		}, []*autogenTestCaseFunc{}},
		{0x3B, 0x3E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x3F, 0x3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0x40, 0x40, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x41, 0x5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x5B, 0x5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0x5C, 0x5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5D, 0x5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0x5E, 0x60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x61, 0x7A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x7B, 0x7B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0x7C, 0x7C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x7D, 0x7D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0x7E, 0x84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x85, 0x85, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSe,
		}, []*autogenTestCaseFunc{}},
		{0x86, 0x9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA0, 0xA0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSp,
		}, []*autogenTestCaseFunc{}},
		{0xA1, 0xA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAA, 0xAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xAB, 0xAB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0xAC, 0xAC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAD, 0xAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbFo,
		}, []*autogenTestCaseFunc{}},
		{0xAE, 0xB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB5, 0xB5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xB6, 0xB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBA, 0xBA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xBB, 0xBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0xBC, 0xBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC0, 0xD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xD7, 0xD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD8, 0xDE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xDF, 0xF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xF7, 0xF7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF8, 0xFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x100, 0x100, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x101, 0x101, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x102, 0x102, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x103, 0x103, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x104, 0x104, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x105, 0x105, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x106, 0x106, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x107, 0x107, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x108, 0x108, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x109, 0x109, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x10A, 0x10A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x10B, 0x10B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x10C, 0x10C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x10D, 0x10D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x10E, 0x10E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x10F, 0x10F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x110, 0x110, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x111, 0x111, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x112, 0x112, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x113, 0x113, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x114, 0x114, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x115, 0x115, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x116, 0x116, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x117, 0x117, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x118, 0x118, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x119, 0x119, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x11A, 0x11A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x11B, 0x11B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x11C, 0x11C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x11D, 0x11D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x11E, 0x11E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x11F, 0x11F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x120, 0x120, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x121, 0x121, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x122, 0x122, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x123, 0x123, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x124, 0x124, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x125, 0x125, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x126, 0x126, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x127, 0x127, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x128, 0x128, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x129, 0x129, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x12A, 0x12A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x12B, 0x12B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x12C, 0x12C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x12D, 0x12D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x12E, 0x12E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x12F, 0x12F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x130, 0x130, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x131, 0x131, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x132, 0x132, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x133, 0x133, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x134, 0x134, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x135, 0x135, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x136, 0x136, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x137, 0x138, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x139, 0x139, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x13A, 0x13A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x13B, 0x13B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x13C, 0x13C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x13D, 0x13D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x13E, 0x13E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x13F, 0x13F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x140, 0x140, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x141, 0x141, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x142, 0x142, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x143, 0x143, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x144, 0x144, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x145, 0x145, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x146, 0x146, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x147, 0x147, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x148, 0x149, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x14A, 0x14A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x14B, 0x14B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x14C, 0x14C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x14D, 0x14D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x14E, 0x14E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x14F, 0x14F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x150, 0x150, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x151, 0x151, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x152, 0x152, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x153, 0x153, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x154, 0x154, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x155, 0x155, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x156, 0x156, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x157, 0x157, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x158, 0x158, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x159, 0x159, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x15A, 0x15A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x15B, 0x15B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x15C, 0x15C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x15D, 0x15D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x15E, 0x15E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x15F, 0x15F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x160, 0x160, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x161, 0x161, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x162, 0x162, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x163, 0x163, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x164, 0x164, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x165, 0x165, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x166, 0x166, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x167, 0x167, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x168, 0x168, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x169, 0x169, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x16A, 0x16A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x16B, 0x16B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x16C, 0x16C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x16D, 0x16D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x16E, 0x16E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x16F, 0x16F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x170, 0x170, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x171, 0x171, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x172, 0x172, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x173, 0x173, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x174, 0x174, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x175, 0x175, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x176, 0x176, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x177, 0x177, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x178, 0x179, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x17A, 0x17A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x17B, 0x17B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x17C, 0x17C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x17D, 0x17D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x17E, 0x180, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x181, 0x182, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x183, 0x183, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x184, 0x184, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x185, 0x185, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x186, 0x187, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x188, 0x188, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x189, 0x18B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x18C, 0x18D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x18E, 0x191, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x192, 0x192, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x193, 0x194, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x195, 0x195, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x196, 0x198, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x199, 0x19B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x19C, 0x19D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x19E, 0x19E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x19F, 0x1A0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1A1, 0x1A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1A2, 0x1A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1A3, 0x1A3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1A4, 0x1A4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1A5, 0x1A5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1A6, 0x1A7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1A8, 0x1A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1A9, 0x1A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1AA, 0x1AB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1AC, 0x1AC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1AD, 0x1AD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1AE, 0x1AF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1B0, 0x1B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1B1, 0x1B3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1B4, 0x1B4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1B5, 0x1B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1B6, 0x1B6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1B7, 0x1B8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1B9, 0x1BA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1BB, 0x1BB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1BC, 0x1BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1BD, 0x1BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1C0, 0x1C3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1C4, 0x1C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1C6, 0x1C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1C7, 0x1C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1C9, 0x1C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1CA, 0x1CB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1CC, 0x1CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1CD, 0x1CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1CE, 0x1CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1CF, 0x1CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1D0, 0x1D0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1D1, 0x1D1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1D2, 0x1D2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1D3, 0x1D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1D4, 0x1D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1D5, 0x1D5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1D6, 0x1D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1D7, 0x1D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1D8, 0x1D8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1D9, 0x1D9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1DA, 0x1DA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1DB, 0x1DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1DC, 0x1DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1DE, 0x1DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1DF, 0x1DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E0, 0x1E0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E1, 0x1E1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E2, 0x1E2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E3, 0x1E3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E4, 0x1E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E5, 0x1E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E6, 0x1E6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E7, 0x1E7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E8, 0x1E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E9, 0x1E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1EA, 0x1EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1EB, 0x1EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1EC, 0x1EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1ED, 0x1ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1EE, 0x1EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1EF, 0x1F0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1F1, 0x1F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1F3, 0x1F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1F4, 0x1F4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1F5, 0x1F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1F6, 0x1F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1F9, 0x1F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1FA, 0x1FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1FB, 0x1FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1FC, 0x1FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1FD, 0x1FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1FE, 0x1FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1FF, 0x1FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x200, 0x200, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x201, 0x201, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x202, 0x202, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x203, 0x203, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x204, 0x204, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x205, 0x205, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x206, 0x206, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x207, 0x207, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x208, 0x208, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x209, 0x209, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x20A, 0x20A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x20B, 0x20B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x20C, 0x20C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x20D, 0x20D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x20E, 0x20E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x20F, 0x20F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x210, 0x210, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x211, 0x211, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x212, 0x212, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x213, 0x213, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x214, 0x214, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x215, 0x215, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x216, 0x216, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x217, 0x217, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x218, 0x218, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x219, 0x219, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x21A, 0x21A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x21B, 0x21B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x21C, 0x21C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x21D, 0x21D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x21E, 0x21E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x21F, 0x21F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x220, 0x220, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x221, 0x221, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x222, 0x222, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x223, 0x223, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x224, 0x224, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x225, 0x225, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x226, 0x226, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x227, 0x227, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x228, 0x228, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x229, 0x229, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x22A, 0x22A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x22B, 0x22B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x22C, 0x22C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x22D, 0x22D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x22E, 0x22E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x22F, 0x22F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x230, 0x230, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x231, 0x231, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x232, 0x232, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x233, 0x239, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x23A, 0x23B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x23C, 0x23C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x23D, 0x23E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x23F, 0x240, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x241, 0x241, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x242, 0x242, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x243, 0x246, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x247, 0x247, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x248, 0x248, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x249, 0x249, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x24A, 0x24A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x24B, 0x24B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x24C, 0x24C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x24D, 0x24D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x24E, 0x24E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x24F, 0x293, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x294, 0x294, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x295, 0x2B8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2B9, 0x2BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x2C0, 0x2C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2C2, 0x2C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2C6, 0x2D1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x2D2, 0x2DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2E0, 0x2E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2E5, 0x2EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2EC, 0x2EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x2ED, 0x2ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2EE, 0x2EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x2EF, 0x2FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x300, 0x36F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x370, 0x370, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x371, 0x371, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x372, 0x372, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x373, 0x373, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x374, 0x374, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x375, 0x375, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x376, 0x376, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x377, 0x377, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x378, 0x379, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x37A, 0x37D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x37E, 0x37E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x37F, 0x37F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x380, 0x385, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x386, 0x386, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x387, 0x387, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x388, 0x38A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x38B, 0x38B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x38C, 0x38C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x38D, 0x38D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x38E, 0x38F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x390, 0x390, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x391, 0x3A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3A2, 0x3A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x3A3, 0x3AB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3AC, 0x3CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3CF, 0x3CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3D0, 0x3D1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3D2, 0x3D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3D5, 0x3D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3D8, 0x3D8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3D9, 0x3D9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3DA, 0x3DA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3DB, 0x3DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3DC, 0x3DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3DD, 0x3DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3DE, 0x3DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3DF, 0x3DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3E0, 0x3E0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3E1, 0x3E1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3E2, 0x3E2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3E3, 0x3E3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3E4, 0x3E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3E5, 0x3E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3E6, 0x3E6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3E7, 0x3E7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3E8, 0x3E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3E9, 0x3E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3EA, 0x3EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3EB, 0x3EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3EC, 0x3EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3ED, 0x3ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3EE, 0x3EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3EF, 0x3F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3F4, 0x3F4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3F5, 0x3F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3F6, 0x3F6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x3F7, 0x3F7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3F8, 0x3F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3F9, 0x3FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x3FB, 0x3FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x3FD, 0x42F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x430, 0x45F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x460, 0x460, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x461, 0x461, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x462, 0x462, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x463, 0x463, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x464, 0x464, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x465, 0x465, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x466, 0x466, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x467, 0x467, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x468, 0x468, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x469, 0x469, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x46A, 0x46A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x46B, 0x46B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x46C, 0x46C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x46D, 0x46D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x46E, 0x46E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x46F, 0x46F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x470, 0x470, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x471, 0x471, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x472, 0x472, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x473, 0x473, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x474, 0x474, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x475, 0x475, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x476, 0x476, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x477, 0x477, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x478, 0x478, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x479, 0x479, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x47A, 0x47A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x47B, 0x47B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x47C, 0x47C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x47D, 0x47D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x47E, 0x47E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x47F, 0x47F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x480, 0x480, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x481, 0x481, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x482, 0x482, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x483, 0x489, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x48A, 0x48A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x48B, 0x48B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x48C, 0x48C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x48D, 0x48D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x48E, 0x48E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x48F, 0x48F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x490, 0x490, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x491, 0x491, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x492, 0x492, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x493, 0x493, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x494, 0x494, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x495, 0x495, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x496, 0x496, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x497, 0x497, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x498, 0x498, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x499, 0x499, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x49A, 0x49A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x49B, 0x49B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x49C, 0x49C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x49D, 0x49D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x49E, 0x49E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x49F, 0x49F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4A0, 0x4A0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4A1, 0x4A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4A2, 0x4A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4A3, 0x4A3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4A4, 0x4A4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4A5, 0x4A5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4A6, 0x4A6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4A7, 0x4A7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4A8, 0x4A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4A9, 0x4A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4AA, 0x4AA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4AB, 0x4AB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4AC, 0x4AC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4AD, 0x4AD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4AE, 0x4AE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4AF, 0x4AF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4B0, 0x4B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4B1, 0x4B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4B2, 0x4B2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4B3, 0x4B3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4B4, 0x4B4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4B5, 0x4B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4B6, 0x4B6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4B7, 0x4B7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4B8, 0x4B8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4B9, 0x4B9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4BA, 0x4BA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4BB, 0x4BB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4BC, 0x4BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4BD, 0x4BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4BE, 0x4BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4BF, 0x4BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x4C0, 0x4C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4C9, 0x4C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4D2, 0x4D2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4DA, 0x4DA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4E2, 0x4E2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4EA, 0x4EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4F2, 0x4F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x4FA, 0x4FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x502, 0x502, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x50A, 0x50A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x512, 0x512, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x51A, 0x51A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x522, 0x522, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x52A, 0x52A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x557, 0x558, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x591, 0x5BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x5C7, 0x5C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x60C, 0x60D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSc,
		}, []*autogenTestCaseFunc{}},
		{0x660, 0x669, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0x6D5, 0x6D5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x6EA, 0x6ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x70F, 0x70F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbFo,
		}, []*autogenTestCaseFunc{}},
		{0x7B1, 0x7B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x7F9, 0x7F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0x81B, 0x823, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x839, 0x839, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0x86B, 0x86F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x8A0, 0x8C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x950, 0x950, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x981, 0x983, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x9AA, 0x9B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x9BE, 0x9C4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x9D8, 0x9DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9F2, 0x9FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA0B, 0xA0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA34, 0xA34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA43, 0xA46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA5D, 0xA5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA81, 0xA83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xAAA, 0xAB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xABE, 0xAC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xAE0, 0xAE1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB01, 0xB03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xB2A, 0xB30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB3E, 0xB44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xB5C, 0xB5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB72, 0xB81, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB92, 0xB95, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xBA3, 0xBA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xBC6, 0xBC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xBE6, 0xBEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0xC29, 0xC29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC49, 0xC49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC5E, 0xC5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC84, 0xC84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCB4, 0xCB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCC9, 0xCC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCE2, 0xCE3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xD0D, 0xD0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD46, 0xD48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xD5F, 0xD61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xD84, 0xD84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDBE, 0xDBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDD7, 0xDD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xE31, 0xE31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xE5A, 0xE80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xEA4, 0xEA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xEBE, 0xEBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xEDA, 0xEDB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF2A, 0xF34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF40, 0xF47, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xF8D, 0xF97, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x103F, 0x103F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1061, 0x1061, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x108E, 0x108E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10C8, 0x10CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x124E, 0x124F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1289, 0x1289, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x12BF, 0x12BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1311, 0x1311, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1363, 0x1366, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x13FE, 0x1400, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x169D, 0x169F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x171F, 0x1731, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x176D, 0x176D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x17D7, 0x17D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1803, 0x1803, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0x1810, 0x1819, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0x18AA, 0x18AA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1930, 0x193B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x1980, 0x19AB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1A1C, 0x1A1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1A8A, 0x1A8F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1B00, 0x1B04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x1B5E, 0x1B5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0x1BA1, 0x1BAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x1C38, 0x1C3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1C7E, 0x1C7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0x1CD3, 0x1CD3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1CFA, 0x1CFA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1E04, 0x1E04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E0C, 0x1E0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E14, 0x1E14, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E1C, 0x1E1C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E24, 0x1E24, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E2C, 0x1E2C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E34, 0x1E34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E3C, 0x1E3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E44, 0x1E44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E4C, 0x1E4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E54, 0x1E54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E5C, 0x1E5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E64, 0x1E64, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E6C, 0x1E6C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E74, 0x1E74, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E7C, 0x1E7C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E84, 0x1E84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E8C, 0x1E8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1E94, 0x1E94, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1EA4, 0x1EA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1EAC, 0x1EAC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1EB4, 0x1EB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1EBC, 0x1EBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1EC4, 0x1EC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1ECC, 0x1ECC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1ED4, 0x1ED4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1EDC, 0x1EDC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1EE4, 0x1EE4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1EEC, 0x1EEC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1EF4, 0x1EF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1EFC, 0x1EFC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1F1E, 0x1F1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1F4E, 0x1F4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1F5E, 0x1F5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1F90, 0x1F97, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1FBD, 0x1FBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1FD0, 0x1FD3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1FF2, 0x1FF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x200E, 0x200F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbFo,
		}, []*autogenTestCaseFunc{}},
		{0x2028, 0x2029, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSe,
		}, []*autogenTestCaseFunc{}},
		{0x2045, 0x2046, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0x2071, 0x2071, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x209D, 0x20CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x210B, 0x210D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x211E, 0x2123, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x212E, 0x212E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x213E, 0x213F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2170, 0x217F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2329, 0x232A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0x2776, 0x27C4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x29DC, 0x29FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2C65, 0x2C66, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2C71, 0x2C71, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2C83, 0x2C83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2C8B, 0x2C8B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2C93, 0x2C93, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2C9B, 0x2C9B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2CA3, 0x2CA3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2CAB, 0x2CAB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2CB3, 0x2CB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2CBB, 0x2CBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2CC3, 0x2CC3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2CCB, 0x2CCB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2CD3, 0x2CD3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2CDB, 0x2CDB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2CE3, 0x2CE4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2CF3, 0x2CF3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2D30, 0x2D67, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x2DA7, 0x2DA7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2DC7, 0x2DC7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2E00, 0x2E0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0x2E30, 0x2E3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x3000, 0x3000, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSp,
		}, []*autogenTestCaseFunc{}},
		{0x301C, 0x301C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x3038, 0x303C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x30A1, 0x30FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x31A0, 0x31BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA4D0, 0xA4FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA62A, 0xA62B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA646, 0xA646, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA64E, 0xA64E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA656, 0xA656, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA65E, 0xA65E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA666, 0xA666, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA66E, 0xA66E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA682, 0xA682, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA68A, 0xA68A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA692, 0xA692, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA69A, 0xA69A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA6F7, 0xA6F7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0xA726, 0xA726, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA72E, 0xA72E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA738, 0xA738, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA740, 0xA740, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA748, 0xA748, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA750, 0xA750, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA758, 0xA758, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA760, 0xA760, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA768, 0xA768, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA779, 0xA779, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA782, 0xA782, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA78B, 0xA78B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA793, 0xA795, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA79D, 0xA79D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA7A5, 0xA7A5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA7B5, 0xA7B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA7BD, 0xA7BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA7C8, 0xA7C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA7D4, 0xA7D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA7F5, 0xA7F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA807, 0xA80A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA874, 0xA875, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA8D0, 0xA8D9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0xA8FF, 0xA8FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xA954, 0xA95F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA9CA, 0xA9CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA9FA, 0xA9FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAA4C, 0xAA4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xAA7B, 0xAA7D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xAABE, 0xAABF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xAAEB, 0xAAEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xAB0F, 0xAB10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAB5B, 0xAB5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xABEC, 0xABED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xD7CB, 0xD7FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xFB13, 0xFB17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xFB38, 0xFB3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xFB46, 0xFBB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xFDC8, 0xFDEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFE17, 0xFE18, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0xFE47, 0xFE48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0xFE59, 0xFE5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0xFEFF, 0xFEFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbFo,
		}, []*autogenTestCaseFunc{}},
		{0xFF0F, 0xFF0F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFF3C, 0xFF3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFF5F, 0xFF60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0xFFBF, 0xFFC1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFFDD, 0xFFF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1003B, 0x1003B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x100FB, 0x1013F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x102D1, 0x102DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10376, 0x1037A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x103D1, 0x103D5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x104B0, 0x104D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x10570, 0x1057A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x10597, 0x105A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x10600, 0x10736, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10783, 0x10785, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x10808, 0x10808, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1083F, 0x10855, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x108F4, 0x108F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x109BE, 0x109BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10A10, 0x10A13, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10A3F, 0x10A3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x10AC0, 0x10AC7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10B56, 0x10B5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10CB3, 0x10CBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10E80, 0x10EA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10F27, 0x10F27, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10F82, 0x10F85, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11003, 0x11037, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11075, 0x11075, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x110C2, 0x110C2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11100, 0x11102, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11145, 0x11146, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11180, 0x11182, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x111CE, 0x111CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11200, 0x11211, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1123E, 0x1123E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x1128F, 0x1129D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x112F0, 0x112F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0x11313, 0x11328, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1133B, 0x1133C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11350, 0x11350, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1136D, 0x1136F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11450, 0x11459, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0x114C6, 0x114C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x115B8, 0x115C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11600, 0x1162F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11680, 0x116AA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1171D, 0x1172B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11800, 0x1182B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11907, 0x11908, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11930, 0x11935, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11942, 0x11943, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x119A8, 0x119A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x119E4, 0x119E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11A3F, 0x11A41, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11A8A, 0x11A99, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11C09, 0x11C09, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11C50, 0x11C59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0x11D00, 0x11D06, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11D3B, 0x11D3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11D5A, 0x11D5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11D90, 0x11D91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11EF3, 0x11EF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x1246F, 0x1247F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x13439, 0x143FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16A6A, 0x16A6D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16AF0, 0x16AF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x16B44, 0x16B44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0x16E40, 0x16E5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x16F50, 0x16F50, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x16FE3, 0x16FE3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x18CD6, 0x18CFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1AFFF, 0x1AFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1B2FC, 0x1BBFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1BC9A, 0x1BC9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1CF47, 0x1D164, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D18C, 0x1D1A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D44E, 0x1D454, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1D4A0, 0x1D4A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D4B6, 0x1D4B9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1D4EA, 0x1D503, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1D51D, 0x1D51D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D546, 0x1D546, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1D5BA, 0x1D5D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1D68A, 0x1D6A5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1D6FB, 0x1D6FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D750, 0x1D755, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1D7AA, 0x1D7C2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1DA00, 0x1DA36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x1DA88, 0x1DA88, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0x1DF0B, 0x1DF1E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E023, 0x1E024, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x1E13E, 0x1E13F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E2C0, 0x1E2EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1E7ED, 0x1E7EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1E900, 0x1E921, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1EE04, 0x1EE04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE28, 0x1EE28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE3C, 0x1EE41, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE4C, 0x1EE4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE58, 0x1EE58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE60, 0x1EE60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE73, 0x1EE73, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE8A, 0x1EE8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EEBC, 0x1F12F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1F679, 0x1FBEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2B81E, 0x2B81F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x3134B, 0xE0000, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbAt,
			&autogenTestCasesFuncSbCl,
			&autogenTestCasesFuncSbCr,
			&autogenTestCasesFuncSbEx,
			&autogenTestCasesFuncSbFo,
			&autogenTestCasesFuncSbLe,
			&autogenTestCasesFuncSbLf,
			&autogenTestCasesFuncSbLo,
			&autogenTestCasesFuncSbNu,
			&autogenTestCasesFuncSbSc,
			&autogenTestCasesFuncSbSe,
			&autogenTestCasesFuncSbSp,
			&autogenTestCasesFuncSbSt,
			&autogenTestCasesFuncSbUp,
			&autogenTestCasesFuncSbXx,
		}},
	},
	{ // Wb
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbCr,
			&autogenTestCasesFuncWbDq,
			&autogenTestCasesFuncWbEx,
			&autogenTestCasesFuncWbExtend,
			&autogenTestCasesFuncWbFo,
			&autogenTestCasesFuncWbHl,
			&autogenTestCasesFuncWbKa,
			&autogenTestCasesFuncWbLe,
			&autogenTestCasesFuncWbLf,
			&autogenTestCasesFuncWbMb,
			&autogenTestCasesFuncWbMl,
			&autogenTestCasesFuncWbMn,
			&autogenTestCasesFuncWbNl,
			&autogenTestCasesFuncWbNu,
			&autogenTestCasesFuncWbRi,
			&autogenTestCasesFuncWbSq,
			&autogenTestCasesFuncWbWsegspace,
			&autogenTestCasesFuncWbXx,
			&autogenTestCasesFuncWbZwj,
		}},
		{0x00, 0x09, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x0A, 0x0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLf,
		}, []*autogenTestCaseFunc{}},
		{0x0B, 0x0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNl,
		}, []*autogenTestCaseFunc{}},
		{0x0D, 0x0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbCr,
		}, []*autogenTestCaseFunc{}},
		{0x0E, 0x1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x20, 0x20, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbWsegspace,
		}, []*autogenTestCaseFunc{}},
		{0x21, 0x21, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x22, 0x22, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbDq,
		}, []*autogenTestCaseFunc{}},
		{0x23, 0x26, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x27, 0x27, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbSq,
		}, []*autogenTestCaseFunc{}},
		{0x28, 0x2B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2C, 0x2C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMn,
		}, []*autogenTestCaseFunc{}},
		{0x2D, 0x2D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2E, 0x2E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMb,
		}, []*autogenTestCaseFunc{}},
		{0x2F, 0x2F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x30, 0x39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0x3A, 0x3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMl,
		}, []*autogenTestCaseFunc{}},
		{0x3B, 0x3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMn,
		}, []*autogenTestCaseFunc{}},
		{0x3C, 0x40, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x41, 0x5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x5B, 0x5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5F, 0x5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbEx,
		}, []*autogenTestCaseFunc{}},
		{0x60, 0x60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x61, 0x7A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x7B, 0x84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x85, 0x85, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNl,
		}, []*autogenTestCaseFunc{}},
		{0x86, 0xA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAA, 0xAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAB, 0xAC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAD, 0xAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbFo,
		}, []*autogenTestCaseFunc{}},
		{0xAE, 0xB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB5, 0xB5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB6, 0xB6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB7, 0xB7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMl,
		}, []*autogenTestCaseFunc{}},
		{0xB8, 0xB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBA, 0xBA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xBB, 0xBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC0, 0xD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xD7, 0xD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD8, 0xF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xF7, 0xF7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF8, 0x2D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x2D8, 0x2DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2DE, 0x2FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x300, 0x36F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x370, 0x374, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x375, 0x375, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x376, 0x377, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x378, 0x379, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x37A, 0x37D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x37E, 0x37E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMn,
		}, []*autogenTestCaseFunc{}},
		{0x37F, 0x37F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x380, 0x385, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x386, 0x386, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x387, 0x387, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMl,
		}, []*autogenTestCaseFunc{}},
		{0x388, 0x38A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x38B, 0x38B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x38C, 0x38C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x38D, 0x38D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x38E, 0x3A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x3A2, 0x3A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x3A3, 0x3F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x3F6, 0x3F6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x3F7, 0x481, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x482, 0x482, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x483, 0x489, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x48A, 0x52F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x530, 0x530, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x531, 0x556, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x557, 0x558, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x559, 0x55C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x55D, 0x55D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x55E, 0x55E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x55F, 0x55F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMl,
		}, []*autogenTestCaseFunc{}},
		{0x560, 0x588, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x589, 0x589, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMn,
		}, []*autogenTestCaseFunc{}},
		{0x58A, 0x58A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x58B, 0x590, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x591, 0x5BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x5BE, 0x5BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5BF, 0x5BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x5C0, 0x5C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5C1, 0x5C2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x5C3, 0x5C3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5C4, 0x5C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x5C6, 0x5C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5C7, 0x5C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x5C8, 0x5CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5D0, 0x5EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbHl,
		}, []*autogenTestCaseFunc{}},
		{0x5EB, 0x5EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x5EF, 0x5F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbHl,
		}, []*autogenTestCaseFunc{}},
		{0x5F3, 0x5F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x5F4, 0x5F4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMl,
		}, []*autogenTestCaseFunc{}},
		{0x5F5, 0x5FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x600, 0x605, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbFo,
		}, []*autogenTestCaseFunc{}},
		{0x606, 0x60B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x60C, 0x60D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMn,
		}, []*autogenTestCaseFunc{}},
		{0x60E, 0x60F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x610, 0x61A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x61B, 0x61B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x61C, 0x61C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbFo,
		}, []*autogenTestCaseFunc{}},
		{0x61D, 0x61F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x620, 0x64A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x64B, 0x65F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x660, 0x669, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0x66A, 0x66A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x66B, 0x66B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0x66C, 0x66C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMn,
		}, []*autogenTestCaseFunc{}},
		{0x66D, 0x66D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x66E, 0x66F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x670, 0x670, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x671, 0x6D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x6D4, 0x6D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x6D5, 0x6D5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x6D6, 0x6DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x6DD, 0x6DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbFo,
		}, []*autogenTestCaseFunc{}},
		{0x6DE, 0x6DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x6DF, 0x6E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x6E5, 0x6E6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x6E7, 0x6E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x6E9, 0x6E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x6EA, 0x6ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x6EE, 0x6EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x6F0, 0x6F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0x6FA, 0x6FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x6FD, 0x6FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x6FF, 0x6FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x700, 0x70E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x70F, 0x70F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbFo,
		}, []*autogenTestCaseFunc{}},
		{0x710, 0x710, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x711, 0x711, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x712, 0x72F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x730, 0x74A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x74B, 0x74C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x74D, 0x7A5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x7A6, 0x7B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x7B1, 0x7B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x7B2, 0x7BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x7C0, 0x7C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0x7CA, 0x7EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x7EB, 0x7F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x7F4, 0x7F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x7F6, 0x7F7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x7F8, 0x7F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMn,
		}, []*autogenTestCaseFunc{}},
		{0x7F9, 0x7F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x7FA, 0x7FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x7FB, 0x7FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x7FD, 0x7FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x7FE, 0x7FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x800, 0x815, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x816, 0x819, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x81A, 0x81A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x81B, 0x823, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x824, 0x824, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x825, 0x827, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x828, 0x828, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x829, 0x82D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x82E, 0x83F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x840, 0x858, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x859, 0x85B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x85C, 0x85F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x860, 0x86A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x86B, 0x86F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x870, 0x887, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x888, 0x888, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x889, 0x88E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x88F, 0x88F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x890, 0x891, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbFo,
		}, []*autogenTestCaseFunc{}},
		{0x892, 0x897, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x898, 0x89F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x8A0, 0x8C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x8CA, 0x8E1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x8E2, 0x8E2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbFo,
		}, []*autogenTestCaseFunc{}},
		{0x8E3, 0x903, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x904, 0x939, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x93A, 0x93C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x93D, 0x93D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x93E, 0x94F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x950, 0x950, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x951, 0x957, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x958, 0x961, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x962, 0x963, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x964, 0x965, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x966, 0x96F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0x970, 0x970, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x971, 0x980, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x981, 0x983, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x984, 0x984, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x985, 0x98C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x98D, 0x98E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x98F, 0x990, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x991, 0x992, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x993, 0x9A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x9A9, 0x9A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9AA, 0x9B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x9B1, 0x9B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9B2, 0x9B2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x9B3, 0x9B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9B6, 0x9B9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x9BA, 0x9BB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9BC, 0x9BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x9BD, 0x9BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x9BE, 0x9C4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x9C5, 0x9C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9C7, 0x9C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x9C9, 0x9CA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9CB, 0x9CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x9CE, 0x9CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x9CF, 0x9D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9D7, 0x9D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x9D8, 0x9DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9DC, 0x9DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x9DE, 0x9DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9DF, 0x9E1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x9E2, 0x9E3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x9E4, 0x9E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9E6, 0x9EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0x9F0, 0x9F1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x9F2, 0x9FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9FC, 0x9FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x9FD, 0x9FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x9FE, 0x9FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x9FF, 0xA00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA01, 0xA03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xA04, 0xA04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA05, 0xA0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA0B, 0xA0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA0F, 0xA10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA11, 0xA12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA13, 0xA28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA29, 0xA29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA2A, 0xA30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA31, 0xA31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA32, 0xA33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA34, 0xA34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA35, 0xA36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA37, 0xA37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA38, 0xA39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA3A, 0xA3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA3C, 0xA3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xA3D, 0xA3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA3E, 0xA42, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xA43, 0xA46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA47, 0xA48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xA49, 0xA4A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA4B, 0xA4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xA4E, 0xA50, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA51, 0xA51, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xA52, 0xA58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA59, 0xA5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA5D, 0xA5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA5E, 0xA5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA5F, 0xA65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA66, 0xA6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0xA70, 0xA71, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xA72, 0xA74, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA75, 0xA75, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xA76, 0xA80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA81, 0xA83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xA84, 0xA84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA85, 0xA8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA8E, 0xA8E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA8F, 0xA91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA92, 0xA92, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA93, 0xAA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAA9, 0xAA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAAA, 0xAB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAB1, 0xAB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAB2, 0xAB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAB4, 0xAB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAB5, 0xAB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xABA, 0xABB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xABC, 0xABC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xABD, 0xABD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xABE, 0xAC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xAC6, 0xAC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAC7, 0xAC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xACA, 0xACA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xACB, 0xACD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xACE, 0xACF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAD0, 0xAD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAD1, 0xADF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAE0, 0xAE1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAE2, 0xAE3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xAE4, 0xAE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAE6, 0xAEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0xAF0, 0xAF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAF9, 0xAF9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAFA, 0xAFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xB00, 0xB00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB01, 0xB03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xB04, 0xB04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB05, 0xB0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB0D, 0xB0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB0F, 0xB10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB11, 0xB12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB13, 0xB28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB29, 0xB29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB2A, 0xB30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB31, 0xB31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB32, 0xB33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB34, 0xB34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB35, 0xB39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB3A, 0xB3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB3C, 0xB3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xB3D, 0xB3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB3E, 0xB44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xB45, 0xB46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB47, 0xB48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xB49, 0xB4A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB4B, 0xB4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xB4E, 0xB54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB55, 0xB57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xB58, 0xB5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB5C, 0xB5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB5E, 0xB5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB5F, 0xB61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB62, 0xB63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xB64, 0xB65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB66, 0xB6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0xB70, 0xB70, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB71, 0xB71, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB72, 0xB81, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB82, 0xB82, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xB83, 0xB83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB84, 0xB84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB85, 0xB8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB8B, 0xB8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB8E, 0xB90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB91, 0xB91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB92, 0xB95, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB96, 0xB98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB99, 0xB9A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB9B, 0xB9B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB9C, 0xB9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xB9D, 0xB9D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xB9E, 0xB9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xBA0, 0xBA2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBA3, 0xBA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xBA5, 0xBA7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBA8, 0xBAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xBAB, 0xBAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBAE, 0xBB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xBBA, 0xBBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBBE, 0xBC2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xBC3, 0xBC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBC6, 0xBC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xBC9, 0xBC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBCA, 0xBCD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xBCE, 0xBCF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBD0, 0xBD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xBD1, 0xBD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBD7, 0xBD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xBD8, 0xBE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xBE6, 0xBEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0xBF0, 0xBFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC00, 0xC04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xC05, 0xC0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xC0D, 0xC0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC0E, 0xC10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xC11, 0xC11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC12, 0xC28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xC29, 0xC29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC2A, 0xC39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xC3A, 0xC3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC3C, 0xC3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xC3D, 0xC3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xC3E, 0xC44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xC45, 0xC45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC46, 0xC48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xC49, 0xC49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC4A, 0xC4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xC4E, 0xC54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC55, 0xC56, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xC57, 0xC57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC58, 0xC5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xC5B, 0xC5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC5D, 0xC5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xC5E, 0xC5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC60, 0xC61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xC62, 0xC63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xC64, 0xC65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC66, 0xC6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0xC70, 0xC7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC80, 0xC80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xC81, 0xC83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xC84, 0xC84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC85, 0xC8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xC8D, 0xC8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC8E, 0xC90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xC91, 0xC91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xC92, 0xCA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xCA9, 0xCA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCAA, 0xCB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xCB4, 0xCB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCB5, 0xCB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xCBA, 0xCBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCBC, 0xCBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xCBD, 0xCBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xCBE, 0xCC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xCC5, 0xCC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCC6, 0xCC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xCC9, 0xCC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCCA, 0xCCD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xCCE, 0xCD4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCD5, 0xCD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xCD7, 0xCDC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCDD, 0xCDE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xCDF, 0xCDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCE0, 0xCE1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xCE2, 0xCE3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xCE4, 0xCE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCE6, 0xCEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0xCF0, 0xCF0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xCF1, 0xCF2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xCF3, 0xCFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD00, 0xD03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xD04, 0xD0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xD0D, 0xD0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD0E, 0xD10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xD11, 0xD11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD12, 0xD3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xD3B, 0xD3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xD3D, 0xD3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xD3E, 0xD44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xD45, 0xD45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD46, 0xD48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xD49, 0xD49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD4A, 0xD4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xD4E, 0xD4E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xD4F, 0xD53, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD54, 0xD56, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xD57, 0xD57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xD58, 0xD5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD5F, 0xD61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xD62, 0xD63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xD64, 0xD65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD66, 0xD6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0xD70, 0xD79, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD7A, 0xD7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xD80, 0xD80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD81, 0xD83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xD84, 0xD84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD85, 0xD96, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xD97, 0xD99, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xD9A, 0xDB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xDB2, 0xDB2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDB3, 0xDBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xDBC, 0xDBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDBD, 0xDBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xDBE, 0xDBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDC0, 0xDC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xDC7, 0xDC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDCA, 0xDCA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xDCB, 0xDCE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDCF, 0xDD4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xDD5, 0xDD5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDD6, 0xDD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xDD7, 0xDD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDD8, 0xDDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xDE0, 0xDE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDE6, 0xDEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0xDF0, 0xDF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xDF2, 0xDF3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xDF4, 0xE30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xE31, 0xE31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xE32, 0xE33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xE34, 0xE3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xE3B, 0xE46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xE47, 0xE4E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xE4F, 0xE4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xE50, 0xE59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0xE5A, 0xEB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xEB1, 0xEB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xEB2, 0xEB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xEB4, 0xEBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xEBD, 0xEC7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xEC8, 0xECD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xECE, 0xECF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xED0, 0xED9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0xEDA, 0xEFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF00, 0xF00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xF01, 0xF17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF18, 0xF19, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xF1A, 0xF1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF20, 0xF29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0xF2A, 0xF34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF35, 0xF35, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xF36, 0xF36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF37, 0xF37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xF38, 0xF38, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF39, 0xF39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xF3A, 0xF3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF3E, 0xF3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xF40, 0xF47, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xF48, 0xF48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF49, 0xF6C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xF6D, 0xF70, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF71, 0xF84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xF85, 0xF85, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xF86, 0xF87, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x102B, 0x103E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x1062, 0x1064, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x108F, 0x108F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x10CD, 0x10CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1250, 0x1256, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x128A, 0x128D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x12C0, 0x12C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1312, 0x1315, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x13A0, 0x13F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1681, 0x169A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1716, 0x171E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x176D, 0x176D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x17DE, 0x17DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1820, 0x1878, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x18B0, 0x18F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1946, 0x194F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0x1A5F, 0x1A5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1AB0, 0x1ACE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x1B5A, 0x1B6A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1BBA, 0x1BE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1C4D, 0x1C4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1CBD, 0x1CBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1CF4, 0x1CF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x1F16, 0x1F17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1F58, 0x1F58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1F7E, 0x1F7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1FC5, 0x1FC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1FED, 0x1FF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x200B, 0x200B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2025, 0x2026, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2044, 0x2044, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMn,
		}, []*autogenTestCaseFunc{}},
		{0x2070, 0x2070, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x20F1, 0x2101, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2116, 0x2118, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2129, 0x2129, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x214A, 0x214D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2CE5, 0x2CEA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2D28, 0x2D2C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2D80, 0x2D96, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x2DB8, 0x2DBE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x2DD8, 0x2DDE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x3005, 0x3005, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x3099, 0x309A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x3130, 0x3130, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x32FF, 0x32FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA60D, 0xA60F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA674, 0xA67D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xA7CB, 0xA7CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA802, 0xA802, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xA82C, 0xA82C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xA8D0, 0xA8D9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0xA8FF, 0xA8FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xA960, 0xA97C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA9DA, 0xA9E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAA40, 0xAA42, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAA7E, 0xAAAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAAC0, 0xAAC0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAAF7, 0xAB00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAB27, 0xAB27, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xABEC, 0xABED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xD7CB, 0xD7FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xFB1F, 0xFB28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbHl,
		}, []*autogenTestCaseFunc{}},
		{0xFB40, 0xFB41, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbHl,
		}, []*autogenTestCaseFunc{}},
		{0xFD3E, 0xFD4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFE10, 0xFE10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbMn,
		}, []*autogenTestCaseFunc{}},
		{0xFE35, 0xFE4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFE56, 0xFE6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFF08, 0xFF0B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFF1C, 0xFF20, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFF9E, 0xFF9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0xFFD8, 0xFFD9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10027, 0x10027, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1005E, 0x1007F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1029D, 0x1029F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1034B, 0x1034F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x103C8, 0x103CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x104B0, 0x104D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10570, 0x1057A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10597, 0x105A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10600, 0x10736, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10787, 0x107B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1080A, 0x10835, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10860, 0x10876, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10900, 0x10915, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10A00, 0x10A00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10A15, 0x10A17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10A60, 0x10A7C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10AE7, 0x10AFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10B92, 0x10BFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10D24, 0x10D27, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x10EB0, 0x10EB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10F51, 0x10F6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11000, 0x11002, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11075, 0x11075, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x110C2, 0x110C2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11100, 0x11102, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11147, 0x11147, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11183, 0x111B2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x111DA, 0x111DA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11238, 0x1123D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1128E, 0x1128E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x112F0, 0x112F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0x11313, 0x11328, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1133B, 0x1133C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11350, 0x11350, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1136D, 0x1136F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1145A, 0x1145D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x114C7, 0x114C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x115C1, 0x115D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11645, 0x1164F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x116CA, 0x1171C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x118A0, 0x118DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11914, 0x11914, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1193B, 0x1193E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x119A0, 0x119A7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x119E3, 0x119E3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11A3B, 0x11A3E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11A9A, 0x11A9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11C2F, 0x11C36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11C90, 0x11C91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11D0A, 0x11D0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11D3F, 0x11D45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11D67, 0x11D68, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11D98, 0x11D98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11FB1, 0x11FFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x12FF1, 0x12FFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16A39, 0x16A3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16ACA, 0x16ACF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16B40, 0x16B43, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x16E40, 0x16E7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x16F8F, 0x16F92, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x16FF0, 0x16FF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x1B000, 0x1B000, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbKa,
		}, []*autogenTestCaseFunc{}},
		{0x1BC70, 0x1BC7C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1BCA0, 0x1BCA3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbFo,
		}, []*autogenTestCaseFunc{}},
		{0x1D16D, 0x1D172, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x1D242, 0x1D244, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x1D4A2, 0x1D4A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1D4BB, 0x1D4BB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1D50D, 0x1D514, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1D540, 0x1D544, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1D6A8, 0x1D6C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1D716, 0x1D734, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1D78A, 0x1D7A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1DA00, 0x1DA36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x1DA9B, 0x1DA9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x1E008, 0x1E018, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x1E100, 0x1E12C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1E14F, 0x1E28F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E7E0, 0x1E7E6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1E800, 0x1E8C4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1E950, 0x1E959, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0x1EE24, 0x1EE24, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1EE39, 0x1EE39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1EE49, 0x1EE49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1EE54, 0x1EE54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1EE5D, 0x1EE5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1EE67, 0x1EE6A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1EE7E, 0x1EE7E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1EEA5, 0x1EEA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1F170, 0x1F189, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0xE0001, 0xE0001, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbFo,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbCr,
			&autogenTestCasesFuncWbDq,
			&autogenTestCasesFuncWbEx,
			&autogenTestCasesFuncWbExtend,
			&autogenTestCasesFuncWbFo,
			&autogenTestCasesFuncWbHl,
			&autogenTestCasesFuncWbKa,
			&autogenTestCasesFuncWbLe,
			&autogenTestCasesFuncWbLf,
			&autogenTestCasesFuncWbMb,
			&autogenTestCasesFuncWbMl,
			&autogenTestCasesFuncWbMn,
			&autogenTestCasesFuncWbNl,
			&autogenTestCasesFuncWbNu,
			&autogenTestCasesFuncWbRi,
			&autogenTestCasesFuncWbSq,
			&autogenTestCasesFuncWbWsegspace,
			&autogenTestCasesFuncWbXx,
			&autogenTestCasesFuncWbZwj,
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

// autogenBreakTestEntry describes a break test entry.
type autogenBreakTestEntry struct {
	data  []rune // sample data
	ebbs  []int  // expected break boundaries
	label string // test case description
}

// autogenExtGraphemeCluster was derived from UCD data at https://www.unicode.org/ucd/.
var autogenExtGraphemeCluster = []autogenBreakTestEntry{
	{[]rune{0x0020, 0x0020}, []int{0, 1, 2}, "?? 0020 ?? 0020 ??"},
	{[]rune{0x0020, 0x0308, 0x0020}, []int{0, 2, 3}, "?? 0020 ?? 0308 ?? 0020 ??"},
	{[]rune{0x0020, 0x000D}, []int{0, 1, 2}, "?? 0020 ?? 000D ??"},
	{[]rune{0x0020, 0x0308, 0x000D}, []int{0, 2, 3}, "?? 0020 ?? 0308 ?? 000D ??"},
	{[]rune{0x0020, 0x000A}, []int{0, 1, 2}, "?? 0020 ?? 000A ??"},
	{[]rune{0x0020, 0x0308, 0x000A}, []int{0, 2, 3}, "?? 0020 ?? 0308 ?? 000A ??"},
	{[]rune{0x0020, 0x0001}, []int{0, 1, 2}, "?? 0020 ?? 0001 ??"},
	{[]rune{0x0020, 0x0308, 0x0001}, []int{0, 2, 3}, "?? 0020 ?? 0308 ?? 0001 ??"},
	{[]rune{0x0020, 0x034F}, []int{0, 2}, "?? 0020 ?? 034F ??"},
	{[]rune{0x0020, 0x0308, 0x034F}, []int{0, 3}, "?? 0020 ?? 0308 ?? 034F ??"},
	{[]rune{0x0020, 0x1F1E6}, []int{0, 1, 2}, "?? 0020 ?? 1F1E6 ??"},
	{[]rune{0x0020, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? 0020 ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x0020, 0x0600}, []int{0, 1, 2}, "?? 0020 ?? 0600 ??"},
	{[]rune{0x0020, 0x0308, 0x0600}, []int{0, 2, 3}, "?? 0020 ?? 0308 ?? 0600 ??"},
	{[]rune{0x0020, 0x0903}, []int{0, 2}, "?? 0020 ?? 0903 ??"},
	{[]rune{0x0020, 0x0308, 0x0903}, []int{0, 3}, "?? 0020 ?? 0308 ?? 0903 ??"},
	{[]rune{0x0020, 0x1100}, []int{0, 1, 2}, "?? 0020 ?? 1100 ??"},
	{[]rune{0x0020, 0x0308, 0x1100}, []int{0, 2, 3}, "?? 0020 ?? 0308 ?? 1100 ??"},
	{[]rune{0x0020, 0x1160}, []int{0, 1, 2}, "?? 0020 ?? 1160 ??"},
	{[]rune{0x0020, 0x0308, 0x1160}, []int{0, 2, 3}, "?? 0020 ?? 0308 ?? 1160 ??"},
	{[]rune{0x0020, 0x11A8}, []int{0, 1, 2}, "?? 0020 ?? 11A8 ??"},
	{[]rune{0x0020, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? 0020 ?? 0308 ?? 11A8 ??"},
	{[]rune{0x0020, 0xAC00}, []int{0, 1, 2}, "?? 0020 ?? AC00 ??"},
	{[]rune{0x0020, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? 0020 ?? 0308 ?? AC00 ??"},
	{[]rune{0x0020, 0xAC01}, []int{0, 1, 2}, "?? 0020 ?? AC01 ??"},
	{[]rune{0x0020, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? 0020 ?? 0308 ?? AC01 ??"},
	{[]rune{0x0020, 0x231A}, []int{0, 1, 2}, "?? 0020 ?? 231A ??"},
	{[]rune{0x0020, 0x0308, 0x231A}, []int{0, 2, 3}, "?? 0020 ?? 0308 ?? 231A ??"},
	{[]rune{0x0020, 0x0300}, []int{0, 2}, "?? 0020 ?? 0300 ??"},
	{[]rune{0x0020, 0x0308, 0x0300}, []int{0, 3}, "?? 0020 ?? 0308 ?? 0300 ??"},
	{[]rune{0x0020, 0x200D}, []int{0, 2}, "?? 0020 ?? 200D ??"},
	{[]rune{0x0020, 0x0308, 0x200D}, []int{0, 3}, "?? 0020 ?? 0308 ?? 200D ??"},
	{[]rune{0x0020, 0x0378}, []int{0, 1, 2}, "?? 0020 ?? 0378 ??"},
	{[]rune{0x0020, 0x0308, 0x0378}, []int{0, 2, 3}, "?? 0020 ?? 0308 ?? 0378 ??"},
	{[]rune{0x000D, 0x0020}, []int{0, 1, 2}, "?? 000D ?? 0020 ??"},
	{[]rune{0x000D, 0x0308, 0x0020}, []int{0, 1, 2, 3}, "?? 000D ?? 0308 ?? 0020 ??"},
	{[]rune{0x000D, 0x000D}, []int{0, 1, 2}, "?? 000D ?? 000D ??"},
	{[]rune{0x000D, 0x0308, 0x000D}, []int{0, 1, 2, 3}, "?? 000D ?? 0308 ?? 000D ??"},
	{[]rune{0x000D, 0x000A}, []int{0, 2}, "?? 000D ?? 000A ??"},
	{[]rune{0x000D, 0x0308, 0x000A}, []int{0, 1, 2, 3}, "?? 000D ?? 0308 ?? 000A ??"},
	{[]rune{0x000D, 0x0001}, []int{0, 1, 2}, "?? 000D ?? 0001 ??"},
	{[]rune{0x000D, 0x0308, 0x0001}, []int{0, 1, 2, 3}, "?? 000D ?? 0308 ?? 0001 ??"},
	{[]rune{0x000D, 0x034F}, []int{0, 1, 2}, "?? 000D ?? 034F ??"},
	{[]rune{0x000D, 0x0308, 0x034F}, []int{0, 1, 3}, "?? 000D ?? 0308 ?? 034F ??"},
	{[]rune{0x000D, 0x1F1E6}, []int{0, 1, 2}, "?? 000D ?? 1F1E6 ??"},
	{[]rune{0x000D, 0x0308, 0x1F1E6}, []int{0, 1, 2, 3}, "?? 000D ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x000D, 0x0600}, []int{0, 1, 2}, "?? 000D ?? 0600 ??"},
	{[]rune{0x000D, 0x0308, 0x0600}, []int{0, 1, 2, 3}, "?? 000D ?? 0308 ?? 0600 ??"},
	{[]rune{0x000D, 0x0903}, []int{0, 1, 2}, "?? 000D ?? 0903 ??"},
	{[]rune{0x000D, 0x0308, 0x0903}, []int{0, 1, 3}, "?? 000D ?? 0308 ?? 0903 ??"},
	{[]rune{0x000D, 0x1100}, []int{0, 1, 2}, "?? 000D ?? 1100 ??"},
	{[]rune{0x000D, 0x0308, 0x1100}, []int{0, 1, 2, 3}, "?? 000D ?? 0308 ?? 1100 ??"},
	{[]rune{0x000D, 0x1160}, []int{0, 1, 2}, "?? 000D ?? 1160 ??"},
	{[]rune{0x000D, 0x0308, 0x1160}, []int{0, 1, 2, 3}, "?? 000D ?? 0308 ?? 1160 ??"},
	{[]rune{0x000D, 0x11A8}, []int{0, 1, 2}, "?? 000D ?? 11A8 ??"},
	{[]rune{0x000D, 0x0308, 0x11A8}, []int{0, 1, 2, 3}, "?? 000D ?? 0308 ?? 11A8 ??"},
	{[]rune{0x000D, 0xAC00}, []int{0, 1, 2}, "?? 000D ?? AC00 ??"},
	{[]rune{0x000D, 0x0308, 0xAC00}, []int{0, 1, 2, 3}, "?? 000D ?? 0308 ?? AC00 ??"},
	{[]rune{0x000D, 0xAC01}, []int{0, 1, 2}, "?? 000D ?? AC01 ??"},
	{[]rune{0x000D, 0x0308, 0xAC01}, []int{0, 1, 2, 3}, "?? 000D ?? 0308 ?? AC01 ??"},
	{[]rune{0x000D, 0x231A}, []int{0, 1, 2}, "?? 000D ?? 231A ??"},
	{[]rune{0x000D, 0x0308, 0x231A}, []int{0, 1, 2, 3}, "?? 000D ?? 0308 ?? 231A ??"},
	{[]rune{0x000D, 0x0300}, []int{0, 1, 2}, "?? 000D ?? 0300 ??"},
	{[]rune{0x000D, 0x0308, 0x0300}, []int{0, 1, 3}, "?? 000D ?? 0308 ?? 0300 ??"},
	{[]rune{0x000D, 0x200D}, []int{0, 1, 2}, "?? 000D ?? 200D ??"},
	{[]rune{0x000D, 0x0308, 0x200D}, []int{0, 1, 3}, "?? 000D ?? 0308 ?? 200D ??"},
	{[]rune{0x000D, 0x0378}, []int{0, 1, 2}, "?? 000D ?? 0378 ??"},
	{[]rune{0x000D, 0x0308, 0x0378}, []int{0, 1, 2, 3}, "?? 000D ?? 0308 ?? 0378 ??"},
	{[]rune{0x000A, 0x0020}, []int{0, 1, 2}, "?? 000A ?? 0020 ??"},
	{[]rune{0x000A, 0x0308, 0x0020}, []int{0, 1, 2, 3}, "?? 000A ?? 0308 ?? 0020 ??"},
	{[]rune{0x000A, 0x000D}, []int{0, 1, 2}, "?? 000A ?? 000D ??"},
	{[]rune{0x000A, 0x0308, 0x000D}, []int{0, 1, 2, 3}, "?? 000A ?? 0308 ?? 000D ??"},
	{[]rune{0x000A, 0x000A}, []int{0, 1, 2}, "?? 000A ?? 000A ??"},
	{[]rune{0x000A, 0x0308, 0x000A}, []int{0, 1, 2, 3}, "?? 000A ?? 0308 ?? 000A ??"},
	{[]rune{0x000A, 0x0001}, []int{0, 1, 2}, "?? 000A ?? 0001 ??"},
	{[]rune{0x000A, 0x0308, 0x0001}, []int{0, 1, 2, 3}, "?? 000A ?? 0308 ?? 0001 ??"},
	{[]rune{0x000A, 0x034F}, []int{0, 1, 2}, "?? 000A ?? 034F ??"},
	{[]rune{0x000A, 0x0308, 0x034F}, []int{0, 1, 3}, "?? 000A ?? 0308 ?? 034F ??"},
	{[]rune{0x000A, 0x1F1E6}, []int{0, 1, 2}, "?? 000A ?? 1F1E6 ??"},
	{[]rune{0x000A, 0x0308, 0x1F1E6}, []int{0, 1, 2, 3}, "?? 000A ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x000A, 0x0600}, []int{0, 1, 2}, "?? 000A ?? 0600 ??"},
	{[]rune{0x000A, 0x0308, 0x0600}, []int{0, 1, 2, 3}, "?? 000A ?? 0308 ?? 0600 ??"},
	{[]rune{0x000A, 0x0903}, []int{0, 1, 2}, "?? 000A ?? 0903 ??"},
	{[]rune{0x000A, 0x0308, 0x0903}, []int{0, 1, 3}, "?? 000A ?? 0308 ?? 0903 ??"},
	{[]rune{0x000A, 0x1100}, []int{0, 1, 2}, "?? 000A ?? 1100 ??"},
	{[]rune{0x000A, 0x0308, 0x1100}, []int{0, 1, 2, 3}, "?? 000A ?? 0308 ?? 1100 ??"},
	{[]rune{0x000A, 0x1160}, []int{0, 1, 2}, "?? 000A ?? 1160 ??"},
	{[]rune{0x000A, 0x0308, 0x1160}, []int{0, 1, 2, 3}, "?? 000A ?? 0308 ?? 1160 ??"},
	{[]rune{0x000A, 0x11A8}, []int{0, 1, 2}, "?? 000A ?? 11A8 ??"},
	{[]rune{0x000A, 0x0308, 0x11A8}, []int{0, 1, 2, 3}, "?? 000A ?? 0308 ?? 11A8 ??"},
	{[]rune{0x000A, 0xAC00}, []int{0, 1, 2}, "?? 000A ?? AC00 ??"},
	{[]rune{0x000A, 0x0308, 0xAC00}, []int{0, 1, 2, 3}, "?? 000A ?? 0308 ?? AC00 ??"},
	{[]rune{0x000A, 0xAC01}, []int{0, 1, 2}, "?? 000A ?? AC01 ??"},
	{[]rune{0x000A, 0x0308, 0xAC01}, []int{0, 1, 2, 3}, "?? 000A ?? 0308 ?? AC01 ??"},
	{[]rune{0x000A, 0x231A}, []int{0, 1, 2}, "?? 000A ?? 231A ??"},
	{[]rune{0x000A, 0x0308, 0x231A}, []int{0, 1, 2, 3}, "?? 000A ?? 0308 ?? 231A ??"},
	{[]rune{0x000A, 0x0300}, []int{0, 1, 2}, "?? 000A ?? 0300 ??"},
	{[]rune{0x000A, 0x0308, 0x0300}, []int{0, 1, 3}, "?? 000A ?? 0308 ?? 0300 ??"},
	{[]rune{0x000A, 0x200D}, []int{0, 1, 2}, "?? 000A ?? 200D ??"},
	{[]rune{0x000A, 0x0308, 0x200D}, []int{0, 1, 3}, "?? 000A ?? 0308 ?? 200D ??"},
	{[]rune{0x000A, 0x0378}, []int{0, 1, 2}, "?? 000A ?? 0378 ??"},
	{[]rune{0x000A, 0x0308, 0x0378}, []int{0, 1, 2, 3}, "?? 000A ?? 0308 ?? 0378 ??"},
	{[]rune{0x0001, 0x0020}, []int{0, 1, 2}, "?? 0001 ?? 0020 ??"},
	{[]rune{0x0001, 0x0308, 0x0020}, []int{0, 1, 2, 3}, "?? 0001 ?? 0308 ?? 0020 ??"},
	{[]rune{0x0001, 0x000D}, []int{0, 1, 2}, "?? 0001 ?? 000D ??"},
	{[]rune{0x0001, 0x0308, 0x000D}, []int{0, 1, 2, 3}, "?? 0001 ?? 0308 ?? 000D ??"},
	{[]rune{0x0001, 0x000A}, []int{0, 1, 2}, "?? 0001 ?? 000A ??"},
	{[]rune{0x0001, 0x0308, 0x000A}, []int{0, 1, 2, 3}, "?? 0001 ?? 0308 ?? 000A ??"},
	{[]rune{0x0001, 0x0001}, []int{0, 1, 2}, "?? 0001 ?? 0001 ??"},
	{[]rune{0x0001, 0x0308, 0x0001}, []int{0, 1, 2, 3}, "?? 0001 ?? 0308 ?? 0001 ??"},
	{[]rune{0x0001, 0x034F}, []int{0, 1, 2}, "?? 0001 ?? 034F ??"},
	{[]rune{0x0001, 0x0308, 0x034F}, []int{0, 1, 3}, "?? 0001 ?? 0308 ?? 034F ??"},
	{[]rune{0x0001, 0x1F1E6}, []int{0, 1, 2}, "?? 0001 ?? 1F1E6 ??"},
	{[]rune{0x0001, 0x0308, 0x1F1E6}, []int{0, 1, 2, 3}, "?? 0001 ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x0001, 0x0600}, []int{0, 1, 2}, "?? 0001 ?? 0600 ??"},
	{[]rune{0x0001, 0x0308, 0x0600}, []int{0, 1, 2, 3}, "?? 0001 ?? 0308 ?? 0600 ??"},
	{[]rune{0x0001, 0x0903}, []int{0, 1, 2}, "?? 0001 ?? 0903 ??"},
	{[]rune{0x0001, 0x0308, 0x0903}, []int{0, 1, 3}, "?? 0001 ?? 0308 ?? 0903 ??"},
	{[]rune{0x0001, 0x1100}, []int{0, 1, 2}, "?? 0001 ?? 1100 ??"},
	{[]rune{0x0001, 0x0308, 0x1100}, []int{0, 1, 2, 3}, "?? 0001 ?? 0308 ?? 1100 ??"},
	{[]rune{0x0001, 0x1160}, []int{0, 1, 2}, "?? 0001 ?? 1160 ??"},
	{[]rune{0x0001, 0x0308, 0x1160}, []int{0, 1, 2, 3}, "?? 0001 ?? 0308 ?? 1160 ??"},
	{[]rune{0x0001, 0x11A8}, []int{0, 1, 2}, "?? 0001 ?? 11A8 ??"},
	{[]rune{0x0001, 0x0308, 0x11A8}, []int{0, 1, 2, 3}, "?? 0001 ?? 0308 ?? 11A8 ??"},
	{[]rune{0x0001, 0xAC00}, []int{0, 1, 2}, "?? 0001 ?? AC00 ??"},
	{[]rune{0x0001, 0x0308, 0xAC00}, []int{0, 1, 2, 3}, "?? 0001 ?? 0308 ?? AC00 ??"},
	{[]rune{0x0001, 0xAC01}, []int{0, 1, 2}, "?? 0001 ?? AC01 ??"},
	{[]rune{0x0001, 0x0308, 0xAC01}, []int{0, 1, 2, 3}, "?? 0001 ?? 0308 ?? AC01 ??"},
	{[]rune{0x0001, 0x231A}, []int{0, 1, 2}, "?? 0001 ?? 231A ??"},
	{[]rune{0x0001, 0x0308, 0x231A}, []int{0, 1, 2, 3}, "?? 0001 ?? 0308 ?? 231A ??"},
	{[]rune{0x0001, 0x0300}, []int{0, 1, 2}, "?? 0001 ?? 0300 ??"},
	{[]rune{0x0001, 0x0308, 0x0300}, []int{0, 1, 3}, "?? 0001 ?? 0308 ?? 0300 ??"},
	{[]rune{0x0001, 0x200D}, []int{0, 1, 2}, "?? 0001 ?? 200D ??"},
	{[]rune{0x0001, 0x0308, 0x200D}, []int{0, 1, 3}, "?? 0001 ?? 0308 ?? 200D ??"},
	{[]rune{0x0001, 0x0378}, []int{0, 1, 2}, "?? 0001 ?? 0378 ??"},
	{[]rune{0x0001, 0x0308, 0x0378}, []int{0, 1, 2, 3}, "?? 0001 ?? 0308 ?? 0378 ??"},
	{[]rune{0x034F, 0x0020}, []int{0, 1, 2}, "?? 034F ?? 0020 ??"},
	{[]rune{0x034F, 0x0308, 0x0020}, []int{0, 2, 3}, "?? 034F ?? 0308 ?? 0020 ??"},
	{[]rune{0x034F, 0x000D}, []int{0, 1, 2}, "?? 034F ?? 000D ??"},
	{[]rune{0x034F, 0x0308, 0x000D}, []int{0, 2, 3}, "?? 034F ?? 0308 ?? 000D ??"},
	{[]rune{0x034F, 0x000A}, []int{0, 1, 2}, "?? 034F ?? 000A ??"},
	{[]rune{0x034F, 0x0308, 0x000A}, []int{0, 2, 3}, "?? 034F ?? 0308 ?? 000A ??"},
	{[]rune{0x034F, 0x0001}, []int{0, 1, 2}, "?? 034F ?? 0001 ??"},
	{[]rune{0x034F, 0x0308, 0x0001}, []int{0, 2, 3}, "?? 034F ?? 0308 ?? 0001 ??"},
	{[]rune{0x034F, 0x034F}, []int{0, 2}, "?? 034F ?? 034F ??"},
	{[]rune{0x034F, 0x0308, 0x034F}, []int{0, 3}, "?? 034F ?? 0308 ?? 034F ??"},
	{[]rune{0x034F, 0x1F1E6}, []int{0, 1, 2}, "?? 034F ?? 1F1E6 ??"},
	{[]rune{0x034F, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? 034F ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x034F, 0x0600}, []int{0, 1, 2}, "?? 034F ?? 0600 ??"},
	{[]rune{0x034F, 0x0308, 0x0600}, []int{0, 2, 3}, "?? 034F ?? 0308 ?? 0600 ??"},
	{[]rune{0x034F, 0x0903}, []int{0, 2}, "?? 034F ?? 0903 ??"},
	{[]rune{0x034F, 0x0308, 0x0903}, []int{0, 3}, "?? 034F ?? 0308 ?? 0903 ??"},
	{[]rune{0x034F, 0x1100}, []int{0, 1, 2}, "?? 034F ?? 1100 ??"},
	{[]rune{0x034F, 0x0308, 0x1100}, []int{0, 2, 3}, "?? 034F ?? 0308 ?? 1100 ??"},
	{[]rune{0x034F, 0x1160}, []int{0, 1, 2}, "?? 034F ?? 1160 ??"},
	{[]rune{0x034F, 0x0308, 0x1160}, []int{0, 2, 3}, "?? 034F ?? 0308 ?? 1160 ??"},
	{[]rune{0x034F, 0x11A8}, []int{0, 1, 2}, "?? 034F ?? 11A8 ??"},
	{[]rune{0x034F, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? 034F ?? 0308 ?? 11A8 ??"},
	{[]rune{0x034F, 0xAC00}, []int{0, 1, 2}, "?? 034F ?? AC00 ??"},
	{[]rune{0x034F, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? 034F ?? 0308 ?? AC00 ??"},
	{[]rune{0x034F, 0xAC01}, []int{0, 1, 2}, "?? 034F ?? AC01 ??"},
	{[]rune{0x034F, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? 034F ?? 0308 ?? AC01 ??"},
	{[]rune{0x034F, 0x231A}, []int{0, 1, 2}, "?? 034F ?? 231A ??"},
	{[]rune{0x034F, 0x0308, 0x231A}, []int{0, 2, 3}, "?? 034F ?? 0308 ?? 231A ??"},
	{[]rune{0x034F, 0x0300}, []int{0, 2}, "?? 034F ?? 0300 ??"},
	{[]rune{0x034F, 0x0308, 0x0300}, []int{0, 3}, "?? 034F ?? 0308 ?? 0300 ??"},
	{[]rune{0x034F, 0x200D}, []int{0, 2}, "?? 034F ?? 200D ??"},
	{[]rune{0x034F, 0x0308, 0x200D}, []int{0, 3}, "?? 034F ?? 0308 ?? 200D ??"},
	{[]rune{0x034F, 0x0378}, []int{0, 1, 2}, "?? 034F ?? 0378 ??"},
	{[]rune{0x034F, 0x0308, 0x0378}, []int{0, 2, 3}, "?? 034F ?? 0308 ?? 0378 ??"},
	{[]rune{0x1F1E6, 0x0020}, []int{0, 1, 2}, "?? 1F1E6 ?? 0020 ??"},
	{[]rune{0x1F1E6, 0x0308, 0x0020}, []int{0, 2, 3}, "?? 1F1E6 ?? 0308 ?? 0020 ??"},
	{[]rune{0x1F1E6, 0x000D}, []int{0, 1, 2}, "?? 1F1E6 ?? 000D ??"},
	{[]rune{0x1F1E6, 0x0308, 0x000D}, []int{0, 2, 3}, "?? 1F1E6 ?? 0308 ?? 000D ??"},
	{[]rune{0x1F1E6, 0x000A}, []int{0, 1, 2}, "?? 1F1E6 ?? 000A ??"},
	{[]rune{0x1F1E6, 0x0308, 0x000A}, []int{0, 2, 3}, "?? 1F1E6 ?? 0308 ?? 000A ??"},
	{[]rune{0x1F1E6, 0x0001}, []int{0, 1, 2}, "?? 1F1E6 ?? 0001 ??"},
	{[]rune{0x1F1E6, 0x0308, 0x0001}, []int{0, 2, 3}, "?? 1F1E6 ?? 0308 ?? 0001 ??"},
	{[]rune{0x1F1E6, 0x034F}, []int{0, 2}, "?? 1F1E6 ?? 034F ??"},
	{[]rune{0x1F1E6, 0x0308, 0x034F}, []int{0, 3}, "?? 1F1E6 ?? 0308 ?? 034F ??"},
	{[]rune{0x1F1E6, 0x1F1E6}, []int{0, 2}, "?? 1F1E6 ?? 1F1E6 ??"},
	{[]rune{0x1F1E6, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? 1F1E6 ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x1F1E6, 0x0600}, []int{0, 1, 2}, "?? 1F1E6 ?? 0600 ??"},
	{[]rune{0x1F1E6, 0x0308, 0x0600}, []int{0, 2, 3}, "?? 1F1E6 ?? 0308 ?? 0600 ??"},
	{[]rune{0x1F1E6, 0x0903}, []int{0, 2}, "?? 1F1E6 ?? 0903 ??"},
	{[]rune{0x1F1E6, 0x0308, 0x0903}, []int{0, 3}, "?? 1F1E6 ?? 0308 ?? 0903 ??"},
	{[]rune{0x1F1E6, 0x1100}, []int{0, 1, 2}, "?? 1F1E6 ?? 1100 ??"},
	{[]rune{0x1F1E6, 0x0308, 0x1100}, []int{0, 2, 3}, "?? 1F1E6 ?? 0308 ?? 1100 ??"},
	{[]rune{0x1F1E6, 0x1160}, []int{0, 1, 2}, "?? 1F1E6 ?? 1160 ??"},
	{[]rune{0x1F1E6, 0x0308, 0x1160}, []int{0, 2, 3}, "?? 1F1E6 ?? 0308 ?? 1160 ??"},
	{[]rune{0x1F1E6, 0x11A8}, []int{0, 1, 2}, "?? 1F1E6 ?? 11A8 ??"},
	{[]rune{0x1F1E6, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? 1F1E6 ?? 0308 ?? 11A8 ??"},
	{[]rune{0x1F1E6, 0xAC00}, []int{0, 1, 2}, "?? 1F1E6 ?? AC00 ??"},
	{[]rune{0x1F1E6, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? 1F1E6 ?? 0308 ?? AC00 ??"},
	{[]rune{0x1F1E6, 0xAC01}, []int{0, 1, 2}, "?? 1F1E6 ?? AC01 ??"},
	{[]rune{0x1F1E6, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? 1F1E6 ?? 0308 ?? AC01 ??"},
	{[]rune{0x1F1E6, 0x231A}, []int{0, 1, 2}, "?? 1F1E6 ?? 231A ??"},
	{[]rune{0x1F1E6, 0x0308, 0x231A}, []int{0, 2, 3}, "?? 1F1E6 ?? 0308 ?? 231A ??"},
	{[]rune{0x1F1E6, 0x0300}, []int{0, 2}, "?? 1F1E6 ?? 0300 ??"},
	{[]rune{0x1F1E6, 0x0308, 0x0300}, []int{0, 3}, "?? 1F1E6 ?? 0308 ?? 0300 ??"},
	{[]rune{0x1F1E6, 0x200D}, []int{0, 2}, "?? 1F1E6 ?? 200D ??"},
	{[]rune{0x1F1E6, 0x0308, 0x200D}, []int{0, 3}, "?? 1F1E6 ?? 0308 ?? 200D ??"},
	{[]rune{0x1F1E6, 0x0378}, []int{0, 1, 2}, "?? 1F1E6 ?? 0378 ??"},
	{[]rune{0x1F1E6, 0x0308, 0x0378}, []int{0, 2, 3}, "?? 1F1E6 ?? 0308 ?? 0378 ??"},
	{[]rune{0x0600, 0x0020}, []int{0, 2}, "?? 0600 ?? 0020 ??"},
	{[]rune{0x0600, 0x0308, 0x0020}, []int{0, 2, 3}, "?? 0600 ?? 0308 ?? 0020 ??"},
	{[]rune{0x0600, 0x000D}, []int{0, 1, 2}, "?? 0600 ?? 000D ??"},
	{[]rune{0x0600, 0x0308, 0x000D}, []int{0, 2, 3}, "?? 0600 ?? 0308 ?? 000D ??"},
	{[]rune{0x0600, 0x000A}, []int{0, 1, 2}, "?? 0600 ?? 000A ??"},
	{[]rune{0x0600, 0x0308, 0x000A}, []int{0, 2, 3}, "?? 0600 ?? 0308 ?? 000A ??"},
	{[]rune{0x0600, 0x0001}, []int{0, 1, 2}, "?? 0600 ?? 0001 ??"},
	{[]rune{0x0600, 0x0308, 0x0001}, []int{0, 2, 3}, "?? 0600 ?? 0308 ?? 0001 ??"},
	{[]rune{0x0600, 0x034F}, []int{0, 2}, "?? 0600 ?? 034F ??"},
	{[]rune{0x0600, 0x0308, 0x034F}, []int{0, 3}, "?? 0600 ?? 0308 ?? 034F ??"},
	{[]rune{0x0600, 0x1F1E6}, []int{0, 2}, "?? 0600 ?? 1F1E6 ??"},
	{[]rune{0x0600, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? 0600 ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x0600, 0x0600}, []int{0, 2}, "?? 0600 ?? 0600 ??"},
	{[]rune{0x0600, 0x0308, 0x0600}, []int{0, 2, 3}, "?? 0600 ?? 0308 ?? 0600 ??"},
	{[]rune{0x0600, 0x0903}, []int{0, 2}, "?? 0600 ?? 0903 ??"},
	{[]rune{0x0600, 0x0308, 0x0903}, []int{0, 3}, "?? 0600 ?? 0308 ?? 0903 ??"},
	{[]rune{0x0600, 0x1100}, []int{0, 2}, "?? 0600 ?? 1100 ??"},
	{[]rune{0x0600, 0x0308, 0x1100}, []int{0, 2, 3}, "?? 0600 ?? 0308 ?? 1100 ??"},
	{[]rune{0x0600, 0x1160}, []int{0, 2}, "?? 0600 ?? 1160 ??"},
	{[]rune{0x0600, 0x0308, 0x1160}, []int{0, 2, 3}, "?? 0600 ?? 0308 ?? 1160 ??"},
	{[]rune{0x0600, 0x11A8}, []int{0, 2}, "?? 0600 ?? 11A8 ??"},
	{[]rune{0x0600, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? 0600 ?? 0308 ?? 11A8 ??"},
	{[]rune{0x0600, 0xAC00}, []int{0, 2}, "?? 0600 ?? AC00 ??"},
	{[]rune{0x0600, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? 0600 ?? 0308 ?? AC00 ??"},
	{[]rune{0x0600, 0xAC01}, []int{0, 2}, "?? 0600 ?? AC01 ??"},
	{[]rune{0x0600, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? 0600 ?? 0308 ?? AC01 ??"},
	{[]rune{0x0600, 0x231A}, []int{0, 2}, "?? 0600 ?? 231A ??"},
	{[]rune{0x0600, 0x0308, 0x231A}, []int{0, 2, 3}, "?? 0600 ?? 0308 ?? 231A ??"},
	{[]rune{0x0600, 0x0300}, []int{0, 2}, "?? 0600 ?? 0300 ??"},
	{[]rune{0x0600, 0x0308, 0x0300}, []int{0, 3}, "?? 0600 ?? 0308 ?? 0300 ??"},
	{[]rune{0x0600, 0x200D}, []int{0, 2}, "?? 0600 ?? 200D ??"},
	{[]rune{0x0600, 0x0308, 0x200D}, []int{0, 3}, "?? 0600 ?? 0308 ?? 200D ??"},
	{[]rune{0x0600, 0x0378}, []int{0, 2}, "?? 0600 ?? 0378 ??"},
	{[]rune{0x0600, 0x0308, 0x0378}, []int{0, 2, 3}, "?? 0600 ?? 0308 ?? 0378 ??"},
	{[]rune{0x0903, 0x0020}, []int{0, 1, 2}, "?? 0903 ?? 0020 ??"},
	{[]rune{0x0903, 0x0308, 0x0020}, []int{0, 2, 3}, "?? 0903 ?? 0308 ?? 0020 ??"},
	{[]rune{0x0903, 0x000D}, []int{0, 1, 2}, "?? 0903 ?? 000D ??"},
	{[]rune{0x0903, 0x0308, 0x000D}, []int{0, 2, 3}, "?? 0903 ?? 0308 ?? 000D ??"},
	{[]rune{0x0903, 0x000A}, []int{0, 1, 2}, "?? 0903 ?? 000A ??"},
	{[]rune{0x0903, 0x0308, 0x000A}, []int{0, 2, 3}, "?? 0903 ?? 0308 ?? 000A ??"},
	{[]rune{0x0903, 0x0001}, []int{0, 1, 2}, "?? 0903 ?? 0001 ??"},
	{[]rune{0x0903, 0x0308, 0x0001}, []int{0, 2, 3}, "?? 0903 ?? 0308 ?? 0001 ??"},
	{[]rune{0x0903, 0x034F}, []int{0, 2}, "?? 0903 ?? 034F ??"},
	{[]rune{0x0903, 0x0308, 0x034F}, []int{0, 3}, "?? 0903 ?? 0308 ?? 034F ??"},
	{[]rune{0x0903, 0x1F1E6}, []int{0, 1, 2}, "?? 0903 ?? 1F1E6 ??"},
	{[]rune{0x0903, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? 0903 ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x0903, 0x0600}, []int{0, 1, 2}, "?? 0903 ?? 0600 ??"},
	{[]rune{0x0903, 0x0308, 0x0600}, []int{0, 2, 3}, "?? 0903 ?? 0308 ?? 0600 ??"},
	{[]rune{0x0903, 0x0903}, []int{0, 2}, "?? 0903 ?? 0903 ??"},
	{[]rune{0x0903, 0x0308, 0x0903}, []int{0, 3}, "?? 0903 ?? 0308 ?? 0903 ??"},
	{[]rune{0x0903, 0x1100}, []int{0, 1, 2}, "?? 0903 ?? 1100 ??"},
	{[]rune{0x0903, 0x0308, 0x1100}, []int{0, 2, 3}, "?? 0903 ?? 0308 ?? 1100 ??"},
	{[]rune{0x0903, 0x1160}, []int{0, 1, 2}, "?? 0903 ?? 1160 ??"},
	{[]rune{0x0903, 0x0308, 0x1160}, []int{0, 2, 3}, "?? 0903 ?? 0308 ?? 1160 ??"},
	{[]rune{0x0903, 0x11A8}, []int{0, 1, 2}, "?? 0903 ?? 11A8 ??"},
	{[]rune{0x0903, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? 0903 ?? 0308 ?? 11A8 ??"},
	{[]rune{0x0903, 0xAC00}, []int{0, 1, 2}, "?? 0903 ?? AC00 ??"},
	{[]rune{0x0903, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? 0903 ?? 0308 ?? AC00 ??"},
	{[]rune{0x0903, 0xAC01}, []int{0, 1, 2}, "?? 0903 ?? AC01 ??"},
	{[]rune{0x0903, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? 0903 ?? 0308 ?? AC01 ??"},
	{[]rune{0x0903, 0x231A}, []int{0, 1, 2}, "?? 0903 ?? 231A ??"},
	{[]rune{0x0903, 0x0308, 0x231A}, []int{0, 2, 3}, "?? 0903 ?? 0308 ?? 231A ??"},
	{[]rune{0x0903, 0x0300}, []int{0, 2}, "?? 0903 ?? 0300 ??"},
	{[]rune{0x0903, 0x0308, 0x0300}, []int{0, 3}, "?? 0903 ?? 0308 ?? 0300 ??"},
	{[]rune{0x0903, 0x200D}, []int{0, 2}, "?? 0903 ?? 200D ??"},
	{[]rune{0x0903, 0x0308, 0x200D}, []int{0, 3}, "?? 0903 ?? 0308 ?? 200D ??"},
	{[]rune{0x0903, 0x0378}, []int{0, 1, 2}, "?? 0903 ?? 0378 ??"},
	{[]rune{0x0903, 0x0308, 0x0378}, []int{0, 2, 3}, "?? 0903 ?? 0308 ?? 0378 ??"},
	{[]rune{0x1100, 0x0020}, []int{0, 1, 2}, "?? 1100 ?? 0020 ??"},
	{[]rune{0x1100, 0x0308, 0x0020}, []int{0, 2, 3}, "?? 1100 ?? 0308 ?? 0020 ??"},
	{[]rune{0x1100, 0x000D}, []int{0, 1, 2}, "?? 1100 ?? 000D ??"},
	{[]rune{0x1100, 0x0308, 0x000D}, []int{0, 2, 3}, "?? 1100 ?? 0308 ?? 000D ??"},
	{[]rune{0x1100, 0x000A}, []int{0, 1, 2}, "?? 1100 ?? 000A ??"},
	{[]rune{0x1100, 0x0308, 0x000A}, []int{0, 2, 3}, "?? 1100 ?? 0308 ?? 000A ??"},
	{[]rune{0x1100, 0x0001}, []int{0, 1, 2}, "?? 1100 ?? 0001 ??"},
	{[]rune{0x1100, 0x0308, 0x0001}, []int{0, 2, 3}, "?? 1100 ?? 0308 ?? 0001 ??"},
	{[]rune{0x1100, 0x034F}, []int{0, 2}, "?? 1100 ?? 034F ??"},
	{[]rune{0x1100, 0x0308, 0x034F}, []int{0, 3}, "?? 1100 ?? 0308 ?? 034F ??"},
	{[]rune{0x1100, 0x1F1E6}, []int{0, 1, 2}, "?? 1100 ?? 1F1E6 ??"},
	{[]rune{0x1100, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? 1100 ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x1100, 0x0600}, []int{0, 1, 2}, "?? 1100 ?? 0600 ??"},
	{[]rune{0x1100, 0x0308, 0x0600}, []int{0, 2, 3}, "?? 1100 ?? 0308 ?? 0600 ??"},
	{[]rune{0x1100, 0x0903}, []int{0, 2}, "?? 1100 ?? 0903 ??"},
	{[]rune{0x1100, 0x0308, 0x0903}, []int{0, 3}, "?? 1100 ?? 0308 ?? 0903 ??"},
	{[]rune{0x1100, 0x1100}, []int{0, 2}, "?? 1100 ?? 1100 ??"},
	{[]rune{0x1100, 0x0308, 0x1100}, []int{0, 2, 3}, "?? 1100 ?? 0308 ?? 1100 ??"},
	{[]rune{0x1100, 0x1160}, []int{0, 2}, "?? 1100 ?? 1160 ??"},
	{[]rune{0x1100, 0x0308, 0x1160}, []int{0, 2, 3}, "?? 1100 ?? 0308 ?? 1160 ??"},
	{[]rune{0x1100, 0x11A8}, []int{0, 1, 2}, "?? 1100 ?? 11A8 ??"},
	{[]rune{0x1100, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? 1100 ?? 0308 ?? 11A8 ??"},
	{[]rune{0x1100, 0xAC00}, []int{0, 2}, "?? 1100 ?? AC00 ??"},
	{[]rune{0x1100, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? 1100 ?? 0308 ?? AC00 ??"},
	{[]rune{0x1100, 0xAC01}, []int{0, 2}, "?? 1100 ?? AC01 ??"},
	{[]rune{0x1100, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? 1100 ?? 0308 ?? AC01 ??"},
	{[]rune{0x1100, 0x231A}, []int{0, 1, 2}, "?? 1100 ?? 231A ??"},
	{[]rune{0x1100, 0x0308, 0x231A}, []int{0, 2, 3}, "?? 1100 ?? 0308 ?? 231A ??"},
	{[]rune{0x1100, 0x0300}, []int{0, 2}, "?? 1100 ?? 0300 ??"},
	{[]rune{0x1100, 0x0308, 0x0300}, []int{0, 3}, "?? 1100 ?? 0308 ?? 0300 ??"},
	{[]rune{0x1100, 0x200D}, []int{0, 2}, "?? 1100 ?? 200D ??"},
	{[]rune{0x1100, 0x0308, 0x200D}, []int{0, 3}, "?? 1100 ?? 0308 ?? 200D ??"},
	{[]rune{0x1100, 0x0378}, []int{0, 1, 2}, "?? 1100 ?? 0378 ??"},
	{[]rune{0x1100, 0x0308, 0x0378}, []int{0, 2, 3}, "?? 1100 ?? 0308 ?? 0378 ??"},
	{[]rune{0x1160, 0x0020}, []int{0, 1, 2}, "?? 1160 ?? 0020 ??"},
	{[]rune{0x1160, 0x0308, 0x0020}, []int{0, 2, 3}, "?? 1160 ?? 0308 ?? 0020 ??"},
	{[]rune{0x1160, 0x000D}, []int{0, 1, 2}, "?? 1160 ?? 000D ??"},
	{[]rune{0x1160, 0x0308, 0x000D}, []int{0, 2, 3}, "?? 1160 ?? 0308 ?? 000D ??"},
	{[]rune{0x1160, 0x000A}, []int{0, 1, 2}, "?? 1160 ?? 000A ??"},
	{[]rune{0x1160, 0x0308, 0x000A}, []int{0, 2, 3}, "?? 1160 ?? 0308 ?? 000A ??"},
	{[]rune{0x1160, 0x0001}, []int{0, 1, 2}, "?? 1160 ?? 0001 ??"},
	{[]rune{0x1160, 0x0308, 0x0001}, []int{0, 2, 3}, "?? 1160 ?? 0308 ?? 0001 ??"},
	{[]rune{0x1160, 0x034F}, []int{0, 2}, "?? 1160 ?? 034F ??"},
	{[]rune{0x1160, 0x0308, 0x034F}, []int{0, 3}, "?? 1160 ?? 0308 ?? 034F ??"},
	{[]rune{0x1160, 0x1F1E6}, []int{0, 1, 2}, "?? 1160 ?? 1F1E6 ??"},
	{[]rune{0x1160, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? 1160 ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x1160, 0x0600}, []int{0, 1, 2}, "?? 1160 ?? 0600 ??"},
	{[]rune{0x1160, 0x0308, 0x0600}, []int{0, 2, 3}, "?? 1160 ?? 0308 ?? 0600 ??"},
	{[]rune{0x1160, 0x0903}, []int{0, 2}, "?? 1160 ?? 0903 ??"},
	{[]rune{0x1160, 0x0308, 0x0903}, []int{0, 3}, "?? 1160 ?? 0308 ?? 0903 ??"},
	{[]rune{0x1160, 0x1100}, []int{0, 1, 2}, "?? 1160 ?? 1100 ??"},
	{[]rune{0x1160, 0x0308, 0x1100}, []int{0, 2, 3}, "?? 1160 ?? 0308 ?? 1100 ??"},
	{[]rune{0x1160, 0x1160}, []int{0, 2}, "?? 1160 ?? 1160 ??"},
	{[]rune{0x1160, 0x0308, 0x1160}, []int{0, 2, 3}, "?? 1160 ?? 0308 ?? 1160 ??"},
	{[]rune{0x1160, 0x11A8}, []int{0, 2}, "?? 1160 ?? 11A8 ??"},
	{[]rune{0x1160, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? 1160 ?? 0308 ?? 11A8 ??"},
	{[]rune{0x1160, 0xAC00}, []int{0, 1, 2}, "?? 1160 ?? AC00 ??"},
	{[]rune{0x1160, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? 1160 ?? 0308 ?? AC00 ??"},
	{[]rune{0x1160, 0xAC01}, []int{0, 1, 2}, "?? 1160 ?? AC01 ??"},
	{[]rune{0x1160, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? 1160 ?? 0308 ?? AC01 ??"},
	{[]rune{0x1160, 0x231A}, []int{0, 1, 2}, "?? 1160 ?? 231A ??"},
	{[]rune{0x1160, 0x0308, 0x231A}, []int{0, 2, 3}, "?? 1160 ?? 0308 ?? 231A ??"},
	{[]rune{0x1160, 0x0300}, []int{0, 2}, "?? 1160 ?? 0300 ??"},
	{[]rune{0x1160, 0x0308, 0x0300}, []int{0, 3}, "?? 1160 ?? 0308 ?? 0300 ??"},
	{[]rune{0x1160, 0x200D}, []int{0, 2}, "?? 1160 ?? 200D ??"},
	{[]rune{0x1160, 0x0308, 0x200D}, []int{0, 3}, "?? 1160 ?? 0308 ?? 200D ??"},
	{[]rune{0x1160, 0x0378}, []int{0, 1, 2}, "?? 1160 ?? 0378 ??"},
	{[]rune{0x1160, 0x0308, 0x0378}, []int{0, 2, 3}, "?? 1160 ?? 0308 ?? 0378 ??"},
	{[]rune{0x11A8, 0x0020}, []int{0, 1, 2}, "?? 11A8 ?? 0020 ??"},
	{[]rune{0x11A8, 0x0308, 0x0020}, []int{0, 2, 3}, "?? 11A8 ?? 0308 ?? 0020 ??"},
	{[]rune{0x11A8, 0x000D}, []int{0, 1, 2}, "?? 11A8 ?? 000D ??"},
	{[]rune{0x11A8, 0x0308, 0x000D}, []int{0, 2, 3}, "?? 11A8 ?? 0308 ?? 000D ??"},
	{[]rune{0x11A8, 0x000A}, []int{0, 1, 2}, "?? 11A8 ?? 000A ??"},
	{[]rune{0x11A8, 0x0308, 0x000A}, []int{0, 2, 3}, "?? 11A8 ?? 0308 ?? 000A ??"},
	{[]rune{0x11A8, 0x0001}, []int{0, 1, 2}, "?? 11A8 ?? 0001 ??"},
	{[]rune{0x11A8, 0x0308, 0x0001}, []int{0, 2, 3}, "?? 11A8 ?? 0308 ?? 0001 ??"},
	{[]rune{0x11A8, 0x034F}, []int{0, 2}, "?? 11A8 ?? 034F ??"},
	{[]rune{0x11A8, 0x0308, 0x034F}, []int{0, 3}, "?? 11A8 ?? 0308 ?? 034F ??"},
	{[]rune{0x11A8, 0x1F1E6}, []int{0, 1, 2}, "?? 11A8 ?? 1F1E6 ??"},
	{[]rune{0x11A8, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? 11A8 ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x11A8, 0x0600}, []int{0, 1, 2}, "?? 11A8 ?? 0600 ??"},
	{[]rune{0x11A8, 0x0308, 0x0600}, []int{0, 2, 3}, "?? 11A8 ?? 0308 ?? 0600 ??"},
	{[]rune{0x11A8, 0x0903}, []int{0, 2}, "?? 11A8 ?? 0903 ??"},
	{[]rune{0x11A8, 0x0308, 0x0903}, []int{0, 3}, "?? 11A8 ?? 0308 ?? 0903 ??"},
	{[]rune{0x11A8, 0x1100}, []int{0, 1, 2}, "?? 11A8 ?? 1100 ??"},
	{[]rune{0x11A8, 0x0308, 0x1100}, []int{0, 2, 3}, "?? 11A8 ?? 0308 ?? 1100 ??"},
	{[]rune{0x11A8, 0x1160}, []int{0, 1, 2}, "?? 11A8 ?? 1160 ??"},
	{[]rune{0x11A8, 0x0308, 0x1160}, []int{0, 2, 3}, "?? 11A8 ?? 0308 ?? 1160 ??"},
	{[]rune{0x11A8, 0x11A8}, []int{0, 2}, "?? 11A8 ?? 11A8 ??"},
	{[]rune{0x11A8, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? 11A8 ?? 0308 ?? 11A8 ??"},
	{[]rune{0x11A8, 0xAC00}, []int{0, 1, 2}, "?? 11A8 ?? AC00 ??"},
	{[]rune{0x11A8, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? 11A8 ?? 0308 ?? AC00 ??"},
	{[]rune{0x11A8, 0xAC01}, []int{0, 1, 2}, "?? 11A8 ?? AC01 ??"},
	{[]rune{0x11A8, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? 11A8 ?? 0308 ?? AC01 ??"},
	{[]rune{0x11A8, 0x231A}, []int{0, 1, 2}, "?? 11A8 ?? 231A ??"},
	{[]rune{0x11A8, 0x0308, 0x231A}, []int{0, 2, 3}, "?? 11A8 ?? 0308 ?? 231A ??"},
	{[]rune{0x11A8, 0x0300}, []int{0, 2}, "?? 11A8 ?? 0300 ??"},
	{[]rune{0x11A8, 0x0308, 0x0300}, []int{0, 3}, "?? 11A8 ?? 0308 ?? 0300 ??"},
	{[]rune{0x11A8, 0x200D}, []int{0, 2}, "?? 11A8 ?? 200D ??"},
	{[]rune{0x11A8, 0x0308, 0x200D}, []int{0, 3}, "?? 11A8 ?? 0308 ?? 200D ??"},
	{[]rune{0x11A8, 0x0378}, []int{0, 1, 2}, "?? 11A8 ?? 0378 ??"},
	{[]rune{0x11A8, 0x0308, 0x0378}, []int{0, 2, 3}, "?? 11A8 ?? 0308 ?? 0378 ??"},
	{[]rune{0xAC00, 0x0020}, []int{0, 1, 2}, "?? AC00 ?? 0020 ??"},
	{[]rune{0xAC00, 0x0308, 0x0020}, []int{0, 2, 3}, "?? AC00 ?? 0308 ?? 0020 ??"},
	{[]rune{0xAC00, 0x000D}, []int{0, 1, 2}, "?? AC00 ?? 000D ??"},
	{[]rune{0xAC00, 0x0308, 0x000D}, []int{0, 2, 3}, "?? AC00 ?? 0308 ?? 000D ??"},
	{[]rune{0xAC00, 0x000A}, []int{0, 1, 2}, "?? AC00 ?? 000A ??"},
	{[]rune{0xAC00, 0x0308, 0x000A}, []int{0, 2, 3}, "?? AC00 ?? 0308 ?? 000A ??"},
	{[]rune{0xAC00, 0x0001}, []int{0, 1, 2}, "?? AC00 ?? 0001 ??"},
	{[]rune{0xAC00, 0x0308, 0x0001}, []int{0, 2, 3}, "?? AC00 ?? 0308 ?? 0001 ??"},
	{[]rune{0xAC00, 0x034F}, []int{0, 2}, "?? AC00 ?? 034F ??"},
	{[]rune{0xAC00, 0x0308, 0x034F}, []int{0, 3}, "?? AC00 ?? 0308 ?? 034F ??"},
	{[]rune{0xAC00, 0x1F1E6}, []int{0, 1, 2}, "?? AC00 ?? 1F1E6 ??"},
	{[]rune{0xAC00, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? AC00 ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0xAC00, 0x0600}, []int{0, 1, 2}, "?? AC00 ?? 0600 ??"},
	{[]rune{0xAC00, 0x0308, 0x0600}, []int{0, 2, 3}, "?? AC00 ?? 0308 ?? 0600 ??"},
	{[]rune{0xAC00, 0x0903}, []int{0, 2}, "?? AC00 ?? 0903 ??"},
	{[]rune{0xAC00, 0x0308, 0x0903}, []int{0, 3}, "?? AC00 ?? 0308 ?? 0903 ??"},
	{[]rune{0xAC00, 0x1100}, []int{0, 1, 2}, "?? AC00 ?? 1100 ??"},
	{[]rune{0xAC00, 0x0308, 0x1100}, []int{0, 2, 3}, "?? AC00 ?? 0308 ?? 1100 ??"},
	{[]rune{0xAC00, 0x1160}, []int{0, 2}, "?? AC00 ?? 1160 ??"},
	{[]rune{0xAC00, 0x0308, 0x1160}, []int{0, 2, 3}, "?? AC00 ?? 0308 ?? 1160 ??"},
	{[]rune{0xAC00, 0x11A8}, []int{0, 2}, "?? AC00 ?? 11A8 ??"},
	{[]rune{0xAC00, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? AC00 ?? 0308 ?? 11A8 ??"},
	{[]rune{0xAC00, 0xAC00}, []int{0, 1, 2}, "?? AC00 ?? AC00 ??"},
	{[]rune{0xAC00, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? AC00 ?? 0308 ?? AC00 ??"},
	{[]rune{0xAC00, 0xAC01}, []int{0, 1, 2}, "?? AC00 ?? AC01 ??"},
	{[]rune{0xAC00, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? AC00 ?? 0308 ?? AC01 ??"},
	{[]rune{0xAC00, 0x231A}, []int{0, 1, 2}, "?? AC00 ?? 231A ??"},
	{[]rune{0xAC00, 0x0308, 0x231A}, []int{0, 2, 3}, "?? AC00 ?? 0308 ?? 231A ??"},
	{[]rune{0xAC00, 0x0300}, []int{0, 2}, "?? AC00 ?? 0300 ??"},
	{[]rune{0xAC00, 0x0308, 0x0300}, []int{0, 3}, "?? AC00 ?? 0308 ?? 0300 ??"},
	{[]rune{0xAC00, 0x200D}, []int{0, 2}, "?? AC00 ?? 200D ??"},
	{[]rune{0xAC00, 0x0308, 0x200D}, []int{0, 3}, "?? AC00 ?? 0308 ?? 200D ??"},
	{[]rune{0xAC00, 0x0378}, []int{0, 1, 2}, "?? AC00 ?? 0378 ??"},
	{[]rune{0xAC00, 0x0308, 0x0378}, []int{0, 2, 3}, "?? AC00 ?? 0308 ?? 0378 ??"},
	{[]rune{0xAC01, 0x0020}, []int{0, 1, 2}, "?? AC01 ?? 0020 ??"},
	{[]rune{0xAC01, 0x0308, 0x0020}, []int{0, 2, 3}, "?? AC01 ?? 0308 ?? 0020 ??"},
	{[]rune{0xAC01, 0x000D}, []int{0, 1, 2}, "?? AC01 ?? 000D ??"},
	{[]rune{0xAC01, 0x0308, 0x000D}, []int{0, 2, 3}, "?? AC01 ?? 0308 ?? 000D ??"},
	{[]rune{0xAC01, 0x000A}, []int{0, 1, 2}, "?? AC01 ?? 000A ??"},
	{[]rune{0xAC01, 0x0308, 0x000A}, []int{0, 2, 3}, "?? AC01 ?? 0308 ?? 000A ??"},
	{[]rune{0xAC01, 0x0001}, []int{0, 1, 2}, "?? AC01 ?? 0001 ??"},
	{[]rune{0xAC01, 0x0308, 0x0001}, []int{0, 2, 3}, "?? AC01 ?? 0308 ?? 0001 ??"},
	{[]rune{0xAC01, 0x034F}, []int{0, 2}, "?? AC01 ?? 034F ??"},
	{[]rune{0xAC01, 0x0308, 0x034F}, []int{0, 3}, "?? AC01 ?? 0308 ?? 034F ??"},
	{[]rune{0xAC01, 0x1F1E6}, []int{0, 1, 2}, "?? AC01 ?? 1F1E6 ??"},
	{[]rune{0xAC01, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? AC01 ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0xAC01, 0x0600}, []int{0, 1, 2}, "?? AC01 ?? 0600 ??"},
	{[]rune{0xAC01, 0x0308, 0x0600}, []int{0, 2, 3}, "?? AC01 ?? 0308 ?? 0600 ??"},
	{[]rune{0xAC01, 0x0903}, []int{0, 2}, "?? AC01 ?? 0903 ??"},
	{[]rune{0xAC01, 0x0308, 0x0903}, []int{0, 3}, "?? AC01 ?? 0308 ?? 0903 ??"},
	{[]rune{0xAC01, 0x1100}, []int{0, 1, 2}, "?? AC01 ?? 1100 ??"},
	{[]rune{0xAC01, 0x0308, 0x1100}, []int{0, 2, 3}, "?? AC01 ?? 0308 ?? 1100 ??"},
	{[]rune{0xAC01, 0x1160}, []int{0, 1, 2}, "?? AC01 ?? 1160 ??"},
	{[]rune{0xAC01, 0x0308, 0x1160}, []int{0, 2, 3}, "?? AC01 ?? 0308 ?? 1160 ??"},
	{[]rune{0xAC01, 0x11A8}, []int{0, 2}, "?? AC01 ?? 11A8 ??"},
	{[]rune{0xAC01, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? AC01 ?? 0308 ?? 11A8 ??"},
	{[]rune{0xAC01, 0xAC00}, []int{0, 1, 2}, "?? AC01 ?? AC00 ??"},
	{[]rune{0xAC01, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? AC01 ?? 0308 ?? AC00 ??"},
	{[]rune{0xAC01, 0xAC01}, []int{0, 1, 2}, "?? AC01 ?? AC01 ??"},
	{[]rune{0xAC01, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? AC01 ?? 0308 ?? AC01 ??"},
	{[]rune{0xAC01, 0x231A}, []int{0, 1, 2}, "?? AC01 ?? 231A ??"},
	{[]rune{0xAC01, 0x0308, 0x231A}, []int{0, 2, 3}, "?? AC01 ?? 0308 ?? 231A ??"},
	{[]rune{0xAC01, 0x0300}, []int{0, 2}, "?? AC01 ?? 0300 ??"},
	{[]rune{0xAC01, 0x0308, 0x0300}, []int{0, 3}, "?? AC01 ?? 0308 ?? 0300 ??"},
	{[]rune{0xAC01, 0x200D}, []int{0, 2}, "?? AC01 ?? 200D ??"},
	{[]rune{0xAC01, 0x0308, 0x200D}, []int{0, 3}, "?? AC01 ?? 0308 ?? 200D ??"},
	{[]rune{0xAC01, 0x0378}, []int{0, 1, 2}, "?? AC01 ?? 0378 ??"},
	{[]rune{0xAC01, 0x0308, 0x0378}, []int{0, 2, 3}, "?? AC01 ?? 0308 ?? 0378 ??"},
	{[]rune{0x231A, 0x0020}, []int{0, 1, 2}, "?? 231A ?? 0020 ??"},
	{[]rune{0x231A, 0x0308, 0x0020}, []int{0, 2, 3}, "?? 231A ?? 0308 ?? 0020 ??"},
	{[]rune{0x231A, 0x000D}, []int{0, 1, 2}, "?? 231A ?? 000D ??"},
	{[]rune{0x231A, 0x0308, 0x000D}, []int{0, 2, 3}, "?? 231A ?? 0308 ?? 000D ??"},
	{[]rune{0x231A, 0x000A}, []int{0, 1, 2}, "?? 231A ?? 000A ??"},
	{[]rune{0x231A, 0x0308, 0x000A}, []int{0, 2, 3}, "?? 231A ?? 0308 ?? 000A ??"},
	{[]rune{0x231A, 0x0001}, []int{0, 1, 2}, "?? 231A ?? 0001 ??"},
	{[]rune{0x231A, 0x0308, 0x0001}, []int{0, 2, 3}, "?? 231A ?? 0308 ?? 0001 ??"},
	{[]rune{0x231A, 0x034F}, []int{0, 2}, "?? 231A ?? 034F ??"},
	{[]rune{0x231A, 0x0308, 0x034F}, []int{0, 3}, "?? 231A ?? 0308 ?? 034F ??"},
	{[]rune{0x231A, 0x1F1E6}, []int{0, 1, 2}, "?? 231A ?? 1F1E6 ??"},
	{[]rune{0x231A, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? 231A ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x231A, 0x0600}, []int{0, 1, 2}, "?? 231A ?? 0600 ??"},
	{[]rune{0x231A, 0x0308, 0x0600}, []int{0, 2, 3}, "?? 231A ?? 0308 ?? 0600 ??"},
	{[]rune{0x231A, 0x0903}, []int{0, 2}, "?? 231A ?? 0903 ??"},
	{[]rune{0x231A, 0x0308, 0x0903}, []int{0, 3}, "?? 231A ?? 0308 ?? 0903 ??"},
	{[]rune{0x231A, 0x1100}, []int{0, 1, 2}, "?? 231A ?? 1100 ??"},
	{[]rune{0x231A, 0x0308, 0x1100}, []int{0, 2, 3}, "?? 231A ?? 0308 ?? 1100 ??"},
	{[]rune{0x231A, 0x1160}, []int{0, 1, 2}, "?? 231A ?? 1160 ??"},
	{[]rune{0x231A, 0x0308, 0x1160}, []int{0, 2, 3}, "?? 231A ?? 0308 ?? 1160 ??"},
	{[]rune{0x231A, 0x11A8}, []int{0, 1, 2}, "?? 231A ?? 11A8 ??"},
	{[]rune{0x231A, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? 231A ?? 0308 ?? 11A8 ??"},
	{[]rune{0x231A, 0xAC00}, []int{0, 1, 2}, "?? 231A ?? AC00 ??"},
	{[]rune{0x231A, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? 231A ?? 0308 ?? AC00 ??"},
	{[]rune{0x231A, 0xAC01}, []int{0, 1, 2}, "?? 231A ?? AC01 ??"},
	{[]rune{0x231A, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? 231A ?? 0308 ?? AC01 ??"},
	{[]rune{0x231A, 0x231A}, []int{0, 1, 2}, "?? 231A ?? 231A ??"},
	{[]rune{0x231A, 0x0308, 0x231A}, []int{0, 2, 3}, "?? 231A ?? 0308 ?? 231A ??"},
	{[]rune{0x231A, 0x0300}, []int{0, 2}, "?? 231A ?? 0300 ??"},
	{[]rune{0x231A, 0x0308, 0x0300}, []int{0, 3}, "?? 231A ?? 0308 ?? 0300 ??"},
	{[]rune{0x231A, 0x200D}, []int{0, 2}, "?? 231A ?? 200D ??"},
	{[]rune{0x231A, 0x0308, 0x200D}, []int{0, 3}, "?? 231A ?? 0308 ?? 200D ??"},
	{[]rune{0x231A, 0x0378}, []int{0, 1, 2}, "?? 231A ?? 0378 ??"},
	{[]rune{0x231A, 0x0308, 0x0378}, []int{0, 2, 3}, "?? 231A ?? 0308 ?? 0378 ??"},
	{[]rune{0x0300, 0x0020}, []int{0, 1, 2}, "?? 0300 ?? 0020 ??"},
	{[]rune{0x0300, 0x0308, 0x0020}, []int{0, 2, 3}, "?? 0300 ?? 0308 ?? 0020 ??"},
	{[]rune{0x0300, 0x000D}, []int{0, 1, 2}, "?? 0300 ?? 000D ??"},
	{[]rune{0x0300, 0x0308, 0x000D}, []int{0, 2, 3}, "?? 0300 ?? 0308 ?? 000D ??"},
	{[]rune{0x0300, 0x000A}, []int{0, 1, 2}, "?? 0300 ?? 000A ??"},
	{[]rune{0x0300, 0x0308, 0x000A}, []int{0, 2, 3}, "?? 0300 ?? 0308 ?? 000A ??"},
	{[]rune{0x0300, 0x0001}, []int{0, 1, 2}, "?? 0300 ?? 0001 ??"},
	{[]rune{0x0300, 0x0308, 0x0001}, []int{0, 2, 3}, "?? 0300 ?? 0308 ?? 0001 ??"},
	{[]rune{0x0300, 0x034F}, []int{0, 2}, "?? 0300 ?? 034F ??"},
	{[]rune{0x0300, 0x0308, 0x034F}, []int{0, 3}, "?? 0300 ?? 0308 ?? 034F ??"},
	{[]rune{0x0300, 0x1F1E6}, []int{0, 1, 2}, "?? 0300 ?? 1F1E6 ??"},
	{[]rune{0x0300, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? 0300 ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x0300, 0x0600}, []int{0, 1, 2}, "?? 0300 ?? 0600 ??"},
	{[]rune{0x0300, 0x0308, 0x0600}, []int{0, 2, 3}, "?? 0300 ?? 0308 ?? 0600 ??"},
	{[]rune{0x0300, 0x0903}, []int{0, 2}, "?? 0300 ?? 0903 ??"},
	{[]rune{0x0300, 0x0308, 0x0903}, []int{0, 3}, "?? 0300 ?? 0308 ?? 0903 ??"},
	{[]rune{0x0300, 0x1100}, []int{0, 1, 2}, "?? 0300 ?? 1100 ??"},
	{[]rune{0x0300, 0x0308, 0x1100}, []int{0, 2, 3}, "?? 0300 ?? 0308 ?? 1100 ??"},
	{[]rune{0x0300, 0x1160}, []int{0, 1, 2}, "?? 0300 ?? 1160 ??"},
	{[]rune{0x0300, 0x0308, 0x1160}, []int{0, 2, 3}, "?? 0300 ?? 0308 ?? 1160 ??"},
	{[]rune{0x0300, 0x11A8}, []int{0, 1, 2}, "?? 0300 ?? 11A8 ??"},
	{[]rune{0x0300, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? 0300 ?? 0308 ?? 11A8 ??"},
	{[]rune{0x0300, 0xAC00}, []int{0, 1, 2}, "?? 0300 ?? AC00 ??"},
	{[]rune{0x0300, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? 0300 ?? 0308 ?? AC00 ??"},
	{[]rune{0x0300, 0xAC01}, []int{0, 1, 2}, "?? 0300 ?? AC01 ??"},
	{[]rune{0x0300, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? 0300 ?? 0308 ?? AC01 ??"},
	{[]rune{0x0300, 0x231A}, []int{0, 1, 2}, "?? 0300 ?? 231A ??"},
	{[]rune{0x0300, 0x0308, 0x231A}, []int{0, 2, 3}, "?? 0300 ?? 0308 ?? 231A ??"},
	{[]rune{0x0300, 0x0300}, []int{0, 2}, "?? 0300 ?? 0300 ??"},
	{[]rune{0x0300, 0x0308, 0x0300}, []int{0, 3}, "?? 0300 ?? 0308 ?? 0300 ??"},
	{[]rune{0x0300, 0x200D}, []int{0, 2}, "?? 0300 ?? 200D ??"},
	{[]rune{0x0300, 0x0308, 0x200D}, []int{0, 3}, "?? 0300 ?? 0308 ?? 200D ??"},
	{[]rune{0x0300, 0x0378}, []int{0, 1, 2}, "?? 0300 ?? 0378 ??"},
	{[]rune{0x0300, 0x0308, 0x0378}, []int{0, 2, 3}, "?? 0300 ?? 0308 ?? 0378 ??"},
	{[]rune{0x200D, 0x0020}, []int{0, 1, 2}, "?? 200D ?? 0020 ??"},
	{[]rune{0x200D, 0x0308, 0x0020}, []int{0, 2, 3}, "?? 200D ?? 0308 ?? 0020 ??"},
	{[]rune{0x200D, 0x000D}, []int{0, 1, 2}, "?? 200D ?? 000D ??"},
	{[]rune{0x200D, 0x0308, 0x000D}, []int{0, 2, 3}, "?? 200D ?? 0308 ?? 000D ??"},
	{[]rune{0x200D, 0x000A}, []int{0, 1, 2}, "?? 200D ?? 000A ??"},
	{[]rune{0x200D, 0x0308, 0x000A}, []int{0, 2, 3}, "?? 200D ?? 0308 ?? 000A ??"},
	{[]rune{0x200D, 0x0001}, []int{0, 1, 2}, "?? 200D ?? 0001 ??"},
	{[]rune{0x200D, 0x0308, 0x0001}, []int{0, 2, 3}, "?? 200D ?? 0308 ?? 0001 ??"},
	{[]rune{0x200D, 0x034F}, []int{0, 2}, "?? 200D ?? 034F ??"},
	{[]rune{0x200D, 0x0308, 0x034F}, []int{0, 3}, "?? 200D ?? 0308 ?? 034F ??"},
	{[]rune{0x200D, 0x1F1E6}, []int{0, 1, 2}, "?? 200D ?? 1F1E6 ??"},
	{[]rune{0x200D, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? 200D ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x200D, 0x0600}, []int{0, 1, 2}, "?? 200D ?? 0600 ??"},
	{[]rune{0x200D, 0x0308, 0x0600}, []int{0, 2, 3}, "?? 200D ?? 0308 ?? 0600 ??"},
	{[]rune{0x200D, 0x0903}, []int{0, 2}, "?? 200D ?? 0903 ??"},
	{[]rune{0x200D, 0x0308, 0x0903}, []int{0, 3}, "?? 200D ?? 0308 ?? 0903 ??"},
	{[]rune{0x200D, 0x1100}, []int{0, 1, 2}, "?? 200D ?? 1100 ??"},
	{[]rune{0x200D, 0x0308, 0x1100}, []int{0, 2, 3}, "?? 200D ?? 0308 ?? 1100 ??"},
	{[]rune{0x200D, 0x1160}, []int{0, 1, 2}, "?? 200D ?? 1160 ??"},
	{[]rune{0x200D, 0x0308, 0x1160}, []int{0, 2, 3}, "?? 200D ?? 0308 ?? 1160 ??"},
	{[]rune{0x200D, 0x11A8}, []int{0, 1, 2}, "?? 200D ?? 11A8 ??"},
	{[]rune{0x200D, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? 200D ?? 0308 ?? 11A8 ??"},
	{[]rune{0x200D, 0xAC00}, []int{0, 1, 2}, "?? 200D ?? AC00 ??"},
	{[]rune{0x200D, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? 200D ?? 0308 ?? AC00 ??"},
	{[]rune{0x200D, 0xAC01}, []int{0, 1, 2}, "?? 200D ?? AC01 ??"},
	{[]rune{0x200D, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? 200D ?? 0308 ?? AC01 ??"},
	{[]rune{0x200D, 0x231A}, []int{0, 1, 2}, "?? 200D ?? 231A ??"},
	{[]rune{0x200D, 0x0308, 0x231A}, []int{0, 2, 3}, "?? 200D ?? 0308 ?? 231A ??"},
	{[]rune{0x200D, 0x0300}, []int{0, 2}, "?? 200D ?? 0300 ??"},
	{[]rune{0x200D, 0x0308, 0x0300}, []int{0, 3}, "?? 200D ?? 0308 ?? 0300 ??"},
	{[]rune{0x200D, 0x200D}, []int{0, 2}, "?? 200D ?? 200D ??"},
	{[]rune{0x200D, 0x0308, 0x200D}, []int{0, 3}, "?? 200D ?? 0308 ?? 200D ??"},
	{[]rune{0x200D, 0x0378}, []int{0, 1, 2}, "?? 200D ?? 0378 ??"},
	{[]rune{0x200D, 0x0308, 0x0378}, []int{0, 2, 3}, "?? 200D ?? 0308 ?? 0378 ??"},
	{[]rune{0x0378, 0x0020}, []int{0, 1, 2}, "?? 0378 ?? 0020 ??"},
	{[]rune{0x0378, 0x0308, 0x0020}, []int{0, 2, 3}, "?? 0378 ?? 0308 ?? 0020 ??"},
	{[]rune{0x0378, 0x000D}, []int{0, 1, 2}, "?? 0378 ?? 000D ??"},
	{[]rune{0x0378, 0x0308, 0x000D}, []int{0, 2, 3}, "?? 0378 ?? 0308 ?? 000D ??"},
	{[]rune{0x0378, 0x000A}, []int{0, 1, 2}, "?? 0378 ?? 000A ??"},
	{[]rune{0x0378, 0x0308, 0x000A}, []int{0, 2, 3}, "?? 0378 ?? 0308 ?? 000A ??"},
	{[]rune{0x0378, 0x0001}, []int{0, 1, 2}, "?? 0378 ?? 0001 ??"},
	{[]rune{0x0378, 0x0308, 0x0001}, []int{0, 2, 3}, "?? 0378 ?? 0308 ?? 0001 ??"},
	{[]rune{0x0378, 0x034F}, []int{0, 2}, "?? 0378 ?? 034F ??"},
	{[]rune{0x0378, 0x0308, 0x034F}, []int{0, 3}, "?? 0378 ?? 0308 ?? 034F ??"},
	{[]rune{0x0378, 0x1F1E6}, []int{0, 1, 2}, "?? 0378 ?? 1F1E6 ??"},
	{[]rune{0x0378, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "?? 0378 ?? 0308 ?? 1F1E6 ??"},
	{[]rune{0x0378, 0x0600}, []int{0, 1, 2}, "?? 0378 ?? 0600 ??"},
	{[]rune{0x0378, 0x0308, 0x0600}, []int{0, 2, 3}, "?? 0378 ?? 0308 ?? 0600 ??"},
	{[]rune{0x0378, 0x0903}, []int{0, 2}, "?? 0378 ?? 0903 ??"},
	{[]rune{0x0378, 0x0308, 0x0903}, []int{0, 3}, "?? 0378 ?? 0308 ?? 0903 ??"},
	{[]rune{0x0378, 0x1100}, []int{0, 1, 2}, "?? 0378 ?? 1100 ??"},
	{[]rune{0x0378, 0x0308, 0x1100}, []int{0, 2, 3}, "?? 0378 ?? 0308 ?? 1100 ??"},
	{[]rune{0x0378, 0x1160}, []int{0, 1, 2}, "?? 0378 ?? 1160 ??"},
	{[]rune{0x0378, 0x0308, 0x1160}, []int{0, 2, 3}, "?? 0378 ?? 0308 ?? 1160 ??"},
	{[]rune{0x0378, 0x11A8}, []int{0, 1, 2}, "?? 0378 ?? 11A8 ??"},
	{[]rune{0x0378, 0x0308, 0x11A8}, []int{0, 2, 3}, "?? 0378 ?? 0308 ?? 11A8 ??"},
	{[]rune{0x0378, 0xAC00}, []int{0, 1, 2}, "?? 0378 ?? AC00 ??"},
	{[]rune{0x0378, 0x0308, 0xAC00}, []int{0, 2, 3}, "?? 0378 ?? 0308 ?? AC00 ??"},
	{[]rune{0x0378, 0xAC01}, []int{0, 1, 2}, "?? 0378 ?? AC01 ??"},
	{[]rune{0x0378, 0x0308, 0xAC01}, []int{0, 2, 3}, "?? 0378 ?? 0308 ?? AC01 ??"},
	{[]rune{0x0378, 0x231A}, []int{0, 1, 2}, "?? 0378 ?? 231A ??"},
	{[]rune{0x0378, 0x0308, 0x231A}, []int{0, 2, 3}, "?? 0378 ?? 0308 ?? 231A ??"},
	{[]rune{0x0378, 0x0300}, []int{0, 2}, "?? 0378 ?? 0300 ??"},
	{[]rune{0x0378, 0x0308, 0x0300}, []int{0, 3}, "?? 0378 ?? 0308 ?? 0300 ??"},
	{[]rune{0x0378, 0x200D}, []int{0, 2}, "?? 0378 ?? 200D ??"},
	{[]rune{0x0378, 0x0308, 0x200D}, []int{0, 3}, "?? 0378 ?? 0308 ?? 200D ??"},
	{[]rune{0x0378, 0x0378}, []int{0, 1, 2}, "?? 0378 ?? 0378 ??"},
	{[]rune{0x0378, 0x0308, 0x0378}, []int{0, 2, 3}, "?? 0378 ?? 0308 ?? 0378 ??"},
	{[]rune{0x000D, 0x000A, 0x0061, 0x000A, 0x0308}, []int{0, 2, 3, 4, 5}, "?? 000D ?? 000A ?? 0061 ?? 000A ?? 0308 ??"},
	{[]rune{0x0061, 0x0308}, []int{0, 2}, "?? 0061 ?? 0308 ??"},
	{[]rune{0x0020, 0x200D, 0x0646}, []int{0, 2, 3}, "?? 0020 ?? 200D ?? 0646 ??"},
	{[]rune{0x0646, 0x200D, 0x0020}, []int{0, 2, 3}, "?? 0646 ?? 200D ?? 0020 ??"},
	{[]rune{0x1100, 0x1100}, []int{0, 2}, "?? 1100 ?? 1100 ??"},
	{[]rune{0xAC00, 0x11A8, 0x1100}, []int{0, 2, 3}, "?? AC00 ?? 11A8 ?? 1100 ??"},
	{[]rune{0xAC01, 0x11A8, 0x1100}, []int{0, 2, 3}, "?? AC01 ?? 11A8 ?? 1100 ??"},
	{[]rune{0x1F1E6, 0x1F1E7, 0x1F1E8, 0x0062}, []int{0, 2, 3, 4}, "?? 1F1E6 ?? 1F1E7 ?? 1F1E8 ?? 0062 ??"},
	{[]rune{0x0061, 0x1F1E6, 0x1F1E7, 0x1F1E8, 0x0062}, []int{0, 1, 3, 4, 5}, "?? 0061 ?? 1F1E6 ?? 1F1E7 ?? 1F1E8 ?? 0062 ??"},
	{[]rune{0x0061, 0x1F1E6, 0x1F1E7, 0x200D, 0x1F1E8, 0x0062}, []int{0, 1, 4, 5, 6}, "?? 0061 ?? 1F1E6 ?? 1F1E7 ?? 200D ?? 1F1E8 ?? 0062 ??"},
	{[]rune{0x0061, 0x1F1E6, 0x200D, 0x1F1E7, 0x1F1E8, 0x0062}, []int{0, 1, 3, 5, 6}, "?? 0061 ?? 1F1E6 ?? 200D ?? 1F1E7 ?? 1F1E8 ?? 0062 ??"},
	{[]rune{0x0061, 0x1F1E6, 0x1F1E7, 0x1F1E8, 0x1F1E9, 0x0062}, []int{0, 1, 3, 5, 6}, "?? 0061 ?? 1F1E6 ?? 1F1E7 ?? 1F1E8 ?? 1F1E9 ?? 0062 ??"},
	{[]rune{0x0061, 0x200D}, []int{0, 2}, "?? 0061 ?? 200D ??"},
	{[]rune{0x0061, 0x0308, 0x0062}, []int{0, 2, 3}, "?? 0061 ?? 0308 ?? 0062 ??"},
	{[]rune{0x0061, 0x0903, 0x0062}, []int{0, 2, 3}, "?? 0061 ?? 0903 ?? 0062 ??"},
	{[]rune{0x0061, 0x0600, 0x0062}, []int{0, 1, 3}, "?? 0061 ?? 0600 ?? 0062 ??"},
	{[]rune{0x1F476, 0x1F3FF, 0x1F476}, []int{0, 2, 3}, "?? 1F476 ?? 1F3FF ?? 1F476 ??"},
	{[]rune{0x0061, 0x1F3FF, 0x1F476}, []int{0, 2, 3}, "?? 0061 ?? 1F3FF ?? 1F476 ??"},
	{[]rune{0x0061, 0x1F3FF, 0x1F476, 0x200D, 0x1F6D1}, []int{0, 2, 5}, "?? 0061 ?? 1F3FF ?? 1F476 ?? 200D ?? 1F6D1 ??"},
	{[]rune{0x1F476, 0x1F3FF, 0x0308, 0x200D, 0x1F476, 0x1F3FF}, []int{0, 6}, "?? 1F476 ?? 1F3FF ?? 0308 ?? 200D ?? 1F476 ?? 1F3FF ??"},
	{[]rune{0x1F6D1, 0x200D, 0x1F6D1}, []int{0, 3}, "?? 1F6D1 ?? 200D ?? 1F6D1 ??"},
	{[]rune{0x0061, 0x200D, 0x1F6D1}, []int{0, 2, 3}, "?? 0061 ?? 200D ?? 1F6D1 ??"},
	{[]rune{0x2701, 0x200D, 0x2701}, []int{0, 3}, "?? 2701 ?? 200D ?? 2701 ??"},
	{[]rune{0x0061, 0x200D, 0x2701}, []int{0, 2, 3}, "?? 0061 ?? 200D ?? 2701 ??"},
}

func TestAutogenExtGraphemeCluster(t *testing.T) {
	for i := range autogenExtGraphemeCluster {
		te := &autogenExtGraphemeCluster[i]
		dataOffset := 0
		ebbsOffset := 1
		for {
			d1, m1 := useg.NextExtGraphemeClusterBreak(te.data, dataOffset)
			d2, m2 := useg.NextExtGraphemeClusterBreakLen(te.data, dataOffset, len(te.data))
			if d1 != d2 {
				t.Errorf("NextExtGraphemeClusterBreak: \"%s\", offset=%d: conflicting results: %d and %d", te.label, dataOffset, d1, d2)
				break
			}
			if m1 != m2 {
				t.Errorf("NextExtGraphemeClusterBreak: \"%s\", offset=%d: conflicting mandatory results: %v and %v", te.label, dataOffset, m1, m2)
				break
			}
			if d1 == 0 && ebbsOffset < len(te.ebbs) {
				t.Errorf("NextExtGraphemeClusterBreak: \"%s\", offset=%d: = 0, want %d", te.label, dataOffset, te.ebbs[ebbsOffset]-dataOffset)
				break
			}
			if d1 != 0 {
				if ebbsOffset >= len(te.ebbs) {
					t.Errorf("NextExtGraphemeClusterBreak: \"%s\", offset=%d: = %d, want 0", te.label, dataOffset, d1)
					break
				}
				if d1+dataOffset != te.ebbs[ebbsOffset] {
					t.Errorf("NextExtGraphemeClusterBreak: \"%s\", offset=%d: = %d, want %d", te.label, dataOffset, d1, te.ebbs[ebbsOffset]-dataOffset)
					break
				}
				dataOffset += d1
				ebbsOffset++
				break
			}
			break
		}
	}
}
