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
	case 125260, 125261, 125262, 125263, 125274, 125275, 125276, 125277:
		return false
	}
	return r >= 125184 && r <= 125279
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
	switch r {
	case 71451, 71452, 71468, 71469, 71470, 71471:
		return false
	}
	return r >= 71424 && r <= 71494
}

// HasScxArab returns true if r has Unicode property Script_Extensions = Arabic.
func HasScxArab(r rune) bool {
	switch r {
	case 1536, 1537, 1538, 1539, 1540, 1563, 1564, 1565, 1566, 2192, 2193, 64975,
		65021, 65022, 65023, 65136, 65137, 65138, 65139, 65140, 126464, 126465,
		126466, 126467, 126497, 126498, 126500, 126503, 126516, 126517, 126518,
		126519, 126521, 126523, 126530, 126535, 126537, 126539, 126541, 126542,
		126543, 126545, 126546, 126548, 126551, 126553, 126555, 126557, 126559,
		126561, 126562, 126564, 126567, 126568, 126569, 126570, 126572, 126573,
		126574, 126575, 126576, 126577, 126578, 126580, 126581, 126582, 126583,
		126585, 126586, 126587, 126588, 126590, 126625, 126626, 126627, 126629,
		126630, 126631, 126632, 126633, 126704, 126705:
		return true
	case 1757, 2191, 2194, 2195, 2196, 2197, 2198, 2199, 2274, 64912, 64913,
		64968, 64969, 64970, 64971, 64972, 64973, 64974, 65141, 126468, 126496,
		126499, 126501, 126502, 126504, 126515, 126520, 126522, 126524, 126525,
		126526, 126527, 126528, 126529, 126531, 126532, 126533, 126534, 126536,
		126538, 126540, 126544, 126547, 126549, 126550, 126552, 126554, 126556,
		126558, 126560, 126563, 126565, 126566, 126571, 126579, 126584, 126589,
		126591, 126602, 126620, 126621, 126622, 126623, 126624, 126628, 126634:
		return false
	}
	return r >= 1542 && (r <= 64911 && (r <= 1919 && (r <= 1791 || r >= 1872) ||
		r >= 2160 && (r <= 2303 || r >= 64336 && (r <= 64450 || r >= 64467))) ||
		r >= 64914 && (r <= 65276 && (r <= 64975 || r >= 65008 && (r <= 65023 ||
			r >= 65142)) || r >= 66272 && (r <= 66299 || r >= 69216 && (r <= 69246 ||
			r >= 126469)))) && r <= 126651
}

// HasScxArmi returns true if r has Unicode property Script_Extensions = Imperial_Aramaic.
func HasScxArmi(r rune) bool {
	switch r {
	case 67670:
		return false
	}
	return r >= 67648 && r <= 67679
}

// HasScxArmn returns true if r has Unicode property Script_Extensions = Armenian.
func HasScxArmn(r rune) bool {
	switch r {
	case 1367, 1368, 1419, 1420:
		return false
	case 1421, 1422, 1423, 64275, 64276, 64277, 64278, 64279:
		return true
	}
	return r >= 1329 && r <= 1423
}

// HasScxAvst returns true if r has Unicode property Script_Extensions = Avestan.
func HasScxAvst(r rune) bool {
	switch r {
	case 68406, 68407, 68408:
		return false
	case 68409, 68410, 68411, 68412, 68413, 68414, 68415:
		return true
	}
	return r >= 68352 && r <= 68415
}

// HasScxBali returns true if r has Unicode property Script_Extensions = Balinese.
func HasScxBali(r rune) bool {
	switch r {
	case 6989, 6990, 6991:
		return false
	}
	return r >= 6912 && r <= 7038
}

// HasScxBamu returns true if r has Unicode property Script_Extensions = Bamum.
func HasScxBamu(r rune) bool {
	return r >= 42656 && (r <= 42743 || r >= 92160) && r <= 92728
}

// HasScxBass returns true if r has Unicode property Script_Extensions = Bassa_Vah.
func HasScxBass(r rune) bool {
	switch r {
	case 92910, 92911:
		return false
	case 92912, 92913, 92914, 92915, 92916, 92917:
		return true
	}
	return r >= 92880 && r <= 92917
}

// HasScxBatk returns true if r has Unicode property Script_Extensions = Batak.
func HasScxBatk(r rune) bool {
	switch r {
	case 7156, 7157, 7158, 7159, 7160, 7161, 7162, 7163:
		return false
	case 7164, 7165, 7166, 7167:
		return true
	}
	return r >= 7104 && r <= 7167
}

// HasScxBeng returns true if r has Unicode property Script_Extensions = Bengali.
func HasScxBeng(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 2432, 2433, 2434, 2435, 2437, 2438, 2439, 2440,
		2441, 2442, 2443, 2444, 2447, 2448, 2474, 2475, 2476, 2477, 2478, 2479,
		2480, 2482, 2486, 2487, 2488, 2489, 2503, 2504, 2507, 2508, 2509, 2510,
		2519, 2524, 2525, 2527, 2528, 2529, 2530, 2531, 7376, 7378, 7381, 7382,
		7384, 7393, 7402, 7405, 7410, 7413, 7414, 7415, 43249:
		return true
	case 2436, 2445, 2446, 2449, 2450, 2473, 2481, 2483, 2484, 2485, 2490, 2491,
		2501, 2502, 2505, 2506, 2511, 2512, 2513, 2514, 2515, 2516, 2517, 2518,
		2520, 2521, 2522, 2523, 2526, 2532, 2533, 7377, 7379, 7380, 7383, 7385,
		7386, 7387, 7388, 7389, 7390, 7391, 7392, 7394, 7395, 7396, 7397, 7398,
		7399, 7400, 7401, 7403, 7404, 7406, 7407, 7408, 7409, 7411, 7412:
		return false
	}
	return r >= 2451 && r <= 2558
}

// HasScxBhks returns true if r has Unicode property Script_Extensions = Bhaiksuki.
func HasScxBhks(r rune) bool {
	switch r {
	case 72713, 72759:
		return false
	}
	return r >= 72704 && (r <= 72773 || r >= 72784) && r <= 72812
}

// HasScxBopo returns true if r has Unicode property Script_Extensions = Bopomofo.
func HasScxBopo(r rune) bool {
	switch r {
	case 746, 747, 12289, 12290, 12291, 12316, 12317, 12318, 12319, 12330, 12331,
		12332, 12333, 12336, 12343, 12539, 65093, 65094, 65377, 65378, 65379, 65380,
		65381:
		return true
	case 12292, 12293, 12294, 12295, 12306, 12334, 12335, 12337, 12338, 12339,
		12340, 12341, 12342:
		return false
	}
	return r >= 12296 && (r <= 12319 || r >= 12549 && (r <= 12591 || r >= 12704)) &&
		r <= 12735
}

// HasScxBrah returns true if r has Unicode property Script_Extensions = Brahmi.
func HasScxBrah(r rune) bool {
	switch r {
	case 69710, 69711, 69712, 69713:
		return false
	case 69759:
		return true
	}
	return r >= 69632 && r <= 69749
}

// HasScxBrai returns true if r has Unicode property Script_Extensions = Braille.
func HasScxBrai(r rune) bool {
	return r >= 10240 && r <= 10495
}

// HasScxBugi returns true if r has Unicode property Script_Extensions = Buginese.
func HasScxBugi(r rune) bool {
	switch r {
	case 6684, 6685:
		return false
	case 6686, 6687, 43471:
		return true
	}
	return r >= 6656 && r <= 6687
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
	switch r {
	case 69941:
		return false
	}
	return r >= 2534 && (r <= 2543 || r >= 4160 && (r <= 4169 || r >= 69888)) &&
		r <= 69959
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
	case 43598, 43599, 43610, 43611:
		return false
	case 43612, 43613, 43614, 43615:
		return true
	}
	return r >= 43520 && (r <= 43574 || r >= 43584) && r <= 43615
}

// HasScxCher returns true if r has Unicode property Script_Extensions = Cherokee.
func HasScxCher(r rune) bool {
	switch r {
	case 5110, 5111:
		return false
	case 5112, 5113, 5114, 5115, 5116, 5117:
		return true
	}
	return r >= 5024 && (r <= 5117 || r >= 43888) && r <= 43967
}

// HasScxChrs returns true if r has Unicode property Script_Extensions = Chorasmian.
func HasScxChrs(r rune) bool {
	return r >= 69552 && r <= 69579
}

// HasScxCopt returns true if r has Unicode property Script_Extensions = Coptic.
func HasScxCopt(r rune) bool {
	switch r {
	case 11508, 11509, 11510, 11511, 11512:
		return false
	case 11513, 11514, 11515, 11516, 11517, 11518, 11519:
		return true
	}
	return r >= 994 && (r <= 1007 || r >= 11392 && (r <= 11519 || r >= 66272)) &&
		r <= 66299
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
	case 65792, 65793, 65794, 67584, 67585, 67586, 67587, 67588, 67589, 67592,
		67639, 67640, 67644, 67647:
		return true
	case 65844, 65845, 65846, 67590, 67591, 67593, 67638, 67641, 67642, 67643,
		67645, 67646:
		return false
	}
	return r >= 65799 && (r <= 65855 || r >= 67594) && r <= 67647
}

// HasScxCyrl returns true if r has Unicode property Script_Extensions = Cyrillic.
func HasScxCyrl(r rune) bool {
	switch r {
	case 7467, 7544, 7672, 11843, 65070, 65071:
		return true
	}
	return r >= 1024 && (r <= 7304 && (r <= 1327 || r >= 7296) || r >= 11744 &&
		(r <= 11775 || r >= 42560)) && r <= 42655
}

// HasScxDeva returns true if r has Unicode property Script_Extensions = Devanagari.
func HasScxDeva(r rune) bool {
	switch r {
	case 2387, 2388, 7415:
		return false
	case 7376, 7377, 7378, 7379, 7380, 7381, 7382, 7383, 7384, 7385, 7386, 7387,
		7388, 7389, 7390, 7391, 7392, 7402, 7403, 7404, 7405, 7406, 7407, 7408,
		7409, 7410, 7411, 7412, 7413, 7414, 7416, 7417, 8432:
		return true
	}
	return r >= 2304 && (r <= 7401 && (r <= 2431 || r >= 7393) || r >= 43056 &&
		(r <= 43065 || r >= 43232)) && r <= 43263
}

// HasScxDiak returns true if r has Unicode property Script_Extensions = Dives_Akuru.
func HasScxDiak(r rune) bool {
	switch r {
	case 71936, 71937, 71938, 71939, 71940, 71941, 71942, 71945, 71948, 71949,
		71950, 71951, 71952, 71953, 71954, 71955, 71957, 71958, 71991, 71992:
		return true
	case 71990, 71993, 71994:
		return false
	}
	return r >= 71960 && (r <= 72006 || r >= 72016) && r <= 72025
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
	switch r {
	case 113771, 113772, 113773, 113774, 113775, 113789, 113790, 113791, 113801,
		113802, 113803, 113804, 113805, 113806, 113807, 113818, 113819:
		return false
	case 113820, 113821, 113822, 113823, 113824, 113825, 113826, 113827:
		return true
	}
	return r >= 113664 && r <= 113827
}

