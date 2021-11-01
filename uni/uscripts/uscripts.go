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
***************************************************************************/

/*
Package upatids provides queries of common Unicode(R) properties.

For example, the function, HasGcLu(rune r) bool, returns true only if the
code point, r, has the Unicode General_Category value of Lu.

The file, README.md, provides a complete list of all functions provided by
this package along with which Unicode properties and values it tests for.

The file, SUMMARY.txt, in the parent directory (top level directory of this
module), can be perused to look up Unicode property names/values and which
functions test it within this and other packages of this module.

Unicode and the Unicode Logo are registered trademarks of Unicode, Inc. in the
United States and other countries.
*/
package uscripts

// UnicodeVersion specifies the version of the Unicode Standard that this
// package is compliant with if it has no bugs.
const UnicodeVersion = "14.0.0"

// HasScxAdlm returns true if r has Unicode property Script_Extensions = Adlam.
func HasScxAdlm(r rune) bool {
	switch r {
	case 1567, 1600, 125278, 125279:
		return true
	}
	return r >= 125184 && (r <= 125259 || r >= 125264) && r <= 125273
}

// HasScxAghb returns true if r has Unicode property Script_Extensions = Caucasian_Albanian.
func HasScxAghb(r rune) bool {
	switch r {
	case 66927:
		return true
	}
	return r >= 66864 && r <= 66915
}

// HasScxAhom returns true if r has Unicode property Script_Extensions = Ahom.
func HasScxAhom(r rune) bool {
	return r >= 71424 && (r <= 71450 || r >= 71453 && (r <= 71467 || r >= 71472)) &&
		r <= 71494
}

// HasScxArab returns true if r has Unicode property Script_Extensions = Arabic.
func HasScxArab(r rune) bool {
	switch r {
	case 1536, 1537, 1538, 1539, 1540, 1563, 1564, 1565, 1566, 2192, 2193, 64975,
		65021, 65022, 65023, 65136, 65137, 65138, 65139, 65140, 126464, 126465,
		126466, 126467, 126497, 126498, 126500, 126503, 126516, 126517, 126518,
		126519, 126521, 126523, 126530, 126535, 126537, 126539, 126541, 126542,
		126543, 126545, 126546, 126548, 126551, 126553, 126555, 126557, 126559,
		126561, 126562, 126564, 126567, 126568, 126569, 126570, 126580, 126581,
		126582, 126583, 126585, 126586, 126587, 126588, 126590, 126625, 126626,
		126627, 126629, 126630, 126631, 126632, 126633, 126704, 126705:
		return true
	}
	return r >= 1542 && (r <= 64911 && (r <= 1791 && (r <= 1610 && (r <= 1599 ||
		r >= 1600) || r >= 1611 && (r <= 1647 || r >= 1648 && (r <= 1756 || r >= 1758))) ||
		r >= 1872 && (r <= 2273 && (r <= 1919 || r >= 2160 && (r <= 2190 || r >= 2200)) ||
			r >= 2275 && (r <= 2303 || r >= 64336 && (r <= 64450 || r >= 64467)))) ||
		r >= 64914 && (r <= 69246 && (r <= 65020 && (r <= 64967 || r >= 65008) ||
			r >= 65142 && (r <= 65276 || r >= 66272 && (r <= 66299 || r >= 69216))) ||
			r >= 126469 && (r <= 126578 && (r <= 126495 || r >= 126505 && (r <= 126514 ||
				r >= 126572)) || r >= 126592 && (r <= 126601 || r >= 126603 && (r <= 126619 ||
				r >= 126635)) && r <= 126651) && r <= 126651) && r <= 126651) && r <= 126651
}

// HasScxArmi returns true if r has Unicode property Script_Extensions = Imperial_Aramaic.
func HasScxArmi(r rune) bool {
	return r >= 67648 && (r <= 67669 || r >= 67671) && r <= 67679
}

// HasScxArmn returns true if r has Unicode property Script_Extensions = Armenian.
func HasScxArmn(r rune) bool {
	switch r {
	case 1421, 1422, 1423, 64275, 64276, 64277, 64278, 64279:
		return true
	}
	return r >= 1329 && (r <= 1366 || r >= 1369) && r <= 1418
}

// HasScxAvst returns true if r has Unicode property Script_Extensions = Avestan.
func HasScxAvst(r rune) bool {
	return r >= 68352 && (r <= 68405 || r >= 68409) && r <= 68415
}

// HasScxBali returns true if r has Unicode property Script_Extensions = Balinese.
func HasScxBali(r rune) bool {
	return r >= 6912 && (r <= 6988 || r >= 6992) && r <= 7038
}

// HasScxBamu returns true if r has Unicode property Script_Extensions = Bamum.
func HasScxBamu(r rune) bool {
	return r >= 42656 && (r <= 42743 || r >= 92160) && r <= 92728
}

// HasScxBass returns true if r has Unicode property Script_Extensions = Bassa_Vah.
func HasScxBass(r rune) bool {
	return r >= 92880 && (r <= 92909 || r >= 92912) && r <= 92917
}

// HasScxBatk returns true if r has Unicode property Script_Extensions = Batak.
func HasScxBatk(r rune) bool {
	switch r {
	case 7164, 7165, 7166, 7167:
		return true
	}
	return r >= 7104 && r <= 7155
}

// HasScxBeng returns true if r has Unicode property Script_Extensions = Bengali.
func HasScxBeng(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 2432, 2433, 2434, 2435, 2447, 2448, 2482, 2486,
		2487, 2488, 2489, 2503, 2504, 2507, 2508, 2509, 2510, 2519, 2524, 2525,
		2527, 2528, 2529, 2530, 2531, 7376, 7378, 7381, 7382, 7384, 7393, 7402,
		7405, 7410, 7413, 7414, 7415, 43249:
		return true
	}
	return r >= 2437 && (r <= 2472 && (r <= 2444 || r >= 2451) || r >= 2474 &&
		(r <= 2480 || r >= 2492 && (r <= 2500 || r >= 2534)) && r <= 2558) && r <= 2558
}

// HasScxBhks returns true if r has Unicode property Script_Extensions = Bhaiksuki.
func HasScxBhks(r rune) bool {
	return r >= 72704 && (r <= 72758 && (r <= 72712 || r >= 72714) || r >= 72760 &&
		(r <= 72773 || r >= 72784) && r <= 72812) && r <= 72812
}

// HasScxBopo returns true if r has Unicode property Script_Extensions = Bopomofo.
func HasScxBopo(r rune) bool {
	switch r {
	case 746, 747, 12289, 12290, 12291, 12316, 12317, 12318, 12319, 12330, 12331,
		12332, 12333, 12336, 12343, 12539, 65093, 65094, 65377, 65378, 65379, 65380,
		65381:
		return true
	}
	return r >= 12296 && (r <= 12315 && (r <= 12305 || r >= 12307) || r >= 12549 &&
		(r <= 12591 || r >= 12704) && r <= 12735) && r <= 12735
}

// HasScxBrah returns true if r has Unicode property Script_Extensions = Brahmi.
func HasScxBrah(r rune) bool {
	switch r {
	case 69759:
		return true
	}
	return r >= 69632 && (r <= 69709 || r >= 69714) && r <= 69749
}

// HasScxBrai returns true if r has Unicode property Script_Extensions = Braille.
func HasScxBrai(r rune) bool {
	return r >= 10240 && r <= 10495
}

// HasScxBugi returns true if r has Unicode property Script_Extensions = Buginese.
func HasScxBugi(r rune) bool {
	switch r {
	case 6686, 6687, 43471:
		return true
	}
	return r >= 6656 && r <= 6683
}

// HasScxBuhd returns true if r has Unicode property Script_Extensions = Buhid.
func HasScxBuhd(r rune) bool {
	switch r {
	case 5941, 5942:
		return true
	}
	return r >= 5952 && r <= 5971
}

// HasScxCakm returns true if r has Unicode property Script_Extensions = Chakma.
func HasScxCakm(r rune) bool {
	return r >= 2534 && (r <= 4169 && (r <= 2543 || r >= 4160) || r >= 69888 &&
		(r <= 69940 || r >= 69942) && r <= 69959) && r <= 69959
}

// HasScxCans returns true if r has Unicode property Script_Extensions = Canadian_Aboriginal.
func HasScxCans(r rune) bool {
	return r >= 5120 && (r <= 5759 || r >= 6320 && (r <= 6389 || r >= 72368)) &&
		r <= 72383
}

// HasScxCari returns true if r has Unicode property Script_Extensions = Carian.
func HasScxCari(r rune) bool {
	return r >= 66208 && r <= 66256
}

// HasScxCham returns true if r has Unicode property Script_Extensions = Cham.
func HasScxCham(r rune) bool {
	switch r {
	case 43612, 43613, 43614, 43615:
		return true
	}
	return r >= 43520 && (r <= 43574 || r >= 43584 && (r <= 43597 || r >= 43600)) &&
		r <= 43609
}

// HasScxCher returns true if r has Unicode property Script_Extensions = Cherokee.
func HasScxCher(r rune) bool {
	return r >= 5024 && (r <= 5109 || r >= 5112 && (r <= 5117 || r >= 43888)) &&
		r <= 43967
}

// HasScxChrs returns true if r has Unicode property Script_Extensions = Chorasmian.
func HasScxChrs(r rune) bool {
	return r >= 69552 && r <= 69579
}

// HasScxCopt returns true if r has Unicode property Script_Extensions = Coptic.
func HasScxCopt(r rune) bool {
	return r >= 994 && (r <= 11507 && (r <= 1007 || r >= 11392) || r >= 11513 &&
		(r <= 11519 || r >= 66272) && r <= 66299) && r <= 66299
}

// HasScxCpmn returns true if r has Unicode property Script_Extensions = Cypro_Minoan.
func HasScxCpmn(r rune) bool {
	switch r {
	case 65792, 65793:
		return true
	}
	return r >= 77712 && r <= 77810
}

// HasScxCprt returns true if r has Unicode property Script_Extensions = Cypriot.
func HasScxCprt(r rune) bool {
	switch r {
	case 65792, 65793, 65794, 67592, 67639, 67640, 67644, 67647:
		return true
	}
	return r >= 65799 && (r <= 65855 && (r <= 65843 || r >= 65847) || r >= 67584 &&
		(r <= 67589 || r >= 67594) && r <= 67637) && r <= 67637
}

// HasScxCyrl returns true if r has Unicode property Script_Extensions = Cyrillic.
func HasScxCyrl(r rune) bool {
	switch r {
	case 7467, 7544, 7672, 11843, 65070, 65071:
		return true
	}
	return r >= 1024 && (r <= 7304 && (r <= 1327 || r >= 7296) || r >= 11744 &&
		(r <= 11775 || r >= 42560) && r <= 42655) && r <= 42655
}

