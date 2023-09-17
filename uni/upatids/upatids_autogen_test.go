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

package upatids_test

import "testing"
import "github.com/orkvozku/go/uni/upatids"

// autogenTestCaseFunc contains a function to test.
type autogenTestCaseFunc struct {
	label string
	f     func(rune) bool
}

// autogenTestCasesFunc... describes a function to test.
var (
	autogenTestCasesFuncIdcN    = autogenTestCaseFunc{"HasIdcN", upatids.HasIdcN}
	autogenTestCasesFuncIdcY    = autogenTestCaseFunc{"HasIdcY", upatids.HasIdcY}
	autogenTestCasesFuncIdsN    = autogenTestCaseFunc{"HasIdsN", upatids.HasIdsN}
	autogenTestCasesFuncIdsY    = autogenTestCaseFunc{"HasIdsY", upatids.HasIdsY}
	autogenTestCasesFuncPatsynN = autogenTestCaseFunc{"HasPatsynN", upatids.HasPatsynN}
	autogenTestCasesFuncPatsynY = autogenTestCaseFunc{"HasPatsynY", upatids.HasPatsynY}
	autogenTestCasesFuncPatwsN  = autogenTestCaseFunc{"HasPatwsN", upatids.HasPatwsN}
	autogenTestCasesFuncPatwsY  = autogenTestCaseFunc{"HasPatwsY", upatids.HasPatwsY}
	autogenTestCasesFuncXidcN   = autogenTestCaseFunc{"HasXidcN", upatids.HasXidcN}
	autogenTestCasesFuncXidcY   = autogenTestCaseFunc{"HasXidcY", upatids.HasXidcY}
	autogenTestCasesFuncXidsN   = autogenTestCaseFunc{"HasXidsN", upatids.HasXidsN}
	autogenTestCasesFuncXidsY   = autogenTestCaseFunc{"HasXidsY", upatids.HasXidsY}
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
	{ // Idc
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
			&autogenTestCasesFuncIdcY,
		}},
		{0x00, 0x2F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x30, 0x39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x3A, 0x40, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x41, 0x5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x5B, 0x5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x5F, 0x5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x60, 0x60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x61, 0x7A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x7B, 0xA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xAA, 0xAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xAB, 0xB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB5, 0xB5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB6, 0xB6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB7, 0xB7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB8, 0xB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xBA, 0xBA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xBB, 0xBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC0, 0xD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xD7, 0xD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD8, 0xF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xF7, 0xF7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xF8, 0x2C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x2C2, 0x2C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x2C6, 0x2D1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x2D2, 0x2DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x2E0, 0x2E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x2E5, 0x2EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x2EC, 0x2EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x2ED, 0x2ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x2EE, 0x2EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x2EF, 0x2FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x300, 0x374, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x375, 0x375, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x376, 0x377, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x378, 0x379, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x37A, 0x37D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x37E, 0x37E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x37F, 0x37F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x380, 0x385, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x386, 0x38A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x38B, 0x38B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x38C, 0x38C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x38D, 0x38D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x38E, 0x3A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x3A2, 0x3A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x3A3, 0x3F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x3F6, 0x3F6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x3F7, 0x481, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x482, 0x482, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x483, 0x487, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x488, 0x489, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x48A, 0x52F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x530, 0x530, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x531, 0x556, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x557, 0x558, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x559, 0x559, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x55A, 0x55F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x560, 0x588, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x589, 0x590, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x591, 0x5BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x5BE, 0x5BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x5BF, 0x5BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x5C0, 0x5C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x5C1, 0x5C2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x5C3, 0x5C3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x5C4, 0x5C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x5C6, 0x5C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x5C7, 0x5C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x5C8, 0x5CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x5D0, 0x5EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x5EB, 0x5EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x5EF, 0x5F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x5F3, 0x60F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x610, 0x61A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x61B, 0x61F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x620, 0x669, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x66A, 0x66D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x66E, 0x6D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x6D4, 0x6D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x6D5, 0x6DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x6DD, 0x6DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x6DF, 0x6E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x6E9, 0x6E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x6EA, 0x6FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x6FD, 0x6FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x6FF, 0x6FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x700, 0x70F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x710, 0x74A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x74B, 0x74C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x74D, 0x7B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x7B2, 0x7BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x7C0, 0x7F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x7F6, 0x7F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x7FA, 0x7FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x7FB, 0x7FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x7FD, 0x7FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x7FE, 0x7FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x800, 0x82D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x82E, 0x83F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x840, 0x85B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x85C, 0x85F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x860, 0x86A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x86B, 0x86F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x870, 0x887, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x888, 0x888, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x889, 0x88E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x88F, 0x897, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x898, 0x8E1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x8E2, 0x8E2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x8E3, 0x963, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x964, 0x965, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x966, 0x96F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x970, 0x970, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x971, 0x983, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x984, 0x984, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x985, 0x98C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x98D, 0x98E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x98F, 0x990, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x991, 0x992, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x993, 0x9A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x9A9, 0x9A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x9AA, 0x9B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x9B1, 0x9B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x9B2, 0x9B2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x9B3, 0x9B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x9B6, 0x9B9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x9BA, 0x9BB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x9BC, 0x9C4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x9C5, 0x9C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x9C7, 0x9C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x9C9, 0x9CA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x9CB, 0x9CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x9CF, 0x9D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x9D7, 0x9D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x9D8, 0x9DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x9DC, 0x9DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x9DE, 0x9DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x9DF, 0x9E3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x9E4, 0x9E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x9E6, 0x9F1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x9F2, 0x9FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x9FC, 0x9FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x9FD, 0x9FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x9FE, 0x9FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x9FF, 0xA00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA01, 0xA03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA04, 0xA04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA05, 0xA0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA0B, 0xA0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA0F, 0xA10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA11, 0xA12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA13, 0xA28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA29, 0xA29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA2A, 0xA30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA31, 0xA31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA32, 0xA33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA34, 0xA34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA35, 0xA36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA37, 0xA37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA38, 0xA39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA3A, 0xA3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA3C, 0xA3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA3D, 0xA3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA3E, 0xA42, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA43, 0xA46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA47, 0xA48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA49, 0xA4A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA4B, 0xA4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA4E, 0xA50, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA51, 0xA51, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA52, 0xA58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA59, 0xA5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA5D, 0xA5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA5E, 0xA5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA5F, 0xA65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA66, 0xA75, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA76, 0xA80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA81, 0xA83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA84, 0xA84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA85, 0xA8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA8E, 0xA8E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA8F, 0xA91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xA92, 0xA92, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA93, 0xAA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xAA9, 0xAA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xAAA, 0xAB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xAB1, 0xAB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xAB2, 0xAB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xAB4, 0xAB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xAB5, 0xAB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xABA, 0xABB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xABC, 0xAC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xAC6, 0xAC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xAC7, 0xAC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xACA, 0xACA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xACB, 0xACD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xACE, 0xACF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xAD0, 0xAD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xAD1, 0xADF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xAE0, 0xAE3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xAE4, 0xAE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xAE6, 0xAEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xAF0, 0xAF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xAF9, 0xAFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB00, 0xB00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB01, 0xB03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB04, 0xB04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB05, 0xB0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB0D, 0xB0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB0F, 0xB10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB11, 0xB12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB13, 0xB28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB29, 0xB29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB2A, 0xB30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB31, 0xB31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB32, 0xB33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB34, 0xB34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB35, 0xB39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB3A, 0xB3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB3C, 0xB44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB45, 0xB46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB47, 0xB48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB49, 0xB4A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB4B, 0xB4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB4E, 0xB54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB55, 0xB57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB58, 0xB5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB5C, 0xB5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB5E, 0xB5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB5F, 0xB63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB64, 0xB65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB66, 0xB6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB70, 0xB70, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB71, 0xB71, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB72, 0xB81, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB82, 0xB83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB84, 0xB84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB85, 0xB8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB8B, 0xB8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB8E, 0xB90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB91, 0xB91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB92, 0xB95, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB96, 0xB98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB99, 0xB9A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB9B, 0xB9B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB9C, 0xB9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xB9D, 0xB9D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xB9E, 0xB9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xBA0, 0xBA2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xBA3, 0xBA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xBA5, 0xBA7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xBA8, 0xBAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xBAB, 0xBAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xBAE, 0xBB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xBBA, 0xBBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xBBE, 0xBC2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xBC3, 0xBC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xBC6, 0xBC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xBC9, 0xBC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xBCA, 0xBCD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xBCE, 0xBCF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xBD0, 0xBD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xBD1, 0xBD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xBD7, 0xBD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xBD8, 0xBE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xBE6, 0xBEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xBF0, 0xBFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC00, 0xC0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC0D, 0xC0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC0E, 0xC10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC11, 0xC11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC12, 0xC28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC29, 0xC29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC2A, 0xC39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC3A, 0xC3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC3C, 0xC44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC45, 0xC45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC46, 0xC48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC49, 0xC49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC4A, 0xC4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC4E, 0xC54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC55, 0xC56, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC57, 0xC57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC58, 0xC5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC5B, 0xC5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC5D, 0xC5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC5E, 0xC5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC60, 0xC63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC64, 0xC65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC66, 0xC6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC70, 0xC7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC80, 0xC83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC84, 0xC84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC85, 0xC8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC8D, 0xC8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC8E, 0xC90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xC91, 0xC91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xC92, 0xCA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xCA9, 0xCA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xCAA, 0xCB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xCB4, 0xCB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xCB5, 0xCB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xCBA, 0xCBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xCBC, 0xCC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xCC5, 0xCC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xCC6, 0xCC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xCC9, 0xCC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xCCA, 0xCCD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xCCE, 0xCD4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xCD5, 0xCD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xCD7, 0xCDC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xCDD, 0xCDE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xCDF, 0xCDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xCE0, 0xCE3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xCE4, 0xCE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xCE6, 0xCEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xCF0, 0xCF0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xCF1, 0xCF3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xCF4, 0xCFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD00, 0xD0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xD0D, 0xD0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD0E, 0xD10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xD11, 0xD11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD12, 0xD44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xD45, 0xD45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD46, 0xD48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xD49, 0xD49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD4A, 0xD4E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xD4F, 0xD53, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD54, 0xD57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xD58, 0xD5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD5F, 0xD63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xD64, 0xD65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD66, 0xD6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xD70, 0xD79, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD7A, 0xD7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xD80, 0xD80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD81, 0xD83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xD84, 0xD84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD85, 0xD96, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xD97, 0xD99, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD9A, 0xDB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xDB2, 0xDB2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xDB3, 0xDBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xDBC, 0xDBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xDBD, 0xDBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xDBE, 0xDBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xDC0, 0xDC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xDC7, 0xDC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xDCA, 0xDCA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xDCB, 0xDCE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xDCF, 0xDD4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xDD5, 0xDD5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xDD6, 0xDD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xDD7, 0xDD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xDD8, 0xDDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xDE0, 0xDE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xDE6, 0xDEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xDF0, 0xDF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xDF2, 0xDF3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xDF4, 0xE00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xE01, 0xE3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xE3B, 0xE3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xE40, 0xE4E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xE4F, 0xE4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xE50, 0xE59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xE5A, 0xE80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xE81, 0xE82, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xE83, 0xE83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xE84, 0xE84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xE85, 0xE85, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xE86, 0xE8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xE8B, 0xE8B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xE8C, 0xEA3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xEA4, 0xEA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xEA5, 0xEA5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xEA6, 0xEA6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xEA7, 0xEBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xEBE, 0xEBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xEC0, 0xEC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xEC5, 0xEC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xEC6, 0xEC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xEC7, 0xEC7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xEC8, 0xECE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xECF, 0xECF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xED0, 0xED9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xEDA, 0xEDB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xEDC, 0xEDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xEE0, 0xEFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xF00, 0xF00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xF01, 0xF17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xF18, 0xF19, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xF1A, 0xF1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xF20, 0xF29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xF2A, 0xF34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xF35, 0xF35, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xF36, 0xF36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xF37, 0xF37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xF38, 0xF38, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xF39, 0xF39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xF3A, 0xF3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xF3E, 0xF47, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xF48, 0xF48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xF49, 0xF6C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xF6D, 0xF70, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xF71, 0xF84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xF85, 0xF85, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xF86, 0xF97, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xF98, 0xF98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xF99, 0xFBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xFBD, 0xFC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xFC6, 0xFC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0xFC7, 0xFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1000, 0x1049, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x104A, 0x104F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1050, 0x109D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x109E, 0x109F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10A0, 0x10C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x10C6, 0x10C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10C7, 0x10C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x10C8, 0x10CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10CD, 0x10CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x10CE, 0x10CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10D0, 0x10FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x10FB, 0x10FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10FC, 0x1248, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1249, 0x1249, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x124A, 0x124D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x124E, 0x124F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1250, 0x1256, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1257, 0x1257, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1258, 0x1258, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1259, 0x1259, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x125A, 0x125D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x125E, 0x125F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1260, 0x1288, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1289, 0x1289, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x128A, 0x128D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x128E, 0x128F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1290, 0x12B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x12B1, 0x12B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x12B2, 0x12B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x12B6, 0x12B7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x12B8, 0x12BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x12BF, 0x12BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x12C0, 0x12C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x12C1, 0x12C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x12C2, 0x12C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x12C6, 0x12C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x12C8, 0x12D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x12D7, 0x12D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x12D8, 0x1310, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1311, 0x1311, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1312, 0x1315, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1316, 0x1317, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1318, 0x135A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x135B, 0x135C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x135D, 0x135F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1360, 0x1368, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1369, 0x1371, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1372, 0x137F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1380, 0x138F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1390, 0x139F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x13A0, 0x13F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x13F6, 0x13F7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x13F8, 0x13FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x13FE, 0x1400, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1401, 0x166C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x166D, 0x166E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x166F, 0x167F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1680, 0x1680, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1681, 0x169A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x169B, 0x169F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x16A0, 0x16EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x16EB, 0x16ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x16EE, 0x16F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x16F9, 0x16FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1700, 0x1715, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1716, 0x171E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x171F, 0x1734, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1735, 0x173F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1740, 0x1753, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1754, 0x175F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1760, 0x176C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x176D, 0x176D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x176E, 0x1770, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1771, 0x1771, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1772, 0x1773, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x1774, 0x177F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1780, 0x17D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x17D4, 0x17D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x17D7, 0x17D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x17D8, 0x17DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x17DC, 0x17DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x17DE, 0x17DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x17E0, 0x17E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x17EA, 0x180A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x180B, 0x180D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x180E, 0x180E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x180F, 0x1819, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcY,
		}, []*autogenTestCaseFunc{}},
		{0x181A, 0x181F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x191F, 0x191F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1975, 0x197F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1A1C, 0x1A1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1A9A, 0x1AA6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1B4D, 0x1B4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1C38, 0x1C3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1CBB, 0x1CBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1F16, 0x1F17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1F58, 0x1F58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1F7E, 0x1F7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1FC5, 0x1FC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1FED, 0x1FF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x2041, 0x2053, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x209D, 0x20CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x2103, 0x2106, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x211E, 0x2123, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x213A, 0x213B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x2189, 0x2BFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x2D28, 0x2D2C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x2D97, 0x2D9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x2DBF, 0x2DBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x2DDF, 0x2DDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x3036, 0x3037, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x3100, 0x3104, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x3200, 0x33FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA60D, 0xA60F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA6F2, 0xA716, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA7D2, 0xA7D2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA82D, 0xA83F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA8F8, 0xA8FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xA97D, 0xA97F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xAA37, 0xAA3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xAAC3, 0xAADA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xAB07, 0xAB08, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xAB2F, 0xAB2F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xABEE, 0xABEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xD7FC, 0xF8FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xFB18, 0xFB1C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xFB3F, 0xFB3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xFD3E, 0xFD4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xFE10, 0xFE1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xFE75, 0xFE75, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xFF40, 0xFF40, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0xFFD0, 0xFFD1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10027, 0x10027, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1005E, 0x1007F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1029D, 0x1029F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1034B, 0x1034F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x103D0, 0x103D0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x104D4, 0x104D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1057B, 0x1057B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x105A2, 0x105A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10737, 0x1073F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x107B1, 0x107B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10836, 0x10836, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10877, 0x1087F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10916, 0x1091F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10A04, 0x10A04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10A36, 0x10A37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10A9D, 0x10ABF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10B56, 0x10B5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10CB3, 0x10CBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10EAA, 0x10EAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10F28, 0x10F2F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x10FF7, 0x10FFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x110C3, 0x110CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11140, 0x11143, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x111C5, 0x111C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11212, 0x11212, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11289, 0x11289, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x112EB, 0x112EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11311, 0x11312, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1133A, 0x1133A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11351, 0x11356, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11375, 0x113FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x114C6, 0x114C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x115C1, 0x115D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1165A, 0x1167F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1172C, 0x1172F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x118EA, 0x118FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11917, 0x11917, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1195A, 0x1199F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x119E5, 0x119FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11A9E, 0x11AAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11C41, 0x11C4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11CB7, 0x11CFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11D3B, 0x11D3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11D66, 0x11D66, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11D99, 0x11D9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x11F3B, 0x11F3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1239A, 0x123FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x13430, 0x1343F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x16A5F, 0x16A5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x16AEE, 0x16AEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x16B5A, 0x16B62, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x16F4B, 0x16F4E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x16FE5, 0x16FEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x18D09, 0x1AFEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1B123, 0x1B131, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1B168, 0x1B16F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1BC89, 0x1BC8F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1CF47, 0x1D164, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D18C, 0x1D1A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D49D, 0x1D49D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D4AD, 0x1D4AD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D506, 0x1D506, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D53A, 0x1D53A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D551, 0x1D551, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D6FB, 0x1D6FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D76F, 0x1D76F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D7CC, 0x1D7CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1DA76, 0x1DA83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1DF1F, 0x1DF24, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1E022, 0x1E022, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1E090, 0x1E0FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1E14F, 0x1E28F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1E7E7, 0x1E7E7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1E8C5, 0x1E8CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE04, 0x1EE04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE28, 0x1EE28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE3C, 0x1EE41, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE4C, 0x1EE4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE58, 0x1EE58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE60, 0x1EE60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE73, 0x1EE73, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE8A, 0x1EE8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EEBC, 0x1FBEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x2B81E, 0x2B81F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x2FA1E, 0x2FFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdcN,
			&autogenTestCasesFuncIdcY,
		}},
	},
	{ // Ids
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
			&autogenTestCasesFuncIdsY,
		}},
		{0x00, 0x40, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x41, 0x5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x5B, 0x60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x61, 0x7A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x7B, 0xA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xAA, 0xAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xAB, 0xB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB5, 0xB5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB6, 0xB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xBA, 0xBA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xBB, 0xBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xC0, 0xD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xD7, 0xD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xD8, 0xF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xF7, 0xF7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xF8, 0x2C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x2C2, 0x2C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x2C6, 0x2D1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x2D2, 0x2DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x2E0, 0x2E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x2E5, 0x2EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x2EC, 0x2EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x2ED, 0x2ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x2EE, 0x2EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x2EF, 0x36F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x370, 0x374, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x375, 0x375, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x376, 0x377, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x378, 0x379, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x37A, 0x37D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x37E, 0x37E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x37F, 0x37F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x380, 0x385, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x386, 0x386, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x387, 0x387, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x388, 0x38A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x38B, 0x38B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x38C, 0x38C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x38D, 0x38D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x38E, 0x3A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x3A2, 0x3A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x3A3, 0x3F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x3F6, 0x3F6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x3F7, 0x481, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x482, 0x489, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x48A, 0x52F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x530, 0x530, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x531, 0x556, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x557, 0x558, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x559, 0x559, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x55A, 0x55F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x560, 0x588, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x589, 0x5CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x5D0, 0x5EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x5EB, 0x5EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x5EF, 0x5F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x5F3, 0x61F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x620, 0x64A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x64B, 0x66D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x66E, 0x66F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x670, 0x670, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x671, 0x6D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x6D4, 0x6D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x6D5, 0x6D5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x6D6, 0x6E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x6E5, 0x6E6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x6E7, 0x6ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x6EE, 0x6EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x6F0, 0x6F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x6FA, 0x6FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x6FD, 0x6FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x6FF, 0x6FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x700, 0x70F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x710, 0x710, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x711, 0x711, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x712, 0x72F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x730, 0x74C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x74D, 0x7A5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x7A6, 0x7B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x7B1, 0x7B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x7B2, 0x7C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x7CA, 0x7EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x7EB, 0x7F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x7F4, 0x7F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x7F6, 0x7F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x7FA, 0x7FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x7FB, 0x7FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x800, 0x815, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x816, 0x819, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x81A, 0x81A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x81B, 0x823, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x824, 0x824, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x825, 0x827, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x828, 0x828, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x829, 0x83F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x840, 0x858, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x859, 0x85F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x860, 0x86A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x86B, 0x86F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x870, 0x887, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x888, 0x888, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x889, 0x88E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x88F, 0x89F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x8A0, 0x8C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x8CA, 0x903, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x904, 0x939, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x93A, 0x93C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x93D, 0x93D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x93E, 0x94F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x950, 0x950, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x951, 0x957, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x958, 0x961, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x962, 0x970, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x971, 0x980, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x981, 0x984, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x985, 0x98C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x98D, 0x98E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x98F, 0x990, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x991, 0x992, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x993, 0x9A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x9A9, 0x9A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x9AA, 0x9B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x9B1, 0x9B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x9B2, 0x9B2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x9B3, 0x9B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x9B6, 0x9B9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x9BA, 0x9BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x9BD, 0x9BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x9BE, 0x9CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x9CE, 0x9CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x9CF, 0x9DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x9DC, 0x9DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x9DE, 0x9DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x9DF, 0x9E1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x9E2, 0x9EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x9F0, 0x9F1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x9F2, 0x9FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x9FC, 0x9FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x9FD, 0xA04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA05, 0xA0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xA0B, 0xA0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA0F, 0xA10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xA11, 0xA12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA13, 0xA28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xA29, 0xA29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA2A, 0xA30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xA31, 0xA31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA32, 0xA33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xA34, 0xA34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA35, 0xA36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xA37, 0xA37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA38, 0xA39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xA3A, 0xA58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA59, 0xA5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xA5D, 0xA5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA5E, 0xA5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xA5F, 0xA71, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA72, 0xA74, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xA75, 0xA84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA85, 0xA8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xA8E, 0xA8E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA8F, 0xA91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xA92, 0xA92, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA93, 0xAA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xAA9, 0xAA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xAAA, 0xAB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xAB1, 0xAB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xAB2, 0xAB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xAB4, 0xAB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xAB5, 0xAB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xABA, 0xABC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xABD, 0xABD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xABE, 0xACF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xAD0, 0xAD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xAD1, 0xADF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xAE0, 0xAE1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xAE2, 0xAF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xAF9, 0xAF9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xAFA, 0xB04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB05, 0xB0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB0D, 0xB0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB0F, 0xB10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB11, 0xB12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB13, 0xB28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB29, 0xB29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB2A, 0xB30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB31, 0xB31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB32, 0xB33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB34, 0xB34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB35, 0xB39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB3A, 0xB3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB3D, 0xB3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB3E, 0xB5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB5C, 0xB5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB5E, 0xB5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB5F, 0xB61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB62, 0xB70, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB71, 0xB71, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB72, 0xB82, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB83, 0xB83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB84, 0xB84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB85, 0xB8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB8B, 0xB8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB8E, 0xB90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB91, 0xB91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB92, 0xB95, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB96, 0xB98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB99, 0xB9A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB9B, 0xB9B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB9C, 0xB9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xB9D, 0xB9D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xB9E, 0xB9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xBA0, 0xBA2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xBA3, 0xBA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xBA5, 0xBA7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xBA8, 0xBAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xBAB, 0xBAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xBAE, 0xBB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xBBA, 0xBCF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xBD0, 0xBD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xBD1, 0xC04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xC05, 0xC0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xC0D, 0xC0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xC0E, 0xC10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xC11, 0xC11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xC12, 0xC28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xC29, 0xC29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xC2A, 0xC39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xC3A, 0xC3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xC3D, 0xC3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xC3E, 0xC57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xC58, 0xC5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xC5B, 0xC5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xC5D, 0xC5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xC5E, 0xC5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xC60, 0xC61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xC62, 0xC7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xC80, 0xC80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xC81, 0xC84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xC85, 0xC8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xC8D, 0xC8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xC8E, 0xC90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xC91, 0xC91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xC92, 0xCA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xCA9, 0xCA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xCAA, 0xCB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xCB4, 0xCB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xCB5, 0xCB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xCBA, 0xCBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xCBD, 0xCBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xCBE, 0xCDC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xCDD, 0xCDE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xCDF, 0xCDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xCE0, 0xCE1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xCE2, 0xCF0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xCF1, 0xCF2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xCF3, 0xD03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xD04, 0xD0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xD0D, 0xD0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xD0E, 0xD10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xD11, 0xD11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xD12, 0xD3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xD3B, 0xD3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xD3D, 0xD3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xD3E, 0xD4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xD4E, 0xD4E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xD4F, 0xD53, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xD54, 0xD56, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xD57, 0xD5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xD5F, 0xD61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xD62, 0xD79, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xD7A, 0xD7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xD80, 0xD84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xD85, 0xD96, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xD97, 0xD99, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xD9A, 0xDB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xDB2, 0xDB2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xDB3, 0xDBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xDBC, 0xDBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xDBD, 0xDBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xDBE, 0xDBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xDC0, 0xDC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xDC7, 0xE00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xE01, 0xE30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xE31, 0xE31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xE32, 0xE33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xE34, 0xE3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xE40, 0xE46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xE47, 0xE80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xE81, 0xE82, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xE83, 0xE83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xE84, 0xE84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xE85, 0xE85, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xE86, 0xE8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xE8B, 0xE8B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xE8C, 0xEA3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xEA4, 0xEA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xEA5, 0xEA5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xEA6, 0xEA6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xEA7, 0xEB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xEB1, 0xEB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xEB2, 0xEB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xEB4, 0xEBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xEBD, 0xEBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xEBE, 0xEBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xEC0, 0xEC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xEC5, 0xEC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xEC6, 0xEC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xEC7, 0xEDB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xEDC, 0xEDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xEE0, 0xEFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xF00, 0xF00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xF01, 0xF3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xF40, 0xF47, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xF48, 0xF48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xF49, 0xF6C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xF6D, 0xF87, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xF88, 0xF8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0xF8D, 0xFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1000, 0x102A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x102B, 0x103E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x103F, 0x103F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1040, 0x104F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1050, 0x1055, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1056, 0x1059, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x105A, 0x105D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x105E, 0x1060, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1061, 0x1061, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1062, 0x1064, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1065, 0x1066, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1067, 0x106D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x106E, 0x1070, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1071, 0x1074, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1075, 0x1081, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1082, 0x108D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x108E, 0x108E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x108F, 0x109F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10A0, 0x10C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x10C6, 0x10C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10C7, 0x10C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x10C8, 0x10CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10CD, 0x10CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x10CE, 0x10CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10D0, 0x10FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x10FB, 0x10FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10FC, 0x1248, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1249, 0x1249, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x124A, 0x124D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x124E, 0x124F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1250, 0x1256, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1257, 0x1257, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1258, 0x1258, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1259, 0x1259, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x125A, 0x125D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x125E, 0x125F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1260, 0x1288, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1289, 0x1289, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x128A, 0x128D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x128E, 0x128F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1290, 0x12B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x12B1, 0x12B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x12B2, 0x12B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x12B6, 0x12B7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x12B8, 0x12BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x12BF, 0x12BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x12C0, 0x12C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x12C1, 0x12C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x12C2, 0x12C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x12C6, 0x12C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x12C8, 0x12D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x12D7, 0x12D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x12D8, 0x1310, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1311, 0x1311, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1312, 0x1315, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1316, 0x1317, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1318, 0x135A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x135B, 0x137F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1380, 0x138F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1390, 0x139F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x13A0, 0x13F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x13F6, 0x13F7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x13F8, 0x13FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x13FE, 0x1400, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1401, 0x166C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x166D, 0x166E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x166F, 0x167F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1680, 0x1680, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1681, 0x169A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x169B, 0x169F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x16A0, 0x16EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x16EB, 0x16ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x16EE, 0x16F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x16F9, 0x16FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1700, 0x1711, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1712, 0x171E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x171F, 0x1731, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1732, 0x173F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1740, 0x1751, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1752, 0x175F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1760, 0x176C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x176D, 0x176D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x176E, 0x1770, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1771, 0x177F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1780, 0x17B3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x17B4, 0x17D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x17D7, 0x17D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x17D8, 0x17DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x17DC, 0x17DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x17DD, 0x181F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1820, 0x1878, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1879, 0x187F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1880, 0x18A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x18A9, 0x18A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x18AA, 0x18AA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x18AB, 0x18AF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x18B0, 0x18F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x18F6, 0x18FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1900, 0x191E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x191F, 0x194F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1950, 0x196D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x196E, 0x196F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1970, 0x1974, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1975, 0x197F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1980, 0x19AB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x19AC, 0x19AF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x19B0, 0x19C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x19CA, 0x19FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1A00, 0x1A16, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1A17, 0x1A1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1A20, 0x1A54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1A55, 0x1AA6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1AA7, 0x1AA7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1AA8, 0x1B04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1B05, 0x1B33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1B34, 0x1B44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1B45, 0x1B4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1B4D, 0x1B82, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1B83, 0x1BA0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1BA1, 0x1BAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1BAE, 0x1BAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1BB0, 0x1BB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1BBA, 0x1BE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1BE6, 0x1BFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1C00, 0x1C23, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1C24, 0x1C4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1C4D, 0x1C4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1C50, 0x1C59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1C5A, 0x1C7D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1C7E, 0x1C7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1C80, 0x1C88, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1C89, 0x1C8F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1C90, 0x1CBA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1CBB, 0x1CBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1CBD, 0x1CBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1CC0, 0x1CE8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1CE9, 0x1CEC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1CED, 0x1CED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1CEE, 0x1CF3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1CF4, 0x1CF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1CF5, 0x1CF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1CF7, 0x1CF9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1CFA, 0x1CFA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1CFB, 0x1CFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D00, 0x1DBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1DC0, 0x1DFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1E00, 0x1F15, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F16, 0x1F17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F18, 0x1F1D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F1E, 0x1F1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F20, 0x1F45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F46, 0x1F47, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F48, 0x1F4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F4E, 0x1F4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F50, 0x1F57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F58, 0x1F58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F59, 0x1F59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5A, 0x1F5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5B, 0x1F5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5C, 0x1F5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5D, 0x1F5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5E, 0x1F5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5F, 0x1F7D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F7E, 0x1F7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F80, 0x1FB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FB5, 0x1FB5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FB6, 0x1FBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FBD, 0x1FBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FBE, 0x1FBE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FBF, 0x1FC1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FC2, 0x1FC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FC5, 0x1FC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FC6, 0x1FCC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FCD, 0x1FCF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FD0, 0x1FD3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FD4, 0x1FD5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FD6, 0x1FDB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FDC, 0x1FDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FE0, 0x1FEC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FED, 0x1FF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FF2, 0x1FF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FF5, 0x1FF5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x209D, 0x2101, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x2116, 0x2117, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x2129, 0x2129, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x214F, 0x215F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x2CF4, 0x2CFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x2D68, 0x2D6E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x2DAF, 0x2DAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x2DCF, 0x2DCF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x302A, 0x3030, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x30A0, 0x30A0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x318F, 0x319F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA48D, 0xA4CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA62C, 0xA63F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA720, 0xA721, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA7D4, 0xA7D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA80B, 0xA80B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA8F8, 0xA8FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA947, 0xA95F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xA9E5, 0xA9E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xAA43, 0xAA43, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xAAB0, 0xAAB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xAAC1, 0xAAC1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xAAF5, 0xAB00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xAB27, 0xAB27, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xABE3, 0xABFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xFA6E, 0xFA6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xFB1E, 0xFB1E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xFB3F, 0xFB3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xFD3E, 0xFD4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xFE75, 0xFE75, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xFFBF, 0xFFC1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0xFFDD, 0xFFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1003E, 0x1003E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10175, 0x1027F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1034B, 0x1034F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x103D0, 0x103D0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x104FC, 0x104FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1058B, 0x1058B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x105B2, 0x105B2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10756, 0x1075F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x107BB, 0x107FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10839, 0x1083B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1089F, 0x108DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1093A, 0x1097F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10A14, 0x10A14, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10A9D, 0x10ABF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10B56, 0x10B5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10CB3, 0x10CBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10EB2, 0x10EFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x10F82, 0x10FAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11073, 0x11074, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11127, 0x11143, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11177, 0x11182, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x111DD, 0x111FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11287, 0x11287, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x112A9, 0x112AF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11329, 0x11329, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1133E, 0x1134F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1144B, 0x1145E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x114C8, 0x1157F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11645, 0x1167F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11747, 0x117FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1190A, 0x1190B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11940, 0x11940, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x119E2, 0x119E2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11A3B, 0x11A4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11AF9, 0x11BFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11C90, 0x11CFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11D47, 0x11D5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11D99, 0x11EDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x11F34, 0x11FAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x12544, 0x12F8F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x14647, 0x167FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x16AEE, 0x16AFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x16B90, 0x16E3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x16FA0, 0x16FDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x18CD6, 0x18CFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1AFFF, 0x1AFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1B156, 0x1B163, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1BC7D, 0x1BC7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D49D, 0x1D49D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D4AD, 0x1D4AD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D506, 0x1D506, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D53A, 0x1D53A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D551, 0x1D551, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D6FB, 0x1D6FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D76F, 0x1D76F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D7CC, 0x1DEFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1E12D, 0x1E136, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1E2EC, 0x1E4CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1E7EF, 0x1E7EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1E94C, 0x1EDFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE25, 0x1EE26, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE3A, 0x1EE3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE4A, 0x1EE4A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE55, 0x1EE56, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE5E, 0x1EE5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE6B, 0x1EE6B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE7F, 0x1EE7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EEAA, 0x1EEAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x2B81E, 0x2B81F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x2FA1E, 0x2FFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIdsN,
			&autogenTestCasesFuncIdsY,
		}},
	},
	{ // Patsyn
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
			&autogenTestCasesFuncPatsynY,
		}},
		{0x00, 0x20, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x21, 0x2F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x30, 0x39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x3A, 0x40, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x41, 0x5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x5B, 0x5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x5F, 0x5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x60, 0x60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x61, 0x7A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x7B, 0x7E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x7F, 0xA0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0xA1, 0xA7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0xA8, 0xA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0xA9, 0xA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0xAA, 0xAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0xAB, 0xAC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0xAD, 0xAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0xAE, 0xAE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0xAF, 0xAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0xB0, 0xB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0xB2, 0xB5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0xB6, 0xB6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0xB7, 0xBA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0xBB, 0xBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0xBC, 0xBE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0xBF, 0xBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0xC0, 0xD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0xD7, 0xD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0xD8, 0xF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0xF7, 0xF7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0xF8, 0x200F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x2010, 0x2027, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x2028, 0x202F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x2030, 0x203E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x203F, 0x2040, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x2041, 0x2053, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x2054, 0x2054, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x2055, 0x205E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x205F, 0x218F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x2190, 0x245F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x2460, 0x24FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x2500, 0x2775, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x2776, 0x2793, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x2794, 0x2BFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x2C00, 0x2DFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x2E00, 0x2E7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x2E80, 0x3000, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x3001, 0x3003, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x3004, 0x3007, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x3008, 0x3020, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x3021, 0x302F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x3030, 0x3030, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0x3031, 0xFD3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0xFD3E, 0xFD3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0xFD40, 0xFE44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0xFE45, 0xFE46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynY,
		}, []*autogenTestCaseFunc{}},
		{0xFE47, 0x10FFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatsynN,
			&autogenTestCasesFuncPatsynY,
		}},
	},
	{ // Patws
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatwsN,
			&autogenTestCasesFuncPatwsY,
		}},
		{0x00, 0x08, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatwsN,
		}, []*autogenTestCaseFunc{}},
		{0x09, 0x0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatwsY,
		}, []*autogenTestCaseFunc{}},
		{0x0E, 0x1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatwsN,
		}, []*autogenTestCaseFunc{}},
		{0x20, 0x20, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatwsY,
		}, []*autogenTestCaseFunc{}},
		{0x21, 0x84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatwsN,
		}, []*autogenTestCaseFunc{}},
		{0x85, 0x85, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatwsY,
		}, []*autogenTestCaseFunc{}},
		{0x86, 0x200D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatwsN,
		}, []*autogenTestCaseFunc{}},
		{0x200E, 0x200F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatwsY,
		}, []*autogenTestCaseFunc{}},
		{0x2010, 0x2027, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatwsN,
		}, []*autogenTestCaseFunc{}},
		{0x2028, 0x2029, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatwsY,
		}, []*autogenTestCaseFunc{}},
		{0x202A, 0x10FFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatwsN,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncPatwsN,
			&autogenTestCasesFuncPatwsY,
		}},
	},
	{ // Xidc
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
			&autogenTestCasesFuncXidcY,
		}},
		{0x00, 0x2F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x30, 0x39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x3A, 0x40, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x41, 0x5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x5B, 0x5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x5F, 0x5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x60, 0x60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x61, 0x7A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x7B, 0xA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xAA, 0xAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xAB, 0xB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB5, 0xB5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB6, 0xB6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB7, 0xB7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB8, 0xB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xBA, 0xBA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xBB, 0xBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC0, 0xD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xD7, 0xD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD8, 0xF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xF7, 0xF7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xF8, 0x2C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x2C2, 0x2C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x2C6, 0x2D1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x2D2, 0x2DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x2E0, 0x2E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x2E5, 0x2EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x2EC, 0x2EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x2ED, 0x2ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x2EE, 0x2EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x2EF, 0x2FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x300, 0x374, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x375, 0x375, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x376, 0x377, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x378, 0x37A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x37B, 0x37D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x37E, 0x37E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x37F, 0x37F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x380, 0x385, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x386, 0x38A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x38B, 0x38B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x38C, 0x38C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x38D, 0x38D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x38E, 0x3A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x3A2, 0x3A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x3A3, 0x3F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x3F6, 0x3F6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x3F7, 0x481, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x482, 0x482, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x483, 0x487, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x488, 0x489, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x48A, 0x52F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x530, 0x530, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x531, 0x556, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x557, 0x558, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x559, 0x559, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x55A, 0x55F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x560, 0x588, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x589, 0x590, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x591, 0x5BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x5BE, 0x5BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x5BF, 0x5BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x5C0, 0x5C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x5C1, 0x5C2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x5C3, 0x5C3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x5C4, 0x5C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x5C6, 0x5C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x5C7, 0x5C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x5C8, 0x5CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x5D0, 0x5EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x5EB, 0x5EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x5EF, 0x5F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x5F3, 0x60F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x610, 0x61A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x61B, 0x61F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x620, 0x669, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x66A, 0x66D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x66E, 0x6D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x6D4, 0x6D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x6D5, 0x6DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x6DD, 0x6DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x6DF, 0x6E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x6E9, 0x6E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x6EA, 0x6FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x6FD, 0x6FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x6FF, 0x6FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x700, 0x70F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x710, 0x74A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x74B, 0x74C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x74D, 0x7B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x7B2, 0x7BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x7C0, 0x7F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x7F6, 0x7F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x7FA, 0x7FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x7FB, 0x7FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x7FD, 0x7FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x7FE, 0x7FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x800, 0x82D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x82E, 0x83F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x840, 0x85B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x85C, 0x85F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x860, 0x86A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x86B, 0x86F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x870, 0x887, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x888, 0x888, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x889, 0x88E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x88F, 0x897, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x898, 0x8E1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x8E2, 0x8E2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x8E3, 0x963, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x964, 0x965, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x966, 0x96F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x970, 0x970, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x971, 0x983, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x984, 0x984, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x985, 0x98C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x98D, 0x98E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x98F, 0x990, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x991, 0x992, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x993, 0x9A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x9A9, 0x9A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x9AA, 0x9B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x9B1, 0x9B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x9B2, 0x9B2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x9B3, 0x9B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x9B6, 0x9B9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x9BA, 0x9BB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x9BC, 0x9C4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x9C5, 0x9C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x9C7, 0x9C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x9C9, 0x9CA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x9CB, 0x9CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x9CF, 0x9D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x9D7, 0x9D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x9D8, 0x9DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x9DC, 0x9DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x9DE, 0x9DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x9DF, 0x9E3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x9E4, 0x9E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x9E6, 0x9F1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x9F2, 0x9FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x9FC, 0x9FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x9FD, 0x9FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x9FE, 0x9FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x9FF, 0xA00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA01, 0xA03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA04, 0xA04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA05, 0xA0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA0B, 0xA0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA0F, 0xA10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA11, 0xA12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA13, 0xA28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA29, 0xA29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA2A, 0xA30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA31, 0xA31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA32, 0xA33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA34, 0xA34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA35, 0xA36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA37, 0xA37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA38, 0xA39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA3A, 0xA3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA3C, 0xA3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA3D, 0xA3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA3E, 0xA42, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA43, 0xA46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA47, 0xA48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA49, 0xA4A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA4B, 0xA4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA4E, 0xA50, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA51, 0xA51, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA52, 0xA58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA59, 0xA5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA5D, 0xA5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA5E, 0xA5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA5F, 0xA65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA66, 0xA75, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA76, 0xA80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA81, 0xA83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA84, 0xA84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA85, 0xA8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA8E, 0xA8E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA8F, 0xA91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xA92, 0xA92, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA93, 0xAA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xAA9, 0xAA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xAAA, 0xAB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xAB1, 0xAB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xAB2, 0xAB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xAB4, 0xAB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xAB5, 0xAB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xABA, 0xABB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xABC, 0xAC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xAC6, 0xAC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xAC7, 0xAC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xACA, 0xACA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xACB, 0xACD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xACE, 0xACF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xAD0, 0xAD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xAD1, 0xADF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xAE0, 0xAE3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xAE4, 0xAE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xAE6, 0xAEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xAF0, 0xAF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xAF9, 0xAFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB00, 0xB00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB01, 0xB03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB04, 0xB04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB05, 0xB0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB0D, 0xB0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB0F, 0xB10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB11, 0xB12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB13, 0xB28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB29, 0xB29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB2A, 0xB30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB31, 0xB31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB32, 0xB33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB34, 0xB34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB35, 0xB39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB3A, 0xB3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB3C, 0xB44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB45, 0xB46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB47, 0xB48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB49, 0xB4A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB4B, 0xB4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB4E, 0xB54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB55, 0xB57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB58, 0xB5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB5C, 0xB5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB5E, 0xB5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB5F, 0xB63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB64, 0xB65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB66, 0xB6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB70, 0xB70, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB71, 0xB71, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB72, 0xB81, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB82, 0xB83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB84, 0xB84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB85, 0xB8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB8B, 0xB8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB8E, 0xB90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB91, 0xB91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB92, 0xB95, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB96, 0xB98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB99, 0xB9A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB9B, 0xB9B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB9C, 0xB9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xB9D, 0xB9D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xB9E, 0xB9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xBA0, 0xBA2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xBA3, 0xBA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xBA5, 0xBA7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xBA8, 0xBAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xBAB, 0xBAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xBAE, 0xBB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xBBA, 0xBBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xBBE, 0xBC2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xBC3, 0xBC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xBC6, 0xBC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xBC9, 0xBC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xBCA, 0xBCD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xBCE, 0xBCF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xBD0, 0xBD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xBD1, 0xBD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xBD7, 0xBD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xBD8, 0xBE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xBE6, 0xBEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xBF0, 0xBFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC00, 0xC0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC0D, 0xC0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC0E, 0xC10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC11, 0xC11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC12, 0xC28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC29, 0xC29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC2A, 0xC39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC3A, 0xC3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC3C, 0xC44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC45, 0xC45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC46, 0xC48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC49, 0xC49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC4A, 0xC4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC4E, 0xC54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC55, 0xC56, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC57, 0xC57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC58, 0xC5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC5B, 0xC5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC5D, 0xC5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC5E, 0xC5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC60, 0xC63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC64, 0xC65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC66, 0xC6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC70, 0xC7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC80, 0xC83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC84, 0xC84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC85, 0xC8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC8D, 0xC8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC8E, 0xC90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xC91, 0xC91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xC92, 0xCA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xCA9, 0xCA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xCAA, 0xCB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xCB4, 0xCB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xCB5, 0xCB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xCBA, 0xCBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xCBC, 0xCC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xCC5, 0xCC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xCC6, 0xCC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xCC9, 0xCC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xCCA, 0xCCD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xCCE, 0xCD4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xCD5, 0xCD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xCD7, 0xCDC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xCDD, 0xCDE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xCDF, 0xCDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xCE0, 0xCE3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xCE4, 0xCE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xCE6, 0xCEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xCF0, 0xCF0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xCF1, 0xCF3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xCF4, 0xCFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD00, 0xD0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xD0D, 0xD0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD0E, 0xD10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xD11, 0xD11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD12, 0xD44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xD45, 0xD45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD46, 0xD48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xD49, 0xD49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD4A, 0xD4E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xD4F, 0xD53, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD54, 0xD57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xD58, 0xD5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD5F, 0xD63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xD64, 0xD65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD66, 0xD6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xD70, 0xD79, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD7A, 0xD7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xD80, 0xD80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD81, 0xD83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xD84, 0xD84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD85, 0xD96, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xD97, 0xD99, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD9A, 0xDB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xDB2, 0xDB2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xDB3, 0xDBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xDBC, 0xDBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xDBD, 0xDBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xDBE, 0xDBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xDC0, 0xDC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xDC7, 0xDC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xDCA, 0xDCA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xDCB, 0xDCE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xDCF, 0xDD4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xDD5, 0xDD5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xDD6, 0xDD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xDD7, 0xDD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xDD8, 0xDDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xDE0, 0xDE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xDE6, 0xDEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xDF0, 0xDF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xDF2, 0xDF3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xDF4, 0xE00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xE01, 0xE3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xE3B, 0xE3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xE40, 0xE4E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xE4F, 0xE4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xE50, 0xE59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xE5A, 0xE80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xE81, 0xE82, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xE83, 0xE83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xE84, 0xE84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xE85, 0xE85, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xE86, 0xE8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xE8B, 0xE8B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xE8C, 0xEA3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xEA4, 0xEA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xEA5, 0xEA5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xEA6, 0xEA6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xEA7, 0xEBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xEBE, 0xEBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xEC0, 0xEC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xEC5, 0xEC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xEC6, 0xEC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xEC7, 0xEC7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xEC8, 0xECE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xECF, 0xECF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xED0, 0xED9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xEDA, 0xEDB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xEDC, 0xEDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xEE0, 0xEFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xF00, 0xF00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xF01, 0xF17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xF18, 0xF19, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xF1A, 0xF1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xF20, 0xF29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xF2A, 0xF34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xF35, 0xF35, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xF36, 0xF36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xF37, 0xF37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xF38, 0xF38, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xF39, 0xF39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xF3A, 0xF3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xF3E, 0xF47, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xF48, 0xF48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xF49, 0xF6C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xF6D, 0xF70, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xF71, 0xF84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xF85, 0xF85, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xF86, 0xF97, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xF98, 0xF98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xF99, 0xFBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xFBD, 0xFC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xFC6, 0xFC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0xFC7, 0xFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1000, 0x1049, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x104A, 0x104F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1050, 0x109D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x109E, 0x109F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10A0, 0x10C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x10C6, 0x10C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10C7, 0x10C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x10C8, 0x10CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10CD, 0x10CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x10CE, 0x10CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10D0, 0x10FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x10FB, 0x10FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10FC, 0x1248, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1249, 0x1249, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x124A, 0x124D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x124E, 0x124F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1250, 0x1256, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1257, 0x1257, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1258, 0x1258, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1259, 0x1259, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x125A, 0x125D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x125E, 0x125F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1260, 0x1288, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1289, 0x1289, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x128A, 0x128D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x128E, 0x128F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1290, 0x12B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x12B1, 0x12B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x12B2, 0x12B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x12B6, 0x12B7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x12B8, 0x12BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x12BF, 0x12BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x12C0, 0x12C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x12C1, 0x12C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x12C2, 0x12C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x12C6, 0x12C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x12C8, 0x12D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x12D7, 0x12D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x12D8, 0x1310, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1311, 0x1311, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1312, 0x1315, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1316, 0x1317, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1318, 0x135A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x135B, 0x135C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x135D, 0x135F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1360, 0x1368, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1369, 0x1371, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1372, 0x137F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1380, 0x138F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1390, 0x139F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x13A0, 0x13F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x13F6, 0x13F7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x13F8, 0x13FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x13FE, 0x1400, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1401, 0x166C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x166D, 0x166E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x166F, 0x167F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1680, 0x1680, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1681, 0x169A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x169B, 0x169F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x16A0, 0x16EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x16EB, 0x16ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x16EE, 0x16F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x16F9, 0x16FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1700, 0x1715, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1716, 0x171E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x171F, 0x1734, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1735, 0x173F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1740, 0x1753, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1754, 0x175F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1760, 0x176C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x176D, 0x176D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x176E, 0x1770, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1771, 0x1771, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1772, 0x1773, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x1774, 0x177F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1780, 0x17D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x17D4, 0x17D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x17D7, 0x17D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x17D8, 0x17DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x17DC, 0x17DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x17DE, 0x17DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x17E0, 0x17E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x17EA, 0x180A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x180B, 0x180D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x180E, 0x180E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x180F, 0x1819, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcY,
		}, []*autogenTestCaseFunc{}},
		{0x181A, 0x181F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x191F, 0x191F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1975, 0x197F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1A1C, 0x1A1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1A9A, 0x1AA6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1B4D, 0x1B4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1C38, 0x1C3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1CBB, 0x1CBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1F16, 0x1F17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1F58, 0x1F58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1F7E, 0x1F7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1FC5, 0x1FC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1FED, 0x1FF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x2041, 0x2053, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x209D, 0x20CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x2103, 0x2106, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x211E, 0x2123, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x213A, 0x213B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x2189, 0x2BFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x2D28, 0x2D2C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x2D97, 0x2D9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x2DBF, 0x2DBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x2DDF, 0x2DDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x3036, 0x3037, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x30A0, 0x30A0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x31C0, 0x31EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA4FE, 0xA4FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA67E, 0xA67E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA7CB, 0xA7CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA828, 0xA82B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA8DA, 0xA8DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA954, 0xA95F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xA9FF, 0xA9FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xAA77, 0xAA79, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xAAF7, 0xAB00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xAB27, 0xAB27, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xABEB, 0xABEB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xD7C7, 0xD7CA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xFB07, 0xFB12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xFB3D, 0xFB3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xFBB2, 0xFBD2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xFDC8, 0xFDEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xFE35, 0xFE4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xFE78, 0xFE78, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xFEFD, 0xFF0F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xFF5B, 0xFF64, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0xFFD8, 0xFFD9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1003B, 0x1003B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x100FB, 0x1013F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x102D1, 0x102DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1037B, 0x1037F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x103D6, 0x103FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x104FC, 0x104FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1058B, 0x1058B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x105B2, 0x105B2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10756, 0x1075F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x107BB, 0x107FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10839, 0x1083B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1089F, 0x108DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1093A, 0x1097F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10A07, 0x10A0B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10A3B, 0x10A3E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10AC8, 0x10AC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10B73, 0x10B7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10CF3, 0x10CFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10EAD, 0x10EAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x10F51, 0x10F6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11047, 0x11065, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x110E9, 0x110EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11148, 0x1114F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x111CD, 0x111CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11238, 0x1123D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1128E, 0x1128E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x112FA, 0x112FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11329, 0x11329, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11345, 0x11346, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11358, 0x1135C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1144B, 0x1144F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x114C8, 0x114CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x115DE, 0x115FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x116B9, 0x116BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1173A, 0x1173F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11907, 0x11908, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11936, 0x11936, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x119A8, 0x119A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11A3F, 0x11A46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11AF9, 0x11BFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11C5A, 0x11C71, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11D07, 0x11D07, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11D3E, 0x11D3E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11D69, 0x11D69, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11DAA, 0x11EDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x11F43, 0x11F4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1246F, 0x1247F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x13456, 0x143FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x16A6A, 0x16A6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x16AF5, 0x16AFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x16B78, 0x16B7C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x16F88, 0x16F8E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x16FF2, 0x16FFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1AFF4, 0x1AFF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1B133, 0x1B14F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1B2FC, 0x1BBFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1BC9A, 0x1BC9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D16A, 0x1D16C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D1AE, 0x1D241, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D4A0, 0x1D4A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D4BA, 0x1D4BA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D50B, 0x1D50C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D53F, 0x1D53F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D6A6, 0x1D6A7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D715, 0x1D715, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D789, 0x1D789, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1D800, 0x1D9FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1DA85, 0x1DA9A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1DF2B, 0x1DFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1E025, 0x1E025, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1E12D, 0x1E12F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1E2AF, 0x1E2BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1E7EC, 0x1E7EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1E8D7, 0x1E8FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE20, 0x1EE20, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE33, 0x1EE33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE43, 0x1EE46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE50, 0x1EE50, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE5A, 0x1EE5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE63, 0x1EE63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE78, 0x1EE78, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE9C, 0x1EEA0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x1FBFA, 0x1FFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x2CEA2, 0x2CEAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x3134B, 0x3134F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidcN,
			&autogenTestCasesFuncXidcY,
		}},
	},
	{ // Xids
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
			&autogenTestCasesFuncXidsY,
		}},
		{0x00, 0x40, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x41, 0x5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x5B, 0x60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x61, 0x7A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x7B, 0xA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xAA, 0xAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xAB, 0xB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB5, 0xB5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB6, 0xB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xBA, 0xBA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xBB, 0xBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xC0, 0xD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xD7, 0xD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xD8, 0xF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xF7, 0xF7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xF8, 0x2C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x2C2, 0x2C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x2C6, 0x2D1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x2D2, 0x2DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x2E0, 0x2E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x2E5, 0x2EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x2EC, 0x2EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x2ED, 0x2ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x2EE, 0x2EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x2EF, 0x36F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x370, 0x374, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x375, 0x375, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x376, 0x377, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x378, 0x37A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x37B, 0x37D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x37E, 0x37E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x37F, 0x37F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x380, 0x385, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x386, 0x386, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x387, 0x387, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x388, 0x38A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x38B, 0x38B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x38C, 0x38C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x38D, 0x38D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x38E, 0x3A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x3A2, 0x3A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x3A3, 0x3F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x3F6, 0x3F6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x3F7, 0x481, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x482, 0x489, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x48A, 0x52F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x530, 0x530, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x531, 0x556, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x557, 0x558, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x559, 0x559, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x55A, 0x55F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x560, 0x588, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x589, 0x5CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x5D0, 0x5EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x5EB, 0x5EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x5EF, 0x5F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x5F3, 0x61F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x620, 0x64A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x64B, 0x66D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x66E, 0x66F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x670, 0x670, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x671, 0x6D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x6D4, 0x6D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x6D5, 0x6D5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x6D6, 0x6E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x6E5, 0x6E6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x6E7, 0x6ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x6EE, 0x6EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x6F0, 0x6F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x6FA, 0x6FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x6FD, 0x6FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x6FF, 0x6FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x700, 0x70F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x710, 0x710, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x711, 0x711, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x712, 0x72F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x730, 0x74C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x74D, 0x7A5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x7A6, 0x7B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x7B1, 0x7B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x7B2, 0x7C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x7CA, 0x7EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x7EB, 0x7F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x7F4, 0x7F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x7F6, 0x7F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x7FA, 0x7FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x7FB, 0x7FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x800, 0x815, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x816, 0x819, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x81A, 0x81A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x81B, 0x823, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x824, 0x824, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x825, 0x827, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x828, 0x828, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x829, 0x83F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x840, 0x858, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x859, 0x85F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x860, 0x86A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x86B, 0x86F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x870, 0x887, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x888, 0x888, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x889, 0x88E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x88F, 0x89F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x8A0, 0x8C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x8CA, 0x903, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x904, 0x939, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x93A, 0x93C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x93D, 0x93D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x93E, 0x94F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x950, 0x950, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x951, 0x957, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x958, 0x961, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x962, 0x970, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x971, 0x980, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x981, 0x984, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x985, 0x98C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x98D, 0x98E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x98F, 0x990, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x991, 0x992, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x993, 0x9A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x9A9, 0x9A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x9AA, 0x9B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x9B1, 0x9B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x9B2, 0x9B2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x9B3, 0x9B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x9B6, 0x9B9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x9BA, 0x9BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x9BD, 0x9BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x9BE, 0x9CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x9CE, 0x9CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x9CF, 0x9DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x9DC, 0x9DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x9DE, 0x9DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x9DF, 0x9E1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x9E2, 0x9EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x9F0, 0x9F1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x9F2, 0x9FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x9FC, 0x9FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x9FD, 0xA04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA05, 0xA0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xA0B, 0xA0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA0F, 0xA10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xA11, 0xA12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA13, 0xA28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xA29, 0xA29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA2A, 0xA30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xA31, 0xA31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA32, 0xA33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xA34, 0xA34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA35, 0xA36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xA37, 0xA37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA38, 0xA39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xA3A, 0xA58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA59, 0xA5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xA5D, 0xA5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA5E, 0xA5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xA5F, 0xA71, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA72, 0xA74, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xA75, 0xA84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA85, 0xA8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xA8E, 0xA8E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA8F, 0xA91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xA92, 0xA92, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA93, 0xAA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xAA9, 0xAA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xAAA, 0xAB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xAB1, 0xAB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xAB2, 0xAB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xAB4, 0xAB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xAB5, 0xAB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xABA, 0xABC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xABD, 0xABD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xABE, 0xACF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xAD0, 0xAD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xAD1, 0xADF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xAE0, 0xAE1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xAE2, 0xAF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xAF9, 0xAF9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xAFA, 0xB04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB05, 0xB0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB0D, 0xB0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB0F, 0xB10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB11, 0xB12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB13, 0xB28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB29, 0xB29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB2A, 0xB30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB31, 0xB31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB32, 0xB33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB34, 0xB34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB35, 0xB39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB3A, 0xB3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB3D, 0xB3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB3E, 0xB5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB5C, 0xB5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB5E, 0xB5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB5F, 0xB61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB62, 0xB70, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB71, 0xB71, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB72, 0xB82, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB83, 0xB83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB84, 0xB84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB85, 0xB8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB8B, 0xB8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB8E, 0xB90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB91, 0xB91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB92, 0xB95, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB96, 0xB98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB99, 0xB9A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB9B, 0xB9B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB9C, 0xB9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xB9D, 0xB9D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xB9E, 0xB9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xBA0, 0xBA2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xBA3, 0xBA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xBA5, 0xBA7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xBA8, 0xBAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xBAB, 0xBAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xBAE, 0xBB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xBBA, 0xBCF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xBD0, 0xBD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xBD1, 0xC04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xC05, 0xC0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xC0D, 0xC0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xC0E, 0xC10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xC11, 0xC11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xC12, 0xC28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xC29, 0xC29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xC2A, 0xC39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xC3A, 0xC3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xC3D, 0xC3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xC3E, 0xC57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xC58, 0xC5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xC5B, 0xC5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xC5D, 0xC5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xC5E, 0xC5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xC60, 0xC61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xC62, 0xC7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xC80, 0xC80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xC81, 0xC84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xC85, 0xC8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xC8D, 0xC8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xC8E, 0xC90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xC91, 0xC91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xC92, 0xCA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xCA9, 0xCA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xCAA, 0xCB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xCB4, 0xCB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xCB5, 0xCB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xCBA, 0xCBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xCBD, 0xCBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xCBE, 0xCDC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xCDD, 0xCDE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xCDF, 0xCDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xCE0, 0xCE1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xCE2, 0xCF0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xCF1, 0xCF2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xCF3, 0xD03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xD04, 0xD0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xD0D, 0xD0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xD0E, 0xD10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xD11, 0xD11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xD12, 0xD3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xD3B, 0xD3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xD3D, 0xD3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xD3E, 0xD4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xD4E, 0xD4E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xD4F, 0xD53, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xD54, 0xD56, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xD57, 0xD5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xD5F, 0xD61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xD62, 0xD79, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xD7A, 0xD7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xD80, 0xD84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xD85, 0xD96, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xD97, 0xD99, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xD9A, 0xDB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xDB2, 0xDB2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xDB3, 0xDBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xDBC, 0xDBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xDBD, 0xDBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xDBE, 0xDBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xDC0, 0xDC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xDC7, 0xE00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xE01, 0xE30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xE31, 0xE31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xE32, 0xE32, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xE33, 0xE3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xE40, 0xE46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xE47, 0xE80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xE81, 0xE82, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xE83, 0xE83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xE84, 0xE84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xE85, 0xE85, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xE86, 0xE8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xE8B, 0xE8B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xE8C, 0xEA3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xEA4, 0xEA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xEA5, 0xEA5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xEA6, 0xEA6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xEA7, 0xEB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xEB1, 0xEB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xEB2, 0xEB2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xEB3, 0xEBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xEBD, 0xEBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xEBE, 0xEBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xEC0, 0xEC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xEC5, 0xEC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xEC6, 0xEC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xEC7, 0xEDB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xEDC, 0xEDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xEE0, 0xEFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xF00, 0xF00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xF01, 0xF3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xF40, 0xF47, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xF48, 0xF48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xF49, 0xF6C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xF6D, 0xF87, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xF88, 0xF8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0xF8D, 0xFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1000, 0x102A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x102B, 0x103E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x103F, 0x103F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1040, 0x104F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1050, 0x1055, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1056, 0x1059, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x105A, 0x105D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x105E, 0x1060, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1061, 0x1061, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1062, 0x1064, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1065, 0x1066, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1067, 0x106D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x106E, 0x1070, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1071, 0x1074, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1075, 0x1081, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1082, 0x108D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x108E, 0x108E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x108F, 0x109F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10A0, 0x10C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x10C6, 0x10C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10C7, 0x10C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x10C8, 0x10CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10CD, 0x10CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x10CE, 0x10CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10D0, 0x10FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x10FB, 0x10FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10FC, 0x1248, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1249, 0x1249, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x124A, 0x124D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x124E, 0x124F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1250, 0x1256, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1257, 0x1257, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1258, 0x1258, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1259, 0x1259, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x125A, 0x125D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x125E, 0x125F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1260, 0x1288, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1289, 0x1289, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x128A, 0x128D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x128E, 0x128F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1290, 0x12B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x12B1, 0x12B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x12B2, 0x12B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x12B6, 0x12B7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x12B8, 0x12BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x12BF, 0x12BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x12C0, 0x12C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x12C1, 0x12C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x12C2, 0x12C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x12C6, 0x12C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x12C8, 0x12D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x12D7, 0x12D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x12D8, 0x1310, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1311, 0x1311, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1312, 0x1315, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1316, 0x1317, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1318, 0x135A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x135B, 0x137F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1380, 0x138F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1390, 0x139F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x13A0, 0x13F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x13F6, 0x13F7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x13F8, 0x13FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x13FE, 0x1400, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1401, 0x166C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x166D, 0x166E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x166F, 0x167F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1680, 0x1680, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1681, 0x169A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x169B, 0x169F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x16A0, 0x16EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x16EB, 0x16ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x16EE, 0x16F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x16F9, 0x16FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1700, 0x1711, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1712, 0x171E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x171F, 0x1731, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1732, 0x173F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1740, 0x1751, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1752, 0x175F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1760, 0x176C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x176D, 0x176D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x176E, 0x1770, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1771, 0x177F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1780, 0x17B3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x17B4, 0x17D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x17D7, 0x17D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x17D8, 0x17DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x17DC, 0x17DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x17DD, 0x181F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1820, 0x1878, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1879, 0x187F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1880, 0x18A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x18A9, 0x18A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x18AA, 0x18AA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x18AB, 0x18AF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x18B0, 0x18F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x18F6, 0x18FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1900, 0x191E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x191F, 0x194F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1950, 0x196D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x196E, 0x196F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1970, 0x1974, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1975, 0x197F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1980, 0x19AB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x19AC, 0x19AF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x19B0, 0x19C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x19CA, 0x19FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1A00, 0x1A16, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1A17, 0x1A1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1A20, 0x1A54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1A55, 0x1AA6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1AA7, 0x1AA7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1AA8, 0x1B04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1B05, 0x1B33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1B34, 0x1B44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1B45, 0x1B4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1B4D, 0x1B82, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1B83, 0x1BA0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1BA1, 0x1BAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1BAE, 0x1BAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1BB0, 0x1BB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1BBA, 0x1BE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1BE6, 0x1BFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1C00, 0x1C23, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1C24, 0x1C4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1C4D, 0x1C4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1C50, 0x1C59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1C5A, 0x1C7D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1C7E, 0x1C7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1C80, 0x1C88, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1C89, 0x1C8F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1C90, 0x1CBA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1CBB, 0x1CBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1CBD, 0x1CBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1CC0, 0x1CE8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1CE9, 0x1CEC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1CED, 0x1CED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1CEE, 0x1CF3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1CF4, 0x1CF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1CF5, 0x1CF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1CF7, 0x1CF9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1CFA, 0x1CFA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1CFB, 0x1CFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D00, 0x1DBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1DC0, 0x1DFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1E00, 0x1F15, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F16, 0x1F17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F18, 0x1F1D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F1E, 0x1F1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F20, 0x1F45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F46, 0x1F47, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F48, 0x1F4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F4E, 0x1F4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F50, 0x1F57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F58, 0x1F58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F59, 0x1F59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5A, 0x1F5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5B, 0x1F5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5C, 0x1F5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5D, 0x1F5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F5E, 0x1F5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F5F, 0x1F7D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1F7E, 0x1F7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1F80, 0x1FB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FB5, 0x1FB5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FB6, 0x1FBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FBD, 0x1FBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FBE, 0x1FBE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FBF, 0x1FC1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FC2, 0x1FC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FC5, 0x1FC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FC6, 0x1FCC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FCD, 0x1FCF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FD0, 0x1FD3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FD4, 0x1FD5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FD6, 0x1FDB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FDC, 0x1FDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FE0, 0x1FEC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FED, 0x1FF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1FF2, 0x1FF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsY,
		}, []*autogenTestCaseFunc{}},
		{0x1FF5, 0x1FF5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x209D, 0x2101, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x2116, 0x2117, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x2129, 0x2129, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x214F, 0x215F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x2CF4, 0x2CFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x2D68, 0x2D6E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x2DAF, 0x2DAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x2DCF, 0x2DCF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x302A, 0x3030, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x30A0, 0x30A0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x318F, 0x319F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA48D, 0xA4CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA62C, 0xA63F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA720, 0xA721, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA7D4, 0xA7D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA80B, 0xA80B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA8F8, 0xA8FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA947, 0xA95F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xA9E5, 0xA9E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xAA43, 0xAA43, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xAAB0, 0xAAB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xAAC1, 0xAAC1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xAAF5, 0xAB00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xAB27, 0xAB27, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xABE3, 0xABFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xFA6E, 0xFA6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xFB1E, 0xFB1E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xFB3F, 0xFB3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xFC5E, 0xFC63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xFDFA, 0xFE70, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xFE7A, 0xFE7A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xFF3B, 0xFF40, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0xFFC8, 0xFFC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1000C, 0x1000C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1004E, 0x1004F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1029D, 0x1029F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10376, 0x1037F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x103D6, 0x103FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10528, 0x1052F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10593, 0x10593, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x105BA, 0x105BA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10768, 0x1077F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10806, 0x10807, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1083D, 0x1083E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x108F3, 0x108F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x109B8, 0x109BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10A18, 0x10A18, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10AC8, 0x10AC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10B73, 0x10B7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10CF3, 0x10CFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10F1D, 0x10F26, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x10FC5, 0x10FDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11076, 0x11082, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11145, 0x11146, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x111B3, 0x111C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11212, 0x11212, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11289, 0x11289, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x112DF, 0x11304, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11331, 0x11331, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11351, 0x1135C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11462, 0x1147F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x115AF, 0x115D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x116AB, 0x116B7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1182C, 0x1189F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11914, 0x11914, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11942, 0x1199F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x119E4, 0x119FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11A51, 0x11A5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11C09, 0x11C09, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11D07, 0x11D07, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11D66, 0x11D66, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11EF3, 0x11F01, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x11FB1, 0x11FFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x12FF1, 0x12FFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x16A39, 0x16A3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x16B30, 0x16B3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x16E80, 0x16EFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x16FE2, 0x16FE2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x18D09, 0x1AFEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1B123, 0x1B131, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1B168, 0x1B16F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1BC89, 0x1BC8F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D4A0, 0x1D4A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D4BA, 0x1D4BA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D50B, 0x1D50C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D53F, 0x1D53F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D6A6, 0x1D6A7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D715, 0x1D715, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1D789, 0x1D789, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1DF1F, 0x1DF24, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1E13E, 0x1E14D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1E4EC, 0x1E7DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1E7FF, 0x1E7FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE04, 0x1EE04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE28, 0x1EE28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE3C, 0x1EE41, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE4C, 0x1EE4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE58, 0x1EE58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE60, 0x1EE60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE73, 0x1EE73, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EE8A, 0x1EE8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x1EEBC, 0x1FFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x2CEA2, 0x2CEAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x3134B, 0x3134F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncXidsN,
			&autogenTestCasesFuncXidsY,
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