// HasScxEgyp returns true if r has Unicode property Script_Extensions = Egyptian_Hieroglyphs.
func HasScxEgyp(r rune) bool {
	switch r {
	case 78895:
		return false
	}
	return r >= 77824 && r <= 78904
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
	case 4681, 4686, 4687, 4695, 4697, 4702, 4703, 4745, 4750, 4751, 4785, 4790,
		4791, 4799, 4801, 4806, 4807, 4823, 4881, 4886, 4887, 4955, 4956, 4989,
		4990, 4991, 11687, 11695, 11703, 11711, 11719, 11727, 11735, 43783, 43784,
		43791, 43792, 43815, 124903, 124908, 124911:
		return false
	case 4682, 4683, 4684, 4685, 4688, 4689, 4690, 4691, 4692, 4693, 4694, 4696,
		4698, 4699, 4700, 4701, 4746, 4747, 4748, 4749, 4786, 4787, 4788, 4789,
		4792, 4793, 4794, 4795, 4796, 4797, 4798, 4800, 4802, 4803, 4804, 4805,
		4882, 4883, 4884, 4885, 11680, 11681, 11682, 11683, 11684, 11685, 11686,
		11688, 11689, 11690, 11691, 11692, 11693, 11694, 11696, 11697, 11698, 11699,
		11700, 11701, 11702, 11704, 11705, 11706, 11707, 11708, 11709, 11710, 11712,
		11713, 11714, 11715, 11716, 11717, 11718, 11720, 11721, 11722, 11723, 11724,
		11725, 11726, 11728, 11729, 11730, 11731, 11732, 11733, 11734, 11736, 11737,
		11738, 11739, 11740, 11741, 11742, 43777, 43778, 43779, 43780, 43781, 43782,
		43785, 43786, 43787, 43788, 43789, 43790, 43793, 43794, 43795, 43796, 43797,
		43798, 43808, 43809, 43810, 43811, 43812, 43813, 43814, 43816, 43817, 43818,
		43819, 43820, 43821, 43822, 124896, 124897, 124898, 124899, 124900, 124901,
		124902, 124904, 124905, 124906, 124907, 124909, 124910:
		return true
	}
	return r >= 4608 && (r <= 5017 || r >= 11648 && (r <= 11670 || r >= 124912)) &&
		r <= 124926
}

// HasScxGeor returns true if r has Unicode property Script_Extensions = Georgian.
func HasScxGeor(r rune) bool {
	switch r {
	case 4294, 4296, 4297, 4298, 4299, 4300, 4302, 4303, 7355, 7356, 11558,
		11560, 11561, 11562, 11563, 11564:
		return false
	case 4295, 4301, 7357, 7358, 7359, 11559, 11565:
		return true
	}
	return r >= 4256 && (r <= 4351 || r >= 7312 && (r <= 7359 || r >= 11520)) &&
		r <= 11565
}

// HasScxGlag returns true if r has Unicode property Script_Extensions = Glagolitic.
func HasScxGlag(r rune) bool {
	switch r {
	case 122887, 122905, 122906, 122914, 122917:
		return false
	case 1156, 1159, 11843, 42607, 122880, 122881, 122882, 122883, 122884,
		122885, 122886, 122907, 122908, 122909, 122910, 122911, 122912, 122913,
		122915, 122916, 122918, 122919, 122920, 122921, 122922:
		return true
	}
	return r >= 11264 && (r <= 11359 || r >= 122888) && r <= 122922
}

// HasScxGong returns true if r has Unicode property Script_Extensions = Gunjala_Gondi.
func HasScxGong(r rune) bool {
	switch r {
	case 2404, 2405, 73056, 73057, 73058, 73059, 73060, 73061, 73063, 73064,
		73104, 73105, 73107, 73108, 73109, 73110, 73111, 73112:
		return true
	case 73062, 73065, 73103, 73106, 73113, 73114, 73115, 73116, 73117, 73118,
		73119:
		return false
	}
	return r >= 73066 && r <= 73129
}

// HasScxGonm returns true if r has Unicode property Script_Extensions = Masaram_Gondi.
func HasScxGonm(r rune) bool {
	switch r {
	case 72967, 72970, 73015, 73016, 73017, 73019, 73022, 73032, 73033, 73034,
		73035, 73036, 73037, 73038, 73039:
		return false
	case 2404, 2405, 72960, 72961, 72962, 72963, 72964, 72965, 72966, 72968,
		72969, 73018, 73020, 73021:
		return true
	}
	return r >= 72971 && r <= 73049
}

// HasScxGoth returns true if r has Unicode property Script_Extensions = Gothic.
func HasScxGoth(r rune) bool {
	return r >= 66352 && r <= 66378
}

// HasScxGran returns true if r has Unicode property Script_Extensions = Grantha.
func HasScxGran(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 7376, 7378, 7379, 7410, 7411, 7412, 7416, 7417,
		8432, 70400, 70401, 70402, 70403, 70405, 70406, 70407, 70408, 70409, 70410,
		70411, 70412, 70415, 70416, 70442, 70443, 70444, 70445, 70446, 70447, 70448,
		70450, 70451, 70453, 70454, 70455, 70456, 70457, 70471, 70472, 70475, 70476,
		70477, 70480, 70487, 70493, 70494, 70495, 70496, 70497, 70498, 70499, 70502,
		70503, 70504, 70505, 70506, 70507, 70508, 70512, 70513, 70514, 70515, 70516,
		73680, 73681, 73683:
		return true
	case 7377, 7413, 7414, 7415, 70404, 70413, 70414, 70417, 70418, 70441, 70449,
		70452, 70458, 70469, 70470, 70473, 70474, 70478, 70479, 70481, 70482, 70483,
		70484, 70485, 70486, 70488, 70489, 70490, 70491, 70492, 70500, 70501, 70509,
		70510, 70511, 73682:
		return false
	}
	return r >= 3046 && (r <= 3059 || r >= 70419) && r <= 70516
}

// HasScxGrek returns true if r has Unicode property Script_Extensions = Greek.
func HasScxGrek(r rune) bool {
	switch r {
	case 834, 837, 880, 881, 882, 883, 885, 886, 887, 890, 891, 892, 893, 895,
		900, 902, 904, 905, 906, 908, 7462, 7463, 7464, 7465, 7466, 7517, 7518,
		7519, 7520, 7521, 7526, 7527, 7528, 7529, 7530, 7615, 7616, 7617, 7960,
		7961, 7962, 7963, 7964, 7965, 8008, 8009, 8010, 8011, 8012, 8013, 8016,
		8017, 8018, 8019, 8020, 8021, 8022, 8023, 8025, 8027, 8029, 8150, 8151,
		8152, 8153, 8154, 8155, 8178, 8179, 8180, 8486, 43877, 65952:
		return true
	case 884, 888, 889, 894, 896, 897, 898, 899, 901, 903, 907, 909, 930, 7522,
		7523, 7524, 7525, 7958, 7959, 7966, 7967, 8006, 8007, 8014, 8015, 8024,
		8026, 8028, 8030, 8062, 8063, 8117, 8133, 8148, 8149, 8156, 8176, 8177,
		8181:
		return false
	}
	return r >= 910 && (r <= 1023 && (r <= 993 || r >= 1008) || r >= 7936 &&
		(r <= 8190 || r >= 65856 && (r <= 65934 || r >= 119296))) && r <= 119365
}

// HasScxGujr returns true if r has Unicode property Script_Extensions = Gujarati.
func HasScxGujr(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 2689, 2690, 2691, 2703, 2704, 2705, 2730, 2731,
		2732, 2733, 2734, 2735, 2736, 2738, 2739, 2741, 2742, 2743, 2744, 2745,
		2759, 2760, 2761, 2763, 2764, 2765, 2768, 2784, 2785, 2786, 2787, 2809,
		2810, 2811, 2812, 2813, 2814, 2815:
		return true
	case 2692, 2702, 2706, 2729, 2737, 2740, 2746, 2747, 2758, 2762, 2766, 2767,
		2788, 2789, 2802, 2803, 2804, 2805, 2806, 2807, 2808:
		return false
	}
	return r >= 2693 && (r <= 2768 || r >= 2790 && (r <= 2815 || r >= 43056)) &&
		r <= 43065
}

// HasScxGuru returns true if r has Unicode property Script_Extensions = Gurmukhi.
func HasScxGuru(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 2561, 2562, 2563, 2565, 2566, 2567, 2568, 2569,
		2570, 2575, 2576, 2602, 2603, 2604, 2605, 2606, 2607, 2608, 2610, 2611,
		2613, 2614, 2616, 2617, 2620, 2622, 2623, 2624, 2625, 2626, 2631, 2632,
		2635, 2636, 2637, 2641, 2649, 2650, 2651, 2652, 2654:
		return true
	case 2564, 2571, 2572, 2573, 2574, 2577, 2578, 2601, 2609, 2612, 2615, 2618,
		2619, 2621, 2627, 2628, 2629, 2630, 2633, 2634, 2638, 2639, 2640, 2642,
		2643, 2644, 2645, 2646, 2647, 2648, 2653, 2655, 2656, 2657, 2658, 2659,
		2660, 2661:
		return false
	}
	return r >= 2579 && (r <= 2678 || r >= 43056) && r <= 43065
}

// HasScxHang returns true if r has Unicode property Script_Extensions = Hangul.
func HasScxHang(r rune) bool {
	switch r {
	case 12289, 12290, 12291, 12316, 12317, 12318, 12319, 12334, 12335, 12336,
		12343, 12539, 65093, 65094, 65377, 65378, 65379, 65380, 65381, 65474, 65475,
		65476, 65477, 65478, 65479, 65482, 65483, 65484, 65485, 65486, 65487, 65490,
		65491, 65492, 65493, 65494, 65495, 65498, 65499, 65500:
		return true
	case 12292, 12293, 12294, 12295, 12306, 12337, 12338, 12339, 12340, 12341,
		12342, 55239, 55240, 55241, 55242, 65471, 65472, 65473, 65480, 65481, 65488,
		65489, 65496, 65497:
		return false
	}
	return r >= 4352 && (r <= 12830 && (r <= 12305 && (r <= 4607 || r >= 12296) ||
		r >= 12307 && (r <= 12319 || r >= 12593 && (r <= 12686 || r >= 12800))) ||
		r >= 12896 && (r <= 43388 && (r <= 12926 || r >= 43360) || r >= 44032 &&
			(r <= 55203 || r >= 55216 && (r <= 55291 || r >= 65440)))) && r <= 65500
}

// HasScxHani returns true if r has Unicode property Script_Extensions = Han.
func HasScxHani(r rune) bool {
	switch r {
	case 11930, 12292, 12306, 12320, 12334, 12335, 12337, 12338, 12339, 12340,
		12341, 12342, 13311, 64110, 64111, 177977, 177978, 177979, 177980, 177981,
		177982, 177983, 178206, 178207:
		return false
	case 12289, 12290, 12291, 12293, 12294, 12295, 12316, 12317, 12318, 12319,
		12336, 12343, 12344, 12345, 12346, 12347, 12348, 12349, 12350, 12351, 12539,
		13055, 13179, 13180, 13181, 13182, 13183, 42752, 42753, 42754, 42755, 42756,
		42757, 42758, 42759, 65093, 65094, 65377, 65378, 65379, 65380, 65381, 94178,
		94179, 94192, 94193, 127568, 127569:
		return true
	}
	return r >= 11904 && (r <= 13168 && (r <= 12703 && (r <= 12245 && (r <= 12019 ||
		r >= 12032) || r >= 12296 && (r <= 12351 || r >= 12688)) || r >= 12736 &&
		(r <= 12871 && (r <= 12771 || r >= 12832) || r >= 12928 && (r <= 12976 ||
			r >= 12992 && (r <= 13003 || r >= 13144)))) || r >= 13280 && (r <= 119665 &&
		(r <= 40959 && (r <= 19903 || r >= 19968) || r >= 63744 && (r <= 64217 ||
			r >= 119648)) || r >= 131072 && (r <= 177976 && (r <= 173791 || r >= 173824) ||
		r >= 177984 && (r <= 191456 && (r <= 183969 || r >= 183984) || r >= 194560 &&
			(r <= 195101 || r >= 196608))))) && r <= 201546
}