// HasScxDeva returns true if r has Unicode property Script_Extensions = Devanagari.
func HasScxDeva(r rune) bool {
	switch r {
	case 7384, 7385, 7386, 7387, 7388, 7389, 7390, 7391, 7392, 7402, 7403, 7404,
		7405, 7406, 7407, 7408, 7409, 7410, 7411, 7412, 7413, 7414, 7416, 7417,
		8432:
		return true
	}
	return r >= 2304 && (r <= 2405 && (r <= 2386 || r >= 2389) || r >= 2406 &&
		(r <= 7383 && (r <= 2431 || r >= 7376) || r >= 7393 && (r <= 7414 || r >= 43056 &&
			(r <= 43065 || r >= 43232)) && r <= 43263) && r <= 43263) && r <= 43263
}

// HasScxDiak returns true if r has Unicode property Script_Extensions = Dives_Akuru.
func HasScxDiak(r rune) bool {
	switch r {
	case 71945, 71957, 71958, 71991, 71992:
		return true
	}
	return r >= 71936 && (r <= 71955 && (r <= 71942 || r >= 71948) || r >= 71960 &&
		(r <= 71989 || r >= 71995 && (r <= 72006 || r >= 72016)) && r <= 72025) && r <= 72025
}

// HasScxDogr returns true if r has Unicode property Script_Extensions = Dogra.
func HasScxDogr(r rune) bool {
	return r >= 2404 && (r <= 2415 || r >= 43056 && (r <= 43065 || r >= 71680)) &&
		r <= 71739
}

// HasScxDsrt returns true if r has Unicode property Script_Extensions = Deseret.
func HasScxDsrt(r rune) bool {
	return r >= 66560 && r <= 66639
}

// HasScxDupl returns true if r has Unicode property Script_Extensions = Duployan.
func HasScxDupl(r rune) bool {
	return r >= 113664 && (r <= 113788 && (r <= 113770 || r >= 113776) ||
		r >= 113792 && (r <= 113800 || r >= 113808 && (r <= 113817 || r >= 113820)) &&
			r <= 113827) && r <= 113827
}

// HasScxEgyp returns true if r has Unicode property Script_Extensions = Egyptian_Hieroglyphs.
func HasScxEgyp(r rune) bool {
	return r >= 77824 && (r <= 78894 || r >= 78896) && r <= 78904
}

// HasScxElba returns true if r has Unicode property Script_Extensions = Elbasan.
func HasScxElba(r rune) bool {
	return r >= 66816 && r <= 66855
}

// HasScxElym returns true if r has Unicode property Script_Extensions = Elymaic.
func HasScxElym(r rune) bool {
	return r >= 69600 && r <= 69622
}

// HasScxEthi returns true if r has Unicode property Script_Extensions = Ethiopic.
func HasScxEthi(r rune) bool {
	switch r {
	case 4682, 4683, 4684, 4685, 4696, 4698, 4699, 4700, 4701, 4746, 4747, 4748,
		4749, 4786, 4787, 4788, 4789, 4800, 4802, 4803, 4804, 4805, 4882, 4883,
		4884, 4885, 124904, 124905, 124906, 124907, 124909, 124910:
		return true
	}
	return r >= 4608 && (r <= 11694 && (r <= 4822 && (r <= 4744 && (r <= 4680 ||
		r >= 4688 && (r <= 4694 || r >= 4704)) || r >= 4752 && (r <= 4784 || r >= 4792 &&
		(r <= 4798 || r >= 4808))) || r >= 4824 && (r <= 4988 && (r <= 4880 || r >= 4888 &&
		(r <= 4954 || r >= 4957)) || r >= 4992 && (r <= 11670 && (r <= 5017 ||
		r >= 11648) || r >= 11680 && (r <= 11686 || r >= 11688)))) || r >= 11696 &&
		(r <= 11742 && (r <= 11718 && (r <= 11702 || r >= 11704 && (r <= 11710 ||
			r >= 11712)) || r >= 11720 && (r <= 11726 || r >= 11728 && (r <= 11734 ||
			r >= 11736))) || r >= 43777 && (r <= 43798 && (r <= 43782 || r >= 43785 &&
			(r <= 43790 || r >= 43793)) || r >= 43808 && (r <= 43822 && (r <= 43814 ||
			r >= 43816) || r >= 124896 && (r <= 124902 || r >= 124912) && r <= 124926) &&
			r <= 124926) && r <= 124926) && r <= 124926) && r <= 124926
}

// HasScxGeor returns true if r has Unicode property Script_Extensions = Georgian.
func HasScxGeor(r rune) bool {
	switch r {
	case 4295, 4301, 7357, 7358, 7359, 11559, 11565:
		return true
	}
	return r >= 4256 && (r <= 4351 && (r <= 4293 || r >= 4304) || r >= 7312 &&
		(r <= 7354 || r >= 11520) && r <= 11557) && r <= 11557
}

// HasScxGlag returns true if r has Unicode property Script_Extensions = Glagolitic.
func HasScxGlag(r rune) bool {
	switch r {
	case 1156, 1159, 11843, 42607, 122915, 122916, 122918, 122919, 122920,
		122921, 122922:
		return true
	}
	return r >= 11264 && (r <= 122886 && (r <= 11359 || r >= 122880) || r >= 122888 &&
		(r <= 122904 || r >= 122907) && r <= 122913) && r <= 122913
}

// HasScxGong returns true if r has Unicode property Script_Extensions = Gunjala_Gondi.
func HasScxGong(r rune) bool {
	switch r {
	case 2404, 2405, 73063, 73064, 73104, 73105:
		return true
	}
	return r >= 73056 && (r <= 73102 && (r <= 73061 || r >= 73066) || r >= 73107 &&
		(r <= 73112 || r >= 73120) && r <= 73129) && r <= 73129
}

// HasScxGonm returns true if r has Unicode property Script_Extensions = Masaram_Gondi.
func HasScxGonm(r rune) bool {
	switch r {
	case 2404, 2405, 72968, 72969, 73018, 73020, 73021:
		return true
	}
	return r >= 72960 && (r <= 73014 && (r <= 72966 || r >= 72971) || r >= 73023 &&
		(r <= 73031 || r >= 73040) && r <= 73049) && r <= 73049
}

// HasScxGoth returns true if r has Unicode property Script_Extensions = Gothic.
func HasScxGoth(r rune) bool {
	return r >= 66352 && r <= 66378
}

// HasScxGran returns true if r has Unicode property Script_Extensions = Grantha.
func HasScxGran(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 7376, 7378, 7379, 7410, 7411, 7412, 7416, 7417,
		8432, 70400, 70401, 70402, 70403, 70415, 70416, 70450, 70451, 70453, 70454,
		70455, 70456, 70457, 70471, 70472, 70475, 70476, 70477, 70480, 70487, 70512,
		70513, 70514, 70515, 70516, 73680, 73681, 73683:
		return true
	}
	return r >= 3046 && (r <= 70440 && (r <= 3059 || r >= 70405 && (r <= 70412 ||
		r >= 70419)) || r >= 70442 && (r <= 70468 && (r <= 70448 || r >= 70459) ||
		r >= 70493 && (r <= 70499 || r >= 70502) && r <= 70508) && r <= 70508) && r <= 70508
}

// HasScxGrek returns true if r has Unicode property Script_Extensions = Greek.
func HasScxGrek(r rune) bool {
	switch r {
	case 834, 837, 880, 881, 882, 883, 885, 886, 887, 890, 891, 892, 893, 895,
		900, 902, 904, 905, 906, 908, 7462, 7463, 7464, 7465, 7466, 7517, 7518,
		7519, 7520, 7521, 7526, 7527, 7528, 7529, 7530, 7615, 7616, 7617, 8025,
		8027, 8029, 8178, 8179, 8180, 8486, 43877, 65952:
		return true
	}
	return r >= 910 && (r <= 8023 && (r <= 7957 && (r <= 993 && (r <= 929 || r >= 931) ||
		r >= 1008 && (r <= 1023 || r >= 7936)) || r >= 7960 && (r <= 8005 && (r <= 7965 ||
		r >= 7968) || r >= 8008 && (r <= 8013 || r >= 8016))) || r >= 8031 && (r <= 8147 &&
		(r <= 8116 && (r <= 8061 || r >= 8064) || r >= 8118 && (r <= 8132 || r >= 8134)) ||
		r >= 8150 && (r <= 8175 && (r <= 8155 || r >= 8157) || r >= 8182 && (r <= 8190 ||
			r >= 65856 && (r <= 65934 || r >= 119296)) && r <= 119365) && r <= 119365) &&
		r <= 119365) && r <= 119365
}

// HasScxGujr returns true if r has Unicode property Script_Extensions = Gujarati.
func HasScxGujr(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 2689, 2690, 2691, 2703, 2704, 2705, 2738, 2739,
		2741, 2742, 2743, 2744, 2745, 2759, 2760, 2761, 2763, 2764, 2765, 2768,
		2784, 2785, 2786, 2787:
		return true
	}
	return r >= 2693 && (r <= 2736 && (r <= 2701 || r >= 2707 && (r <= 2728 ||
		r >= 2730)) || r >= 2748 && (r <= 2801 && (r <= 2757 || r >= 2790) || r >= 2809 &&
		(r <= 2815 || r >= 43056) && r <= 43065) && r <= 43065) && r <= 43065
}

// HasScxGuru returns true if r has Unicode property Script_Extensions = Gurmukhi.
func HasScxGuru(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 2561, 2562, 2563, 2575, 2576, 2610, 2611, 2613,
		2614, 2616, 2617, 2620, 2622, 2623, 2624, 2625, 2626, 2631, 2632, 2635,
		2636, 2637, 2641, 2649, 2650, 2651, 2652, 2654:
		return true
	}
	return r >= 2565 && (r <= 2600 && (r <= 2570 || r >= 2579) || r >= 2602 &&
		(r <= 2608 || r >= 2662 && (r <= 2678 || r >= 43056)) && r <= 43065) && r <= 43065
}

// HasScxHang returns true if r has Unicode property Script_Extensions = Hangul.
func HasScxHang(r rune) bool {
	switch r {
	case 12289, 12290, 12291, 12316, 12317, 12318, 12319, 12334, 12335, 12336,
		12343, 12539, 65093, 65094, 65377, 65378, 65379, 65380, 65381, 65498, 65499,
		65500:
		return true
	}
	return r >= 4352 && (r <= 43388 && (r <= 12315 && (r <= 4607 || r >= 12296 &&
		(r <= 12305 || r >= 12307)) || r >= 12593 && (r <= 12830 && (r <= 12686 ||
		r >= 12800) || r >= 12896 && (r <= 12926 || r >= 43360))) || r >= 44032 &&
		(r <= 55291 && (r <= 55203 || r >= 55216 && (r <= 55238 || r >= 55243)) ||
			r >= 65440 && (r <= 65479 && (r <= 65470 || r >= 65474) || r >= 65482 &&
				(r <= 65487 || r >= 65490) && r <= 65495) && r <= 65495) && r <= 65495) && r <= 65495
}

