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

package unorm_test

import "testing"
import "github.com/orkvozku/go/uni/unorm"
import "io"
import "strings"

// normTestCase contains a function to test.
type normTestCase struct {
	label string
	data  []rune
	nfc   []rune
	nfd   []rune
	nfkc  []rune
	nfkd  []rune
}

var normTestCases = []normTestCase{
	// Trivial - ASCII:
	{"0041.0042.0043",
		[]rune{0x0041, 0x0042, 0x0043},
		[]rune{0x0041, 0x0042, 0x0043},
		[]rune{0x0041, 0x0042, 0x0043},
		[]rune{0x0041, 0x0042, 0x0043},
		[]rune{0x0041, 0x0042, 0x0043}},
	// Hangul:
	{"D4DB",
		[]rune{0xD4DB},
		[]rune{0xD4DB},
		[]rune{0x1111, 0x1171, 0x11B6},
		[]rune{0xD4DB},
		[]rune{0x1111, 0x1171, 0x11B6}},
	{"1111.1171.11B6",
		[]rune{0xD4DB},
		[]rune{0xD4DB},
		[]rune{0x1111, 0x1171, 0x11B6},
		[]rune{0xD4DB},
		[]rune{0x1111, 0x1171, 0x11B6}},
	{"D4CC.11B6",
		[]rune{0xD4DB},
		[]rune{0xD4DB},
		[]rune{0x1111, 0x1171, 0x11B6},
		[]rune{0xD4DB},
		[]rune{0x1111, 0x1171, 0x11B6}},
	{"0041.D4DB.0042",
		[]rune{0x0041, 0xD4DB, 0x0042},
		[]rune{0x0041, 0xD4DB, 0x0042},
		[]rune{0x0041, 0x1111, 0x1171, 0x11B6, 0x0042},
		[]rune{0x0041, 0xD4DB, 0x0042},
		[]rune{0x0041, 0x1111, 0x1171, 0x11B6, 0x0042}},
	{"0041.1111.1171.11B6.0042",
		[]rune{0x0041, 0xD4DB, 0x0042},
		[]rune{0x0041, 0xD4DB, 0x0042},
		[]rune{0x0041, 0x1111, 0x1171, 0x11B6, 0x0042},
		[]rune{0x0041, 0xD4DB, 0x0042},
		[]rune{0x0041, 0x1111, 0x1171, 0x11B6, 0x0042}},
	{"0041.D4CC.11B6.0042",
		[]rune{0x0041, 0xD4DB, 0x0042},
		[]rune{0x0041, 0xD4DB, 0x0042},
		[]rune{0x0041, 0x1111, 0x1171, 0x11B6, 0x0042},
		[]rune{0x0041, 0xD4DB, 0x0042},
		[]rune{0x0041, 0x1111, 0x1171, 0x11B6, 0x0042}},
	// Hangul - Incomplete Fragments:
	{"1111",
		[]rune{0x1111},
		[]rune{0x1111},
		[]rune{0x1111},
		[]rune{0x1111},
		[]rune{0x1111}},
	{"1111.1171",
		[]rune{0x1111, 0x1171},
		[]rune{0xD4CC},
		[]rune{0x1111, 0x1171},
		[]rune{0xD4CC},
		[]rune{0x1111, 0x1171}},
	{"1171.11B6",
		[]rune{0x1171, 0x11B6},
		[]rune{0x1171, 0x11B6},
		[]rune{0x1171, 0x11B6},
		[]rune{0x1171, 0x11B6},
		[]rune{0x1171, 0x11B6}},
	{"11B6",
		[]rune{0x11B6},
		[]rune{0x11B6},
		[]rune{0x11B6},
		[]rune{0x11B6},
		[]rune{0x11B6}},
	{"0041.1111.0042",
		[]rune{0x0041, 0x1111, 0x0042},
		[]rune{0x0041, 0x1111, 0x0042},
		[]rune{0x0041, 0x1111, 0x0042},
		[]rune{0x0041, 0x1111, 0x0042},
		[]rune{0x0041, 0x1111, 0x0042}},
	{"0041.1111.1171.0042",
		[]rune{0x0041, 0x1111, 0x1171, 0x0042},
		[]rune{0x0041, 0xD4CC, 0x0042},
		[]rune{0x0041, 0x1111, 0x1171, 0x0042},
		[]rune{0x0041, 0xD4CC, 0x0042},
		[]rune{0x0041, 0x1111, 0x1171, 0x0042}},
	{"0041.1171.11B6.0042",
		[]rune{0x0041, 0x1171, 0x11B6, 0x0042},
		[]rune{0x0041, 0x1171, 0x11B6, 0x0042},
		[]rune{0x0041, 0x1171, 0x11B6, 0x0042},
		[]rune{0x0041, 0x1171, 0x11B6, 0x0042},
		[]rune{0x0041, 0x1171, 0x11B6, 0x0042}},
	{"0041.11B6.0042",
		[]rune{0x0041, 0x11B6, 0x0042},
		[]rune{0x0041, 0x11B6, 0x0042},
		[]rune{0x0041, 0x11B6, 0x0042},
		[]rune{0x0041, 0x11B6, 0x0042},
		[]rune{0x0041, 0x11B6, 0x0042}},
	// Singletons:
	{"2126",
		[]rune{0x2126},
		[]rune{0x03A9},
		[]rune{0x03A9},
		[]rune{0x03A9},
		[]rune{0x03A9}},
	{"212B",
		[]rune{0x212B},
		[]rune{0x00C5},
		[]rune{0x0041, 0x030A},
		[]rune{0x00C5},
		[]rune{0x0041, 0x030A}},
	{"0041.2126.0042",
		[]rune{0x0041, 0x2126, 0x0042},
		[]rune{0x0041, 0x03A9, 0x0042},
		[]rune{0x0041, 0x03A9, 0x0042},
		[]rune{0x0041, 0x03A9, 0x0042},
		[]rune{0x0041, 0x03A9, 0x0042}},
	{"0041.212B.0042",
		[]rune{0x0041, 0x212B, 0x0042},
		[]rune{0x0041, 0x00C5, 0x0042},
		[]rune{0x0041, 0x0041, 0x030A, 0x0042},
		[]rune{0x0041, 0x00C5, 0x0042},
		[]rune{0x0041, 0x0041, 0x030A, 0x0042}},
	// Canonical Composites:
	{"00C5",
		[]rune{0x00C5},
		[]rune{0x00C5},
		[]rune{0x0041, 0x030A},
		[]rune{0x00C5},
		[]rune{0x0041, 0x030A}},
	{"00F4",
		[]rune{0x00F4},
		[]rune{0x00F4},
		[]rune{0x006F, 0x0302},
		[]rune{0x00F4},
		[]rune{0x006F, 0x0302}},
	{"0041.00C5.0042",
		[]rune{0x0041, 0x00C5, 0x0042},
		[]rune{0x0041, 0x00C5, 0x0042},
		[]rune{0x0041, 0x0041, 0x030A, 0x0042},
		[]rune{0x0041, 0x00C5, 0x0042},
		[]rune{0x0041, 0x0041, 0x030A, 0x0042}},
	{"0041.00F4.0042",
		[]rune{0x0041, 0x00F4, 0x0042},
		[]rune{0x0041, 0x00F4, 0x0042},
		[]rune{0x0041, 0x006F, 0x0302, 0x0042},
		[]rune{0x0041, 0x00F4, 0x0042},
		[]rune{0x0041, 0x006F, 0x0302, 0x0042}},
	// Decomposed Canonical Composites:
	{"0041.030A",
		[]rune{0x0041, 0x030A},
		[]rune{0x00C5},
		[]rune{0x0041, 0x030A},
		[]rune{0x00C5},
		[]rune{0x0041, 0x030A}},
	{"006F.0302",
		[]rune{0x006F, 0x0302},
		[]rune{0x00F4},
		[]rune{0x006F, 0x0302},
		[]rune{0x00F4},
		[]rune{0x006F, 0x0302}},
	{"0041.0041.030A.0042",
		[]rune{0x0041, 0x0041, 0x030A, 0x0042},
		[]rune{0x0041, 0x00C5, 0x0042},
		[]rune{0x0041, 0x0041, 0x030A, 0x0042},
		[]rune{0x0041, 0x00C5, 0x0042},
		[]rune{0x0041, 0x0041, 0x030A, 0x0042}},
	{"0041.006F.0302.0042",
		[]rune{0x0041, 0x006F, 0x0302, 0x0042},
		[]rune{0x0041, 0x00F4, 0x0042},
		[]rune{0x0041, 0x006F, 0x0302, 0x0042},
		[]rune{0x0041, 0x00F4, 0x0042},
		[]rune{0x0041, 0x006F, 0x0302, 0x0042}},
	// Multiple Combining Marks:
	{"1E69",
		[]rune{0x1E69},
		[]rune{0x1E69},
		[]rune{0x0073, 0x0323, 0x0307},
		[]rune{0x1E69},
		[]rune{0x0073, 0x0323, 0x0307}},
	{"0073.0323.0307",
		[]rune{0x0073, 0x0323, 0x0307},
		[]rune{0x1E69},
		[]rune{0x0073, 0x0323, 0x0307},
		[]rune{0x1E69},
		[]rune{0x0073, 0x0323, 0x0307}},
	{"0073.0307.0323",
		[]rune{0x0073, 0x0307, 0x0323},
		[]rune{0x1E69},
		[]rune{0x0073, 0x0323, 0x0307},
		[]rune{0x1E69},
		[]rune{0x0073, 0x0323, 0x0307}},
	{"0064.0323.0307",
		[]rune{0x0064, 0x0323, 0x0307},
		[]rune{0x1E0D, 0x0307},
		[]rune{0x0064, 0x0323, 0x0307},
		[]rune{0x1E0D, 0x0307},
		[]rune{0x0064, 0x0323, 0x0307}},
	{"0064.0307.0323",
		[]rune{0x0064, 0x0307, 0x0323},
		[]rune{0x1E0D, 0x0307},
		[]rune{0x0064, 0x0323, 0x0307},
		[]rune{0x1E0D, 0x0307},
		[]rune{0x0064, 0x0323, 0x0307}},
	{"1E0B.0323",
		[]rune{0x1E0B, 0x0323},
		[]rune{0x1E0D, 0x0307},
		[]rune{0x0064, 0x0323, 0x0307},
		[]rune{0x1E0D, 0x0307},
		[]rune{0x0064, 0x0323, 0x0307}},
	{"0071.0307.0323",
		[]rune{0x0071, 0x0307, 0x0323},
		[]rune{0x0071, 0x0323, 0x0307},
		[]rune{0x0071, 0x0323, 0x0307},
		[]rune{0x0071, 0x0323, 0x0307},
		[]rune{0x0071, 0x0323, 0x0307}},
	{"1EA5.0328",
		[]rune{0x1EA5, 0x0328},
		[]rune{0x0105, 0x0302, 0x0301},
		[]rune{0x0061, 0x0328, 0x0302, 0x0301},
		[]rune{0x0105, 0x0302, 0x0301},
		[]rune{0x0061, 0x0328, 0x0302, 0x0301}},
	{"0051.0300.0323",
		[]rune{0x0051, 0x0300, 0x0323},
		[]rune{0x0051, 0x0323, 0x0300},
		[]rune{0x0051, 0x0323, 0x0300},
		[]rune{0x0051, 0x0323, 0x0300},
		[]rune{0x0051, 0x0323, 0x0300}},
	{"0041.1E69.0042",
		[]rune{0x0041, 0x1E69, 0x0042},
		[]rune{0x0041, 0x1E69, 0x0042},
		[]rune{0x0041, 0x0073, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x1E69, 0x0042},
		[]rune{0x0041, 0x0073, 0x0323, 0x0307, 0x0042}},
	{"0041.0073.0323.0307.0042",
		[]rune{0x0041, 0x0073, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x1E69, 0x0042},
		[]rune{0x0041, 0x0073, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x1E69, 0x0042},
		[]rune{0x0041, 0x0073, 0x0323, 0x0307, 0x0042}},
	{"0041.0073.0307.0323.0042",
		[]rune{0x0041, 0x0073, 0x0307, 0x0323, 0x0042},
		[]rune{0x0041, 0x1E69, 0x0042},
		[]rune{0x0041, 0x0073, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x1E69, 0x0042},
		[]rune{0x0041, 0x0073, 0x0323, 0x0307, 0x0042}},
	{"0041.0064.0323.0307.0042",
		[]rune{0x0041, 0x0064, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x1E0D, 0x0307, 0x0042},
		[]rune{0x0041, 0x0064, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x1E0D, 0x0307, 0x0042},
		[]rune{0x0041, 0x0064, 0x0323, 0x0307, 0x0042}},
	{"0041.0064.0307.0323.0042",
		[]rune{0x0041, 0x0064, 0x0307, 0x0323, 0x0042},
		[]rune{0x0041, 0x1E0D, 0x0307, 0x0042},
		[]rune{0x0041, 0x0064, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x1E0D, 0x0307, 0x0042},
		[]rune{0x0041, 0x0064, 0x0323, 0x0307, 0x0042}},
	{"0041.1E0B.0323.0042",
		[]rune{0x0041, 0x1E0B, 0x0323, 0x0042},
		[]rune{0x0041, 0x1E0D, 0x0307, 0x0042},
		[]rune{0x0041, 0x0064, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x1E0D, 0x0307, 0x0042},
		[]rune{0x0041, 0x0064, 0x0323, 0x0307, 0x0042}},
	{"0041.0071.0307.0323.0042",
		[]rune{0x0041, 0x0071, 0x0307, 0x0323, 0x0042},
		[]rune{0x0041, 0x0071, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x0071, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x0071, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x0071, 0x0323, 0x0307, 0x0042}},
	{"0041.1EA5.0328.0042",
		[]rune{0x0041, 0x1EA5, 0x0328, 0x0042},
		[]rune{0x0041, 0x0105, 0x0302, 0x0301, 0x0042},
		[]rune{0x0041, 0x0061, 0x0328, 0x0302, 0x0301, 0x0042},
		[]rune{0x0041, 0x0105, 0x0302, 0x0301, 0x0042},
		[]rune{0x0041, 0x0061, 0x0328, 0x0302, 0x0301, 0x0042}},
	{"0041.0051.0300.0323.0042",
		[]rune{0x0041, 0x0051, 0x0300, 0x0323, 0x0042},
		[]rune{0x0041, 0x0051, 0x0323, 0x0300, 0x0042},
		[]rune{0x0041, 0x0051, 0x0323, 0x0300, 0x0042},
		[]rune{0x0041, 0x0051, 0x0323, 0x0300, 0x0042},
		[]rune{0x0041, 0x0051, 0x0323, 0x0300, 0x0042}},
	// Compatibility Composites:
	{"FB01",
		[]rune{0xFB01},
		[]rune{0xFB01},
		[]rune{0xFB01},
		[]rune{0x0066, 0x0069},
		[]rune{0x0066, 0x0069}},
	{"0032.2075",
		[]rune{0x0032, 0x2075},
		[]rune{0x0032, 0x2075},
		[]rune{0x0032, 0x2075},
		[]rune{0x0032, 0x0035},
		[]rune{0x0032, 0x0035}},
	{"1E9B.0323",
		[]rune{0x1E9B, 0x0323},
		[]rune{0x1E9B, 0x0323},
		[]rune{0x017F, 0x0323, 0x0307},
		[]rune{0x1E69},
		[]rune{0x0073, 0x0323, 0x0307}},
	{"017F.0323.0307",
		[]rune{0x017F, 0x323, 0x0307},
		[]rune{0x1E9B, 0x0323},
		[]rune{0x017F, 0x0323, 0x0307},
		[]rune{0x1E69},
		[]rune{0x0073, 0x0323, 0x0307}},
	{"017F.0307.0323",
		[]rune{0x017F, 0x0307, 0x0323},
		[]rune{0x1E9B, 0x0323},
		[]rune{0x017F, 0x0323, 0x0307},
		[]rune{0x1E69},
		[]rune{0x0073, 0x0323, 0x0307}},
	{"0041.FB01.0042",
		[]rune{0x0041, 0xFB01, 0x0042},
		[]rune{0x0041, 0xFB01, 0x0042},
		[]rune{0x0041, 0xFB01, 0x0042},
		[]rune{0x0041, 0x0066, 0x0069, 0x0042},
		[]rune{0x0041, 0x0066, 0x0069, 0x0042}},
	{"0041.0032.2075.0042",
		[]rune{0x0041, 0x0032, 0x2075, 0x0042},
		[]rune{0x0041, 0x0032, 0x2075, 0x0042},
		[]rune{0x0041, 0x0032, 0x2075, 0x0042},
		[]rune{0x0041, 0x0032, 0x0035, 0x0042},
		[]rune{0x0041, 0x0032, 0x0035, 0x0042}},
	{"0041.1E9B.0323.0042",
		[]rune{0x0041, 0x1E9B, 0x0323, 0x0042},
		[]rune{0x0041, 0x1E9B, 0x0323, 0x0042},
		[]rune{0x0041, 0x017F, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x1E69, 0x0042},
		[]rune{0x0041, 0x0073, 0x0323, 0x0307, 0x0042}},
	{"0041.017F.0323.0307.0042",
		[]rune{0x0041, 0x017F, 0x323, 0x0307, 0x0042},
		[]rune{0x0041, 0x1E9B, 0x0323, 0x0042},
		[]rune{0x0041, 0x017F, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x1E69, 0x0042},
		[]rune{0x0041, 0x0073, 0x0323, 0x0307, 0x0042}},
	{"0041.017F.0307.0323.0042",
		[]rune{0x0041, 0x017F, 0x0307, 0x0323, 0x0042},
		[]rune{0x0041, 0x1E9B, 0x0323, 0x0042},
		[]rune{0x0041, 0x017F, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x1E69, 0x0042},
		[]rune{0x0041, 0x0073, 0x0323, 0x0307, 0x0042}},
	// Composition Exclusions:
	{"0958",
		[]rune{0x958},
		[]rune{0x0915, 0x093C},
		[]rune{0x0915, 0x093C},
		[]rune{0x0915, 0x093C},
		[]rune{0x0915, 0x093C}},
	{"0041.0958.0042",
		[]rune{0x0041, 0x958, 0x0042},
		[]rune{0x0041, 0x0915, 0x093C, 0x0042},
		[]rune{0x0041, 0x0915, 0x093C, 0x0042},
		[]rune{0x0041, 0x0915, 0x093C, 0x0042},
		[]rune{0x0041, 0x0915, 0x093C, 0x0042}},
	// Canoncial Ordering and Combining Cases - Other Examples:
	{"1E0B.0323",
		[]rune{0x1E0B, 0x0323},
		[]rune{0x1E0D, 0x0307},
		[]rune{0x0064, 0x0323, 0x0307},
		[]rune{0x1E0D, 0x0307},
		[]rune{0x0064, 0x0323, 0x0307}},
	{"0045.0344",
		[]rune{0x0045, 0x0344},
		[]rune{0x00CB, 0x0301},
		[]rune{0x0045, 0x0308, 0x0301},
		[]rune{0x00CB, 0x0301},
		[]rune{0x0045, 0x0308, 0x0301}},
	{"0041.1E0B.0323.0042",
		[]rune{0x0041, 0x1E0B, 0x0323, 0x0042},
		[]rune{0x0041, 0x1E0D, 0x0307, 0x0042},
		[]rune{0x0041, 0x0064, 0x0323, 0x0307, 0x0042},
		[]rune{0x0041, 0x1E0D, 0x0307, 0x0042},
		[]rune{0x0041, 0x0064, 0x0323, 0x0307, 0x0042}},
	{"0041.0045.0344.0042",
		[]rune{0x0041, 0x0045, 0x0344, 0x0042},
		[]rune{0x0041, 0x00CB, 0x0301, 0x0042},
		[]rune{0x0041, 0x0045, 0x0308, 0x0301, 0x0042},
		[]rune{0x0041, 0x00CB, 0x0301, 0x0042},
		[]rune{0x0041, 0x0045, 0x0308, 0x0301, 0x0042}},
	// CCC Sorting:
	{"0591.059A", // CCC: 220, 222
		[]rune{0x0591, 0x059A},
		[]rune{0x0591, 0x059A},
		[]rune{0x0591, 0x059A},
		[]rune{0x0591, 0x059A},
		[]rune{0x0591, 0x059A}},
	{"059A.0591", // CCC: 222, 220
		[]rune{0x059A, 0x0591},
		[]rune{0x0591, 0x059A},
		[]rune{0x0591, 0x059A},
		[]rune{0x0591, 0x059A},
		[]rune{0x0591, 0x059A}},
	{"0591.0596", // CCC: 220, 220
		[]rune{0x0591, 0x0596},
		[]rune{0x0591, 0x0596},
		[]rune{0x0591, 0x0596},
		[]rune{0x0591, 0x0596},
		[]rune{0x0591, 0x0596}},
	{"0596.0591", // CCC: 220, 220
		[]rune{0x0596, 0x0591},
		[]rune{0x0596, 0x0591},
		[]rune{0x0596, 0x0591},
		[]rune{0x0596, 0x0591},
		[]rune{0x0596, 0x0591}},
	{"0591.059A.05AE", // CCC: 220, 222, 228
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE}},
	{"0591.05AE.059A", // CCC: 220, 228, 222
		[]rune{0x0591, 0x05AE, 0x059A},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE}},
	{"059A.0591.05AE", // CCC: 222, 220, 228
		[]rune{0x059A, 0x0591, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE}},
	{"059A.05AE.0591", // CCC: 222, 228, 220
		[]rune{0x059A, 0x05AE, 0x0591},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE}},
	{"05AE.0591.059A", // CCC: 228, 220, 222
		[]rune{0x05AE, 0x0591, 0x059A},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE}},
	{"05AE.059A.0591", // CCC: 228, 222, 220
		[]rune{0x05AE, 0x059A, 0x0591},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE},
		[]rune{0x0591, 0x059A, 0x05AE}},
	{"059A.05AD.0592", // CCC: 222, 222, 230
		[]rune{0x059A, 0x05AD, 0x0592},
		[]rune{0x059A, 0x05AD, 0x0592},
		[]rune{0x059A, 0x05AD, 0x0592},
		[]rune{0x059A, 0x05AD, 0x0592},
		[]rune{0x059A, 0x05AD, 0x0592}},
	{"059A.0592.05AD", // CCC: 222, 230, 222
		[]rune{0x059A, 0x0592, 0x05AD},
		[]rune{0x059A, 0x05AD, 0x0592},
		[]rune{0x059A, 0x05AD, 0x0592},
		[]rune{0x059A, 0x05AD, 0x0592},
		[]rune{0x059A, 0x05AD, 0x0592}},
	{"0592.059A.05AD", // CCC: 230, 222, 222
		[]rune{0x0592, 0x059A, 0x05AD},
		[]rune{0x059A, 0x05AD, 0x0592},
		[]rune{0x059A, 0x05AD, 0x0592},
		[]rune{0x059A, 0x05AD, 0x0592},
		[]rune{0x059A, 0x05AD, 0x0592}},
	{"059A.05AD.0591", // CCC: 222, 222, 220
		[]rune{0x059A, 0x05AD, 0x0591},
		[]rune{0x0591, 0x059A, 0x05AD},
		[]rune{0x0591, 0x059A, 0x05AD},
		[]rune{0x0591, 0x059A, 0x05AD},
		[]rune{0x0591, 0x059A, 0x05AD}},
	{"059A.0591.05AD", // CCC: 222, 220, 222
		[]rune{0x059A, 0x0591, 0x05AD},
		[]rune{0x0591, 0x059A, 0x05AD},
		[]rune{0x0591, 0x059A, 0x05AD},
		[]rune{0x0591, 0x059A, 0x05AD},
		[]rune{0x0591, 0x059A, 0x05AD}},
	{"0591.059A.05AD", // CCC: 220, 222, 222
		[]rune{0x0591, 0x059A, 0x05AD},
		[]rune{0x0591, 0x059A, 0x05AD},
		[]rune{0x0591, 0x059A, 0x05AD},
		[]rune{0x0591, 0x059A, 0x05AD},
		[]rune{0x0591, 0x059A, 0x05AD}},
	{"0592.0593.0594", // CCC: 230, 230, 230
		[]rune{0x0592, 0x0593, 0x0594},
		[]rune{0x0592, 0x0593, 0x0594},
		[]rune{0x0592, 0x0593, 0x0594},
		[]rune{0x0592, 0x0593, 0x0594},
		[]rune{0x0592, 0x0593, 0x0594}},
	{"0594.0593.0592", // CCC: 230, 230, 230
		[]rune{0x0594, 0x0593, 0x0592},
		[]rune{0x0594, 0x0593, 0x0592},
		[]rune{0x0594, 0x0593, 0x0592},
		[]rune{0x0594, 0x0593, 0x0592},
		[]rune{0x0594, 0x0593, 0x0592}},
	{"05B9.0591.059A.0592", // CCC: 19, 220, 222, 230
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"05B9.0591.0592.059A", // CCC: 19, 220, 230, 222
		[]rune{0x05B9, 0x0591, 0x0592, 0x059A},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"05B9.059A.0591.0592", // CCC: 19, 222, 220, 230
		[]rune{0x05B9, 0x059A, 0x0591, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"05B9.059A.0592.0591", // CCC: 19, 222, 230, 220
		[]rune{0x05B9, 0x059A, 0x0592, 0x0591},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"05B9.0592.059A.0591", // CCC: 19, 230, 222, 220
		[]rune{0x05B9, 0x0592, 0x059A, 0x0591},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"05B9.0592.0591.059A", // CCC: 19, 230, 220, 222
		[]rune{0x05B9, 0x0592, 0x0591, 0x059A},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"0591.05B9.059A.0592", // CCC: 220, 19, 222, 230
		[]rune{0x0591, 0x05B9, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"0591.05B9.0592.059A", // CCC: 220, 19, 230, 222
		[]rune{0x0591, 0x05B9, 0x0592, 0x059A},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"0591.059A.05B9.0592", // CCC: 220, 222, 19, 230
		[]rune{0x0591, 0x059A, 0x05B9, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"0591.059A.0592.05B9", // CCC: 220, 222, 230, 19
		[]rune{0x0591, 0x059A, 0x0592, 0x05B9},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"0591.0592.05B9.059A", // CCC: 220, 230, 19, 222
		[]rune{0x0591, 0x0592, 0x05B9, 0x059A},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"0591.0592.059A.05B9", // CCC: 220, 230, 222, 19
		[]rune{0x0591, 0x0592, 0x059A, 0x05B9},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"059A.05B9.0591.0592", // CCC: 222, 19, 220, 230
		[]rune{0x059A, 0x05B9, 0x0591, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"059A.05B9.0592.0591", // CCC: 222, 19, 230, 220
		[]rune{0x059A, 0x05B9, 0x0592, 0x0591},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"059A.0591.05B9.0592", // CCC: 222, 220, 19, 230
		[]rune{0x059A, 0x0591, 0x05B9, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"059A.0591.0592.05B9", // CCC: 222, 220, 230, 19
		[]rune{0x059A, 0x0591, 0x0592, 0x05B9},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"059A.0592.05B9.0591", // CCC: 222, 230, 19, 220
		[]rune{0x059A, 0x0592, 0x05B9, 0x0591},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"059A.0592.0591.05B9", // CCC: 222, 230, 220, 19
		[]rune{0x059A, 0x0592, 0x0591, 0x05B9},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"0592.05B9.0591.059A", // CCC: 230, 19, 220, 222
		[]rune{0x0592, 0x05B9, 0x0591, 0x059A},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"0592.05B9.059A.0591", // CCC: 230, 19, 222, 220
		[]rune{0x0592, 0x05B9, 0x059A, 0x0591},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"0592.0591.05B9.059A", // CCC: 230, 220, 19, 222
		[]rune{0x0592, 0x0591, 0x05B9, 0x059A},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"0592.0591.059A.05B9", // CCC: 230, 220, 222, 19
		[]rune{0x0592, 0x0591, 0x059A, 0x05B9},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"0592.059A.05B9.0591", // CCC: 230, 222, 19, 220
		[]rune{0x0592, 0x059A, 0x05B9, 0x0591},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"0592.059A.0591.05B9", // CCC: 230, 222, 220, 19
		[]rune{0x0592, 0x059A, 0x0591, 0x05B9},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592},
		[]rune{0x05B9, 0x0591, 0x059A, 0x0592}},
	{"0596.0591.059A.0592", // CCC: 220, 220, 222, 230
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592}},
	{"0596.0591.0592.059A", // CCC: 220, 220, 230, 222
		[]rune{0x0596, 0x0591, 0x0592, 0x059A},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592}},
	{"0596.059A.0591.0592", // CCC: 220, 222, 220, 230
		[]rune{0x0596, 0x059A, 0x0591, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592}},
	{"0596.0592.0591.059A", // CCC: 220, 230, 220, 222
		[]rune{0x0596, 0x0592, 0x0591, 0x059A},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592}},
	{"0596.059A.0592.0591", // CCC: 220, 222, 230, 220
		[]rune{0x0596, 0x059A, 0x0592, 0x0591},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592}},
	{"0596.0592.059A.0591", // CCC: 220, 230, 222, 220
		[]rune{0x0596, 0x0592, 0x059A, 0x0591},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592}},
	{"059A.0596.0591.0592", // CCC: 222, 220, 220, 230
		[]rune{0x059A, 0x0596, 0x0591, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592}},
	{"0592.0596.0591.059A", // CCC: 230, 220, 220, 222
		[]rune{0x0592, 0x0596, 0x0591, 0x059A},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592}},
	{"059A.0596.0592.0591", // CCC: 222, 220, 230, 220
		[]rune{0x059A, 0x0596, 0x0592, 0x0591},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592}},
	{"0592.0596.059A.0591", // CCC: 230, 220, 222, 220
		[]rune{0x0592, 0x0596, 0x059A, 0x0591},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592}},
	{"059A.0592.0596.0591", // CCC: 222, 230, 220, 220
		[]rune{0x059A, 0x0592, 0x0596, 0x0591},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592}},
	{"0592.059A.0596.0591", // CCC: 230, 222, 220, 220
		[]rune{0x0592, 0x059A, 0x0596, 0x0591},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592},
		[]rune{0x0596, 0x0591, 0x059A, 0x0592}},
	{"05B9.0591.0596.0592", // CCC: 19, 220, 220, 230
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592}},
	{"0592.0591.0596.05B9", // CCC: 230, 220, 220, 19
		[]rune{0x0592, 0x0591, 0x0596, 0x05B9},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592}},
	{"0591.0596.05B9.0592", // CCC: 220, 220, 19, 230
		[]rune{0x0591, 0x0596, 0x05B9, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592}},
	{"0591.0596.0592.05B9", // CCC: 220, 220, 230, 19
		[]rune{0x0591, 0x0596, 0x0592, 0x05B9},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592}},
	{"0591.05B9.0596.0592", // CCC: 220, 19, 220, 230
		[]rune{0x0591, 0x05B9, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592}},
	{"0591.0592.0596.05B9", // CCC: 220, 230, 220, 19
		[]rune{0x0591, 0x0592, 0x0596, 0x05B9},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592}},
	{"0591.05B9.0592.0596", // CCC: 220, 19, 230, 220
		[]rune{0x0591, 0x05B9, 0x0592, 0x0596},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592}},
	{"0591.0592.05B9.0596", // CCC: 220, 230, 19, 220
		[]rune{0x0591, 0x0592, 0x05B9, 0x0596},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592}},
	{"05B9.0591.0592.0596", // CCC: 19, 220, 230, 220
		[]rune{0x05B9, 0x0591, 0x0592, 0x0596},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592}},
	{"0592.0591.05B9.0596", // CCC: 230, 220, 19, 220
		[]rune{0x0592, 0x0591, 0x05B9, 0x0596},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592}},
	{"05B9.0592.0591.0596", // CCC: 19, 230, 220, 220
		[]rune{0x05B9, 0x0592, 0x0591, 0x0596},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592}},
	{"0592.05B9.0591.0596", // CCC: 230, 19, 220, 220
		[]rune{0x0592, 0x05B9, 0x0591, 0x0596},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0596, 0x0592}},
	{"05B9.0591.0593.0592", // CCC: 19, 220, 230, 230
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592}},
	{"0591.05B9.0593.0592", // CCC: 220, 19, 230, 230
		[]rune{0x0591, 0x05B9, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592}},
	{"0593.0592.05B9.0591", // CCC: 230, 230, 19, 220
		[]rune{0x0593, 0x0592, 0x05B9, 0x0591},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592}},
	{"0593.0592.0591.05B9", // CCC: 230, 230, 220, 19
		[]rune{0x0593, 0x0592, 0x0591, 0x05B9},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592}},
	{"0593.05B9.0592.0591", // CCC: 230, 19, 230, 220
		[]rune{0x0593, 0x05B9, 0x0592, 0x0591},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592}},
	{"0593.0591.0592.05B9", // CCC: 230, 220, 230, 19
		[]rune{0x0593, 0x0591, 0x0592, 0x05B9},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592}},
	{"0593.05B9.0591.0592", // CCC: 230, 19, 220, 230
		[]rune{0x0593, 0x05B9, 0x0591, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592}},
	{"0593.0591.05B9.0592", // CCC: 230, 220, 19, 230
		[]rune{0x0593, 0x0591, 0x05B9, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592}},
	{"05B9.0593.0592.0591", // CCC: 19, 230, 230, 220
		[]rune{0x05B9, 0x0593, 0x0592, 0x0591},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592}},
	{"0591.0593.0592.05B9", // CCC: 220, 230, 230, 19
		[]rune{0x0591, 0x0593, 0x0592, 0x05B9},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592}},
	{"05B9.0593.0591.0592", // CCC: 19, 230, 220, 230
		[]rune{0x05B9, 0x0593, 0x0591, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592}},
	{"0591.0593.05B9.0592", // CCC: 220, 230, 19, 230
		[]rune{0x0591, 0x0593, 0x05B9, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592},
		[]rune{0x05B9, 0x0591, 0x0593, 0x0592}},
	{"0592.0593.0594.0595", // CCC: 230, 230, 230, 230
		[]rune{0x0592, 0x0593, 0x0594, 0x0595},
		[]rune{0x0592, 0x0593, 0x0594, 0x0595},
		[]rune{0x0592, 0x0593, 0x0594, 0x0595},
		[]rune{0x0592, 0x0593, 0x0594, 0x0595},
		[]rune{0x0592, 0x0593, 0x0594, 0x0595}},
	// Wierd Cases - Large Compatibility Decompositions:
	{"321D",
		[]rune{0x321D},
		[]rune{0x321D},
		[]rune{0x321D},
		[]rune{0x0028, 0xC624, 0xC804, 0x0029},
		[]rune{0x0028, 0x110B, 0x1169, 0x110C, 0x1165, 0x11AB, 0x0029}},
	{"FDFB",
		[]rune{0xFDFB},
		[]rune{0xFDFB},
		[]rune{0xFDFB},
		[]rune{0x062C, 0x0644, 0x0020, 0x062C, 0x0644, 0x0627, 0x0644, 0x0647},
		[]rune{0x062C, 0x0644, 0x0020, 0x062C, 0x0644, 0x0627, 0x0644, 0x0647}},
	{"FDFA",
		[]rune{0xFDFA},
		[]rune{0xFDFA},
		[]rune{0xFDFA},
		[]rune{0x0635, 0x0644, 0x0649, 0x0020, 0x0627, 0x0644,
			0x0644, 0x0647, 0x0020, 0x0639, 0x0644, 0x064A,
			0x0647, 0x0020, 0x0648, 0x0633, 0x0644, 0x0645},
		[]rune{0x0635, 0x0644, 0x0649, 0x0020, 0x0627, 0x0644,
			0x0644, 0x0647, 0x0020, 0x0639, 0x0644, 0x064A,
			0x0647, 0x0020, 0x0648, 0x0633, 0x0644, 0x0645}},
	{"0041.321D.0042",
		[]rune{0x0041, 0x321D, 0x0042},
		[]rune{0x0041, 0x321D, 0x0042},
		[]rune{0x0041, 0x321D, 0x0042},
		[]rune{0x0041, 0x0028, 0xC624, 0xC804, 0x0029, 0x0042},
		[]rune{0x0041, 0x0028, 0x110B, 0x1169, 0x110C, 0x1165, 0x11AB, 0x0029, 0x0042}},
	{"0041.FDFB.0042",
		[]rune{0x0041, 0xFDFB, 0x0042},
		[]rune{0x0041, 0xFDFB, 0x0042},
		[]rune{0x0041, 0xFDFB, 0x0042},
		[]rune{0x0041, 0x062C, 0x0644, 0x0020, 0x062C, 0x0644, 0x0627, 0x0644, 0x0647, 0x0042},
		[]rune{0x0041, 0x062C, 0x0644, 0x0020, 0x062C, 0x0644, 0x0627, 0x0644, 0x0647, 0x0042}},
	{"0041.FDFA.0042",
		[]rune{0x0041, 0xFDFA, 0x0042},
		[]rune{0x0041, 0xFDFA, 0x0042},
		[]rune{0x0041, 0xFDFA, 0x0042},
		[]rune{0x0041, 0x0635, 0x0644, 0x0649, 0x0020, 0x0627, 0x0644,
			0x0644, 0x0647, 0x0020, 0x0639, 0x0644, 0x064A,
			0x0647, 0x0020, 0x0648, 0x0633, 0x0644, 0x0645, 0x0042},
		[]rune{0x0041, 0x0635, 0x0644, 0x0649, 0x0020, 0x0627, 0x0644,
			0x0644, 0x0647, 0x0020, 0x0639, 0x0644, 0x064A,
			0x0647, 0x0020, 0x0648, 0x0633, 0x0644, 0x0645, 0x0042}},
}