// HasScxHano returns true if r has Unicode property Script_Extensions = Hanunoo.
func HasScxHano(r rune) bool {
	return r >= 5920 && r <= 5942
}

// HasScxHatr returns true if r has Unicode property Script_Extensions = Hatran.
func HasScxHatr(r rune) bool {
	switch r {
	case 67827, 67830, 67831, 67832, 67833, 67834:
		return false
	case 67828, 67829, 67835, 67836, 67837, 67838, 67839:
		return true
	}
	return r >= 67808 && r <= 67839
}

// HasScxHebr returns true if r has Unicode property Script_Extensions = Hebrew.
func HasScxHebr(r rune) bool {
	switch r {
	case 1480, 1481, 1482, 1483, 1484, 1485, 1486, 1487, 1515, 1516, 1517, 1518,
		64311, 64317, 64319, 64322, 64325:
		return false
	case 1519, 1520, 1521, 1522, 1523, 1524, 64312, 64313, 64314, 64315, 64316,
		64318, 64320, 64321, 64323, 64324:
		return true
	}
	return r >= 1425 && (r <= 1524 || r >= 64285) && r <= 64335
}

// HasScxHira returns true if r has Unicode property Script_Extensions = Hiragana.
func HasScxHira(r rune) bool {
	switch r {
	case 12289, 12290, 12291, 12316, 12317, 12318, 12319, 12336, 12337, 12338,
		12339, 12340, 12341, 12343, 12348, 12349, 12441, 12442, 12443, 12444, 12445,
		12446, 12447, 12448, 12539, 12540, 65093, 65094, 65377, 65378, 65379, 65380,
		65381, 65392, 65438, 65439, 110928, 110929, 110930, 127488:
		return true
	case 12306, 12342, 12344, 12345, 12346, 12347, 12350, 12351, 12352, 12439,
		12440:
		return false
	}
	return r >= 12296 && (r <= 12319 || r >= 12353 && (r <= 12448 || r >= 110593)) &&
		r <= 110879
}

// HasScxHluw returns true if r has Unicode property Script_Extensions = Anatolian_Hieroglyphs.
func HasScxHluw(r rune) bool {
	return r >= 82944 && r <= 83526
}

// HasScxHmng returns true if r has Unicode property Script_Extensions = Pahawh_Hmong.
func HasScxHmng(r rune) bool {
	switch r {
	case 93018, 93026, 93048, 93049, 93050, 93051, 93052:
		return false
	case 93019, 93020, 93021, 93022, 93023, 93024, 93025:
		return true
	}
	return r >= 92928 && (r <= 92997 || r >= 93008) && r <= 93071
}

// HasScxHmnp returns true if r has Unicode property Script_Extensions = Nyiakeng_Puachue_Hmong.
func HasScxHmnp(r rune) bool {
	switch r {
	case 123181, 123182, 123183, 123198, 123199, 123210, 123211, 123212, 123213:
		return false
	case 123214, 123215:
		return true
	}
	return r >= 123136 && r <= 123215
}

// HasScxHung returns true if r has Unicode property Script_Extensions = Old_Hungarian.
func HasScxHung(r rune) bool {
	switch r {
	case 68851, 68852, 68853, 68854, 68855, 68856, 68857:
		return false
	case 68858, 68859, 68860, 68861, 68862, 68863:
		return true
	}
	return r >= 68736 && (r <= 68786 || r >= 68800) && r <= 68863
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
	case 43470, 43482, 43483, 43484, 43485:
		return false
	}
	return r >= 43392 && r <= 43487
}

// HasScxKali returns true if r has Unicode property Script_Extensions = Kayah_Li.
func HasScxKali(r rune) bool {
	return r >= 43264 && r <= 43311
}

// HasScxKana returns true if r has Unicode property Script_Extensions = Katakana.
func HasScxKana(r rune) bool {
	switch r {
	case 12289, 12290, 12291, 12316, 12317, 12318, 12319, 12336, 12337, 12338,
		12339, 12340, 12341, 12343, 12348, 12349, 12441, 12442, 12443, 12444, 12539,
		12540, 12541, 12542, 12543, 65093, 65094, 65438, 65439, 110576, 110577,
		110578, 110579, 110581, 110582, 110583, 110584, 110585, 110586, 110587,
		110589, 110590, 110592, 110880, 110881, 110882, 110948, 110949, 110950,
		110951:
		return true
	case 12306, 12342, 12344, 12345, 12346, 12347, 12445, 12446, 12447, 13055,
		110580, 110588, 110591:
		return false
	}
	return r >= 12296 && (r <= 12538 && (r <= 12319 || r >= 12448) || r >= 12784 &&
		(r <= 12799 || r >= 13008 && (r <= 13143 || r >= 65377))) && r <= 65439
}

// HasScxKhar returns true if r has Unicode property Script_Extensions = Kharoshthi.
func HasScxKhar(r rune) bool {
	switch r {
	case 68096, 68097, 68098, 68099, 68101, 68102, 68108, 68109, 68110, 68111,
		68112, 68113, 68114, 68115, 68117, 68118, 68119, 68152, 68153, 68154:
		return true
	case 68150, 68151, 68155, 68156, 68157, 68158, 68169, 68170, 68171, 68172,
		68173, 68174, 68175:
		return false
	}
	return r >= 68121 && r <= 68184
}

// HasScxKhmr returns true if r has Unicode property Script_Extensions = Khmer.
func HasScxKhmr(r rune) bool {
	switch r {
	case 6110, 6111, 6122, 6123, 6124, 6125, 6126, 6127:
		return false
	}
	return r >= 6016 && (r <= 6137 || r >= 6624) && r <= 6655
}

// HasScxKhoj returns true if r has Unicode property Script_Extensions = Khojki.
func HasScxKhoj(r rune) bool {
	switch r {
	case 70162:
		return false
	}
	return r >= 2790 && (r <= 2799 || r >= 43056 && (r <= 43065 || r >= 70144)) &&
		r <= 70206
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
		3297, 3298, 3299, 3313, 3314, 7376, 7378, 7386, 7410, 7412, 43056, 43057,
		43058, 43059, 43060, 43061:
		return true
	case 3213, 3217, 3241, 3252, 3258, 3259, 3269, 3273, 3278, 3279, 3280, 3281,
		3282, 3283, 3284, 3287, 3288, 3289, 3290, 3291, 3292, 3295, 3300, 3301,
		3312, 7377, 7379, 7380, 7381, 7382, 7383, 7384, 7385, 7411:
		return false
	}
	return r >= 3200 && r <= 3314
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
	switch r {
	case 6751, 6781, 6782, 6794, 6795, 6796, 6797, 6798, 6799, 6810, 6811, 6812,
		6813, 6814, 6815:
		return false
	}
	return r >= 6688 && r <= 6829
}

// HasScxLaoo returns true if r has Unicode property Script_Extensions = Lao.
func HasScxLaoo(r rune) bool {
	switch r {
	case 3713, 3714, 3716, 3718, 3719, 3720, 3721, 3722, 3749, 3776, 3777, 3778,
		3779, 3780, 3782, 3784, 3785, 3786, 3787, 3788, 3789, 3804, 3805, 3806,
		3807:
		return true
	case 3748, 3750, 3774, 3775, 3781, 3783, 3790, 3791, 3802, 3803:
		return false
	}
	return r >= 3724 && r <= 3807
}

// HasScxLatn returns true if r has Unicode property Script_Extensions = Latin.
func HasScxLatn(r rune) bool {
	switch r {
	case 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82,
		83, 84, 85, 86, 87, 88, 89, 90, 97, 98, 99, 100, 101, 102, 103, 104, 105,
		106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
		121, 122, 170, 186, 736, 737, 738, 739, 740, 1157, 1158, 2385, 2386, 4347,
		7522, 7523, 7524, 7525, 8239, 8305, 8319, 8432, 8490, 8491, 8498, 8526,
		42752, 42753, 42754, 42755, 42756, 42757, 42758, 42759, 42960, 42961, 42963,
		42965, 42966, 42967, 42968, 42969, 43310, 43878, 43879, 43880, 43881, 64256,
		64257, 64258, 64259, 64260, 64261, 64262, 67456, 67457, 67458, 67459, 67460,
		67461:
		return true
	case 187, 188, 189, 190, 191, 215, 247, 7462, 7463, 7464, 7465, 7466, 7467,
		7517, 7518, 7519, 7520, 7521, 7526, 7527, 7528, 7529, 7530, 7544, 8492,
		8493, 8494, 8495, 8496, 8497, 42888, 42889, 42890, 42955, 42956, 42957,
		42958, 42959, 42962, 42964, 43867, 43877, 65339, 65340, 65341, 65342, 65343,
		65344, 67462, 67505:
		return false
	}
	return r >= 192 && (r <= 8584 && (r <= 7461 && (r <= 696 || r >= 867 && (r <= 879 ||
		r >= 7424)) || r >= 7468 && (r <= 7935 && (r <= 7614 || r >= 7680) || r >= 8336 &&
		(r <= 8348 || r >= 8544))) || r >= 11360 && (r <= 43007 && (r <= 11391 ||
		r >= 42786 && (r <= 42969 || r >= 42994)) || r >= 43824 && (r <= 65338 &&
		(r <= 43881 || r >= 65313) || r >= 65345 && (r <= 65370 || r >= 67463 &&
		(r <= 67514 || r >= 122624))))) && r <= 122654
}

// HasScxLepc returns true if r has Unicode property Script_Extensions = Lepcha.
func HasScxLepc(r rune) bool {
	switch r {
	case 7245, 7246, 7247:
		return true
	case 7224, 7225, 7226, 7242, 7243, 7244:
		return false
	}
	return r >= 7168 && r <= 7247
}

// HasScxLimb returns true if r has Unicode property Script_Extensions = Limbu.
func HasScxLimb(r rune) bool {
	switch r {
	case 2405, 6464:
		return true
	case 6431, 6444, 6445, 6446, 6447, 6460, 6461, 6462, 6463, 6465, 6466, 6467:
		return false
	}
	return r >= 6400 && r <= 6479
}

// HasScxLina returns true if r has Unicode property Script_Extensions = Linear_A.
func HasScxLina(r rune) bool {
	switch r {
	case 67424, 67425, 67426, 67427, 67428, 67429, 67430, 67431:
		return true
	}
	return r >= 65799 && (r <= 65843 || r >= 67072 && (r <= 67382 || r >= 67392)) &&
		r <= 67413
}

// HasScxLinb returns true if r has Unicode property Script_Extensions = Linear_B.
func HasScxLinb(r rune) bool {
	switch r {
	case 65548, 65575, 65595, 65598, 65614, 65615, 65787, 65788, 65789, 65790,
		65791, 65795, 65796, 65797, 65798, 65844, 65845, 65846:
		return false
	case 65596, 65597, 65792, 65793, 65794:
		return true
	}
	return r >= 65536 && (r <= 65629 || r >= 65664) && r <= 65855
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
	case 67898, 67899, 67900, 67901, 67902:
		return false
	case 67903:
		return true
	}
	return r >= 67872 && r <= 67903
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
	case 2140, 2141:
		return false
	}
	return r >= 2112 && r <= 2142
}

// HasScxMani returns true if r has Unicode property Script_Extensions = Manichaean.
func HasScxMani(r rune) bool {
	switch r {
	case 1600:
		return true
	case 68327, 68328, 68329, 68330:
		return false
	}
	return r >= 68288 && r <= 68342
}

// HasScxMarc returns true if r has Unicode property Script_Extensions = Marchen.
func HasScxMarc(r rune) bool {
	switch r {
	case 72848, 72849, 72872:
		return false
	}
	return r >= 72816 && r <= 72886
}

