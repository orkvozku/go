/***************************************************************************
  Copyright 2021-2022 https://github.com user @orkvozku

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
	autogenTestCasesFuncExtpictN    = autogenTestCaseFunc{"HasExtpictN",    useg.HasExtpictN}
	autogenTestCasesFuncExtpictY    = autogenTestCaseFunc{"HasExtpictY",    useg.HasExtpictY}
	autogenTestCasesFuncGcbCn       = autogenTestCaseFunc{"HasGcbCn",       useg.HasGcbCn}
	autogenTestCasesFuncGcbCr       = autogenTestCaseFunc{"HasGcbCr",       useg.HasGcbCr}
	autogenTestCasesFuncGcbEx       = autogenTestCaseFunc{"HasGcbEx",       useg.HasGcbEx}
	autogenTestCasesFuncGcbL        = autogenTestCaseFunc{"HasGcbL",        useg.HasGcbL}
	autogenTestCasesFuncGcbLf       = autogenTestCaseFunc{"HasGcbLf",       useg.HasGcbLf}
	autogenTestCasesFuncGcbLv       = autogenTestCaseFunc{"HasGcbLv",       useg.HasGcbLv}
	autogenTestCasesFuncGcbLvt      = autogenTestCaseFunc{"HasGcbLvt",      useg.HasGcbLvt}
	autogenTestCasesFuncGcbPp       = autogenTestCaseFunc{"HasGcbPp",       useg.HasGcbPp}
	autogenTestCasesFuncGcbRi       = autogenTestCaseFunc{"HasGcbRi",       useg.HasGcbRi}
	autogenTestCasesFuncGcbSm       = autogenTestCaseFunc{"HasGcbSm",       useg.HasGcbSm}
	autogenTestCasesFuncGcbT        = autogenTestCaseFunc{"HasGcbT",        useg.HasGcbT}
	autogenTestCasesFuncGcbV        = autogenTestCaseFunc{"HasGcbV",        useg.HasGcbV}
	autogenTestCasesFuncGcbXx       = autogenTestCaseFunc{"HasGcbXx",       useg.HasGcbXx}
	autogenTestCasesFuncGcbZwj      = autogenTestCaseFunc{"HasGcbZwj",      useg.HasGcbZwj}
	autogenTestCasesFuncLbAi        = autogenTestCaseFunc{"HasLbAi",        useg.HasLbAi}
	autogenTestCasesFuncLbAl        = autogenTestCaseFunc{"HasLbAl",        useg.HasLbAl}
	autogenTestCasesFuncLbB2        = autogenTestCaseFunc{"HasLbB2",        useg.HasLbB2}
	autogenTestCasesFuncLbBa        = autogenTestCaseFunc{"HasLbBa",        useg.HasLbBa}
	autogenTestCasesFuncLbBb        = autogenTestCaseFunc{"HasLbBb",        useg.HasLbBb}
	autogenTestCasesFuncLbBk        = autogenTestCaseFunc{"HasLbBk",        useg.HasLbBk}
	autogenTestCasesFuncLbCb        = autogenTestCaseFunc{"HasLbCb",        useg.HasLbCb}
	autogenTestCasesFuncLbCj        = autogenTestCaseFunc{"HasLbCj",        useg.HasLbCj}
	autogenTestCasesFuncLbCl        = autogenTestCaseFunc{"HasLbCl",        useg.HasLbCl}
	autogenTestCasesFuncLbCm        = autogenTestCaseFunc{"HasLbCm",        useg.HasLbCm}
	autogenTestCasesFuncLbCp        = autogenTestCaseFunc{"HasLbCp",        useg.HasLbCp}
	autogenTestCasesFuncLbCr        = autogenTestCaseFunc{"HasLbCr",        useg.HasLbCr}
	autogenTestCasesFuncLbEb        = autogenTestCaseFunc{"HasLbEb",        useg.HasLbEb}
	autogenTestCasesFuncLbEm        = autogenTestCaseFunc{"HasLbEm",        useg.HasLbEm}
	autogenTestCasesFuncLbEx        = autogenTestCaseFunc{"HasLbEx",        useg.HasLbEx}
	autogenTestCasesFuncLbGl        = autogenTestCaseFunc{"HasLbGl",        useg.HasLbGl}
	autogenTestCasesFuncLbH2        = autogenTestCaseFunc{"HasLbH2",        useg.HasLbH2}
	autogenTestCasesFuncLbH3        = autogenTestCaseFunc{"HasLbH3",        useg.HasLbH3}
	autogenTestCasesFuncLbHl        = autogenTestCaseFunc{"HasLbHl",        useg.HasLbHl}
	autogenTestCasesFuncLbHy        = autogenTestCaseFunc{"HasLbHy",        useg.HasLbHy}
	autogenTestCasesFuncLbId        = autogenTestCaseFunc{"HasLbId",        useg.HasLbId}
	autogenTestCasesFuncLbIn        = autogenTestCaseFunc{"HasLbIn",        useg.HasLbIn}
	autogenTestCasesFuncLbIs        = autogenTestCaseFunc{"HasLbIs",        useg.HasLbIs}
	autogenTestCasesFuncLbJl        = autogenTestCaseFunc{"HasLbJl",        useg.HasLbJl}
	autogenTestCasesFuncLbJt        = autogenTestCaseFunc{"HasLbJt",        useg.HasLbJt}
	autogenTestCasesFuncLbJv        = autogenTestCaseFunc{"HasLbJv",        useg.HasLbJv}
	autogenTestCasesFuncLbLf        = autogenTestCaseFunc{"HasLbLf",        useg.HasLbLf}
	autogenTestCasesFuncLbNl        = autogenTestCaseFunc{"HasLbNl",        useg.HasLbNl}
	autogenTestCasesFuncLbNs        = autogenTestCaseFunc{"HasLbNs",        useg.HasLbNs}
	autogenTestCasesFuncLbNu        = autogenTestCaseFunc{"HasLbNu",        useg.HasLbNu}
	autogenTestCasesFuncLbOp        = autogenTestCaseFunc{"HasLbOp",        useg.HasLbOp}
	autogenTestCasesFuncLbPo        = autogenTestCaseFunc{"HasLbPo",        useg.HasLbPo}
	autogenTestCasesFuncLbPr        = autogenTestCaseFunc{"HasLbPr",        useg.HasLbPr}
	autogenTestCasesFuncLbQu        = autogenTestCaseFunc{"HasLbQu",        useg.HasLbQu}
	autogenTestCasesFuncLbRi        = autogenTestCaseFunc{"HasLbRi",        useg.HasLbRi}
	autogenTestCasesFuncLbSa        = autogenTestCaseFunc{"HasLbSa",        useg.HasLbSa}
	autogenTestCasesFuncLbSg        = autogenTestCaseFunc{"HasLbSg",        useg.HasLbSg}
	autogenTestCasesFuncLbSp        = autogenTestCaseFunc{"HasLbSp",        useg.HasLbSp}
	autogenTestCasesFuncLbSy        = autogenTestCaseFunc{"HasLbSy",        useg.HasLbSy}
	autogenTestCasesFuncLbWj        = autogenTestCaseFunc{"HasLbWj",        useg.HasLbWj}
	autogenTestCasesFuncLbXx        = autogenTestCaseFunc{"HasLbXx",        useg.HasLbXx}
	autogenTestCasesFuncLbZw        = autogenTestCaseFunc{"HasLbZw",        useg.HasLbZw}
	autogenTestCasesFuncLbZwj       = autogenTestCaseFunc{"HasLbZwj",       useg.HasLbZwj}
	autogenTestCasesFuncSbAt        = autogenTestCaseFunc{"HasSbAt",        useg.HasSbAt}
	autogenTestCasesFuncSbCl        = autogenTestCaseFunc{"HasSbCl",        useg.HasSbCl}
	autogenTestCasesFuncSbCr        = autogenTestCaseFunc{"HasSbCr",        useg.HasSbCr}
	autogenTestCasesFuncSbEx        = autogenTestCaseFunc{"HasSbEx",        useg.HasSbEx}
	autogenTestCasesFuncSbFo        = autogenTestCaseFunc{"HasSbFo",        useg.HasSbFo}
	autogenTestCasesFuncSbLe        = autogenTestCaseFunc{"HasSbLe",        useg.HasSbLe}
	autogenTestCasesFuncSbLf        = autogenTestCaseFunc{"HasSbLf",        useg.HasSbLf}
	autogenTestCasesFuncSbLo        = autogenTestCaseFunc{"HasSbLo",        useg.HasSbLo}
	autogenTestCasesFuncSbNu        = autogenTestCaseFunc{"HasSbNu",        useg.HasSbNu}
	autogenTestCasesFuncSbSc        = autogenTestCaseFunc{"HasSbSc",        useg.HasSbSc}
	autogenTestCasesFuncSbSe        = autogenTestCaseFunc{"HasSbSe",        useg.HasSbSe}
	autogenTestCasesFuncSbSp        = autogenTestCaseFunc{"HasSbSp",        useg.HasSbSp}
	autogenTestCasesFuncSbSt        = autogenTestCaseFunc{"HasSbSt",        useg.HasSbSt}
	autogenTestCasesFuncSbUp        = autogenTestCaseFunc{"HasSbUp",        useg.HasSbUp}
	autogenTestCasesFuncSbXx        = autogenTestCaseFunc{"HasSbXx",        useg.HasSbXx}
	autogenTestCasesFuncWbCr        = autogenTestCaseFunc{"HasWbCr",        useg.HasWbCr}
	autogenTestCasesFuncWbDq        = autogenTestCaseFunc{"HasWbDq",        useg.HasWbDq}
	autogenTestCasesFuncWbEx        = autogenTestCaseFunc{"HasWbEx",        useg.HasWbEx}
	autogenTestCasesFuncWbExtend    = autogenTestCaseFunc{"HasWbExtend",    useg.HasWbExtend}
	autogenTestCasesFuncWbFo        = autogenTestCaseFunc{"HasWbFo",        useg.HasWbFo}
	autogenTestCasesFuncWbHl        = autogenTestCaseFunc{"HasWbHl",        useg.HasWbHl}
	autogenTestCasesFuncWbKa        = autogenTestCaseFunc{"HasWbKa",        useg.HasWbKa}
	autogenTestCasesFuncWbLe        = autogenTestCaseFunc{"HasWbLe",        useg.HasWbLe}
	autogenTestCasesFuncWbLf        = autogenTestCaseFunc{"HasWbLf",        useg.HasWbLf}
	autogenTestCasesFuncWbMb        = autogenTestCaseFunc{"HasWbMb",        useg.HasWbMb}
	autogenTestCasesFuncWbMl        = autogenTestCaseFunc{"HasWbMl",        useg.HasWbMl}
	autogenTestCasesFuncWbMn        = autogenTestCaseFunc{"HasWbMn",        useg.HasWbMn}
	autogenTestCasesFuncWbNl        = autogenTestCaseFunc{"HasWbNl",        useg.HasWbNl}
	autogenTestCasesFuncWbNu        = autogenTestCaseFunc{"HasWbNu",        useg.HasWbNu}
	autogenTestCasesFuncWbRi        = autogenTestCaseFunc{"HasWbRi",        useg.HasWbRi}
	autogenTestCasesFuncWbSq        = autogenTestCaseFunc{"HasWbSq",        useg.HasWbSq}
	autogenTestCasesFuncWbWsegspace = autogenTestCaseFunc{"HasWbWsegspace", useg.HasWbWsegspace}
	autogenTestCasesFuncWbXx        = autogenTestCaseFunc{"HasWbXx",        useg.HasWbXx}
	autogenTestCasesFuncWbZwj       = autogenTestCaseFunc{"HasWbZwj",       useg.HasWbZwj}
)

// autogenTestCase describes an entry in autogenTestCases.
type autogenTestCase struct {
	rMin       rune                   // code point range: minimum
	rMax       rune                   // code point range: maximum
	funcsTrue  []*autogenTestCaseFunc // test functions: true
	funcsFalse []*autogenTestCaseFunc // test functions: false
}

// autogenTestCases describes all of our test cases.
var autogenTestCases = [][]autogenTestCase {
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
		{0xDCF, 0xDCF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbEx,
			}, []*autogenTestCaseFunc{}},
		{0xF85, 0xF85, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1360, 0x1711, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbXx,
			}, []*autogenTestCaseFunc{}},
		{0x193C, 0x1A16, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1BA2, 0x1BA5, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbEx,
			}, []*autogenTestCaseFunc{}},
		{0x200E, 0x200F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbCn,
			}, []*autogenTestCaseFunc{}},
		{0xA880, 0xA881, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbSm,
			}, []*autogenTestCaseFunc{}},
		{0xAA4E, 0xAA7B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbXx,
			}, []*autogenTestCaseFunc{}},
		{0xAC71, 0xAC8B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xAE85, 0xAE9F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xB099, 0xB0B3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xB2AD, 0xB2C7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xB4C1, 0xB4DB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xB6D5, 0xB6EF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xB8E9, 0xB903, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xBAFD, 0xBB17, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xBD11, 0xBD2B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xBF25, 0xBF3F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xC139, 0xC153, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xC34D, 0xC367, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xC561, 0xC57B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xC775, 0xC78F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xC989, 0xC9A3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xCB9D, 0xCBB7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xCDB1, 0xCDCB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xCFC5, 0xCFDF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xD1D9, 0xD1F3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xD3ED, 0xD407, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xD601, 0xD61B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbLvt,
			}, []*autogenTestCaseFunc{}},
		{0xFE20, 0xFE2F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbEx,
			}, []*autogenTestCaseFunc{}},
		{0x11002, 0x11002, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbSm,
			}, []*autogenTestCaseFunc{}},
		{0x111C1, 0x111C1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11358, 0x11361, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbXx,
			}, []*autogenTestCaseFunc{}},
		{0x115DE, 0x1162F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11942, 0x11942, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbSm,
			}, []*autogenTestCaseFunc{}},
		{0x11CA8, 0x11CA8, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11F40, 0x11F40, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbEx,
			}, []*autogenTestCaseFunc{}},
		{0x1D17B, 0x1D182, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbEx,
			}, []*autogenTestCaseFunc{}},
		{0x1E4EC, 0x1E4EF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncGcbEx,
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
		{0xCF3, 0xCF3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0xCF4, 0xCFF, []*autogenTestCaseFunc{
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
		{0xE83, 0xE83, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xEA6, 0xEA6, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xECF, 0xECF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xF06, 0xF07, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbBb,
			}, []*autogenTestCaseFunc{}},
		{0xF14, 0xF14, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbEx,
			}, []*autogenTestCaseFunc{}},
		{0xF36, 0xF36, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0xF3E, 0xF3F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0xF85, 0xF85, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbBa,
			}, []*autogenTestCaseFunc{}},
		{0xFC0, 0xFC5, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0xFD4, 0xFD8, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1090, 0x1099, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbNu,
			}, []*autogenTestCaseFunc{}},
		{0x10D0, 0x10FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1250, 0x1256, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x128A, 0x128D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x12C0, 0x12C0, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1312, 0x1315, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x137D, 0x137F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1401, 0x167F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x16EE, 0x16F8, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1737, 0x173F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1772, 0x1773, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x17DA, 0x17DA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbBa,
			}, []*autogenTestCaseFunc{}},
		{0x1800, 0x1801, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x180E, 0x180E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbGl,
			}, []*autogenTestCaseFunc{}},
		{0x1887, 0x18A8, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1920, 0x192B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x1950, 0x196D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbSa,
			}, []*autogenTestCaseFunc{}},
		{0x19D0, 0x19D9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbNu,
			}, []*autogenTestCaseFunc{}},
		{0x1A20, 0x1A5E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbSa,
			}, []*autogenTestCaseFunc{}},
		{0x1A9A, 0x1A9F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1B45, 0x1B4C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1B74, 0x1B7C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1BBA, 0x1BE5, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1C4A, 0x1C4C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1CBB, 0x1CBC, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1CEE, 0x1CF3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1DCD, 0x1DCD, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbGl,
			}, []*autogenTestCaseFunc{}},
		{0x1F20, 0x1F45, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1F5B, 0x1F5B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1FB6, 0x1FC4, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1FF2, 0x1FF4, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x2008, 0x200A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbBa,
			}, []*autogenTestCaseFunc{}},
		{0x2014, 0x2014, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbB2,
			}, []*autogenTestCaseFunc{}},
		{0x2020, 0x2021, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAi,
			}, []*autogenTestCaseFunc{}},
		{0x2038, 0x2038, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x2047, 0x2049, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbNs,
			}, []*autogenTestCaseFunc{}},
		{0x2061, 0x2064, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x207E, 0x207E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCl,
			}, []*autogenTestCaseFunc{}},
		{0x2090, 0x209C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x20BC, 0x20BD, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbPr,
			}, []*autogenTestCaseFunc{}},
		{0x2103, 0x2103, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbPo,
			}, []*autogenTestCaseFunc{}},
		{0x2116, 0x2116, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbPr,
			}, []*autogenTestCaseFunc{}},
		{0x215B, 0x215B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAi,
			}, []*autogenTestCaseFunc{}},
		{0x2189, 0x2189, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAi,
			}, []*autogenTestCaseFunc{}},
		{0x21D5, 0x21FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x220C, 0x220E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x221A, 0x221A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAi,
			}, []*autogenTestCaseFunc{}},
		{0x2227, 0x222C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAi,
			}, []*autogenTestCaseFunc{}},
		{0x2248, 0x2248, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAi,
			}, []*autogenTestCaseFunc{}},
		{0x2264, 0x2267, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAi,
			}, []*autogenTestCaseFunc{}},
		{0x2286, 0x2287, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAi,
			}, []*autogenTestCaseFunc{}},
		{0x22BF, 0x22BF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAi,
			}, []*autogenTestCaseFunc{}},
		{0x230C, 0x2311, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x23F0, 0x23F3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x254C, 0x254F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x25A2, 0x25A2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x25BE, 0x25BF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x25D2, 0x25E1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x2607, 0x2608, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x2619, 0x2619, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x2641, 0x2641, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x2668, 0x2668, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x2680, 0x269D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x26D2, 0x26D2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAi,
			}, []*autogenTestCaseFunc{}},
		{0x26E2, 0x26E2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x26F7, 0x26F8, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x270E, 0x2756, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x2768, 0x2768, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbOp,
			}, []*autogenTestCaseFunc{}},
		{0x2770, 0x2770, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbOp,
			}, []*autogenTestCaseFunc{}},
		{0x27C5, 0x27C5, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbOp,
			}, []*autogenTestCaseFunc{}},
		{0x27EB, 0x27EB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCl,
			}, []*autogenTestCaseFunc{}},
		{0x2985, 0x2985, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbOp,
			}, []*autogenTestCaseFunc{}},
		{0x298D, 0x298D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbOp,
			}, []*autogenTestCaseFunc{}},
		{0x2995, 0x2995, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbOp,
			}, []*autogenTestCaseFunc{}},
		{0x29DB, 0x29DB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCl,
			}, []*autogenTestCaseFunc{}},
		{0x2B76, 0x2B95, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x2CFD, 0x2CFD, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x2D2E, 0x2D2F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2D97, 0x2D9F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2DBF, 0x2DBF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2DDF, 0x2DDF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2E1A, 0x2E1B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x2E26, 0x2E26, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbOp,
			}, []*autogenTestCaseFunc{}},
		{0x2E32, 0x2E32, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x2E43, 0x2E4A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbBa,
			}, []*autogenTestCaseFunc{}},
		{0x2E56, 0x2E56, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCl,
			}, []*autogenTestCaseFunc{}},
		{0x2E5E, 0x2E7F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2FFC, 0x2FFF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x300A, 0x300A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbOp,
			}, []*autogenTestCaseFunc{}},
		{0x3012, 0x3013, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x301B, 0x301B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCl,
			}, []*autogenTestCaseFunc{}},
		{0x3036, 0x303A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x3045, 0x3045, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCj,
			}, []*autogenTestCaseFunc{}},
		{0x3083, 0x3083, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCj,
			}, []*autogenTestCaseFunc{}},
		{0x3095, 0x3096, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCj,
			}, []*autogenTestCaseFunc{}},
		{0x30A3, 0x30A3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCj,
			}, []*autogenTestCaseFunc{}},
		{0x30C3, 0x30C3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCj,
			}, []*autogenTestCaseFunc{}},
		{0x30EE, 0x30EE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCj,
			}, []*autogenTestCaseFunc{}},
		{0x3100, 0x3104, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x3200, 0x321E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0xA016, 0xA48C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0xA60E, 0xA60E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbEx,
			}, []*autogenTestCaseFunc{}},
		{0xA673, 0xA673, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0xA6F8, 0xA6FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xA7DA, 0xA7F1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xA823, 0xA827, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0xA840, 0xA873, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0xA8CE, 0xA8CF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbBa,
			}, []*autogenTestCaseFunc{}},
		{0xA900, 0xA909, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbNu,
			}, []*autogenTestCaseFunc{}},
		{0xA960, 0xA97C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbJl,
			}, []*autogenTestCaseFunc{}},
		{0xA9CE, 0xA9CE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xA9FF, 0xA9FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xAA4E, 0xAA4F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xAAE0, 0xAAEA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0xAB09, 0xAB0E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0xAB30, 0xAB6B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0xABFA, 0xABFF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xAC55, 0xAC6F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xACC5, 0xACDF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xAD35, 0xAD4F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xADA5, 0xADBF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xAE15, 0xAE2F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xAE85, 0xAE9F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xAEF5, 0xAF0F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xAF65, 0xAF7F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xAFD5, 0xAFEF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB045, 0xB05F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB0B5, 0xB0CF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB125, 0xB13F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB195, 0xB1AF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB205, 0xB21F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB275, 0xB28F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB2E5, 0xB2FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB355, 0xB36F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB3C5, 0xB3DF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB435, 0xB44F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB4A5, 0xB4BF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB515, 0xB52F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB585, 0xB59F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB5F5, 0xB60F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB665, 0xB67F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB6D5, 0xB6EF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB745, 0xB75F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB7B5, 0xB7CF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB825, 0xB83F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB895, 0xB8AF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB905, 0xB91F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB975, 0xB98F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xB9E5, 0xB9FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xBA55, 0xBA6F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xBAC5, 0xBADF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xBB35, 0xBB4F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xBBA5, 0xBBBF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xBC15, 0xBC2F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xBC85, 0xBC9F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xBCF5, 0xBD0F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xBD65, 0xBD7F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xBDD5, 0xBDEF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xBE45, 0xBE5F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xBEB5, 0xBECF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xBF25, 0xBF3F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xBF95, 0xBFAF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC005, 0xC01F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC075, 0xC08F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC0E5, 0xC0FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC155, 0xC16F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC1C5, 0xC1DF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC235, 0xC24F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC2A5, 0xC2BF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC315, 0xC32F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC385, 0xC39F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC3F5, 0xC40F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC465, 0xC47F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC4D5, 0xC4EF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC545, 0xC55F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC5B5, 0xC5CF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC625, 0xC63F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC695, 0xC6AF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC705, 0xC71F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC775, 0xC78F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC7E5, 0xC7FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC855, 0xC86F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC8C5, 0xC8DF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC935, 0xC94F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xC9A5, 0xC9BF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCA15, 0xCA2F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCA85, 0xCA9F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCAF5, 0xCB0F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCB65, 0xCB7F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCBD5, 0xCBEF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCC45, 0xCC5F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCCB5, 0xCCCF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCD25, 0xCD3F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCD95, 0xCDAF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCE05, 0xCE1F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCE75, 0xCE8F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCEE5, 0xCEFF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCF55, 0xCF6F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xCFC5, 0xCFDF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD035, 0xD04F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD0A5, 0xD0BF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD115, 0xD12F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD185, 0xD19F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD1F5, 0xD20F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD265, 0xD27F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD2D5, 0xD2EF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD345, 0xD35F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD3B5, 0xD3CF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD425, 0xD43F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD495, 0xD4AF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD505, 0xD51F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD575, 0xD58F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD5E5, 0xD5FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD655, 0xD66F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD6C5, 0xD6DF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD735, 0xD74F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbH3,
			}, []*autogenTestCaseFunc{}},
		{0xD7B0, 0xD7C6, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbJv,
			}, []*autogenTestCaseFunc{}},
		{0xFB07, 0xFB12, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFB37, 0xFB37, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFB45, 0xFB45, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFD90, 0xFD91, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFE00, 0xFE0F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0xFE1A, 0xFE1F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFE3A, 0xFE3A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCl,
			}, []*autogenTestCaseFunc{}},
		{0xFE42, 0xFE42, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCl,
			}, []*autogenTestCaseFunc{}},
		{0xFE51, 0xFE51, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0xFE5B, 0xFE5B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbOp,
			}, []*autogenTestCaseFunc{}},
		{0xFE6A, 0xFE6A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbPo,
			}, []*autogenTestCaseFunc{}},
		{0xFF00, 0xFF00, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFF0A, 0xFF0B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0xFF20, 0xFF3A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0xFF5E, 0xFF5E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0xFF71, 0xFF9D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0xFFD2, 0xFFD7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0xFFE7, 0xFFE7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1000C, 0x1000C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1004E, 0x1004F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10134, 0x10136, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x101FD, 0x101FD, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x102FC, 0x102FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10380, 0x1039D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x103D6, 0x103FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x104FC, 0x104FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1058B, 0x1058B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x105B2, 0x105B2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10756, 0x1075F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x107BB, 0x107FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10839, 0x1083B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x108A7, 0x108AF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1091F, 0x1091F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbBa,
			}, []*autogenTestCaseFunc{}},
		{0x109D0, 0x109D1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10A14, 0x10A14, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10A40, 0x10A48, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x10AE5, 0x10AE6, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x10B39, 0x10B3F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbBa,
			}, []*autogenTestCaseFunc{}},
		{0x10B9D, 0x10BA8, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10CF3, 0x10CF9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10E80, 0x10EA9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x10F00, 0x10F27, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x10F86, 0x10F89, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x11038, 0x11046, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x11073, 0x11074, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x110C2, 0x110C2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x11100, 0x11102, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x11147, 0x11147, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x11180, 0x11182, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x111CD, 0x111CD, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x111E1, 0x111F4, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1123B, 0x1123C, []*autogenTestCaseFunc{
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
		{0x11C00, 0x11C08, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x11C46, 0x11C4F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11C92, 0x11CA7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x11D0B, 0x11D30, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x11D46, 0x11D46, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x11D69, 0x11D69, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11D99, 0x11D9F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11F02, 0x11F02, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x11F43, 0x11F44, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbBa,
			}, []*autogenTestCaseFunc{}},
		{0x11FE1, 0x11FF1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x12475, 0x1247F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1325E, 0x13281, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x13379, 0x13379, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbOp,
			}, []*autogenTestCaseFunc{}},
		{0x1343D, 0x1343D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCl,
			}, []*autogenTestCaseFunc{}},
		{0x145CE, 0x145CE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbOp,
			}, []*autogenTestCaseFunc{}},
		{0x16A60, 0x16A69, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbNu,
			}, []*autogenTestCaseFunc{}},
		{0x16AEE, 0x16AEF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x16B44, 0x16B44, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbBa,
			}, []*autogenTestCaseFunc{}},
		{0x16B78, 0x16B7C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x16F4B, 0x16F4E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x16FE0, 0x16FE3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbNs,
			}, []*autogenTestCaseFunc{}},
		{0x18B00, 0x18CD5, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1AFFD, 0x1AFFE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1B155, 0x1B155, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCj,
			}, []*autogenTestCaseFunc{}},
		{0x1BC70, 0x1BC7C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1BC9F, 0x1BC9F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbBa,
			}, []*autogenTestCaseFunc{}},
		{0x1CFC4, 0x1CFFF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1D16D, 0x1D182, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x1D242, 0x1D244, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x1D357, 0x1D35F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1D4A0, 0x1D4A1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1D4BA, 0x1D4BA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1D50B, 0x1D50C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1D53F, 0x1D53F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1D6A6, 0x1D6A7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1DA6D, 0x1DA74, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1DA9B, 0x1DA9F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x1E000, 0x1E006, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x1E026, 0x1E02A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x1E130, 0x1E136, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x1E2AE, 0x1E2AE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
			}, []*autogenTestCaseFunc{}},
		{0x1E4D0, 0x1E4EB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1E7ED, 0x1E7EE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAl,
			}, []*autogenTestCaseFunc{}},
		{0x1E8D7, 0x1E8FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1E960, 0x1EC70, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1ED3E, 0x1EDFF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1EE25, 0x1EE26, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1EE3A, 0x1EE3A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1EE4A, 0x1EE4A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1EE55, 0x1EE56, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1EE5E, 0x1EE5E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1EE6B, 0x1EE6B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1EE7F, 0x1EE7F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1EEAA, 0x1EEAA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1F110, 0x1F12D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbAi,
			}, []*autogenTestCaseFunc{}},
		{0x1F200, 0x1F384, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x1F3BD, 0x1F3C1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x1F400, 0x1F441, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x1F47D, 0x1F480, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x1F492, 0x1F49F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x1F4AB, 0x1F4AE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x1F525, 0x1F531, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x1F591, 0x1F594, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x1F648, 0x1F64A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x1F6A4, 0x1F6B3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x1F774, 0x1F77F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x1F85A, 0x1F85F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbId,
			}, []*autogenTestCaseFunc{}},
		{0x1F90F, 0x1F90F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbEb,
			}, []*autogenTestCaseFunc{}},
		{0x1F93C, 0x1F93E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbEb,
			}, []*autogenTestCaseFunc{}},
		{0x1F9BB, 0x1F9BB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbEb,
			}, []*autogenTestCaseFunc{}},
		{0x1FAC3, 0x1FAC5, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbEb,
			}, []*autogenTestCaseFunc{}},
		{0x1FBF0, 0x1FBF9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbNu,
			}, []*autogenTestCaseFunc{}},
		{0xE0001, 0xE0001, []*autogenTestCaseFunc{
				&autogenTestCasesFuncLbCm,
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
		{0xD04, 0xD0C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xD45, 0xD45, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xD58, 0xD5E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xD81, 0xD83, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0xDBD, 0xDBD, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xDD6, 0xDD6, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0xE01, 0xE30, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xE50, 0xE59, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbNu,
			}, []*autogenTestCaseFunc{}},
		{0xE8C, 0xEA3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xEBD, 0xEBD, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xED0, 0xED9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbNu,
			}, []*autogenTestCaseFunc{}},
		{0xF20, 0xF29, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbNu,
			}, []*autogenTestCaseFunc{}},
		{0xF3E, 0xF3F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0xF88, 0xF8C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x102B, 0x103E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x105E, 0x1060, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x1082, 0x108D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x10C7, 0x10C7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1249, 0x1249, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x125E, 0x125F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x12B6, 0x12B7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x12D7, 0x12D7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1360, 0x1361, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x13F6, 0x13F7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1681, 0x169A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1712, 0x1715, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x1754, 0x175F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x17B4, 0x17D3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x17EA, 0x1801, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x180E, 0x180E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbFo,
			}, []*autogenTestCaseFunc{}},
		{0x1887, 0x18A8, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1920, 0x192B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x1970, 0x1974, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1A00, 0x1A16, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1A7F, 0x1A7F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x1AB0, 0x1ACE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x1B5A, 0x1B5B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbSt,
			}, []*autogenTestCaseFunc{}},
		{0x1B80, 0x1B82, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x1C00, 0x1C23, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1C50, 0x1C59, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbNu,
			}, []*autogenTestCaseFunc{}},
		{0x1CC0, 0x1CCF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1CF5, 0x1CF6, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1E02, 0x1E02, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E0A, 0x1E0A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E12, 0x1E12, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E1A, 0x1E1A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E22, 0x1E22, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E2A, 0x1E2A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E32, 0x1E32, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E3A, 0x1E3A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E42, 0x1E42, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E4A, 0x1E4A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E52, 0x1E52, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E5A, 0x1E5A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E62, 0x1E62, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E6A, 0x1E6A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E72, 0x1E72, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E7A, 0x1E7A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E82, 0x1E82, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E8A, 0x1E8A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1E92, 0x1E92, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1EA2, 0x1EA2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1EAA, 0x1EAA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1EB2, 0x1EB2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1EBA, 0x1EBA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1EC2, 0x1EC2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1ECA, 0x1ECA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1ED2, 0x1ED2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1EDA, 0x1EDA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1EE2, 0x1EE2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1EEA, 0x1EEA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1EF2, 0x1EF2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1EFA, 0x1EFA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1F16, 0x1F17, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1F46, 0x1F47, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1F5C, 0x1F5C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1F80, 0x1F87, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x1FB6, 0x1FB7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x1FC8, 0x1FCC, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1FE8, 0x1FEC, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x200B, 0x200B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbFo,
			}, []*autogenTestCaseFunc{}},
		{0x2024, 0x2024, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbAt,
			}, []*autogenTestCaseFunc{}},
		{0x203C, 0x203D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbSt,
			}, []*autogenTestCaseFunc{}},
		{0x2066, 0x206F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbFo,
			}, []*autogenTestCaseFunc{}},
		{0x208F, 0x208F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2108, 0x2109, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2116, 0x2118, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2129, 0x2129, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x213A, 0x213B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x214F, 0x215F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2308, 0x230B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbCl,
			}, []*autogenTestCaseFunc{}},
		{0x2761, 0x2767, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2999, 0x29D7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2C61, 0x2C61, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2C6C, 0x2C6C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2C81, 0x2C81, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2C89, 0x2C89, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2C91, 0x2C91, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2C99, 0x2C99, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2CA1, 0x2CA1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2CA9, 0x2CA9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2CB1, 0x2CB1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2CB9, 0x2CB9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2CC1, 0x2CC1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2CC9, 0x2CC9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2CD1, 0x2CD1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2CD9, 0x2CD9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2CE1, 0x2CE1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2CEF, 0x2CF1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x2D2D, 0x2D2D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x2D97, 0x2D9F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2DBF, 0x2DBF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2DDF, 0x2DDF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2E2E, 0x2E2E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbSt,
			}, []*autogenTestCaseFunc{}},
		{0x2E55, 0x2E5C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbCl,
			}, []*autogenTestCaseFunc{}},
		{0x3012, 0x3013, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x3031, 0x3035, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x309D, 0x309F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x3131, 0x318E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x4E00, 0xA48C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xA610, 0xA61F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xA644, 0xA644, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA64C, 0xA64C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA654, 0xA654, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA65C, 0xA65C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA664, 0xA664, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA66C, 0xA66C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA680, 0xA680, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA688, 0xA688, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA690, 0xA690, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA698, 0xA698, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA6F3, 0xA6F3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbSt,
			}, []*autogenTestCaseFunc{}},
		{0xA724, 0xA724, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA72C, 0xA72C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA736, 0xA736, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA73E, 0xA73E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA746, 0xA746, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA74E, 0xA74E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA756, 0xA756, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA75E, 0xA75E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA766, 0xA766, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA76E, 0xA76E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA780, 0xA780, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0xA788, 0xA788, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xA791, 0xA791, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0xA79B, 0xA79B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0xA7A3, 0xA7A3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0xA7AF, 0xA7AF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0xA7BB, 0xA7BB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0xA7C3, 0xA7C3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0xA7D2, 0xA7D2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xA7DA, 0xA7F1, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xA803, 0xA805, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xA82D, 0xA83F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xA8C6, 0xA8CD, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xA8FC, 0xA8FC, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xA930, 0xA946, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xA9C1, 0xA9C7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xA9E6, 0xA9EF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xAA43, 0xAA43, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0xAA77, 0xAA79, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xAAB7, 0xAAB8, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0xAADE, 0xAADF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xAB07, 0xAB08, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xAB2F, 0xAB2F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xABEB, 0xABEB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbSt,
			}, []*autogenTestCaseFunc{}},
		{0xD7C7, 0xD7CA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFB07, 0xFB12, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFB37, 0xFB37, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFB45, 0xFB45, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFD92, 0xFDC7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xFE14, 0xFE16, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFE45, 0xFE46, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFE58, 0xFE58, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbSc,
			}, []*autogenTestCaseFunc{}},
		{0xFEFD, 0xFEFE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFF0E, 0xFF0E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbAt,
			}, []*autogenTestCaseFunc{}},
		{0xFF3B, 0xFF3B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbCl,
			}, []*autogenTestCaseFunc{}},
		{0xFF5E, 0xFF5E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFFA0, 0xFFBE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xFFDA, 0xFFDC, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x10028, 0x1003A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x10080, 0x100FA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x102A0, 0x102D0, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x10350, 0x10375, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x103D0, 0x103D0, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x104AA, 0x104AF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10564, 0x1056F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10596, 0x10596, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x105BD, 0x105FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10781, 0x10782, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x10806, 0x10807, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1083D, 0x1083E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x108F3, 0x108F3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x109B8, 0x109BD, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10A0C, 0x10A0F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x10A3B, 0x10A3E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10A9D, 0x10ABF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10B40, 0x10B55, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x10C80, 0x10CB2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x10D3A, 0x10E7F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10F00, 0x10F1C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x10F5A, 0x10F6F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10FF7, 0x10FFF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11071, 0x11072, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x110BD, 0x110BD, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbFo,
			}, []*autogenTestCaseFunc{}},
		{0x110F0, 0x110F9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbNu,
			}, []*autogenTestCaseFunc{}},
		{0x11141, 0x11143, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbSt,
			}, []*autogenTestCaseFunc{}},
		{0x11176, 0x11176, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x111C9, 0x111CC, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x111DE, 0x111DF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbSt,
			}, []*autogenTestCaseFunc{}},
		{0x1123B, 0x1123C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbSt,
			}, []*autogenTestCaseFunc{}},
		{0x11288, 0x11288, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x112AA, 0x112AF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11305, 0x1130C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11332, 0x11333, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11347, 0x11348, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x1135D, 0x11361, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11435, 0x11446, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x11462, 0x1147F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x114DA, 0x1157F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x115C9, 0x115D7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbSt,
			}, []*autogenTestCaseFunc{}},
		{0x11644, 0x11644, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x116C0, 0x116C9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbNu,
			}, []*autogenTestCaseFunc{}},
		{0x1173C, 0x1173E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbSt,
			}, []*autogenTestCaseFunc{}},
		{0x118C0, 0x118DF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x11914, 0x11914, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1193B, 0x1193E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x11947, 0x1194F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x119DA, 0x119E0, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x11A0B, 0x11A32, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11A48, 0x11A4F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11A9E, 0x11AAF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11C38, 0x11C3F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x11C92, 0x11CA7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x11D0B, 0x11D30, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11D46, 0x11D46, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11D69, 0x11D69, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11D99, 0x11D9F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11F02, 0x11F02, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11F43, 0x11F44, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbSt,
			}, []*autogenTestCaseFunc{}},
		{0x12400, 0x1246E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x13440, 0x13440, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x16A40, 0x16A5E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x16ACA, 0x16ACF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x16B37, 0x16B38, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbSt,
			}, []*autogenTestCaseFunc{}},
		{0x16B78, 0x16B7C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x16F00, 0x16F4A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x16FA0, 0x16FDF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x17000, 0x187F7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1AFF5, 0x1AFFB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1B150, 0x1B152, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1BC00, 0x1BC6A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1BC9D, 0x1BC9E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x1D165, 0x1D169, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x1D1AA, 0x1D1AD, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
			}, []*autogenTestCaseFunc{}},
		{0x1D455, 0x1D455, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1D4A2, 0x1D4A2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1D4BA, 0x1D4BA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1D504, 0x1D505, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1D51E, 0x1D537, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x1D547, 0x1D549, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1D5D4, 0x1D5ED, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1D6A6, 0x1D6A7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1D6FC, 0x1D714, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLo,
			}, []*autogenTestCaseFunc{}},
		{0x1D756, 0x1D76E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbUp,
			}, []*autogenTestCaseFunc{}},
		{0x1D7C3, 0x1D7C3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1DA37, 0x1DA3A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1DA89, 0x1DA9A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1DF1F, 0x1DF24, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1E022, 0x1E022, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1E090, 0x1E0FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1E14E, 0x1E14E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1E2FA, 0x1E4CF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1E7EC, 0x1E7EC, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1E8D7, 0x1E8FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1EE00, 0x1EE03, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE27, 0x1EE27, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE3B, 0x1EE3B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE4B, 0x1EE4B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE57, 0x1EE57, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE5F, 0x1EE5F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE6C, 0x1EE72, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE80, 0x1EE89, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EEAB, 0x1EEBB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1F676, 0x1F678, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbCl,
			}, []*autogenTestCaseFunc{}},
		{0x2B740, 0x2B81D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0x30000, 0x3134A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbLe,
			}, []*autogenTestCaseFunc{}},
		{0xE0100, 0xE01EF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncSbEx,
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
		{0xCF3, 0xCF3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0xCF4, 0xCFF, []*autogenTestCaseFunc{
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
		{0xEC8, 0xECE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0xECF, 0xECF, []*autogenTestCaseFunc{
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
		{0xFC7, 0x102A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1061, 0x1061, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x108E, 0x108E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10C8, 0x10CC, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x124E, 0x124F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1289, 0x1289, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x12BF, 0x12BF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1311, 0x1311, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1390, 0x139F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1680, 0x1680, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbWsegspace,
			}, []*autogenTestCaseFunc{}},
		{0x1712, 0x1715, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x1760, 0x176C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x17DD, 0x17DD, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x181A, 0x181F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x18AB, 0x18AF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x193C, 0x1945, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1A55, 0x1A5E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x1A9A, 0x1AAF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1B50, 0x1B59, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbNu,
			}, []*autogenTestCaseFunc{}},
		{0x1BB0, 0x1BB9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbNu,
			}, []*autogenTestCaseFunc{}},
		{0x1C4A, 0x1C4C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1CBB, 0x1CBC, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1CEE, 0x1CF3, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1E00, 0x1F15, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1F50, 0x1F57, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1F5F, 0x1F7D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1FC2, 0x1FC4, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1FE0, 0x1FEC, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x2008, 0x200A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbWsegspace,
			}, []*autogenTestCaseFunc{}},
		{0x2024, 0x2024, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbMb,
			}, []*autogenTestCaseFunc{}},
		{0x2041, 0x2043, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2066, 0x206F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbFo,
			}, []*autogenTestCaseFunc{}},
		{0x20D0, 0x20F0, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x2115, 0x2115, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x2128, 0x2128, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x2145, 0x2149, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x2C00, 0x2CE4, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x2D27, 0x2D27, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x2D7F, 0x2D7F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x2DB7, 0x2DB7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x2DD7, 0x2DD7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x3001, 0x3004, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x303D, 0x3098, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x3105, 0x312F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x32D0, 0x32FE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbKa,
			}, []*autogenTestCaseFunc{}},
		{0xA500, 0xA60C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0xA673, 0xA673, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0xA708, 0xA7CA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0xA7F2, 0xA801, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0xA828, 0xA82B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0xA8C6, 0xA8CF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0xA8FD, 0xA8FE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0xA954, 0xA95F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0xA9D0, 0xA9D9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbNu,
			}, []*autogenTestCaseFunc{}},
		{0xAA37, 0xAA3F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0xAA7B, 0xAA7D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0xAABE, 0xAABF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0xAAF5, 0xAAF6, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0xAB20, 0xAB26, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0xABEB, 0xABEB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0xD7C7, 0xD7CA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFB1E, 0xFB1E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0xFB3F, 0xFB3F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0xFBD3, 0xFD3D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0xFE00, 0xFE0F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0xFE33, 0xFE34, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbEx,
			}, []*autogenTestCaseFunc{}},
		{0xFE55, 0xFE55, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbMl,
			}, []*autogenTestCaseFunc{}},
		{0xFF07, 0xFF07, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbMb,
			}, []*autogenTestCaseFunc{}},
		{0xFF1B, 0xFF1B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbMn,
			}, []*autogenTestCaseFunc{}},
		{0xFF66, 0xFF9D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbKa,
			}, []*autogenTestCaseFunc{}},
		{0xFFD2, 0xFFD7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1000D, 0x10026, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x10050, 0x1005D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x10280, 0x1029C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1032D, 0x1034A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x103C4, 0x103C7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x104AA, 0x104AF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10564, 0x1056F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10596, 0x10596, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x105BD, 0x105FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10786, 0x10786, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10809, 0x10809, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10856, 0x1085F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x108F6, 0x108FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x109C0, 0x109FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10A14, 0x10A14, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10A40, 0x10A5F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10AE5, 0x10AE6, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x10B80, 0x10B91, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x10D00, 0x10D23, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x10EAD, 0x10EAF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x10F30, 0x10F45, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x10FE0, 0x10FF6, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11071, 0x11072, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x110BD, 0x110BD, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbFo,
			}, []*autogenTestCaseFunc{}},
		{0x110F0, 0x110F9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbNu,
			}, []*autogenTestCaseFunc{}},
		{0x11144, 0x11144, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11177, 0x1117F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x111CE, 0x111CF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x11213, 0x1122B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11287, 0x11287, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x112A9, 0x112AF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11305, 0x1130C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11332, 0x11333, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11347, 0x11348, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x1135D, 0x11361, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11435, 0x11446, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x11480, 0x114AF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11580, 0x115AE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11600, 0x1162F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x116AB, 0x116B7, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x1173A, 0x117FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11907, 0x11908, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11930, 0x11935, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x11942, 0x11943, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x119D8, 0x119D9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11A01, 0x11A0A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x11A50, 0x11A50, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11AF9, 0x11BFF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11C41, 0x11C4F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11CB7, 0x11CFF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11D3A, 0x11D3A, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x11D50, 0x11D59, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbNu,
			}, []*autogenTestCaseFunc{}},
		{0x11D8F, 0x11D8F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x11EE0, 0x11EF2, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11F12, 0x11F33, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x11FB1, 0x11FFF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x12FF1, 0x12FFF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x14647, 0x167FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x16ABF, 0x16ABF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x16B30, 0x16B36, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x16B7D, 0x16B8F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x16F51, 0x16F87, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x16FE4, 0x16FE4, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x1AFFD, 0x1AFFE, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbKa,
			}, []*autogenTestCaseFunc{}},
		{0x1B164, 0x1B167, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbKa,
			}, []*autogenTestCaseFunc{}},
		{0x1BC90, 0x1BC99, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1CF30, 0x1CF46, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x1D185, 0x1D18B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x1D456, 0x1D49C, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1D4A9, 0x1D4AC, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1D4C5, 0x1D505, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1D51E, 0x1D539, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1D54A, 0x1D550, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1D6DC, 0x1D6FA, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1D750, 0x1D76E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1D7C4, 0x1D7CB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1DA75, 0x1DA75, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x1DF00, 0x1DF1E, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1E01B, 0x1E021, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x1E08F, 0x1E08F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x1E14A, 0x1E14D, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbXx,
			}, []*autogenTestCaseFunc{}},
		{0x1E2F0, 0x1E2F9, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbNu,
			}, []*autogenTestCaseFunc{}},
		{0x1E7E8, 0x1E7EB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1E8D0, 0x1E8D6, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
			}, []*autogenTestCaseFunc{}},
		{0x1EE00, 0x1EE03, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE27, 0x1EE27, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE3B, 0x1EE3B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE4B, 0x1EE4B, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE57, 0x1EE57, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE5F, 0x1EE5F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE6C, 0x1EE72, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EE80, 0x1EE89, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1EEAB, 0x1EEBB, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbLe,
			}, []*autogenTestCaseFunc{}},
		{0x1F1E6, 0x1F1FF, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbRi,
			}, []*autogenTestCaseFunc{}},
		{0xE0020, 0xE007F, []*autogenTestCaseFunc{
				&autogenTestCasesFuncWbExtend,
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
	label string  // test case description
}

// autogenExtGraphemeCluster was derived from UCD data at https://www.unicode.org/ucd/.
var autogenExtGraphemeCluster = []autogenBreakTestEntry{
	{[]rune{0x0020,0x0020}, []int{0,1,2}, "÷ 0020 ÷ 0020 ÷"},
	{[]rune{0x0020,0x0308,0x0020}, []int{0,2,3}, "÷ 0020 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0020,0x000D}, []int{0,1,2}, "÷ 0020 ÷ 000D ÷"},
	{[]rune{0x0020,0x0308,0x000D}, []int{0,2,3}, "÷ 0020 × 0308 ÷ 000D ÷"},
	{[]rune{0x0020,0x000A}, []int{0,1,2}, "÷ 0020 ÷ 000A ÷"},
	{[]rune{0x0020,0x0308,0x000A}, []int{0,2,3}, "÷ 0020 × 0308 ÷ 000A ÷"},
	{[]rune{0x0020,0x0001}, []int{0,1,2}, "÷ 0020 ÷ 0001 ÷"},
	{[]rune{0x0020,0x0308,0x0001}, []int{0,2,3}, "÷ 0020 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0020,0x034F}, []int{0,2}, "÷ 0020 × 034F ÷"},
	{[]rune{0x0020,0x0308,0x034F}, []int{0,3}, "÷ 0020 × 0308 × 034F ÷"},
	{[]rune{0x0020,0x1F1E6}, []int{0,1,2}, "÷ 0020 ÷ 1F1E6 ÷"},
	{[]rune{0x0020,0x0308,0x1F1E6}, []int{0,2,3}, "÷ 0020 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0020,0x0600}, []int{0,1,2}, "÷ 0020 ÷ 0600 ÷"},
	{[]rune{0x0020,0x0308,0x0600}, []int{0,2,3}, "÷ 0020 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0020,0x0903}, []int{0,2}, "÷ 0020 × 0903 ÷"},
	{[]rune{0x0020,0x0308,0x0903}, []int{0,3}, "÷ 0020 × 0308 × 0903 ÷"},
	{[]rune{0x0020,0x1100}, []int{0,1,2}, "÷ 0020 ÷ 1100 ÷"},
	{[]rune{0x0020,0x0308,0x1100}, []int{0,2,3}, "÷ 0020 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0020,0x1160}, []int{0,1,2}, "÷ 0020 ÷ 1160 ÷"},
	{[]rune{0x0020,0x0308,0x1160}, []int{0,2,3}, "÷ 0020 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0020,0x11A8}, []int{0,1,2}, "÷ 0020 ÷ 11A8 ÷"},
	{[]rune{0x0020,0x0308,0x11A8}, []int{0,2,3}, "÷ 0020 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0020,0xAC00}, []int{0,1,2}, "÷ 0020 ÷ AC00 ÷"},
	{[]rune{0x0020,0x0308,0xAC00}, []int{0,2,3}, "÷ 0020 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0020,0xAC01}, []int{0,1,2}, "÷ 0020 ÷ AC01 ÷"},
	{[]rune{0x0020,0x0308,0xAC01}, []int{0,2,3}, "÷ 0020 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0020,0x231A}, []int{0,1,2}, "÷ 0020 ÷ 231A ÷"},
	{[]rune{0x0020,0x0308,0x231A}, []int{0,2,3}, "÷ 0020 × 0308 ÷ 231A ÷"},
	{[]rune{0x0020,0x0300}, []int{0,2}, "÷ 0020 × 0300 ÷"},
	{[]rune{0x0020,0x0308,0x0300}, []int{0,3}, "÷ 0020 × 0308 × 0300 ÷"},
	{[]rune{0x0020,0x200D}, []int{0,2}, "÷ 0020 × 200D ÷"},
	{[]rune{0x0020,0x0308,0x200D}, []int{0,3}, "÷ 0020 × 0308 × 200D ÷"},
	{[]rune{0x0020,0x0378}, []int{0,1,2}, "÷ 0020 ÷ 0378 ÷"},
	{[]rune{0x0020,0x0308,0x0378}, []int{0,2,3}, "÷ 0020 × 0308 ÷ 0378 ÷"},
	{[]rune{0x000D,0x0020}, []int{0,1,2}, "÷ 000D ÷ 0020 ÷"},
	{[]rune{0x000D,0x0308,0x0020}, []int{0,1,2,3}, "÷ 000D ÷ 0308 ÷ 0020 ÷"},
	{[]rune{0x000D,0x000D}, []int{0,1,2}, "÷ 000D ÷ 000D ÷"},
	{[]rune{0x000D,0x0308,0x000D}, []int{0,1,2,3}, "÷ 000D ÷ 0308 ÷ 000D ÷"},
	{[]rune{0x000D,0x000A}, []int{0,2}, "÷ 000D × 000A ÷"},
	{[]rune{0x000D,0x0308,0x000A}, []int{0,1,2,3}, "÷ 000D ÷ 0308 ÷ 000A ÷"},
	{[]rune{0x000D,0x0001}, []int{0,1,2}, "÷ 000D ÷ 0001 ÷"},
	{[]rune{0x000D,0x0308,0x0001}, []int{0,1,2,3}, "÷ 000D ÷ 0308 ÷ 0001 ÷"},
	{[]rune{0x000D,0x034F}, []int{0,1,2}, "÷ 000D ÷ 034F ÷"},
	{[]rune{0x000D,0x0308,0x034F}, []int{0,1,3}, "÷ 000D ÷ 0308 × 034F ÷"},
	{[]rune{0x000D,0x1F1E6}, []int{0,1,2}, "÷ 000D ÷ 1F1E6 ÷"},
	{[]rune{0x000D,0x0308,0x1F1E6}, []int{0,1,2,3}, "÷ 000D ÷ 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x000D,0x0600}, []int{0,1,2}, "÷ 000D ÷ 0600 ÷"},
	{[]rune{0x000D,0x0308,0x0600}, []int{0,1,2,3}, "÷ 000D ÷ 0308 ÷ 0600 ÷"},
	{[]rune{0x000D,0x0903}, []int{0,1,2}, "÷ 000D ÷ 0903 ÷"},
	{[]rune{0x000D,0x0308,0x0903}, []int{0,1,3}, "÷ 000D ÷ 0308 × 0903 ÷"},
	{[]rune{0x000D,0x1100}, []int{0,1,2}, "÷ 000D ÷ 1100 ÷"},
	{[]rune{0x000D,0x0308,0x1100}, []int{0,1,2,3}, "÷ 000D ÷ 0308 ÷ 1100 ÷"},
	{[]rune{0x000D,0x1160}, []int{0,1,2}, "÷ 000D ÷ 1160 ÷"},
	{[]rune{0x000D,0x0308,0x1160}, []int{0,1,2,3}, "÷ 000D ÷ 0308 ÷ 1160 ÷"},
	{[]rune{0x000D,0x11A8}, []int{0,1,2}, "÷ 000D ÷ 11A8 ÷"},
	{[]rune{0x000D,0x0308,0x11A8}, []int{0,1,2,3}, "÷ 000D ÷ 0308 ÷ 11A8 ÷"},
	{[]rune{0x000D,0xAC00}, []int{0,1,2}, "÷ 000D ÷ AC00 ÷"},
	{[]rune{0x000D,0x0308,0xAC00}, []int{0,1,2,3}, "÷ 000D ÷ 0308 ÷ AC00 ÷"},
	{[]rune{0x000D,0xAC01}, []int{0,1,2}, "÷ 000D ÷ AC01 ÷"},
	{[]rune{0x000D,0x0308,0xAC01}, []int{0,1,2,3}, "÷ 000D ÷ 0308 ÷ AC01 ÷"},
	{[]rune{0x000D,0x231A}, []int{0,1,2}, "÷ 000D ÷ 231A ÷"},
	{[]rune{0x000D,0x0308,0x231A}, []int{0,1,2,3}, "÷ 000D ÷ 0308 ÷ 231A ÷"},
	{[]rune{0x000D,0x0300}, []int{0,1,2}, "÷ 000D ÷ 0300 ÷"},
	{[]rune{0x000D,0x0308,0x0300}, []int{0,1,3}, "÷ 000D ÷ 0308 × 0300 ÷"},
	{[]rune{0x000D,0x200D}, []int{0,1,2}, "÷ 000D ÷ 200D ÷"},
	{[]rune{0x000D,0x0308,0x200D}, []int{0,1,3}, "÷ 000D ÷ 0308 × 200D ÷"},
	{[]rune{0x000D,0x0378}, []int{0,1,2}, "÷ 000D ÷ 0378 ÷"},
	{[]rune{0x000D,0x0308,0x0378}, []int{0,1,2,3}, "÷ 000D ÷ 0308 ÷ 0378 ÷"},
	{[]rune{0x000A,0x0020}, []int{0,1,2}, "÷ 000A ÷ 0020 ÷"},
	{[]rune{0x000A,0x0308,0x0020}, []int{0,1,2,3}, "÷ 000A ÷ 0308 ÷ 0020 ÷"},
	{[]rune{0x000A,0x000D}, []int{0,1,2}, "÷ 000A ÷ 000D ÷"},
	{[]rune{0x000A,0x0308,0x000D}, []int{0,1,2,3}, "÷ 000A ÷ 0308 ÷ 000D ÷"},
	{[]rune{0x000A,0x000A}, []int{0,1,2}, "÷ 000A ÷ 000A ÷"},
	{[]rune{0x000A,0x0308,0x000A}, []int{0,1,2,3}, "÷ 000A ÷ 0308 ÷ 000A ÷"},
	{[]rune{0x000A,0x0001}, []int{0,1,2}, "÷ 000A ÷ 0001 ÷"},
	{[]rune{0x000A,0x0308,0x0001}, []int{0,1,2,3}, "÷ 000A ÷ 0308 ÷ 0001 ÷"},
	{[]rune{0x000A,0x034F}, []int{0,1,2}, "÷ 000A ÷ 034F ÷"},
	{[]rune{0x000A,0x0308,0x034F}, []int{0,1,3}, "÷ 000A ÷ 0308 × 034F ÷"},
	{[]rune{0x000A,0x1F1E6}, []int{0,1,2}, "÷ 000A ÷ 1F1E6 ÷"},
	{[]rune{0x000A,0x0308,0x1F1E6}, []int{0,1,2,3}, "÷ 000A ÷ 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x000A,0x0600}, []int{0,1,2}, "÷ 000A ÷ 0600 ÷"},
	{[]rune{0x000A,0x0308,0x0600}, []int{0,1,2,3}, "÷ 000A ÷ 0308 ÷ 0600 ÷"},
	{[]rune{0x000A,0x0903}, []int{0,1,2}, "÷ 000A ÷ 0903 ÷"},
	{[]rune{0x000A,0x0308,0x0903}, []int{0,1,3}, "÷ 000A ÷ 0308 × 0903 ÷"},
	{[]rune{0x000A,0x1100}, []int{0,1,2}, "÷ 000A ÷ 1100 ÷"},
	{[]rune{0x000A,0x0308,0x1100}, []int{0,1,2,3}, "÷ 000A ÷ 0308 ÷ 1100 ÷"},
	{[]rune{0x000A,0x1160}, []int{0,1,2}, "÷ 000A ÷ 1160 ÷"},
	{[]rune{0x000A,0x0308,0x1160}, []int{0,1,2,3}, "÷ 000A ÷ 0308 ÷ 1160 ÷"},
	{[]rune{0x000A,0x11A8}, []int{0,1,2}, "÷ 000A ÷ 11A8 ÷"},
	{[]rune{0x000A,0x0308,0x11A8}, []int{0,1,2,3}, "÷ 000A ÷ 0308 ÷ 11A8 ÷"},
	{[]rune{0x000A,0xAC00}, []int{0,1,2}, "÷ 000A ÷ AC00 ÷"},
	{[]rune{0x000A,0x0308,0xAC00}, []int{0,1,2,3}, "÷ 000A ÷ 0308 ÷ AC00 ÷"},
	{[]rune{0x000A,0xAC01}, []int{0,1,2}, "÷ 000A ÷ AC01 ÷"},
	{[]rune{0x000A,0x0308,0xAC01}, []int{0,1,2,3}, "÷ 000A ÷ 0308 ÷ AC01 ÷"},
	{[]rune{0x000A,0x231A}, []int{0,1,2}, "÷ 000A ÷ 231A ÷"},
	{[]rune{0x000A,0x0308,0x231A}, []int{0,1,2,3}, "÷ 000A ÷ 0308 ÷ 231A ÷"},
	{[]rune{0x000A,0x0300}, []int{0,1,2}, "÷ 000A ÷ 0300 ÷"},
	{[]rune{0x000A,0x0308,0x0300}, []int{0,1,3}, "÷ 000A ÷ 0308 × 0300 ÷"},
	{[]rune{0x000A,0x200D}, []int{0,1,2}, "÷ 000A ÷ 200D ÷"},
	{[]rune{0x000A,0x0308,0x200D}, []int{0,1,3}, "÷ 000A ÷ 0308 × 200D ÷"},
	{[]rune{0x000A,0x0378}, []int{0,1,2}, "÷ 000A ÷ 0378 ÷"},
	{[]rune{0x000A,0x0308,0x0378}, []int{0,1,2,3}, "÷ 000A ÷ 0308 ÷ 0378 ÷"},
	{[]rune{0x0001,0x0020}, []int{0,1,2}, "÷ 0001 ÷ 0020 ÷"},
	{[]rune{0x0001,0x0308,0x0020}, []int{0,1,2,3}, "÷ 0001 ÷ 0308 ÷ 0020 ÷"},
	{[]rune{0x0001,0x000D}, []int{0,1,2}, "÷ 0001 ÷ 000D ÷"},
	{[]rune{0x0001,0x0308,0x000D}, []int{0,1,2,3}, "÷ 0001 ÷ 0308 ÷ 000D ÷"},
	{[]rune{0x0001,0x000A}, []int{0,1,2}, "÷ 0001 ÷ 000A ÷"},
	{[]rune{0x0001,0x0308,0x000A}, []int{0,1,2,3}, "÷ 0001 ÷ 0308 ÷ 000A ÷"},
	{[]rune{0x0001,0x0001}, []int{0,1,2}, "÷ 0001 ÷ 0001 ÷"},
	{[]rune{0x0001,0x0308,0x0001}, []int{0,1,2,3}, "÷ 0001 ÷ 0308 ÷ 0001 ÷"},
	{[]rune{0x0001,0x034F}, []int{0,1,2}, "÷ 0001 ÷ 034F ÷"},
	{[]rune{0x0001,0x0308,0x034F}, []int{0,1,3}, "÷ 0001 ÷ 0308 × 034F ÷"},
	{[]rune{0x0001,0x1F1E6}, []int{0,1,2}, "÷ 0001 ÷ 1F1E6 ÷"},
	{[]rune{0x0001,0x0308,0x1F1E6}, []int{0,1,2,3}, "÷ 0001 ÷ 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0001,0x0600}, []int{0,1,2}, "÷ 0001 ÷ 0600 ÷"},
	{[]rune{0x0001,0x0308,0x0600}, []int{0,1,2,3}, "÷ 0001 ÷ 0308 ÷ 0600 ÷"},
	{[]rune{0x0001,0x0903}, []int{0,1,2}, "÷ 0001 ÷ 0903 ÷"},
	{[]rune{0x0001,0x0308,0x0903}, []int{0,1,3}, "÷ 0001 ÷ 0308 × 0903 ÷"},
	{[]rune{0x0001,0x1100}, []int{0,1,2}, "÷ 0001 ÷ 1100 ÷"},
	{[]rune{0x0001,0x0308,0x1100}, []int{0,1,2,3}, "÷ 0001 ÷ 0308 ÷ 1100 ÷"},
	{[]rune{0x0001,0x1160}, []int{0,1,2}, "÷ 0001 ÷ 1160 ÷"},
	{[]rune{0x0001,0x0308,0x1160}, []int{0,1,2,3}, "÷ 0001 ÷ 0308 ÷ 1160 ÷"},
	{[]rune{0x0001,0x11A8}, []int{0,1,2}, "÷ 0001 ÷ 11A8 ÷"},
	{[]rune{0x0001,0x0308,0x11A8}, []int{0,1,2,3}, "÷ 0001 ÷ 0308 ÷ 11A8 ÷"},
	{[]rune{0x0001,0xAC00}, []int{0,1,2}, "÷ 0001 ÷ AC00 ÷"},
	{[]rune{0x0001,0x0308,0xAC00}, []int{0,1,2,3}, "÷ 0001 ÷ 0308 ÷ AC00 ÷"},
	{[]rune{0x0001,0xAC01}, []int{0,1,2}, "÷ 0001 ÷ AC01 ÷"},
	{[]rune{0x0001,0x0308,0xAC01}, []int{0,1,2,3}, "÷ 0001 ÷ 0308 ÷ AC01 ÷"},
	{[]rune{0x0001,0x231A}, []int{0,1,2}, "÷ 0001 ÷ 231A ÷"},
	{[]rune{0x0001,0x0308,0x231A}, []int{0,1,2,3}, "÷ 0001 ÷ 0308 ÷ 231A ÷"},
	{[]rune{0x0001,0x0300}, []int{0,1,2}, "÷ 0001 ÷ 0300 ÷"},
	{[]rune{0x0001,0x0308,0x0300}, []int{0,1,3}, "÷ 0001 ÷ 0308 × 0300 ÷"},
	{[]rune{0x0001,0x200D}, []int{0,1,2}, "÷ 0001 ÷ 200D ÷"},
	{[]rune{0x0001,0x0308,0x200D}, []int{0,1,3}, "÷ 0001 ÷ 0308 × 200D ÷"},
	{[]rune{0x0001,0x0378}, []int{0,1,2}, "÷ 0001 ÷ 0378 ÷"},
	{[]rune{0x0001,0x0308,0x0378}, []int{0,1,2,3}, "÷ 0001 ÷ 0308 ÷ 0378 ÷"},
	{[]rune{0x034F,0x0020}, []int{0,1,2}, "÷ 034F ÷ 0020 ÷"},
	{[]rune{0x034F,0x0308,0x0020}, []int{0,2,3}, "÷ 034F × 0308 ÷ 0020 ÷"},
	{[]rune{0x034F,0x000D}, []int{0,1,2}, "÷ 034F ÷ 000D ÷"},
	{[]rune{0x034F,0x0308,0x000D}, []int{0,2,3}, "÷ 034F × 0308 ÷ 000D ÷"},
	{[]rune{0x034F,0x000A}, []int{0,1,2}, "÷ 034F ÷ 000A ÷"},
	{[]rune{0x034F,0x0308,0x000A}, []int{0,2,3}, "÷ 034F × 0308 ÷ 000A ÷"},
	{[]rune{0x034F,0x0001}, []int{0,1,2}, "÷ 034F ÷ 0001 ÷"},
	{[]rune{0x034F,0x0308,0x0001}, []int{0,2,3}, "÷ 034F × 0308 ÷ 0001 ÷"},
	{[]rune{0x034F,0x034F}, []int{0,2}, "÷ 034F × 034F ÷"},
	{[]rune{0x034F,0x0308,0x034F}, []int{0,3}, "÷ 034F × 0308 × 034F ÷"},
	{[]rune{0x034F,0x1F1E6}, []int{0,1,2}, "÷ 034F ÷ 1F1E6 ÷"},
	{[]rune{0x034F,0x0308,0x1F1E6}, []int{0,2,3}, "÷ 034F × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x034F,0x0600}, []int{0,1,2}, "÷ 034F ÷ 0600 ÷"},
	{[]rune{0x034F,0x0308,0x0600}, []int{0,2,3}, "÷ 034F × 0308 ÷ 0600 ÷"},
	{[]rune{0x034F,0x0903}, []int{0,2}, "÷ 034F × 0903 ÷"},
	{[]rune{0x034F,0x0308,0x0903}, []int{0,3}, "÷ 034F × 0308 × 0903 ÷"},
	{[]rune{0x034F,0x1100}, []int{0,1,2}, "÷ 034F ÷ 1100 ÷"},
	{[]rune{0x034F,0x0308,0x1100}, []int{0,2,3}, "÷ 034F × 0308 ÷ 1100 ÷"},
	{[]rune{0x034F,0x1160}, []int{0,1,2}, "÷ 034F ÷ 1160 ÷"},
	{[]rune{0x034F,0x0308,0x1160}, []int{0,2,3}, "÷ 034F × 0308 ÷ 1160 ÷"},
	{[]rune{0x034F,0x11A8}, []int{0,1,2}, "÷ 034F ÷ 11A8 ÷"},
	{[]rune{0x034F,0x0308,0x11A8}, []int{0,2,3}, "÷ 034F × 0308 ÷ 11A8 ÷"},
	{[]rune{0x034F,0xAC00}, []int{0,1,2}, "÷ 034F ÷ AC00 ÷"},
	{[]rune{0x034F,0x0308,0xAC00}, []int{0,2,3}, "÷ 034F × 0308 ÷ AC00 ÷"},
	{[]rune{0x034F,0xAC01}, []int{0,1,2}, "÷ 034F ÷ AC01 ÷"},
	{[]rune{0x034F,0x0308,0xAC01}, []int{0,2,3}, "÷ 034F × 0308 ÷ AC01 ÷"},
	{[]rune{0x034F,0x231A}, []int{0,1,2}, "÷ 034F ÷ 231A ÷"},
	{[]rune{0x034F,0x0308,0x231A}, []int{0,2,3}, "÷ 034F × 0308 ÷ 231A ÷"},
	{[]rune{0x034F,0x0300}, []int{0,2}, "÷ 034F × 0300 ÷"},
	{[]rune{0x034F,0x0308,0x0300}, []int{0,3}, "÷ 034F × 0308 × 0300 ÷"},
	{[]rune{0x034F,0x200D}, []int{0,2}, "÷ 034F × 200D ÷"},
	{[]rune{0x034F,0x0308,0x200D}, []int{0,3}, "÷ 034F × 0308 × 200D ÷"},
	{[]rune{0x034F,0x0378}, []int{0,1,2}, "÷ 034F ÷ 0378 ÷"},
	{[]rune{0x034F,0x0308,0x0378}, []int{0,2,3}, "÷ 034F × 0308 ÷ 0378 ÷"},
	{[]rune{0x1F1E6,0x0020}, []int{0,1,2}, "÷ 1F1E6 ÷ 0020 ÷"},
	{[]rune{0x1F1E6,0x0308,0x0020}, []int{0,2,3}, "÷ 1F1E6 × 0308 ÷ 0020 ÷"},
	{[]rune{0x1F1E6,0x000D}, []int{0,1,2}, "÷ 1F1E6 ÷ 000D ÷"},
	{[]rune{0x1F1E6,0x0308,0x000D}, []int{0,2,3}, "÷ 1F1E6 × 0308 ÷ 000D ÷"},
	{[]rune{0x1F1E6,0x000A}, []int{0,1,2}, "÷ 1F1E6 ÷ 000A ÷"},
	{[]rune{0x1F1E6,0x0308,0x000A}, []int{0,2,3}, "÷ 1F1E6 × 0308 ÷ 000A ÷"},
	{[]rune{0x1F1E6,0x0001}, []int{0,1,2}, "÷ 1F1E6 ÷ 0001 ÷"},
	{[]rune{0x1F1E6,0x0308,0x0001}, []int{0,2,3}, "÷ 1F1E6 × 0308 ÷ 0001 ÷"},
	{[]rune{0x1F1E6,0x034F}, []int{0,2}, "÷ 1F1E6 × 034F ÷"},
	{[]rune{0x1F1E6,0x0308,0x034F}, []int{0,3}, "÷ 1F1E6 × 0308 × 034F ÷"},
	{[]rune{0x1F1E6,0x1F1E6}, []int{0,2}, "÷ 1F1E6 × 1F1E6 ÷"},
	{[]rune{0x1F1E6,0x0308,0x1F1E6}, []int{0,2,3}, "÷ 1F1E6 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x1F1E6,0x0600}, []int{0,1,2}, "÷ 1F1E6 ÷ 0600 ÷"},
	{[]rune{0x1F1E6,0x0308,0x0600}, []int{0,2,3}, "÷ 1F1E6 × 0308 ÷ 0600 ÷"},
	{[]rune{0x1F1E6,0x0903}, []int{0,2}, "÷ 1F1E6 × 0903 ÷"},
	{[]rune{0x1F1E6,0x0308,0x0903}, []int{0,3}, "÷ 1F1E6 × 0308 × 0903 ÷"},
	{[]rune{0x1F1E6,0x1100}, []int{0,1,2}, "÷ 1F1E6 ÷ 1100 ÷"},
	{[]rune{0x1F1E6,0x0308,0x1100}, []int{0,2,3}, "÷ 1F1E6 × 0308 ÷ 1100 ÷"},
	{[]rune{0x1F1E6,0x1160}, []int{0,1,2}, "÷ 1F1E6 ÷ 1160 ÷"},
	{[]rune{0x1F1E6,0x0308,0x1160}, []int{0,2,3}, "÷ 1F1E6 × 0308 ÷ 1160 ÷"},
	{[]rune{0x1F1E6,0x11A8}, []int{0,1,2}, "÷ 1F1E6 ÷ 11A8 ÷"},
	{[]rune{0x1F1E6,0x0308,0x11A8}, []int{0,2,3}, "÷ 1F1E6 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x1F1E6,0xAC00}, []int{0,1,2}, "÷ 1F1E6 ÷ AC00 ÷"},
	{[]rune{0x1F1E6,0x0308,0xAC00}, []int{0,2,3}, "÷ 1F1E6 × 0308 ÷ AC00 ÷"},
	{[]rune{0x1F1E6,0xAC01}, []int{0,1,2}, "÷ 1F1E6 ÷ AC01 ÷"},
	{[]rune{0x1F1E6,0x0308,0xAC01}, []int{0,2,3}, "÷ 1F1E6 × 0308 ÷ AC01 ÷"},
	{[]rune{0x1F1E6,0x231A}, []int{0,1,2}, "÷ 1F1E6 ÷ 231A ÷"},
	{[]rune{0x1F1E6,0x0308,0x231A}, []int{0,2,3}, "÷ 1F1E6 × 0308 ÷ 231A ÷"},
	{[]rune{0x1F1E6,0x0300}, []int{0,2}, "÷ 1F1E6 × 0300 ÷"},
	{[]rune{0x1F1E6,0x0308,0x0300}, []int{0,3}, "÷ 1F1E6 × 0308 × 0300 ÷"},
	{[]rune{0x1F1E6,0x200D}, []int{0,2}, "÷ 1F1E6 × 200D ÷"},
	{[]rune{0x1F1E6,0x0308,0x200D}, []int{0,3}, "÷ 1F1E6 × 0308 × 200D ÷"},
	{[]rune{0x1F1E6,0x0378}, []int{0,1,2}, "÷ 1F1E6 ÷ 0378 ÷"},
	{[]rune{0x1F1E6,0x0308,0x0378}, []int{0,2,3}, "÷ 1F1E6 × 0308 ÷ 0378 ÷"},
	{[]rune{0x0600,0x0020}, []int{0,2}, "÷ 0600 × 0020 ÷"},
	{[]rune{0x0600,0x0308,0x0020}, []int{0,2,3}, "÷ 0600 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0600,0x000D}, []int{0,1,2}, "÷ 0600 ÷ 000D ÷"},
	{[]rune{0x0600,0x0308,0x000D}, []int{0,2,3}, "÷ 0600 × 0308 ÷ 000D ÷"},
	{[]rune{0x0600,0x000A}, []int{0,1,2}, "÷ 0600 ÷ 000A ÷"},
	{[]rune{0x0600,0x0308,0x000A}, []int{0,2,3}, "÷ 0600 × 0308 ÷ 000A ÷"},
	{[]rune{0x0600,0x0001}, []int{0,1,2}, "÷ 0600 ÷ 0001 ÷"},
	{[]rune{0x0600,0x0308,0x0001}, []int{0,2,3}, "÷ 0600 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0600,0x034F}, []int{0,2}, "÷ 0600 × 034F ÷"},
	{[]rune{0x0600,0x0308,0x034F}, []int{0,3}, "÷ 0600 × 0308 × 034F ÷"},
	{[]rune{0x0600,0x1F1E6}, []int{0,2}, "÷ 0600 × 1F1E6 ÷"},
	{[]rune{0x0600,0x0308,0x1F1E6}, []int{0,2,3}, "÷ 0600 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0600,0x0600}, []int{0,2}, "÷ 0600 × 0600 ÷"},
	{[]rune{0x0600,0x0308,0x0600}, []int{0,2,3}, "÷ 0600 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0600,0x0903}, []int{0,2}, "÷ 0600 × 0903 ÷"},
	{[]rune{0x0600,0x0308,0x0903}, []int{0,3}, "÷ 0600 × 0308 × 0903 ÷"},
	{[]rune{0x0600,0x1100}, []int{0,2}, "÷ 0600 × 1100 ÷"},
	{[]rune{0x0600,0x0308,0x1100}, []int{0,2,3}, "÷ 0600 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0600,0x1160}, []int{0,2}, "÷ 0600 × 1160 ÷"},
	{[]rune{0x0600,0x0308,0x1160}, []int{0,2,3}, "÷ 0600 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0600,0x11A8}, []int{0,2}, "÷ 0600 × 11A8 ÷"},
	{[]rune{0x0600,0x0308,0x11A8}, []int{0,2,3}, "÷ 0600 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0600,0xAC00}, []int{0,2}, "÷ 0600 × AC00 ÷"},
	{[]rune{0x0600,0x0308,0xAC00}, []int{0,2,3}, "÷ 0600 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0600,0xAC01}, []int{0,2}, "÷ 0600 × AC01 ÷"},
	{[]rune{0x0600,0x0308,0xAC01}, []int{0,2,3}, "÷ 0600 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0600,0x231A}, []int{0,2}, "÷ 0600 × 231A ÷"},
	{[]rune{0x0600,0x0308,0x231A}, []int{0,2,3}, "÷ 0600 × 0308 ÷ 231A ÷"},
	{[]rune{0x0600,0x0300}, []int{0,2}, "÷ 0600 × 0300 ÷"},
	{[]rune{0x0600,0x0308,0x0300}, []int{0,3}, "÷ 0600 × 0308 × 0300 ÷"},
	{[]rune{0x0600,0x200D}, []int{0,2}, "÷ 0600 × 200D ÷"},
	{[]rune{0x0600,0x0308,0x200D}, []int{0,3}, "÷ 0600 × 0308 × 200D ÷"},
	{[]rune{0x0600,0x0378}, []int{0,2}, "÷ 0600 × 0378 ÷"},
	{[]rune{0x0600,0x0308,0x0378}, []int{0,2,3}, "÷ 0600 × 0308 ÷ 0378 ÷"},
	{[]rune{0x0903,0x0020}, []int{0,1,2}, "÷ 0903 ÷ 0020 ÷"},
	{[]rune{0x0903,0x0308,0x0020}, []int{0,2,3}, "÷ 0903 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0903,0x000D}, []int{0,1,2}, "÷ 0903 ÷ 000D ÷"},
	{[]rune{0x0903,0x0308,0x000D}, []int{0,2,3}, "÷ 0903 × 0308 ÷ 000D ÷"},
	{[]rune{0x0903,0x000A}, []int{0,1,2}, "÷ 0903 ÷ 000A ÷"},
	{[]rune{0x0903,0x0308,0x000A}, []int{0,2,3}, "÷ 0903 × 0308 ÷ 000A ÷"},
	{[]rune{0x0903,0x0001}, []int{0,1,2}, "÷ 0903 ÷ 0001 ÷"},
	{[]rune{0x0903,0x0308,0x0001}, []int{0,2,3}, "÷ 0903 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0903,0x034F}, []int{0,2}, "÷ 0903 × 034F ÷"},
	{[]rune{0x0903,0x0308,0x034F}, []int{0,3}, "÷ 0903 × 0308 × 034F ÷"},
	{[]rune{0x0903,0x1F1E6}, []int{0,1,2}, "÷ 0903 ÷ 1F1E6 ÷"},
	{[]rune{0x0903,0x0308,0x1F1E6}, []int{0,2,3}, "÷ 0903 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0903,0x0600}, []int{0,1,2}, "÷ 0903 ÷ 0600 ÷"},
	{[]rune{0x0903,0x0308,0x0600}, []int{0,2,3}, "÷ 0903 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0903,0x0903}, []int{0,2}, "÷ 0903 × 0903 ÷"},
	{[]rune{0x0903,0x0308,0x0903}, []int{0,3}, "÷ 0903 × 0308 × 0903 ÷"},
	{[]rune{0x0903,0x1100}, []int{0,1,2}, "÷ 0903 ÷ 1100 ÷"},
	{[]rune{0x0903,0x0308,0x1100}, []int{0,2,3}, "÷ 0903 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0903,0x1160}, []int{0,1,2}, "÷ 0903 ÷ 1160 ÷"},
	{[]rune{0x0903,0x0308,0x1160}, []int{0,2,3}, "÷ 0903 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0903,0x11A8}, []int{0,1,2}, "÷ 0903 ÷ 11A8 ÷"},
	{[]rune{0x0903,0x0308,0x11A8}, []int{0,2,3}, "÷ 0903 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0903,0xAC00}, []int{0,1,2}, "÷ 0903 ÷ AC00 ÷"},
	{[]rune{0x0903,0x0308,0xAC00}, []int{0,2,3}, "÷ 0903 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0903,0xAC01}, []int{0,1,2}, "÷ 0903 ÷ AC01 ÷"},
	{[]rune{0x0903,0x0308,0xAC01}, []int{0,2,3}, "÷ 0903 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0903,0x231A}, []int{0,1,2}, "÷ 0903 ÷ 231A ÷"},
	{[]rune{0x0903,0x0308,0x231A}, []int{0,2,3}, "÷ 0903 × 0308 ÷ 231A ÷"},
	{[]rune{0x0903,0x0300}, []int{0,2}, "÷ 0903 × 0300 ÷"},
	{[]rune{0x0903,0x0308,0x0300}, []int{0,3}, "÷ 0903 × 0308 × 0300 ÷"},
	{[]rune{0x0903,0x200D}, []int{0,2}, "÷ 0903 × 200D ÷"},
	{[]rune{0x0903,0x0308,0x200D}, []int{0,3}, "÷ 0903 × 0308 × 200D ÷"},
	{[]rune{0x0903,0x0378}, []int{0,1,2}, "÷ 0903 ÷ 0378 ÷"},
	{[]rune{0x0903,0x0308,0x0378}, []int{0,2,3}, "÷ 0903 × 0308 ÷ 0378 ÷"},
	{[]rune{0x1100,0x0020}, []int{0,1,2}, "÷ 1100 ÷ 0020 ÷"},
	{[]rune{0x1100,0x0308,0x0020}, []int{0,2,3}, "÷ 1100 × 0308 ÷ 0020 ÷"},
	{[]rune{0x1100,0x000D}, []int{0,1,2}, "÷ 1100 ÷ 000D ÷"},
	{[]rune{0x1100,0x0308,0x000D}, []int{0,2,3}, "÷ 1100 × 0308 ÷ 000D ÷"},
	{[]rune{0x1100,0x000A}, []int{0,1,2}, "÷ 1100 ÷ 000A ÷"},
	{[]rune{0x1100,0x0308,0x000A}, []int{0,2,3}, "÷ 1100 × 0308 ÷ 000A ÷"},
	{[]rune{0x1100,0x0001}, []int{0,1,2}, "÷ 1100 ÷ 0001 ÷"},
	{[]rune{0x1100,0x0308,0x0001}, []int{0,2,3}, "÷ 1100 × 0308 ÷ 0001 ÷"},
	{[]rune{0x1100,0x034F}, []int{0,2}, "÷ 1100 × 034F ÷"},
	{[]rune{0x1100,0x0308,0x034F}, []int{0,3}, "÷ 1100 × 0308 × 034F ÷"},
	{[]rune{0x1100,0x1F1E6}, []int{0,1,2}, "÷ 1100 ÷ 1F1E6 ÷"},
	{[]rune{0x1100,0x0308,0x1F1E6}, []int{0,2,3}, "÷ 1100 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x1100,0x0600}, []int{0,1,2}, "÷ 1100 ÷ 0600 ÷"},
	{[]rune{0x1100,0x0308,0x0600}, []int{0,2,3}, "÷ 1100 × 0308 ÷ 0600 ÷"},
	{[]rune{0x1100,0x0903}, []int{0,2}, "÷ 1100 × 0903 ÷"},
	{[]rune{0x1100,0x0308,0x0903}, []int{0,3}, "÷ 1100 × 0308 × 0903 ÷"},
	{[]rune{0x1100,0x1100}, []int{0,2}, "÷ 1100 × 1100 ÷"},
	{[]rune{0x1100,0x0308,0x1100}, []int{0,2,3}, "÷ 1100 × 0308 ÷ 1100 ÷"},
	{[]rune{0x1100,0x1160}, []int{0,2}, "÷ 1100 × 1160 ÷"},
	{[]rune{0x1100,0x0308,0x1160}, []int{0,2,3}, "÷ 1100 × 0308 ÷ 1160 ÷"},
	{[]rune{0x1100,0x11A8}, []int{0,1,2}, "÷ 1100 ÷ 11A8 ÷"},
	{[]rune{0x1100,0x0308,0x11A8}, []int{0,2,3}, "÷ 1100 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x1100,0xAC00}, []int{0,2}, "÷ 1100 × AC00 ÷"},
	{[]rune{0x1100,0x0308,0xAC00}, []int{0,2,3}, "÷ 1100 × 0308 ÷ AC00 ÷"},
	{[]rune{0x1100,0xAC01}, []int{0,2}, "÷ 1100 × AC01 ÷"},
	{[]rune{0x1100,0x0308,0xAC01}, []int{0,2,3}, "÷ 1100 × 0308 ÷ AC01 ÷"},
	{[]rune{0x1100,0x231A}, []int{0,1,2}, "÷ 1100 ÷ 231A ÷"},
	{[]rune{0x1100,0x0308,0x231A}, []int{0,2,3}, "÷ 1100 × 0308 ÷ 231A ÷"},
	{[]rune{0x1100,0x0300}, []int{0,2}, "÷ 1100 × 0300 ÷"},
	{[]rune{0x1100,0x0308,0x0300}, []int{0,3}, "÷ 1100 × 0308 × 0300 ÷"},
	{[]rune{0x1100,0x200D}, []int{0,2}, "÷ 1100 × 200D ÷"},
	{[]rune{0x1100,0x0308,0x200D}, []int{0,3}, "÷ 1100 × 0308 × 200D ÷"},
	{[]rune{0x1100,0x0378}, []int{0,1,2}, "÷ 1100 ÷ 0378 ÷"},
	{[]rune{0x1100,0x0308,0x0378}, []int{0,2,3}, "÷ 1100 × 0308 ÷ 0378 ÷"},
	{[]rune{0x1160,0x0020}, []int{0,1,2}, "÷ 1160 ÷ 0020 ÷"},
	{[]rune{0x1160,0x0308,0x0020}, []int{0,2,3}, "÷ 1160 × 0308 ÷ 0020 ÷"},
	{[]rune{0x1160,0x000D}, []int{0,1,2}, "÷ 1160 ÷ 000D ÷"},
	{[]rune{0x1160,0x0308,0x000D}, []int{0,2,3}, "÷ 1160 × 0308 ÷ 000D ÷"},
	{[]rune{0x1160,0x000A}, []int{0,1,2}, "÷ 1160 ÷ 000A ÷"},
	{[]rune{0x1160,0x0308,0x000A}, []int{0,2,3}, "÷ 1160 × 0308 ÷ 000A ÷"},
	{[]rune{0x1160,0x0001}, []int{0,1,2}, "÷ 1160 ÷ 0001 ÷"},
	{[]rune{0x1160,0x0308,0x0001}, []int{0,2,3}, "÷ 1160 × 0308 ÷ 0001 ÷"},
	{[]rune{0x1160,0x034F}, []int{0,2}, "÷ 1160 × 034F ÷"},
	{[]rune{0x1160,0x0308,0x034F}, []int{0,3}, "÷ 1160 × 0308 × 034F ÷"},
	{[]rune{0x1160,0x1F1E6}, []int{0,1,2}, "÷ 1160 ÷ 1F1E6 ÷"},
	{[]rune{0x1160,0x0308,0x1F1E6}, []int{0,2,3}, "÷ 1160 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x1160,0x0600}, []int{0,1,2}, "÷ 1160 ÷ 0600 ÷"},
	{[]rune{0x1160,0x0308,0x0600}, []int{0,2,3}, "÷ 1160 × 0308 ÷ 0600 ÷"},
	{[]rune{0x1160,0x0903}, []int{0,2}, "÷ 1160 × 0903 ÷"},
	{[]rune{0x1160,0x0308,0x0903}, []int{0,3}, "÷ 1160 × 0308 × 0903 ÷"},
	{[]rune{0x1160,0x1100}, []int{0,1,2}, "÷ 1160 ÷ 1100 ÷"},
	{[]rune{0x1160,0x0308,0x1100}, []int{0,2,3}, "÷ 1160 × 0308 ÷ 1100 ÷"},
	{[]rune{0x1160,0x1160}, []int{0,2}, "÷ 1160 × 1160 ÷"},
	{[]rune{0x1160,0x0308,0x1160}, []int{0,2,3}, "÷ 1160 × 0308 ÷ 1160 ÷"},
	{[]rune{0x1160,0x11A8}, []int{0,2}, "÷ 1160 × 11A8 ÷"},
	{[]rune{0x1160,0x0308,0x11A8}, []int{0,2,3}, "÷ 1160 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x1160,0xAC00}, []int{0,1,2}, "÷ 1160 ÷ AC00 ÷"},
	{[]rune{0x1160,0x0308,0xAC00}, []int{0,2,3}, "÷ 1160 × 0308 ÷ AC00 ÷"},
	{[]rune{0x1160,0xAC01}, []int{0,1,2}, "÷ 1160 ÷ AC01 ÷"},
	{[]rune{0x1160,0x0308,0xAC01}, []int{0,2,3}, "÷ 1160 × 0308 ÷ AC01 ÷"},
	{[]rune{0x1160,0x231A}, []int{0,1,2}, "÷ 1160 ÷ 231A ÷"},
	{[]rune{0x1160,0x0308,0x231A}, []int{0,2,3}, "÷ 1160 × 0308 ÷ 231A ÷"},
	{[]rune{0x1160,0x0300}, []int{0,2}, "÷ 1160 × 0300 ÷"},
	{[]rune{0x1160,0x0308,0x0300}, []int{0,3}, "÷ 1160 × 0308 × 0300 ÷"},
	{[]rune{0x1160,0x200D}, []int{0,2}, "÷ 1160 × 200D ÷"},
	{[]rune{0x1160,0x0308,0x200D}, []int{0,3}, "÷ 1160 × 0308 × 200D ÷"},
	{[]rune{0x1160,0x0378}, []int{0,1,2}, "÷ 1160 ÷ 0378 ÷"},
	{[]rune{0x1160,0x0308,0x0378}, []int{0,2,3}, "÷ 1160 × 0308 ÷ 0378 ÷"},
	{[]rune{0x11A8,0x0020}, []int{0,1,2}, "÷ 11A8 ÷ 0020 ÷"},
	{[]rune{0x11A8,0x0308,0x0020}, []int{0,2,3}, "÷ 11A8 × 0308 ÷ 0020 ÷"},
	{[]rune{0x11A8,0x000D}, []int{0,1,2}, "÷ 11A8 ÷ 000D ÷"},
	{[]rune{0x11A8,0x0308,0x000D}, []int{0,2,3}, "÷ 11A8 × 0308 ÷ 000D ÷"},
	{[]rune{0x11A8,0x000A}, []int{0,1,2}, "÷ 11A8 ÷ 000A ÷"},
	{[]rune{0x11A8,0x0308,0x000A}, []int{0,2,3}, "÷ 11A8 × 0308 ÷ 000A ÷"},
	{[]rune{0x11A8,0x0001}, []int{0,1,2}, "÷ 11A8 ÷ 0001 ÷"},
	{[]rune{0x11A8,0x0308,0x0001}, []int{0,2,3}, "÷ 11A8 × 0308 ÷ 0001 ÷"},
	{[]rune{0x11A8,0x034F}, []int{0,2}, "÷ 11A8 × 034F ÷"},
	{[]rune{0x11A8,0x0308,0x034F}, []int{0,3}, "÷ 11A8 × 0308 × 034F ÷"},
	{[]rune{0x11A8,0x1F1E6}, []int{0,1,2}, "÷ 11A8 ÷ 1F1E6 ÷"},
	{[]rune{0x11A8,0x0308,0x1F1E6}, []int{0,2,3}, "÷ 11A8 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x11A8,0x0600}, []int{0,1,2}, "÷ 11A8 ÷ 0600 ÷"},
	{[]rune{0x11A8,0x0308,0x0600}, []int{0,2,3}, "÷ 11A8 × 0308 ÷ 0600 ÷"},
	{[]rune{0x11A8,0x0903}, []int{0,2}, "÷ 11A8 × 0903 ÷"},
	{[]rune{0x11A8,0x0308,0x0903}, []int{0,3}, "÷ 11A8 × 0308 × 0903 ÷"},
	{[]rune{0x11A8,0x1100}, []int{0,1,2}, "÷ 11A8 ÷ 1100 ÷"},
	{[]rune{0x11A8,0x0308,0x1100}, []int{0,2,3}, "÷ 11A8 × 0308 ÷ 1100 ÷"},
	{[]rune{0x11A8,0x1160}, []int{0,1,2}, "÷ 11A8 ÷ 1160 ÷"},
	{[]rune{0x11A8,0x0308,0x1160}, []int{0,2,3}, "÷ 11A8 × 0308 ÷ 1160 ÷"},
	{[]rune{0x11A8,0x11A8}, []int{0,2}, "÷ 11A8 × 11A8 ÷"},
	{[]rune{0x11A8,0x0308,0x11A8}, []int{0,2,3}, "÷ 11A8 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x11A8,0xAC00}, []int{0,1,2}, "÷ 11A8 ÷ AC00 ÷"},
	{[]rune{0x11A8,0x0308,0xAC00}, []int{0,2,3}, "÷ 11A8 × 0308 ÷ AC00 ÷"},
	{[]rune{0x11A8,0xAC01}, []int{0,1,2}, "÷ 11A8 ÷ AC01 ÷"},
	{[]rune{0x11A8,0x0308,0xAC01}, []int{0,2,3}, "÷ 11A8 × 0308 ÷ AC01 ÷"},
	{[]rune{0x11A8,0x231A}, []int{0,1,2}, "÷ 11A8 ÷ 231A ÷"},
	{[]rune{0x11A8,0x0308,0x231A}, []int{0,2,3}, "÷ 11A8 × 0308 ÷ 231A ÷"},
	{[]rune{0x11A8,0x0300}, []int{0,2}, "÷ 11A8 × 0300 ÷"},
	{[]rune{0x11A8,0x0308,0x0300}, []int{0,3}, "÷ 11A8 × 0308 × 0300 ÷"},
	{[]rune{0x11A8,0x200D}, []int{0,2}, "÷ 11A8 × 200D ÷"},
	{[]rune{0x11A8,0x0308,0x200D}, []int{0,3}, "÷ 11A8 × 0308 × 200D ÷"},
	{[]rune{0x11A8,0x0378}, []int{0,1,2}, "÷ 11A8 ÷ 0378 ÷"},
	{[]rune{0x11A8,0x0308,0x0378}, []int{0,2,3}, "÷ 11A8 × 0308 ÷ 0378 ÷"},
	{[]rune{0xAC00,0x0020}, []int{0,1,2}, "÷ AC00 ÷ 0020 ÷"},
	{[]rune{0xAC00,0x0308,0x0020}, []int{0,2,3}, "÷ AC00 × 0308 ÷ 0020 ÷"},
	{[]rune{0xAC00,0x000D}, []int{0,1,2}, "÷ AC00 ÷ 000D ÷"},
	{[]rune{0xAC00,0x0308,0x000D}, []int{0,2,3}, "÷ AC00 × 0308 ÷ 000D ÷"},
	{[]rune{0xAC00,0x000A}, []int{0,1,2}, "÷ AC00 ÷ 000A ÷"},
	{[]rune{0xAC00,0x0308,0x000A}, []int{0,2,3}, "÷ AC00 × 0308 ÷ 000A ÷"},
	{[]rune{0xAC00,0x0001}, []int{0,1,2}, "÷ AC00 ÷ 0001 ÷"},
	{[]rune{0xAC00,0x0308,0x0001}, []int{0,2,3}, "÷ AC00 × 0308 ÷ 0001 ÷"},
	{[]rune{0xAC00,0x034F}, []int{0,2}, "÷ AC00 × 034F ÷"},
	{[]rune{0xAC00,0x0308,0x034F}, []int{0,3}, "÷ AC00 × 0308 × 034F ÷"},
	{[]rune{0xAC00,0x1F1E6}, []int{0,1,2}, "÷ AC00 ÷ 1F1E6 ÷"},
	{[]rune{0xAC00,0x0308,0x1F1E6}, []int{0,2,3}, "÷ AC00 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0xAC00,0x0600}, []int{0,1,2}, "÷ AC00 ÷ 0600 ÷"},
	{[]rune{0xAC00,0x0308,0x0600}, []int{0,2,3}, "÷ AC00 × 0308 ÷ 0600 ÷"},
	{[]rune{0xAC00,0x0903}, []int{0,2}, "÷ AC00 × 0903 ÷"},
	{[]rune{0xAC00,0x0308,0x0903}, []int{0,3}, "÷ AC00 × 0308 × 0903 ÷"},
	{[]rune{0xAC00,0x1100}, []int{0,1,2}, "÷ AC00 ÷ 1100 ÷"},
	{[]rune{0xAC00,0x0308,0x1100}, []int{0,2,3}, "÷ AC00 × 0308 ÷ 1100 ÷"},
	{[]rune{0xAC00,0x1160}, []int{0,2}, "÷ AC00 × 1160 ÷"},
	{[]rune{0xAC00,0x0308,0x1160}, []int{0,2,3}, "÷ AC00 × 0308 ÷ 1160 ÷"},
	{[]rune{0xAC00,0x11A8}, []int{0,2}, "÷ AC00 × 11A8 ÷"},
	{[]rune{0xAC00,0x0308,0x11A8}, []int{0,2,3}, "÷ AC00 × 0308 ÷ 11A8 ÷"},
	{[]rune{0xAC00,0xAC00}, []int{0,1,2}, "÷ AC00 ÷ AC00 ÷"},
	{[]rune{0xAC00,0x0308,0xAC00}, []int{0,2,3}, "÷ AC00 × 0308 ÷ AC00 ÷"},
	{[]rune{0xAC00,0xAC01}, []int{0,1,2}, "÷ AC00 ÷ AC01 ÷"},
	{[]rune{0xAC00,0x0308,0xAC01}, []int{0,2,3}, "÷ AC00 × 0308 ÷ AC01 ÷"},
	{[]rune{0xAC00,0x231A}, []int{0,1,2}, "÷ AC00 ÷ 231A ÷"},
	{[]rune{0xAC00,0x0308,0x231A}, []int{0,2,3}, "÷ AC00 × 0308 ÷ 231A ÷"},
	{[]rune{0xAC00,0x0300}, []int{0,2}, "÷ AC00 × 0300 ÷"},
	{[]rune{0xAC00,0x0308,0x0300}, []int{0,3}, "÷ AC00 × 0308 × 0300 ÷"},
	{[]rune{0xAC00,0x200D}, []int{0,2}, "÷ AC00 × 200D ÷"},
	{[]rune{0xAC00,0x0308,0x200D}, []int{0,3}, "÷ AC00 × 0308 × 200D ÷"},
	{[]rune{0xAC00,0x0378}, []int{0,1,2}, "÷ AC00 ÷ 0378 ÷"},
	{[]rune{0xAC00,0x0308,0x0378}, []int{0,2,3}, "÷ AC00 × 0308 ÷ 0378 ÷"},
	{[]rune{0xAC01,0x0020}, []int{0,1,2}, "÷ AC01 ÷ 0020 ÷"},
	{[]rune{0xAC01,0x0308,0x0020}, []int{0,2,3}, "÷ AC01 × 0308 ÷ 0020 ÷"},
	{[]rune{0xAC01,0x000D}, []int{0,1,2}, "÷ AC01 ÷ 000D ÷"},
	{[]rune{0xAC01,0x0308,0x000D}, []int{0,2,3}, "÷ AC01 × 0308 ÷ 000D ÷"},
	{[]rune{0xAC01,0x000A}, []int{0,1,2}, "÷ AC01 ÷ 000A ÷"},
	{[]rune{0xAC01,0x0308,0x000A}, []int{0,2,3}, "÷ AC01 × 0308 ÷ 000A ÷"},
	{[]rune{0xAC01,0x0001}, []int{0,1,2}, "÷ AC01 ÷ 0001 ÷"},
	{[]rune{0xAC01,0x0308,0x0001}, []int{0,2,3}, "÷ AC01 × 0308 ÷ 0001 ÷"},
	{[]rune{0xAC01,0x034F}, []int{0,2}, "÷ AC01 × 034F ÷"},
	{[]rune{0xAC01,0x0308,0x034F}, []int{0,3}, "÷ AC01 × 0308 × 034F ÷"},
	{[]rune{0xAC01,0x1F1E6}, []int{0,1,2}, "÷ AC01 ÷ 1F1E6 ÷"},
	{[]rune{0xAC01,0x0308,0x1F1E6}, []int{0,2,3}, "÷ AC01 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0xAC01,0x0600}, []int{0,1,2}, "÷ AC01 ÷ 0600 ÷"},
	{[]rune{0xAC01,0x0308,0x0600}, []int{0,2,3}, "÷ AC01 × 0308 ÷ 0600 ÷"},
	{[]rune{0xAC01,0x0903}, []int{0,2}, "÷ AC01 × 0903 ÷"},
	{[]rune{0xAC01,0x0308,0x0903}, []int{0,3}, "÷ AC01 × 0308 × 0903 ÷"},
	{[]rune{0xAC01,0x1100}, []int{0,1,2}, "÷ AC01 ÷ 1100 ÷"},
	{[]rune{0xAC01,0x0308,0x1100}, []int{0,2,3}, "÷ AC01 × 0308 ÷ 1100 ÷"},
	{[]rune{0xAC01,0x1160}, []int{0,1,2}, "÷ AC01 ÷ 1160 ÷"},
	{[]rune{0xAC01,0x0308,0x1160}, []int{0,2,3}, "÷ AC01 × 0308 ÷ 1160 ÷"},
	{[]rune{0xAC01,0x11A8}, []int{0,2}, "÷ AC01 × 11A8 ÷"},
	{[]rune{0xAC01,0x0308,0x11A8}, []int{0,2,3}, "÷ AC01 × 0308 ÷ 11A8 ÷"},
	{[]rune{0xAC01,0xAC00}, []int{0,1,2}, "÷ AC01 ÷ AC00 ÷"},
	{[]rune{0xAC01,0x0308,0xAC00}, []int{0,2,3}, "÷ AC01 × 0308 ÷ AC00 ÷"},
	{[]rune{0xAC01,0xAC01}, []int{0,1,2}, "÷ AC01 ÷ AC01 ÷"},
	{[]rune{0xAC01,0x0308,0xAC01}, []int{0,2,3}, "÷ AC01 × 0308 ÷ AC01 ÷"},
	{[]rune{0xAC01,0x231A}, []int{0,1,2}, "÷ AC01 ÷ 231A ÷"},
	{[]rune{0xAC01,0x0308,0x231A}, []int{0,2,3}, "÷ AC01 × 0308 ÷ 231A ÷"},
	{[]rune{0xAC01,0x0300}, []int{0,2}, "÷ AC01 × 0300 ÷"},
	{[]rune{0xAC01,0x0308,0x0300}, []int{0,3}, "÷ AC01 × 0308 × 0300 ÷"},
	{[]rune{0xAC01,0x200D}, []int{0,2}, "÷ AC01 × 200D ÷"},
	{[]rune{0xAC01,0x0308,0x200D}, []int{0,3}, "÷ AC01 × 0308 × 200D ÷"},
	{[]rune{0xAC01,0x0378}, []int{0,1,2}, "÷ AC01 ÷ 0378 ÷"},
	{[]rune{0xAC01,0x0308,0x0378}, []int{0,2,3}, "÷ AC01 × 0308 ÷ 0378 ÷"},
	{[]rune{0x231A,0x0020}, []int{0,1,2}, "÷ 231A ÷ 0020 ÷"},
	{[]rune{0x231A,0x0308,0x0020}, []int{0,2,3}, "÷ 231A × 0308 ÷ 0020 ÷"},
	{[]rune{0x231A,0x000D}, []int{0,1,2}, "÷ 231A ÷ 000D ÷"},
	{[]rune{0x231A,0x0308,0x000D}, []int{0,2,3}, "÷ 231A × 0308 ÷ 000D ÷"},
	{[]rune{0x231A,0x000A}, []int{0,1,2}, "÷ 231A ÷ 000A ÷"},
	{[]rune{0x231A,0x0308,0x000A}, []int{0,2,3}, "÷ 231A × 0308 ÷ 000A ÷"},
	{[]rune{0x231A,0x0001}, []int{0,1,2}, "÷ 231A ÷ 0001 ÷"},
	{[]rune{0x231A,0x0308,0x0001}, []int{0,2,3}, "÷ 231A × 0308 ÷ 0001 ÷"},
	{[]rune{0x231A,0x034F}, []int{0,2}, "÷ 231A × 034F ÷"},
	{[]rune{0x231A,0x0308,0x034F}, []int{0,3}, "÷ 231A × 0308 × 034F ÷"},
	{[]rune{0x231A,0x1F1E6}, []int{0,1,2}, "÷ 231A ÷ 1F1E6 ÷"},
	{[]rune{0x231A,0x0308,0x1F1E6}, []int{0,2,3}, "÷ 231A × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x231A,0x0600}, []int{0,1,2}, "÷ 231A ÷ 0600 ÷"},
	{[]rune{0x231A,0x0308,0x0600}, []int{0,2,3}, "÷ 231A × 0308 ÷ 0600 ÷"},
	{[]rune{0x231A,0x0903}, []int{0,2}, "÷ 231A × 0903 ÷"},
	{[]rune{0x231A,0x0308,0x0903}, []int{0,3}, "÷ 231A × 0308 × 0903 ÷"},
	{[]rune{0x231A,0x1100}, []int{0,1,2}, "÷ 231A ÷ 1100 ÷"},
	{[]rune{0x231A,0x0308,0x1100}, []int{0,2,3}, "÷ 231A × 0308 ÷ 1100 ÷"},
	{[]rune{0x231A,0x1160}, []int{0,1,2}, "÷ 231A ÷ 1160 ÷"},
	{[]rune{0x231A,0x0308,0x1160}, []int{0,2,3}, "÷ 231A × 0308 ÷ 1160 ÷"},
	{[]rune{0x231A,0x11A8}, []int{0,1,2}, "÷ 231A ÷ 11A8 ÷"},
	{[]rune{0x231A,0x0308,0x11A8}, []int{0,2,3}, "÷ 231A × 0308 ÷ 11A8 ÷"},
	{[]rune{0x231A,0xAC00}, []int{0,1,2}, "÷ 231A ÷ AC00 ÷"},
	{[]rune{0x231A,0x0308,0xAC00}, []int{0,2,3}, "÷ 231A × 0308 ÷ AC00 ÷"},
	{[]rune{0x231A,0xAC01}, []int{0,1,2}, "÷ 231A ÷ AC01 ÷"},
	{[]rune{0x231A,0x0308,0xAC01}, []int{0,2,3}, "÷ 231A × 0308 ÷ AC01 ÷"},
	{[]rune{0x231A,0x231A}, []int{0,1,2}, "÷ 231A ÷ 231A ÷"},
	{[]rune{0x231A,0x0308,0x231A}, []int{0,2,3}, "÷ 231A × 0308 ÷ 231A ÷"},
	{[]rune{0x231A,0x0300}, []int{0,2}, "÷ 231A × 0300 ÷"},
	{[]rune{0x231A,0x0308,0x0300}, []int{0,3}, "÷ 231A × 0308 × 0300 ÷"},
	{[]rune{0x231A,0x200D}, []int{0,2}, "÷ 231A × 200D ÷"},
	{[]rune{0x231A,0x0308,0x200D}, []int{0,3}, "÷ 231A × 0308 × 200D ÷"},
	{[]rune{0x231A,0x0378}, []int{0,1,2}, "÷ 231A ÷ 0378 ÷"},
	{[]rune{0x231A,0x0308,0x0378}, []int{0,2,3}, "÷ 231A × 0308 ÷ 0378 ÷"},
	{[]rune{0x0300,0x0020}, []int{0,1,2}, "÷ 0300 ÷ 0020 ÷"},
	{[]rune{0x0300,0x0308,0x0020}, []int{0,2,3}, "÷ 0300 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0300,0x000D}, []int{0,1,2}, "÷ 0300 ÷ 000D ÷"},
	{[]rune{0x0300,0x0308,0x000D}, []int{0,2,3}, "÷ 0300 × 0308 ÷ 000D ÷"},
	{[]rune{0x0300,0x000A}, []int{0,1,2}, "÷ 0300 ÷ 000A ÷"},
	{[]rune{0x0300,0x0308,0x000A}, []int{0,2,3}, "÷ 0300 × 0308 ÷ 000A ÷"},
	{[]rune{0x0300,0x0001}, []int{0,1,2}, "÷ 0300 ÷ 0001 ÷"},
	{[]rune{0x0300,0x0308,0x0001}, []int{0,2,3}, "÷ 0300 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0300,0x034F}, []int{0,2}, "÷ 0300 × 034F ÷"},
	{[]rune{0x0300,0x0308,0x034F}, []int{0,3}, "÷ 0300 × 0308 × 034F ÷"},
	{[]rune{0x0300,0x1F1E6}, []int{0,1,2}, "÷ 0300 ÷ 1F1E6 ÷"},
	{[]rune{0x0300,0x0308,0x1F1E6}, []int{0,2,3}, "÷ 0300 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0300,0x0600}, []int{0,1,2}, "÷ 0300 ÷ 0600 ÷"},
	{[]rune{0x0300,0x0308,0x0600}, []int{0,2,3}, "÷ 0300 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0300,0x0903}, []int{0,2}, "÷ 0300 × 0903 ÷"},
	{[]rune{0x0300,0x0308,0x0903}, []int{0,3}, "÷ 0300 × 0308 × 0903 ÷"},
	{[]rune{0x0300,0x1100}, []int{0,1,2}, "÷ 0300 ÷ 1100 ÷"},
	{[]rune{0x0300,0x0308,0x1100}, []int{0,2,3}, "÷ 0300 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0300,0x1160}, []int{0,1,2}, "÷ 0300 ÷ 1160 ÷"},
	{[]rune{0x0300,0x0308,0x1160}, []int{0,2,3}, "÷ 0300 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0300,0x11A8}, []int{0,1,2}, "÷ 0300 ÷ 11A8 ÷"},
	{[]rune{0x0300,0x0308,0x11A8}, []int{0,2,3}, "÷ 0300 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0300,0xAC00}, []int{0,1,2}, "÷ 0300 ÷ AC00 ÷"},
	{[]rune{0x0300,0x0308,0xAC00}, []int{0,2,3}, "÷ 0300 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0300,0xAC01}, []int{0,1,2}, "÷ 0300 ÷ AC01 ÷"},
	{[]rune{0x0300,0x0308,0xAC01}, []int{0,2,3}, "÷ 0300 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0300,0x231A}, []int{0,1,2}, "÷ 0300 ÷ 231A ÷"},
	{[]rune{0x0300,0x0308,0x231A}, []int{0,2,3}, "÷ 0300 × 0308 ÷ 231A ÷"},
	{[]rune{0x0300,0x0300}, []int{0,2}, "÷ 0300 × 0300 ÷"},
	{[]rune{0x0300,0x0308,0x0300}, []int{0,3}, "÷ 0300 × 0308 × 0300 ÷"},
	{[]rune{0x0300,0x200D}, []int{0,2}, "÷ 0300 × 200D ÷"},
	{[]rune{0x0300,0x0308,0x200D}, []int{0,3}, "÷ 0300 × 0308 × 200D ÷"},
	{[]rune{0x0300,0x0378}, []int{0,1,2}, "÷ 0300 ÷ 0378 ÷"},
	{[]rune{0x0300,0x0308,0x0378}, []int{0,2,3}, "÷ 0300 × 0308 ÷ 0378 ÷"},
	{[]rune{0x200D,0x0020}, []int{0,1,2}, "÷ 200D ÷ 0020 ÷"},
	{[]rune{0x200D,0x0308,0x0020}, []int{0,2,3}, "÷ 200D × 0308 ÷ 0020 ÷"},
	{[]rune{0x200D,0x000D}, []int{0,1,2}, "÷ 200D ÷ 000D ÷"},
	{[]rune{0x200D,0x0308,0x000D}, []int{0,2,3}, "÷ 200D × 0308 ÷ 000D ÷"},
	{[]rune{0x200D,0x000A}, []int{0,1,2}, "÷ 200D ÷ 000A ÷"},
	{[]rune{0x200D,0x0308,0x000A}, []int{0,2,3}, "÷ 200D × 0308 ÷ 000A ÷"},
	{[]rune{0x200D,0x0001}, []int{0,1,2}, "÷ 200D ÷ 0001 ÷"},
	{[]rune{0x200D,0x0308,0x0001}, []int{0,2,3}, "÷ 200D × 0308 ÷ 0001 ÷"},
	{[]rune{0x200D,0x034F}, []int{0,2}, "÷ 200D × 034F ÷"},
	{[]rune{0x200D,0x0308,0x034F}, []int{0,3}, "÷ 200D × 0308 × 034F ÷"},
	{[]rune{0x200D,0x1F1E6}, []int{0,1,2}, "÷ 200D ÷ 1F1E6 ÷"},
	{[]rune{0x200D,0x0308,0x1F1E6}, []int{0,2,3}, "÷ 200D × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x200D,0x0600}, []int{0,1,2}, "÷ 200D ÷ 0600 ÷"},
	{[]rune{0x200D,0x0308,0x0600}, []int{0,2,3}, "÷ 200D × 0308 ÷ 0600 ÷"},
	{[]rune{0x200D,0x0903}, []int{0,2}, "÷ 200D × 0903 ÷"},
	{[]rune{0x200D,0x0308,0x0903}, []int{0,3}, "÷ 200D × 0308 × 0903 ÷"},
	{[]rune{0x200D,0x1100}, []int{0,1,2}, "÷ 200D ÷ 1100 ÷"},
	{[]rune{0x200D,0x0308,0x1100}, []int{0,2,3}, "÷ 200D × 0308 ÷ 1100 ÷"},
	{[]rune{0x200D,0x1160}, []int{0,1,2}, "÷ 200D ÷ 1160 ÷"},
	{[]rune{0x200D,0x0308,0x1160}, []int{0,2,3}, "÷ 200D × 0308 ÷ 1160 ÷"},
	{[]rune{0x200D,0x11A8}, []int{0,1,2}, "÷ 200D ÷ 11A8 ÷"},
	{[]rune{0x200D,0x0308,0x11A8}, []int{0,2,3}, "÷ 200D × 0308 ÷ 11A8 ÷"},
	{[]rune{0x200D,0xAC00}, []int{0,1,2}, "÷ 200D ÷ AC00 ÷"},
	{[]rune{0x200D,0x0308,0xAC00}, []int{0,2,3}, "÷ 200D × 0308 ÷ AC00 ÷"},
	{[]rune{0x200D,0xAC01}, []int{0,1,2}, "÷ 200D ÷ AC01 ÷"},
	{[]rune{0x200D,0x0308,0xAC01}, []int{0,2,3}, "÷ 200D × 0308 ÷ AC01 ÷"},
	{[]rune{0x200D,0x231A}, []int{0,1,2}, "÷ 200D ÷ 231A ÷"},
	{[]rune{0x200D,0x0308,0x231A}, []int{0,2,3}, "÷ 200D × 0308 ÷ 231A ÷"},
	{[]rune{0x200D,0x0300}, []int{0,2}, "÷ 200D × 0300 ÷"},
	{[]rune{0x200D,0x0308,0x0300}, []int{0,3}, "÷ 200D × 0308 × 0300 ÷"},
	{[]rune{0x200D,0x200D}, []int{0,2}, "÷ 200D × 200D ÷"},
	{[]rune{0x200D,0x0308,0x200D}, []int{0,3}, "÷ 200D × 0308 × 200D ÷"},
	{[]rune{0x200D,0x0378}, []int{0,1,2}, "÷ 200D ÷ 0378 ÷"},
	{[]rune{0x200D,0x0308,0x0378}, []int{0,2,3}, "÷ 200D × 0308 ÷ 0378 ÷"},
	{[]rune{0x0378,0x0020}, []int{0,1,2}, "÷ 0378 ÷ 0020 ÷"},
	{[]rune{0x0378,0x0308,0x0020}, []int{0,2,3}, "÷ 0378 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0378,0x000D}, []int{0,1,2}, "÷ 0378 ÷ 000D ÷"},
	{[]rune{0x0378,0x0308,0x000D}, []int{0,2,3}, "÷ 0378 × 0308 ÷ 000D ÷"},
	{[]rune{0x0378,0x000A}, []int{0,1,2}, "÷ 0378 ÷ 000A ÷"},
	{[]rune{0x0378,0x0308,0x000A}, []int{0,2,3}, "÷ 0378 × 0308 ÷ 000A ÷"},
	{[]rune{0x0378,0x0001}, []int{0,1,2}, "÷ 0378 ÷ 0001 ÷"},
	{[]rune{0x0378,0x0308,0x0001}, []int{0,2,3}, "÷ 0378 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0378,0x034F}, []int{0,2}, "÷ 0378 × 034F ÷"},
	{[]rune{0x0378,0x0308,0x034F}, []int{0,3}, "÷ 0378 × 0308 × 034F ÷"},
	{[]rune{0x0378,0x1F1E6}, []int{0,1,2}, "÷ 0378 ÷ 1F1E6 ÷"},
	{[]rune{0x0378,0x0308,0x1F1E6}, []int{0,2,3}, "÷ 0378 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0378,0x0600}, []int{0,1,2}, "÷ 0378 ÷ 0600 ÷"},
	{[]rune{0x0378,0x0308,0x0600}, []int{0,2,3}, "÷ 0378 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0378,0x0903}, []int{0,2}, "÷ 0378 × 0903 ÷"},
	{[]rune{0x0378,0x0308,0x0903}, []int{0,3}, "÷ 0378 × 0308 × 0903 ÷"},
	{[]rune{0x0378,0x1100}, []int{0,1,2}, "÷ 0378 ÷ 1100 ÷"},
	{[]rune{0x0378,0x0308,0x1100}, []int{0,2,3}, "÷ 0378 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0378,0x1160}, []int{0,1,2}, "÷ 0378 ÷ 1160 ÷"},
	{[]rune{0x0378,0x0308,0x1160}, []int{0,2,3}, "÷ 0378 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0378,0x11A8}, []int{0,1,2}, "÷ 0378 ÷ 11A8 ÷"},
	{[]rune{0x0378,0x0308,0x11A8}, []int{0,2,3}, "÷ 0378 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0378,0xAC00}, []int{0,1,2}, "÷ 0378 ÷ AC00 ÷"},
	{[]rune{0x0378,0x0308,0xAC00}, []int{0,2,3}, "÷ 0378 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0378,0xAC01}, []int{0,1,2}, "÷ 0378 ÷ AC01 ÷"},
	{[]rune{0x0378,0x0308,0xAC01}, []int{0,2,3}, "÷ 0378 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0378,0x231A}, []int{0,1,2}, "÷ 0378 ÷ 231A ÷"},
	{[]rune{0x0378,0x0308,0x231A}, []int{0,2,3}, "÷ 0378 × 0308 ÷ 231A ÷"},
	{[]rune{0x0378,0x0300}, []int{0,2}, "÷ 0378 × 0300 ÷"},
	{[]rune{0x0378,0x0308,0x0300}, []int{0,3}, "÷ 0378 × 0308 × 0300 ÷"},
	{[]rune{0x0378,0x200D}, []int{0,2}, "÷ 0378 × 200D ÷"},
	{[]rune{0x0378,0x0308,0x200D}, []int{0,3}, "÷ 0378 × 0308 × 200D ÷"},
	{[]rune{0x0378,0x0378}, []int{0,1,2}, "÷ 0378 ÷ 0378 ÷"},
	{[]rune{0x0378,0x0308,0x0378}, []int{0,2,3}, "÷ 0378 × 0308 ÷ 0378 ÷"},
	{[]rune{0x000D,0x000A,0x0061,0x000A,0x0308}, []int{0,2,3,4,5}, "÷ 000D × 000A ÷ 0061 ÷ 000A ÷ 0308 ÷"},
	{[]rune{0x0061,0x0308}, []int{0,2}, "÷ 0061 × 0308 ÷"},
	{[]rune{0x0020,0x200D,0x0646}, []int{0,2,3}, "÷ 0020 × 200D ÷ 0646 ÷"},
	{[]rune{0x0646,0x200D,0x0020}, []int{0,2,3}, "÷ 0646 × 200D ÷ 0020 ÷"},
	{[]rune{0x1100,0x1100}, []int{0,2}, "÷ 1100 × 1100 ÷"},
	{[]rune{0xAC00,0x11A8,0x1100}, []int{0,2,3}, "÷ AC00 × 11A8 ÷ 1100 ÷"},
	{[]rune{0xAC01,0x11A8,0x1100}, []int{0,2,3}, "÷ AC01 × 11A8 ÷ 1100 ÷"},
	{[]rune{0x1F1E6,0x1F1E7,0x1F1E8,0x0062}, []int{0,2,3,4}, "÷ 1F1E6 × 1F1E7 ÷ 1F1E8 ÷ 0062 ÷"},
	{[]rune{0x0061,0x1F1E6,0x1F1E7,0x1F1E8,0x0062}, []int{0,1,3,4,5}, "÷ 0061 ÷ 1F1E6 × 1F1E7 ÷ 1F1E8 ÷ 0062 ÷"},
	{[]rune{0x0061,0x1F1E6,0x1F1E7,0x200D,0x1F1E8,0x0062}, []int{0,1,4,5,6}, "÷ 0061 ÷ 1F1E6 × 1F1E7 × 200D ÷ 1F1E8 ÷ 0062 ÷"},
	{[]rune{0x0061,0x1F1E6,0x200D,0x1F1E7,0x1F1E8,0x0062}, []int{0,1,3,5,6}, "÷ 0061 ÷ 1F1E6 × 200D ÷ 1F1E7 × 1F1E8 ÷ 0062 ÷"},
	{[]rune{0x0061,0x1F1E6,0x1F1E7,0x1F1E8,0x1F1E9,0x0062}, []int{0,1,3,5,6}, "÷ 0061 ÷ 1F1E6 × 1F1E7 ÷ 1F1E8 × 1F1E9 ÷ 0062 ÷"},
	{[]rune{0x0061,0x200D}, []int{0,2}, "÷ 0061 × 200D ÷"},
	{[]rune{0x0061,0x0308,0x0062}, []int{0,2,3}, "÷ 0061 × 0308 ÷ 0062 ÷"},
	{[]rune{0x0061,0x0903,0x0062}, []int{0,2,3}, "÷ 0061 × 0903 ÷ 0062 ÷"},
	{[]rune{0x0061,0x0600,0x0062}, []int{0,1,3}, "÷ 0061 ÷ 0600 × 0062 ÷"},
	{[]rune{0x1F476,0x1F3FF,0x1F476}, []int{0,2,3}, "÷ 1F476 × 1F3FF ÷ 1F476 ÷"},
	{[]rune{0x0061,0x1F3FF,0x1F476}, []int{0,2,3}, "÷ 0061 × 1F3FF ÷ 1F476 ÷"},
	{[]rune{0x0061,0x1F3FF,0x1F476,0x200D,0x1F6D1}, []int{0,2,5}, "÷ 0061 × 1F3FF ÷ 1F476 × 200D × 1F6D1 ÷"},
	{[]rune{0x1F476,0x1F3FF,0x0308,0x200D,0x1F476,0x1F3FF}, []int{0,6}, "÷ 1F476 × 1F3FF × 0308 × 200D × 1F476 × 1F3FF ÷"},
	{[]rune{0x1F6D1,0x200D,0x1F6D1}, []int{0,3}, "÷ 1F6D1 × 200D × 1F6D1 ÷"},
	{[]rune{0x0061,0x200D,0x1F6D1}, []int{0,2,3}, "÷ 0061 × 200D ÷ 1F6D1 ÷"},
	{[]rune{0x2701,0x200D,0x2701}, []int{0,3}, "÷ 2701 × 200D × 2701 ÷"},
	{[]rune{0x0061,0x200D,0x2701}, []int{0,2,3}, "÷ 0061 × 200D ÷ 2701 ÷"},
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