// TestAutogenNormCases walks through autogenNormTestCases and tests
// the various normalizations.
func TestNormCases(t *testing.T) {
	for i := range normTestCases {
		tc := normTestCases[i]
		testNormCase(t, tc.label, "NFC", tc.data, tc.nfc, unorm.NextNFCLen)
		testNormCase(t, tc.label, "NFD", tc.data, tc.nfd, unorm.NextNFDLen)
		testNormCase(t, tc.label, "NFKC", tc.data, tc.nfkc, unorm.NextNFKCLen)
		testNormCase(t, tc.label, "NFKD", tc.data, tc.nfkd, unorm.NextNFKDLen)
		testNormCaseReader(t, tc.label, "NFC", tc.data, tc.nfc, unorm.NFCReader)
		testNormCaseReader(t, tc.label, "NFD", tc.data, tc.nfd, unorm.NFDReader)
		testNormCaseReader(t, tc.label, "NFKC", tc.data, tc.nfkc, unorm.NFKCReader)
		testNormCaseReader(t, tc.label, "NFKD", tc.data, tc.nfkd, unorm.NFKDReader)
	}
	//	testNormCaseReader(t, normTestCases[0].label, "NFC", normTestCases[0].data, normTestCases[0].nfc, unorm.NFCReader)
}