// HasScxHani returns true if r has Unicode property Script_Extensions = Han.
func HasScxHani(r rune) bool {
	switch r {
	case 12289, 12290, 12291, 12293, 12294, 12295, 12316, 12317, 12318, 12319,
		12336, 12343, 12344, 12345, 12346, 12347, 12348, 12349, 12350, 12351, 12539,
		13055, 13179, 13180, 13181, 13182, 13183, 65093, 65094, 65377, 65378, 65379,
		65380, 65381, 94178, 94179, 94192, 94193, 127568, 127569:
		return true
	}
	return r >= 11904 && (r <= 13310 && (r <= 12333 && (r <= 12245 && (r <= 11929 ||
		r >= 11931 && (r <= 12019 || r >= 12032)) || r >= 12296 && (r <= 12305 ||
		r >= 12307 && (r <= 12319 || r >= 12321))) || r >= 12688 && (r <= 12871 &&
		(r <= 12703 || r >= 12736 && (r <= 12771 || r >= 12832)) || r >= 12928 &&
		(r <= 13003 && (r <= 12976 || r >= 12992) || r >= 13144 && (r <= 13168 ||
			r >= 13280)))) || r >= 13312 && (r <= 119665 && (r <= 42759 && (r <= 19903 ||
		r >= 19968 && (r <= 40959 || r >= 42752)) || r >= 63744 && (r <= 64109 ||
		r >= 64112 && (r <= 64217 || r >= 119648))) || r >= 131072 && (r <= 178205 &&
		(r <= 173791 || r >= 173824 && (r <= 177976 || r >= 177984)) || r >= 178208 &&
		(r <= 191456 && (r <= 183969 || r >= 183984) || r >= 194560 && (r <= 195101 ||
			r >= 196608) && r <= 201546) && r <= 201546) && r <= 201546) && r <= 201546) &&
		r <= 201546
}

// HasScxHano returns true if r has Unicode property Script_Extensions = Hanunoo.
func HasScxHano(r rune) bool {
	return r >= 5920 && r <= 5942
}

// HasScxHatr returns true if r has Unicode property Script_Extensions = Hatran.
func HasScxHatr(r rune) bool {
	switch r {
	case 67828, 67829, 67835, 67836, 67837, 67838, 67839:
		return true
	}
	return r >= 67808 && r <= 67826
}

// HasScxHebr returns true if r has Unicode property Script_Extensions = Hebrew.
func HasScxHebr(r rune) bool {
	switch r {
	case 64312, 64313, 64314, 64315, 64316, 64318, 64320, 64321, 64323, 64324:
		return true
	}
	return r >= 1425 && (r <= 1514 && (r <= 1479 || r >= 1488) || r >= 1519 &&
		(r <= 1524 || r >= 64285 && (r <= 64310 || r >= 64326)) && r <= 64335) && r <= 64335
}

// HasScxHira returns true if r has Unicode property Script_Extensions = Hiragana.
func HasScxHira(r rune) bool {
	switch r {
	case 12289, 12290, 12291, 12316, 12317, 12318, 12319, 12343, 12348, 12349,
		12448, 12539, 12540, 65093, 65094, 65377, 65378, 65379, 65380, 65381, 65392,
		65438, 65439, 110928, 110929, 110930, 127488:
		return true
	}
	return r >= 12296 && (r <= 12341 && (r <= 12305 || r >= 12307 && (r <= 12319 ||
		r >= 12336)) || r >= 12353 && (r <= 12438 || r >= 12441 && (r <= 12448 ||
		r >= 110593)) && r <= 110879) && r <= 110879
}

// HasScxHluw returns true if r has Unicode property Script_Extensions = Anatolian_Hieroglyphs.
func HasScxHluw(r rune) bool {
	return r >= 82944 && r <= 83526
}

// HasScxHmng returns true if r has Unicode property Script_Extensions = Pahawh_Hmong.
func HasScxHmng(r rune) bool {
	return r >= 92928 && (r <= 93017 && (r <= 92997 || r >= 93008) || r >= 93019 &&
		(r <= 93025 || r >= 93027 && (r <= 93047 || r >= 93053)) && r <= 93071) && r <= 93071
}

// HasScxHmnp returns true if r has Unicode property Script_Extensions = Nyiakeng_Puachue_Hmong.
func HasScxHmnp(r rune) bool {
	switch r {
	case 123214, 123215:
		return true
	}
	return r >= 123136 && (r <= 123180 || r >= 123184 && (r <= 123197 ||
		r >= 123200)) && r <= 123209
}

// HasScxHung returns true if r has Unicode property Script_Extensions = Old_Hungarian.
func HasScxHung(r rune) bool {
	return r >= 68736 && (r <= 68786 || r >= 68800 && (r <= 68850 || r >= 68858)) &&
		r <= 68863
}

// HasScxItal returns true if r has Unicode property Script_Extensions = Old_Italic.
func HasScxItal(r rune) bool {
	switch r {
	case 66349, 66350, 66351:
		return true
	}
	return r >= 66304 && r <= 66339
}

// HasScxJava returns true if r has Unicode property Script_Extensions = Javanese.
func HasScxJava(r rune) bool {
	switch r {
	case 43486, 43487:
		return true
	}
	return r >= 43392 && (r <= 43469 || r >= 43471) && r <= 43481
}

// HasScxKali returns true if r has Unicode property Script_Extensions = Kayah_Li.
func HasScxKali(r rune) bool {
	return r >= 43264 && r <= 43311
}

// HasScxKana returns true if r has Unicode property Script_Extensions = Katakana.
func HasScxKana(r rune) bool {
	switch r {
	case 12289, 12290, 12291, 12316, 12317, 12318, 12319, 12343, 12348, 12349,
		12441, 12442, 12443, 12444, 12539, 12540, 12541, 12542, 12543, 65093, 65094,
		65438, 65439, 110576, 110577, 110578, 110579, 110589, 110590, 110592,
		110880, 110881, 110882, 110948, 110949, 110950, 110951:
		return true
	}
	return r >= 12296 && (r <= 12538 && (r <= 12315 && (r <= 12305 || r >= 12307) ||
		r >= 12336 && (r <= 12341 || r >= 12448)) || r >= 12784 && (r <= 13054 &&
		(r <= 12799 || r >= 13008) || r >= 13056 && (r <= 13143 || r >= 65377 &&
		(r <= 65437 || r >= 110581)) && r <= 110587) && r <= 110587) && r <= 110587
}

// HasScxKhar returns true if r has Unicode property Script_Extensions = Kharoshthi.
func HasScxKhar(r rune) bool {
	switch r {
	case 68096, 68097, 68098, 68099, 68101, 68102, 68117, 68118, 68119, 68152,
		68153, 68154:
		return true
	}
	return r >= 68108 && (r <= 68149 && (r <= 68115 || r >= 68121) || r >= 68159 &&
		(r <= 68168 || r >= 68176) && r <= 68184) && r <= 68184
}

// HasScxKhmr returns true if r has Unicode property Script_Extensions = Khmer.
func HasScxKhmr(r rune) bool {
	return r >= 6016 && (r <= 6121 && (r <= 6109 || r >= 6112) || r >= 6128 &&
		(r <= 6137 || r >= 6624) && r <= 6655) && r <= 6655
}

// HasScxKhoj returns true if r has Unicode property Script_Extensions = Khojki.
func HasScxKhoj(r rune) bool {
	return r >= 2790 && (r <= 43065 && (r <= 2799 || r >= 43056) || r >= 70144 &&
		(r <= 70161 || r >= 70163) && r <= 70206) && r <= 70206
}

// HasScxKits returns true if r has Unicode property Script_Extensions = Khitan_Small_Script.
func HasScxKits(r rune) bool {
	switch r {
	case 94180:
		return true
	}
	return r >= 101120 && r <= 101589
}

// HasScxKnda returns true if r has Unicode property Script_Extensions = Kannada.
func HasScxKnda(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 3214, 3215, 3216, 3253, 3254, 3255, 3256, 3257,
		3270, 3271, 3272, 3274, 3275, 3276, 3277, 3285, 3286, 3293, 3294, 3296,
		3297, 3298, 3299, 3313, 3314, 7376, 7378, 7386, 7410, 7412:
		return true
	}
	return r >= 3200 && (r <= 3251 && (r <= 3212 || r >= 3218 && (r <= 3240 ||
		r >= 3242)) || r >= 3260 && (r <= 3268 || r >= 3302 && (r <= 3311 || r >= 43056)) &&
		r <= 43061) && r <= 43061
}

// HasScxKthi returns true if r has Unicode property Script_Extensions = Kaithi.
func HasScxKthi(r rune) bool {
	switch r {
	case 69837:
		return true
	}
	return r >= 2406 && (r <= 2415 || r >= 43056 && (r <= 43065 || r >= 69760)) &&
		r <= 69826
}

// HasScxLana returns true if r has Unicode property Script_Extensions = Tai_Tham.
func HasScxLana(r rune) bool {
	return r >= 6688 && (r <= 6780 && (r <= 6750 || r >= 6752) || r >= 6783 &&
		(r <= 6793 || r >= 6800 && (r <= 6809 || r >= 6816)) && r <= 6829) && r <= 6829
}

// HasScxLaoo returns true if r has Unicode property Script_Extensions = Lao.
func HasScxLaoo(r rune) bool {
	switch r {
	case 3713, 3714, 3716, 3718, 3719, 3720, 3721, 3722, 3749, 3776, 3777, 3778,
		3779, 3780, 3782, 3804, 3805, 3806, 3807:
		return true
	}
	return r >= 3724 && (r <= 3773 && (r <= 3747 || r >= 3751) || r >= 3784 &&
		(r <= 3789 || r >= 3792) && r <= 3801) && r <= 3801
}

