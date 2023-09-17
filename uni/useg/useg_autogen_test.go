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

package useg_test

import "testing"
import "github.com/orkvozku/go/uni/useg"

func autogenTestCaseFuncGcbCn(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValueCn
}

func autogenTestCaseFuncGcbCr(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValueCr
}

func autogenTestCaseFuncGcbEx(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValueEx
}

func autogenTestCaseFuncGcbL(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValueL
}

func autogenTestCaseFuncGcbLf(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValueLf
}

func autogenTestCaseFuncGcbLv(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValueLv
}

func autogenTestCaseFuncGcbLvt(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValueLvt
}

func autogenTestCaseFuncGcbPp(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValuePp
}

func autogenTestCaseFuncGcbRi(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValueRi
}

func autogenTestCaseFuncGcbSm(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValueSm
}

func autogenTestCaseFuncGcbT(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValueT
}

func autogenTestCaseFuncGcbV(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValueV
}

func autogenTestCaseFuncGcbXx(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValueXx
}

func autogenTestCaseFuncGcbZwj(r rune) bool {
	return useg.ValueOfGcb(r, -2147483648) == useg.GcbValueZwj
}

func autogenTestCaseFuncIncbConsonant(r rune) bool {
	return useg.ValueOfIncb(r, -2147483648) == useg.IncbValueConsonant
}

func autogenTestCaseFuncIncbExtend(r rune) bool {
	return useg.ValueOfIncb(r, -2147483648) == useg.IncbValueExtend
}

func autogenTestCaseFuncIncbLinker(r rune) bool {
	return useg.ValueOfIncb(r, -2147483648) == useg.IncbValueLinker
}

func autogenTestCaseFuncIncbNone(r rune) bool {
	return useg.ValueOfIncb(r, -2147483648) == useg.IncbValueNone
}

func autogenTestCaseFuncLbAi(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueAi
}

func autogenTestCaseFuncLbAk(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueAk
}

func autogenTestCaseFuncLbAl(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueAl
}

func autogenTestCaseFuncLbAp(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueAp
}

func autogenTestCaseFuncLbAs(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueAs
}

func autogenTestCaseFuncLbB2(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueB2
}

func autogenTestCaseFuncLbBa(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueBa
}

func autogenTestCaseFuncLbBb(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueBb
}

func autogenTestCaseFuncLbBk(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueBk
}

func autogenTestCaseFuncLbCb(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueCb
}

func autogenTestCaseFuncLbCj(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueCj
}

func autogenTestCaseFuncLbCl(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueCl
}

func autogenTestCaseFuncLbCm(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueCm
}

func autogenTestCaseFuncLbCp(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueCp
}

func autogenTestCaseFuncLbCr(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueCr
}

func autogenTestCaseFuncLbEb(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueEb
}

func autogenTestCaseFuncLbEm(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueEm
}

func autogenTestCaseFuncLbEx(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueEx
}

func autogenTestCaseFuncLbGl(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueGl
}

func autogenTestCaseFuncLbH2(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueH2
}

func autogenTestCaseFuncLbH3(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueH3
}

func autogenTestCaseFuncLbHl(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueHl
}

func autogenTestCaseFuncLbHy(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueHy
}

func autogenTestCaseFuncLbId(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueId
}

func autogenTestCaseFuncLbIn(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueIn
}

func autogenTestCaseFuncLbIs(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueIs
}

func autogenTestCaseFuncLbJl(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueJl
}

func autogenTestCaseFuncLbJt(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueJt
}

func autogenTestCaseFuncLbJv(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueJv
}

func autogenTestCaseFuncLbLf(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueLf
}

func autogenTestCaseFuncLbNl(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueNl
}

func autogenTestCaseFuncLbNs(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueNs
}

func autogenTestCaseFuncLbNu(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueNu
}

func autogenTestCaseFuncLbOp(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueOp
}

func autogenTestCaseFuncLbPo(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValuePo
}

func autogenTestCaseFuncLbPr(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValuePr
}

func autogenTestCaseFuncLbQu(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueQu
}

func autogenTestCaseFuncLbRi(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueRi
}

func autogenTestCaseFuncLbSa(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueSa
}

func autogenTestCaseFuncLbSg(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueSg
}

func autogenTestCaseFuncLbSp(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueSp
}

func autogenTestCaseFuncLbSy(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueSy
}

func autogenTestCaseFuncLbVf(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueVf
}

func autogenTestCaseFuncLbVi(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueVi
}

func autogenTestCaseFuncLbWj(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueWj
}

func autogenTestCaseFuncLbXx(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueXx
}

func autogenTestCaseFuncLbZw(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueZw
}

func autogenTestCaseFuncLbZwj(r rune) bool {
	return useg.ValueOfLb(r, -2147483648) == useg.LbValueZwj
}

func autogenTestCaseFuncSbAt(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueAt
}

func autogenTestCaseFuncSbCl(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueCl
}

func autogenTestCaseFuncSbCr(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueCr
}

func autogenTestCaseFuncSbEx(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueEx
}

func autogenTestCaseFuncSbFo(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueFo
}

func autogenTestCaseFuncSbLe(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueLe
}

func autogenTestCaseFuncSbLf(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueLf
}

func autogenTestCaseFuncSbLo(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueLo
}

func autogenTestCaseFuncSbNu(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueNu
}

func autogenTestCaseFuncSbSc(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueSc
}

func autogenTestCaseFuncSbSe(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueSe
}

func autogenTestCaseFuncSbSp(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueSp
}

func autogenTestCaseFuncSbSt(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueSt
}

func autogenTestCaseFuncSbUp(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueUp
}

func autogenTestCaseFuncSbXx(r rune) bool {
	return useg.ValueOfSb(r, -2147483648) == useg.SbValueXx
}

func autogenTestCaseFuncWbCr(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueCr
}

func autogenTestCaseFuncWbDq(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueDq
}

func autogenTestCaseFuncWbEx(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueEx
}

func autogenTestCaseFuncWbExtend(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueExtend
}

func autogenTestCaseFuncWbFo(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueFo
}

func autogenTestCaseFuncWbHl(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueHl
}

func autogenTestCaseFuncWbKa(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueKa
}

func autogenTestCaseFuncWbLe(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueLe
}

func autogenTestCaseFuncWbLf(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueLf
}

func autogenTestCaseFuncWbMb(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueMb
}

func autogenTestCaseFuncWbMl(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueMl
}

func autogenTestCaseFuncWbMn(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueMn
}

func autogenTestCaseFuncWbNl(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueNl
}

func autogenTestCaseFuncWbNu(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueNu
}

func autogenTestCaseFuncWbRi(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueRi
}

func autogenTestCaseFuncWbSq(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueSq
}

func autogenTestCaseFuncWbWsegspace(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueWsegspace
}

func autogenTestCaseFuncWbXx(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueXx
}

func autogenTestCaseFuncWbZwj(r rune) bool {
	return useg.ValueOfWb(r, -2147483648) == useg.WbValueZwj
}

// autogenTestCaseFunc contains a function to test.
type autogenTestCaseFunc struct {
	label string
	f     func(rune) bool
}