func testNormCase(t *testing.T, label, nType string, dataIn, dataOut []rune, testFunc func([]rune, int, int, bool) ([]rune, int)) {
	ofsIn := 0
	ofsOut := 0
	dataInLen := len(dataIn)
	dataOutLen := len(dataOut)
	for {
		r1, d1 := testFunc(dataIn, ofsIn, dataInLen, false)
		r2, d2 := testFunc(dataIn, ofsIn, dataInLen, true)
		if d1 != d2 && ofsIn+d1 < dataInLen {
			t.Errorf("\"%s\"(\"%s\",%d,%d): delta abandon = %d != delta %d", nType, label, ofsIn, dataInLen, d2, d1)
		}
		if d1 == 0 {
			if len(r1) == 0 && len(r2) == 0 && ofsIn == dataInLen && ofsOut == dataOutLen {
				return
			}
			t.Errorf("\"%s\"(\"%s\",%d,%d): unexpected results at end", nType, label, ofsIn, dataInLen)
			return
		}
		if len(r1) == 0 {
			t.Errorf("\"%s\"(\"%s\",%d,%d): len(r1)=0, want >0", nType, label, ofsIn, dataInLen)
			return
		}
		if ofsIn+d2 >= dataInLen {
			if len(r2) != 0 {
				t.Errorf("\"%s\"(\"%s\",%d,%d): len(r2)=%d, want 0", nType, label, ofsIn, dataInLen, len(r2))
				return
			}
		} else {
			if len(r2) == 0 {
				t.Errorf("\"%s\"(\"%s\",%d,%d): len(r2)=0, want >0", nType, label, ofsIn, dataInLen)
				return
			}
		}
		if len(r1) == 0 {
			t.Errorf("\"%s\"(\"%s\",%d,%d): len(r1)=0, want >0", nType, label, ofsIn, dataInLen)
			return
		}
		if ofsIn+d1 > dataInLen {
			t.Errorf("\"%s\"(\"%s\",%d,%d): ofs+d1=%d+%d=%d, want >%d: %U", nType, label, ofsIn, dataInLen, ofsIn, d1, ofsIn+d1, dataInLen, r1)
			return
		}
		if ofsOut+len(r1) > dataOutLen {
			t.Errorf("\"%s\"(\"%s\",%d,%d): ofs+len(r1)=%d+%d=%d, want <=%d: %U", nType, label, ofsIn, dataInLen, ofsOut, len(r1), ofsOut+len(r1), dataOutLen, r1)
			return
		}
		for i := range r1 {
			if r1[i] != dataOut[ofsOut+i] {
				t.Errorf("\"%s\"(\"%s\",%d,%d): r1[%d]=%02X, want %02X", nType, label, ofsIn, dataInLen, i, r1[i], dataOut[ofsOut+i])
				return
			}
		}
		ofsIn += d1
		ofsOut += len(r1)
	}
}

func testNormCaseReader(t *testing.T, label, nType string, dataIn, dataOut []rune, testFunc func(io.RuneReader) io.RuneReader) {
	dataOutLen := len(dataOut)
	rr := testFunc(strings.NewReader(string(dataIn)))
	for i := 0; ; i++ {
		r, _, err := rr.ReadRune()
		if i >= dataOutLen {
			if err != io.EOF {
				t.Errorf("\"%s\"(\"%s\") %d: expected io.EOF", nType, label, i)
			}
			break
		}
		if err != nil {
			t.Errorf("\"%s\"(\"%s\") %d: unexpected error: %s", nType, label, i, err)
			return
		}
		if r != dataOut[i] {
			t.Errorf("\"%s\"(\"%s\") %d: = %04X, want %04X", nType, label, i, r, dataOut[i])
		}
	}
}