// HasScxLatn returns true if r has Unicode property Script_Extensions = Latin.
func HasScxLatn(r rune) bool {
	switch r {
	case 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82,
		83, 84, 85, 86, 87, 88, 89, 90, 97, 98, 99, 100, 101, 102, 103, 104, 105,
		106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
		121, 122, 170, 186, 736, 737, 738, 739, 740, 1157, 1158, 2385, 2386, 4347,
		7522, 7523, 7524, 7525, 8239, 8305, 8319, 8432, 8490, 8491, 8498, 8526,
		42960, 42961, 42963, 42965, 42966, 42967, 42968, 42969, 43310, 43878, 43879,
		43880, 43881:
		return true
	}
	return r >= 192 && (r <= 11391 && (r <= 7516 && (r <= 696 && (r <= 214 || r >= 216 &&
		(r <= 246 || r >= 248)) || r >= 867 && (r <= 879 || r >= 7424 && (r <= 7461 ||
		r >= 7468))) || r >= 7531 && (r <= 7935 && (r <= 7543 || r >= 7545 && (r <= 7614 ||
		r >= 7680)) || r >= 8336 && (r <= 8348 || r >= 8544 && (r <= 8584 ||
		r >= 11360)))) || r >= 42752 && (r <= 43876 && (r <= 42954 && (r <= 42759 ||
		r >= 42786 && (r <= 42887 || r >= 42891)) || r >= 42994 && (r <= 43007 ||
		r >= 43824 && (r <= 43866 || r >= 43868))) || r >= 64256 && (r <= 65370 &&
		(r <= 64262 || r >= 65313 && (r <= 65338 || r >= 65345)) || r >= 67456 &&
		(r <= 67504 && (r <= 67461 || r >= 67463) || r >= 67506 && (r <= 67514 ||
			r >= 122624) && r <= 122654) && r <= 122654) && r <= 122654) && r <= 122654) &&
		r <= 122654
}

// HasScxLepc returns true if r has Unicode property Script_Extensions = Lepcha.
func HasScxLepc(r rune) bool {
	switch r {
	case 7245, 7246, 7247:
		return true
	}
	return r >= 7168 && (r <= 7223 || r >= 7227) && r <= 7241
}

// HasScxLimb returns true if r has Unicode property Script_Extensions = Limbu.
func HasScxLimb(r rune) bool {
	switch r {
	case 2405, 6464:
		return true
	}
	return r >= 6400 && (r <= 6443 && (r <= 6430 || r >= 6432) || r >= 6448 &&
		(r <= 6459 || r >= 6468) && r <= 6479) && r <= 6479
}

// HasScxLina returns true if r has Unicode property Script_Extensions = Linear_A.
func HasScxLina(r rune) bool {
	return r >= 65799 && (r <= 67382 && (r <= 65843 || r >= 67072) || r >= 67392 &&
		(r <= 67413 || r >= 67424) && r <= 67431) && r <= 67431
}

// HasScxLinb returns true if r has Unicode property Script_Extensions = Linear_B.
func HasScxLinb(r rune) bool {
	switch r {
	case 65596, 65597, 65792, 65793, 65794:
		return true
	}
	return r >= 65536 && (r <= 65613 && (r <= 65574 && (r <= 65547 || r >= 65549) ||
		r >= 65576 && (r <= 65594 || r >= 65599)) || r >= 65616 && (r <= 65786 &&
		(r <= 65629 || r >= 65664) || r >= 65799 && (r <= 65843 || r >= 65847) &&
		r <= 65855) && r <= 65855) && r <= 65855
}

// HasScxLisu returns true if r has Unicode property Script_Extensions = Lisu.
func HasScxLisu(r rune) bool {
	switch r {
	case 73648:
		return true
	}
	return r >= 42192 && r <= 42239
}

// HasScxLyci returns true if r has Unicode property Script_Extensions = Lycian.
func HasScxLyci(r rune) bool {
	return r >= 66176 && r <= 66204
}

// HasScxLydi returns true if r has Unicode property Script_Extensions = Lydian.
func HasScxLydi(r rune) bool {
	switch r {
	case 67903:
		return true
	}
	return r >= 67872 && r <= 67897
}

// HasScxMahj returns true if r has Unicode property Script_Extensions = Mahajani.
func HasScxMahj(r rune) bool {
	return r >= 2404 && (r <= 2415 || r >= 43056 && (r <= 43065 || r >= 69968)) &&
		r <= 70006
}

// HasScxMaka returns true if r has Unicode property Script_Extensions = Makasar.
func HasScxMaka(r rune) bool {
	return r >= 73440 && r <= 73464
}

// HasScxMand returns true if r has Unicode property Script_Extensions = Mandaic.
func HasScxMand(r rune) bool {
	switch r {
	case 1600, 2142:
		return true
	}
	return r >= 2112 && r <= 2139
}

// HasScxMani returns true if r has Unicode property Script_Extensions = Manichaean.
func HasScxMani(r rune) bool {
	switch r {
	case 1600:
		return true
	}
	return r >= 68288 && (r <= 68326 || r >= 68331) && r <= 68342
}

// HasScxMarc returns true if r has Unicode property Script_Extensions = Marchen.
func HasScxMarc(r rune) bool {
	return r >= 72816 && (r <= 72847 || r >= 72850 && (r <= 72871 || r >= 72873)) &&
		r <= 72886
}

// HasScxMedf returns true if r has Unicode property Script_Extensions = Medefaidrin.
func HasScxMedf(r rune) bool {
	return r >= 93760 && r <= 93850
}

// HasScxMend returns true if r has Unicode property Script_Extensions = Mende_Kikakui.
func HasScxMend(r rune) bool {
	return r >= 124928 && (r <= 125124 || r >= 125127) && r <= 125142
}

// HasScxMerc returns true if r has Unicode property Script_Extensions = Meroitic_Cursive.
func HasScxMerc(r rune) bool {
	return r >= 68000 && (r <= 68023 || r >= 68028 && (r <= 68047 || r >= 68050)) &&
		r <= 68095
}

// HasScxMero returns true if r has Unicode property Script_Extensions = Meroitic_Hieroglyphs.
func HasScxMero(r rune) bool {
	return r >= 67968 && r <= 67999
}

// HasScxMlym returns true if r has Unicode property Script_Extensions = Malayalam.
func HasScxMlym(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 3342, 3343, 3344, 3398, 3399, 3400, 7386, 43056,
		43057, 43058:
		return true
	}
	return r >= 3328 && (r <= 3396 && (r <= 3340 || r >= 3346) || r >= 3402 &&
		(r <= 3407 || r >= 3412 && (r <= 3427 || r >= 3430)) && r <= 3455) && r <= 3455
}

// HasScxModi returns true if r has Unicode property Script_Extensions = Modi.
func HasScxModi(r rune) bool {
	return r >= 43056 && (r <= 43065 || r >= 71168 && (r <= 71236 || r >= 71248)) &&
		r <= 71257
}

// HasScxMong returns true if r has Unicode property Script_Extensions = Mongolian.
func HasScxMong(r rune) bool {
	switch r {
	case 8239:
		return true
	}
	return r >= 6144 && (r <= 6264 && (r <= 6169 || r >= 6176) || r >= 6272 &&
		(r <= 6314 || r >= 71264) && r <= 71276) && r <= 71276
}

// HasScxMroo returns true if r has Unicode property Script_Extensions = Mro.
func HasScxMroo(r rune) bool {
	switch r {
	case 92782, 92783:
		return true
	}
	return r >= 92736 && (r <= 92766 || r >= 92768) && r <= 92777
}

// HasScxMtei returns true if r has Unicode property Script_Extensions = Meetei_Mayek.
func HasScxMtei(r rune) bool {
	return r >= 43744 && (r <= 43766 || r >= 43968 && (r <= 44013 || r >= 44016)) &&
		r <= 44025
}

// HasScxMult returns true if r has Unicode property Script_Extensions = Multani.
func HasScxMult(r rune) bool {
	switch r {
	case 70280, 70282, 70283, 70284, 70285:
		return true
	}
	return r >= 2662 && (r <= 70278 && (r <= 2671 || r >= 70272) || r >= 70287 &&
		(r <= 70301 || r >= 70303) && r <= 70313) && r <= 70313
}

// HasScxMymr returns true if r has Unicode property Script_Extensions = Myanmar.
func HasScxMymr(r rune) bool {
	switch r {
	case 43310:
		return true
	}
	return r >= 4096 && (r <= 4255 || r >= 43488 && (r <= 43518 || r >= 43616)) &&
		r <= 43647
}

// HasScxNand returns true if r has Unicode property Script_Extensions = Nandinagari.
func HasScxNand(r rune) bool {
	switch r {
	case 2404, 2405, 7401, 7410, 7418:
		return true
	}
	return r >= 3302 && (r <= 43061 && (r <= 3311 || r >= 43056) || r >= 72096 &&
		(r <= 72103 || r >= 72106 && (r <= 72151 || r >= 72154)) && r <= 72164) && r <= 72164
}

// HasScxNarb returns true if r has Unicode property Script_Extensions = Old_North_Arabian.
func HasScxNarb(r rune) bool {
	return r >= 68224 && r <= 68255
}

// HasScxNbat returns true if r has Unicode property Script_Extensions = Nabataean.
func HasScxNbat(r rune) bool {
	return r >= 67712 && (r <= 67742 || r >= 67751) && r <= 67759
}

// HasScxNewa returns true if r has Unicode property Script_Extensions = Newa.
func HasScxNewa(r rune) bool {
	switch r {
	case 70749, 70750, 70751, 70752, 70753:
		return true
	}
	return r >= 70656 && r <= 70747
}

// HasScxNkoo returns true if r has Unicode property Script_Extensions = Nko.
func HasScxNkoo(r rune) bool {
	switch r {
	case 1548, 1563, 1567, 2045, 2046, 2047, 64830, 64831:
		return true
	}
	return r >= 1984 && r <= 2042
}

// HasScxNshu returns true if r has Unicode property Script_Extensions = Nushu.
func HasScxNshu(r rune) bool {
	switch r {
	case 94177:
		return true
	}
	return r >= 110960 && r <= 111355
}

// HasScxOgam returns true if r has Unicode property Script_Extensions = Ogham.
func HasScxOgam(r rune) bool {
	return r >= 5760 && r <= 5788
}

// HasScxOlck returns true if r has Unicode property Script_Extensions = Ol_Chiki.
func HasScxOlck(r rune) bool {
	return r >= 7248 && r <= 7295
}

// HasScxOrkh returns true if r has Unicode property Script_Extensions = Old_Turkic.
func HasScxOrkh(r rune) bool {
	return r >= 68608 && r <= 68680
}

// HasScxOrya returns true if r has Unicode property Script_Extensions = Oriya.
func HasScxOrya(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 2817, 2818, 2819, 2831, 2832, 2866, 2867, 2869,
		2870, 2871, 2872, 2873, 2887, 2888, 2891, 2892, 2893, 2901, 2902, 2903,
		2908, 2909, 2911, 2912, 2913, 2914, 2915, 7386, 7410:
		return true
	}
	return r >= 2821 && (r <= 2856 && (r <= 2828 || r >= 2835) || r >= 2858 &&
		(r <= 2864 || r >= 2876 && (r <= 2884 || r >= 2918)) && r <= 2935) && r <= 2935
}