// HasScxMedf returns true if r has Unicode property Script_Extensions = Medefaidrin.
func HasScxMedf(r rune) bool {
	return r >= 93760 && r <= 93850
}

// HasScxMend returns true if r has Unicode property Script_Extensions = Mende_Kikakui.
func HasScxMend(r rune) bool {
	switch r {
	case 125125, 125126:
		return false
	}
	return r >= 124928 && r <= 125142
}

// HasScxMerc returns true if r has Unicode property Script_Extensions = Meroitic_Cursive.
func HasScxMerc(r rune) bool {
	switch r {
	case 68024, 68025, 68026, 68027, 68048, 68049:
		return false
	}
	return r >= 68000 && r <= 68095
}

// HasScxMero returns true if r has Unicode property Script_Extensions = Meroitic_Hieroglyphs.
func HasScxMero(r rune) bool {
	return r >= 67968 && r <= 67999
}

// HasScxMlym returns true if r has Unicode property Script_Extensions = Malayalam.
func HasScxMlym(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 3342, 3343, 3344, 3398, 3399, 3400, 3402, 3403,
		3404, 3405, 3406, 3407, 7386, 43056, 43057, 43058:
		return true
	case 3341, 3345, 3397, 3401, 3408, 3409, 3410, 3411, 3428, 3429:
		return false
	}
	return r >= 3328 && r <= 3455
}

// HasScxModi returns true if r has Unicode property Script_Extensions = Modi.
func HasScxModi(r rune) bool {
	return r >= 43056 && (r <= 43065 || r >= 71168 && (r <= 71236 || r >= 71248)) &&
		r <= 71257
}

// HasScxMong returns true if r has Unicode property Script_Extensions = Mongolian.
func HasScxMong(r rune) bool {
	switch r {
	case 6170, 6171, 6172, 6173, 6174, 6175, 6265, 6266, 6267, 6268, 6269, 6270,
		6271:
		return false
	case 8239:
		return true
	}
	return r >= 6144 && (r <= 6314 || r >= 71264) && r <= 71276
}

// HasScxMroo returns true if r has Unicode property Script_Extensions = Mro.
func HasScxMroo(r rune) bool {
	switch r {
	case 92767, 92778, 92779, 92780, 92781:
		return false
	case 92782, 92783:
		return true
	}
	return r >= 92736 && r <= 92783
}

// HasScxMtei returns true if r has Unicode property Script_Extensions = Meetei_Mayek.
func HasScxMtei(r rune) bool {
	switch r {
	case 44014, 44015:
		return false
	}
	return r >= 43744 && (r <= 43766 || r >= 43968) && r <= 44025
}

// HasScxMult returns true if r has Unicode property Script_Extensions = Multani.
func HasScxMult(r rune) bool {
	switch r {
	case 70272, 70273, 70274, 70275, 70276, 70277, 70278, 70280, 70282, 70283,
		70284, 70285:
		return true
	case 70279, 70281, 70286, 70302:
		return false
	}
	return r >= 2662 && (r <= 2671 || r >= 70287) && r <= 70313
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
	case 2404, 2405, 7401, 7410, 7418, 43056, 43057, 43058, 43059, 43060, 43061,
		72096, 72097, 72098, 72099, 72100, 72101, 72102, 72103:
		return true
	case 7402, 7403, 7404, 7405, 7406, 7407, 7408, 7409, 7411, 7412, 7413, 7414,
		7415, 7416, 7417, 72104, 72105, 72152, 72153:
		return false
	}
	return r >= 3302 && (r <= 3311 || r >= 72106) && r <= 72164
}

// HasScxNarb returns true if r has Unicode property Script_Extensions = Old_North_Arabian.
func HasScxNarb(r rune) bool {
	return r >= 68224 && r <= 68255
}

// HasScxNbat returns true if r has Unicode property Script_Extensions = Nabataean.
func HasScxNbat(r rune) bool {
	switch r {
	case 67743, 67744, 67745, 67746, 67747, 67748, 67749, 67750:
		return false
	}
	return r >= 67712 && r <= 67759
}

// HasScxNewa returns true if r has Unicode property Script_Extensions = Newa.
func HasScxNewa(r rune) bool {
	switch r {
	case 70748:
		return false
	case 70749, 70750, 70751, 70752, 70753:
		return true
	}
	return r >= 70656 && r <= 70753
}

// HasScxNkoo returns true if r has Unicode property Script_Extensions = Nko.
func HasScxNkoo(r rune) bool {
	switch r {
	case 1548, 1563, 1567, 2045, 2046, 2047, 64830, 64831:
		return true
	case 1564, 1565, 1566, 2043, 2044:
		return false
	}
	return r >= 1984 && r <= 2047
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
	case 2385, 2386, 2404, 2405, 2817, 2818, 2819, 2821, 2822, 2823, 2824, 2825,
		2826, 2827, 2828, 2831, 2832, 2858, 2859, 2860, 2861, 2862, 2863, 2864,
		2866, 2867, 2869, 2870, 2871, 2872, 2873, 2887, 2888, 2891, 2892, 2893,
		2901, 2902, 2903, 2908, 2909, 2911, 2912, 2913, 2914, 2915, 7386, 7410:
		return true
	case 2820, 2829, 2830, 2833, 2834, 2857, 2865, 2868, 2874, 2875, 2885, 2886,
		2889, 2890, 2894, 2895, 2896, 2897, 2898, 2899, 2900, 2904, 2905, 2906,
		2907, 2910, 2916, 2917:
		return false
	}
	return r >= 2835 && r <= 2935
}

// HasScxOsge returns true if r has Unicode property Script_Extensions = Osage.
func HasScxOsge(r rune) bool {
	switch r {
	case 66772, 66773, 66774, 66775:
		return false
	}
	return r >= 66736 && r <= 66811
}

// HasScxOsma returns true if r has Unicode property Script_Extensions = Osmanya.
func HasScxOsma(r rune) bool {
	switch r {
	case 66718, 66719:
		return false
	}
	return r >= 66688 && r <= 66729
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
	switch r {
	case 68467, 68468, 68469, 68470, 68471:
		return false
	case 68472, 68473, 68474, 68475, 68476, 68477, 68478, 68479:
		return true
	}
	return r >= 68448 && r <= 68479
}

// HasScxPhlp returns true if r has Unicode property Script_Extensions = Psalter_Pahlavi.
func HasScxPhlp(r rune) bool {
	switch r {
	case 1600, 68505, 68506, 68507, 68508, 68521, 68522, 68523, 68524, 68525,
		68526, 68527:
		return true
	case 68498, 68499, 68500, 68501, 68502, 68503, 68504:
		return false
	}
	return r >= 68480 && r <= 68508
}

// HasScxPhnx returns true if r has Unicode property Script_Extensions = Phoenician.
func HasScxPhnx(r rune) bool {
	switch r {
	case 67868, 67869, 67870:
		return false
	case 67871:
		return true
	}
	return r >= 67840 && r <= 67871
}

// HasScxPlrd returns true if r has Unicode property Script_Extensions = Miao.
func HasScxPlrd(r rune) bool {
	switch r {
	case 94027, 94028, 94029, 94030, 94088, 94089, 94090, 94091, 94092, 94093,
		94094:
		return false
	}
	return r >= 93952 && r <= 94111
}

// HasScxPrti returns true if r has Unicode property Script_Extensions = Inscriptional_Parthian.
func HasScxPrti(r rune) bool {
	switch r {
	case 68438, 68439:
		return false
	case 68440, 68441, 68442, 68443, 68444, 68445, 68446, 68447:
		return true
	}
	return r >= 68416 && r <= 68447
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
	case 1564, 1565, 1566, 68904, 68905, 68906, 68907, 68908, 68909, 68910,
		68911:
		return false
	}
	return r >= 68864 && r <= 68921
}

// HasScxRunr returns true if r has Unicode property Script_Extensions = Runic.
func HasScxRunr(r rune) bool {
	switch r {
	case 5867, 5868, 5869:
		return false
	}
	return r >= 5792 && r <= 5880
}

// HasScxSamr returns true if r has Unicode property Script_Extensions = Samaritan.
func HasScxSamr(r rune) bool {
	switch r {
	case 2094, 2095:
		return false
	}
	return r >= 2048 && r <= 2110
}

// HasScxSarb returns true if r has Unicode property Script_Extensions = Old_South_Arabian.
func HasScxSarb(r rune) bool {
	return r >= 68192 && r <= 68223
}

// HasScxSaur returns true if r has Unicode property Script_Extensions = Saurashtra.
func HasScxSaur(r rune) bool {
	switch r {
	case 43206, 43207, 43208, 43209, 43210, 43211, 43212, 43213:
		return false
	}
	return r >= 43136 && r <= 43225
}

// HasScxSgnw returns true if r has Unicode property Script_Extensions = SignWriting.
func HasScxSgnw(r rune) bool {
	switch r {
	case 121499, 121500, 121501, 121502, 121503:
		return true
	case 121504:
		return false
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
	case 7384, 7386, 7387, 7390, 7391:
		return false
	}
	return r >= 70016 && r <= 70111
}

// HasScxSidd returns true if r has Unicode property Script_Extensions = Siddham.
func HasScxSidd(r rune) bool {
	switch r {
	case 71094, 71095:
		return false
	}
	return r >= 71040 && r <= 71133
}

// HasScxSind returns true if r has Unicode property Script_Extensions = Khudawadi.
func HasScxSind(r rune) bool {
	switch r {
	case 2404, 2405:
		return true
	case 70379, 70380, 70381, 70382, 70383:
		return false
	}
	return r >= 43056 && (r <= 43065 || r >= 70320) && r <= 70393
}

// HasScxSinh returns true if r has Unicode property Script_Extensions = Sinhala.
func HasScxSinh(r rune) bool {
	switch r {
	case 2404, 2405, 3457, 3458, 3459, 3517, 3520, 3521, 3522, 3523, 3524, 3525,
		3526, 3530, 3535, 3536, 3537, 3538, 3539, 3540, 3542, 3544, 3545, 3546,
		3547, 3548, 3549, 3550, 3551, 3570, 3571, 3572:
		return true
	case 3460, 3479, 3480, 3481, 3506, 3516, 3518, 3519, 3527, 3528, 3529, 3531,
		3532, 3533, 3534, 3541, 3543, 3552, 3553, 3554, 3555, 3556, 3557, 3568,
		3569:
		return false
	}
	return r >= 3461 && (r <= 3572 || r >= 70113) && r <= 70132
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
	switch r {
	case 69865, 69866, 69867, 69868, 69869, 69870, 69871:
		return false
	}
	return r >= 69840 && r <= 69881
}

// HasScxSoyo returns true if r has Unicode property Script_Extensions = Soyombo.
func HasScxSoyo(r rune) bool {
	return r >= 72272 && r <= 72354
}

// HasScxSund returns true if r has Unicode property Script_Extensions = Sundanese.
func HasScxSund(r rune) bool {
	switch r {
	case 7360, 7361, 7362, 7363, 7364, 7365, 7366, 7367:
		return true
	}
	return r >= 7040 && r <= 7103
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
	case 1565, 1566, 1806, 1867, 1868, 7673:
		return false
	}
	return r >= 1611 && (r <= 1621 || r >= 1792 && (r <= 1871 || r >= 2144)) && r <= 2154
}

// HasScxTagb returns true if r has Unicode property Script_Extensions = Tagbanwa.
func HasScxTagb(r rune) bool {
	switch r {
	case 5997, 6001:
		return false
	case 5941, 5942, 5998, 5999, 6000, 6002, 6003:
		return true
	}
	return r >= 5984 && r <= 6003
}