// autogenTestCasesFunc... describes a function to test.
var (
	autogenTestCasesFuncExtpictN      = autogenTestCaseFunc{"HasExtpictN", useg.HasExtpictN}
	autogenTestCasesFuncExtpictY      = autogenTestCaseFunc{"HasExtpictY", useg.HasExtpictY}
	autogenTestCasesFuncGcbCn         = autogenTestCaseFunc{"ValueOfGcb=Cn", autogenTestCaseFuncGcbCn}
	autogenTestCasesFuncGcbCr         = autogenTestCaseFunc{"ValueOfGcb=Cr", autogenTestCaseFuncGcbCr}
	autogenTestCasesFuncGcbEx         = autogenTestCaseFunc{"ValueOfGcb=Ex", autogenTestCaseFuncGcbEx}
	autogenTestCasesFuncGcbL          = autogenTestCaseFunc{"ValueOfGcb=L", autogenTestCaseFuncGcbL}
	autogenTestCasesFuncGcbLf         = autogenTestCaseFunc{"ValueOfGcb=Lf", autogenTestCaseFuncGcbLf}
	autogenTestCasesFuncGcbLv         = autogenTestCaseFunc{"ValueOfGcb=Lv", autogenTestCaseFuncGcbLv}
	autogenTestCasesFuncGcbLvt        = autogenTestCaseFunc{"ValueOfGcb=Lvt", autogenTestCaseFuncGcbLvt}
	autogenTestCasesFuncGcbPp         = autogenTestCaseFunc{"ValueOfGcb=Pp", autogenTestCaseFuncGcbPp}
	autogenTestCasesFuncGcbRi         = autogenTestCaseFunc{"ValueOfGcb=Ri", autogenTestCaseFuncGcbRi}
	autogenTestCasesFuncGcbSm         = autogenTestCaseFunc{"ValueOfGcb=Sm", autogenTestCaseFuncGcbSm}
	autogenTestCasesFuncGcbT          = autogenTestCaseFunc{"ValueOfGcb=T", autogenTestCaseFuncGcbT}
	autogenTestCasesFuncGcbV          = autogenTestCaseFunc{"ValueOfGcb=V", autogenTestCaseFuncGcbV}
	autogenTestCasesFuncGcbXx         = autogenTestCaseFunc{"ValueOfGcb=Xx", autogenTestCaseFuncGcbXx}
	autogenTestCasesFuncGcbZwj        = autogenTestCaseFunc{"ValueOfGcb=Zwj", autogenTestCaseFuncGcbZwj}
	autogenTestCasesFuncIncbConsonant = autogenTestCaseFunc{"ValueOfIncb=Consonant", autogenTestCaseFuncIncbConsonant}
	autogenTestCasesFuncIncbExtend    = autogenTestCaseFunc{"ValueOfIncb=Extend", autogenTestCaseFuncIncbExtend}
	autogenTestCasesFuncIncbLinker    = autogenTestCaseFunc{"ValueOfIncb=Linker", autogenTestCaseFuncIncbLinker}
	autogenTestCasesFuncIncbNone      = autogenTestCaseFunc{"ValueOfIncb=None", autogenTestCaseFuncIncbNone}
	autogenTestCasesFuncLbAi          = autogenTestCaseFunc{"ValueOfLb=Ai", autogenTestCaseFuncLbAi}
	autogenTestCasesFuncLbAk          = autogenTestCaseFunc{"ValueOfLb=Ak", autogenTestCaseFuncLbAk}
	autogenTestCasesFuncLbAl          = autogenTestCaseFunc{"ValueOfLb=Al", autogenTestCaseFuncLbAl}
	autogenTestCasesFuncLbAp          = autogenTestCaseFunc{"ValueOfLb=Ap", autogenTestCaseFuncLbAp}
	autogenTestCasesFuncLbAs          = autogenTestCaseFunc{"ValueOfLb=As", autogenTestCaseFuncLbAs}
	autogenTestCasesFuncLbB2          = autogenTestCaseFunc{"ValueOfLb=B2", autogenTestCaseFuncLbB2}
	autogenTestCasesFuncLbBa          = autogenTestCaseFunc{"ValueOfLb=Ba", autogenTestCaseFuncLbBa}
	autogenTestCasesFuncLbBb          = autogenTestCaseFunc{"ValueOfLb=Bb", autogenTestCaseFuncLbBb}
	autogenTestCasesFuncLbBk          = autogenTestCaseFunc{"ValueOfLb=Bk", autogenTestCaseFuncLbBk}
	autogenTestCasesFuncLbCb          = autogenTestCaseFunc{"ValueOfLb=Cb", autogenTestCaseFuncLbCb}
	autogenTestCasesFuncLbCj          = autogenTestCaseFunc{"ValueOfLb=Cj", autogenTestCaseFuncLbCj}
	autogenTestCasesFuncLbCl          = autogenTestCaseFunc{"ValueOfLb=Cl", autogenTestCaseFuncLbCl}
	autogenTestCasesFuncLbCm          = autogenTestCaseFunc{"ValueOfLb=Cm", autogenTestCaseFuncLbCm}
	autogenTestCasesFuncLbCp          = autogenTestCaseFunc{"ValueOfLb=Cp", autogenTestCaseFuncLbCp}
	autogenTestCasesFuncLbCr          = autogenTestCaseFunc{"ValueOfLb=Cr", autogenTestCaseFuncLbCr}
	autogenTestCasesFuncLbEb          = autogenTestCaseFunc{"ValueOfLb=Eb", autogenTestCaseFuncLbEb}
	autogenTestCasesFuncLbEm          = autogenTestCaseFunc{"ValueOfLb=Em", autogenTestCaseFuncLbEm}
	autogenTestCasesFuncLbEx          = autogenTestCaseFunc{"ValueOfLb=Ex", autogenTestCaseFuncLbEx}
	autogenTestCasesFuncLbGl          = autogenTestCaseFunc{"ValueOfLb=Gl", autogenTestCaseFuncLbGl}
	autogenTestCasesFuncLbH2          = autogenTestCaseFunc{"ValueOfLb=H2", autogenTestCaseFuncLbH2}
	autogenTestCasesFuncLbH3          = autogenTestCaseFunc{"ValueOfLb=H3", autogenTestCaseFuncLbH3}
	autogenTestCasesFuncLbHl          = autogenTestCaseFunc{"ValueOfLb=Hl", autogenTestCaseFuncLbHl}
	autogenTestCasesFuncLbHy          = autogenTestCaseFunc{"ValueOfLb=Hy", autogenTestCaseFuncLbHy}
	autogenTestCasesFuncLbId          = autogenTestCaseFunc{"ValueOfLb=Id", autogenTestCaseFuncLbId}
	autogenTestCasesFuncLbIn          = autogenTestCaseFunc{"ValueOfLb=In", autogenTestCaseFuncLbIn}
	autogenTestCasesFuncLbIs          = autogenTestCaseFunc{"ValueOfLb=Is", autogenTestCaseFuncLbIs}
	autogenTestCasesFuncLbJl          = autogenTestCaseFunc{"ValueOfLb=Jl", autogenTestCaseFuncLbJl}
	autogenTestCasesFuncLbJt          = autogenTestCaseFunc{"ValueOfLb=Jt", autogenTestCaseFuncLbJt}
	autogenTestCasesFuncLbJv          = autogenTestCaseFunc{"ValueOfLb=Jv", autogenTestCaseFuncLbJv}
	autogenTestCasesFuncLbLf          = autogenTestCaseFunc{"ValueOfLb=Lf", autogenTestCaseFuncLbLf}
	autogenTestCasesFuncLbNl          = autogenTestCaseFunc{"ValueOfLb=Nl", autogenTestCaseFuncLbNl}
	autogenTestCasesFuncLbNs          = autogenTestCaseFunc{"ValueOfLb=Ns", autogenTestCaseFuncLbNs}
	autogenTestCasesFuncLbNu          = autogenTestCaseFunc{"ValueOfLb=Nu", autogenTestCaseFuncLbNu}
	autogenTestCasesFuncLbOp          = autogenTestCaseFunc{"ValueOfLb=Op", autogenTestCaseFuncLbOp}
	autogenTestCasesFuncLbPo          = autogenTestCaseFunc{"ValueOfLb=Po", autogenTestCaseFuncLbPo}
	autogenTestCasesFuncLbPr          = autogenTestCaseFunc{"ValueOfLb=Pr", autogenTestCaseFuncLbPr}
	autogenTestCasesFuncLbQu          = autogenTestCaseFunc{"ValueOfLb=Qu", autogenTestCaseFuncLbQu}
	autogenTestCasesFuncLbRi          = autogenTestCaseFunc{"ValueOfLb=Ri", autogenTestCaseFuncLbRi}
	autogenTestCasesFuncLbSa          = autogenTestCaseFunc{"ValueOfLb=Sa", autogenTestCaseFuncLbSa}
	autogenTestCasesFuncLbSg          = autogenTestCaseFunc{"ValueOfLb=Sg", autogenTestCaseFuncLbSg}
	autogenTestCasesFuncLbSp          = autogenTestCaseFunc{"ValueOfLb=Sp", autogenTestCaseFuncLbSp}
	autogenTestCasesFuncLbSy          = autogenTestCaseFunc{"ValueOfLb=Sy", autogenTestCaseFuncLbSy}
	autogenTestCasesFuncLbVf          = autogenTestCaseFunc{"ValueOfLb=Vf", autogenTestCaseFuncLbVf}
	autogenTestCasesFuncLbVi          = autogenTestCaseFunc{"ValueOfLb=Vi", autogenTestCaseFuncLbVi}
	autogenTestCasesFuncLbWj          = autogenTestCaseFunc{"ValueOfLb=Wj", autogenTestCaseFuncLbWj}
	autogenTestCasesFuncLbXx          = autogenTestCaseFunc{"ValueOfLb=Xx", autogenTestCaseFuncLbXx}
	autogenTestCasesFuncLbZw          = autogenTestCaseFunc{"ValueOfLb=Zw", autogenTestCaseFuncLbZw}
	autogenTestCasesFuncLbZwj         = autogenTestCaseFunc{"ValueOfLb=Zwj", autogenTestCaseFuncLbZwj}
	autogenTestCasesFuncSbAt          = autogenTestCaseFunc{"ValueOfSb=At", autogenTestCaseFuncSbAt}
	autogenTestCasesFuncSbCl          = autogenTestCaseFunc{"ValueOfSb=Cl", autogenTestCaseFuncSbCl}
	autogenTestCasesFuncSbCr          = autogenTestCaseFunc{"ValueOfSb=Cr", autogenTestCaseFuncSbCr}
	autogenTestCasesFuncSbEx          = autogenTestCaseFunc{"ValueOfSb=Ex", autogenTestCaseFuncSbEx}
	autogenTestCasesFuncSbFo          = autogenTestCaseFunc{"ValueOfSb=Fo", autogenTestCaseFuncSbFo}
	autogenTestCasesFuncSbLe          = autogenTestCaseFunc{"ValueOfSb=Le", autogenTestCaseFuncSbLe}
	autogenTestCasesFuncSbLf          = autogenTestCaseFunc{"ValueOfSb=Lf", autogenTestCaseFuncSbLf}
	autogenTestCasesFuncSbLo          = autogenTestCaseFunc{"ValueOfSb=Lo", autogenTestCaseFuncSbLo}
	autogenTestCasesFuncSbNu          = autogenTestCaseFunc{"ValueOfSb=Nu", autogenTestCaseFuncSbNu}
	autogenTestCasesFuncSbSc          = autogenTestCaseFunc{"ValueOfSb=Sc", autogenTestCaseFuncSbSc}
	autogenTestCasesFuncSbSe          = autogenTestCaseFunc{"ValueOfSb=Se", autogenTestCaseFuncSbSe}
	autogenTestCasesFuncSbSp          = autogenTestCaseFunc{"ValueOfSb=Sp", autogenTestCaseFuncSbSp}
	autogenTestCasesFuncSbSt          = autogenTestCaseFunc{"ValueOfSb=St", autogenTestCaseFuncSbSt}
	autogenTestCasesFuncSbUp          = autogenTestCaseFunc{"ValueOfSb=Up", autogenTestCaseFuncSbUp}
	autogenTestCasesFuncSbXx          = autogenTestCaseFunc{"ValueOfSb=Xx", autogenTestCaseFuncSbXx}
	autogenTestCasesFuncWbCr          = autogenTestCaseFunc{"ValueOfWb=Cr", autogenTestCaseFuncWbCr}
	autogenTestCasesFuncWbDq          = autogenTestCaseFunc{"ValueOfWb=Dq", autogenTestCaseFuncWbDq}
	autogenTestCasesFuncWbEx          = autogenTestCaseFunc{"ValueOfWb=Ex", autogenTestCaseFuncWbEx}
	autogenTestCasesFuncWbExtend      = autogenTestCaseFunc{"ValueOfWb=Extend", autogenTestCaseFuncWbExtend}
	autogenTestCasesFuncWbFo          = autogenTestCaseFunc{"ValueOfWb=Fo", autogenTestCaseFuncWbFo}
	autogenTestCasesFuncWbHl          = autogenTestCaseFunc{"ValueOfWb=Hl", autogenTestCaseFuncWbHl}
	autogenTestCasesFuncWbKa          = autogenTestCaseFunc{"ValueOfWb=Ka", autogenTestCaseFuncWbKa}
	autogenTestCasesFuncWbLe          = autogenTestCaseFunc{"ValueOfWb=Le", autogenTestCaseFuncWbLe}
	autogenTestCasesFuncWbLf          = autogenTestCaseFunc{"ValueOfWb=Lf", autogenTestCaseFuncWbLf}
	autogenTestCasesFuncWbMb          = autogenTestCaseFunc{"ValueOfWb=Mb", autogenTestCaseFuncWbMb}
	autogenTestCasesFuncWbMl          = autogenTestCaseFunc{"ValueOfWb=Ml", autogenTestCaseFuncWbMl}
	autogenTestCasesFuncWbMn          = autogenTestCaseFunc{"ValueOfWb=Mn", autogenTestCaseFuncWbMn}
	autogenTestCasesFuncWbNl          = autogenTestCaseFunc{"ValueOfWb=Nl", autogenTestCaseFuncWbNl}
	autogenTestCasesFuncWbNu          = autogenTestCaseFunc{"ValueOfWb=Nu", autogenTestCaseFuncWbNu}
	autogenTestCasesFuncWbRi          = autogenTestCaseFunc{"ValueOfWb=Ri", autogenTestCaseFuncWbRi}
	autogenTestCasesFuncWbSq          = autogenTestCaseFunc{"ValueOfWb=Sq", autogenTestCaseFuncWbSq}
	autogenTestCasesFuncWbWsegspace   = autogenTestCaseFunc{"ValueOfWb=Wsegspace", autogenTestCaseFuncWbWsegspace}
	autogenTestCasesFuncWbXx          = autogenTestCaseFunc{"ValueOfWb=Xx", autogenTestCaseFuncWbXx}
	autogenTestCasesFuncWbZwj         = autogenTestCaseFunc{"ValueOfWb=Zwj", autogenTestCaseFuncWbZwj}
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
	{ // Incb
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbConsonant,
			&autogenTestCasesFuncIncbExtend,
			&autogenTestCasesFuncIncbLinker,
			&autogenTestCasesFuncIncbNone,
		}},
		{0x00, 0x2FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x300, 0x34E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x34F, 0x34F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x350, 0x36F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x370, 0x482, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x483, 0x487, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x488, 0x590, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x591, 0x5BD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x5BE, 0x5BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x5BF, 0x5BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x5C0, 0x5C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x5C1, 0x5C2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x5C3, 0x5C3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x5C4, 0x5C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x5C6, 0x5C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x5C7, 0x5C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x5C8, 0x60F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x610, 0x61A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x61B, 0x64A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x64B, 0x65F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x660, 0x66F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x670, 0x670, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x671, 0x6D5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x6D6, 0x6DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x6DD, 0x6DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x6DF, 0x6E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x6E5, 0x6E6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x6E7, 0x6E8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x6E9, 0x6E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x6EA, 0x6ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x6EE, 0x710, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x711, 0x711, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x712, 0x72F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x730, 0x74A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x74B, 0x7EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x7EB, 0x7F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x7F4, 0x7FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x7FD, 0x7FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x7FE, 0x815, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x816, 0x819, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x81A, 0x81A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x81B, 0x823, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x824, 0x824, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x825, 0x827, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x828, 0x828, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x829, 0x82D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x82E, 0x858, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x859, 0x85B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x85C, 0x897, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0x898, 0x89F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x8A0, 0x8C9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0xA3D, 0xA94, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0xC3D, 0xC4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbNone,
		}, []*autogenTestCaseFunc{}},
		{0xF82, 0xF84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x1B6B, 0x1B73, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x3099, 0x309A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x102E0, 0x102E0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11236, 0x11236, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11D97, 0x11D97, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x1E2AE, 0x1E2AE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncIncbConsonant,
			&autogenTestCasesFuncIncbExtend,
			&autogenTestCasesFuncIncbLinker,
			&autogenTestCasesFuncIncbNone,
		}},
	},
	{ // Lb
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
			&autogenTestCasesFuncLbAk,
			&autogenTestCasesFuncLbAl,
			&autogenTestCasesFuncLbAp,
			&autogenTestCasesFuncLbAs,
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
			&autogenTestCasesFuncLbVf,
			&autogenTestCasesFuncLbVi,
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
		{0x600, 0x605, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x606, 0x608, []*autogenTestCaseFunc{
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
		{0x6DD, 0x6DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x6DE, 0x6DE, []*autogenTestCaseFunc{
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
			&autogenTestCasesFuncLbNu,
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
			&autogenTestCasesFuncLbNu,
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
		{0xE81, 0xE82, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSa,
		}, []*autogenTestCaseFunc{}},
		{0xEA5, 0xEA5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSa,
		}, []*autogenTestCaseFunc{}},
		{0xEC8, 0xECE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSa,
		}, []*autogenTestCaseFunc{}},
		{0xF05, 0xF05, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xF13, 0xF13, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xF35, 0xF35, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xF3D, 0xF3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0xF80, 0xF84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xFBE, 0xFBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0xFD3, 0xFD3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBb,
		}, []*autogenTestCaseFunc{}},
		{0x1050, 0x108F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSa,
		}, []*autogenTestCaseFunc{}},
		{0x10CE, 0x10CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x124E, 0x124F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1289, 0x1289, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x12BF, 0x12BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1311, 0x1311, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1362, 0x137C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1400, 0x1400, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x16EB, 0x16ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x1735, 0x1736, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x1771, 0x1771, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x17D9, 0x17D9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x17FA, 0x17FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x180B, 0x180D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x1885, 0x1886, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x191F, 0x191F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1946, 0x194F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x19CA, 0x19CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1A1E, 0x1A1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1A90, 0x1A99, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x1B34, 0x1B43, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x1B61, 0x1B6A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1BAE, 0x1BAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1C24, 0x1C37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x1C7E, 0x1C7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x1CD3, 0x1CD3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1CFA, 0x1CFA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1E00, 0x1F15, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1F50, 0x1F57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1F5F, 0x1F7D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1FD6, 0x1FDB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1FFE, 0x1FFE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x200E, 0x200F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x201A, 0x201A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x2028, 0x2029, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBk,
		}, []*autogenTestCaseFunc{}},
		{0x203E, 0x2043, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2058, 0x205B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x2072, 0x2073, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2085, 0x208C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x20A8, 0x20B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0x20C1, 0x20CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0x2109, 0x2109, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPo,
		}, []*autogenTestCaseFunc{}},
		{0x212B, 0x212B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2160, 0x216B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x219A, 0x21D1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2204, 0x2206, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2212, 0x2213, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0x2223, 0x2223, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2234, 0x2237, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2252, 0x2252, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x226E, 0x226F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2299, 0x2299, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x2308, 0x2308, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x231C, 0x2328, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x244B, 0x245F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2590, 0x2591, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x25B4, 0x25B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x25C9, 0x25CA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x25F0, 0x25FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2610, 0x2613, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2620, 0x2638, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2662, 0x2662, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x266E, 0x266E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x26C9, 0x26CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x26DA, 0x26DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
		}, []*autogenTestCaseFunc{}},
		{0x26EA, 0x26EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x26FD, 0x2704, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x2761, 0x2761, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x276C, 0x276C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x2774, 0x2774, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x27E7, 0x27E7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0x27EF, 0x27EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0x2989, 0x2989, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x2991, 0x2991, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x2999, 0x29D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x29FE, 0x2B54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2CF2, 0x2CF3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2D26, 0x2D26, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2D70, 0x2D70, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x2DAF, 0x2DAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2DCF, 0x2DCF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2E16, 0x2E16, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x2E22, 0x2E22, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x2E2A, 0x2E2D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x2E3C, 0x2E3E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x2E4E, 0x2E4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x2E5A, 0x2E5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0x2EF4, 0x2EFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
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
		{0x3250, 0x4DBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0xA4D0, 0xA4FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA62A, 0xA62B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA6A0, 0xA6EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA7D2, 0xA7D2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA806, 0xA806, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA830, 0xA837, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA880, 0xA881, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA8F2, 0xA8FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA930, 0xA946, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xA9B3, 0xA9BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xA9DA, 0xA9DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAA37, 0xAA3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xAA5C, 0xAA5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0xAAF2, 0xAAF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xAB17, 0xAB1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xABE3, 0xABEA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0xAC1C, 0xAC1C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xAC8C, 0xAC8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xACFC, 0xACFC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xAD6C, 0xAD6C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xADDC, 0xADDC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xAE4C, 0xAE4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xAEBC, 0xAEBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xAF2C, 0xAF2C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xAF9C, 0xAF9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB00C, 0xB00C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB07C, 0xB07C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB0EC, 0xB0EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB15C, 0xB15C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB1CC, 0xB1CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB23C, 0xB23C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB2AC, 0xB2AC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB31C, 0xB31C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB38C, 0xB38C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB3FC, 0xB3FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB46C, 0xB46C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB4DC, 0xB4DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB54C, 0xB54C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB5BC, 0xB5BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB62C, 0xB62C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB69C, 0xB69C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB70C, 0xB70C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB77C, 0xB77C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB7EC, 0xB7EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB85C, 0xB85C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB8CC, 0xB8CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB93C, 0xB93C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xB9AC, 0xB9AC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBA1C, 0xBA1C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBA8C, 0xBA8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBAFC, 0xBAFC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBB6C, 0xBB6C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBBDC, 0xBBDC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBC4C, 0xBC4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBCBC, 0xBCBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBD2C, 0xBD2C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBD9C, 0xBD9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBE0C, 0xBE0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBE7C, 0xBE7C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBEEC, 0xBEEC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBF5C, 0xBF5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xBFCC, 0xBFCC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC03C, 0xC03C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC0AC, 0xC0AC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC11C, 0xC11C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC18C, 0xC18C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC1FC, 0xC1FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC26C, 0xC26C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC2DC, 0xC2DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC34C, 0xC34C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC3BC, 0xC3BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC42C, 0xC42C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC49C, 0xC49C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC50C, 0xC50C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC57C, 0xC57C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC5EC, 0xC5EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC65C, 0xC65C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC6CC, 0xC6CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC73C, 0xC73C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC7AC, 0xC7AC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC81C, 0xC81C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC88C, 0xC88C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC8FC, 0xC8FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC96C, 0xC96C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xC9DC, 0xC9DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCA4C, 0xCA4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCABC, 0xCABC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCB2C, 0xCB2C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCB9C, 0xCB9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCC0C, 0xCC0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCC7C, 0xCC7C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCCEC, 0xCCEC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCD5C, 0xCD5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCDCC, 0xCDCC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCE3C, 0xCE3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCEAC, 0xCEAC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCF1C, 0xCF1C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCF8C, 0xCF8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xCFFC, 0xCFFC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD06C, 0xD06C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD0DC, 0xD0DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD14C, 0xD14C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD1BC, 0xD1BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD22C, 0xD22C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD29C, 0xD29C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD30C, 0xD30C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD37C, 0xD37C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD3EC, 0xD3EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD45C, 0xD45C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD4CC, 0xD4CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD53C, 0xD53C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD5AC, 0xD5AC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD61C, 0xD61C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD68C, 0xD68C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD6FC, 0xD6FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD76C, 0xD76C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbH2,
		}, []*autogenTestCaseFunc{}},
		{0xD7FC, 0xD7FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFB1D, 0xFB1D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbHl,
		}, []*autogenTestCaseFunc{}},
		{0xFB3E, 0xFB3E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbHl,
		}, []*autogenTestCaseFunc{}},
		{0xFBC3, 0xFBD2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFDCF, 0xFDCF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xFE13, 0xFE14, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbIs,
		}, []*autogenTestCaseFunc{}},
		{0xFE35, 0xFE35, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0xFE3D, 0xFE3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0xFE45, 0xFE46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0xFE54, 0xFE55, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNs,
		}, []*autogenTestCaseFunc{}},
		{0xFE5E, 0xFE5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0xFE70, 0xFE74, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0xFF04, 0xFF04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbPr,
		}, []*autogenTestCaseFunc{}},
		{0xFF0E, 0xFF0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0xFF3D, 0xFF3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCl,
		}, []*autogenTestCaseFunc{}},
		{0xFF62, 0xFF62, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0xFFBF, 0xFFC1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFFDD, 0xFFDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFFF9, 0xFFFB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x10028, 0x1003A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10080, 0x100FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10190, 0x1019C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1029D, 0x1029F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1032D, 0x1034A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x103A0, 0x103C3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x104A0, 0x104A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x10530, 0x10563, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10594, 0x10595, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x105BB, 0x105BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10780, 0x10785, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10808, 0x10808, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1083F, 0x10855, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x108F3, 0x108F3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1093F, 0x1093F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10A04, 0x10A04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10A19, 0x10A35, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10A58, 0x10A58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10AF0, 0x10AF5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x10B58, 0x10B72, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10C00, 0x10C48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x10D28, 0x10D2F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10EAD, 0x10EAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x10F46, 0x10F50, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x10FCC, 0x10FDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11047, 0x11048, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x11075, 0x11075, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAk,
		}, []*autogenTestCaseFunc{}},
		{0x110BE, 0x110C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x110FA, 0x110FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11145, 0x11146, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x11177, 0x1117F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x111C9, 0x111CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x111E0, 0x111E0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1123A, 0x1123A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11287, 0x11287, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x112A9, 0x112A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
		}, []*autogenTestCaseFunc{}},
		{0x11304, 0x11304, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11331, 0x11331, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11345, 0x11346, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11357, 0x11357, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x1136D, 0x1136F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11450, 0x11459, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbNu,
		}, []*autogenTestCaseFunc{}},
		{0x114B0, 0x114C3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x115B8, 0x115C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x115DE, 0x115FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11660, 0x1166C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBb,
		}, []*autogenTestCaseFunc{}},
		{0x11700, 0x1171A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbSa,
		}, []*autogenTestCaseFunc{}},
		{0x11747, 0x117FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x118F3, 0x118FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11915, 0x11916, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAk,
		}, []*autogenTestCaseFunc{}},
		{0x1193E, 0x1193E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbVi,
		}, []*autogenTestCaseFunc{}},
		{0x1195A, 0x1199F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x119E2, 0x119E2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBb,
		}, []*autogenTestCaseFunc{}},
		{0x11A3A, 0x11A3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x11A48, 0x11A4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11AA1, 0x11AA2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbBa,
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
		{0x11F03, 0x11F03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
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
		{0x1343C, 0x1343C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x14400, 0x145CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x16A5F, 0x16A5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16AD0, 0x16AED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x16B3A, 0x16B43, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x16B63, 0x16B77, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x16F00, 0x16F4A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x16FA0, 0x16FDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x18800, 0x18AFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1AFFC, 0x1AFFC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1B153, 0x1B154, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1BC6B, 0x1BC6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1BC9D, 0x1BC9E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x1CF50, 0x1CFC3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1D16A, 0x1D16C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1D200, 0x1D241, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1D300, 0x1D356, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1D49E, 0x1D49F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1D4AE, 0x1D4B9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1D507, 0x1D50A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1D53B, 0x1D53E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1D552, 0x1D6A5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1DA3B, 0x1DA6C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x1DA8C, 0x1DA9A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1DF2B, 0x1DFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E025, 0x1E025, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E12D, 0x1E12F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E290, 0x1E2AD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1E300, 0x1E4CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E7EC, 0x1E7EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E8D0, 0x1E8D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbCm,
		}, []*autogenTestCaseFunc{}},
		{0x1E95E, 0x1E95F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbOp,
		}, []*autogenTestCaseFunc{}},
		{0x1ED01, 0x1ED3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1EE24, 0x1EE24, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1EE39, 0x1EE39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1EE49, 0x1EE49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1EE54, 0x1EE54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1EE5D, 0x1EE5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1EE67, 0x1EE6A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1EE7E, 0x1EE7E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1EEA5, 0x1EEA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1F10D, 0x1F10F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F1E6, 0x1F1FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbRi,
		}, []*autogenTestCaseFunc{}},
		{0x1F3BC, 0x1F3BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1F3FB, 0x1F3FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEm,
		}, []*autogenTestCaseFunc{}},
		{0x1F47C, 0x1F47C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEb,
		}, []*autogenTestCaseFunc{}},
		{0x1F491, 0x1F491, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEb,
		}, []*autogenTestCaseFunc{}},
		{0x1F4AA, 0x1F4AA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEb,
		}, []*autogenTestCaseFunc{}},
		{0x1F517, 0x1F524, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1F590, 0x1F590, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEb,
		}, []*autogenTestCaseFunc{}},
		{0x1F645, 0x1F647, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEb,
		}, []*autogenTestCaseFunc{}},
		{0x1F6A3, 0x1F6A3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbEb,
		}, []*autogenTestCaseFunc{}},
		{0x1F700, 0x1F773, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1F850, 0x1F859, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAl,
		}, []*autogenTestCaseFunc{}},
		{0x1F90D, 0x1F90E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F93A, 0x1F93B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1F9BA, 0x1F9BA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1FA54, 0x1FAC2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbId,
		}, []*autogenTestCaseFunc{}},
		{0x1FBCB, 0x1FBEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x3FFFE, 0xE0000, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10FFFE, 0x10FFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbXx,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncLbAi,
			&autogenTestCasesFuncLbAk,
			&autogenTestCasesFuncLbAl,
			&autogenTestCasesFuncLbAp,
			&autogenTestCasesFuncLbAs,
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
			&autogenTestCasesFuncLbVf,
			&autogenTestCasesFuncLbVi,
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
		{0x17E0, 0x17E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0x180B, 0x180D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x1885, 0x1886, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x191F, 0x191F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x196E, 0x196F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x19DA, 0x19FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1A7D, 0x1A7E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1AAC, 0x1AAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1B50, 0x1B59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0x1B7F, 0x1B7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1BF4, 0x1BFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1C4D, 0x1C4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1CBD, 0x1CBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1CF4, 0x1CF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x1E01, 0x1E01, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E09, 0x1E09, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E11, 0x1E11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E19, 0x1E19, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E21, 0x1E21, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E29, 0x1E29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E31, 0x1E31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E39, 0x1E39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E41, 0x1E41, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E49, 0x1E49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E51, 0x1E51, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E59, 0x1E59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E61, 0x1E61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E69, 0x1E69, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E71, 0x1E71, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E79, 0x1E79, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E81, 0x1E81, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E89, 0x1E89, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1E91, 0x1E91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1EA1, 0x1EA1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1EA9, 0x1EA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1EB1, 0x1EB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1EB9, 0x1EB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1EC1, 0x1EC1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1EC9, 0x1EC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1ED1, 0x1ED1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1ED9, 0x1ED9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1EE1, 0x1EE1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1EE9, 0x1EE9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1EF1, 0x1EF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1EF9, 0x1EF9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1F10, 0x1F15, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1F40, 0x1F45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1F5B, 0x1F5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1F7E, 0x1F7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1FB5, 0x1FB5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1FC6, 0x1FC7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x1FE0, 0x1FE7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2000, 0x200A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSp,
		}, []*autogenTestCaseFunc{}},
		{0x2020, 0x2023, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x203B, 0x203B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2065, 0x2065, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x208D, 0x208E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0x2107, 0x2107, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2115, 0x2115, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2128, 0x2128, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2139, 0x2139, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x214E, 0x214E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2189, 0x2307, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x275B, 0x2760, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0x2983, 0x2998, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0x2C60, 0x2C60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2C6B, 0x2C6B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2C7E, 0x2C80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2C88, 0x2C88, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2C90, 0x2C90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2C98, 0x2C98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2CA0, 0x2CA0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2CA8, 0x2CA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2CB0, 0x2CB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2CB8, 0x2CB8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2CC0, 0x2CC0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2CC8, 0x2CC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2CD0, 0x2CD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2CD8, 0x2CD8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2CE0, 0x2CE0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x2CEE, 0x2CEE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x2D28, 0x2D2C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2D80, 0x2D96, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x2DB8, 0x2DBE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x2DD8, 0x2DDE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x2E2A, 0x2E2D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2E53, 0x2E54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0x3008, 0x3011, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0x3030, 0x3030, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x309B, 0x309C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x3130, 0x3130, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x4DC0, 0x4DFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA60E, 0xA60F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0xA643, 0xA643, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA64B, 0xA64B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA653, 0xA653, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA65B, 0xA65B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA663, 0xA663, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA66B, 0xA66B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA67F, 0xA67F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA687, 0xA687, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA68F, 0xA68F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA697, 0xA697, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA6F2, 0xA6F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xA723, 0xA723, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA72B, 0xA72B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA735, 0xA735, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA73D, 0xA73D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA745, 0xA745, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA74D, 0xA74D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA755, 0xA755, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA75D, 0xA75D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA765, 0xA765, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA76D, 0xA76D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA77F, 0xA77F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA787, 0xA787, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA790, 0xA790, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA79A, 0xA79A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA7A2, 0xA7A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA7AA, 0xA7AE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA7BA, 0xA7BA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA7C2, 0xA7C2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xA7D1, 0xA7D1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA7D9, 0xA7D9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xA802, 0xA802, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xA82C, 0xA82C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xA8B4, 0xA8C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xA8FB, 0xA8FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xA92F, 0xA92F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0xA9B3, 0xA9C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xA9E5, 0xA9E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xAA40, 0xAA42, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAA60, 0xAA76, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAAB5, 0xAAB6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAADB, 0xAADD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAB01, 0xAB06, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xAB28, 0xAB2E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xABE3, 0xABEA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xD7B0, 0xD7C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xFB00, 0xFB06, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0xFB2A, 0xFB36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xFB43, 0xFB44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xFD90, 0xFD91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xFE13, 0xFE13, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSc,
		}, []*autogenTestCaseFunc{}},
		{0xFE35, 0xFE44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0xFE56, 0xFE57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0xFE76, 0xFEFC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0xFF0C, 0xFF0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSc,
		}, []*autogenTestCaseFunc{}},
		{0xFF21, 0xFF3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0xFF5D, 0xFF5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbCl,
		}, []*autogenTestCaseFunc{}},
		{0xFF9E, 0xFF9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0xFFD8, 0xFFD9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10027, 0x10027, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1005E, 0x1007F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1029D, 0x1029F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1034B, 0x1034F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x103C8, 0x103CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x104A0, 0x104A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0x10530, 0x10563, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10594, 0x10595, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x105BB, 0x105BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x10780, 0x10780, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLo,
		}, []*autogenTestCaseFunc{}},
		{0x10800, 0x10805, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1083C, 0x1083C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x108E0, 0x108F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10980, 0x109B7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10A07, 0x10A0B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10A38, 0x10A3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x10A80, 0x10A9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x10B36, 0x10B3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10C49, 0x10C7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x10D30, 0x10D39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0x10EFD, 0x10EFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x10F55, 0x10F59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0x10FE0, 0x10FF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11070, 0x11070, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x110BB, 0x110BC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x110E9, 0x110EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11140, 0x11140, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11174, 0x11175, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x111C7, 0x111C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x111DD, 0x111DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1123A, 0x1123A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11287, 0x11287, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x112A9, 0x112A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0x11304, 0x11304, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11331, 0x11331, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11345, 0x11346, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11358, 0x1135C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11400, 0x11434, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1145F, 0x11461, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x114D0, 0x114D9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0x115C4, 0x115C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11643, 0x11643, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x116B9, 0x116BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1173A, 0x1173B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x118A0, 0x118BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbUp,
		}, []*autogenTestCaseFunc{}},
		{0x1190C, 0x11913, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11939, 0x1193A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11946, 0x11946, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbSt,
		}, []*autogenTestCaseFunc{}},
		{0x119D8, 0x119D9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11A01, 0x11A0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11A47, 0x11A47, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11A9D, 0x11A9D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11C37, 0x11C37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11C90, 0x11C91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11D0A, 0x11D0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11D3F, 0x11D45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11D67, 0x11D68, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11D98, 0x11D98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11F00, 0x11F01, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x11F3E, 0x11F42, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x1239A, 0x123FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x13430, 0x1343F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbFo,
		}, []*autogenTestCaseFunc{}},
		{0x16A39, 0x16A3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16AC0, 0x16AC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0x16B30, 0x16B36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x16B63, 0x16B77, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x16E99, 0x16EFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16F93, 0x16F9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x16FF2, 0x16FFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1AFF4, 0x1AFF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1B133, 0x1B14F, []*autogenTestCaseFunc{
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
		{0x1E01B, 0x1E021, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x1E08F, 0x1E08F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x1E14A, 0x1E14D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E2F0, 0x1E2F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbNu,
		}, []*autogenTestCaseFunc{}},
		{0x1E7E8, 0x1E7EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1E8D0, 0x1E8D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbEx,
		}, []*autogenTestCaseFunc{}},
		{0x1E95A, 0x1EDFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE25, 0x1EE26, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE3A, 0x1EE3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE4A, 0x1EE4A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE55, 0x1EE56, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE5E, 0x1EE5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE6B, 0x1EE6B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE7F, 0x1EE7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EEAA, 0x1EEAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1F18A, 0x1F675, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2B73A, 0x2B73F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0x2EE5E, 0x2F7FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncSbXx,
		}, []*autogenTestCaseFunc{}},
		{0xE0002, 0xE001F, []*autogenTestCaseFunc{
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
			&autogenTestCasesFuncWbNu,
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
			&autogenTestCasesFuncWbNu,
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
		{0x70F, 0x710, []*autogenTestCaseFunc{
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
			&autogenTestCasesFuncWbNu,
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
			&autogenTestCasesFuncWbNu,
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
		{0x10F46, 0x10F50, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x10FF7, 0x10FFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11073, 0x11074, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x110BE, 0x110C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x110FA, 0x110FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11145, 0x11146, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11180, 0x11182, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x111D0, 0x111D9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0x1122C, 0x11237, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11288, 0x11288, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x112B0, 0x112DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1130D, 0x1130E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11334, 0x11334, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11349, 0x1134A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11362, 0x11363, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11447, 0x1144A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x114B0, 0x114C3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x115AF, 0x115B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11630, 0x11640, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x116B8, 0x116B8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11800, 0x1182B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11909, 0x11909, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11936, 0x11936, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11944, 0x1194F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x119DA, 0x119E0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11A0B, 0x11A32, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11A51, 0x11A5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11C00, 0x11C08, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11C50, 0x11C59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0x11D00, 0x11D06, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x11D3B, 0x11D3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11D5A, 0x11D5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x11D90, 0x11D91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11EF3, 0x11EF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x11F34, 0x11F3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbExtend,
		}, []*autogenTestCaseFunc{}},
		{0x12000, 0x12399, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x13000, 0x1342F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x16800, 0x16A38, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x16AC0, 0x16AC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbNu,
		}, []*autogenTestCaseFunc{}},
		{0x16B37, 0x16B3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16B90, 0x16E3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16F88, 0x16F8E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x16FE5, 0x16FEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1AFFF, 0x1AFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1B168, 0x1BBFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1BC9A, 0x1BC9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1CF47, 0x1D164, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D18C, 0x1D1A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D49D, 0x1D49D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D4AD, 0x1D4AD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D506, 0x1D506, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D53A, 0x1D53A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D551, 0x1D551, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D6FB, 0x1D6FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D76F, 0x1D76F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1D7CC, 0x1D7CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1DA76, 0x1DA83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1DF1F, 0x1DF24, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E022, 0x1E022, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E090, 0x1E0FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E14E, 0x1E14E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbLe,
		}, []*autogenTestCaseFunc{}},
		{0x1E2FA, 0x1E4CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E7EC, 0x1E7EC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1E8D7, 0x1E8FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE04, 0x1EE04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE28, 0x1EE28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE3C, 0x1EE41, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE4C, 0x1EE4C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE58, 0x1EE58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE60, 0x1EE60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE73, 0x1EE73, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EE8A, 0x1EE8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1EEBC, 0x1F12F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0x1F200, 0x1F3FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
		}, []*autogenTestCaseFunc{}},
		{0xE0080, 0xE00FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncWbXx,
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
	{[]rune{0x0020, 0x0020}, []int{0, 1, 2}, "÷ 0020 ÷ 0020 ÷"},
	{[]rune{0x0020, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0020, 0x000D}, []int{0, 1, 2}, "÷ 0020 ÷ 000D ÷"},
	{[]rune{0x0020, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 000D ÷"},
	{[]rune{0x0020, 0x000A}, []int{0, 1, 2}, "÷ 0020 ÷ 000A ÷"},
	{[]rune{0x0020, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 000A ÷"},
	{[]rune{0x0020, 0x0001}, []int{0, 1, 2}, "÷ 0020 ÷ 0001 ÷"},
	{[]rune{0x0020, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0020, 0x034F}, []int{0, 2}, "÷ 0020 × 034F ÷"},
	{[]rune{0x0020, 0x0308, 0x034F}, []int{0, 3}, "÷ 0020 × 0308 × 034F ÷"},
	{[]rune{0x0020, 0x1F1E6}, []int{0, 1, 2}, "÷ 0020 ÷ 1F1E6 ÷"},
	{[]rune{0x0020, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0020, 0x0600}, []int{0, 1, 2}, "÷ 0020 ÷ 0600 ÷"},
	{[]rune{0x0020, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0020, 0x0A03}, []int{0, 2}, "÷ 0020 × 0A03 ÷"},
	{[]rune{0x0020, 0x0308, 0x0A03}, []int{0, 3}, "÷ 0020 × 0308 × 0A03 ÷"},
	{[]rune{0x0020, 0x1100}, []int{0, 1, 2}, "÷ 0020 ÷ 1100 ÷"},
	{[]rune{0x0020, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0020, 0x1160}, []int{0, 1, 2}, "÷ 0020 ÷ 1160 ÷"},
	{[]rune{0x0020, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0020, 0x11A8}, []int{0, 1, 2}, "÷ 0020 ÷ 11A8 ÷"},
	{[]rune{0x0020, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0020, 0xAC00}, []int{0, 1, 2}, "÷ 0020 ÷ AC00 ÷"},
	{[]rune{0x0020, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0020, 0xAC01}, []int{0, 1, 2}, "÷ 0020 ÷ AC01 ÷"},
	{[]rune{0x0020, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0020, 0x0900}, []int{0, 2}, "÷ 0020 × 0900 ÷"},
	{[]rune{0x0020, 0x0308, 0x0900}, []int{0, 3}, "÷ 0020 × 0308 × 0900 ÷"},
	{[]rune{0x0020, 0x0903}, []int{0, 2}, "÷ 0020 × 0903 ÷"},
	{[]rune{0x0020, 0x0308, 0x0903}, []int{0, 3}, "÷ 0020 × 0308 × 0903 ÷"},
	{[]rune{0x0020, 0x0904}, []int{0, 1, 2}, "÷ 0020 ÷ 0904 ÷"},
	{[]rune{0x0020, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 0904 ÷"},
	{[]rune{0x0020, 0x0D4E}, []int{0, 1, 2}, "÷ 0020 ÷ 0D4E ÷"},
	{[]rune{0x0020, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 0D4E ÷"},
	{[]rune{0x0020, 0x0915}, []int{0, 1, 2}, "÷ 0020 ÷ 0915 ÷"},
	{[]rune{0x0020, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 0915 ÷"},
	{[]rune{0x0020, 0x231A}, []int{0, 1, 2}, "÷ 0020 ÷ 231A ÷"},
	{[]rune{0x0020, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 231A ÷"},
	{[]rune{0x0020, 0x0300}, []int{0, 2}, "÷ 0020 × 0300 ÷"},
	{[]rune{0x0020, 0x0308, 0x0300}, []int{0, 3}, "÷ 0020 × 0308 × 0300 ÷"},
	{[]rune{0x0020, 0x093C}, []int{0, 2}, "÷ 0020 × 093C ÷"},
	{[]rune{0x0020, 0x0308, 0x093C}, []int{0, 3}, "÷ 0020 × 0308 × 093C ÷"},
	{[]rune{0x0020, 0x094D}, []int{0, 2}, "÷ 0020 × 094D ÷"},
	{[]rune{0x0020, 0x0308, 0x094D}, []int{0, 3}, "÷ 0020 × 0308 × 094D ÷"},
	{[]rune{0x0020, 0x200D}, []int{0, 2}, "÷ 0020 × 200D ÷"},
	{[]rune{0x0020, 0x0308, 0x200D}, []int{0, 3}, "÷ 0020 × 0308 × 200D ÷"},
	{[]rune{0x0020, 0x0378}, []int{0, 1, 2}, "÷ 0020 ÷ 0378 ÷"},
	{[]rune{0x0020, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 0020 × 0308 ÷ 0378 ÷"},
	{[]rune{0x000D, 0x0020}, []int{0, 1, 2}, "÷ 000D ÷ 0020 ÷"},
	{[]rune{0x000D, 0x0308, 0x0020}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 0020 ÷"},
	{[]rune{0x000D, 0x000D}, []int{0, 1, 2}, "÷ 000D ÷ 000D ÷"},
	{[]rune{0x000D, 0x0308, 0x000D}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 000D ÷"},
	{[]rune{0x000D, 0x000A}, []int{0, 2}, "÷ 000D × 000A ÷"},
	{[]rune{0x000D, 0x0308, 0x000A}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 000A ÷"},
	{[]rune{0x000D, 0x0001}, []int{0, 1, 2}, "÷ 000D ÷ 0001 ÷"},
	{[]rune{0x000D, 0x0308, 0x0001}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 0001 ÷"},
	{[]rune{0x000D, 0x034F}, []int{0, 1, 2}, "÷ 000D ÷ 034F ÷"},
	{[]rune{0x000D, 0x0308, 0x034F}, []int{0, 1, 3}, "÷ 000D ÷ 0308 × 034F ÷"},
	{[]rune{0x000D, 0x1F1E6}, []int{0, 1, 2}, "÷ 000D ÷ 1F1E6 ÷"},
	{[]rune{0x000D, 0x0308, 0x1F1E6}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x000D, 0x0600}, []int{0, 1, 2}, "÷ 000D ÷ 0600 ÷"},
	{[]rune{0x000D, 0x0308, 0x0600}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 0600 ÷"},
	{[]rune{0x000D, 0x0A03}, []int{0, 1, 2}, "÷ 000D ÷ 0A03 ÷"},
	{[]rune{0x000D, 0x0308, 0x0A03}, []int{0, 1, 3}, "÷ 000D ÷ 0308 × 0A03 ÷"},
	{[]rune{0x000D, 0x1100}, []int{0, 1, 2}, "÷ 000D ÷ 1100 ÷"},
	{[]rune{0x000D, 0x0308, 0x1100}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 1100 ÷"},
	{[]rune{0x000D, 0x1160}, []int{0, 1, 2}, "÷ 000D ÷ 1160 ÷"},
	{[]rune{0x000D, 0x0308, 0x1160}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 1160 ÷"},
	{[]rune{0x000D, 0x11A8}, []int{0, 1, 2}, "÷ 000D ÷ 11A8 ÷"},
	{[]rune{0x000D, 0x0308, 0x11A8}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 11A8 ÷"},
	{[]rune{0x000D, 0xAC00}, []int{0, 1, 2}, "÷ 000D ÷ AC00 ÷"},
	{[]rune{0x000D, 0x0308, 0xAC00}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ AC00 ÷"},
	{[]rune{0x000D, 0xAC01}, []int{0, 1, 2}, "÷ 000D ÷ AC01 ÷"},
	{[]rune{0x000D, 0x0308, 0xAC01}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ AC01 ÷"},
	{[]rune{0x000D, 0x0900}, []int{0, 1, 2}, "÷ 000D ÷ 0900 ÷"},
	{[]rune{0x000D, 0x0308, 0x0900}, []int{0, 1, 3}, "÷ 000D ÷ 0308 × 0900 ÷"},
	{[]rune{0x000D, 0x0903}, []int{0, 1, 2}, "÷ 000D ÷ 0903 ÷"},
	{[]rune{0x000D, 0x0308, 0x0903}, []int{0, 1, 3}, "÷ 000D ÷ 0308 × 0903 ÷"},
	{[]rune{0x000D, 0x0904}, []int{0, 1, 2}, "÷ 000D ÷ 0904 ÷"},
	{[]rune{0x000D, 0x0308, 0x0904}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 0904 ÷"},
	{[]rune{0x000D, 0x0D4E}, []int{0, 1, 2}, "÷ 000D ÷ 0D4E ÷"},
	{[]rune{0x000D, 0x0308, 0x0D4E}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 0D4E ÷"},
	{[]rune{0x000D, 0x0915}, []int{0, 1, 2}, "÷ 000D ÷ 0915 ÷"},
	{[]rune{0x000D, 0x0308, 0x0915}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 0915 ÷"},
	{[]rune{0x000D, 0x231A}, []int{0, 1, 2}, "÷ 000D ÷ 231A ÷"},
	{[]rune{0x000D, 0x0308, 0x231A}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 231A ÷"},
	{[]rune{0x000D, 0x0300}, []int{0, 1, 2}, "÷ 000D ÷ 0300 ÷"},
	{[]rune{0x000D, 0x0308, 0x0300}, []int{0, 1, 3}, "÷ 000D ÷ 0308 × 0300 ÷"},
	{[]rune{0x000D, 0x093C}, []int{0, 1, 2}, "÷ 000D ÷ 093C ÷"},
	{[]rune{0x000D, 0x0308, 0x093C}, []int{0, 1, 3}, "÷ 000D ÷ 0308 × 093C ÷"},
	{[]rune{0x000D, 0x094D}, []int{0, 1, 2}, "÷ 000D ÷ 094D ÷"},
	{[]rune{0x000D, 0x0308, 0x094D}, []int{0, 1, 3}, "÷ 000D ÷ 0308 × 094D ÷"},
	{[]rune{0x000D, 0x200D}, []int{0, 1, 2}, "÷ 000D ÷ 200D ÷"},
	{[]rune{0x000D, 0x0308, 0x200D}, []int{0, 1, 3}, "÷ 000D ÷ 0308 × 200D ÷"},
	{[]rune{0x000D, 0x0378}, []int{0, 1, 2}, "÷ 000D ÷ 0378 ÷"},
	{[]rune{0x000D, 0x0308, 0x0378}, []int{0, 1, 2, 3}, "÷ 000D ÷ 0308 ÷ 0378 ÷"},
	{[]rune{0x000A, 0x0020}, []int{0, 1, 2}, "÷ 000A ÷ 0020 ÷"},
	{[]rune{0x000A, 0x0308, 0x0020}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 0020 ÷"},
	{[]rune{0x000A, 0x000D}, []int{0, 1, 2}, "÷ 000A ÷ 000D ÷"},
	{[]rune{0x000A, 0x0308, 0x000D}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 000D ÷"},
	{[]rune{0x000A, 0x000A}, []int{0, 1, 2}, "÷ 000A ÷ 000A ÷"},
	{[]rune{0x000A, 0x0308, 0x000A}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 000A ÷"},
	{[]rune{0x000A, 0x0001}, []int{0, 1, 2}, "÷ 000A ÷ 0001 ÷"},
	{[]rune{0x000A, 0x0308, 0x0001}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 0001 ÷"},
	{[]rune{0x000A, 0x034F}, []int{0, 1, 2}, "÷ 000A ÷ 034F ÷"},
	{[]rune{0x000A, 0x0308, 0x034F}, []int{0, 1, 3}, "÷ 000A ÷ 0308 × 034F ÷"},
	{[]rune{0x000A, 0x1F1E6}, []int{0, 1, 2}, "÷ 000A ÷ 1F1E6 ÷"},
	{[]rune{0x000A, 0x0308, 0x1F1E6}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x000A, 0x0600}, []int{0, 1, 2}, "÷ 000A ÷ 0600 ÷"},
	{[]rune{0x000A, 0x0308, 0x0600}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 0600 ÷"},
	{[]rune{0x000A, 0x0A03}, []int{0, 1, 2}, "÷ 000A ÷ 0A03 ÷"},
	{[]rune{0x000A, 0x0308, 0x0A03}, []int{0, 1, 3}, "÷ 000A ÷ 0308 × 0A03 ÷"},
	{[]rune{0x000A, 0x1100}, []int{0, 1, 2}, "÷ 000A ÷ 1100 ÷"},
	{[]rune{0x000A, 0x0308, 0x1100}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 1100 ÷"},
	{[]rune{0x000A, 0x1160}, []int{0, 1, 2}, "÷ 000A ÷ 1160 ÷"},
	{[]rune{0x000A, 0x0308, 0x1160}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 1160 ÷"},
	{[]rune{0x000A, 0x11A8}, []int{0, 1, 2}, "÷ 000A ÷ 11A8 ÷"},
	{[]rune{0x000A, 0x0308, 0x11A8}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 11A8 ÷"},
	{[]rune{0x000A, 0xAC00}, []int{0, 1, 2}, "÷ 000A ÷ AC00 ÷"},
	{[]rune{0x000A, 0x0308, 0xAC00}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ AC00 ÷"},
	{[]rune{0x000A, 0xAC01}, []int{0, 1, 2}, "÷ 000A ÷ AC01 ÷"},
	{[]rune{0x000A, 0x0308, 0xAC01}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ AC01 ÷"},
	{[]rune{0x000A, 0x0900}, []int{0, 1, 2}, "÷ 000A ÷ 0900 ÷"},
	{[]rune{0x000A, 0x0308, 0x0900}, []int{0, 1, 3}, "÷ 000A ÷ 0308 × 0900 ÷"},
	{[]rune{0x000A, 0x0903}, []int{0, 1, 2}, "÷ 000A ÷ 0903 ÷"},
	{[]rune{0x000A, 0x0308, 0x0903}, []int{0, 1, 3}, "÷ 000A ÷ 0308 × 0903 ÷"},
	{[]rune{0x000A, 0x0904}, []int{0, 1, 2}, "÷ 000A ÷ 0904 ÷"},
	{[]rune{0x000A, 0x0308, 0x0904}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 0904 ÷"},
	{[]rune{0x000A, 0x0D4E}, []int{0, 1, 2}, "÷ 000A ÷ 0D4E ÷"},
	{[]rune{0x000A, 0x0308, 0x0D4E}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 0D4E ÷"},
	{[]rune{0x000A, 0x0915}, []int{0, 1, 2}, "÷ 000A ÷ 0915 ÷"},
	{[]rune{0x000A, 0x0308, 0x0915}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 0915 ÷"},
	{[]rune{0x000A, 0x231A}, []int{0, 1, 2}, "÷ 000A ÷ 231A ÷"},
	{[]rune{0x000A, 0x0308, 0x231A}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 231A ÷"},
	{[]rune{0x000A, 0x0300}, []int{0, 1, 2}, "÷ 000A ÷ 0300 ÷"},
	{[]rune{0x000A, 0x0308, 0x0300}, []int{0, 1, 3}, "÷ 000A ÷ 0308 × 0300 ÷"},
	{[]rune{0x000A, 0x093C}, []int{0, 1, 2}, "÷ 000A ÷ 093C ÷"},
	{[]rune{0x000A, 0x0308, 0x093C}, []int{0, 1, 3}, "÷ 000A ÷ 0308 × 093C ÷"},
	{[]rune{0x000A, 0x094D}, []int{0, 1, 2}, "÷ 000A ÷ 094D ÷"},
	{[]rune{0x000A, 0x0308, 0x094D}, []int{0, 1, 3}, "÷ 000A ÷ 0308 × 094D ÷"},
	{[]rune{0x000A, 0x200D}, []int{0, 1, 2}, "÷ 000A ÷ 200D ÷"},
	{[]rune{0x000A, 0x0308, 0x200D}, []int{0, 1, 3}, "÷ 000A ÷ 0308 × 200D ÷"},
	{[]rune{0x000A, 0x0378}, []int{0, 1, 2}, "÷ 000A ÷ 0378 ÷"},
	{[]rune{0x000A, 0x0308, 0x0378}, []int{0, 1, 2, 3}, "÷ 000A ÷ 0308 ÷ 0378 ÷"},
	{[]rune{0x0001, 0x0020}, []int{0, 1, 2}, "÷ 0001 ÷ 0020 ÷"},
	{[]rune{0x0001, 0x0308, 0x0020}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 0020 ÷"},
	{[]rune{0x0001, 0x000D}, []int{0, 1, 2}, "÷ 0001 ÷ 000D ÷"},
	{[]rune{0x0001, 0x0308, 0x000D}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 000D ÷"},
	{[]rune{0x0001, 0x000A}, []int{0, 1, 2}, "÷ 0001 ÷ 000A ÷"},
	{[]rune{0x0001, 0x0308, 0x000A}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 000A ÷"},
	{[]rune{0x0001, 0x0001}, []int{0, 1, 2}, "÷ 0001 ÷ 0001 ÷"},
	{[]rune{0x0001, 0x0308, 0x0001}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 0001 ÷"},
	{[]rune{0x0001, 0x034F}, []int{0, 1, 2}, "÷ 0001 ÷ 034F ÷"},
	{[]rune{0x0001, 0x0308, 0x034F}, []int{0, 1, 3}, "÷ 0001 ÷ 0308 × 034F ÷"},
	{[]rune{0x0001, 0x1F1E6}, []int{0, 1, 2}, "÷ 0001 ÷ 1F1E6 ÷"},
	{[]rune{0x0001, 0x0308, 0x1F1E6}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0001, 0x0600}, []int{0, 1, 2}, "÷ 0001 ÷ 0600 ÷"},
	{[]rune{0x0001, 0x0308, 0x0600}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 0600 ÷"},
	{[]rune{0x0001, 0x0A03}, []int{0, 1, 2}, "÷ 0001 ÷ 0A03 ÷"},
	{[]rune{0x0001, 0x0308, 0x0A03}, []int{0, 1, 3}, "÷ 0001 ÷ 0308 × 0A03 ÷"},
	{[]rune{0x0001, 0x1100}, []int{0, 1, 2}, "÷ 0001 ÷ 1100 ÷"},
	{[]rune{0x0001, 0x0308, 0x1100}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 1100 ÷"},
	{[]rune{0x0001, 0x1160}, []int{0, 1, 2}, "÷ 0001 ÷ 1160 ÷"},
	{[]rune{0x0001, 0x0308, 0x1160}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 1160 ÷"},
	{[]rune{0x0001, 0x11A8}, []int{0, 1, 2}, "÷ 0001 ÷ 11A8 ÷"},
	{[]rune{0x0001, 0x0308, 0x11A8}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 11A8 ÷"},
	{[]rune{0x0001, 0xAC00}, []int{0, 1, 2}, "÷ 0001 ÷ AC00 ÷"},
	{[]rune{0x0001, 0x0308, 0xAC00}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ AC00 ÷"},
	{[]rune{0x0001, 0xAC01}, []int{0, 1, 2}, "÷ 0001 ÷ AC01 ÷"},
	{[]rune{0x0001, 0x0308, 0xAC01}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ AC01 ÷"},
	{[]rune{0x0001, 0x0900}, []int{0, 1, 2}, "÷ 0001 ÷ 0900 ÷"},
	{[]rune{0x0001, 0x0308, 0x0900}, []int{0, 1, 3}, "÷ 0001 ÷ 0308 × 0900 ÷"},
	{[]rune{0x0001, 0x0903}, []int{0, 1, 2}, "÷ 0001 ÷ 0903 ÷"},
	{[]rune{0x0001, 0x0308, 0x0903}, []int{0, 1, 3}, "÷ 0001 ÷ 0308 × 0903 ÷"},
	{[]rune{0x0001, 0x0904}, []int{0, 1, 2}, "÷ 0001 ÷ 0904 ÷"},
	{[]rune{0x0001, 0x0308, 0x0904}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 0904 ÷"},
	{[]rune{0x0001, 0x0D4E}, []int{0, 1, 2}, "÷ 0001 ÷ 0D4E ÷"},
	{[]rune{0x0001, 0x0308, 0x0D4E}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 0D4E ÷"},
	{[]rune{0x0001, 0x0915}, []int{0, 1, 2}, "÷ 0001 ÷ 0915 ÷"},
	{[]rune{0x0001, 0x0308, 0x0915}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 0915 ÷"},
	{[]rune{0x0001, 0x231A}, []int{0, 1, 2}, "÷ 0001 ÷ 231A ÷"},
	{[]rune{0x0001, 0x0308, 0x231A}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 231A ÷"},
	{[]rune{0x0001, 0x0300}, []int{0, 1, 2}, "÷ 0001 ÷ 0300 ÷"},
	{[]rune{0x0001, 0x0308, 0x0300}, []int{0, 1, 3}, "÷ 0001 ÷ 0308 × 0300 ÷"},
	{[]rune{0x0001, 0x093C}, []int{0, 1, 2}, "÷ 0001 ÷ 093C ÷"},
	{[]rune{0x0001, 0x0308, 0x093C}, []int{0, 1, 3}, "÷ 0001 ÷ 0308 × 093C ÷"},
	{[]rune{0x0001, 0x094D}, []int{0, 1, 2}, "÷ 0001 ÷ 094D ÷"},
	{[]rune{0x0001, 0x0308, 0x094D}, []int{0, 1, 3}, "÷ 0001 ÷ 0308 × 094D ÷"},
	{[]rune{0x0001, 0x200D}, []int{0, 1, 2}, "÷ 0001 ÷ 200D ÷"},
	{[]rune{0x0001, 0x0308, 0x200D}, []int{0, 1, 3}, "÷ 0001 ÷ 0308 × 200D ÷"},
	{[]rune{0x0001, 0x0378}, []int{0, 1, 2}, "÷ 0001 ÷ 0378 ÷"},
	{[]rune{0x0001, 0x0308, 0x0378}, []int{0, 1, 2, 3}, "÷ 0001 ÷ 0308 ÷ 0378 ÷"},
	{[]rune{0x034F, 0x0020}, []int{0, 1, 2}, "÷ 034F ÷ 0020 ÷"},
	{[]rune{0x034F, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 0020 ÷"},
	{[]rune{0x034F, 0x000D}, []int{0, 1, 2}, "÷ 034F ÷ 000D ÷"},
	{[]rune{0x034F, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 000D ÷"},
	{[]rune{0x034F, 0x000A}, []int{0, 1, 2}, "÷ 034F ÷ 000A ÷"},
	{[]rune{0x034F, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 000A ÷"},
	{[]rune{0x034F, 0x0001}, []int{0, 1, 2}, "÷ 034F ÷ 0001 ÷"},
	{[]rune{0x034F, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 0001 ÷"},
	{[]rune{0x034F, 0x034F}, []int{0, 2}, "÷ 034F × 034F ÷"},
	{[]rune{0x034F, 0x0308, 0x034F}, []int{0, 3}, "÷ 034F × 0308 × 034F ÷"},
	{[]rune{0x034F, 0x1F1E6}, []int{0, 1, 2}, "÷ 034F ÷ 1F1E6 ÷"},
	{[]rune{0x034F, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x034F, 0x0600}, []int{0, 1, 2}, "÷ 034F ÷ 0600 ÷"},
	{[]rune{0x034F, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 0600 ÷"},
	{[]rune{0x034F, 0x0A03}, []int{0, 2}, "÷ 034F × 0A03 ÷"},
	{[]rune{0x034F, 0x0308, 0x0A03}, []int{0, 3}, "÷ 034F × 0308 × 0A03 ÷"},
	{[]rune{0x034F, 0x1100}, []int{0, 1, 2}, "÷ 034F ÷ 1100 ÷"},
	{[]rune{0x034F, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 1100 ÷"},
	{[]rune{0x034F, 0x1160}, []int{0, 1, 2}, "÷ 034F ÷ 1160 ÷"},
	{[]rune{0x034F, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 1160 ÷"},
	{[]rune{0x034F, 0x11A8}, []int{0, 1, 2}, "÷ 034F ÷ 11A8 ÷"},
	{[]rune{0x034F, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 11A8 ÷"},
	{[]rune{0x034F, 0xAC00}, []int{0, 1, 2}, "÷ 034F ÷ AC00 ÷"},
	{[]rune{0x034F, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ AC00 ÷"},
	{[]rune{0x034F, 0xAC01}, []int{0, 1, 2}, "÷ 034F ÷ AC01 ÷"},
	{[]rune{0x034F, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ AC01 ÷"},
	{[]rune{0x034F, 0x0900}, []int{0, 2}, "÷ 034F × 0900 ÷"},
	{[]rune{0x034F, 0x0308, 0x0900}, []int{0, 3}, "÷ 034F × 0308 × 0900 ÷"},
	{[]rune{0x034F, 0x0903}, []int{0, 2}, "÷ 034F × 0903 ÷"},
	{[]rune{0x034F, 0x0308, 0x0903}, []int{0, 3}, "÷ 034F × 0308 × 0903 ÷"},
	{[]rune{0x034F, 0x0904}, []int{0, 1, 2}, "÷ 034F ÷ 0904 ÷"},
	{[]rune{0x034F, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 0904 ÷"},
	{[]rune{0x034F, 0x0D4E}, []int{0, 1, 2}, "÷ 034F ÷ 0D4E ÷"},
	{[]rune{0x034F, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 0D4E ÷"},
	{[]rune{0x034F, 0x0915}, []int{0, 1, 2}, "÷ 034F ÷ 0915 ÷"},
	{[]rune{0x034F, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 0915 ÷"},
	{[]rune{0x034F, 0x231A}, []int{0, 1, 2}, "÷ 034F ÷ 231A ÷"},
	{[]rune{0x034F, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 231A ÷"},
	{[]rune{0x034F, 0x0300}, []int{0, 2}, "÷ 034F × 0300 ÷"},
	{[]rune{0x034F, 0x0308, 0x0300}, []int{0, 3}, "÷ 034F × 0308 × 0300 ÷"},
	{[]rune{0x034F, 0x093C}, []int{0, 2}, "÷ 034F × 093C ÷"},
	{[]rune{0x034F, 0x0308, 0x093C}, []int{0, 3}, "÷ 034F × 0308 × 093C ÷"},
	{[]rune{0x034F, 0x094D}, []int{0, 2}, "÷ 034F × 094D ÷"},
	{[]rune{0x034F, 0x0308, 0x094D}, []int{0, 3}, "÷ 034F × 0308 × 094D ÷"},
	{[]rune{0x034F, 0x200D}, []int{0, 2}, "÷ 034F × 200D ÷"},
	{[]rune{0x034F, 0x0308, 0x200D}, []int{0, 3}, "÷ 034F × 0308 × 200D ÷"},
	{[]rune{0x034F, 0x0378}, []int{0, 1, 2}, "÷ 034F ÷ 0378 ÷"},
	{[]rune{0x034F, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 034F × 0308 ÷ 0378 ÷"},
	{[]rune{0x1F1E6, 0x0020}, []int{0, 1, 2}, "÷ 1F1E6 ÷ 0020 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 0020 ÷"},
	{[]rune{0x1F1E6, 0x000D}, []int{0, 1, 2}, "÷ 1F1E6 ÷ 000D ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 000D ÷"},
	{[]rune{0x1F1E6, 0x000A}, []int{0, 1, 2}, "÷ 1F1E6 ÷ 000A ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 000A ÷"},
	{[]rune{0x1F1E6, 0x0001}, []int{0, 1, 2}, "÷ 1F1E6 ÷ 0001 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 0001 ÷"},
	{[]rune{0x1F1E6, 0x034F}, []int{0, 2}, "÷ 1F1E6 × 034F ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x034F}, []int{0, 3}, "÷ 1F1E6 × 0308 × 034F ÷"},
	{[]rune{0x1F1E6, 0x1F1E6}, []int{0, 2}, "÷ 1F1E6 × 1F1E6 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x1F1E6, 0x0600}, []int{0, 1, 2}, "÷ 1F1E6 ÷ 0600 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 0600 ÷"},
	{[]rune{0x1F1E6, 0x0A03}, []int{0, 2}, "÷ 1F1E6 × 0A03 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x0A03}, []int{0, 3}, "÷ 1F1E6 × 0308 × 0A03 ÷"},
	{[]rune{0x1F1E6, 0x1100}, []int{0, 1, 2}, "÷ 1F1E6 ÷ 1100 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 1100 ÷"},
	{[]rune{0x1F1E6, 0x1160}, []int{0, 1, 2}, "÷ 1F1E6 ÷ 1160 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 1160 ÷"},
	{[]rune{0x1F1E6, 0x11A8}, []int{0, 1, 2}, "÷ 1F1E6 ÷ 11A8 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x1F1E6, 0xAC00}, []int{0, 1, 2}, "÷ 1F1E6 ÷ AC00 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ AC00 ÷"},
	{[]rune{0x1F1E6, 0xAC01}, []int{0, 1, 2}, "÷ 1F1E6 ÷ AC01 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ AC01 ÷"},
	{[]rune{0x1F1E6, 0x0900}, []int{0, 2}, "÷ 1F1E6 × 0900 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x0900}, []int{0, 3}, "÷ 1F1E6 × 0308 × 0900 ÷"},
	{[]rune{0x1F1E6, 0x0903}, []int{0, 2}, "÷ 1F1E6 × 0903 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x0903}, []int{0, 3}, "÷ 1F1E6 × 0308 × 0903 ÷"},
	{[]rune{0x1F1E6, 0x0904}, []int{0, 1, 2}, "÷ 1F1E6 ÷ 0904 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 0904 ÷"},
	{[]rune{0x1F1E6, 0x0D4E}, []int{0, 1, 2}, "÷ 1F1E6 ÷ 0D4E ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 0D4E ÷"},
	{[]rune{0x1F1E6, 0x0915}, []int{0, 1, 2}, "÷ 1F1E6 ÷ 0915 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 0915 ÷"},
	{[]rune{0x1F1E6, 0x231A}, []int{0, 1, 2}, "÷ 1F1E6 ÷ 231A ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 231A ÷"},
	{[]rune{0x1F1E6, 0x0300}, []int{0, 2}, "÷ 1F1E6 × 0300 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x0300}, []int{0, 3}, "÷ 1F1E6 × 0308 × 0300 ÷"},
	{[]rune{0x1F1E6, 0x093C}, []int{0, 2}, "÷ 1F1E6 × 093C ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x093C}, []int{0, 3}, "÷ 1F1E6 × 0308 × 093C ÷"},
	{[]rune{0x1F1E6, 0x094D}, []int{0, 2}, "÷ 1F1E6 × 094D ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x094D}, []int{0, 3}, "÷ 1F1E6 × 0308 × 094D ÷"},
	{[]rune{0x1F1E6, 0x200D}, []int{0, 2}, "÷ 1F1E6 × 200D ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x200D}, []int{0, 3}, "÷ 1F1E6 × 0308 × 200D ÷"},
	{[]rune{0x1F1E6, 0x0378}, []int{0, 1, 2}, "÷ 1F1E6 ÷ 0378 ÷"},
	{[]rune{0x1F1E6, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 1F1E6 × 0308 ÷ 0378 ÷"},
	{[]rune{0x0600, 0x0020}, []int{0, 2}, "÷ 0600 × 0020 ÷"},
	{[]rune{0x0600, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0600, 0x000D}, []int{0, 1, 2}, "÷ 0600 ÷ 000D ÷"},
	{[]rune{0x0600, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 000D ÷"},
	{[]rune{0x0600, 0x000A}, []int{0, 1, 2}, "÷ 0600 ÷ 000A ÷"},
	{[]rune{0x0600, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 000A ÷"},
	{[]rune{0x0600, 0x0001}, []int{0, 1, 2}, "÷ 0600 ÷ 0001 ÷"},
	{[]rune{0x0600, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0600, 0x034F}, []int{0, 2}, "÷ 0600 × 034F ÷"},
	{[]rune{0x0600, 0x0308, 0x034F}, []int{0, 3}, "÷ 0600 × 0308 × 034F ÷"},
	{[]rune{0x0600, 0x1F1E6}, []int{0, 2}, "÷ 0600 × 1F1E6 ÷"},
	{[]rune{0x0600, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0600, 0x0600}, []int{0, 2}, "÷ 0600 × 0600 ÷"},
	{[]rune{0x0600, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0600, 0x0A03}, []int{0, 2}, "÷ 0600 × 0A03 ÷"},
	{[]rune{0x0600, 0x0308, 0x0A03}, []int{0, 3}, "÷ 0600 × 0308 × 0A03 ÷"},
	{[]rune{0x0600, 0x1100}, []int{0, 2}, "÷ 0600 × 1100 ÷"},
	{[]rune{0x0600, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0600, 0x1160}, []int{0, 2}, "÷ 0600 × 1160 ÷"},
	{[]rune{0x0600, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0600, 0x11A8}, []int{0, 2}, "÷ 0600 × 11A8 ÷"},
	{[]rune{0x0600, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0600, 0xAC00}, []int{0, 2}, "÷ 0600 × AC00 ÷"},
	{[]rune{0x0600, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0600, 0xAC01}, []int{0, 2}, "÷ 0600 × AC01 ÷"},
	{[]rune{0x0600, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0600, 0x0900}, []int{0, 2}, "÷ 0600 × 0900 ÷"},
	{[]rune{0x0600, 0x0308, 0x0900}, []int{0, 3}, "÷ 0600 × 0308 × 0900 ÷"},
	{[]rune{0x0600, 0x0903}, []int{0, 2}, "÷ 0600 × 0903 ÷"},
	{[]rune{0x0600, 0x0308, 0x0903}, []int{0, 3}, "÷ 0600 × 0308 × 0903 ÷"},
	{[]rune{0x0600, 0x0904}, []int{0, 2}, "÷ 0600 × 0904 ÷"},
	{[]rune{0x0600, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 0904 ÷"},
	{[]rune{0x0600, 0x0D4E}, []int{0, 2}, "÷ 0600 × 0D4E ÷"},
	{[]rune{0x0600, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 0D4E ÷"},
	{[]rune{0x0600, 0x0915}, []int{0, 2}, "÷ 0600 × 0915 ÷"},
	{[]rune{0x0600, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 0915 ÷"},
	{[]rune{0x0600, 0x231A}, []int{0, 2}, "÷ 0600 × 231A ÷"},
	{[]rune{0x0600, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 231A ÷"},
	{[]rune{0x0600, 0x0300}, []int{0, 2}, "÷ 0600 × 0300 ÷"},
	{[]rune{0x0600, 0x0308, 0x0300}, []int{0, 3}, "÷ 0600 × 0308 × 0300 ÷"},
	{[]rune{0x0600, 0x093C}, []int{0, 2}, "÷ 0600 × 093C ÷"},
	{[]rune{0x0600, 0x0308, 0x093C}, []int{0, 3}, "÷ 0600 × 0308 × 093C ÷"},
	{[]rune{0x0600, 0x094D}, []int{0, 2}, "÷ 0600 × 094D ÷"},
	{[]rune{0x0600, 0x0308, 0x094D}, []int{0, 3}, "÷ 0600 × 0308 × 094D ÷"},
	{[]rune{0x0600, 0x200D}, []int{0, 2}, "÷ 0600 × 200D ÷"},
	{[]rune{0x0600, 0x0308, 0x200D}, []int{0, 3}, "÷ 0600 × 0308 × 200D ÷"},
	{[]rune{0x0600, 0x0378}, []int{0, 2}, "÷ 0600 × 0378 ÷"},
	{[]rune{0x0600, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 0600 × 0308 ÷ 0378 ÷"},
	{[]rune{0x0A03, 0x0020}, []int{0, 1, 2}, "÷ 0A03 ÷ 0020 ÷"},
	{[]rune{0x0A03, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0A03, 0x000D}, []int{0, 1, 2}, "÷ 0A03 ÷ 000D ÷"},
	{[]rune{0x0A03, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 000D ÷"},
	{[]rune{0x0A03, 0x000A}, []int{0, 1, 2}, "÷ 0A03 ÷ 000A ÷"},
	{[]rune{0x0A03, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 000A ÷"},
	{[]rune{0x0A03, 0x0001}, []int{0, 1, 2}, "÷ 0A03 ÷ 0001 ÷"},
	{[]rune{0x0A03, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0A03, 0x034F}, []int{0, 2}, "÷ 0A03 × 034F ÷"},
	{[]rune{0x0A03, 0x0308, 0x034F}, []int{0, 3}, "÷ 0A03 × 0308 × 034F ÷"},
	{[]rune{0x0A03, 0x1F1E6}, []int{0, 1, 2}, "÷ 0A03 ÷ 1F1E6 ÷"},
	{[]rune{0x0A03, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0A03, 0x0600}, []int{0, 1, 2}, "÷ 0A03 ÷ 0600 ÷"},
	{[]rune{0x0A03, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0A03, 0x0A03}, []int{0, 2}, "÷ 0A03 × 0A03 ÷"},
	{[]rune{0x0A03, 0x0308, 0x0A03}, []int{0, 3}, "÷ 0A03 × 0308 × 0A03 ÷"},
	{[]rune{0x0A03, 0x1100}, []int{0, 1, 2}, "÷ 0A03 ÷ 1100 ÷"},
	{[]rune{0x0A03, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0A03, 0x1160}, []int{0, 1, 2}, "÷ 0A03 ÷ 1160 ÷"},
	{[]rune{0x0A03, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0A03, 0x11A8}, []int{0, 1, 2}, "÷ 0A03 ÷ 11A8 ÷"},
	{[]rune{0x0A03, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0A03, 0xAC00}, []int{0, 1, 2}, "÷ 0A03 ÷ AC00 ÷"},
	{[]rune{0x0A03, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0A03, 0xAC01}, []int{0, 1, 2}, "÷ 0A03 ÷ AC01 ÷"},
	{[]rune{0x0A03, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0A03, 0x0900}, []int{0, 2}, "÷ 0A03 × 0900 ÷"},
	{[]rune{0x0A03, 0x0308, 0x0900}, []int{0, 3}, "÷ 0A03 × 0308 × 0900 ÷"},
	{[]rune{0x0A03, 0x0903}, []int{0, 2}, "÷ 0A03 × 0903 ÷"},
	{[]rune{0x0A03, 0x0308, 0x0903}, []int{0, 3}, "÷ 0A03 × 0308 × 0903 ÷"},
	{[]rune{0x0A03, 0x0904}, []int{0, 1, 2}, "÷ 0A03 ÷ 0904 ÷"},
	{[]rune{0x0A03, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 0904 ÷"},
	{[]rune{0x0A03, 0x0D4E}, []int{0, 1, 2}, "÷ 0A03 ÷ 0D4E ÷"},
	{[]rune{0x0A03, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 0D4E ÷"},
	{[]rune{0x0A03, 0x0915}, []int{0, 1, 2}, "÷ 0A03 ÷ 0915 ÷"},
	{[]rune{0x0A03, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 0915 ÷"},
	{[]rune{0x0A03, 0x231A}, []int{0, 1, 2}, "÷ 0A03 ÷ 231A ÷"},
	{[]rune{0x0A03, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 231A ÷"},
	{[]rune{0x0A03, 0x0300}, []int{0, 2}, "÷ 0A03 × 0300 ÷"},
	{[]rune{0x0A03, 0x0308, 0x0300}, []int{0, 3}, "÷ 0A03 × 0308 × 0300 ÷"},
	{[]rune{0x0A03, 0x093C}, []int{0, 2}, "÷ 0A03 × 093C ÷"},
	{[]rune{0x0A03, 0x0308, 0x093C}, []int{0, 3}, "÷ 0A03 × 0308 × 093C ÷"},
	{[]rune{0x0A03, 0x094D}, []int{0, 2}, "÷ 0A03 × 094D ÷"},
	{[]rune{0x0A03, 0x0308, 0x094D}, []int{0, 3}, "÷ 0A03 × 0308 × 094D ÷"},
	{[]rune{0x0A03, 0x200D}, []int{0, 2}, "÷ 0A03 × 200D ÷"},
	{[]rune{0x0A03, 0x0308, 0x200D}, []int{0, 3}, "÷ 0A03 × 0308 × 200D ÷"},
	{[]rune{0x0A03, 0x0378}, []int{0, 1, 2}, "÷ 0A03 ÷ 0378 ÷"},
	{[]rune{0x0A03, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 0A03 × 0308 ÷ 0378 ÷"},
	{[]rune{0x1100, 0x0020}, []int{0, 1, 2}, "÷ 1100 ÷ 0020 ÷"},
	{[]rune{0x1100, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 0020 ÷"},
	{[]rune{0x1100, 0x000D}, []int{0, 1, 2}, "÷ 1100 ÷ 000D ÷"},
	{[]rune{0x1100, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 000D ÷"},
	{[]rune{0x1100, 0x000A}, []int{0, 1, 2}, "÷ 1100 ÷ 000A ÷"},
	{[]rune{0x1100, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 000A ÷"},
	{[]rune{0x1100, 0x0001}, []int{0, 1, 2}, "÷ 1100 ÷ 0001 ÷"},
	{[]rune{0x1100, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 0001 ÷"},
	{[]rune{0x1100, 0x034F}, []int{0, 2}, "÷ 1100 × 034F ÷"},
	{[]rune{0x1100, 0x0308, 0x034F}, []int{0, 3}, "÷ 1100 × 0308 × 034F ÷"},
	{[]rune{0x1100, 0x1F1E6}, []int{0, 1, 2}, "÷ 1100 ÷ 1F1E6 ÷"},
	{[]rune{0x1100, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x1100, 0x0600}, []int{0, 1, 2}, "÷ 1100 ÷ 0600 ÷"},
	{[]rune{0x1100, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 0600 ÷"},
	{[]rune{0x1100, 0x0A03}, []int{0, 2}, "÷ 1100 × 0A03 ÷"},
	{[]rune{0x1100, 0x0308, 0x0A03}, []int{0, 3}, "÷ 1100 × 0308 × 0A03 ÷"},
	{[]rune{0x1100, 0x1100}, []int{0, 2}, "÷ 1100 × 1100 ÷"},
	{[]rune{0x1100, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 1100 ÷"},
	{[]rune{0x1100, 0x1160}, []int{0, 2}, "÷ 1100 × 1160 ÷"},
	{[]rune{0x1100, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 1160 ÷"},
	{[]rune{0x1100, 0x11A8}, []int{0, 1, 2}, "÷ 1100 ÷ 11A8 ÷"},
	{[]rune{0x1100, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x1100, 0xAC00}, []int{0, 2}, "÷ 1100 × AC00 ÷"},
	{[]rune{0x1100, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ AC00 ÷"},
	{[]rune{0x1100, 0xAC01}, []int{0, 2}, "÷ 1100 × AC01 ÷"},
	{[]rune{0x1100, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ AC01 ÷"},
	{[]rune{0x1100, 0x0900}, []int{0, 2}, "÷ 1100 × 0900 ÷"},
	{[]rune{0x1100, 0x0308, 0x0900}, []int{0, 3}, "÷ 1100 × 0308 × 0900 ÷"},
	{[]rune{0x1100, 0x0903}, []int{0, 2}, "÷ 1100 × 0903 ÷"},
	{[]rune{0x1100, 0x0308, 0x0903}, []int{0, 3}, "÷ 1100 × 0308 × 0903 ÷"},
	{[]rune{0x1100, 0x0904}, []int{0, 1, 2}, "÷ 1100 ÷ 0904 ÷"},
	{[]rune{0x1100, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 0904 ÷"},
	{[]rune{0x1100, 0x0D4E}, []int{0, 1, 2}, "÷ 1100 ÷ 0D4E ÷"},
	{[]rune{0x1100, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 0D4E ÷"},
	{[]rune{0x1100, 0x0915}, []int{0, 1, 2}, "÷ 1100 ÷ 0915 ÷"},
	{[]rune{0x1100, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 0915 ÷"},
	{[]rune{0x1100, 0x231A}, []int{0, 1, 2}, "÷ 1100 ÷ 231A ÷"},
	{[]rune{0x1100, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 231A ÷"},
	{[]rune{0x1100, 0x0300}, []int{0, 2}, "÷ 1100 × 0300 ÷"},
	{[]rune{0x1100, 0x0308, 0x0300}, []int{0, 3}, "÷ 1100 × 0308 × 0300 ÷"},
	{[]rune{0x1100, 0x093C}, []int{0, 2}, "÷ 1100 × 093C ÷"},
	{[]rune{0x1100, 0x0308, 0x093C}, []int{0, 3}, "÷ 1100 × 0308 × 093C ÷"},
	{[]rune{0x1100, 0x094D}, []int{0, 2}, "÷ 1100 × 094D ÷"},
	{[]rune{0x1100, 0x0308, 0x094D}, []int{0, 3}, "÷ 1100 × 0308 × 094D ÷"},
	{[]rune{0x1100, 0x200D}, []int{0, 2}, "÷ 1100 × 200D ÷"},
	{[]rune{0x1100, 0x0308, 0x200D}, []int{0, 3}, "÷ 1100 × 0308 × 200D ÷"},
	{[]rune{0x1100, 0x0378}, []int{0, 1, 2}, "÷ 1100 ÷ 0378 ÷"},
	{[]rune{0x1100, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 1100 × 0308 ÷ 0378 ÷"},
	{[]rune{0x1160, 0x0020}, []int{0, 1, 2}, "÷ 1160 ÷ 0020 ÷"},
	{[]rune{0x1160, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 0020 ÷"},
	{[]rune{0x1160, 0x000D}, []int{0, 1, 2}, "÷ 1160 ÷ 000D ÷"},
	{[]rune{0x1160, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 000D ÷"},
	{[]rune{0x1160, 0x000A}, []int{0, 1, 2}, "÷ 1160 ÷ 000A ÷"},
	{[]rune{0x1160, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 000A ÷"},
	{[]rune{0x1160, 0x0001}, []int{0, 1, 2}, "÷ 1160 ÷ 0001 ÷"},
	{[]rune{0x1160, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 0001 ÷"},
	{[]rune{0x1160, 0x034F}, []int{0, 2}, "÷ 1160 × 034F ÷"},
	{[]rune{0x1160, 0x0308, 0x034F}, []int{0, 3}, "÷ 1160 × 0308 × 034F ÷"},
	{[]rune{0x1160, 0x1F1E6}, []int{0, 1, 2}, "÷ 1160 ÷ 1F1E6 ÷"},
	{[]rune{0x1160, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x1160, 0x0600}, []int{0, 1, 2}, "÷ 1160 ÷ 0600 ÷"},
	{[]rune{0x1160, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 0600 ÷"},
	{[]rune{0x1160, 0x0A03}, []int{0, 2}, "÷ 1160 × 0A03 ÷"},
	{[]rune{0x1160, 0x0308, 0x0A03}, []int{0, 3}, "÷ 1160 × 0308 × 0A03 ÷"},
	{[]rune{0x1160, 0x1100}, []int{0, 1, 2}, "÷ 1160 ÷ 1100 ÷"},
	{[]rune{0x1160, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 1100 ÷"},
	{[]rune{0x1160, 0x1160}, []int{0, 2}, "÷ 1160 × 1160 ÷"},
	{[]rune{0x1160, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 1160 ÷"},
	{[]rune{0x1160, 0x11A8}, []int{0, 2}, "÷ 1160 × 11A8 ÷"},
	{[]rune{0x1160, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x1160, 0xAC00}, []int{0, 1, 2}, "÷ 1160 ÷ AC00 ÷"},
	{[]rune{0x1160, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ AC00 ÷"},
	{[]rune{0x1160, 0xAC01}, []int{0, 1, 2}, "÷ 1160 ÷ AC01 ÷"},
	{[]rune{0x1160, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ AC01 ÷"},
	{[]rune{0x1160, 0x0900}, []int{0, 2}, "÷ 1160 × 0900 ÷"},
	{[]rune{0x1160, 0x0308, 0x0900}, []int{0, 3}, "÷ 1160 × 0308 × 0900 ÷"},
	{[]rune{0x1160, 0x0903}, []int{0, 2}, "÷ 1160 × 0903 ÷"},
	{[]rune{0x1160, 0x0308, 0x0903}, []int{0, 3}, "÷ 1160 × 0308 × 0903 ÷"},
	{[]rune{0x1160, 0x0904}, []int{0, 1, 2}, "÷ 1160 ÷ 0904 ÷"},
	{[]rune{0x1160, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 0904 ÷"},
	{[]rune{0x1160, 0x0D4E}, []int{0, 1, 2}, "÷ 1160 ÷ 0D4E ÷"},
	{[]rune{0x1160, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 0D4E ÷"},
	{[]rune{0x1160, 0x0915}, []int{0, 1, 2}, "÷ 1160 ÷ 0915 ÷"},
	{[]rune{0x1160, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 0915 ÷"},
	{[]rune{0x1160, 0x231A}, []int{0, 1, 2}, "÷ 1160 ÷ 231A ÷"},
	{[]rune{0x1160, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 231A ÷"},
	{[]rune{0x1160, 0x0300}, []int{0, 2}, "÷ 1160 × 0300 ÷"},
	{[]rune{0x1160, 0x0308, 0x0300}, []int{0, 3}, "÷ 1160 × 0308 × 0300 ÷"},
	{[]rune{0x1160, 0x093C}, []int{0, 2}, "÷ 1160 × 093C ÷"},
	{[]rune{0x1160, 0x0308, 0x093C}, []int{0, 3}, "÷ 1160 × 0308 × 093C ÷"},
	{[]rune{0x1160, 0x094D}, []int{0, 2}, "÷ 1160 × 094D ÷"},
	{[]rune{0x1160, 0x0308, 0x094D}, []int{0, 3}, "÷ 1160 × 0308 × 094D ÷"},
	{[]rune{0x1160, 0x200D}, []int{0, 2}, "÷ 1160 × 200D ÷"},
	{[]rune{0x1160, 0x0308, 0x200D}, []int{0, 3}, "÷ 1160 × 0308 × 200D ÷"},
	{[]rune{0x1160, 0x0378}, []int{0, 1, 2}, "÷ 1160 ÷ 0378 ÷"},
	{[]rune{0x1160, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 1160 × 0308 ÷ 0378 ÷"},
	{[]rune{0x11A8, 0x0020}, []int{0, 1, 2}, "÷ 11A8 ÷ 0020 ÷"},
	{[]rune{0x11A8, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 0020 ÷"},
	{[]rune{0x11A8, 0x000D}, []int{0, 1, 2}, "÷ 11A8 ÷ 000D ÷"},
	{[]rune{0x11A8, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 000D ÷"},
	{[]rune{0x11A8, 0x000A}, []int{0, 1, 2}, "÷ 11A8 ÷ 000A ÷"},
	{[]rune{0x11A8, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 000A ÷"},
	{[]rune{0x11A8, 0x0001}, []int{0, 1, 2}, "÷ 11A8 ÷ 0001 ÷"},
	{[]rune{0x11A8, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 0001 ÷"},
	{[]rune{0x11A8, 0x034F}, []int{0, 2}, "÷ 11A8 × 034F ÷"},
	{[]rune{0x11A8, 0x0308, 0x034F}, []int{0, 3}, "÷ 11A8 × 0308 × 034F ÷"},
	{[]rune{0x11A8, 0x1F1E6}, []int{0, 1, 2}, "÷ 11A8 ÷ 1F1E6 ÷"},
	{[]rune{0x11A8, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x11A8, 0x0600}, []int{0, 1, 2}, "÷ 11A8 ÷ 0600 ÷"},
	{[]rune{0x11A8, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 0600 ÷"},
	{[]rune{0x11A8, 0x0A03}, []int{0, 2}, "÷ 11A8 × 0A03 ÷"},
	{[]rune{0x11A8, 0x0308, 0x0A03}, []int{0, 3}, "÷ 11A8 × 0308 × 0A03 ÷"},
	{[]rune{0x11A8, 0x1100}, []int{0, 1, 2}, "÷ 11A8 ÷ 1100 ÷"},
	{[]rune{0x11A8, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 1100 ÷"},
	{[]rune{0x11A8, 0x1160}, []int{0, 1, 2}, "÷ 11A8 ÷ 1160 ÷"},
	{[]rune{0x11A8, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 1160 ÷"},
	{[]rune{0x11A8, 0x11A8}, []int{0, 2}, "÷ 11A8 × 11A8 ÷"},
	{[]rune{0x11A8, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x11A8, 0xAC00}, []int{0, 1, 2}, "÷ 11A8 ÷ AC00 ÷"},
	{[]rune{0x11A8, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ AC00 ÷"},
	{[]rune{0x11A8, 0xAC01}, []int{0, 1, 2}, "÷ 11A8 ÷ AC01 ÷"},
	{[]rune{0x11A8, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ AC01 ÷"},
	{[]rune{0x11A8, 0x0900}, []int{0, 2}, "÷ 11A8 × 0900 ÷"},
	{[]rune{0x11A8, 0x0308, 0x0900}, []int{0, 3}, "÷ 11A8 × 0308 × 0900 ÷"},
	{[]rune{0x11A8, 0x0903}, []int{0, 2}, "÷ 11A8 × 0903 ÷"},
	{[]rune{0x11A8, 0x0308, 0x0903}, []int{0, 3}, "÷ 11A8 × 0308 × 0903 ÷"},
	{[]rune{0x11A8, 0x0904}, []int{0, 1, 2}, "÷ 11A8 ÷ 0904 ÷"},
	{[]rune{0x11A8, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 0904 ÷"},
	{[]rune{0x11A8, 0x0D4E}, []int{0, 1, 2}, "÷ 11A8 ÷ 0D4E ÷"},
	{[]rune{0x11A8, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 0D4E ÷"},
	{[]rune{0x11A8, 0x0915}, []int{0, 1, 2}, "÷ 11A8 ÷ 0915 ÷"},
	{[]rune{0x11A8, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 0915 ÷"},
	{[]rune{0x11A8, 0x231A}, []int{0, 1, 2}, "÷ 11A8 ÷ 231A ÷"},
	{[]rune{0x11A8, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 231A ÷"},
	{[]rune{0x11A8, 0x0300}, []int{0, 2}, "÷ 11A8 × 0300 ÷"},
	{[]rune{0x11A8, 0x0308, 0x0300}, []int{0, 3}, "÷ 11A8 × 0308 × 0300 ÷"},
	{[]rune{0x11A8, 0x093C}, []int{0, 2}, "÷ 11A8 × 093C ÷"},
	{[]rune{0x11A8, 0x0308, 0x093C}, []int{0, 3}, "÷ 11A8 × 0308 × 093C ÷"},
	{[]rune{0x11A8, 0x094D}, []int{0, 2}, "÷ 11A8 × 094D ÷"},
	{[]rune{0x11A8, 0x0308, 0x094D}, []int{0, 3}, "÷ 11A8 × 0308 × 094D ÷"},
	{[]rune{0x11A8, 0x200D}, []int{0, 2}, "÷ 11A8 × 200D ÷"},
	{[]rune{0x11A8, 0x0308, 0x200D}, []int{0, 3}, "÷ 11A8 × 0308 × 200D ÷"},
	{[]rune{0x11A8, 0x0378}, []int{0, 1, 2}, "÷ 11A8 ÷ 0378 ÷"},
	{[]rune{0x11A8, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 11A8 × 0308 ÷ 0378 ÷"},
	{[]rune{0xAC00, 0x0020}, []int{0, 1, 2}, "÷ AC00 ÷ 0020 ÷"},
	{[]rune{0xAC00, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 0020 ÷"},
	{[]rune{0xAC00, 0x000D}, []int{0, 1, 2}, "÷ AC00 ÷ 000D ÷"},
	{[]rune{0xAC00, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 000D ÷"},
	{[]rune{0xAC00, 0x000A}, []int{0, 1, 2}, "÷ AC00 ÷ 000A ÷"},
	{[]rune{0xAC00, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 000A ÷"},
	{[]rune{0xAC00, 0x0001}, []int{0, 1, 2}, "÷ AC00 ÷ 0001 ÷"},
	{[]rune{0xAC00, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 0001 ÷"},
	{[]rune{0xAC00, 0x034F}, []int{0, 2}, "÷ AC00 × 034F ÷"},
	{[]rune{0xAC00, 0x0308, 0x034F}, []int{0, 3}, "÷ AC00 × 0308 × 034F ÷"},
	{[]rune{0xAC00, 0x1F1E6}, []int{0, 1, 2}, "÷ AC00 ÷ 1F1E6 ÷"},
	{[]rune{0xAC00, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0xAC00, 0x0600}, []int{0, 1, 2}, "÷ AC00 ÷ 0600 ÷"},
	{[]rune{0xAC00, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 0600 ÷"},
	{[]rune{0xAC00, 0x0A03}, []int{0, 2}, "÷ AC00 × 0A03 ÷"},
	{[]rune{0xAC00, 0x0308, 0x0A03}, []int{0, 3}, "÷ AC00 × 0308 × 0A03 ÷"},
	{[]rune{0xAC00, 0x1100}, []int{0, 1, 2}, "÷ AC00 ÷ 1100 ÷"},
	{[]rune{0xAC00, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 1100 ÷"},
	{[]rune{0xAC00, 0x1160}, []int{0, 2}, "÷ AC00 × 1160 ÷"},
	{[]rune{0xAC00, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 1160 ÷"},
	{[]rune{0xAC00, 0x11A8}, []int{0, 2}, "÷ AC00 × 11A8 ÷"},
	{[]rune{0xAC00, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 11A8 ÷"},
	{[]rune{0xAC00, 0xAC00}, []int{0, 1, 2}, "÷ AC00 ÷ AC00 ÷"},
	{[]rune{0xAC00, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ AC00 ÷"},
	{[]rune{0xAC00, 0xAC01}, []int{0, 1, 2}, "÷ AC00 ÷ AC01 ÷"},
	{[]rune{0xAC00, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ AC01 ÷"},
	{[]rune{0xAC00, 0x0900}, []int{0, 2}, "÷ AC00 × 0900 ÷"},
	{[]rune{0xAC00, 0x0308, 0x0900}, []int{0, 3}, "÷ AC00 × 0308 × 0900 ÷"},
	{[]rune{0xAC00, 0x0903}, []int{0, 2}, "÷ AC00 × 0903 ÷"},
	{[]rune{0xAC00, 0x0308, 0x0903}, []int{0, 3}, "÷ AC00 × 0308 × 0903 ÷"},
	{[]rune{0xAC00, 0x0904}, []int{0, 1, 2}, "÷ AC00 ÷ 0904 ÷"},
	{[]rune{0xAC00, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 0904 ÷"},
	{[]rune{0xAC00, 0x0D4E}, []int{0, 1, 2}, "÷ AC00 ÷ 0D4E ÷"},
	{[]rune{0xAC00, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 0D4E ÷"},
	{[]rune{0xAC00, 0x0915}, []int{0, 1, 2}, "÷ AC00 ÷ 0915 ÷"},
	{[]rune{0xAC00, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 0915 ÷"},
	{[]rune{0xAC00, 0x231A}, []int{0, 1, 2}, "÷ AC00 ÷ 231A ÷"},
	{[]rune{0xAC00, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 231A ÷"},
	{[]rune{0xAC00, 0x0300}, []int{0, 2}, "÷ AC00 × 0300 ÷"},
	{[]rune{0xAC00, 0x0308, 0x0300}, []int{0, 3}, "÷ AC00 × 0308 × 0300 ÷"},
	{[]rune{0xAC00, 0x093C}, []int{0, 2}, "÷ AC00 × 093C ÷"},
	{[]rune{0xAC00, 0x0308, 0x093C}, []int{0, 3}, "÷ AC00 × 0308 × 093C ÷"},
	{[]rune{0xAC00, 0x094D}, []int{0, 2}, "÷ AC00 × 094D ÷"},
	{[]rune{0xAC00, 0x0308, 0x094D}, []int{0, 3}, "÷ AC00 × 0308 × 094D ÷"},
	{[]rune{0xAC00, 0x200D}, []int{0, 2}, "÷ AC00 × 200D ÷"},
	{[]rune{0xAC00, 0x0308, 0x200D}, []int{0, 3}, "÷ AC00 × 0308 × 200D ÷"},
	{[]rune{0xAC00, 0x0378}, []int{0, 1, 2}, "÷ AC00 ÷ 0378 ÷"},
	{[]rune{0xAC00, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ AC00 × 0308 ÷ 0378 ÷"},
	{[]rune{0xAC01, 0x0020}, []int{0, 1, 2}, "÷ AC01 ÷ 0020 ÷"},
	{[]rune{0xAC01, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 0020 ÷"},
	{[]rune{0xAC01, 0x000D}, []int{0, 1, 2}, "÷ AC01 ÷ 000D ÷"},
	{[]rune{0xAC01, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 000D ÷"},
	{[]rune{0xAC01, 0x000A}, []int{0, 1, 2}, "÷ AC01 ÷ 000A ÷"},
	{[]rune{0xAC01, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 000A ÷"},
	{[]rune{0xAC01, 0x0001}, []int{0, 1, 2}, "÷ AC01 ÷ 0001 ÷"},
	{[]rune{0xAC01, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 0001 ÷"},
	{[]rune{0xAC01, 0x034F}, []int{0, 2}, "÷ AC01 × 034F ÷"},
	{[]rune{0xAC01, 0x0308, 0x034F}, []int{0, 3}, "÷ AC01 × 0308 × 034F ÷"},
	{[]rune{0xAC01, 0x1F1E6}, []int{0, 1, 2}, "÷ AC01 ÷ 1F1E6 ÷"},
	{[]rune{0xAC01, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0xAC01, 0x0600}, []int{0, 1, 2}, "÷ AC01 ÷ 0600 ÷"},
	{[]rune{0xAC01, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 0600 ÷"},
	{[]rune{0xAC01, 0x0A03}, []int{0, 2}, "÷ AC01 × 0A03 ÷"},
	{[]rune{0xAC01, 0x0308, 0x0A03}, []int{0, 3}, "÷ AC01 × 0308 × 0A03 ÷"},
	{[]rune{0xAC01, 0x1100}, []int{0, 1, 2}, "÷ AC01 ÷ 1100 ÷"},
	{[]rune{0xAC01, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 1100 ÷"},
	{[]rune{0xAC01, 0x1160}, []int{0, 1, 2}, "÷ AC01 ÷ 1160 ÷"},
	{[]rune{0xAC01, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 1160 ÷"},
	{[]rune{0xAC01, 0x11A8}, []int{0, 2}, "÷ AC01 × 11A8 ÷"},
	{[]rune{0xAC01, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 11A8 ÷"},
	{[]rune{0xAC01, 0xAC00}, []int{0, 1, 2}, "÷ AC01 ÷ AC00 ÷"},
	{[]rune{0xAC01, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ AC00 ÷"},
	{[]rune{0xAC01, 0xAC01}, []int{0, 1, 2}, "÷ AC01 ÷ AC01 ÷"},
	{[]rune{0xAC01, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ AC01 ÷"},
	{[]rune{0xAC01, 0x0900}, []int{0, 2}, "÷ AC01 × 0900 ÷"},
	{[]rune{0xAC01, 0x0308, 0x0900}, []int{0, 3}, "÷ AC01 × 0308 × 0900 ÷"},
	{[]rune{0xAC01, 0x0903}, []int{0, 2}, "÷ AC01 × 0903 ÷"},
	{[]rune{0xAC01, 0x0308, 0x0903}, []int{0, 3}, "÷ AC01 × 0308 × 0903 ÷"},
	{[]rune{0xAC01, 0x0904}, []int{0, 1, 2}, "÷ AC01 ÷ 0904 ÷"},
	{[]rune{0xAC01, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 0904 ÷"},
	{[]rune{0xAC01, 0x0D4E}, []int{0, 1, 2}, "÷ AC01 ÷ 0D4E ÷"},
	{[]rune{0xAC01, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 0D4E ÷"},
	{[]rune{0xAC01, 0x0915}, []int{0, 1, 2}, "÷ AC01 ÷ 0915 ÷"},
	{[]rune{0xAC01, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 0915 ÷"},
	{[]rune{0xAC01, 0x231A}, []int{0, 1, 2}, "÷ AC01 ÷ 231A ÷"},
	{[]rune{0xAC01, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 231A ÷"},
	{[]rune{0xAC01, 0x0300}, []int{0, 2}, "÷ AC01 × 0300 ÷"},
	{[]rune{0xAC01, 0x0308, 0x0300}, []int{0, 3}, "÷ AC01 × 0308 × 0300 ÷"},
	{[]rune{0xAC01, 0x093C}, []int{0, 2}, "÷ AC01 × 093C ÷"},
	{[]rune{0xAC01, 0x0308, 0x093C}, []int{0, 3}, "÷ AC01 × 0308 × 093C ÷"},
	{[]rune{0xAC01, 0x094D}, []int{0, 2}, "÷ AC01 × 094D ÷"},
	{[]rune{0xAC01, 0x0308, 0x094D}, []int{0, 3}, "÷ AC01 × 0308 × 094D ÷"},
	{[]rune{0xAC01, 0x200D}, []int{0, 2}, "÷ AC01 × 200D ÷"},
	{[]rune{0xAC01, 0x0308, 0x200D}, []int{0, 3}, "÷ AC01 × 0308 × 200D ÷"},
	{[]rune{0xAC01, 0x0378}, []int{0, 1, 2}, "÷ AC01 ÷ 0378 ÷"},
	{[]rune{0xAC01, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ AC01 × 0308 ÷ 0378 ÷"},
	{[]rune{0x0900, 0x0020}, []int{0, 1, 2}, "÷ 0900 ÷ 0020 ÷"},
	{[]rune{0x0900, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0900, 0x000D}, []int{0, 1, 2}, "÷ 0900 ÷ 000D ÷"},
	{[]rune{0x0900, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 000D ÷"},
	{[]rune{0x0900, 0x000A}, []int{0, 1, 2}, "÷ 0900 ÷ 000A ÷"},
	{[]rune{0x0900, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 000A ÷"},
	{[]rune{0x0900, 0x0001}, []int{0, 1, 2}, "÷ 0900 ÷ 0001 ÷"},
	{[]rune{0x0900, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0900, 0x034F}, []int{0, 2}, "÷ 0900 × 034F ÷"},
	{[]rune{0x0900, 0x0308, 0x034F}, []int{0, 3}, "÷ 0900 × 0308 × 034F ÷"},
	{[]rune{0x0900, 0x1F1E6}, []int{0, 1, 2}, "÷ 0900 ÷ 1F1E6 ÷"},
	{[]rune{0x0900, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0900, 0x0600}, []int{0, 1, 2}, "÷ 0900 ÷ 0600 ÷"},
	{[]rune{0x0900, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0900, 0x0A03}, []int{0, 2}, "÷ 0900 × 0A03 ÷"},
	{[]rune{0x0900, 0x0308, 0x0A03}, []int{0, 3}, "÷ 0900 × 0308 × 0A03 ÷"},
	{[]rune{0x0900, 0x1100}, []int{0, 1, 2}, "÷ 0900 ÷ 1100 ÷"},
	{[]rune{0x0900, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0900, 0x1160}, []int{0, 1, 2}, "÷ 0900 ÷ 1160 ÷"},
	{[]rune{0x0900, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0900, 0x11A8}, []int{0, 1, 2}, "÷ 0900 ÷ 11A8 ÷"},
	{[]rune{0x0900, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0900, 0xAC00}, []int{0, 1, 2}, "÷ 0900 ÷ AC00 ÷"},
	{[]rune{0x0900, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0900, 0xAC01}, []int{0, 1, 2}, "÷ 0900 ÷ AC01 ÷"},
	{[]rune{0x0900, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0900, 0x0900}, []int{0, 2}, "÷ 0900 × 0900 ÷"},
	{[]rune{0x0900, 0x0308, 0x0900}, []int{0, 3}, "÷ 0900 × 0308 × 0900 ÷"},
	{[]rune{0x0900, 0x0903}, []int{0, 2}, "÷ 0900 × 0903 ÷"},
	{[]rune{0x0900, 0x0308, 0x0903}, []int{0, 3}, "÷ 0900 × 0308 × 0903 ÷"},
	{[]rune{0x0900, 0x0904}, []int{0, 1, 2}, "÷ 0900 ÷ 0904 ÷"},
	{[]rune{0x0900, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 0904 ÷"},
	{[]rune{0x0900, 0x0D4E}, []int{0, 1, 2}, "÷ 0900 ÷ 0D4E ÷"},
	{[]rune{0x0900, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 0D4E ÷"},
	{[]rune{0x0900, 0x0915}, []int{0, 1, 2}, "÷ 0900 ÷ 0915 ÷"},
	{[]rune{0x0900, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 0915 ÷"},
	{[]rune{0x0900, 0x231A}, []int{0, 1, 2}, "÷ 0900 ÷ 231A ÷"},
	{[]rune{0x0900, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 231A ÷"},
	{[]rune{0x0900, 0x0300}, []int{0, 2}, "÷ 0900 × 0300 ÷"},
	{[]rune{0x0900, 0x0308, 0x0300}, []int{0, 3}, "÷ 0900 × 0308 × 0300 ÷"},
	{[]rune{0x0900, 0x093C}, []int{0, 2}, "÷ 0900 × 093C ÷"},
	{[]rune{0x0900, 0x0308, 0x093C}, []int{0, 3}, "÷ 0900 × 0308 × 093C ÷"},
	{[]rune{0x0900, 0x094D}, []int{0, 2}, "÷ 0900 × 094D ÷"},
	{[]rune{0x0900, 0x0308, 0x094D}, []int{0, 3}, "÷ 0900 × 0308 × 094D ÷"},
	{[]rune{0x0900, 0x200D}, []int{0, 2}, "÷ 0900 × 200D ÷"},
	{[]rune{0x0900, 0x0308, 0x200D}, []int{0, 3}, "÷ 0900 × 0308 × 200D ÷"},
	{[]rune{0x0900, 0x0378}, []int{0, 1, 2}, "÷ 0900 ÷ 0378 ÷"},
	{[]rune{0x0900, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 0900 × 0308 ÷ 0378 ÷"},
	{[]rune{0x0903, 0x0020}, []int{0, 1, 2}, "÷ 0903 ÷ 0020 ÷"},
	{[]rune{0x0903, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0903, 0x000D}, []int{0, 1, 2}, "÷ 0903 ÷ 000D ÷"},
	{[]rune{0x0903, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 000D ÷"},
	{[]rune{0x0903, 0x000A}, []int{0, 1, 2}, "÷ 0903 ÷ 000A ÷"},
	{[]rune{0x0903, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 000A ÷"},
	{[]rune{0x0903, 0x0001}, []int{0, 1, 2}, "÷ 0903 ÷ 0001 ÷"},
	{[]rune{0x0903, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0903, 0x034F}, []int{0, 2}, "÷ 0903 × 034F ÷"},
	{[]rune{0x0903, 0x0308, 0x034F}, []int{0, 3}, "÷ 0903 × 0308 × 034F ÷"},
	{[]rune{0x0903, 0x1F1E6}, []int{0, 1, 2}, "÷ 0903 ÷ 1F1E6 ÷"},
	{[]rune{0x0903, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0903, 0x0600}, []int{0, 1, 2}, "÷ 0903 ÷ 0600 ÷"},
	{[]rune{0x0903, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0903, 0x0A03}, []int{0, 2}, "÷ 0903 × 0A03 ÷"},
	{[]rune{0x0903, 0x0308, 0x0A03}, []int{0, 3}, "÷ 0903 × 0308 × 0A03 ÷"},
	{[]rune{0x0903, 0x1100}, []int{0, 1, 2}, "÷ 0903 ÷ 1100 ÷"},
	{[]rune{0x0903, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0903, 0x1160}, []int{0, 1, 2}, "÷ 0903 ÷ 1160 ÷"},
	{[]rune{0x0903, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0903, 0x11A8}, []int{0, 1, 2}, "÷ 0903 ÷ 11A8 ÷"},
	{[]rune{0x0903, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0903, 0xAC00}, []int{0, 1, 2}, "÷ 0903 ÷ AC00 ÷"},
	{[]rune{0x0903, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0903, 0xAC01}, []int{0, 1, 2}, "÷ 0903 ÷ AC01 ÷"},
	{[]rune{0x0903, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0903, 0x0900}, []int{0, 2}, "÷ 0903 × 0900 ÷"},
	{[]rune{0x0903, 0x0308, 0x0900}, []int{0, 3}, "÷ 0903 × 0308 × 0900 ÷"},
	{[]rune{0x0903, 0x0903}, []int{0, 2}, "÷ 0903 × 0903 ÷"},
	{[]rune{0x0903, 0x0308, 0x0903}, []int{0, 3}, "÷ 0903 × 0308 × 0903 ÷"},
	{[]rune{0x0903, 0x0904}, []int{0, 1, 2}, "÷ 0903 ÷ 0904 ÷"},
	{[]rune{0x0903, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 0904 ÷"},
	{[]rune{0x0903, 0x0D4E}, []int{0, 1, 2}, "÷ 0903 ÷ 0D4E ÷"},
	{[]rune{0x0903, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 0D4E ÷"},
	{[]rune{0x0903, 0x0915}, []int{0, 1, 2}, "÷ 0903 ÷ 0915 ÷"},
	{[]rune{0x0903, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 0915 ÷"},
	{[]rune{0x0903, 0x231A}, []int{0, 1, 2}, "÷ 0903 ÷ 231A ÷"},
	{[]rune{0x0903, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 231A ÷"},
	{[]rune{0x0903, 0x0300}, []int{0, 2}, "÷ 0903 × 0300 ÷"},
	{[]rune{0x0903, 0x0308, 0x0300}, []int{0, 3}, "÷ 0903 × 0308 × 0300 ÷"},
	{[]rune{0x0903, 0x093C}, []int{0, 2}, "÷ 0903 × 093C ÷"},
	{[]rune{0x0903, 0x0308, 0x093C}, []int{0, 3}, "÷ 0903 × 0308 × 093C ÷"},
	{[]rune{0x0903, 0x094D}, []int{0, 2}, "÷ 0903 × 094D ÷"},
	{[]rune{0x0903, 0x0308, 0x094D}, []int{0, 3}, "÷ 0903 × 0308 × 094D ÷"},
	{[]rune{0x0903, 0x200D}, []int{0, 2}, "÷ 0903 × 200D ÷"},
	{[]rune{0x0903, 0x0308, 0x200D}, []int{0, 3}, "÷ 0903 × 0308 × 200D ÷"},
	{[]rune{0x0903, 0x0378}, []int{0, 1, 2}, "÷ 0903 ÷ 0378 ÷"},
	{[]rune{0x0903, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 0903 × 0308 ÷ 0378 ÷"},
	{[]rune{0x0904, 0x0020}, []int{0, 1, 2}, "÷ 0904 ÷ 0020 ÷"},
	{[]rune{0x0904, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0904, 0x000D}, []int{0, 1, 2}, "÷ 0904 ÷ 000D ÷"},
	{[]rune{0x0904, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 000D ÷"},
	{[]rune{0x0904, 0x000A}, []int{0, 1, 2}, "÷ 0904 ÷ 000A ÷"},
	{[]rune{0x0904, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 000A ÷"},
	{[]rune{0x0904, 0x0001}, []int{0, 1, 2}, "÷ 0904 ÷ 0001 ÷"},
	{[]rune{0x0904, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0904, 0x034F}, []int{0, 2}, "÷ 0904 × 034F ÷"},
	{[]rune{0x0904, 0x0308, 0x034F}, []int{0, 3}, "÷ 0904 × 0308 × 034F ÷"},
	{[]rune{0x0904, 0x1F1E6}, []int{0, 1, 2}, "÷ 0904 ÷ 1F1E6 ÷"},
	{[]rune{0x0904, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0904, 0x0600}, []int{0, 1, 2}, "÷ 0904 ÷ 0600 ÷"},
	{[]rune{0x0904, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0904, 0x0A03}, []int{0, 2}, "÷ 0904 × 0A03 ÷"},
	{[]rune{0x0904, 0x0308, 0x0A03}, []int{0, 3}, "÷ 0904 × 0308 × 0A03 ÷"},
	{[]rune{0x0904, 0x1100}, []int{0, 1, 2}, "÷ 0904 ÷ 1100 ÷"},
	{[]rune{0x0904, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0904, 0x1160}, []int{0, 1, 2}, "÷ 0904 ÷ 1160 ÷"},
	{[]rune{0x0904, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0904, 0x11A8}, []int{0, 1, 2}, "÷ 0904 ÷ 11A8 ÷"},
	{[]rune{0x0904, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0904, 0xAC00}, []int{0, 1, 2}, "÷ 0904 ÷ AC00 ÷"},
	{[]rune{0x0904, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0904, 0xAC01}, []int{0, 1, 2}, "÷ 0904 ÷ AC01 ÷"},
	{[]rune{0x0904, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0904, 0x0900}, []int{0, 2}, "÷ 0904 × 0900 ÷"},
	{[]rune{0x0904, 0x0308, 0x0900}, []int{0, 3}, "÷ 0904 × 0308 × 0900 ÷"},
	{[]rune{0x0904, 0x0903}, []int{0, 2}, "÷ 0904 × 0903 ÷"},
	{[]rune{0x0904, 0x0308, 0x0903}, []int{0, 3}, "÷ 0904 × 0308 × 0903 ÷"},
	{[]rune{0x0904, 0x0904}, []int{0, 1, 2}, "÷ 0904 ÷ 0904 ÷"},
	{[]rune{0x0904, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 0904 ÷"},
	{[]rune{0x0904, 0x0D4E}, []int{0, 1, 2}, "÷ 0904 ÷ 0D4E ÷"},
	{[]rune{0x0904, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 0D4E ÷"},
	{[]rune{0x0904, 0x0915}, []int{0, 1, 2}, "÷ 0904 ÷ 0915 ÷"},
	{[]rune{0x0904, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 0915 ÷"},
	{[]rune{0x0904, 0x231A}, []int{0, 1, 2}, "÷ 0904 ÷ 231A ÷"},
	{[]rune{0x0904, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 231A ÷"},
	{[]rune{0x0904, 0x0300}, []int{0, 2}, "÷ 0904 × 0300 ÷"},
	{[]rune{0x0904, 0x0308, 0x0300}, []int{0, 3}, "÷ 0904 × 0308 × 0300 ÷"},
	{[]rune{0x0904, 0x093C}, []int{0, 2}, "÷ 0904 × 093C ÷"},
	{[]rune{0x0904, 0x0308, 0x093C}, []int{0, 3}, "÷ 0904 × 0308 × 093C ÷"},
	{[]rune{0x0904, 0x094D}, []int{0, 2}, "÷ 0904 × 094D ÷"},
	{[]rune{0x0904, 0x0308, 0x094D}, []int{0, 3}, "÷ 0904 × 0308 × 094D ÷"},
	{[]rune{0x0904, 0x200D}, []int{0, 2}, "÷ 0904 × 200D ÷"},
	{[]rune{0x0904, 0x0308, 0x200D}, []int{0, 3}, "÷ 0904 × 0308 × 200D ÷"},
	{[]rune{0x0904, 0x0378}, []int{0, 1, 2}, "÷ 0904 ÷ 0378 ÷"},
	{[]rune{0x0904, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 0904 × 0308 ÷ 0378 ÷"},
	{[]rune{0x0D4E, 0x0020}, []int{0, 2}, "÷ 0D4E × 0020 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 0020 ÷"},
	{[]rune{0x0D4E, 0x000D}, []int{0, 1, 2}, "÷ 0D4E ÷ 000D ÷"},
	{[]rune{0x0D4E, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 000D ÷"},
	{[]rune{0x0D4E, 0x000A}, []int{0, 1, 2}, "÷ 0D4E ÷ 000A ÷"},
	{[]rune{0x0D4E, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 000A ÷"},
	{[]rune{0x0D4E, 0x0001}, []int{0, 1, 2}, "÷ 0D4E ÷ 0001 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 0001 ÷"},
	{[]rune{0x0D4E, 0x034F}, []int{0, 2}, "÷ 0D4E × 034F ÷"},
	{[]rune{0x0D4E, 0x0308, 0x034F}, []int{0, 3}, "÷ 0D4E × 0308 × 034F ÷"},
	{[]rune{0x0D4E, 0x1F1E6}, []int{0, 2}, "÷ 0D4E × 1F1E6 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0D4E, 0x0600}, []int{0, 2}, "÷ 0D4E × 0600 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 0600 ÷"},
	{[]rune{0x0D4E, 0x0A03}, []int{0, 2}, "÷ 0D4E × 0A03 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x0A03}, []int{0, 3}, "÷ 0D4E × 0308 × 0A03 ÷"},
	{[]rune{0x0D4E, 0x1100}, []int{0, 2}, "÷ 0D4E × 1100 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 1100 ÷"},
	{[]rune{0x0D4E, 0x1160}, []int{0, 2}, "÷ 0D4E × 1160 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 1160 ÷"},
	{[]rune{0x0D4E, 0x11A8}, []int{0, 2}, "÷ 0D4E × 11A8 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0D4E, 0xAC00}, []int{0, 2}, "÷ 0D4E × AC00 ÷"},
	{[]rune{0x0D4E, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ AC00 ÷"},
	{[]rune{0x0D4E, 0xAC01}, []int{0, 2}, "÷ 0D4E × AC01 ÷"},
	{[]rune{0x0D4E, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ AC01 ÷"},
	{[]rune{0x0D4E, 0x0900}, []int{0, 2}, "÷ 0D4E × 0900 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x0900}, []int{0, 3}, "÷ 0D4E × 0308 × 0900 ÷"},
	{[]rune{0x0D4E, 0x0903}, []int{0, 2}, "÷ 0D4E × 0903 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x0903}, []int{0, 3}, "÷ 0D4E × 0308 × 0903 ÷"},
	{[]rune{0x0D4E, 0x0904}, []int{0, 2}, "÷ 0D4E × 0904 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 0904 ÷"},
	{[]rune{0x0D4E, 0x0D4E}, []int{0, 2}, "÷ 0D4E × 0D4E ÷"},
	{[]rune{0x0D4E, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 0D4E ÷"},
	{[]rune{0x0D4E, 0x0915}, []int{0, 2}, "÷ 0D4E × 0915 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 0915 ÷"},
	{[]rune{0x0D4E, 0x231A}, []int{0, 2}, "÷ 0D4E × 231A ÷"},
	{[]rune{0x0D4E, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 231A ÷"},
	{[]rune{0x0D4E, 0x0300}, []int{0, 2}, "÷ 0D4E × 0300 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x0300}, []int{0, 3}, "÷ 0D4E × 0308 × 0300 ÷"},
	{[]rune{0x0D4E, 0x093C}, []int{0, 2}, "÷ 0D4E × 093C ÷"},
	{[]rune{0x0D4E, 0x0308, 0x093C}, []int{0, 3}, "÷ 0D4E × 0308 × 093C ÷"},
	{[]rune{0x0D4E, 0x094D}, []int{0, 2}, "÷ 0D4E × 094D ÷"},
	{[]rune{0x0D4E, 0x0308, 0x094D}, []int{0, 3}, "÷ 0D4E × 0308 × 094D ÷"},
	{[]rune{0x0D4E, 0x200D}, []int{0, 2}, "÷ 0D4E × 200D ÷"},
	{[]rune{0x0D4E, 0x0308, 0x200D}, []int{0, 3}, "÷ 0D4E × 0308 × 200D ÷"},
	{[]rune{0x0D4E, 0x0378}, []int{0, 2}, "÷ 0D4E × 0378 ÷"},
	{[]rune{0x0D4E, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 0D4E × 0308 ÷ 0378 ÷"},
	{[]rune{0x0915, 0x0020}, []int{0, 1, 2}, "÷ 0915 ÷ 0020 ÷"},
	{[]rune{0x0915, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0915, 0x000D}, []int{0, 1, 2}, "÷ 0915 ÷ 000D ÷"},
	{[]rune{0x0915, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 000D ÷"},
	{[]rune{0x0915, 0x000A}, []int{0, 1, 2}, "÷ 0915 ÷ 000A ÷"},
	{[]rune{0x0915, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 000A ÷"},
	{[]rune{0x0915, 0x0001}, []int{0, 1, 2}, "÷ 0915 ÷ 0001 ÷"},
	{[]rune{0x0915, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0915, 0x034F}, []int{0, 2}, "÷ 0915 × 034F ÷"},
	{[]rune{0x0915, 0x0308, 0x034F}, []int{0, 3}, "÷ 0915 × 0308 × 034F ÷"},
	{[]rune{0x0915, 0x1F1E6}, []int{0, 1, 2}, "÷ 0915 ÷ 1F1E6 ÷"},
	{[]rune{0x0915, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0915, 0x0600}, []int{0, 1, 2}, "÷ 0915 ÷ 0600 ÷"},
	{[]rune{0x0915, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0915, 0x0A03}, []int{0, 2}, "÷ 0915 × 0A03 ÷"},
	{[]rune{0x0915, 0x0308, 0x0A03}, []int{0, 3}, "÷ 0915 × 0308 × 0A03 ÷"},
	{[]rune{0x0915, 0x1100}, []int{0, 1, 2}, "÷ 0915 ÷ 1100 ÷"},
	{[]rune{0x0915, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0915, 0x1160}, []int{0, 1, 2}, "÷ 0915 ÷ 1160 ÷"},
	{[]rune{0x0915, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0915, 0x11A8}, []int{0, 1, 2}, "÷ 0915 ÷ 11A8 ÷"},
	{[]rune{0x0915, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0915, 0xAC00}, []int{0, 1, 2}, "÷ 0915 ÷ AC00 ÷"},
	{[]rune{0x0915, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0915, 0xAC01}, []int{0, 1, 2}, "÷ 0915 ÷ AC01 ÷"},
	{[]rune{0x0915, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0915, 0x0900}, []int{0, 2}, "÷ 0915 × 0900 ÷"},
	{[]rune{0x0915, 0x0308, 0x0900}, []int{0, 3}, "÷ 0915 × 0308 × 0900 ÷"},
	{[]rune{0x0915, 0x0903}, []int{0, 2}, "÷ 0915 × 0903 ÷"},
	{[]rune{0x0915, 0x0308, 0x0903}, []int{0, 3}, "÷ 0915 × 0308 × 0903 ÷"},
	{[]rune{0x0915, 0x0904}, []int{0, 1, 2}, "÷ 0915 ÷ 0904 ÷"},
	{[]rune{0x0915, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 0904 ÷"},
	{[]rune{0x0915, 0x0D4E}, []int{0, 1, 2}, "÷ 0915 ÷ 0D4E ÷"},
	{[]rune{0x0915, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 0D4E ÷"},
	{[]rune{0x0915, 0x0915}, []int{0, 1, 2}, "÷ 0915 ÷ 0915 ÷"},
	{[]rune{0x0915, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 0915 ÷"},
	{[]rune{0x0915, 0x231A}, []int{0, 1, 2}, "÷ 0915 ÷ 231A ÷"},
	{[]rune{0x0915, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 231A ÷"},
	{[]rune{0x0915, 0x0300}, []int{0, 2}, "÷ 0915 × 0300 ÷"},
	{[]rune{0x0915, 0x0308, 0x0300}, []int{0, 3}, "÷ 0915 × 0308 × 0300 ÷"},
	{[]rune{0x0915, 0x093C}, []int{0, 2}, "÷ 0915 × 093C ÷"},
	{[]rune{0x0915, 0x0308, 0x093C}, []int{0, 3}, "÷ 0915 × 0308 × 093C ÷"},
	{[]rune{0x0915, 0x094D}, []int{0, 2}, "÷ 0915 × 094D ÷"},
	{[]rune{0x0915, 0x0308, 0x094D}, []int{0, 3}, "÷ 0915 × 0308 × 094D ÷"},
	{[]rune{0x0915, 0x200D}, []int{0, 2}, "÷ 0915 × 200D ÷"},
	{[]rune{0x0915, 0x0308, 0x200D}, []int{0, 3}, "÷ 0915 × 0308 × 200D ÷"},
	{[]rune{0x0915, 0x0378}, []int{0, 1, 2}, "÷ 0915 ÷ 0378 ÷"},
	{[]rune{0x0915, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 0915 × 0308 ÷ 0378 ÷"},
	{[]rune{0x231A, 0x0020}, []int{0, 1, 2}, "÷ 231A ÷ 0020 ÷"},
	{[]rune{0x231A, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 0020 ÷"},
	{[]rune{0x231A, 0x000D}, []int{0, 1, 2}, "÷ 231A ÷ 000D ÷"},
	{[]rune{0x231A, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 000D ÷"},
	{[]rune{0x231A, 0x000A}, []int{0, 1, 2}, "÷ 231A ÷ 000A ÷"},
	{[]rune{0x231A, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 000A ÷"},
	{[]rune{0x231A, 0x0001}, []int{0, 1, 2}, "÷ 231A ÷ 0001 ÷"},
	{[]rune{0x231A, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 0001 ÷"},
	{[]rune{0x231A, 0x034F}, []int{0, 2}, "÷ 231A × 034F ÷"},
	{[]rune{0x231A, 0x0308, 0x034F}, []int{0, 3}, "÷ 231A × 0308 × 034F ÷"},
	{[]rune{0x231A, 0x1F1E6}, []int{0, 1, 2}, "÷ 231A ÷ 1F1E6 ÷"},
	{[]rune{0x231A, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x231A, 0x0600}, []int{0, 1, 2}, "÷ 231A ÷ 0600 ÷"},
	{[]rune{0x231A, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 0600 ÷"},
	{[]rune{0x231A, 0x0A03}, []int{0, 2}, "÷ 231A × 0A03 ÷"},
	{[]rune{0x231A, 0x0308, 0x0A03}, []int{0, 3}, "÷ 231A × 0308 × 0A03 ÷"},
	{[]rune{0x231A, 0x1100}, []int{0, 1, 2}, "÷ 231A ÷ 1100 ÷"},
	{[]rune{0x231A, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 1100 ÷"},
	{[]rune{0x231A, 0x1160}, []int{0, 1, 2}, "÷ 231A ÷ 1160 ÷"},
	{[]rune{0x231A, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 1160 ÷"},
	{[]rune{0x231A, 0x11A8}, []int{0, 1, 2}, "÷ 231A ÷ 11A8 ÷"},
	{[]rune{0x231A, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 11A8 ÷"},
	{[]rune{0x231A, 0xAC00}, []int{0, 1, 2}, "÷ 231A ÷ AC00 ÷"},
	{[]rune{0x231A, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ AC00 ÷"},
	{[]rune{0x231A, 0xAC01}, []int{0, 1, 2}, "÷ 231A ÷ AC01 ÷"},
	{[]rune{0x231A, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ AC01 ÷"},
	{[]rune{0x231A, 0x0900}, []int{0, 2}, "÷ 231A × 0900 ÷"},
	{[]rune{0x231A, 0x0308, 0x0900}, []int{0, 3}, "÷ 231A × 0308 × 0900 ÷"},
	{[]rune{0x231A, 0x0903}, []int{0, 2}, "÷ 231A × 0903 ÷"},
	{[]rune{0x231A, 0x0308, 0x0903}, []int{0, 3}, "÷ 231A × 0308 × 0903 ÷"},
	{[]rune{0x231A, 0x0904}, []int{0, 1, 2}, "÷ 231A ÷ 0904 ÷"},
	{[]rune{0x231A, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 0904 ÷"},
	{[]rune{0x231A, 0x0D4E}, []int{0, 1, 2}, "÷ 231A ÷ 0D4E ÷"},
	{[]rune{0x231A, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 0D4E ÷"},
	{[]rune{0x231A, 0x0915}, []int{0, 1, 2}, "÷ 231A ÷ 0915 ÷"},
	{[]rune{0x231A, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 0915 ÷"},
	{[]rune{0x231A, 0x231A}, []int{0, 1, 2}, "÷ 231A ÷ 231A ÷"},
	{[]rune{0x231A, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 231A ÷"},
	{[]rune{0x231A, 0x0300}, []int{0, 2}, "÷ 231A × 0300 ÷"},
	{[]rune{0x231A, 0x0308, 0x0300}, []int{0, 3}, "÷ 231A × 0308 × 0300 ÷"},
	{[]rune{0x231A, 0x093C}, []int{0, 2}, "÷ 231A × 093C ÷"},
	{[]rune{0x231A, 0x0308, 0x093C}, []int{0, 3}, "÷ 231A × 0308 × 093C ÷"},
	{[]rune{0x231A, 0x094D}, []int{0, 2}, "÷ 231A × 094D ÷"},
	{[]rune{0x231A, 0x0308, 0x094D}, []int{0, 3}, "÷ 231A × 0308 × 094D ÷"},
	{[]rune{0x231A, 0x200D}, []int{0, 2}, "÷ 231A × 200D ÷"},
	{[]rune{0x231A, 0x0308, 0x200D}, []int{0, 3}, "÷ 231A × 0308 × 200D ÷"},
	{[]rune{0x231A, 0x0378}, []int{0, 1, 2}, "÷ 231A ÷ 0378 ÷"},
	{[]rune{0x231A, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 231A × 0308 ÷ 0378 ÷"},
	{[]rune{0x0300, 0x0020}, []int{0, 1, 2}, "÷ 0300 ÷ 0020 ÷"},
	{[]rune{0x0300, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0300, 0x000D}, []int{0, 1, 2}, "÷ 0300 ÷ 000D ÷"},
	{[]rune{0x0300, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 000D ÷"},
	{[]rune{0x0300, 0x000A}, []int{0, 1, 2}, "÷ 0300 ÷ 000A ÷"},
	{[]rune{0x0300, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 000A ÷"},
	{[]rune{0x0300, 0x0001}, []int{0, 1, 2}, "÷ 0300 ÷ 0001 ÷"},
	{[]rune{0x0300, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0300, 0x034F}, []int{0, 2}, "÷ 0300 × 034F ÷"},
	{[]rune{0x0300, 0x0308, 0x034F}, []int{0, 3}, "÷ 0300 × 0308 × 034F ÷"},
	{[]rune{0x0300, 0x1F1E6}, []int{0, 1, 2}, "÷ 0300 ÷ 1F1E6 ÷"},
	{[]rune{0x0300, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0300, 0x0600}, []int{0, 1, 2}, "÷ 0300 ÷ 0600 ÷"},
	{[]rune{0x0300, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0300, 0x0A03}, []int{0, 2}, "÷ 0300 × 0A03 ÷"},
	{[]rune{0x0300, 0x0308, 0x0A03}, []int{0, 3}, "÷ 0300 × 0308 × 0A03 ÷"},
	{[]rune{0x0300, 0x1100}, []int{0, 1, 2}, "÷ 0300 ÷ 1100 ÷"},
	{[]rune{0x0300, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0300, 0x1160}, []int{0, 1, 2}, "÷ 0300 ÷ 1160 ÷"},
	{[]rune{0x0300, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0300, 0x11A8}, []int{0, 1, 2}, "÷ 0300 ÷ 11A8 ÷"},
	{[]rune{0x0300, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0300, 0xAC00}, []int{0, 1, 2}, "÷ 0300 ÷ AC00 ÷"},
	{[]rune{0x0300, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0300, 0xAC01}, []int{0, 1, 2}, "÷ 0300 ÷ AC01 ÷"},
	{[]rune{0x0300, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0300, 0x0900}, []int{0, 2}, "÷ 0300 × 0900 ÷"},
	{[]rune{0x0300, 0x0308, 0x0900}, []int{0, 3}, "÷ 0300 × 0308 × 0900 ÷"},
	{[]rune{0x0300, 0x0903}, []int{0, 2}, "÷ 0300 × 0903 ÷"},
	{[]rune{0x0300, 0x0308, 0x0903}, []int{0, 3}, "÷ 0300 × 0308 × 0903 ÷"},
	{[]rune{0x0300, 0x0904}, []int{0, 1, 2}, "÷ 0300 ÷ 0904 ÷"},
	{[]rune{0x0300, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 0904 ÷"},
	{[]rune{0x0300, 0x0D4E}, []int{0, 1, 2}, "÷ 0300 ÷ 0D4E ÷"},
	{[]rune{0x0300, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 0D4E ÷"},
	{[]rune{0x0300, 0x0915}, []int{0, 1, 2}, "÷ 0300 ÷ 0915 ÷"},
	{[]rune{0x0300, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 0915 ÷"},
	{[]rune{0x0300, 0x231A}, []int{0, 1, 2}, "÷ 0300 ÷ 231A ÷"},
	{[]rune{0x0300, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 231A ÷"},
	{[]rune{0x0300, 0x0300}, []int{0, 2}, "÷ 0300 × 0300 ÷"},
	{[]rune{0x0300, 0x0308, 0x0300}, []int{0, 3}, "÷ 0300 × 0308 × 0300 ÷"},
	{[]rune{0x0300, 0x093C}, []int{0, 2}, "÷ 0300 × 093C ÷"},
	{[]rune{0x0300, 0x0308, 0x093C}, []int{0, 3}, "÷ 0300 × 0308 × 093C ÷"},
	{[]rune{0x0300, 0x094D}, []int{0, 2}, "÷ 0300 × 094D ÷"},
	{[]rune{0x0300, 0x0308, 0x094D}, []int{0, 3}, "÷ 0300 × 0308 × 094D ÷"},
	{[]rune{0x0300, 0x200D}, []int{0, 2}, "÷ 0300 × 200D ÷"},
	{[]rune{0x0300, 0x0308, 0x200D}, []int{0, 3}, "÷ 0300 × 0308 × 200D ÷"},
	{[]rune{0x0300, 0x0378}, []int{0, 1, 2}, "÷ 0300 ÷ 0378 ÷"},
	{[]rune{0x0300, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 0300 × 0308 ÷ 0378 ÷"},
	{[]rune{0x093C, 0x0020}, []int{0, 1, 2}, "÷ 093C ÷ 0020 ÷"},
	{[]rune{0x093C, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 0020 ÷"},
	{[]rune{0x093C, 0x000D}, []int{0, 1, 2}, "÷ 093C ÷ 000D ÷"},
	{[]rune{0x093C, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 000D ÷"},
	{[]rune{0x093C, 0x000A}, []int{0, 1, 2}, "÷ 093C ÷ 000A ÷"},
	{[]rune{0x093C, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 000A ÷"},
	{[]rune{0x093C, 0x0001}, []int{0, 1, 2}, "÷ 093C ÷ 0001 ÷"},
	{[]rune{0x093C, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 0001 ÷"},
	{[]rune{0x093C, 0x034F}, []int{0, 2}, "÷ 093C × 034F ÷"},
	{[]rune{0x093C, 0x0308, 0x034F}, []int{0, 3}, "÷ 093C × 0308 × 034F ÷"},
	{[]rune{0x093C, 0x1F1E6}, []int{0, 1, 2}, "÷ 093C ÷ 1F1E6 ÷"},
	{[]rune{0x093C, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x093C, 0x0600}, []int{0, 1, 2}, "÷ 093C ÷ 0600 ÷"},
	{[]rune{0x093C, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 0600 ÷"},
	{[]rune{0x093C, 0x0A03}, []int{0, 2}, "÷ 093C × 0A03 ÷"},
	{[]rune{0x093C, 0x0308, 0x0A03}, []int{0, 3}, "÷ 093C × 0308 × 0A03 ÷"},
	{[]rune{0x093C, 0x1100}, []int{0, 1, 2}, "÷ 093C ÷ 1100 ÷"},
	{[]rune{0x093C, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 1100 ÷"},
	{[]rune{0x093C, 0x1160}, []int{0, 1, 2}, "÷ 093C ÷ 1160 ÷"},
	{[]rune{0x093C, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 1160 ÷"},
	{[]rune{0x093C, 0x11A8}, []int{0, 1, 2}, "÷ 093C ÷ 11A8 ÷"},
	{[]rune{0x093C, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 11A8 ÷"},
	{[]rune{0x093C, 0xAC00}, []int{0, 1, 2}, "÷ 093C ÷ AC00 ÷"},
	{[]rune{0x093C, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ AC00 ÷"},
	{[]rune{0x093C, 0xAC01}, []int{0, 1, 2}, "÷ 093C ÷ AC01 ÷"},
	{[]rune{0x093C, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ AC01 ÷"},
	{[]rune{0x093C, 0x0900}, []int{0, 2}, "÷ 093C × 0900 ÷"},
	{[]rune{0x093C, 0x0308, 0x0900}, []int{0, 3}, "÷ 093C × 0308 × 0900 ÷"},
	{[]rune{0x093C, 0x0903}, []int{0, 2}, "÷ 093C × 0903 ÷"},
	{[]rune{0x093C, 0x0308, 0x0903}, []int{0, 3}, "÷ 093C × 0308 × 0903 ÷"},
	{[]rune{0x093C, 0x0904}, []int{0, 1, 2}, "÷ 093C ÷ 0904 ÷"},
	{[]rune{0x093C, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 0904 ÷"},
	{[]rune{0x093C, 0x0D4E}, []int{0, 1, 2}, "÷ 093C ÷ 0D4E ÷"},
	{[]rune{0x093C, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 0D4E ÷"},
	{[]rune{0x093C, 0x0915}, []int{0, 1, 2}, "÷ 093C ÷ 0915 ÷"},
	{[]rune{0x093C, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 0915 ÷"},
	{[]rune{0x093C, 0x231A}, []int{0, 1, 2}, "÷ 093C ÷ 231A ÷"},
	{[]rune{0x093C, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 231A ÷"},
	{[]rune{0x093C, 0x0300}, []int{0, 2}, "÷ 093C × 0300 ÷"},
	{[]rune{0x093C, 0x0308, 0x0300}, []int{0, 3}, "÷ 093C × 0308 × 0300 ÷"},
	{[]rune{0x093C, 0x093C}, []int{0, 2}, "÷ 093C × 093C ÷"},
	{[]rune{0x093C, 0x0308, 0x093C}, []int{0, 3}, "÷ 093C × 0308 × 093C ÷"},
	{[]rune{0x093C, 0x094D}, []int{0, 2}, "÷ 093C × 094D ÷"},
	{[]rune{0x093C, 0x0308, 0x094D}, []int{0, 3}, "÷ 093C × 0308 × 094D ÷"},
	{[]rune{0x093C, 0x200D}, []int{0, 2}, "÷ 093C × 200D ÷"},
	{[]rune{0x093C, 0x0308, 0x200D}, []int{0, 3}, "÷ 093C × 0308 × 200D ÷"},
	{[]rune{0x093C, 0x0378}, []int{0, 1, 2}, "÷ 093C ÷ 0378 ÷"},
	{[]rune{0x093C, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 093C × 0308 ÷ 0378 ÷"},
	{[]rune{0x094D, 0x0020}, []int{0, 1, 2}, "÷ 094D ÷ 0020 ÷"},
	{[]rune{0x094D, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 0020 ÷"},
	{[]rune{0x094D, 0x000D}, []int{0, 1, 2}, "÷ 094D ÷ 000D ÷"},
	{[]rune{0x094D, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 000D ÷"},
	{[]rune{0x094D, 0x000A}, []int{0, 1, 2}, "÷ 094D ÷ 000A ÷"},
	{[]rune{0x094D, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 000A ÷"},
	{[]rune{0x094D, 0x0001}, []int{0, 1, 2}, "÷ 094D ÷ 0001 ÷"},
	{[]rune{0x094D, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 0001 ÷"},
	{[]rune{0x094D, 0x034F}, []int{0, 2}, "÷ 094D × 034F ÷"},
	{[]rune{0x094D, 0x0308, 0x034F}, []int{0, 3}, "÷ 094D × 0308 × 034F ÷"},
	{[]rune{0x094D, 0x1F1E6}, []int{0, 1, 2}, "÷ 094D ÷ 1F1E6 ÷"},
	{[]rune{0x094D, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x094D, 0x0600}, []int{0, 1, 2}, "÷ 094D ÷ 0600 ÷"},
	{[]rune{0x094D, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 0600 ÷"},
	{[]rune{0x094D, 0x0A03}, []int{0, 2}, "÷ 094D × 0A03 ÷"},
	{[]rune{0x094D, 0x0308, 0x0A03}, []int{0, 3}, "÷ 094D × 0308 × 0A03 ÷"},
	{[]rune{0x094D, 0x1100}, []int{0, 1, 2}, "÷ 094D ÷ 1100 ÷"},
	{[]rune{0x094D, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 1100 ÷"},
	{[]rune{0x094D, 0x1160}, []int{0, 1, 2}, "÷ 094D ÷ 1160 ÷"},
	{[]rune{0x094D, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 1160 ÷"},
	{[]rune{0x094D, 0x11A8}, []int{0, 1, 2}, "÷ 094D ÷ 11A8 ÷"},
	{[]rune{0x094D, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 11A8 ÷"},
	{[]rune{0x094D, 0xAC00}, []int{0, 1, 2}, "÷ 094D ÷ AC00 ÷"},
	{[]rune{0x094D, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ AC00 ÷"},
	{[]rune{0x094D, 0xAC01}, []int{0, 1, 2}, "÷ 094D ÷ AC01 ÷"},
	{[]rune{0x094D, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ AC01 ÷"},
	{[]rune{0x094D, 0x0900}, []int{0, 2}, "÷ 094D × 0900 ÷"},
	{[]rune{0x094D, 0x0308, 0x0900}, []int{0, 3}, "÷ 094D × 0308 × 0900 ÷"},
	{[]rune{0x094D, 0x0903}, []int{0, 2}, "÷ 094D × 0903 ÷"},
	{[]rune{0x094D, 0x0308, 0x0903}, []int{0, 3}, "÷ 094D × 0308 × 0903 ÷"},
	{[]rune{0x094D, 0x0904}, []int{0, 1, 2}, "÷ 094D ÷ 0904 ÷"},
	{[]rune{0x094D, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 0904 ÷"},
	{[]rune{0x094D, 0x0D4E}, []int{0, 1, 2}, "÷ 094D ÷ 0D4E ÷"},
	{[]rune{0x094D, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 0D4E ÷"},
	{[]rune{0x094D, 0x0915}, []int{0, 1, 2}, "÷ 094D ÷ 0915 ÷"},
	{[]rune{0x094D, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 0915 ÷"},
	{[]rune{0x094D, 0x231A}, []int{0, 1, 2}, "÷ 094D ÷ 231A ÷"},
	{[]rune{0x094D, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 231A ÷"},
	{[]rune{0x094D, 0x0300}, []int{0, 2}, "÷ 094D × 0300 ÷"},
	{[]rune{0x094D, 0x0308, 0x0300}, []int{0, 3}, "÷ 094D × 0308 × 0300 ÷"},
	{[]rune{0x094D, 0x093C}, []int{0, 2}, "÷ 094D × 093C ÷"},
	{[]rune{0x094D, 0x0308, 0x093C}, []int{0, 3}, "÷ 094D × 0308 × 093C ÷"},
	{[]rune{0x094D, 0x094D}, []int{0, 2}, "÷ 094D × 094D ÷"},
	{[]rune{0x094D, 0x0308, 0x094D}, []int{0, 3}, "÷ 094D × 0308 × 094D ÷"},
	{[]rune{0x094D, 0x200D}, []int{0, 2}, "÷ 094D × 200D ÷"},
	{[]rune{0x094D, 0x0308, 0x200D}, []int{0, 3}, "÷ 094D × 0308 × 200D ÷"},
	{[]rune{0x094D, 0x0378}, []int{0, 1, 2}, "÷ 094D ÷ 0378 ÷"},
	{[]rune{0x094D, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 094D × 0308 ÷ 0378 ÷"},
	{[]rune{0x200D, 0x0020}, []int{0, 1, 2}, "÷ 200D ÷ 0020 ÷"},
	{[]rune{0x200D, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 0020 ÷"},
	{[]rune{0x200D, 0x000D}, []int{0, 1, 2}, "÷ 200D ÷ 000D ÷"},
	{[]rune{0x200D, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 000D ÷"},
	{[]rune{0x200D, 0x000A}, []int{0, 1, 2}, "÷ 200D ÷ 000A ÷"},
	{[]rune{0x200D, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 000A ÷"},
	{[]rune{0x200D, 0x0001}, []int{0, 1, 2}, "÷ 200D ÷ 0001 ÷"},
	{[]rune{0x200D, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 0001 ÷"},
	{[]rune{0x200D, 0x034F}, []int{0, 2}, "÷ 200D × 034F ÷"},
	{[]rune{0x200D, 0x0308, 0x034F}, []int{0, 3}, "÷ 200D × 0308 × 034F ÷"},
	{[]rune{0x200D, 0x1F1E6}, []int{0, 1, 2}, "÷ 200D ÷ 1F1E6 ÷"},
	{[]rune{0x200D, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x200D, 0x0600}, []int{0, 1, 2}, "÷ 200D ÷ 0600 ÷"},
	{[]rune{0x200D, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 0600 ÷"},
	{[]rune{0x200D, 0x0A03}, []int{0, 2}, "÷ 200D × 0A03 ÷"},
	{[]rune{0x200D, 0x0308, 0x0A03}, []int{0, 3}, "÷ 200D × 0308 × 0A03 ÷"},
	{[]rune{0x200D, 0x1100}, []int{0, 1, 2}, "÷ 200D ÷ 1100 ÷"},
	{[]rune{0x200D, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 1100 ÷"},
	{[]rune{0x200D, 0x1160}, []int{0, 1, 2}, "÷ 200D ÷ 1160 ÷"},
	{[]rune{0x200D, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 1160 ÷"},
	{[]rune{0x200D, 0x11A8}, []int{0, 1, 2}, "÷ 200D ÷ 11A8 ÷"},
	{[]rune{0x200D, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 11A8 ÷"},
	{[]rune{0x200D, 0xAC00}, []int{0, 1, 2}, "÷ 200D ÷ AC00 ÷"},
	{[]rune{0x200D, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ AC00 ÷"},
	{[]rune{0x200D, 0xAC01}, []int{0, 1, 2}, "÷ 200D ÷ AC01 ÷"},
	{[]rune{0x200D, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ AC01 ÷"},
	{[]rune{0x200D, 0x0900}, []int{0, 2}, "÷ 200D × 0900 ÷"},
	{[]rune{0x200D, 0x0308, 0x0900}, []int{0, 3}, "÷ 200D × 0308 × 0900 ÷"},
	{[]rune{0x200D, 0x0903}, []int{0, 2}, "÷ 200D × 0903 ÷"},
	{[]rune{0x200D, 0x0308, 0x0903}, []int{0, 3}, "÷ 200D × 0308 × 0903 ÷"},
	{[]rune{0x200D, 0x0904}, []int{0, 1, 2}, "÷ 200D ÷ 0904 ÷"},
	{[]rune{0x200D, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 0904 ÷"},
	{[]rune{0x200D, 0x0D4E}, []int{0, 1, 2}, "÷ 200D ÷ 0D4E ÷"},
	{[]rune{0x200D, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 0D4E ÷"},
	{[]rune{0x200D, 0x0915}, []int{0, 1, 2}, "÷ 200D ÷ 0915 ÷"},
	{[]rune{0x200D, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 0915 ÷"},
	{[]rune{0x200D, 0x231A}, []int{0, 1, 2}, "÷ 200D ÷ 231A ÷"},
	{[]rune{0x200D, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 231A ÷"},
	{[]rune{0x200D, 0x0300}, []int{0, 2}, "÷ 200D × 0300 ÷"},
	{[]rune{0x200D, 0x0308, 0x0300}, []int{0, 3}, "÷ 200D × 0308 × 0300 ÷"},
	{[]rune{0x200D, 0x093C}, []int{0, 2}, "÷ 200D × 093C ÷"},
	{[]rune{0x200D, 0x0308, 0x093C}, []int{0, 3}, "÷ 200D × 0308 × 093C ÷"},
	{[]rune{0x200D, 0x094D}, []int{0, 2}, "÷ 200D × 094D ÷"},
	{[]rune{0x200D, 0x0308, 0x094D}, []int{0, 3}, "÷ 200D × 0308 × 094D ÷"},
	{[]rune{0x200D, 0x200D}, []int{0, 2}, "÷ 200D × 200D ÷"},
	{[]rune{0x200D, 0x0308, 0x200D}, []int{0, 3}, "÷ 200D × 0308 × 200D ÷"},
	{[]rune{0x200D, 0x0378}, []int{0, 1, 2}, "÷ 200D ÷ 0378 ÷"},
	{[]rune{0x200D, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 200D × 0308 ÷ 0378 ÷"},
	{[]rune{0x0378, 0x0020}, []int{0, 1, 2}, "÷ 0378 ÷ 0020 ÷"},
	{[]rune{0x0378, 0x0308, 0x0020}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 0020 ÷"},
	{[]rune{0x0378, 0x000D}, []int{0, 1, 2}, "÷ 0378 ÷ 000D ÷"},
	{[]rune{0x0378, 0x0308, 0x000D}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 000D ÷"},
	{[]rune{0x0378, 0x000A}, []int{0, 1, 2}, "÷ 0378 ÷ 000A ÷"},
	{[]rune{0x0378, 0x0308, 0x000A}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 000A ÷"},
	{[]rune{0x0378, 0x0001}, []int{0, 1, 2}, "÷ 0378 ÷ 0001 ÷"},
	{[]rune{0x0378, 0x0308, 0x0001}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 0001 ÷"},
	{[]rune{0x0378, 0x034F}, []int{0, 2}, "÷ 0378 × 034F ÷"},
	{[]rune{0x0378, 0x0308, 0x034F}, []int{0, 3}, "÷ 0378 × 0308 × 034F ÷"},
	{[]rune{0x0378, 0x1F1E6}, []int{0, 1, 2}, "÷ 0378 ÷ 1F1E6 ÷"},
	{[]rune{0x0378, 0x0308, 0x1F1E6}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 1F1E6 ÷"},
	{[]rune{0x0378, 0x0600}, []int{0, 1, 2}, "÷ 0378 ÷ 0600 ÷"},
	{[]rune{0x0378, 0x0308, 0x0600}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 0600 ÷"},
	{[]rune{0x0378, 0x0A03}, []int{0, 2}, "÷ 0378 × 0A03 ÷"},
	{[]rune{0x0378, 0x0308, 0x0A03}, []int{0, 3}, "÷ 0378 × 0308 × 0A03 ÷"},
	{[]rune{0x0378, 0x1100}, []int{0, 1, 2}, "÷ 0378 ÷ 1100 ÷"},
	{[]rune{0x0378, 0x0308, 0x1100}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 1100 ÷"},
	{[]rune{0x0378, 0x1160}, []int{0, 1, 2}, "÷ 0378 ÷ 1160 ÷"},
	{[]rune{0x0378, 0x0308, 0x1160}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 1160 ÷"},
	{[]rune{0x0378, 0x11A8}, []int{0, 1, 2}, "÷ 0378 ÷ 11A8 ÷"},
	{[]rune{0x0378, 0x0308, 0x11A8}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 11A8 ÷"},
	{[]rune{0x0378, 0xAC00}, []int{0, 1, 2}, "÷ 0378 ÷ AC00 ÷"},
	{[]rune{0x0378, 0x0308, 0xAC00}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ AC00 ÷"},
	{[]rune{0x0378, 0xAC01}, []int{0, 1, 2}, "÷ 0378 ÷ AC01 ÷"},
	{[]rune{0x0378, 0x0308, 0xAC01}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ AC01 ÷"},
	{[]rune{0x0378, 0x0900}, []int{0, 2}, "÷ 0378 × 0900 ÷"},
	{[]rune{0x0378, 0x0308, 0x0900}, []int{0, 3}, "÷ 0378 × 0308 × 0900 ÷"},
	{[]rune{0x0378, 0x0903}, []int{0, 2}, "÷ 0378 × 0903 ÷"},
	{[]rune{0x0378, 0x0308, 0x0903}, []int{0, 3}, "÷ 0378 × 0308 × 0903 ÷"},
	{[]rune{0x0378, 0x0904}, []int{0, 1, 2}, "÷ 0378 ÷ 0904 ÷"},
	{[]rune{0x0378, 0x0308, 0x0904}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 0904 ÷"},
	{[]rune{0x0378, 0x0D4E}, []int{0, 1, 2}, "÷ 0378 ÷ 0D4E ÷"},
	{[]rune{0x0378, 0x0308, 0x0D4E}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 0D4E ÷"},
	{[]rune{0x0378, 0x0915}, []int{0, 1, 2}, "÷ 0378 ÷ 0915 ÷"},
	{[]rune{0x0378, 0x0308, 0x0915}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 0915 ÷"},
	{[]rune{0x0378, 0x231A}, []int{0, 1, 2}, "÷ 0378 ÷ 231A ÷"},
	{[]rune{0x0378, 0x0308, 0x231A}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 231A ÷"},
	{[]rune{0x0378, 0x0300}, []int{0, 2}, "÷ 0378 × 0300 ÷"},
	{[]rune{0x0378, 0x0308, 0x0300}, []int{0, 3}, "÷ 0378 × 0308 × 0300 ÷"},
	{[]rune{0x0378, 0x093C}, []int{0, 2}, "÷ 0378 × 093C ÷"},
	{[]rune{0x0378, 0x0308, 0x093C}, []int{0, 3}, "÷ 0378 × 0308 × 093C ÷"},
	{[]rune{0x0378, 0x094D}, []int{0, 2}, "÷ 0378 × 094D ÷"},
	{[]rune{0x0378, 0x0308, 0x094D}, []int{0, 3}, "÷ 0378 × 0308 × 094D ÷"},
	{[]rune{0x0378, 0x200D}, []int{0, 2}, "÷ 0378 × 200D ÷"},
	{[]rune{0x0378, 0x0308, 0x200D}, []int{0, 3}, "÷ 0378 × 0308 × 200D ÷"},
	{[]rune{0x0378, 0x0378}, []int{0, 1, 2}, "÷ 0378 ÷ 0378 ÷"},
	{[]rune{0x0378, 0x0308, 0x0378}, []int{0, 2, 3}, "÷ 0378 × 0308 ÷ 0378 ÷"},
	{[]rune{0x000D, 0x000A, 0x0061, 0x000A, 0x0308}, []int{0, 2, 3, 4, 5}, "÷ 000D × 000A ÷ 0061 ÷ 000A ÷ 0308 ÷"},
	{[]rune{0x0061, 0x0308}, []int{0, 2}, "÷ 0061 × 0308 ÷"},
	{[]rune{0x0020, 0x200D, 0x0646}, []int{0, 2, 3}, "÷ 0020 × 200D ÷ 0646 ÷"},
	{[]rune{0x0646, 0x200D, 0x0020}, []int{0, 2, 3}, "÷ 0646 × 200D ÷ 0020 ÷"},
	{[]rune{0x1100, 0x1100}, []int{0, 2}, "÷ 1100 × 1100 ÷"},
	{[]rune{0xAC00, 0x11A8, 0x1100}, []int{0, 2, 3}, "÷ AC00 × 11A8 ÷ 1100 ÷"},
	{[]rune{0xAC01, 0x11A8, 0x1100}, []int{0, 2, 3}, "÷ AC01 × 11A8 ÷ 1100 ÷"},
	{[]rune{0x1F1E6, 0x1F1E7, 0x1F1E8, 0x0062}, []int{0, 2, 3, 4}, "÷ 1F1E6 × 1F1E7 ÷ 1F1E8 ÷ 0062 ÷"},
	{[]rune{0x0061, 0x1F1E6, 0x1F1E7, 0x1F1E8, 0x0062}, []int{0, 1, 3, 4, 5}, "÷ 0061 ÷ 1F1E6 × 1F1E7 ÷ 1F1E8 ÷ 0062 ÷"},
	{[]rune{0x0061, 0x1F1E6, 0x1F1E7, 0x200D, 0x1F1E8, 0x0062}, []int{0, 1, 4, 5, 6}, "÷ 0061 ÷ 1F1E6 × 1F1E7 × 200D ÷ 1F1E8 ÷ 0062 ÷"},
	{[]rune{0x0061, 0x1F1E6, 0x200D, 0x1F1E7, 0x1F1E8, 0x0062}, []int{0, 1, 3, 5, 6}, "÷ 0061 ÷ 1F1E6 × 200D ÷ 1F1E7 × 1F1E8 ÷ 0062 ÷"},
	{[]rune{0x0061, 0x1F1E6, 0x1F1E7, 0x1F1E8, 0x1F1E9, 0x0062}, []int{0, 1, 3, 5, 6}, "÷ 0061 ÷ 1F1E6 × 1F1E7 ÷ 1F1E8 × 1F1E9 ÷ 0062 ÷"},
	{[]rune{0x0061, 0x200D}, []int{0, 2}, "÷ 0061 × 200D ÷"},
	{[]rune{0x0061, 0x0308, 0x0062}, []int{0, 2, 3}, "÷ 0061 × 0308 ÷ 0062 ÷"},
	{[]rune{0x0061, 0x0903, 0x0062}, []int{0, 2, 3}, "÷ 0061 × 0903 ÷ 0062 ÷"},
	{[]rune{0x0061, 0x0600, 0x0062}, []int{0, 1, 3}, "÷ 0061 ÷ 0600 × 0062 ÷"},
	{[]rune{0x1F476, 0x1F3FF, 0x1F476}, []int{0, 2, 3}, "÷ 1F476 × 1F3FF ÷ 1F476 ÷"},
	{[]rune{0x0061, 0x1F3FF, 0x1F476}, []int{0, 2, 3}, "÷ 0061 × 1F3FF ÷ 1F476 ÷"},
	{[]rune{0x0061, 0x1F3FF, 0x1F476, 0x200D, 0x1F6D1}, []int{0, 2, 5}, "÷ 0061 × 1F3FF ÷ 1F476 × 200D × 1F6D1 ÷"},
	{[]rune{0x1F476, 0x1F3FF, 0x0308, 0x200D, 0x1F476, 0x1F3FF}, []int{0, 6}, "÷ 1F476 × 1F3FF × 0308 × 200D × 1F476 × 1F3FF ÷"},
	{[]rune{0x1F6D1, 0x200D, 0x1F6D1}, []int{0, 3}, "÷ 1F6D1 × 200D × 1F6D1 ÷"},
	{[]rune{0x0061, 0x200D, 0x1F6D1}, []int{0, 2, 3}, "÷ 0061 × 200D ÷ 1F6D1 ÷"},
	{[]rune{0x2701, 0x200D, 0x2701}, []int{0, 3}, "÷ 2701 × 200D × 2701 ÷"},
	{[]rune{0x0061, 0x200D, 0x2701}, []int{0, 2, 3}, "÷ 0061 × 200D ÷ 2701 ÷"},
	{[]rune{0x0915, 0x0924}, []int{0, 1, 2}, "÷ 0915 ÷ 0924 ÷"},
	{[]rune{0x0915, 0x094D, 0x0924}, []int{0, 3}, "÷ 0915 × 094D × 0924 ÷"},
	{[]rune{0x0915, 0x094D, 0x094D, 0x0924}, []int{0, 4}, "÷ 0915 × 094D × 094D × 0924 ÷"},
	{[]rune{0x0915, 0x094D, 0x200D, 0x0924}, []int{0, 4}, "÷ 0915 × 094D × 200D × 0924 ÷"},
	{[]rune{0x0915, 0x093C, 0x200D, 0x094D, 0x0924}, []int{0, 5}, "÷ 0915 × 093C × 200D × 094D × 0924 ÷"},
	{[]rune{0x0915, 0x093C, 0x094D, 0x200D, 0x0924}, []int{0, 5}, "÷ 0915 × 093C × 094D × 200D × 0924 ÷"},
	{[]rune{0x0915, 0x094D, 0x0924, 0x094D, 0x092F}, []int{0, 5}, "÷ 0915 × 094D × 0924 × 094D × 092F ÷"},
	{[]rune{0x0915, 0x094D, 0x0061}, []int{0, 2, 3}, "÷ 0915 × 094D ÷ 0061 ÷"},
	{[]rune{0x0061, 0x094D, 0x0924}, []int{0, 2, 3}, "÷ 0061 × 094D ÷ 0924 ÷"},
	{[]rune{0x003F, 0x094D, 0x0924}, []int{0, 2, 3}, "÷ 003F × 094D ÷ 0924 ÷"},
	{[]rune{0x0915, 0x094D, 0x094D, 0x0924}, []int{0, 4}, "÷ 0915 × 094D × 094D × 0924 ÷"},
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