// HasScxOsge returns true if r has Unicode property Script_Extensions = Osage.
func HasScxOsge(r rune) bool {
	return r >= 66736 && (r <= 66771 || r >= 66776) && r <= 66811
}

// HasScxOsma returns true if r has Unicode property Script_Extensions = Osmanya.
func HasScxOsma(r rune) bool {
	return r >= 66688 && (r <= 66717 || r >= 66720) && r <= 66729
}

// HasScxOugr returns true if r has Unicode property Script_Extensions = Old_Uyghur.
func HasScxOugr(r rune) bool {
	switch r {
	case 1600, 68338:
		return true
	}
	return r >= 69488 && r <= 69513
}

// HasScxPalm returns true if r has Unicode property Script_Extensions = Palmyrene.
func HasScxPalm(r rune) bool {
	return r >= 67680 && r <= 67711
}

// HasScxPauc returns true if r has Unicode property Script_Extensions = Pau_Cin_Hau.
func HasScxPauc(r rune) bool {
	return r >= 72384 && r <= 72440
}

// HasScxPerm returns true if r has Unicode property Script_Extensions = Old_Permic.
func HasScxPerm(r rune) bool {
	switch r {
	case 1155:
		return true
	}
	return r >= 66384 && r <= 66426
}

// HasScxPhag returns true if r has Unicode property Script_Extensions = Phags_Pa.
func HasScxPhag(r rune) bool {
	switch r {
	case 6146, 6147, 6149:
		return true
	}
	return r >= 43072 && r <= 43127
}

// HasScxPhli returns true if r has Unicode property Script_Extensions = Inscriptional_Pahlavi.
func HasScxPhli(r rune) bool {
	return r >= 68448 && (r <= 68466 || r >= 68472) && r <= 68479
}

// HasScxPhlp returns true if r has Unicode property Script_Extensions = Psalter_Pahlavi.
func HasScxPhlp(r rune) bool {
	switch r {
	case 1600, 68505, 68506, 68507, 68508:
		return true
	}
	return r >= 68480 && (r <= 68497 || r >= 68521) && r <= 68527
}

// HasScxPhnx returns true if r has Unicode property Script_Extensions = Phoenician.
func HasScxPhnx(r rune) bool {
	switch r {
	case 67871:
		return true
	}
	return r >= 67840 && r <= 67867
}

// HasScxPlrd returns true if r has Unicode property Script_Extensions = Miao.
func HasScxPlrd(r rune) bool {
	return r >= 93952 && (r <= 94026 || r >= 94031 && (r <= 94087 || r >= 94095)) &&
		r <= 94111
}

// HasScxPrti returns true if r has Unicode property Script_Extensions = Inscriptional_Parthian.
func HasScxPrti(r rune) bool {
	return r >= 68416 && (r <= 68437 || r >= 68440) && r <= 68447
}

// HasScxRjng returns true if r has Unicode property Script_Extensions = Rejang.
func HasScxRjng(r rune) bool {
	switch r {
	case 43359:
		return true
	}
	return r >= 43312 && r <= 43347
}

// HasScxRohg returns true if r has Unicode property Script_Extensions = Hanifi_Rohingya.
func HasScxRohg(r rune) bool {
	switch r {
	case 1548, 1563, 1567, 1600, 1748:
		return true
	}
	return r >= 68864 && (r <= 68903 || r >= 68912) && r <= 68921
}

// HasScxRunr returns true if r has Unicode property Script_Extensions = Runic.
func HasScxRunr(r rune) bool {
	return r >= 5792 && (r <= 5866 || r >= 5870) && r <= 5880
}

// HasScxSamr returns true if r has Unicode property Script_Extensions = Samaritan.
func HasScxSamr(r rune) bool {
	return r >= 2048 && (r <= 2093 || r >= 2096) && r <= 2110
}

// HasScxSarb returns true if r has Unicode property Script_Extensions = Old_South_Arabian.
func HasScxSarb(r rune) bool {
	return r >= 68192 && r <= 68223
}

// HasScxSaur returns true if r has Unicode property Script_Extensions = Saurashtra.
func HasScxSaur(r rune) bool {
	return r >= 43136 && (r <= 43205 || r >= 43214) && r <= 43225
}

// HasScxSgnw returns true if r has Unicode property Script_Extensions = SignWriting.
func HasScxSgnw(r rune) bool {
	switch r {
	case 121499, 121500, 121501, 121502, 121503:
		return true
	}
	return r >= 120832 && (r <= 121483 || r >= 121505) && r <= 121519
}

// HasScxShaw returns true if r has Unicode property Script_Extensions = Shavian.
func HasScxShaw(r rune) bool {
	return r >= 66640 && r <= 66687
}

// HasScxShrd returns true if r has Unicode property Script_Extensions = Sharada.
func HasScxShrd(r rune) bool {
	switch r {
	case 2385, 7383, 7385, 7388, 7389, 7392:
		return true
	}
	return r >= 70016 && r <= 70111
}

// HasScxSidd returns true if r has Unicode property Script_Extensions = Siddham.
func HasScxSidd(r rune) bool {
	return r >= 71040 && (r <= 71093 || r >= 71096) && r <= 71133
}

// HasScxSind returns true if r has Unicode property Script_Extensions = Khudawadi.
func HasScxSind(r rune) bool {
	switch r {
	case 2404, 2405:
		return true
	}
	return r >= 43056 && (r <= 43065 || r >= 70320 && (r <= 70378 || r >= 70384)) &&
		r <= 70393
}

// HasScxSinh returns true if r has Unicode property Script_Extensions = Sinhala.
func HasScxSinh(r rune) bool {
	switch r {
	case 2404, 2405, 3457, 3458, 3459, 3517, 3530, 3542, 3570, 3571, 3572:
		return true
	}
	return r >= 3461 && (r <= 3526 && (r <= 3505 && (r <= 3478 || r >= 3482) ||
		r >= 3507 && (r <= 3515 || r >= 3520)) || r >= 3535 && (r <= 3551 && (r <= 3540 ||
		r >= 3544) || r >= 3558 && (r <= 3567 || r >= 70113) && r <= 70132) &&
		r <= 70132) && r <= 70132
}

// HasScxSogd returns true if r has Unicode property Script_Extensions = Sogdian.
func HasScxSogd(r rune) bool {
	switch r {
	case 1600:
		return true
	}
	return r >= 69424 && r <= 69465
}

// HasScxSogo returns true if r has Unicode property Script_Extensions = Old_Sogdian.
func HasScxSogo(r rune) bool {
	return r >= 69376 && r <= 69415
}

// HasScxSora returns true if r has Unicode property Script_Extensions = Sora_Sompeng.
func HasScxSora(r rune) bool {
	return r >= 69840 && (r <= 69864 || r >= 69872) && r <= 69881
}

// HasScxSoyo returns true if r has Unicode property Script_Extensions = Soyombo.
func HasScxSoyo(r rune) bool {
	return r >= 72272 && r <= 72354
}

// HasScxSund returns true if r has Unicode property Script_Extensions = Sundanese.
func HasScxSund(r rune) bool {
	return r >= 7040 && (r <= 7103 || r >= 7360) && r <= 7367
}

// HasScxSylo returns true if r has Unicode property Script_Extensions = Syloti_Nagri.
func HasScxSylo(r rune) bool {
	switch r {
	case 2404, 2405:
		return true
	}
	return r >= 2534 && (r <= 2543 || r >= 43008) && r <= 43052
}

// HasScxSyrc returns true if r has Unicode property Script_Extensions = Syriac.
func HasScxSyrc(r rune) bool {
	switch r {
	case 1548, 1563, 1564, 1567, 1600, 1648, 1869, 1870, 1871, 7672, 7674:
		return true
	}
	return r >= 1611 && (r <= 1805 && (r <= 1621 || r >= 1792) || r >= 1807 &&
		(r <= 1866 || r >= 2144) && r <= 2154) && r <= 2154
}

// HasScxTagb returns true if r has Unicode property Script_Extensions = Tagbanwa.
func HasScxTagb(r rune) bool {
	switch r {
	case 5941, 5942, 5998, 5999, 6000, 6002, 6003:
		return true
	}
	return r >= 5984 && r <= 5996
}

// HasScxTakr returns true if r has Unicode property Script_Extensions = Takri.
func HasScxTakr(r rune) bool {
	switch r {
	case 2404, 2405:
		return true
	}
	return r >= 43056 && (r <= 43065 || r >= 71296 && (r <= 71353 || r >= 71360)) &&
		r <= 71369
}

// HasScxTale returns true if r has Unicode property Script_Extensions = Tai_Le.
func HasScxTale(r rune) bool {
	switch r {
	case 6512, 6513, 6514, 6515, 6516:
		return true
	}
	return r >= 4160 && (r <= 4169 || r >= 6480) && r <= 6509
}

// HasScxTalu returns true if r has Unicode property Script_Extensions = New_Tai_Lue.
func HasScxTalu(r rune) bool {
	switch r {
	case 6622, 6623:
		return true
	}
	return r >= 6528 && (r <= 6571 || r >= 6576 && (r <= 6601 || r >= 6608)) && r <= 6618
}

// HasScxTaml returns true if r has Unicode property Script_Extensions = Tamil.
func HasScxTaml(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 2946, 2947, 2958, 2959, 2960, 2962, 2963, 2964,
		2965, 2969, 2970, 2972, 2974, 2975, 2979, 2980, 2984, 2985, 2986, 3006,
		3007, 3008, 3009, 3010, 3014, 3015, 3016, 3018, 3019, 3020, 3021, 3024,
		3031, 7386, 43251, 70401, 70403, 70459, 70460, 73727:
		return true
	}
	return r >= 2949 && (r <= 3001 && (r <= 2954 || r >= 2990) || r >= 3046 &&
		(r <= 3066 || r >= 73664) && r <= 73713) && r <= 73713
}

// HasScxTang returns true if r has Unicode property Script_Extensions = Tangut.
func HasScxTang(r rune) bool {
	switch r {
	case 94176:
		return true
	}
	return r >= 94208 && (r <= 100343 || r >= 100352 && (r <= 101119 ||
		r >= 101632)) && r <= 101640
}

// HasScxTavt returns true if r has Unicode property Script_Extensions = Tai_Viet.
func HasScxTavt(r rune) bool {
	switch r {
	case 43739, 43740, 43741, 43742, 43743:
		return true
	}
	return r >= 43648 && r <= 43714
}