// HasScxTakr returns true if r has Unicode property Script_Extensions = Takri.
func HasScxTakr(r rune) bool {
	switch r {
	case 71354, 71355, 71356, 71357, 71358, 71359:
		return false
	case 2404, 2405:
		return true
	}
	return r >= 43056 && (r <= 43065 || r >= 71296) && r <= 71369
}

// HasScxTale returns true if r has Unicode property Script_Extensions = Tai_Le.
func HasScxTale(r rune) bool {
	switch r {
	case 6510, 6511:
		return false
	case 6512, 6513, 6514, 6515, 6516:
		return true
	}
	return r >= 4160 && (r <= 4169 || r >= 6480) && r <= 6516
}

// HasScxTalu returns true if r has Unicode property Script_Extensions = New_Tai_Lue.
func HasScxTalu(r rune) bool {
	switch r {
	case 6572, 6573, 6574, 6575, 6602, 6603, 6604, 6605, 6606, 6607, 6619, 6620,
		6621:
		return false
	case 6622, 6623:
		return true
	}
	return r >= 6528 && r <= 6623
}

// HasScxTaml returns true if r has Unicode property Script_Extensions = Tamil.
func HasScxTaml(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 2946, 2947, 2949, 2950, 2951, 2952, 2953, 2954,
		2958, 2959, 2960, 2962, 2963, 2964, 2965, 2969, 2970, 2972, 2974, 2975,
		2979, 2980, 2984, 2985, 2986, 3006, 3007, 3008, 3009, 3010, 3014, 3015,
		3016, 3018, 3019, 3020, 3021, 3024, 3031, 7386, 43251, 70401, 70403, 70459,
		70460, 73727:
		return true
	case 2948, 2955, 2956, 2957, 2961, 2966, 2967, 2968, 2971, 2973, 2976, 2977,
		2978, 2981, 2982, 2983, 2987, 2988, 2989, 3002, 3003, 3004, 3005, 3011,
		3012, 3013, 3017, 3022, 3023, 3025, 3026, 3027, 3028, 3029, 3030, 70402:
		return false
	}
	return r >= 2990 && (r <= 3031 || r >= 3046 && (r <= 3066 || r >= 73664)) &&
		r <= 73713
}

// HasScxTang returns true if r has Unicode property Script_Extensions = Tangut.
func HasScxTang(r rune) bool {
	switch r {
	case 94176:
		return true
	case 100344, 100345, 100346, 100347, 100348, 100349, 100350, 100351:
		return false
	}
	return r >= 94208 && (r <= 101119 || r >= 101632) && r <= 101640
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
	case 3085, 3089, 3113, 3130, 3131, 3141, 3145, 3150, 3151, 3152, 3153, 3154,
		3155, 3156, 3159, 3163, 3164, 3166, 3167, 3172, 3173, 3184, 3185, 3186,
		3187, 3188, 3189, 3190:
		return false
	}
	return r >= 3072 && r <= 3199
}

// HasScxTfng returns true if r has Unicode property Script_Extensions = Tifinagh.
func HasScxTfng(r rune) bool {
	switch r {
	case 11624, 11625, 11626, 11627, 11628, 11629, 11630:
		return false
	case 11631, 11632, 11647:
		return true
	}
	return r >= 11568 && r <= 11632
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
	case 1565, 1566:
		return false
	case 1548, 1563, 1564, 1567, 65010, 65021:
		return true
	}
	return r >= 1632 && (r <= 1641 || r >= 1920) && r <= 1969
}

// HasScxThai returns true if r has Unicode property Script_Extensions = Thai.
func HasScxThai(r rune) bool {
	switch r {
	case 3643, 3644, 3645, 3646, 3647:
		return false
	}
	return r >= 3585 && r <= 3675
}

// HasScxTibt returns true if r has Unicode property Script_Extensions = Tibetan.
func HasScxTibt(r rune) bool {
	switch r {
	case 3912, 3949, 3950, 3951, 3952, 3992, 4029, 4045, 4053, 4054, 4055, 4056:
		return false
	case 4046, 4047, 4048, 4049, 4050, 4051, 4052, 4057, 4058:
		return true
	}
	return r >= 3840 && r <= 4058
}

// HasScxTirh returns true if r has Unicode property Script_Extensions = Tirhuta.
func HasScxTirh(r rune) bool {
	switch r {
	case 2385, 2386, 2404, 2405, 7410:
		return true
	case 70856, 70857, 70858, 70859, 70860, 70861, 70862, 70863:
		return false
	}
	return r >= 43056 && (r <= 43065 || r >= 70784) && r <= 70873
}

// HasScxTnsa returns true if r has Unicode property Script_Extensions = Tangsa.
func HasScxTnsa(r rune) bool {
	switch r {
	case 92863:
		return false
	}
	return r >= 92784 && r <= 92873
}

// HasScxToto returns true if r has Unicode property Script_Extensions = Toto.
func HasScxToto(r rune) bool {
	return r >= 123536 && r <= 123566
}

// HasScxUgar returns true if r has Unicode property Script_Extensions = Ugaritic.
func HasScxUgar(r rune) bool {
	switch r {
	case 66462:
		return false
	case 66463:
		return true
	}
	return r >= 66432 && r <= 66463
}

// HasScxVaii returns true if r has Unicode property Script_Extensions = Vai.
func HasScxVaii(r rune) bool {
	return r >= 42240 && r <= 42539
}

// HasScxVith returns true if r has Unicode property Script_Extensions = Vithkuqi.
func HasScxVith(r rune) bool {
	switch r {
	case 66939, 66955, 66963, 66966, 66978, 66994, 67002:
		return false
	case 66956, 66957, 66958, 66959, 66960, 66961, 66962, 66964, 66965, 66995,
		66996, 66997, 66998, 66999, 67000, 67001, 67003, 67004:
		return true
	}
	return r >= 66928 && r <= 67004
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
	case 123642, 123643, 123644, 123645, 123646:
		return false
	case 123647:
		return true
	}
	return r >= 123584 && r <= 123647
}

// HasScxXpeo returns true if r has Unicode property Script_Extensions = Old_Persian.
func HasScxXpeo(r rune) bool {
	switch r {
	case 66500, 66501, 66502, 66503:
		return false
	}
	return r >= 66464 && r <= 66517
}

// HasScxXsux returns true if r has Unicode property Script_Extensions = Cuneiform.
func HasScxXsux(r rune) bool {
	switch r {
	case 74863:
		return false
	case 74864, 74865, 74866, 74867, 74868:
		return true
	}
	return r >= 73728 && (r <= 74649 || r >= 74752 && (r <= 74868 || r >= 74880)) &&
		r <= 75075
}

// HasScxYezi returns true if r has Unicode property Script_Extensions = Yezidi.
func HasScxYezi(r rune) bool {
	switch r {
	case 1548, 1563, 1567, 69291, 69292, 69293, 69296, 69297:
		return true
	case 1564, 1565, 1566, 69290, 69294, 69295:
		return false
	}
	return r >= 1632 && (r <= 1641 || r >= 69248) && r <= 69297
}

// HasScxYiii returns true if r has Unicode property Script_Extensions = Yi.
func HasScxYiii(r rune) bool {
	switch r {
	case 12289, 12290, 12308, 12309, 12310, 12311, 12312, 12313, 12314, 12315,
		12539, 65377, 65378, 65379, 65380, 65381:
		return true
	case 12306, 12307, 42125, 42126, 42127:
		return false
	}
	return r >= 12296 && (r <= 12315 || r >= 40960) && r <= 42182
}

// HasScxZanb returns true if r has Unicode property Script_Extensions = Zanabazar_Square.
func HasScxZanb(r rune) bool {
	return r >= 72192 && r <= 72263
}

// HasScxZinh returns true if r has Unicode property Script_Extensions = Inherited.
func HasScxZinh(r rune) bool {
	switch r {
	case 834, 837, 7672, 7674, 118574, 118575, 119171, 119172:
		return false
	case 835, 836, 2387, 2388, 7673, 7675, 7676, 7677, 7678, 7679, 8204, 8205,
		66045, 119143, 119144, 119145, 119163, 119164, 119165, 119166, 119167,
		119168, 119169, 119170, 119173, 119174, 119175, 119176, 119177, 119178,
		119179, 119210, 119211, 119212, 119213:
		return true
	}
	return r >= 768 && (r <= 8431 && (r <= 6862 && (r <= 866 || r >= 6832) ||
		r >= 7618 && (r <= 7679 || r >= 8400)) || r >= 65024 && (r <= 65069 &&
		(r <= 65039 || r >= 65056) || r >= 118528 && (r <= 118598 || r >= 917760))) &&
		r <= 917999
}

