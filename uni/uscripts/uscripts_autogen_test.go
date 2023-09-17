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

package uscripts_test

import "testing"
import "github.com/orkvozku/go/uni/uscripts"

// autogenTestCaseFunc contains a function to test.
type autogenTestCaseFunc struct {
	label string
	f     func(rune) bool
}

// autogenTestCasesFunc... describes a function to test.
var (
	autogenTestCasesFuncScxAdlm = autogenTestCaseFunc{"HasScxAdlm", uscripts.HasScxAdlm}
	autogenTestCasesFuncScxAghb = autogenTestCaseFunc{"HasScxAghb", uscripts.HasScxAghb}
	autogenTestCasesFuncScxAhom = autogenTestCaseFunc{"HasScxAhom", uscripts.HasScxAhom}
	autogenTestCasesFuncScxArab = autogenTestCaseFunc{"HasScxArab", uscripts.HasScxArab}
	autogenTestCasesFuncScxArmi = autogenTestCaseFunc{"HasScxArmi", uscripts.HasScxArmi}
	autogenTestCasesFuncScxArmn = autogenTestCaseFunc{"HasScxArmn", uscripts.HasScxArmn}
	autogenTestCasesFuncScxAvst = autogenTestCaseFunc{"HasScxAvst", uscripts.HasScxAvst}
	autogenTestCasesFuncScxBali = autogenTestCaseFunc{"HasScxBali", uscripts.HasScxBali}
	autogenTestCasesFuncScxBamu = autogenTestCaseFunc{"HasScxBamu", uscripts.HasScxBamu}
	autogenTestCasesFuncScxBass = autogenTestCaseFunc{"HasScxBass", uscripts.HasScxBass}
	autogenTestCasesFuncScxBatk = autogenTestCaseFunc{"HasScxBatk", uscripts.HasScxBatk}
	autogenTestCasesFuncScxBeng = autogenTestCaseFunc{"HasScxBeng", uscripts.HasScxBeng}
	autogenTestCasesFuncScxBhks = autogenTestCaseFunc{"HasScxBhks", uscripts.HasScxBhks}
	autogenTestCasesFuncScxBopo = autogenTestCaseFunc{"HasScxBopo", uscripts.HasScxBopo}
	autogenTestCasesFuncScxBrah = autogenTestCaseFunc{"HasScxBrah", uscripts.HasScxBrah}
	autogenTestCasesFuncScxBrai = autogenTestCaseFunc{"HasScxBrai", uscripts.HasScxBrai}
	autogenTestCasesFuncScxBugi = autogenTestCaseFunc{"HasScxBugi", uscripts.HasScxBugi}
	autogenTestCasesFuncScxBuhd = autogenTestCaseFunc{"HasScxBuhd", uscripts.HasScxBuhd}
	autogenTestCasesFuncScxCakm = autogenTestCaseFunc{"HasScxCakm", uscripts.HasScxCakm}
	autogenTestCasesFuncScxCans = autogenTestCaseFunc{"HasScxCans", uscripts.HasScxCans}
	autogenTestCasesFuncScxCari = autogenTestCaseFunc{"HasScxCari", uscripts.HasScxCari}
	autogenTestCasesFuncScxCham = autogenTestCaseFunc{"HasScxCham", uscripts.HasScxCham}
	autogenTestCasesFuncScxCher = autogenTestCaseFunc{"HasScxCher", uscripts.HasScxCher}
	autogenTestCasesFuncScxChrs = autogenTestCaseFunc{"HasScxChrs", uscripts.HasScxChrs}
	autogenTestCasesFuncScxCopt = autogenTestCaseFunc{"HasScxCopt", uscripts.HasScxCopt}
	autogenTestCasesFuncScxCpmn = autogenTestCaseFunc{"HasScxCpmn", uscripts.HasScxCpmn}
	autogenTestCasesFuncScxCprt = autogenTestCaseFunc{"HasScxCprt", uscripts.HasScxCprt}
	autogenTestCasesFuncScxCyrl = autogenTestCaseFunc{"HasScxCyrl", uscripts.HasScxCyrl}
	autogenTestCasesFuncScxDeva = autogenTestCaseFunc{"HasScxDeva", uscripts.HasScxDeva}
	autogenTestCasesFuncScxDiak = autogenTestCaseFunc{"HasScxDiak", uscripts.HasScxDiak}
	autogenTestCasesFuncScxDogr = autogenTestCaseFunc{"HasScxDogr", uscripts.HasScxDogr}
	autogenTestCasesFuncScxDsrt = autogenTestCaseFunc{"HasScxDsrt", uscripts.HasScxDsrt}
	autogenTestCasesFuncScxDupl = autogenTestCaseFunc{"HasScxDupl", uscripts.HasScxDupl}
	autogenTestCasesFuncScxEgyp = autogenTestCaseFunc{"HasScxEgyp", uscripts.HasScxEgyp}
	autogenTestCasesFuncScxElba = autogenTestCaseFunc{"HasScxElba", uscripts.HasScxElba}
	autogenTestCasesFuncScxElym = autogenTestCaseFunc{"HasScxElym", uscripts.HasScxElym}
	autogenTestCasesFuncScxEthi = autogenTestCaseFunc{"HasScxEthi", uscripts.HasScxEthi}
	autogenTestCasesFuncScxGeor = autogenTestCaseFunc{"HasScxGeor", uscripts.HasScxGeor}
	autogenTestCasesFuncScxGlag = autogenTestCaseFunc{"HasScxGlag", uscripts.HasScxGlag}
	autogenTestCasesFuncScxGong = autogenTestCaseFunc{"HasScxGong", uscripts.HasScxGong}
	autogenTestCasesFuncScxGonm = autogenTestCaseFunc{"HasScxGonm", uscripts.HasScxGonm}
	autogenTestCasesFuncScxGoth = autogenTestCaseFunc{"HasScxGoth", uscripts.HasScxGoth}
	autogenTestCasesFuncScxGran = autogenTestCaseFunc{"HasScxGran", uscripts.HasScxGran}
	autogenTestCasesFuncScxGrek = autogenTestCaseFunc{"HasScxGrek", uscripts.HasScxGrek}
	autogenTestCasesFuncScxGujr = autogenTestCaseFunc{"HasScxGujr", uscripts.HasScxGujr}
	autogenTestCasesFuncScxGuru = autogenTestCaseFunc{"HasScxGuru", uscripts.HasScxGuru}
	autogenTestCasesFuncScxHang = autogenTestCaseFunc{"HasScxHang", uscripts.HasScxHang}
	autogenTestCasesFuncScxHani = autogenTestCaseFunc{"HasScxHani", uscripts.HasScxHani}
	autogenTestCasesFuncScxHano = autogenTestCaseFunc{"HasScxHano", uscripts.HasScxHano}
	autogenTestCasesFuncScxHatr = autogenTestCaseFunc{"HasScxHatr", uscripts.HasScxHatr}
	autogenTestCasesFuncScxHebr = autogenTestCaseFunc{"HasScxHebr", uscripts.HasScxHebr}
	autogenTestCasesFuncScxHira = autogenTestCaseFunc{"HasScxHira", uscripts.HasScxHira}
	autogenTestCasesFuncScxHluw = autogenTestCaseFunc{"HasScxHluw", uscripts.HasScxHluw}
	autogenTestCasesFuncScxHmng = autogenTestCaseFunc{"HasScxHmng", uscripts.HasScxHmng}
	autogenTestCasesFuncScxHmnp = autogenTestCaseFunc{"HasScxHmnp", uscripts.HasScxHmnp}
	autogenTestCasesFuncScxHung = autogenTestCaseFunc{"HasScxHung", uscripts.HasScxHung}
	autogenTestCasesFuncScxItal = autogenTestCaseFunc{"HasScxItal", uscripts.HasScxItal}
	autogenTestCasesFuncScxJava = autogenTestCaseFunc{"HasScxJava", uscripts.HasScxJava}
	autogenTestCasesFuncScxKali = autogenTestCaseFunc{"HasScxKali", uscripts.HasScxKali}
	autogenTestCasesFuncScxKana = autogenTestCaseFunc{"HasScxKana", uscripts.HasScxKana}
	autogenTestCasesFuncScxKawi = autogenTestCaseFunc{"HasScxKawi", uscripts.HasScxKawi}
	autogenTestCasesFuncScxKhar = autogenTestCaseFunc{"HasScxKhar", uscripts.HasScxKhar}
	autogenTestCasesFuncScxKhmr = autogenTestCaseFunc{"HasScxKhmr", uscripts.HasScxKhmr}
	autogenTestCasesFuncScxKhoj = autogenTestCaseFunc{"HasScxKhoj", uscripts.HasScxKhoj}
	autogenTestCasesFuncScxKits = autogenTestCaseFunc{"HasScxKits", uscripts.HasScxKits}
	autogenTestCasesFuncScxKnda = autogenTestCaseFunc{"HasScxKnda", uscripts.HasScxKnda}
	autogenTestCasesFuncScxKthi = autogenTestCaseFunc{"HasScxKthi", uscripts.HasScxKthi}
	autogenTestCasesFuncScxLana = autogenTestCaseFunc{"HasScxLana", uscripts.HasScxLana}
	autogenTestCasesFuncScxLaoo = autogenTestCaseFunc{"HasScxLaoo", uscripts.HasScxLaoo}
	autogenTestCasesFuncScxLatn = autogenTestCaseFunc{"HasScxLatn", uscripts.HasScxLatn}
	autogenTestCasesFuncScxLepc = autogenTestCaseFunc{"HasScxLepc", uscripts.HasScxLepc}
	autogenTestCasesFuncScxLimb = autogenTestCaseFunc{"HasScxLimb", uscripts.HasScxLimb}
	autogenTestCasesFuncScxLina = autogenTestCaseFunc{"HasScxLina", uscripts.HasScxLina}
	autogenTestCasesFuncScxLinb = autogenTestCaseFunc{"HasScxLinb", uscripts.HasScxLinb}
	autogenTestCasesFuncScxLisu = autogenTestCaseFunc{"HasScxLisu", uscripts.HasScxLisu}
	autogenTestCasesFuncScxLyci = autogenTestCaseFunc{"HasScxLyci", uscripts.HasScxLyci}
	autogenTestCasesFuncScxLydi = autogenTestCaseFunc{"HasScxLydi", uscripts.HasScxLydi}
	autogenTestCasesFuncScxMahj = autogenTestCaseFunc{"HasScxMahj", uscripts.HasScxMahj}
	autogenTestCasesFuncScxMaka = autogenTestCaseFunc{"HasScxMaka", uscripts.HasScxMaka}
	autogenTestCasesFuncScxMand = autogenTestCaseFunc{"HasScxMand", uscripts.HasScxMand}
	autogenTestCasesFuncScxMani = autogenTestCaseFunc{"HasScxMani", uscripts.HasScxMani}
	autogenTestCasesFuncScxMarc = autogenTestCaseFunc{"HasScxMarc", uscripts.HasScxMarc}
	autogenTestCasesFuncScxMedf = autogenTestCaseFunc{"HasScxMedf", uscripts.HasScxMedf}
	autogenTestCasesFuncScxMend = autogenTestCaseFunc{"HasScxMend", uscripts.HasScxMend}
	autogenTestCasesFuncScxMerc = autogenTestCaseFunc{"HasScxMerc", uscripts.HasScxMerc}
	autogenTestCasesFuncScxMero = autogenTestCaseFunc{"HasScxMero", uscripts.HasScxMero}
	autogenTestCasesFuncScxMlym = autogenTestCaseFunc{"HasScxMlym", uscripts.HasScxMlym}
	autogenTestCasesFuncScxModi = autogenTestCaseFunc{"HasScxModi", uscripts.HasScxModi}
	autogenTestCasesFuncScxMong = autogenTestCaseFunc{"HasScxMong", uscripts.HasScxMong}
	autogenTestCasesFuncScxMroo = autogenTestCaseFunc{"HasScxMroo", uscripts.HasScxMroo}
	autogenTestCasesFuncScxMtei = autogenTestCaseFunc{"HasScxMtei", uscripts.HasScxMtei}
	autogenTestCasesFuncScxMult = autogenTestCaseFunc{"HasScxMult", uscripts.HasScxMult}
	autogenTestCasesFuncScxMymr = autogenTestCaseFunc{"HasScxMymr", uscripts.HasScxMymr}
	autogenTestCasesFuncScxNagm = autogenTestCaseFunc{"HasScxNagm", uscripts.HasScxNagm}
	autogenTestCasesFuncScxNand = autogenTestCaseFunc{"HasScxNand", uscripts.HasScxNand}
	autogenTestCasesFuncScxNarb = autogenTestCaseFunc{"HasScxNarb", uscripts.HasScxNarb}
	autogenTestCasesFuncScxNbat = autogenTestCaseFunc{"HasScxNbat", uscripts.HasScxNbat}
	autogenTestCasesFuncScxNewa = autogenTestCaseFunc{"HasScxNewa", uscripts.HasScxNewa}
	autogenTestCasesFuncScxNkoo = autogenTestCaseFunc{"HasScxNkoo", uscripts.HasScxNkoo}
	autogenTestCasesFuncScxNshu = autogenTestCaseFunc{"HasScxNshu", uscripts.HasScxNshu}
	autogenTestCasesFuncScxOgam = autogenTestCaseFunc{"HasScxOgam", uscripts.HasScxOgam}
	autogenTestCasesFuncScxOlck = autogenTestCaseFunc{"HasScxOlck", uscripts.HasScxOlck}
	autogenTestCasesFuncScxOrkh = autogenTestCaseFunc{"HasScxOrkh", uscripts.HasScxOrkh}
	autogenTestCasesFuncScxOrya = autogenTestCaseFunc{"HasScxOrya", uscripts.HasScxOrya}
	autogenTestCasesFuncScxOsge = autogenTestCaseFunc{"HasScxOsge", uscripts.HasScxOsge}
	autogenTestCasesFuncScxOsma = autogenTestCaseFunc{"HasScxOsma", uscripts.HasScxOsma}
	autogenTestCasesFuncScxOugr = autogenTestCaseFunc{"HasScxOugr", uscripts.HasScxOugr}
	autogenTestCasesFuncScxPalm = autogenTestCaseFunc{"HasScxPalm", uscripts.HasScxPalm}
	autogenTestCasesFuncScxPauc = autogenTestCaseFunc{"HasScxPauc", uscripts.HasScxPauc}
	autogenTestCasesFuncScxPerm = autogenTestCaseFunc{"HasScxPerm", uscripts.HasScxPerm}
	autogenTestCasesFuncScxPhag = autogenTestCaseFunc{"HasScxPhag", uscripts.HasScxPhag}
	autogenTestCasesFuncScxPhli = autogenTestCaseFunc{"HasScxPhli", uscripts.HasScxPhli}
	autogenTestCasesFuncScxPhlp = autogenTestCaseFunc{"HasScxPhlp", uscripts.HasScxPhlp}
	autogenTestCasesFuncScxPhnx = autogenTestCaseFunc{"HasScxPhnx", uscripts.HasScxPhnx}
	autogenTestCasesFuncScxPlrd = autogenTestCaseFunc{"HasScxPlrd", uscripts.HasScxPlrd}
	autogenTestCasesFuncScxPrti = autogenTestCaseFunc{"HasScxPrti", uscripts.HasScxPrti}
	autogenTestCasesFuncScxRjng = autogenTestCaseFunc{"HasScxRjng", uscripts.HasScxRjng}
	autogenTestCasesFuncScxRohg = autogenTestCaseFunc{"HasScxRohg", uscripts.HasScxRohg}
	autogenTestCasesFuncScxRunr = autogenTestCaseFunc{"HasScxRunr", uscripts.HasScxRunr}
	autogenTestCasesFuncScxSamr = autogenTestCaseFunc{"HasScxSamr", uscripts.HasScxSamr}
	autogenTestCasesFuncScxSarb = autogenTestCaseFunc{"HasScxSarb", uscripts.HasScxSarb}
	autogenTestCasesFuncScxSaur = autogenTestCaseFunc{"HasScxSaur", uscripts.HasScxSaur}
	autogenTestCasesFuncScxSgnw = autogenTestCaseFunc{"HasScxSgnw", uscripts.HasScxSgnw}
	autogenTestCasesFuncScxShaw = autogenTestCaseFunc{"HasScxShaw", uscripts.HasScxShaw}
	autogenTestCasesFuncScxShrd = autogenTestCaseFunc{"HasScxShrd", uscripts.HasScxShrd}
	autogenTestCasesFuncScxSidd = autogenTestCaseFunc{"HasScxSidd", uscripts.HasScxSidd}
	autogenTestCasesFuncScxSind = autogenTestCaseFunc{"HasScxSind", uscripts.HasScxSind}
	autogenTestCasesFuncScxSinh = autogenTestCaseFunc{"HasScxSinh", uscripts.HasScxSinh}
	autogenTestCasesFuncScxSogd = autogenTestCaseFunc{"HasScxSogd", uscripts.HasScxSogd}
	autogenTestCasesFuncScxSogo = autogenTestCaseFunc{"HasScxSogo", uscripts.HasScxSogo}
	autogenTestCasesFuncScxSora = autogenTestCaseFunc{"HasScxSora", uscripts.HasScxSora}
	autogenTestCasesFuncScxSoyo = autogenTestCaseFunc{"HasScxSoyo", uscripts.HasScxSoyo}
	autogenTestCasesFuncScxSund = autogenTestCaseFunc{"HasScxSund", uscripts.HasScxSund}
	autogenTestCasesFuncScxSylo = autogenTestCaseFunc{"HasScxSylo", uscripts.HasScxSylo}
	autogenTestCasesFuncScxSyrc = autogenTestCaseFunc{"HasScxSyrc", uscripts.HasScxSyrc}
	autogenTestCasesFuncScxTagb = autogenTestCaseFunc{"HasScxTagb", uscripts.HasScxTagb}
	autogenTestCasesFuncScxTakr = autogenTestCaseFunc{"HasScxTakr", uscripts.HasScxTakr}
	autogenTestCasesFuncScxTale = autogenTestCaseFunc{"HasScxTale", uscripts.HasScxTale}
	autogenTestCasesFuncScxTalu = autogenTestCaseFunc{"HasScxTalu", uscripts.HasScxTalu}
	autogenTestCasesFuncScxTaml = autogenTestCaseFunc{"HasScxTaml", uscripts.HasScxTaml}
	autogenTestCasesFuncScxTang = autogenTestCaseFunc{"HasScxTang", uscripts.HasScxTang}
	autogenTestCasesFuncScxTavt = autogenTestCaseFunc{"HasScxTavt", uscripts.HasScxTavt}
	autogenTestCasesFuncScxTelu = autogenTestCaseFunc{"HasScxTelu", uscripts.HasScxTelu}
	autogenTestCasesFuncScxTfng = autogenTestCaseFunc{"HasScxTfng", uscripts.HasScxTfng}
	autogenTestCasesFuncScxTglg = autogenTestCaseFunc{"HasScxTglg", uscripts.HasScxTglg}
	autogenTestCasesFuncScxThaa = autogenTestCaseFunc{"HasScxThaa", uscripts.HasScxThaa}
	autogenTestCasesFuncScxThai = autogenTestCaseFunc{"HasScxThai", uscripts.HasScxThai}
	autogenTestCasesFuncScxTibt = autogenTestCaseFunc{"HasScxTibt", uscripts.HasScxTibt}
	autogenTestCasesFuncScxTirh = autogenTestCaseFunc{"HasScxTirh", uscripts.HasScxTirh}
	autogenTestCasesFuncScxTnsa = autogenTestCaseFunc{"HasScxTnsa", uscripts.HasScxTnsa}
	autogenTestCasesFuncScxToto = autogenTestCaseFunc{"HasScxToto", uscripts.HasScxToto}
	autogenTestCasesFuncScxUgar = autogenTestCaseFunc{"HasScxUgar", uscripts.HasScxUgar}
	autogenTestCasesFuncScxVaii = autogenTestCaseFunc{"HasScxVaii", uscripts.HasScxVaii}
	autogenTestCasesFuncScxVith = autogenTestCaseFunc{"HasScxVith", uscripts.HasScxVith}
	autogenTestCasesFuncScxWara = autogenTestCaseFunc{"HasScxWara", uscripts.HasScxWara}
	autogenTestCasesFuncScxWcho = autogenTestCaseFunc{"HasScxWcho", uscripts.HasScxWcho}
	autogenTestCasesFuncScxXpeo = autogenTestCaseFunc{"HasScxXpeo", uscripts.HasScxXpeo}
	autogenTestCasesFuncScxXsux = autogenTestCaseFunc{"HasScxXsux", uscripts.HasScxXsux}
	autogenTestCasesFuncScxYezi = autogenTestCaseFunc{"HasScxYezi", uscripts.HasScxYezi}
	autogenTestCasesFuncScxYiii = autogenTestCaseFunc{"HasScxYiii", uscripts.HasScxYiii}
	autogenTestCasesFuncScxZanb = autogenTestCaseFunc{"HasScxZanb", uscripts.HasScxZanb}
	autogenTestCasesFuncScxZinh = autogenTestCaseFunc{"HasScxZinh", uscripts.HasScxZinh}
	autogenTestCasesFuncScxZyyy = autogenTestCaseFunc{"HasScxZyyy", uscripts.HasScxZyyy}
	autogenTestCasesFuncScxZzzz = autogenTestCaseFunc{"HasScxZzzz", uscripts.HasScxZzzz}
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
	{ // Scx
		{-1000, -1, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxAdlm,
			&autogenTestCasesFuncScxAghb,
			&autogenTestCasesFuncScxAhom,
			&autogenTestCasesFuncScxArab,
			&autogenTestCasesFuncScxArmi,
			&autogenTestCasesFuncScxArmn,
			&autogenTestCasesFuncScxAvst,
			&autogenTestCasesFuncScxBali,
			&autogenTestCasesFuncScxBamu,
			&autogenTestCasesFuncScxBass,
			&autogenTestCasesFuncScxBatk,
			&autogenTestCasesFuncScxBeng,
			&autogenTestCasesFuncScxBhks,
			&autogenTestCasesFuncScxBopo,
			&autogenTestCasesFuncScxBrah,
			&autogenTestCasesFuncScxBrai,
			&autogenTestCasesFuncScxBugi,
			&autogenTestCasesFuncScxBuhd,
			&autogenTestCasesFuncScxCakm,
			&autogenTestCasesFuncScxCans,
			&autogenTestCasesFuncScxCari,
			&autogenTestCasesFuncScxCham,
			&autogenTestCasesFuncScxCher,
			&autogenTestCasesFuncScxChrs,
			&autogenTestCasesFuncScxCopt,
			&autogenTestCasesFuncScxCpmn,
			&autogenTestCasesFuncScxCprt,
			&autogenTestCasesFuncScxCyrl,
			&autogenTestCasesFuncScxDeva,
			&autogenTestCasesFuncScxDiak,
			&autogenTestCasesFuncScxDogr,
			&autogenTestCasesFuncScxDsrt,
			&autogenTestCasesFuncScxDupl,
			&autogenTestCasesFuncScxEgyp,
			&autogenTestCasesFuncScxElba,
			&autogenTestCasesFuncScxElym,
			&autogenTestCasesFuncScxEthi,
			&autogenTestCasesFuncScxGeor,
			&autogenTestCasesFuncScxGlag,
			&autogenTestCasesFuncScxGong,
			&autogenTestCasesFuncScxGonm,
			&autogenTestCasesFuncScxGoth,
			&autogenTestCasesFuncScxGran,
			&autogenTestCasesFuncScxGrek,
			&autogenTestCasesFuncScxGujr,
			&autogenTestCasesFuncScxGuru,
			&autogenTestCasesFuncScxHang,
			&autogenTestCasesFuncScxHani,
			&autogenTestCasesFuncScxHano,
			&autogenTestCasesFuncScxHatr,
			&autogenTestCasesFuncScxHebr,
			&autogenTestCasesFuncScxHira,
			&autogenTestCasesFuncScxHluw,
			&autogenTestCasesFuncScxHmng,
			&autogenTestCasesFuncScxHmnp,
			&autogenTestCasesFuncScxHung,
			&autogenTestCasesFuncScxItal,
			&autogenTestCasesFuncScxJava,
			&autogenTestCasesFuncScxKali,
			&autogenTestCasesFuncScxKana,
			&autogenTestCasesFuncScxKawi,
			&autogenTestCasesFuncScxKhar,
			&autogenTestCasesFuncScxKhmr,
			&autogenTestCasesFuncScxKhoj,
			&autogenTestCasesFuncScxKits,
			&autogenTestCasesFuncScxKnda,
			&autogenTestCasesFuncScxKthi,
			&autogenTestCasesFuncScxLana,
			&autogenTestCasesFuncScxLaoo,
			&autogenTestCasesFuncScxLatn,
			&autogenTestCasesFuncScxLepc,
			&autogenTestCasesFuncScxLimb,
			&autogenTestCasesFuncScxLina,
			&autogenTestCasesFuncScxLinb,
			&autogenTestCasesFuncScxLisu,
			&autogenTestCasesFuncScxLyci,
			&autogenTestCasesFuncScxLydi,
			&autogenTestCasesFuncScxMahj,
			&autogenTestCasesFuncScxMaka,
			&autogenTestCasesFuncScxMand,
			&autogenTestCasesFuncScxMani,
			&autogenTestCasesFuncScxMarc,
			&autogenTestCasesFuncScxMedf,
			&autogenTestCasesFuncScxMend,
			&autogenTestCasesFuncScxMerc,
			&autogenTestCasesFuncScxMero,
			&autogenTestCasesFuncScxMlym,
			&autogenTestCasesFuncScxModi,
			&autogenTestCasesFuncScxMong,
			&autogenTestCasesFuncScxMroo,
			&autogenTestCasesFuncScxMtei,
			&autogenTestCasesFuncScxMult,
			&autogenTestCasesFuncScxMymr,
			&autogenTestCasesFuncScxNagm,
			&autogenTestCasesFuncScxNand,
			&autogenTestCasesFuncScxNarb,
			&autogenTestCasesFuncScxNbat,
			&autogenTestCasesFuncScxNewa,
			&autogenTestCasesFuncScxNkoo,
			&autogenTestCasesFuncScxNshu,
			&autogenTestCasesFuncScxOgam,
			&autogenTestCasesFuncScxOlck,
			&autogenTestCasesFuncScxOrkh,
			&autogenTestCasesFuncScxOrya,
			&autogenTestCasesFuncScxOsge,
			&autogenTestCasesFuncScxOsma,
			&autogenTestCasesFuncScxOugr,
			&autogenTestCasesFuncScxPalm,
			&autogenTestCasesFuncScxPauc,
			&autogenTestCasesFuncScxPerm,
			&autogenTestCasesFuncScxPhag,
			&autogenTestCasesFuncScxPhli,
			&autogenTestCasesFuncScxPhlp,
			&autogenTestCasesFuncScxPhnx,
			&autogenTestCasesFuncScxPlrd,
			&autogenTestCasesFuncScxPrti,
			&autogenTestCasesFuncScxRjng,
			&autogenTestCasesFuncScxRohg,
			&autogenTestCasesFuncScxRunr,
			&autogenTestCasesFuncScxSamr,
			&autogenTestCasesFuncScxSarb,
			&autogenTestCasesFuncScxSaur,
			&autogenTestCasesFuncScxSgnw,
			&autogenTestCasesFuncScxShaw,
			&autogenTestCasesFuncScxShrd,
			&autogenTestCasesFuncScxSidd,
			&autogenTestCasesFuncScxSind,
			&autogenTestCasesFuncScxSinh,
			&autogenTestCasesFuncScxSogd,
			&autogenTestCasesFuncScxSogo,
			&autogenTestCasesFuncScxSora,
			&autogenTestCasesFuncScxSoyo,
			&autogenTestCasesFuncScxSund,
			&autogenTestCasesFuncScxSylo,
			&autogenTestCasesFuncScxSyrc,
			&autogenTestCasesFuncScxTagb,
			&autogenTestCasesFuncScxTakr,
			&autogenTestCasesFuncScxTale,
			&autogenTestCasesFuncScxTalu,
			&autogenTestCasesFuncScxTaml,
			&autogenTestCasesFuncScxTang,
			&autogenTestCasesFuncScxTavt,
			&autogenTestCasesFuncScxTelu,
			&autogenTestCasesFuncScxTfng,
			&autogenTestCasesFuncScxTglg,
			&autogenTestCasesFuncScxThaa,
			&autogenTestCasesFuncScxThai,
			&autogenTestCasesFuncScxTibt,
			&autogenTestCasesFuncScxTirh,
			&autogenTestCasesFuncScxTnsa,
			&autogenTestCasesFuncScxToto,
			&autogenTestCasesFuncScxUgar,
			&autogenTestCasesFuncScxVaii,
			&autogenTestCasesFuncScxVith,
			&autogenTestCasesFuncScxWara,
			&autogenTestCasesFuncScxWcho,
			&autogenTestCasesFuncScxXpeo,
			&autogenTestCasesFuncScxXsux,
			&autogenTestCasesFuncScxYezi,
			&autogenTestCasesFuncScxYiii,
			&autogenTestCasesFuncScxZanb,
			&autogenTestCasesFuncScxZinh,
			&autogenTestCasesFuncScxZyyy,
			&autogenTestCasesFuncScxZzzz,
		}},
		{0x00, 0x40, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x41, 0x5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0x5B, 0x60, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x61, 0x7A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0x7B, 0xA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0xAA, 0xAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0xAB, 0xB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0xBA, 0xBA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0xBB, 0xBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0xC0, 0xD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0xD7, 0xD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0xD8, 0xF6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0xF7, 0xF7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0xF8, 0x2B8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0x2B9, 0x2DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x2E0, 0x2E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0x2E5, 0x2E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x2EA, 0x2EB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBopo,
		}, []*autogenTestCaseFunc{}},
		{0x2EC, 0x2FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x300, 0x341, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZinh,
		}, []*autogenTestCaseFunc{}},
		{0x342, 0x342, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x343, 0x344, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZinh,
		}, []*autogenTestCaseFunc{}},
		{0x345, 0x345, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x346, 0x362, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZinh,
		}, []*autogenTestCaseFunc{}},
		{0x363, 0x36F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0x370, 0x373, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x374, 0x374, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x375, 0x377, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x378, 0x379, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x37A, 0x37D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x37E, 0x37E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x37F, 0x37F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x380, 0x383, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x384, 0x384, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x385, 0x385, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x386, 0x386, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x387, 0x387, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x388, 0x38A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x38B, 0x38B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x38C, 0x38C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x38D, 0x38D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x38E, 0x3A1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x3A2, 0x3A2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x3A3, 0x3E1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x3E2, 0x3EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCopt,
		}, []*autogenTestCaseFunc{}},
		{0x3F0, 0x3FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x400, 0x482, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCyrl,
		}, []*autogenTestCaseFunc{}},
		{0x483, 0x483, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCyrl,
			&autogenTestCasesFuncScxPerm,
		}, []*autogenTestCaseFunc{}},
		{0x484, 0x484, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCyrl,
			&autogenTestCasesFuncScxGlag,
		}, []*autogenTestCaseFunc{}},
		{0x485, 0x486, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCyrl,
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0x487, 0x487, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCyrl,
			&autogenTestCasesFuncScxGlag,
		}, []*autogenTestCaseFunc{}},
		{0x488, 0x52F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCyrl,
		}, []*autogenTestCaseFunc{}},
		{0x530, 0x530, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x531, 0x556, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArmn,
		}, []*autogenTestCaseFunc{}},
		{0x557, 0x558, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x559, 0x58A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArmn,
		}, []*autogenTestCaseFunc{}},
		{0x58B, 0x58C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x58D, 0x58F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArmn,
		}, []*autogenTestCaseFunc{}},
		{0x590, 0x590, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x591, 0x5C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHebr,
		}, []*autogenTestCaseFunc{}},
		{0x5C8, 0x5CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x5D0, 0x5EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHebr,
		}, []*autogenTestCaseFunc{}},
		{0x5EB, 0x5EE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x5EF, 0x5F4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHebr,
		}, []*autogenTestCaseFunc{}},
		{0x5F5, 0x5FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x600, 0x604, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x605, 0x605, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x606, 0x60B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x60C, 0x60C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
			&autogenTestCasesFuncScxNkoo,
			&autogenTestCasesFuncScxRohg,
			&autogenTestCasesFuncScxSyrc,
			&autogenTestCasesFuncScxThaa,
			&autogenTestCasesFuncScxYezi,
		}, []*autogenTestCaseFunc{}},
		{0x60D, 0x61A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x61B, 0x61B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
			&autogenTestCasesFuncScxNkoo,
			&autogenTestCasesFuncScxRohg,
			&autogenTestCasesFuncScxSyrc,
			&autogenTestCasesFuncScxThaa,
			&autogenTestCasesFuncScxYezi,
		}, []*autogenTestCaseFunc{}},
		{0x61C, 0x61C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
			&autogenTestCasesFuncScxSyrc,
			&autogenTestCasesFuncScxThaa,
		}, []*autogenTestCaseFunc{}},
		{0x61D, 0x61E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x61F, 0x61F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxAdlm,
			&autogenTestCasesFuncScxArab,
			&autogenTestCasesFuncScxNkoo,
			&autogenTestCasesFuncScxRohg,
			&autogenTestCasesFuncScxSyrc,
			&autogenTestCasesFuncScxThaa,
			&autogenTestCasesFuncScxYezi,
		}, []*autogenTestCaseFunc{}},
		{0x620, 0x63F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x640, 0x640, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxAdlm,
			&autogenTestCasesFuncScxArab,
			&autogenTestCasesFuncScxMand,
			&autogenTestCasesFuncScxMani,
			&autogenTestCasesFuncScxOugr,
			&autogenTestCasesFuncScxPhlp,
			&autogenTestCasesFuncScxRohg,
			&autogenTestCasesFuncScxSogd,
			&autogenTestCasesFuncScxSyrc,
		}, []*autogenTestCaseFunc{}},
		{0x641, 0x64A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x64B, 0x655, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
			&autogenTestCasesFuncScxSyrc,
		}, []*autogenTestCaseFunc{}},
		{0x656, 0x65F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x660, 0x669, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
			&autogenTestCasesFuncScxThaa,
			&autogenTestCasesFuncScxYezi,
		}, []*autogenTestCaseFunc{}},
		{0x66A, 0x66F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x670, 0x670, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
			&autogenTestCasesFuncScxSyrc,
		}, []*autogenTestCaseFunc{}},
		{0x671, 0x6D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x6D4, 0x6D4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
			&autogenTestCasesFuncScxRohg,
		}, []*autogenTestCaseFunc{}},
		{0x6D5, 0x6DC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x6DD, 0x6DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x6DE, 0x6FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x700, 0x70D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSyrc,
		}, []*autogenTestCaseFunc{}},
		{0x70E, 0x70E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x70F, 0x74A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSyrc,
		}, []*autogenTestCaseFunc{}},
		{0x74B, 0x74C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x74D, 0x74F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSyrc,
		}, []*autogenTestCaseFunc{}},
		{0x750, 0x77F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x780, 0x7B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxThaa,
		}, []*autogenTestCaseFunc{}},
		{0x7B2, 0x7BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x7C0, 0x7FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxNkoo,
		}, []*autogenTestCaseFunc{}},
		{0x7FB, 0x7FC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x7FD, 0x7FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxNkoo,
		}, []*autogenTestCaseFunc{}},
		{0x800, 0x82D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSamr,
		}, []*autogenTestCaseFunc{}},
		{0x82E, 0x82F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x830, 0x83E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSamr,
		}, []*autogenTestCaseFunc{}},
		{0x83F, 0x83F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x840, 0x85B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMand,
		}, []*autogenTestCaseFunc{}},
		{0x85C, 0x85D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x85E, 0x85E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMand,
		}, []*autogenTestCaseFunc{}},
		{0x85F, 0x85F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x860, 0x86A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSyrc,
		}, []*autogenTestCaseFunc{}},
		{0x86B, 0x86F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x870, 0x88E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x88F, 0x88F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x890, 0x891, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x892, 0x897, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x898, 0x8E1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x8E2, 0x8E2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x8E3, 0x8FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
		}, []*autogenTestCaseFunc{}},
		{0x900, 0x950, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxDeva,
		}, []*autogenTestCaseFunc{}},
		{0x951, 0x951, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
			&autogenTestCasesFuncScxDeva,
			&autogenTestCasesFuncScxGran,
			&autogenTestCasesFuncScxGujr,
			&autogenTestCasesFuncScxGuru,
			&autogenTestCasesFuncScxKnda,
			&autogenTestCasesFuncScxLatn,
			&autogenTestCasesFuncScxMlym,
			&autogenTestCasesFuncScxOrya,
			&autogenTestCasesFuncScxShrd,
			&autogenTestCasesFuncScxTaml,
			&autogenTestCasesFuncScxTelu,
			&autogenTestCasesFuncScxTirh,
		}, []*autogenTestCaseFunc{}},
		{0x952, 0x952, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
			&autogenTestCasesFuncScxDeva,
			&autogenTestCasesFuncScxGran,
			&autogenTestCasesFuncScxGujr,
			&autogenTestCasesFuncScxGuru,
			&autogenTestCasesFuncScxKnda,
			&autogenTestCasesFuncScxLatn,
			&autogenTestCasesFuncScxMlym,
			&autogenTestCasesFuncScxOrya,
			&autogenTestCasesFuncScxTaml,
			&autogenTestCasesFuncScxTelu,
			&autogenTestCasesFuncScxTirh,
		}, []*autogenTestCaseFunc{}},
		{0x953, 0x954, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZinh,
		}, []*autogenTestCaseFunc{}},
		{0x955, 0x963, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxDeva,
		}, []*autogenTestCaseFunc{}},
		{0x964, 0x964, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
			&autogenTestCasesFuncScxDeva,
			&autogenTestCasesFuncScxDogr,
			&autogenTestCasesFuncScxGong,
			&autogenTestCasesFuncScxGonm,
			&autogenTestCasesFuncScxGran,
			&autogenTestCasesFuncScxGujr,
			&autogenTestCasesFuncScxGuru,
			&autogenTestCasesFuncScxKnda,
			&autogenTestCasesFuncScxMahj,
			&autogenTestCasesFuncScxMlym,
			&autogenTestCasesFuncScxNand,
			&autogenTestCasesFuncScxOrya,
			&autogenTestCasesFuncScxSind,
			&autogenTestCasesFuncScxSinh,
			&autogenTestCasesFuncScxSylo,
			&autogenTestCasesFuncScxTakr,
			&autogenTestCasesFuncScxTaml,
			&autogenTestCasesFuncScxTelu,
			&autogenTestCasesFuncScxTirh,
		}, []*autogenTestCaseFunc{}},
		{0x965, 0x965, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
			&autogenTestCasesFuncScxDeva,
			&autogenTestCasesFuncScxDogr,
			&autogenTestCasesFuncScxGong,
			&autogenTestCasesFuncScxGonm,
			&autogenTestCasesFuncScxGran,
			&autogenTestCasesFuncScxGujr,
			&autogenTestCasesFuncScxGuru,
			&autogenTestCasesFuncScxKnda,
			&autogenTestCasesFuncScxLimb,
			&autogenTestCasesFuncScxMahj,
			&autogenTestCasesFuncScxMlym,
			&autogenTestCasesFuncScxNand,
			&autogenTestCasesFuncScxOrya,
			&autogenTestCasesFuncScxSind,
			&autogenTestCasesFuncScxSinh,
			&autogenTestCasesFuncScxSylo,
			&autogenTestCasesFuncScxTakr,
			&autogenTestCasesFuncScxTaml,
			&autogenTestCasesFuncScxTelu,
			&autogenTestCasesFuncScxTirh,
		}, []*autogenTestCaseFunc{}},
		{0x966, 0x96F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxDeva,
			&autogenTestCasesFuncScxDogr,
			&autogenTestCasesFuncScxKthi,
			&autogenTestCasesFuncScxMahj,
		}, []*autogenTestCaseFunc{}},
		{0x970, 0x97F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxDeva,
		}, []*autogenTestCaseFunc{}},
		{0x980, 0x983, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x984, 0x984, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x985, 0x98C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x98D, 0x98E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x98F, 0x990, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x991, 0x992, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x993, 0x9A8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x9A9, 0x9A9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x9AA, 0x9B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x9B1, 0x9B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x9B2, 0x9B2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x9B3, 0x9B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x9B6, 0x9B9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x9BA, 0x9BB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x9BC, 0x9C4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x9C5, 0x9C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x9C7, 0x9C8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x9C9, 0x9CA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x9CB, 0x9CE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x9CF, 0x9D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x9D7, 0x9D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x9D8, 0x9DB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x9DC, 0x9DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x9DE, 0x9DE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x9DF, 0x9E3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x9E4, 0x9E5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x9E6, 0x9EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
			&autogenTestCasesFuncScxCakm,
			&autogenTestCasesFuncScxSylo,
		}, []*autogenTestCaseFunc{}},
		{0x9F0, 0x9FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x9FF, 0xA00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA01, 0xA03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA04, 0xA04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA05, 0xA0A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA0B, 0xA0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA0F, 0xA10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA11, 0xA12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA13, 0xA28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA29, 0xA29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA2A, 0xA30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA31, 0xA31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA32, 0xA33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA34, 0xA34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA35, 0xA36, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA37, 0xA37, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA38, 0xA39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA3A, 0xA3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA3C, 0xA3C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA3D, 0xA3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA3E, 0xA42, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA43, 0xA46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA47, 0xA48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA49, 0xA4A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA4B, 0xA4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA4E, 0xA50, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA51, 0xA51, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA52, 0xA58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA59, 0xA5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA5D, 0xA5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA5E, 0xA5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA5F, 0xA65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA66, 0xA6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
			&autogenTestCasesFuncScxMult,
		}, []*autogenTestCaseFunc{}},
		{0xA70, 0xA76, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGuru,
		}, []*autogenTestCaseFunc{}},
		{0xA77, 0xA80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA81, 0xA83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xA84, 0xA84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA85, 0xA8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xA8E, 0xA8E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA8F, 0xA91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xA92, 0xA92, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA93, 0xAA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xAA9, 0xAA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xAAA, 0xAB0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xAB1, 0xAB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xAB2, 0xAB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xAB4, 0xAB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xAB5, 0xAB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xABA, 0xABB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xABC, 0xAC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xAC6, 0xAC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xAC7, 0xAC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xACA, 0xACA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xACB, 0xACD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xACE, 0xACF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xAD0, 0xAD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xAD1, 0xADF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xAE0, 0xAE3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xAE4, 0xAE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xAE6, 0xAEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
			&autogenTestCasesFuncScxKhoj,
		}, []*autogenTestCaseFunc{}},
		{0xAF0, 0xAF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xAF2, 0xAF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xAF9, 0xAFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGujr,
		}, []*autogenTestCaseFunc{}},
		{0xB00, 0xB00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB01, 0xB03, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB04, 0xB04, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB05, 0xB0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB0D, 0xB0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB0F, 0xB10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB11, 0xB12, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB13, 0xB28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB29, 0xB29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB2A, 0xB30, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB31, 0xB31, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB32, 0xB33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB34, 0xB34, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB35, 0xB39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB3A, 0xB3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB3C, 0xB44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB45, 0xB46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB47, 0xB48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB49, 0xB4A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB4B, 0xB4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB4E, 0xB54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB55, 0xB57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB58, 0xB5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB5C, 0xB5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB5E, 0xB5E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB5F, 0xB63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB64, 0xB65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB66, 0xB77, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOrya,
		}, []*autogenTestCaseFunc{}},
		{0xB78, 0xB81, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB82, 0xB83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xB84, 0xB84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB85, 0xB8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xB8B, 0xB8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB8E, 0xB90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xB91, 0xB91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB92, 0xB95, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xB96, 0xB98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB99, 0xB9A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xB9B, 0xB9B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB9C, 0xB9C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xB9D, 0xB9D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xB9E, 0xB9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xBA0, 0xBA2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xBA3, 0xBA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xBA5, 0xBA7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xBA8, 0xBAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xBAB, 0xBAD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xBAE, 0xBB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xBBA, 0xBBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xBBE, 0xBC2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xBC3, 0xBC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xBC6, 0xBC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xBC9, 0xBC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xBCA, 0xBCD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xBCE, 0xBCF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xBD0, 0xBD0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xBD1, 0xBD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xBD7, 0xBD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xBD8, 0xBE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xBE6, 0xBF3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGran,
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xBF4, 0xBFA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0xBFB, 0xBFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC00, 0xC0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0xC0D, 0xC0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC0E, 0xC10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0xC11, 0xC11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC12, 0xC28, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0xC29, 0xC29, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC2A, 0xC39, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0xC3A, 0xC3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC3C, 0xC44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0xC45, 0xC45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC46, 0xC48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0xC49, 0xC49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC4A, 0xC4D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0xC4E, 0xC54, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC55, 0xC56, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0xC57, 0xC57, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC58, 0xC5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0xC5B, 0xC5C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC5D, 0xC5D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0xC5E, 0xC5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC60, 0xC63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0xC64, 0xC65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC66, 0xC6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0xC70, 0xC76, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC77, 0xC7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0xC80, 0xC8C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKnda,
		}, []*autogenTestCaseFunc{}},
		{0xC8D, 0xC8D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC8E, 0xC90, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKnda,
		}, []*autogenTestCaseFunc{}},
		{0xC91, 0xC91, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xC92, 0xCA8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKnda,
		}, []*autogenTestCaseFunc{}},
		{0xCA9, 0xCA9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xCAA, 0xCB3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKnda,
		}, []*autogenTestCaseFunc{}},
		{0xCB4, 0xCB4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xCB5, 0xCB9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKnda,
		}, []*autogenTestCaseFunc{}},
		{0xCBA, 0xCBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xCBC, 0xCC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKnda,
		}, []*autogenTestCaseFunc{}},
		{0xCC5, 0xCC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xCC6, 0xCC8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKnda,
		}, []*autogenTestCaseFunc{}},
		{0xCC9, 0xCC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xCCA, 0xCCD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKnda,
		}, []*autogenTestCaseFunc{}},
		{0xCCE, 0xCD4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xCD5, 0xCD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKnda,
		}, []*autogenTestCaseFunc{}},
		{0xCD7, 0xCDC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xCDD, 0xCDE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKnda,
		}, []*autogenTestCaseFunc{}},
		{0xCDF, 0xCDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xCE0, 0xCE3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKnda,
		}, []*autogenTestCaseFunc{}},
		{0xCE4, 0xCE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xCE6, 0xCEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKnda,
			&autogenTestCasesFuncScxNand,
		}, []*autogenTestCaseFunc{}},
		{0xCF0, 0xCF0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xCF1, 0xCF3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKnda,
		}, []*autogenTestCaseFunc{}},
		{0xCF4, 0xCFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xD00, 0xD0C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMlym,
		}, []*autogenTestCaseFunc{}},
		{0xD0D, 0xD0D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xD0E, 0xD10, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMlym,
		}, []*autogenTestCaseFunc{}},
		{0xD11, 0xD11, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xD12, 0xD44, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMlym,
		}, []*autogenTestCaseFunc{}},
		{0xD45, 0xD45, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xD46, 0xD48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMlym,
		}, []*autogenTestCaseFunc{}},
		{0xD49, 0xD49, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xD4A, 0xD4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMlym,
		}, []*autogenTestCaseFunc{}},
		{0xD50, 0xD53, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xD54, 0xD63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMlym,
		}, []*autogenTestCaseFunc{}},
		{0xD64, 0xD65, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xD66, 0xD7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMlym,
		}, []*autogenTestCaseFunc{}},
		{0xD80, 0xD80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xD81, 0xD83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSinh,
		}, []*autogenTestCaseFunc{}},
		{0xD84, 0xD84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xD85, 0xD96, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSinh,
		}, []*autogenTestCaseFunc{}},
		{0xD97, 0xD99, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xD9A, 0xDB1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSinh,
		}, []*autogenTestCaseFunc{}},
		{0xDB2, 0xDB2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xDB3, 0xDBB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSinh,
		}, []*autogenTestCaseFunc{}},
		{0xDBC, 0xDBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xDBD, 0xDBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSinh,
		}, []*autogenTestCaseFunc{}},
		{0xDBE, 0xDBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xDC0, 0xDC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSinh,
		}, []*autogenTestCaseFunc{}},
		{0xDC7, 0xDC9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xDCA, 0xDCA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSinh,
		}, []*autogenTestCaseFunc{}},
		{0xDCB, 0xDCE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xDCF, 0xDD4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSinh,
		}, []*autogenTestCaseFunc{}},
		{0xDD5, 0xDD5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xDD6, 0xDD6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSinh,
		}, []*autogenTestCaseFunc{}},
		{0xDD7, 0xDD7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xDD8, 0xDDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSinh,
		}, []*autogenTestCaseFunc{}},
		{0xDE0, 0xDE5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xDE6, 0xDEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSinh,
		}, []*autogenTestCaseFunc{}},
		{0xDF0, 0xDF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xDF2, 0xDF4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSinh,
		}, []*autogenTestCaseFunc{}},
		{0xDF5, 0xE00, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xE01, 0xE3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxThai,
		}, []*autogenTestCaseFunc{}},
		{0xE3B, 0xE3E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xE3F, 0xE3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0xE40, 0xE5B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxThai,
		}, []*autogenTestCaseFunc{}},
		{0xE5C, 0xE80, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xE81, 0xE82, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLaoo,
		}, []*autogenTestCaseFunc{}},
		{0xE83, 0xE83, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xE84, 0xE84, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLaoo,
		}, []*autogenTestCaseFunc{}},
		{0xE85, 0xE85, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xE86, 0xE8A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLaoo,
		}, []*autogenTestCaseFunc{}},
		{0xE8B, 0xE8B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xE8C, 0xEA3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLaoo,
		}, []*autogenTestCaseFunc{}},
		{0xEA4, 0xEA4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xEA5, 0xEA5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLaoo,
		}, []*autogenTestCaseFunc{}},
		{0xEA6, 0xEA6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xEA7, 0xEBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLaoo,
		}, []*autogenTestCaseFunc{}},
		{0xEBE, 0xEBF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xEC0, 0xEC4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLaoo,
		}, []*autogenTestCaseFunc{}},
		{0xEC5, 0xEC5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xEC6, 0xEC6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLaoo,
		}, []*autogenTestCaseFunc{}},
		{0xEC7, 0xEC7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xEC8, 0xECE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLaoo,
		}, []*autogenTestCaseFunc{}},
		{0xECF, 0xECF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xED0, 0xED9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLaoo,
		}, []*autogenTestCaseFunc{}},
		{0xEDA, 0xEDB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xEDC, 0xEDF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLaoo,
		}, []*autogenTestCaseFunc{}},
		{0xEE0, 0xEFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xF00, 0xF47, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTibt,
		}, []*autogenTestCaseFunc{}},
		{0xF48, 0xF48, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xF49, 0xF6C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTibt,
		}, []*autogenTestCaseFunc{}},
		{0xF6D, 0xF70, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xF71, 0xF97, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTibt,
		}, []*autogenTestCaseFunc{}},
		{0xF98, 0xF98, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xF99, 0xFBC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTibt,
		}, []*autogenTestCaseFunc{}},
		{0xFBD, 0xFBD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xFBE, 0xFCC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTibt,
		}, []*autogenTestCaseFunc{}},
		{0xFCD, 0xFCD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xFCE, 0xFD4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTibt,
		}, []*autogenTestCaseFunc{}},
		{0xFD5, 0xFD8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0xFD9, 0xFDA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTibt,
		}, []*autogenTestCaseFunc{}},
		{0xFDB, 0xFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1000, 0x103F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMymr,
		}, []*autogenTestCaseFunc{}},
		{0x1040, 0x1049, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCakm,
			&autogenTestCasesFuncScxMymr,
			&autogenTestCasesFuncScxTale,
		}, []*autogenTestCaseFunc{}},
		{0x104A, 0x109F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMymr,
		}, []*autogenTestCaseFunc{}},
		{0x10A0, 0x10C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGeor,
		}, []*autogenTestCaseFunc{}},
		{0x10C6, 0x10C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x10C7, 0x10C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGeor,
		}, []*autogenTestCaseFunc{}},
		{0x10C8, 0x10CC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x10CD, 0x10CD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGeor,
		}, []*autogenTestCaseFunc{}},
		{0x10CE, 0x10CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x10D0, 0x10FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGeor,
		}, []*autogenTestCaseFunc{}},
		{0x10FB, 0x10FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGeor,
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0x10FC, 0x10FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGeor,
		}, []*autogenTestCaseFunc{}},
		{0x1100, 0x11FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHang,
		}, []*autogenTestCaseFunc{}},
		{0x1200, 0x1248, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x1249, 0x1249, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x124A, 0x124D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x124E, 0x124F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1250, 0x1256, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x1257, 0x1257, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1258, 0x1258, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x1259, 0x1259, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x125A, 0x125D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x125E, 0x125F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1260, 0x1288, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x1289, 0x1289, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x128A, 0x128D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x128E, 0x128F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1290, 0x12B0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x12B1, 0x12B1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x12B2, 0x12B5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x12B6, 0x12B7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x12B8, 0x12BE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x12BF, 0x12BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x12C0, 0x12C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x12C1, 0x12C1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x12C2, 0x12C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x12C6, 0x12C7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x12C8, 0x12D6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x12D7, 0x12D7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x12D8, 0x1310, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x1311, 0x1311, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1312, 0x1315, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x1316, 0x1317, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1318, 0x135A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x135B, 0x135C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x135D, 0x137C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x137D, 0x137F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1380, 0x1399, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x139A, 0x139F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x13A0, 0x13F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCher,
		}, []*autogenTestCaseFunc{}},
		{0x13F6, 0x13F7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x13F8, 0x13FD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCher,
		}, []*autogenTestCaseFunc{}},
		{0x13FE, 0x13FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1400, 0x167F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCans,
		}, []*autogenTestCaseFunc{}},
		{0x1680, 0x169C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOgam,
		}, []*autogenTestCaseFunc{}},
		{0x169D, 0x169F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x16A0, 0x16EA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxRunr,
		}, []*autogenTestCaseFunc{}},
		{0x16EB, 0x16ED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x16EE, 0x16F8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxRunr,
		}, []*autogenTestCaseFunc{}},
		{0x16F9, 0x16FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1700, 0x1715, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTglg,
		}, []*autogenTestCaseFunc{}},
		{0x1716, 0x171E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x171F, 0x171F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTglg,
		}, []*autogenTestCaseFunc{}},
		{0x1720, 0x1734, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHano,
		}, []*autogenTestCaseFunc{}},
		{0x1735, 0x1736, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBuhd,
			&autogenTestCasesFuncScxHano,
			&autogenTestCasesFuncScxTagb,
			&autogenTestCasesFuncScxTglg,
		}, []*autogenTestCaseFunc{}},
		{0x1737, 0x173F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1740, 0x1753, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBuhd,
		}, []*autogenTestCaseFunc{}},
		{0x1754, 0x175F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1760, 0x176C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTagb,
		}, []*autogenTestCaseFunc{}},
		{0x176D, 0x176D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x176E, 0x1770, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTagb,
		}, []*autogenTestCaseFunc{}},
		{0x1771, 0x1771, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1772, 0x1773, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxTagb,
		}, []*autogenTestCaseFunc{}},
		{0x1774, 0x177F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1780, 0x17DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKhmr,
		}, []*autogenTestCaseFunc{}},
		{0x17DE, 0x17DF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x17E0, 0x17E9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKhmr,
		}, []*autogenTestCaseFunc{}},
		{0x17EA, 0x17EF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x17F0, 0x17F9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKhmr,
		}, []*autogenTestCaseFunc{}},
		{0x17FA, 0x17FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1800, 0x1801, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMong,
		}, []*autogenTestCaseFunc{}},
		{0x1802, 0x1803, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMong,
			&autogenTestCasesFuncScxPhag,
		}, []*autogenTestCaseFunc{}},
		{0x1804, 0x1804, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMong,
		}, []*autogenTestCaseFunc{}},
		{0x1805, 0x1805, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMong,
			&autogenTestCasesFuncScxPhag,
		}, []*autogenTestCaseFunc{}},
		{0x1806, 0x1819, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMong,
		}, []*autogenTestCaseFunc{}},
		{0x181A, 0x181F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1820, 0x1878, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMong,
		}, []*autogenTestCaseFunc{}},
		{0x1879, 0x187F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1880, 0x18AA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMong,
		}, []*autogenTestCaseFunc{}},
		{0x18AB, 0x18AF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x18B0, 0x18F5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCans,
		}, []*autogenTestCaseFunc{}},
		{0x18F6, 0x18FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1900, 0x191E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLimb,
		}, []*autogenTestCaseFunc{}},
		{0x191F, 0x191F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1920, 0x192B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLimb,
		}, []*autogenTestCaseFunc{}},
		{0x196E, 0x196F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x19DB, 0x19DD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1A60, 0x1A7C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLana,
		}, []*autogenTestCaseFunc{}},
		{0x1AB0, 0x1ACE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZinh,
		}, []*autogenTestCaseFunc{}},
		{0x1BF4, 0x1BFB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1C80, 0x1C88, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCyrl,
		}, []*autogenTestCaseFunc{}},
		{0x1CD1, 0x1CD1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxDeva,
		}, []*autogenTestCaseFunc{}},
		{0x1CDA, 0x1CDA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxDeva,
			&autogenTestCasesFuncScxKnda,
			&autogenTestCasesFuncScxMlym,
			&autogenTestCasesFuncScxOrya,
			&autogenTestCasesFuncScxTaml,
			&autogenTestCasesFuncScxTelu,
		}, []*autogenTestCaseFunc{}},
		{0x1CEA, 0x1CEA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
			&autogenTestCasesFuncScxDeva,
		}, []*autogenTestCaseFunc{}},
		{0x1CF7, 0x1CF7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBeng,
		}, []*autogenTestCaseFunc{}},
		{0x1D5D, 0x1D61, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGrek,
		}, []*autogenTestCaseFunc{}},
		{0x1DF8, 0x1DF8, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCyrl,
			&autogenTestCasesFuncScxSyrc,
		}, []*autogenTestCaseFunc{}},
		{0x1F1E, 0x1F1F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1F5A, 0x1F5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1FB5, 0x1FB5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1FF0, 0x1FF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x202F, 0x202F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLatn,
			&autogenTestCasesFuncScxMong,
		}, []*autogenTestCaseFunc{}},
		{0x2080, 0x208E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x20F1, 0x20FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x214E, 0x214E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0x244B, 0x245F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x2C00, 0x2C5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGlag,
		}, []*autogenTestCaseFunc{}},
		{0x2D28, 0x2D2C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x2D80, 0x2D96, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x2DB8, 0x2DBE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x2DD8, 0x2DDE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0x2E9A, 0x2E9A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x3004, 0x3004, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x3021, 0x3029, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHani,
		}, []*autogenTestCaseFunc{}},
		{0x303C, 0x303D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHani,
			&autogenTestCasesFuncScxHira,
			&autogenTestCasesFuncScxKana,
		}, []*autogenTestCaseFunc{}},
		{0x30A1, 0x30FA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKana,
		}, []*autogenTestCaseFunc{}},
		{0x318F, 0x318F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x321F, 0x321F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x32CC, 0x32CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x33E0, 0x33FE, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHani,
		}, []*autogenTestCaseFunc{}},
		{0xA4C7, 0xA4CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA6F8, 0xA6FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA7D2, 0xA7D2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA830, 0xA832, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxDeva,
			&autogenTestCasesFuncScxDogr,
			&autogenTestCasesFuncScxGujr,
			&autogenTestCasesFuncScxGuru,
			&autogenTestCasesFuncScxKhoj,
			&autogenTestCasesFuncScxKnda,
			&autogenTestCasesFuncScxKthi,
			&autogenTestCasesFuncScxMahj,
			&autogenTestCasesFuncScxMlym,
			&autogenTestCasesFuncScxModi,
			&autogenTestCasesFuncScxNand,
			&autogenTestCasesFuncScxShrd,
			&autogenTestCasesFuncScxSind,
			&autogenTestCasesFuncScxTakr,
			&autogenTestCasesFuncScxTirh,
		}, []*autogenTestCaseFunc{}},
		{0xA880, 0xA8C5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSaur,
		}, []*autogenTestCaseFunc{}},
		{0xA8F4, 0xA8FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxDeva,
		}, []*autogenTestCaseFunc{}},
		{0xA97D, 0xA97F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xA9FF, 0xA9FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xAA60, 0xAA7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMymr,
		}, []*autogenTestCaseFunc{}},
		{0xAB09, 0xAB0E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxEthi,
		}, []*autogenTestCaseFunc{}},
		{0xAB30, 0xAB5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0xABC0, 0xABED, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMtei,
		}, []*autogenTestCaseFunc{}},
		{0xD7CB, 0xD7FB, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHang,
		}, []*autogenTestCaseFunc{}},
		{0xFB13, 0xFB17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArmn,
		}, []*autogenTestCaseFunc{}},
		{0xFB40, 0xFB41, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHebr,
		}, []*autogenTestCaseFunc{}},
		{0xFD3E, 0xFD3F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
			&autogenTestCasesFuncScxNkoo,
		}, []*autogenTestCaseFunc{}},
		{0xFDF2, 0xFDF2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArab,
			&autogenTestCasesFuncScxThaa,
		}, []*autogenTestCaseFunc{}},
		{0xFE2E, 0xFE2F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCyrl,
		}, []*autogenTestCaseFunc{}},
		{0xFE6C, 0xFE6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xFF21, 0xFF3A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLatn,
		}, []*autogenTestCaseFunc{}},
		{0xFF9E, 0xFF9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHira,
			&autogenTestCasesFuncScxKana,
		}, []*autogenTestCaseFunc{}},
		{0xFFD8, 0xFFD9, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xFFFE, 0xFFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1003E, 0x1003E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x10102, 0x10102, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCprt,
			&autogenTestCasesFuncScxLinb,
		}, []*autogenTestCaseFunc{}},
		{0x1019D, 0x1019F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x102A0, 0x102D0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxCari,
		}, []*autogenTestCaseFunc{}},
		{0x1034B, 0x1034F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x103C8, 0x103D5, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxXpeo,
		}, []*autogenTestCaseFunc{}},
		{0x104B0, 0x104D3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxOsge,
		}, []*autogenTestCaseFunc{}},
		{0x1056F, 0x1056F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxAghb,
		}, []*autogenTestCaseFunc{}},
		{0x10596, 0x10596, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x105BD, 0x105FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x10786, 0x10786, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x10809, 0x10809, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x10840, 0x10855, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxArmi,
		}, []*autogenTestCaseFunc{}},
		{0x108E0, 0x108F2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHatr,
		}, []*autogenTestCaseFunc{}},
		{0x10920, 0x10939, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxLydi,
		}, []*autogenTestCaseFunc{}},
		{0x109D0, 0x109D1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x10A15, 0x10A17, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKhar,
		}, []*autogenTestCaseFunc{}},
		{0x10A50, 0x10A58, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxKhar,
		}, []*autogenTestCaseFunc{}},
		{0x10AF2, 0x10AF2, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMani,
			&autogenTestCasesFuncScxOugr,
		}, []*autogenTestCaseFunc{}},
		{0x10B58, 0x10B5F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxPrti,
		}, []*autogenTestCaseFunc{}},
		{0x10BA9, 0x10BAF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxPhlp,
		}, []*autogenTestCaseFunc{}},
		{0x10CFA, 0x10CFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHung,
		}, []*autogenTestCaseFunc{}},
		{0x10EAA, 0x10EAA, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x10F30, 0x10F59, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxSogd,
		}, []*autogenTestCaseFunc{}},
		{0x11000, 0x1104D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxBrah,
		}, []*autogenTestCaseFunc{}},
		{0x110CE, 0x110CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x11148, 0x1114F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x11212, 0x11212, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1128E, 0x1128E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x112FA, 0x112FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1130F, 0x11310, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGran,
		}, []*autogenTestCaseFunc{}},
		{0x11335, 0x11339, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGran,
		}, []*autogenTestCaseFunc{}},
		{0x1134E, 0x1134F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1136D, 0x1136F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x114C8, 0x114CF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x11645, 0x1164F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x116CA, 0x116FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1183C, 0x1189F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1190C, 0x11913, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxDiak,
		}, []*autogenTestCaseFunc{}},
		{0x1193B, 0x11946, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxDiak,
		}, []*autogenTestCaseFunc{}},
		{0x119DA, 0x119E4, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxNand,
		}, []*autogenTestCaseFunc{}},
		{0x11AF9, 0x11AFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x11C46, 0x11C4F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x11CB7, 0x11CFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x11D3B, 0x11D3B, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x11D66, 0x11D66, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x11D99, 0x11D9F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x11F3B, 0x11F3D, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x11FD3, 0x11FD3, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxGran,
			&autogenTestCasesFuncScxTaml,
		}, []*autogenTestCaseFunc{}},
		{0x12470, 0x12474, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxXsux,
		}, []*autogenTestCaseFunc{}},
		{0x14400, 0x14646, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHluw,
		}, []*autogenTestCaseFunc{}},
		{0x16A6E, 0x16A6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxMroo,
		}, []*autogenTestCaseFunc{}},
		{0x16AF6, 0x16AFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x16B78, 0x16B7C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x16F88, 0x16F8E, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x16FF0, 0x16FF1, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHani,
		}, []*autogenTestCaseFunc{}},
		{0x18D09, 0x1AFEF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1B001, 0x1B11F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHira,
		}, []*autogenTestCaseFunc{}},
		{0x1B156, 0x1B163, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1BC7D, 0x1BC7F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1CF2E, 0x1CF2F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1D127, 0x1D128, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1D1AA, 0x1D1AD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZinh,
		}, []*autogenTestCaseFunc{}},
		{0x1D2F4, 0x1D2FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1D456, 0x1D49C, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x1D4A9, 0x1D4AC, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x1D4C5, 0x1D505, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x1D51E, 0x1D539, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x1D54A, 0x1D550, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZyyy,
		}, []*autogenTestCaseFunc{}},
		{0x1DA8C, 0x1DA9A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1DF2B, 0x1DFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1E025, 0x1E025, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1E12D, 0x1E12F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1E2AF, 0x1E2BF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1E7E7, 0x1E7E7, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1E8C5, 0x1E8C6, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1E960, 0x1EC70, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1EE20, 0x1EE20, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1EE33, 0x1EE33, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1EE43, 0x1EE46, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1EE50, 0x1EE50, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1EE5A, 0x1EE5A, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1EE63, 0x1EE63, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1EE78, 0x1EE78, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1EE9C, 0x1EEA0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1EEF2, 0x1EFFF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1F0C0, 0x1F0C0, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1F200, 0x1F200, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxHira,
		}, []*autogenTestCaseFunc{}},
		{0x1F252, 0x1F25F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1F6FD, 0x1F6FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1F7F1, 0x1F7FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1F888, 0x1F88F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1FA6E, 0x1FA6F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1FAC6, 0x1FACD, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x1FB93, 0x1FB93, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x2B73A, 0x2B73F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x2EE5E, 0x2F7FF, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0xE0002, 0xE001F, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxZzzz,
		}, []*autogenTestCaseFunc{}},
		{0x110000, 0x110400, []*autogenTestCaseFunc{}, []*autogenTestCaseFunc{
			&autogenTestCasesFuncScxAdlm,
			&autogenTestCasesFuncScxAghb,
			&autogenTestCasesFuncScxAhom,
			&autogenTestCasesFuncScxArab,
			&autogenTestCasesFuncScxArmi,
			&autogenTestCasesFuncScxArmn,
			&autogenTestCasesFuncScxAvst,
			&autogenTestCasesFuncScxBali,
			&autogenTestCasesFuncScxBamu,
			&autogenTestCasesFuncScxBass,
			&autogenTestCasesFuncScxBatk,
			&autogenTestCasesFuncScxBeng,
			&autogenTestCasesFuncScxBhks,
			&autogenTestCasesFuncScxBopo,
			&autogenTestCasesFuncScxBrah,
			&autogenTestCasesFuncScxBrai,
			&autogenTestCasesFuncScxBugi,
			&autogenTestCasesFuncScxBuhd,
			&autogenTestCasesFuncScxCakm,
			&autogenTestCasesFuncScxCans,
			&autogenTestCasesFuncScxCari,
			&autogenTestCasesFuncScxCham,
			&autogenTestCasesFuncScxCher,
			&autogenTestCasesFuncScxChrs,
			&autogenTestCasesFuncScxCopt,
			&autogenTestCasesFuncScxCpmn,
			&autogenTestCasesFuncScxCprt,
			&autogenTestCasesFuncScxCyrl,
			&autogenTestCasesFuncScxDeva,
			&autogenTestCasesFuncScxDiak,
			&autogenTestCasesFuncScxDogr,
			&autogenTestCasesFuncScxDsrt,
			&autogenTestCasesFuncScxDupl,
			&autogenTestCasesFuncScxEgyp,
			&autogenTestCasesFuncScxElba,
			&autogenTestCasesFuncScxElym,
			&autogenTestCasesFuncScxEthi,
			&autogenTestCasesFuncScxGeor,
			&autogenTestCasesFuncScxGlag,
			&autogenTestCasesFuncScxGong,
			&autogenTestCasesFuncScxGonm,
			&autogenTestCasesFuncScxGoth,
			&autogenTestCasesFuncScxGran,
			&autogenTestCasesFuncScxGrek,
			&autogenTestCasesFuncScxGujr,
			&autogenTestCasesFuncScxGuru,
			&autogenTestCasesFuncScxHang,
			&autogenTestCasesFuncScxHani,
			&autogenTestCasesFuncScxHano,
			&autogenTestCasesFuncScxHatr,
			&autogenTestCasesFuncScxHebr,
			&autogenTestCasesFuncScxHira,
			&autogenTestCasesFuncScxHluw,
			&autogenTestCasesFuncScxHmng,
			&autogenTestCasesFuncScxHmnp,
			&autogenTestCasesFuncScxHung,
			&autogenTestCasesFuncScxItal,
			&autogenTestCasesFuncScxJava,
			&autogenTestCasesFuncScxKali,
			&autogenTestCasesFuncScxKana,
			&autogenTestCasesFuncScxKawi,
			&autogenTestCasesFuncScxKhar,
			&autogenTestCasesFuncScxKhmr,
			&autogenTestCasesFuncScxKhoj,
			&autogenTestCasesFuncScxKits,
			&autogenTestCasesFuncScxKnda,
			&autogenTestCasesFuncScxKthi,
			&autogenTestCasesFuncScxLana,
			&autogenTestCasesFuncScxLaoo,
			&autogenTestCasesFuncScxLatn,
			&autogenTestCasesFuncScxLepc,
			&autogenTestCasesFuncScxLimb,
			&autogenTestCasesFuncScxLina,
			&autogenTestCasesFuncScxLinb,
			&autogenTestCasesFuncScxLisu,
			&autogenTestCasesFuncScxLyci,
			&autogenTestCasesFuncScxLydi,
			&autogenTestCasesFuncScxMahj,
			&autogenTestCasesFuncScxMaka,
			&autogenTestCasesFuncScxMand,
			&autogenTestCasesFuncScxMani,
			&autogenTestCasesFuncScxMarc,
			&autogenTestCasesFuncScxMedf,
			&autogenTestCasesFuncScxMend,
			&autogenTestCasesFuncScxMerc,
			&autogenTestCasesFuncScxMero,
			&autogenTestCasesFuncScxMlym,
			&autogenTestCasesFuncScxModi,
			&autogenTestCasesFuncScxMong,
			&autogenTestCasesFuncScxMroo,
			&autogenTestCasesFuncScxMtei,
			&autogenTestCasesFuncScxMult,
			&autogenTestCasesFuncScxMymr,
			&autogenTestCasesFuncScxNagm,
			&autogenTestCasesFuncScxNand,
			&autogenTestCasesFuncScxNarb,
			&autogenTestCasesFuncScxNbat,
			&autogenTestCasesFuncScxNewa,
			&autogenTestCasesFuncScxNkoo,
			&autogenTestCasesFuncScxNshu,
			&autogenTestCasesFuncScxOgam,
			&autogenTestCasesFuncScxOlck,
			&autogenTestCasesFuncScxOrkh,
			&autogenTestCasesFuncScxOrya,
			&autogenTestCasesFuncScxOsge,
			&autogenTestCasesFuncScxOsma,
			&autogenTestCasesFuncScxOugr,
			&autogenTestCasesFuncScxPalm,
			&autogenTestCasesFuncScxPauc,
			&autogenTestCasesFuncScxPerm,
			&autogenTestCasesFuncScxPhag,
			&autogenTestCasesFuncScxPhli,
			&autogenTestCasesFuncScxPhlp,
			&autogenTestCasesFuncScxPhnx,
			&autogenTestCasesFuncScxPlrd,
			&autogenTestCasesFuncScxPrti,
			&autogenTestCasesFuncScxRjng,
			&autogenTestCasesFuncScxRohg,
			&autogenTestCasesFuncScxRunr,
			&autogenTestCasesFuncScxSamr,
			&autogenTestCasesFuncScxSarb,
			&autogenTestCasesFuncScxSaur,
			&autogenTestCasesFuncScxSgnw,
			&autogenTestCasesFuncScxShaw,
			&autogenTestCasesFuncScxShrd,
			&autogenTestCasesFuncScxSidd,
			&autogenTestCasesFuncScxSind,
			&autogenTestCasesFuncScxSinh,
			&autogenTestCasesFuncScxSogd,
			&autogenTestCasesFuncScxSogo,
			&autogenTestCasesFuncScxSora,
			&autogenTestCasesFuncScxSoyo,
			&autogenTestCasesFuncScxSund,
			&autogenTestCasesFuncScxSylo,
			&autogenTestCasesFuncScxSyrc,
			&autogenTestCasesFuncScxTagb,
			&autogenTestCasesFuncScxTakr,
			&autogenTestCasesFuncScxTale,
			&autogenTestCasesFuncScxTalu,
			&autogenTestCasesFuncScxTaml,
			&autogenTestCasesFuncScxTang,
			&autogenTestCasesFuncScxTavt,
			&autogenTestCasesFuncScxTelu,
			&autogenTestCasesFuncScxTfng,
			&autogenTestCasesFuncScxTglg,
			&autogenTestCasesFuncScxThaa,
			&autogenTestCasesFuncScxThai,
			&autogenTestCasesFuncScxTibt,
			&autogenTestCasesFuncScxTirh,
			&autogenTestCasesFuncScxTnsa,
			&autogenTestCasesFuncScxToto,
			&autogenTestCasesFuncScxUgar,
			&autogenTestCasesFuncScxVaii,
			&autogenTestCasesFuncScxVith,
			&autogenTestCasesFuncScxWara,
			&autogenTestCasesFuncScxWcho,
			&autogenTestCasesFuncScxXpeo,
			&autogenTestCasesFuncScxXsux,
			&autogenTestCasesFuncScxYezi,
			&autogenTestCasesFuncScxYiii,
			&autogenTestCasesFuncScxZanb,
			&autogenTestCasesFuncScxZinh,
			&autogenTestCasesFuncScxZyyy,
			&autogenTestCasesFuncScxZzzz,
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