// HasScxTelu returns true if r has Unicode property Script_Extensions = Telugu.
func HasScxTelu(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 3086, 3087, 3088, 3142, 3143, 3144, 3146, 3147,
		3148, 3149, 3157, 3158, 3160, 3161, 3162, 3165, 3168, 3169, 3170, 3171,
		7386, 7410:
		return true
	}
	return r >= 3072 && (r <= 3129 && (r <= 3084 || r >= 3090 && (r <= 3112 ||
		r >= 3114)) || r >= 3132 && (r <= 3140 || r >= 3174 && (r <= 3183 || r >= 3191)) &&
		r <= 3199) && r <= 3199
}

// HasScxTfng returns true if r has Unicode property Script_Extensions = Tifinagh.
func HasScxTfng(r rune) bool {
	switch r {
	case 11631, 11632, 11647:
		return true
	}
	return r >= 11568 && r <= 11623
}

// HasScxTglg returns true if r has Unicode property Script_Extensions = Tagalog.
func HasScxTglg(r rune) bool {
	switch r {
	case 5919, 5941, 5942:
		return true
	}
	return r >= 5888 && r <= 5909
}

// HasScxThaa returns true if r has Unicode property Script_Extensions = Thaana.
func HasScxThaa(r rune) bool {
	switch r {
	case 1548, 1563, 1564, 1567, 65010, 65021:
		return true
	}
	return r >= 1632 && (r <= 1641 || r >= 1920) && r <= 1969
}

// HasScxThai returns true if r has Unicode property Script_Extensions = Thai.
func HasScxThai(r rune) bool {
	return r >= 3585 && (r <= 3642 || r >= 3648) && r <= 3675
}

// HasScxTibt returns true if r has Unicode property Script_Extensions = Tibetan.
func HasScxTibt(r rune) bool {
	switch r {
	case 4057, 4058:
		return true
	}
	return r >= 3840 && (r <= 3991 && (r <= 3911 || r >= 3913 && (r <= 3948 ||
		r >= 3953)) || r >= 3993 && (r <= 4028 || r >= 4030 && (r <= 4044 || r >= 4046)) &&
		r <= 4052) && r <= 4052
}

// HasScxTirh returns true if r has Unicode property Script_Extensions = Tirhuta.
func HasScxTirh(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 7410:
		return true
	}
	return r >= 43056 && (r <= 43065 || r >= 70784 && (r <= 70855 || r >= 70864)) &&
		r <= 70873
}

// HasScxTnsa returns true if r has Unicode property Script_Extensions = Tangsa.
func HasScxTnsa(r rune) bool {
	return r >= 92784 && (r <= 92862 || r >= 92864) && r <= 92873
}

// HasScxToto returns true if r has Unicode property Script_Extensions = Toto.
func HasScxToto(r rune) bool {
	return r >= 123536 && r <= 123566
}

// HasScxUgar returns true if r has Unicode property Script_Extensions = Ugaritic.
func HasScxUgar(r rune) bool {
	switch r {
	case 66463:
		return true
	}
	return r >= 66432 && r <= 66461
}

// HasScxVaii returns true if r has Unicode property Script_Extensions = Vai.
func HasScxVaii(r rune) bool {
	return r >= 42240 && r <= 42539
}

// HasScxVith returns true if r has Unicode property Script_Extensions = Vithkuqi.
func HasScxVith(r rune) bool {
	switch r {
	case 66964, 66965, 67003, 67004:
		return true
	}
	return r >= 66928 && (r <= 66962 && (r <= 66938 || r >= 66940 && (r <= 66954 ||
		r >= 66956)) || r >= 66967 && (r <= 66977 || r >= 66979 && (r <= 66993 ||
		r >= 66995)) && r <= 67001) && r <= 67001
}

// HasScxWara returns true if r has Unicode property Script_Extensions = Warang_Citi.
func HasScxWara(r rune) bool {
	switch r {
	case 71935:
		return true
	}
	return r >= 71840 && r <= 71922
}

// HasScxWcho returns true if r has Unicode property Script_Extensions = Wancho.
func HasScxWcho(r rune) bool {
	switch r {
	case 123647:
		return true
	}
	return r >= 123584 && r <= 123641
}

// HasScxXpeo returns true if r has Unicode property Script_Extensions = Old_Persian.
func HasScxXpeo(r rune) bool {
	return r >= 66464 && (r <= 66499 || r >= 66504) && r <= 66517
}

// HasScxXsux returns true if r has Unicode property Script_Extensions = Cuneiform.
func HasScxXsux(r rune) bool {
	switch r {
	case 74864, 74865, 74866, 74867, 74868:
		return true
	}
	return r >= 73728 && (r <= 74649 || r >= 74752 && (r <= 74862 || r >= 74880)) &&
		r <= 75075
}

// HasScxYezi returns true if r has Unicode property Script_Extensions = Yezidi.
func HasScxYezi(r rune) bool {
	switch r {
	case 1548, 1563, 1567, 69291, 69292, 69293, 69296, 69297:
		return true
	}
	return r >= 1632 && (r <= 1641 || r >= 69248) && r <= 69289
}

// HasScxYiii returns true if r has Unicode property Script_Extensions = Yi.
func HasScxYiii(r rune) bool {
	switch r {
	case 12289, 12290, 12539, 65377, 65378, 65379, 65380, 65381:
		return true
	}
	return r >= 12296 && (r <= 12315 && (r <= 12305 || r >= 12308) || r >= 40960 &&
		(r <= 42124 || r >= 42128) && r <= 42182) && r <= 42182
}

// HasScxZanb returns true if r has Unicode property Script_Extensions = Zanabazar_Square.
func HasScxZanb(r rune) bool {
	return r >= 72192 && r <= 72263
}

// HasScxZinh returns true if r has Unicode property Script_Extensions = Inherited.
func HasScxZinh(r rune) bool {
	switch r {
	case 835, 836, 2387, 2388, 7673, 7675, 7676, 7677, 7678, 7679, 8204, 8205,
		66045, 119143, 119144, 119145, 119210, 119211, 119212, 119213:
		return true
	}
	return r >= 768 && (r <= 65039 && (r <= 6862 && (r <= 833 || r >= 838 && (r <= 866 ||
		r >= 6832)) || r >= 7618 && (r <= 7671 || r >= 8400 && (r <= 8431 ||
		r >= 65024))) || r >= 65056 && (r <= 118598 && (r <= 65069 || r >= 118528 &&
		(r <= 118573 || r >= 118576)) || r >= 119163 && (r <= 119170 || r >= 119173 &&
		(r <= 119179 || r >= 917760)) && r <= 917999) && r <= 917999) && r <= 917999
}

// HasScxZyyy returns true if r has Unicode property Script_Extensions = Common.
func HasScxZyyy(r rune) bool {
	switch r {
	case 91, 92, 93, 94, 95, 96, 187, 188, 189, 190, 191, 215, 247, 741, 742,
		743, 744, 745, 884, 894, 901, 903, 1541, 1757, 2274, 3647, 4053, 4054, 4055,
		4056, 5867, 5868, 5869, 8487, 8488, 8489, 8585, 8586, 8587, 12288, 12292,
		12306, 12320, 12342, 12927, 13004, 13005, 13006, 13007, 13311, 42888, 42889,
		42890, 43867, 43882, 43883, 65128, 65129, 65130, 65131, 65279, 65529, 65530,
		65531, 65532, 65533, 119171, 119172, 119966, 119967, 119970, 119973, 119974,
		119977, 119978, 119979, 119980, 119995, 120071, 120072, 120073, 120074,
		120123, 120124, 120125, 120126, 120128, 120129, 120130, 120131, 120132,
		120134, 127489, 127490, 129008, 129200, 129201, 129648, 129649, 129650,
		129651, 129652, 129656, 129657, 129658, 129659, 129660, 917505:
		return true
	}
	return r >= 0 && (r <= 119539 && (r <= 12283 && (r <= 8384 && (r <= 8203 &&
		(r <= 185 && (r <= 64 || r >= 123 && (r <= 169 || r >= 171)) || r >= 697 &&
			(r <= 735 || r >= 748 && (r <= 767 || r >= 8192))) || r >= 8206 && (r <= 8304 &&
		(r <= 8238 || r >= 8240 && (r <= 8292 || r >= 8294)) || r >= 8308 && (r <= 8318 ||
		r >= 8320 && (r <= 8334 || r >= 8352)))) || r >= 8448 && (r <= 9290 && (r <= 8525 &&
		(r <= 8485 || r >= 8492 && (r <= 8497 || r >= 8499)) || r >= 8527 && (r <= 8543 ||
		r >= 8592 && (r <= 9254 || r >= 9280))) || r >= 9312 && (r <= 11157 &&
		(r <= 10239 || r >= 10496 && (r <= 11123 || r >= 11126)) || r >= 11159 &&
		(r <= 11842 && (r <= 11263 || r >= 11776) || r >= 11844 && (r <= 11869 ||
			r >= 12272))))) || r >= 12872 && (r <= 65344 && (r <= 42785 && (r <= 13178 &&
		(r <= 12895 || r >= 12977 && (r <= 12991 || r >= 13169)) || r >= 13184 &&
		(r <= 13279 || r >= 19904 && (r <= 19967 || r >= 42760))) || r >= 65040 &&
		(r <= 65106 && (r <= 65049 || r >= 65072 && (r <= 65092 || r >= 65095)) ||
			r >= 65108 && (r <= 65126 || r >= 65281 && (r <= 65312 || r >= 65339)))) ||
		r >= 65371 && (r <= 118723 && (r <= 65518 && (r <= 65376 || r >= 65504 &&
			(r <= 65510 || r >= 65512)) || r >= 65936 && (r <= 65948 || r >= 66000 &&
			(r <= 66044 || r >= 118608))) || r >= 118784 && (r <= 119142 && (r <= 119029 ||
			r >= 119040 && (r <= 119078 || r >= 119081)) || r >= 119146 && (r <= 119209 &&
			(r <= 119162 || r >= 119180) || r >= 119214 && (r <= 119274 ||
			r >= 119520)))))) || r >= 119552 && (r <= 127547 && (r <= 120485 &&
		(r <= 120003 && (r <= 119892 && (r <= 119638 || r >= 119666 && (r <= 119672 ||
			r >= 119808)) || r >= 119894 && (r <= 119964 || r >= 119982 && (r <= 119993 ||
			r >= 119997))) || r >= 120005 && (r <= 120092 && (r <= 120069 || r >= 120077 &&
			(r <= 120084 || r >= 120086)) || r >= 120094 && (r <= 120121 || r >= 120138 &&
			(r <= 120144 || r >= 120146)))) || r >= 120488 && (r <= 127123 && (r <= 126132 &&
		(r <= 120779 || r >= 120782 && (r <= 120831 || r >= 126065)) || r >= 126209 &&
		(r <= 126269 || r >= 126976 && (r <= 127019 || r >= 127024))) || r >= 127136 &&
		(r <= 127183 && (r <= 127150 || r >= 127153 && (r <= 127167 || r >= 127169)) ||
			r >= 127185 && (r <= 127405 && (r <= 127221 || r >= 127232) || r >= 127462 &&
				(r <= 127487 || r >= 127504))))) || r >= 127552 && (r <= 129197 && (r <= 128883 &&
		(r <= 128727 && (r <= 127560 || r >= 127584 && (r <= 127589 || r >= 127744)) ||
			r >= 128733 && (r <= 128748 || r >= 128752 && (r <= 128764 || r >= 128768))) ||
		r >= 128896 && (r <= 129035 && (r <= 128984 || r >= 128992 && (r <= 129003 ||
			r >= 129024)) || r >= 129040 && (r <= 129113 && (r <= 129095 || r >= 129104) ||
			r >= 129120 && (r <= 129159 || r >= 129168)))) || r >= 129280 && (r <= 129733 &&
		(r <= 129670 && (r <= 129619 || r >= 129632 && (r <= 129645 || r >= 129664)) ||
			r >= 129680 && (r <= 129708 || r >= 129712 && (r <= 129722 || r >= 129728))) ||
		r >= 129744 && (r <= 129782 && (r <= 129753 || r >= 129760 && (r <= 129767 ||
			r >= 129776)) || r >= 129792 && (r <= 129994 && (r <= 129938 || r >= 129940) ||
			r >= 130032 && (r <= 130041 || r >= 917536) && r <= 917631) && r <= 917631) &&
			r <= 917631) && r <= 917631) && r <= 917631) && r <= 917631) && r <= 917631
}