// HasScxZyyy returns true if r has Unicode property Script_Extensions = Common.
func HasScxZyyy(r rune) bool {
	switch r {
	case 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82,
		83, 84, 85, 86, 87, 88, 89, 90, 97, 98, 99, 100, 101, 102, 103, 104, 105,
		106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
		121, 122, 170, 186, 736, 737, 738, 739, 740, 746, 747, 895, 896, 897, 898,
		899, 900, 902, 8204, 8205, 8239, 8293, 8305, 8306, 8307, 8319, 8486, 8490,
		8491, 8498, 8526, 8588, 8589, 8590, 8591, 11124, 11125, 11158, 11843, 12284,
		12285, 12286, 12287, 12289, 12290, 12291, 13179, 13180, 13181, 13182, 13183,
		65093, 65094, 65107, 65127, 65280, 65511, 119079, 119080, 119143, 119144,
		119145, 119163, 119164, 119165, 119166, 119167, 119168, 119169, 119170,
		119173, 119174, 119175, 119176, 119177, 119178, 119179, 119210, 119211,
		119212, 119213, 119893, 119965, 119968, 119969, 119971, 119972, 119975,
		119976, 119981, 119994, 119996, 120004, 120070, 120075, 120076, 120085,
		120093, 120122, 120127, 120133, 120135, 120136, 120137, 120145, 120486,
		120487, 120780, 120781, 127020, 127021, 127022, 127023, 127151, 127152,
		127168, 127184, 127488, 127548, 127549, 127550, 127551, 128728, 128729,
		128730, 128731, 128732, 128749, 128750, 128751, 128765, 128766, 128767,
		128985, 128986, 128987, 128988, 128989, 128990, 128991, 129004, 129005,
		129006, 129007, 129036, 129037, 129038, 129039, 129096, 129097, 129098,
		129099, 129100, 129101, 129102, 129103, 129114, 129115, 129116, 129117,
		129118, 129119, 129160, 129161, 129162, 129163, 129164, 129165, 129166,
		129167, 129198, 129199, 129646, 129647, 129653, 129654, 129655, 129661,
		129662, 129663, 129709, 129710, 129711, 129723, 129724, 129725, 129726,
		129727, 129754, 129755, 129756, 129757, 129758, 129759, 129768, 129769,
		129770, 129771, 129772, 129773, 129774, 129775, 129939:
		return false
	case 91, 92, 93, 94, 95, 96, 187, 188, 189, 190, 191, 215, 247, 741, 742,
		743, 744, 745, 884, 894, 901, 903, 1541, 1757, 2274, 3647, 4053, 4054, 4055,
		4056, 5867, 5868, 5869, 8487, 8488, 8489, 8492, 8493, 8494, 8495, 8496,
		8497, 8585, 8586, 8587, 12288, 12292, 12306, 12320, 12342, 12927, 13004,
		13005, 13006, 13007, 13311, 42888, 42889, 42890, 43867, 43882, 43883, 65128,
		65129, 65130, 65131, 65279, 65339, 65340, 65341, 65342, 65343, 65344, 65371,
		65372, 65373, 65374, 65375, 65376, 65504, 65505, 65506, 65507, 65508, 65509,
		65510, 65512, 65513, 65514, 65515, 65516, 65517, 65518, 65529, 65530, 65531,
		65532, 65533, 119171, 119172, 119666, 119667, 119668, 119669, 119670,
		119671, 119672, 119966, 119967, 119970, 119973, 119974, 119977, 119978,
		119979, 119980, 119995, 119997, 119998, 119999, 120000, 120001, 120002,
		120003, 120071, 120072, 120073, 120074, 120077, 120078, 120079, 120080,
		120081, 120082, 120083, 120084, 120086, 120087, 120088, 120089, 120090,
		120091, 120092, 120123, 120124, 120125, 120126, 120128, 120129, 120130,
		120131, 120132, 120134, 120138, 120139, 120140, 120141, 120142, 120143,
		120144, 127489, 127490, 127584, 127585, 127586, 127587, 127588, 127589,
		129008, 129200, 129201, 129648, 129649, 129650, 129651, 129652, 129656,
		129657, 129658, 129659, 129660, 129664, 129665, 129666, 129667, 129668,
		129669, 129670, 129728, 129729, 129730, 129731, 129732, 129733, 129760,
		129761, 129762, 129763, 129764, 129765, 129766, 129767, 129776, 129777,
		129778, 129779, 129780, 129781, 129782, 917505:
		return true
	}
	return r >= 0 && (r <= 118723 && (r <= 12283 && (r <= 8485 && (r <= 735 &&
		(r <= 191 || r >= 697) || r >= 748 && (r <= 8203 && (r <= 767 || r >= 8192) ||
		r >= 8206 && (r <= 8334 || r >= 8352 && (r <= 8384 || r >= 8448)))) || r >= 8499 &&
		(r <= 9290 && (r <= 8543 || r >= 8592 && (r <= 9254 || r >= 9280)) || r >= 9312 &&
			(r <= 11123 && (r <= 10239 || r >= 10496) || r >= 11126 && (r <= 11263 ||
				r >= 11776 && (r <= 11869 || r >= 12272))))) || r >= 12872 && (r <= 42785 &&
		(r <= 12991 && (r <= 12895 || r >= 12977) || r >= 13169 && (r <= 13279 ||
			r >= 19904 && (r <= 19967 || r >= 42760))) || r >= 65040 && (r <= 65312 &&
		(r <= 65049 || r >= 65072 && (r <= 65131 || r >= 65281)) || r >= 65936 &&
		(r <= 65948 || r >= 66000 && (r <= 66044 || r >= 118608))))) || r >= 118784 &&
		(r <= 127487 && (r <= 119892 && (r <= 119078 && (r <= 119029 || r >= 119040) ||
			r >= 119081 && (r <= 119539 && (r <= 119274 || r >= 119520) || r >= 119552 &&
				(r <= 119638 || r >= 119808))) || r >= 119894 && (r <= 126269 && (r <= 120831 ||
			r >= 126065 && (r <= 126132 || r >= 126209)) || r >= 126976 && (r <= 127150 &&
			(r <= 127123 || r >= 127136) || r >= 127153 && (r <= 127221 || r >= 127232 &&
			(r <= 127405 || r >= 127462))))) || r >= 127504 && (r <= 129619 && (r <= 128727 &&
			(r <= 127560 || r >= 127744) || r >= 128733 && (r <= 128984 && (r <= 128883 ||
			r >= 128896) || r >= 128992 && (r <= 129008 || r >= 129024 && (r <= 129201 ||
			r >= 129280)))) || r >= 129632 && (r <= 129753 && (r <= 129670 || r >= 129680 &&
			(r <= 129733 || r >= 129744)) || r >= 129792 && (r <= 129994 || r >= 130032 &&
			(r <= 130041 || r >= 917536)))))) && r <= 917631
}

// HasScxZzzz returns true if r has Unicode property Script_Extensions = Unknown.
func HasScxZzzz(r rune) bool {
	switch r {
	case 888, 889, 896, 897, 898, 899, 907, 909, 930, 1328, 1367, 1368, 1419,
		1420, 1424, 1480, 1481, 1482, 1483, 1484, 1485, 1486, 1487, 1515, 1516,
		1517, 1518, 1806, 1867, 1868, 2043, 2044, 2094, 2095, 2111, 2140, 2141,
		2143, 2155, 2156, 2157, 2158, 2159, 2191, 2194, 2195, 2196, 2197, 2198,
		2199, 2436, 2445, 2446, 2449, 2450, 2473, 2481, 2483, 2484, 2485, 2490,
		2491, 2501, 2502, 2505, 2506, 2511, 2512, 2513, 2514, 2515, 2516, 2517,
		2518, 2520, 2521, 2522, 2523, 2526, 2532, 2533, 2559, 2560, 2564, 2571,
		2572, 2573, 2574, 2577, 2578, 2601, 2609, 2612, 2615, 2618, 2619, 2621,
		2627, 2628, 2629, 2630, 2633, 2634, 2638, 2639, 2640, 2642, 2643, 2644,
		2645, 2646, 2647, 2648, 2653, 2655, 2656, 2657, 2658, 2659, 2660, 2661,
		2692, 2702, 2706, 2729, 2737, 2740, 2746, 2747, 2758, 2762, 2766, 2767,
		2788, 2789, 2802, 2803, 2804, 2805, 2806, 2807, 2808, 2816, 2820, 2829,
		2830, 2833, 2834, 2857, 2865, 2868, 2874, 2875, 2885, 2886, 2889, 2890,
		2894, 2895, 2896, 2897, 2898, 2899, 2900, 2904, 2905, 2906, 2907, 2910,
		2916, 2917, 2948, 2955, 2956, 2957, 2961, 2966, 2967, 2968, 2971, 2973,
		2976, 2977, 2978, 2981, 2982, 2983, 2987, 2988, 2989, 3002, 3003, 3004,
		3005, 3011, 3012, 3013, 3017, 3022, 3023, 3025, 3026, 3027, 3028, 3029,
		3030, 3067, 3068, 3069, 3070, 3071, 3085, 3089, 3113, 3130, 3131, 3141,
		3145, 3150, 3151, 3152, 3153, 3154, 3155, 3156, 3159, 3163, 3164, 3166,
		3167, 3172, 3173, 3184, 3185, 3186, 3187, 3188, 3189, 3190, 3213, 3217,
		3241, 3252, 3258, 3259, 3269, 3273, 3278, 3279, 3280, 3281, 3282, 3283,
		3284, 3287, 3288, 3289, 3290, 3291, 3292, 3295, 3300, 3301, 3312, 3341,
		3345, 3397, 3401, 3408, 3409, 3410, 3411, 3428, 3429, 3456, 3460, 3479,
		3480, 3481, 3506, 3516, 3518, 3519, 3527, 3528, 3529, 3531, 3532, 3533,
		3534, 3541, 3543, 3552, 3553, 3554, 3555, 3556, 3557, 3568, 3569, 3643,
		3644, 3645, 3646, 3715, 3717, 3723, 3748, 3750, 3774, 3775, 3781, 3783,
		3790, 3791, 3802, 3803, 3912, 3949, 3950, 3951, 3952, 3992, 4029, 4045,
		4294, 4296, 4297, 4298, 4299, 4300, 4302, 4303, 4681, 4686, 4687, 4695,
		4697, 4702, 4703, 4745, 4750, 4751, 4785, 4790, 4791, 4799, 4801, 4806,
		4807, 4823, 4881, 4886, 4887, 4955, 4956, 4989, 4990, 4991, 5018, 5019,
		5020, 5021, 5022, 5023, 5110, 5111, 5118, 5119, 5789, 5790, 5791, 5881,
		5882, 5883, 5884, 5885, 5886, 5887, 5997, 6001, 6110, 6111, 6122, 6123,
		6124, 6125, 6126, 6127, 6138, 6139, 6140, 6141, 6142, 6143, 6170, 6171,
		6172, 6173, 6174, 6175, 6265, 6266, 6267, 6268, 6269, 6270, 6271, 6315,
		6316, 6317, 6318, 6319, 6431, 6444, 6445, 6446, 6447, 6460, 6461, 6462,
		6463, 6465, 6466, 6467, 6510, 6511, 6572, 6573, 6574, 6575, 6602, 6603,
		6604, 6605, 6606, 6607, 6619, 6620, 6621, 6684, 6685, 6751, 6781, 6782,
		6794, 6795, 6796, 6797, 6798, 6799, 6810, 6811, 6812, 6813, 6814, 6815,
		6830, 6831, 6989, 6990, 6991, 7039, 7156, 7157, 7158, 7159, 7160, 7161,
		7162, 7163, 7224, 7225, 7226, 7242, 7243, 7244, 7305, 7306, 7307, 7308,
		7309, 7310, 7311, 7355, 7356, 7368, 7369, 7370, 7371, 7372, 7373, 7374,
		7375, 7419, 7420, 7421, 7422, 7423, 7958, 7959, 7966, 7967, 8006, 8007,
		8014, 8015, 8024, 8026, 8028, 8030, 8062, 8063, 8117, 8133, 8148, 8149,
		8156, 8176, 8177, 8181, 8191, 8293, 8306, 8307, 8335, 8349, 8350, 8351,
		8588, 8589, 8590, 8591, 11124, 11125, 11158, 11508, 11509, 11510, 11511,
		11512, 11558, 11560, 11561, 11562, 11563, 11564, 11566, 11567, 11624, 11625,
		11626, 11627, 11628, 11629, 11630, 11687, 11695, 11703, 11711, 11719, 11727,
		11735, 11743, 11930, 12284, 12285, 12286, 12287, 12352, 12439, 12440, 12544,
		12545, 12546, 12547, 12548, 12592, 12687, 12831, 42125, 42126, 42127, 42744,
		42745, 42746, 42747, 42748, 42749, 42750, 42751, 42955, 42956, 42957, 42958,
		42959, 42962, 42964, 43053, 43054, 43055, 43066, 43067, 43068, 43069, 43070,
		43071, 43128, 43129, 43130, 43131, 43132, 43133, 43134, 43135, 43206, 43207,
		43208, 43209, 43210, 43211, 43212, 43213, 43226, 43227, 43228, 43229, 43230,
		43231, 43389, 43390, 43391, 43470, 43482, 43483, 43484, 43485, 43519, 43598,
		43599, 43610, 43611, 43783, 43784, 43791, 43792, 43815, 43823, 43884, 43885,
		43886, 43887, 44014, 44015, 44026, 44027, 44028, 44029, 44030, 44031, 55239,
		55240, 55241, 55242, 64110, 64111, 64280, 64281, 64282, 64283, 64284, 64311,
		64317, 64319, 64322, 64325, 64912, 64913, 64968, 64969, 64970, 64971, 64972,
		64973, 64974, 65050, 65051, 65052, 65053, 65054, 65055, 65107, 65127, 65132,
		65133, 65134, 65135, 65141, 65277, 65278, 65280, 65471, 65472, 65473, 65480,
		65481, 65488, 65489, 65496, 65497, 65501, 65502, 65503, 65511, 65534, 65535,
		65548, 65575, 65595, 65598, 65614, 65615, 65787, 65788, 65789, 65790, 65791,
		65795, 65796, 65797, 65798, 65844, 65845, 65846, 65935, 65949, 65950, 65951,
		66205, 66206, 66207, 66300, 66301, 66302, 66303, 66379, 66380, 66381, 66382,
		66383, 66427, 66428, 66429, 66430, 66431, 66462, 66500, 66501, 66502, 66503,
		66718, 66719, 66730, 66731, 66732, 66733, 66734, 66735, 66772, 66773, 66774,
		66775, 66812, 66813, 66814, 66815, 66856, 66857, 66858, 66859, 66860, 66861,
		66862, 66863, 66939, 66955, 66963, 66966, 66978, 66994, 67002, 67462, 67505,
		67590, 67591, 67593, 67638, 67641, 67642, 67643, 67645, 67646, 67670, 67743,
		67744, 67745, 67746, 67747, 67748, 67749, 67750, 67827, 67830, 67831, 67832,
		67833, 67834, 67868, 67869, 67870, 67898, 67899, 67900, 67901, 67902, 68024,
		68025, 68026, 68027, 68048, 68049, 68100, 68103, 68104, 68105, 68106, 68107,
		68116, 68120, 68150, 68151, 68155, 68156, 68157, 68158, 68169, 68170, 68171,
		68172, 68173, 68174, 68175, 68185, 68186, 68187, 68188, 68189, 68190, 68191,
		68327, 68328, 68329, 68330, 68406, 68407, 68408, 68438, 68439, 68467, 68468,
		68469, 68470, 68471, 68498, 68499, 68500, 68501, 68502, 68503, 68504, 68851,
		68852, 68853, 68854, 68855, 68856, 68857, 68904, 68905, 68906, 68907, 68908,
		68909, 68910, 68911, 69247, 69290, 69294, 69295, 69416, 69417, 69418, 69419,
		69420, 69421, 69422, 69423, 69710, 69711, 69712, 69713, 69838, 69839, 69865,
		69866, 69867, 69868, 69869, 69870, 69871, 69882, 69883, 69884, 69885, 69886,
		69887, 69941, 69960, 69961, 69962, 69963, 69964, 69965, 69966, 69967, 70112,
		70162, 70279, 70281, 70286, 70302, 70314, 70315, 70316, 70317, 70318, 70319,
		70379, 70380, 70381, 70382, 70383, 70394, 70395, 70396, 70397, 70398, 70399,
		70404, 70413, 70414, 70417, 70418, 70441, 70449, 70452, 70458, 70469, 70470,
		70473, 70474, 70478, 70479, 70481, 70482, 70483, 70484, 70485, 70486, 70488,
		70489, 70490, 70491, 70492, 70500, 70501, 70509, 70510, 70511, 70748, 70856,
		70857, 70858, 70859, 70860, 70861, 70862, 70863, 71094, 71095, 71258, 71259,
		71260, 71261, 71262, 71263, 71354, 71355, 71356, 71357, 71358, 71359, 71451,
		71452, 71468, 71469, 71470, 71471, 71943, 71944, 71946, 71947, 71956, 71959,
		71990, 71993, 71994, 72104, 72105, 72152, 72153, 72264, 72265, 72266, 72267,
		72268, 72269, 72270, 72271, 72713, 72759, 72813, 72814, 72815, 72848, 72849,
		72872, 72967, 72970, 73015, 73016, 73017, 73019, 73022, 73032, 73033, 73034,
		73035, 73036, 73037, 73038, 73039, 73050, 73051, 73052, 73053, 73054, 73055,
		73062, 73065, 73103, 73106, 73113, 73114, 73115, 73116, 73117, 73118, 73119,
		74863, 78895, 92729, 92730, 92731, 92732, 92733, 92734, 92735, 92767, 92778,
		92779, 92780, 92781, 92863, 92874, 92875, 92876, 92877, 92878, 92879, 92910,
		92911, 93018, 93026, 93048, 93049, 93050, 93051, 93052, 94027, 94028, 94029,
		94030, 94088, 94089, 94090, 94091, 94092, 94093, 94094, 100344, 100345,
		100346, 100347, 100348, 100349, 100350, 100351, 110580, 110588, 110591,
		110952, 110953, 110954, 110955, 110956, 110957, 110958, 110959, 113771,
		113772, 113773, 113774, 113775, 113789, 113790, 113791, 113801, 113802,
		113803, 113804, 113805, 113806, 113807, 113818, 113819, 118574, 118575,
		119079, 119080, 119893, 119965, 119968, 119969, 119971, 119972, 119975,
		119976, 119981, 119994, 119996, 120004, 120070, 120075, 120076, 120085,
		120093, 120122, 120127, 120133, 120135, 120136, 120137, 120145, 120486,
		120487, 120780, 120781, 121504, 122887, 122905, 122906, 122914, 122917,
		123181, 123182, 123183, 123198, 123199, 123210, 123211, 123212, 123213,
		123642, 123643, 123644, 123645, 123646, 124903, 124908, 124911, 124927,
		125125, 125126, 125260, 125261, 125262, 125263, 125274, 125275, 125276,
		125277, 126468, 126496, 126499, 126501, 126502, 126504, 126515, 126520,
		126522, 126524, 126525, 126526, 126527, 126528, 126529, 126531, 126532,
		126533, 126534, 126536, 126538, 126540, 126544, 126547, 126549, 126550,
		126552, 126554, 126556, 126558, 126560, 126563, 126565, 126566, 126571,
		126579, 126584, 126589, 126591, 126602, 126620, 126621, 126622, 126623,
		126624, 126628, 126634, 127020, 127021, 127022, 127023, 127151, 127152,
		127168, 127184, 127548, 127549, 127550, 127551, 127561, 127562, 127563,
		127564, 127565, 127566, 127567, 128728, 128729, 128730, 128731, 128732,
		128749, 128750, 128751, 128765, 128766, 128767, 128985, 128986, 128987,
		128988, 128989, 128990, 128991, 129004, 129005, 129006, 129007, 129036,
		129037, 129038, 129039, 129096, 129097, 129098, 129099, 129100, 129101,
		129102, 129103, 129114, 129115, 129116, 129117, 129118, 129119, 129160,
		129161, 129162, 129163, 129164, 129165, 129166, 129167, 129198, 129199,
		129646, 129647, 129653, 129654, 129655, 129661, 129662, 129663, 129709,
		129710, 129711, 129723, 129724, 129725, 129726, 129727, 129754, 129755,
		129756, 129757, 129758, 129759, 129768, 129769, 129770, 129771, 129772,
		129773, 129774, 129775, 129939, 177977, 177978, 177979, 177980, 177981,
		177982, 177983, 178206, 178207:
		return true
	case 1421, 1422, 1423, 1519, 1520, 1521, 1522, 1523, 1524, 2142, 2192, 2193,
		2437, 2438, 2439, 2440, 2441, 2442, 2443, 2444, 2447, 2448, 2474, 2475,
		2476, 2477, 2478, 2479, 2480, 2482, 2486, 2487, 2488, 2489, 2503, 2504,
		2507, 2508, 2509, 2510, 2519, 2524, 2525, 2527, 2528, 2529, 2530, 2531,
		2561, 2562, 2563, 2565, 2566, 2567, 2568, 2569, 2570, 2575, 2576, 2602,
		2603, 2604, 2605, 2606, 2607, 2608, 2610, 2611, 2613, 2614, 2616, 2617,
		2620, 2622, 2623, 2624, 2625, 2626, 2631, 2632, 2635, 2636, 2637, 2641,
		2649, 2650, 2651, 2652, 2654, 2689, 2690, 2691, 2703, 2704, 2705, 2730,
		2731, 2732, 2733, 2734, 2735, 2736, 2738, 2739, 2741, 2742, 2743, 2744,
		2745, 2759, 2760, 2761, 2763, 2764, 2765, 2768, 2784, 2785, 2786, 2787,
		2809, 2810, 2811, 2812, 2813, 2814, 2815, 2817, 2818, 2819, 2821, 2822,
		2823, 2824, 2825, 2826, 2827, 2828, 2831, 2832, 2858, 2859, 2860, 2861,
		2862, 2863, 2864, 2866, 2867, 2869, 2870, 2871, 2872, 2873, 2887, 2888,
		2891, 2892, 2893, 2901, 2902, 2903, 2908, 2909, 2911, 2912, 2913, 2914,
		2915, 2946, 2947, 2949, 2950, 2951, 2952, 2953, 2954, 2958, 2959, 2960,
		2962, 2963, 2964, 2965, 2969, 2970, 2972, 2974, 2975, 2979, 2980, 2984,
		2985, 2986, 3006, 3007, 3008, 3009, 3010, 3014, 3015, 3016, 3018, 3019,
		3020, 3021, 3024, 3031, 3086, 3087, 3088, 3142, 3143, 3144, 3146, 3147,
		3148, 3149, 3157, 3158, 3160, 3161, 3162, 3165, 3168, 3169, 3170, 3171,
		3214, 3215, 3216, 3253, 3254, 3255, 3256, 3257, 3270, 3271, 3272, 3274,
		3275, 3276, 3277, 3285, 3286, 3293, 3294, 3296, 3297, 3298, 3299, 3313,
		3314, 3342, 3343, 3344, 3398, 3399, 3400, 3402, 3403, 3404, 3405, 3406,
		3407, 3457, 3458, 3459, 3517, 3520, 3521, 3522, 3523, 3524, 3525, 3526,
		3530, 3535, 3536, 3537, 3538, 3539, 3540, 3542, 3544, 3545, 3546, 3547,
		3548, 3549, 3550, 3551, 3570, 3571, 3572, 3713, 3714, 3716, 3718, 3719,
		3720, 3721, 3722, 3749, 3776, 3777, 3778, 3779, 3780, 3782, 3784, 3785,
		3786, 3787, 3788, 3789, 3804, 3805, 3806, 3807, 4295, 4301, 4682, 4683,
		4684, 4685, 4688, 4689, 4690, 4691, 4692, 4693, 4694, 4696, 4698, 4699,
		4700, 4701, 4746, 4747, 4748, 4749, 4786, 4787, 4788, 4789, 4792, 4793,
		4794, 4795, 4796, 4797, 4798, 4800, 4802, 4803, 4804, 4805, 4882, 4883,
		4884, 4885, 5112, 5113, 5114, 5115, 5116, 5117, 5998, 5999, 6000, 6002,
		6003, 6464, 6512, 6513, 6514, 6515, 6516, 7960, 7961, 7962, 7963, 7964,
		7965, 8008, 8009, 8010, 8011, 8012, 8013, 8016, 8017, 8018, 8019, 8020,
		8021, 8022, 8023, 8025, 8027, 8029, 8150, 8151, 8152, 8153, 8154, 8155,
		8178, 8179, 8180, 11559, 11565, 11631, 11632, 11680, 11681, 11682, 11683,
		11684, 11685, 11686, 11688, 11689, 11690, 11691, 11692, 11693, 11694, 11696,
		11697, 11698, 11699, 11700, 11701, 11702, 11704, 11705, 11706, 11707, 11708,
		11709, 11710, 11712, 11713, 11714, 11715, 11716, 11717, 11718, 11720, 11721,
		11722, 11723, 11724, 11725, 11726, 11728, 11729, 11730, 11731, 11732, 11733,
		11734, 11736, 11737, 11738, 11739, 11740, 11741, 11742, 42960, 42961, 42963,
		42965, 42966, 42967, 42968, 42969, 43777, 43778, 43779, 43780, 43781, 43782,
		43785, 43786, 43787, 43788, 43789, 43790, 43793, 43794, 43795, 43796, 43797,
		43798, 43808, 43809, 43810, 43811, 43812, 43813, 43814, 43816, 43817, 43818,
		43819, 43820, 43821, 43822, 64256, 64257, 64258, 64259, 64260, 64261, 64262,
		64275, 64276, 64277, 64278, 64279, 64312, 64313, 64314, 64315, 64316, 64318,
		64320, 64321, 64323, 64324, 64975, 65128, 65129, 65130, 65131, 65136, 65137,
		65138, 65139, 65140, 65279, 65474, 65475, 65476, 65477, 65478, 65479, 65482,
		65483, 65484, 65485, 65486, 65487, 65490, 65491, 65492, 65493, 65494, 65495,
		65498, 65499, 65500, 65504, 65505, 65506, 65507, 65508, 65509, 65510, 65512,
		65513, 65514, 65515, 65516, 65517, 65518, 65529, 65530, 65531, 65532, 65533,
		65596, 65597, 65792, 65793, 65794, 65952, 66956, 66957, 66958, 66959, 66960,
		66961, 66962, 66964, 66965, 66995, 66996, 66997, 66998, 66999, 67000, 67001,
		67003, 67004, 67424, 67425, 67426, 67427, 67428, 67429, 67430, 67431, 67456,
		67457, 67458, 67459, 67460, 67461, 67584, 67585, 67586, 67587, 67588, 67589,
		67592, 67639, 67640, 67644, 67828, 67829, 67903, 68101, 68102, 68108, 68109,
		68110, 68111, 68112, 68113, 68114, 68115, 68117, 68118, 68119, 68152, 68153,
		68154, 68505, 68506, 68507, 68508, 68521, 68522, 68523, 68524, 68525, 68526,
		68527, 69291, 69292, 69293, 69296, 69297, 69837, 70272, 70273, 70274, 70275,
		70276, 70277, 70278, 70280, 70282, 70283, 70284, 70285, 70400, 70401, 70402,
		70403, 70405, 70406, 70407, 70408, 70409, 70410, 70411, 70412, 70415, 70416,
		70442, 70443, 70444, 70445, 70446, 70447, 70448, 70450, 70451, 70453, 70454,
		70455, 70456, 70457, 70471, 70472, 70475, 70476, 70477, 70480, 70487, 70493,
		70494, 70495, 70496, 70497, 70498, 70499, 70502, 70503, 70504, 70505, 70506,
		70507, 70508, 70512, 70513, 70514, 70515, 70516, 70749, 70750, 70751, 70752,
		70753, 71935, 71936, 71937, 71938, 71939, 71940, 71941, 71942, 71945, 71948,
		71949, 71950, 71951, 71952, 71953, 71954, 71955, 71957, 71958, 71991, 71992,
		72096, 72097, 72098, 72099, 72100, 72101, 72102, 72103, 72960, 72961, 72962,
		72963, 72964, 72965, 72966, 72968, 72969, 73018, 73020, 73021, 73056, 73057,
		73058, 73059, 73060, 73061, 73063, 73064, 73104, 73105, 73107, 73108, 73109,
		73110, 73111, 73112, 73648, 74864, 74865, 74866, 74867, 74868, 92912, 92913,
		92914, 92915, 92916, 92917, 93019, 93020, 93021, 93022, 93023, 93024, 93025,
		94176, 94177, 94178, 94179, 94180, 94192, 94193, 110576, 110577, 110578,
		110579, 110581, 110582, 110583, 110584, 110585, 110586, 110587, 110589,
		110590, 110928, 110929, 110930, 110948, 110949, 110950, 110951, 113820,
		113821, 113822, 113823, 113824, 113825, 113826, 113827, 119966, 119967,
		119970, 119973, 119974, 119977, 119978, 119979, 119980, 119995, 119997,
		119998, 119999, 120000, 120001, 120002, 120003, 120071, 120072, 120073,
		120074, 120077, 120078, 120079, 120080, 120081, 120082, 120083, 120084,
		120086, 120087, 120088, 120089, 120090, 120091, 120092, 120123, 120124,
		120125, 120126, 120128, 120129, 120130, 120131, 120132, 120134, 120138,
		120139, 120140, 120141, 120142, 120143, 120144, 121499, 121500, 121501,
		121502, 121503, 122880, 122881, 122882, 122883, 122884, 122885, 122886,
		122907, 122908, 122909, 122910, 122911, 122912, 122913, 122915, 122916,
		122918, 122919, 122920, 122921, 122922, 123214, 123215, 123647, 124896,
		124897, 124898, 124899, 124900, 124901, 124902, 124904, 124905, 124906,
		124907, 124909, 124910, 125278, 125279, 126464, 126465, 126466, 126467,
		126497, 126498, 126500, 126503, 126516, 126517, 126518, 126519, 126521,
		126523, 126530, 126535, 126537, 126539, 126541, 126542, 126543, 126545,
		126546, 126548, 126551, 126553, 126555, 126557, 126559, 126561, 126562,
		126564, 126567, 126568, 126569, 126570, 126572, 126573, 126574, 126575,
		126576, 126577, 126578, 126580, 126581, 126582, 126583, 126585, 126586,
		126587, 126588, 126590, 126625, 126626, 126627, 126629, 126630, 126631,
		126632, 126633, 126704, 126705, 127568, 127569, 127584, 127585, 127586,
		127587, 127588, 127589, 129008, 129200, 129201, 129648, 129649, 129650,
		129651, 129652, 129656, 129657, 129658, 129659, 129660, 129664, 129665,
		129666, 129667, 129668, 129669, 129670, 129728, 129729, 129730, 129731,
		129732, 129733, 129760, 129761, 129762, 129763, 129764, 129765, 129766,
		129767, 129776, 129777, 129778, 129779, 129780, 129781, 129782, 917505:
		return false
	}
	return r >= 1525 && (r <= 71039 && (r <= 55215 && (r <= 6911 && (r <= 3712 &&
		(r <= 2783 && (r <= 1983 && (r <= 1535 || r >= 1970) || r >= 2679 && (r <= 2692 ||
			r >= 2769)) || r >= 2936 && (r <= 3045 && (r <= 2989 || r >= 3032) || r >= 3315 &&
			(r <= 3327 || r >= 3573 && (r <= 3584 || r >= 3676)))) || r >= 3808 && (r <= 5951 &&
		(r <= 4095 && (r <= 3839 || r >= 4059) || r >= 5910 && (r <= 5918 || r >= 5943)) ||
		r >= 5972 && (r <= 6015 && (r <= 5983 || r >= 6004) || r >= 6390 && (r <= 6399 ||
			r >= 6517 && (r <= 6527 || r >= 6863))))) || r >= 8385 && (r <= 12271 &&
		(r <= 9311 && (r <= 8447 && (r <= 8399 || r >= 8433) || r >= 9255 && (r <= 9279 ||
			r >= 9291)) || r >= 11633 && (r <= 11679 && (r <= 11646 || r >= 11671) ||
			r >= 11870 && (r <= 11903 || r >= 12020 && (r <= 12031 || r >= 12246)))) ||
		r >= 12772 && (r <= 42993 && (r <= 42191 && (r <= 12783 || r >= 42183) ||
			r >= 42540 && (r <= 42559 || r >= 42970)) || r >= 43348 && (r <= 43583 &&
			(r <= 43358 || r >= 43575) || r >= 43715 && (r <= 43738 || r >= 43767 &&
			(r <= 43823 || r >= 55204)))))) || r >= 55292 && (r <= 67967 && (r <= 66271 &&
		(r <= 65007 && (r <= 64255 && (r <= 63743 || r >= 64218) || r >= 64263 &&
			(r <= 64284 || r >= 64451 && (r <= 64466 || r >= 64976))) || r >= 65519 &&
			(r <= 65663 && (r <= 65535 || r >= 65630) || r >= 65953 && (r <= 65999 ||
				r >= 66046 && (r <= 66175 || r >= 66257)))) || r >= 66340 && (r <= 67071 &&
		(r <= 66559 && (r <= 66348 || r >= 66518) || r >= 66916 && (r <= 66926 ||
			r >= 67005)) || r >= 67383 && (r <= 67423 && (r <= 67391 || r >= 67414) ||
		r >= 67432 && (r <= 67583 && (r <= 67462 || r >= 67515) || r >= 67760 &&
			(r <= 67807 || r >= 67904))))) || r >= 68256 && (r <= 69551 && (r <= 68735 &&
		(r <= 68351 && (r <= 68287 || r >= 68343) || r >= 68509 && (r <= 68607 ||
			r >= 68681)) || r >= 68787 && (r <= 69215 && (r <= 68799 || r >= 68922) ||
		r >= 69298 && (r <= 69375 || r >= 69466 && (r <= 69487 || r >= 69514)))) ||
		r >= 69580 && (r <= 70015 && (r <= 69631 && (r <= 69599 || r >= 69623) ||
			r >= 69750 && (r <= 69758 || r >= 69827 && (r <= 69839 || r >= 70007))) ||
			r >= 70133 && (r <= 70271 && (r <= 70143 || r >= 70207) || r >= 70517 &&
				(r <= 70655 || r >= 70754 && (r <= 70783 || r >= 70874))))))) || r >= 71134 &&
		(r <= 119295 && (r <= 74751 && (r <= 72095 && (r <= 71423 && (r <= 71247 &&
			(r <= 71167 || r >= 71237) || r >= 71277 && (r <= 71295 || r >= 71370)) ||
			r >= 71495 && (r <= 71839 && (r <= 71679 || r >= 71740) || r >= 71923 &&
				(r <= 71959 || r >= 72007 && (r <= 72015 || r >= 72026)))) || r >= 72165 &&
			(r <= 72783 && (r <= 72367 && (r <= 72191 || r >= 72355) || r >= 72441 &&
				(r <= 72703 || r >= 72774)) || r >= 72887 && (r <= 73439 && (r <= 72970 ||
				r >= 73130) || r >= 73465 && (r <= 73663 || r >= 73714 && (r <= 73726 ||
				r >= 74650))))) || r >= 74869 && (r <= 93951 && (r <= 82943 && (r <= 77711 &&
			(r <= 74879 || r >= 75076) || r >= 77811 && (r <= 77823 || r >= 78905)) ||
			r >= 83527 && (r <= 92927 && (r <= 92159 || r >= 92918) || r >= 92998 &&
				(r <= 93007 || r >= 93072 && (r <= 93759 || r >= 93851)))) || r >= 94112 &&
			(r <= 113663 && (r <= 101631 && (r <= 94207 || r >= 101590) || r >= 101641 &&
				(r <= 110591 || r >= 110883 && (r <= 110959 || r >= 111356))) || r >= 113828 &&
				(r <= 118607 && (r <= 118527 || r >= 118599) || r >= 118724 && (r <= 118783 ||
					r >= 119030 && (r <= 119039 || r >= 119275)))))) || r >= 119366 && (r <= 127231 &&
			(r <= 123535 && (r <= 119807 && (r <= 119551 && (r <= 119519 || r >= 119540) ||
				r >= 119639 && (r <= 119647 || r >= 119673)) || r >= 121484 && (r <= 122623 &&
				(r <= 121504 || r >= 121520) || r >= 122655 && (r <= 122887 || r >= 122923 &&
				(r <= 123135 || r >= 123216)))) || r >= 123567 && (r <= 126064 && (r <= 124895 &&
				(r <= 123583 || r >= 123648) || r >= 125143 && (r <= 125183 || r >= 125280)) ||
				r >= 126133 && (r <= 126463 && (r <= 126208 || r >= 126270) || r >= 126652 &&
					(r <= 126975 || r >= 127124 && (r <= 127135 || r >= 127222))))) || r >= 127406 &&
			(r <= 129743 && (r <= 128895 && (r <= 127503 && (r <= 127461 || r >= 127491) ||
				r >= 127570 && (r <= 127743 || r >= 128884)) || r >= 129009 && (r <= 129279 &&
				(r <= 129023 || r >= 129202) || r >= 129620 && (r <= 129631 || r >= 129671 &&
				(r <= 129679 || r >= 129734)))) || r >= 129783 && (r <= 183983 && (r <= 130031 &&
				(r <= 129791 || r >= 129995) || r >= 130042 && (r <= 131071 || r >= 173792 &&
				(r <= 173823 || r >= 183970))) || r >= 191457 && (r <= 196607 && (r <= 194559 ||
				r >= 195102) || r >= 201547 && (r <= 917535 || r >= 917632 && (r <= 917759 ||
				r >= 918000)))))))) && r <= 1114111
}