// HasScxZzzz returns true if r has Unicode property Script_Extensions = Unknown.
func HasScxZzzz(r rune) bool {
	switch r {
	case 888, 889, 896, 897, 898, 899, 907, 909, 930, 1328, 1367, 1368, 1419,
		1420, 1424, 1515, 1516, 1517, 1518, 1806, 1867, 1868, 2043, 2044, 2094,
		2095, 2111, 2140, 2141, 2143, 2155, 2156, 2157, 2158, 2159, 2191, 2436,
		2445, 2446, 2449, 2450, 2473, 2481, 2483, 2484, 2485, 2490, 2491, 2501,
		2502, 2505, 2506, 2520, 2521, 2522, 2523, 2526, 2532, 2533, 2559, 2560,
		2564, 2571, 2572, 2573, 2574, 2577, 2578, 2601, 2609, 2612, 2615, 2618,
		2619, 2621, 2627, 2628, 2629, 2630, 2633, 2634, 2638, 2639, 2640, 2653,
		2692, 2702, 2706, 2729, 2737, 2740, 2746, 2747, 2758, 2762, 2766, 2767,
		2788, 2789, 2816, 2820, 2829, 2830, 2833, 2834, 2857, 2865, 2868, 2874,
		2875, 2885, 2886, 2889, 2890, 2904, 2905, 2906, 2907, 2910, 2916, 2917,
		2948, 2955, 2956, 2957, 2961, 2966, 2967, 2968, 2971, 2973, 2976, 2977,
		2978, 2981, 2982, 2983, 2987, 2988, 2989, 3002, 3003, 3004, 3005, 3011,
		3012, 3013, 3017, 3022, 3023, 3067, 3068, 3069, 3070, 3071, 3085, 3089,
		3113, 3130, 3131, 3141, 3145, 3159, 3163, 3164, 3166, 3167, 3172, 3173,
		3213, 3217, 3241, 3252, 3258, 3259, 3269, 3273, 3295, 3300, 3301, 3312,
		3341, 3345, 3397, 3401, 3408, 3409, 3410, 3411, 3428, 3429, 3456, 3460,
		3479, 3480, 3481, 3506, 3516, 3518, 3519, 3527, 3528, 3529, 3531, 3532,
		3533, 3534, 3541, 3543, 3568, 3569, 3643, 3644, 3645, 3646, 3715, 3717,
		3723, 3748, 3750, 3774, 3775, 3781, 3783, 3790, 3791, 3802, 3803, 3912,
		3949, 3950, 3951, 3952, 3992, 4029, 4045, 4294, 4296, 4297, 4298, 4299,
		4300, 4302, 4303, 4681, 4686, 4687, 4695, 4697, 4702, 4703, 4745, 4750,
		4751, 4785, 4790, 4791, 4799, 4801, 4806, 4807, 4823, 4881, 4886, 4887,
		4955, 4956, 4989, 4990, 4991, 5110, 5111, 5118, 5119, 5789, 5790, 5791,
		5997, 6001, 6110, 6111, 6315, 6316, 6317, 6318, 6319, 6431, 6444, 6445,
		6446, 6447, 6460, 6461, 6462, 6463, 6465, 6466, 6467, 6510, 6511, 6572,
		6573, 6574, 6575, 6619, 6620, 6621, 6684, 6685, 6751, 6781, 6782, 6830,
		6831, 6989, 6990, 6991, 7039, 7224, 7225, 7226, 7242, 7243, 7244, 7355,
		7356, 7419, 7420, 7421, 7422, 7423, 7958, 7959, 7966, 7967, 8006, 8007,
		8014, 8015, 8024, 8026, 8028, 8030, 8062, 8063, 8117, 8133, 8148, 8149,
		8156, 8176, 8177, 8181, 8191, 8293, 8306, 8307, 8335, 8349, 8350, 8351,
		8588, 8589, 8590, 8591, 11124, 11125, 11158, 11508, 11509, 11510, 11511,
		11512, 11558, 11560, 11561, 11562, 11563, 11564, 11566, 11567, 11687, 11695,
		11703, 11711, 11719, 11727, 11735, 11743, 11930, 12284, 12285, 12286, 12287,
		12352, 12439, 12440, 12544, 12545, 12546, 12547, 12548, 12592, 12687, 12831,
		42125, 42126, 42127, 42955, 42956, 42957, 42958, 42959, 42962, 42964, 43053,
		43054, 43055, 43389, 43390, 43391, 43470, 43482, 43483, 43484, 43485, 43519,
		43598, 43599, 43610, 43611, 43783, 43784, 43791, 43792, 43815, 43823, 43884,
		43885, 43886, 43887, 44014, 44015, 55239, 55240, 55241, 55242, 64110, 64111,
		64280, 64281, 64282, 64283, 64284, 64311, 64317, 64319, 64322, 64325, 64912,
		64913, 65107, 65127, 65132, 65133, 65134, 65135, 65141, 65277, 65278, 65280,
		65471, 65472, 65473, 65480, 65481, 65488, 65489, 65496, 65497, 65501, 65502,
		65503, 65511, 65534, 65535, 65548, 65575, 65595, 65598, 65614, 65615, 65787,
		65788, 65789, 65790, 65791, 65795, 65796, 65797, 65798, 65844, 65845, 65846,
		65935, 65949, 65950, 65951, 66205, 66206, 66207, 66300, 66301, 66302, 66303,
		66379, 66380, 66381, 66382, 66383, 66427, 66428, 66429, 66430, 66431, 66462,
		66500, 66501, 66502, 66503, 66718, 66719, 66772, 66773, 66774, 66775, 66812,
		66813, 66814, 66815, 66939, 66955, 66963, 66966, 66978, 66994, 67002, 67462,
		67505, 67590, 67591, 67593, 67638, 67641, 67642, 67643, 67645, 67646, 67670,
		67827, 67830, 67831, 67832, 67833, 67834, 67868, 67869, 67870, 67898, 67899,
		67900, 67901, 67902, 68024, 68025, 68026, 68027, 68048, 68049, 68100, 68103,
		68104, 68105, 68106, 68107, 68116, 68120, 68150, 68151, 68155, 68156, 68157,
		68158, 68327, 68328, 68329, 68330, 68406, 68407, 68408, 68438, 68439, 68467,
		68468, 68469, 68470, 68471, 69247, 69290, 69294, 69295, 69710, 69711, 69712,
		69713, 69838, 69839, 69941, 70112, 70162, 70279, 70281, 70286, 70302, 70379,
		70380, 70381, 70382, 70383, 70404, 70413, 70414, 70417, 70418, 70441, 70449,
		70452, 70458, 70469, 70470, 70473, 70474, 70478, 70479, 70488, 70489, 70490,
		70491, 70492, 70500, 70501, 70509, 70510, 70511, 70748, 71094, 71095, 71451,
		71452, 71468, 71469, 71470, 71471, 71943, 71944, 71946, 71947, 71956, 71959,
		71990, 71993, 71994, 72104, 72105, 72152, 72153, 72713, 72759, 72813, 72814,
		72815, 72848, 72849, 72872, 72967, 72970, 73015, 73016, 73017, 73019, 73022,
		73062, 73065, 73103, 73106, 74863, 78895, 92767, 92778, 92779, 92780, 92781,
		92863, 92910, 92911, 93018, 93026, 93048, 93049, 93050, 93051, 93052, 94027,
		94028, 94029, 94030, 110580, 110588, 110591, 113771, 113772, 113773, 113774,
		113775, 113789, 113790, 113791, 113818, 113819, 118574, 118575, 119079,
		119080, 119893, 119965, 119968, 119969, 119971, 119972, 119975, 119976,
		119981, 119994, 119996, 120004, 120070, 120075, 120076, 120085, 120093,
		120122, 120127, 120133, 120135, 120136, 120137, 120145, 120486, 120487,
		120780, 120781, 121504, 122887, 122905, 122906, 122914, 122917, 123181,
		123182, 123183, 123198, 123199, 123210, 123211, 123212, 123213, 123642,
		123643, 123644, 123645, 123646, 124903, 124908, 124911, 124927, 125125,
		125126, 125260, 125261, 125262, 125263, 125274, 125275, 125276, 125277,
		126468, 126496, 126499, 126501, 126502, 126504, 126515, 126520, 126522,
		126531, 126532, 126533, 126534, 126536, 126538, 126540, 126544, 126547,
		126549, 126550, 126552, 126554, 126556, 126558, 126560, 126563, 126565,
		126566, 126571, 126579, 126584, 126589, 126591, 126602, 126620, 126621,
		126622, 126623, 126624, 126628, 126634, 127020, 127021, 127022, 127023,
		127151, 127152, 127168, 127184, 127548, 127549, 127550, 127551, 128728,
		128729, 128730, 128731, 128732, 128749, 128750, 128751, 128765, 128766,
		128767, 129004, 129005, 129006, 129007, 129036, 129037, 129038, 129039,
		129198, 129199, 129646, 129647, 129653, 129654, 129655, 129661, 129662,
		129663, 129709, 129710, 129711, 129723, 129724, 129725, 129726, 129727,
		129939, 178206, 178207:
		return true
	}
	return r >= 1480 && (r <= 69836 && (r <= 42751 && (r <= 5951 && (r <= 3045 &&
		(r <= 2661 && (r <= 1983 && (r <= 1487 || r >= 1525 && (r <= 1535 || r >= 1970)) ||
			r >= 2194 && (r <= 2518 && (r <= 2199 || r >= 2511) || r >= 2642 && (r <= 2648 ||
				r >= 2655))) || r >= 2679 && (r <= 2808 && (r <= 2688 || r >= 2769 && (r <= 2783 ||
			r >= 2802)) || r >= 2894 && (r <= 2945 && (r <= 2900 || r >= 2936) || r >= 3025 &&
			(r <= 3030 || r >= 3032)))) || r >= 3150 && (r <= 3584 && (r <= 3284 &&
		(r <= 3156 || r >= 3184 && (r <= 3190 || r >= 3278)) || r >= 3287 && (r <= 3327 &&
		(r <= 3292 || r >= 3315) || r >= 3552 && (r <= 3557 || r >= 3573))) || r >= 3676 &&
		(r <= 4095 && (r <= 3712 || r >= 3808 && (r <= 3839 || r >= 4059)) || r >= 5018 &&
			(r <= 5887 && (r <= 5023 || r >= 5881) || r >= 5910 && (r <= 5918 ||
				r >= 5943))))) || r >= 5972 && (r <= 7311 && (r <= 6399 && (r <= 6127 &&
		(r <= 5983 || r >= 6004 && (r <= 6015 || r >= 6122)) || r >= 6138 && (r <= 6175 &&
		(r <= 6143 || r >= 6170) || r >= 6265 && (r <= 6271 || r >= 6390))) || r >= 6517 &&
		(r <= 6799 && (r <= 6527 || r >= 6602 && (r <= 6607 || r >= 6794)) || r >= 6810 &&
			(r <= 6911 && (r <= 6815 || r >= 6863) || r >= 7156 && (r <= 7163 ||
				r >= 7305)))) || r >= 7368 && (r <= 11646 && (r <= 8447 && (r <= 7375 ||
		r >= 8385 && (r <= 8399 || r >= 8433)) || r >= 9255 && (r <= 9311 && (r <= 9279 ||
		r >= 9291) || r >= 11624 && (r <= 11630 || r >= 11633))) || r >= 11671 &&
		(r <= 12271 && (r <= 11903 && (r <= 11679 || r >= 11870) || r >= 12020 &&
			(r <= 12031 || r >= 12246)) || r >= 12772 && (r <= 42191 && (r <= 12783 ||
			r >= 42183) || r >= 42540 && (r <= 42559 || r >= 42744)))))) || r >= 42970 &&
		(r <= 66863 && (r <= 64255 && (r <= 43583 && (r <= 43135 && (r <= 42993 ||
			r >= 43066 && (r <= 43071 || r >= 43128)) || r >= 43206 && (r <= 43231 &&
			(r <= 43213 || r >= 43226) || r >= 43348 && (r <= 43358 || r >= 43575))) ||
			r >= 43715 && (r <= 43807 && (r <= 43738 || r >= 43767 && (r <= 43776 ||
				r >= 43799)) || r >= 44026 && (r <= 55215 && (r <= 44031 || r >= 55204) ||
				r >= 55292 && (r <= 63743 || r >= 64218)))) || r >= 64263 && (r <= 65663 &&
			(r <= 64974 && (r <= 64274 || r >= 64451 && (r <= 64466 || r >= 64968)) ||
				r >= 64976 && (r <= 65055 && (r <= 65007 || r >= 65050) || r >= 65519 &&
					(r <= 65528 || r >= 65630))) || r >= 65953 && (r <= 66271 && (r <= 65999 ||
			r >= 66046 && (r <= 66175 || r >= 66257)) || r >= 66340 && (r <= 66559 &&
			(r <= 66348 || r >= 66518) || r >= 66730 && (r <= 66735 || r >= 66856))))) ||
			r >= 66916 && (r <= 68504 && (r <= 67750 && (r <= 67391 && (r <= 66926 ||
				r >= 67005 && (r <= 67071 || r >= 67383)) || r >= 67414 && (r <= 67455 &&
				(r <= 67423 || r >= 67432) || r >= 67515 && (r <= 67583 || r >= 67743))) ||
				r >= 67760 && (r <= 68175 && (r <= 67807 || r >= 67904 && (r <= 67967 ||
					r >= 68169)) || r >= 68185 && (r <= 68287 && (r <= 68191 || r >= 68256) ||
					r >= 68343 && (r <= 68351 || r >= 68498)))) || r >= 68509 && (r <= 69215 &&
				(r <= 68735 && (r <= 68520 || r >= 68528 && (r <= 68607 || r >= 68681)) ||
					r >= 68787 && (r <= 68857 && (r <= 68799 || r >= 68851) || r >= 68904 &&
						(r <= 68911 || r >= 68922))) || r >= 69298 && (r <= 69551 && (r <= 69423 &&
				(r <= 69375 || r >= 69416) || r >= 69466 && (r <= 69487 || r >= 69514)) ||
				r >= 69580 && (r <= 69631 && (r <= 69599 || r >= 69623) || r >= 69750 &&
					(r <= 69758 || r >= 69827))))))) || r >= 69865 && (r <= 110927 && (r <= 72703 &&
		(r <= 71167 && (r <= 70319 && (r <= 69967 && (r <= 69871 || r >= 69882 &&
			(r <= 69887 || r >= 69960)) || r >= 70007 && (r <= 70143 && (r <= 70015 ||
			r >= 70133) || r >= 70207 && (r <= 70271 || r >= 70314))) || r >= 70394 &&
			(r <= 70655 && (r <= 70399 || r >= 70481 && (r <= 70486 || r >= 70517)) ||
				r >= 70754 && (r <= 70863 && (r <= 70783 || r >= 70856) || r >= 70874 &&
					(r <= 71039 || r >= 71134)))) || r >= 71237 && (r <= 71839 && (r <= 71295 &&
			(r <= 71247 || r >= 71258 && (r <= 71263 || r >= 71277)) || r >= 71354 &&
			(r <= 71423 && (r <= 71359 || r >= 71370) || r >= 71495 && (r <= 71679 ||
				r >= 71740))) || r >= 71923 && (r <= 72095 && (r <= 71934 || r >= 72007 &&
			(r <= 72015 || r >= 72026)) || r >= 72165 && (r <= 72271 && (r <= 72191 ||
			r >= 72264) || r >= 72355 && (r <= 72367 || r >= 72441))))) || r >= 72774 &&
		(r <= 82943 && (r <= 73647 && (r <= 73039 && (r <= 72783 || r >= 72887 &&
			(r <= 72959 || r >= 73032)) || r >= 73050 && (r <= 73119 && (r <= 73055 ||
			r >= 73113) || r >= 73130 && (r <= 73439 || r >= 73465))) || r >= 73649 &&
			(r <= 74751 && (r <= 73663 || r >= 73714 && (r <= 73726 || r >= 74650)) ||
				r >= 74869 && (r <= 77711 && (r <= 74879 || r >= 75076) || r >= 77811 &&
					(r <= 77823 || r >= 78905)))) || r >= 83527 && (r <= 93951 && (r <= 92879 &&
			(r <= 92159 || r >= 92729 && (r <= 92735 || r >= 92874)) || r >= 92918 &&
			(r <= 93007 && (r <= 92927 || r >= 92998) || r >= 93072 && (r <= 93759 ||
				r >= 93851))) || r >= 94088 && (r <= 94207 && (r <= 94175 && (r <= 94094 ||
			r >= 94112) || r >= 94181 && (r <= 94191 || r >= 94194)) || r >= 100344 &&
			(r <= 101631 && (r <= 100351 || r >= 101590) || r >= 101641 && (r <= 110575 ||
				r >= 110883)))))) || r >= 110931 && (r <= 127231 && (r <= 121498 &&
		(r <= 118783 && (r <= 113663 && (r <= 110947 || r >= 110952 && (r <= 110959 ||
			r >= 111356)) || r >= 113801 && (r <= 118527 && (r <= 113807 || r >= 113828) ||
			r >= 118599 && (r <= 118607 || r >= 118724))) || r >= 119030 && (r <= 119519 &&
			(r <= 119039 || r >= 119275 && (r <= 119295 || r >= 119366)) || r >= 119540 &&
			(r <= 119647 && (r <= 119551 || r >= 119639) || r >= 119673 && (r <= 119807 ||
				r >= 121484)))) || r >= 121520 && (r <= 125183 && (r <= 123135 && (r <= 122623 ||
		r >= 122655 && (r <= 122879 || r >= 122923)) || r >= 123216 && (r <= 123583 &&
		(r <= 123535 || r >= 123567) || r >= 123648 && (r <= 124895 || r >= 125143))) ||
		r >= 125280 && (r <= 126529 && (r <= 126208 && (r <= 126064 || r >= 126133) ||
			r >= 126270 && (r <= 126463 || r >= 126524)) || r >= 126652 && (r <= 126975 &&
			(r <= 126703 || r >= 126706) || r >= 127124 && (r <= 127135 || r >= 127222))))) ||
		r >= 127406 && (r <= 129679 && (r <= 128991 && (r <= 127567 && (r <= 127461 ||
			r >= 127491 && (r <= 127503 || r >= 127561)) || r >= 127570 && (r <= 127743 &&
			(r <= 127583 || r >= 127590) || r >= 128884 && (r <= 128895 || r >= 128985))) ||
			r >= 129009 && (r <= 129119 && (r <= 129023 || r >= 129096 && (r <= 129103 ||
				r >= 129114)) || r >= 129160 && (r <= 129279 && (r <= 129167 || r >= 129202) ||
				r >= 129620 && (r <= 129631 || r >= 129671)))) || r >= 129734 && (r <= 173823 &&
			(r <= 129775 && (r <= 129743 || r >= 129754 && (r <= 129759 || r >= 129768)) ||
				r >= 129783 && (r <= 130031 && (r <= 129791 || r >= 129995) || r >= 130042 &&
					(r <= 131071 || r >= 173792))) || r >= 177977 && (r <= 196607 && (r <= 183983 &&
			(r <= 177983 || r >= 183970) || r >= 191457 && (r <= 194559 || r >= 195102)) ||
			r >= 201547 && (r <= 917535 && (r <= 917504 || r >= 917506) || r >= 917632 &&
				(r <= 917759 || r >= 918000) && r <= 1114111) && r <= 1114111) && r <= 1114111) &&
			r <= 1114111) && r <= 1114111) && r <= 1114111) && r <= 1114111) && r <= 1114111
}
