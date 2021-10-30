<h1 align="center">Unicode® package for script queries</h1>

**uscripts** is a Go package for querying the Unicode **Script_Extensions** property. This property is similar to the more familiar **Script** property and produces the same results for most code points. The significant difference is that, for a given code point, the Script property only assigns at most one script, while the Script_Extensions property may assign multiple scripts. Consider the following examples (from Unicode version 14.0.0).

| Code Points | Script    | Script_Extensions | Sample uscripts test                     |
|-------------|-----------|-------------------|------------------------------------------|
| U+0041 'A'  | Latin     | Latin             | `HasScxLatn(0x41)`                       |
| U+0485      | Inherited | Cyrillic, Latin   | `HasScxCyrl(0x485)`, `HasScxLatn(0x485)` |

For more information about the Script and Script_Extensions properties, consult [UAX #24](https://www.unicode.org/reports/tr24/).

Other Unicode properties are supported by other packages in this module. For a summary of Unicode properties and which functions of which packages of this module provide the corresponding support, peruse the file, **SUMMARY.txt**, in the top level directory of this module.

For example, the function, `HasScxGeor(rune r) bool`, returns true only if the code point, r, has Unicode property, Script_Extensions, value, Georgian (Geor).

_Unicode and the Unicode Logo are registered trademarks of Unicode, Inc. in the United States and other countries._

## Example

If you are new to the Go language, consider perusing the documentation at [golang.org](https://golang.org/doc/) before proceeding.

Download and install it:

```sh
go get github.com/orkvozku/go/uni

```
Usage:

```go
package main

import "github.com/orkvozku/go/uni/uscripts"
import "fmt"

func main() {
    text := "Hi!"
    for _, r := range text {
        if uscripts.HasScxLatn(r) {
            fmt.Printf("'%c' is Latin\n", r)
        } else {
            fmt.Printf("'%c' is not Latin\n", r)
        }
    }

    // output:
    // 'H' is Latin
    // 'i' is Latin
    // '!' is not Latin
}
```
## API
### General Constants
#### UnicodeVersion
UnicodeVersion is the current version of the Unicode specification that this package is intended to comply with.

UnicodeVersion = "14.0.0".
### Unicode Property Queries
#### func HasScxAdlm(r rune) bool
HasScxAdlm returns true if code point, r, has Unicode property, Script_Extensions (scx), of Adlam (Adlm).
#### func HasScxAghb(r rune) bool
HasScxAghb returns true if code point, r, has Unicode property, Script_Extensions (scx), of Caucasian_Albanian (Aghb).
#### func HasScxAhom(r rune) bool
HasScxAhom returns true if code point, r, has Unicode property, Script_Extensions (scx), of Ahom.
#### func HasScxArab(r rune) bool
HasScxArab returns true if code point, r, has Unicode property, Script_Extensions (scx), of Arabic (Arab).
#### func HasScxArmi(r rune) bool
HasScxArmi returns true if code point, r, has Unicode property, Script_Extensions (scx), of Imperial_Aramaic (Armi).
#### func HasScxArmn(r rune) bool
HasScxArmn returns true if code point, r, has Unicode property, Script_Extensions (scx), of Armenian (Armn).
#### func HasScxAvst(r rune) bool
HasScxAvst returns true if code point, r, has Unicode property, Script_Extensions (scx), of Avestan (Avst).
#### func HasScxBali(r rune) bool
HasScxBali returns true if code point, r, has Unicode property, Script_Extensions (scx), of Balinese (Bali).
#### func HasScxBamu(r rune) bool
HasScxBamu returns true if code point, r, has Unicode property, Script_Extensions (scx), of Bamum (Bamu).
#### func HasScxBass(r rune) bool
HasScxBass returns true if code point, r, has Unicode property, Script_Extensions (scx), of Bassa_Vah (Bass).
#### func HasScxBatk(r rune) bool
HasScxBatk returns true if code point, r, has Unicode property, Script_Extensions (scx), of Batak (Batk).
#### func HasScxBeng(r rune) bool
HasScxBeng returns true if code point, r, has Unicode property, Script_Extensions (scx), of Bengali (Beng).
#### func HasScxBhks(r rune) bool
HasScxBhks returns true if code point, r, has Unicode property, Script_Extensions (scx), of Bhaiksuki (Bhks).
#### func HasScxBopo(r rune) bool
HasScxBopo returns true if code point, r, has Unicode property, Script_Extensions (scx), of Bopomofo (Bopo).
#### func HasScxBrah(r rune) bool
HasScxBrah returns true if code point, r, has Unicode property, Script_Extensions (scx), of Brahmi (Brah).
#### func HasScxBrai(r rune) bool
HasScxBrai returns true if code point, r, has Unicode property, Script_Extensions (scx), of Braille (Brai).
#### func HasScxBugi(r rune) bool
HasScxBugi returns true if code point, r, has Unicode property, Script_Extensions (scx), of Buginese (Bugi).
#### func HasScxBuhd(r rune) bool
HasScxBuhd returns true if code point, r, has Unicode property, Script_Extensions (scx), of Buhid (Buhd).
#### func HasScxCakm(r rune) bool
HasScxCakm returns true if code point, r, has Unicode property, Script_Extensions (scx), of Chakma (Cakm).
#### func HasScxCans(r rune) bool
HasScxCans returns true if code point, r, has Unicode property, Script_Extensions (scx), of Canadian_Aboriginal (Cans).
#### func HasScxCari(r rune) bool
HasScxCari returns true if code point, r, has Unicode property, Script_Extensions (scx), of Carian (Cari).
#### func HasScxCham(r rune) bool
HasScxCham returns true if code point, r, has Unicode property, Script_Extensions (scx), of Cham.
#### func HasScxCher(r rune) bool
HasScxCher returns true if code point, r, has Unicode property, Script_Extensions (scx), of Cherokee (Cher).
#### func HasScxChrs(r rune) bool
HasScxChrs returns true if code point, r, has Unicode property, Script_Extensions (scx), of Chorasmian (Chrs).
#### func HasScxCopt(r rune) bool
HasScxCopt returns true if code point, r, has Unicode property, Script_Extensions (scx), of Coptic (Copt).
#### func HasScxCpmn(r rune) bool
HasScxCpmn returns true if code point, r, has Unicode property, Script_Extensions (scx), of Cypro_Minoan (Cpmn).
#### func HasScxCprt(r rune) bool
HasScxCprt returns true if code point, r, has Unicode property, Script_Extensions (scx), of Cypriot (Cprt).
#### func HasScxCyrl(r rune) bool
HasScxCyrl returns true if code point, r, has Unicode property, Script_Extensions (scx), of Cyrillic (Cyrl).
#### func HasScxDeva(r rune) bool
HasScxDeva returns true if code point, r, has Unicode property, Script_Extensions (scx), of Devanagari (Deva).
#### func HasScxDiak(r rune) bool
HasScxDiak returns true if code point, r, has Unicode property, Script_Extensions (scx), of Dives_Akuru (Diak).
#### func HasScxDogr(r rune) bool
HasScxDogr returns true if code point, r, has Unicode property, Script_Extensions (scx), of Dogra (Dogr).
#### func HasScxDsrt(r rune) bool
HasScxDsrt returns true if code point, r, has Unicode property, Script_Extensions (scx), of Deseret (Dsrt).
#### func HasScxDupl(r rune) bool
HasScxDupl returns true if code point, r, has Unicode property, Script_Extensions (scx), of Duployan (Dupl).
#### func HasScxEgyp(r rune) bool
HasScxEgyp returns true if code point, r, has Unicode property, Script_Extensions (scx), of Egyptian_Hieroglyphs (Egyp).
#### func HasScxElba(r rune) bool
HasScxElba returns true if code point, r, has Unicode property, Script_Extensions (scx), of Elbasan (Elba).
#### func HasScxElym(r rune) bool
HasScxElym returns true if code point, r, has Unicode property, Script_Extensions (scx), of Elymaic (Elym).
#### func HasScxEthi(r rune) bool
HasScxEthi returns true if code point, r, has Unicode property, Script_Extensions (scx), of Ethiopic (Ethi).
#### func HasScxGeor(r rune) bool
HasScxGeor returns true if code point, r, has Unicode property, Script_Extensions (scx), of Georgian (Geor).
#### func HasScxGlag(r rune) bool
HasScxGlag returns true if code point, r, has Unicode property, Script_Extensions (scx), of Glagolitic (Glag).
#### func HasScxGong(r rune) bool
HasScxGong returns true if code point, r, has Unicode property, Script_Extensions (scx), of Gunjala_Gondi (Gong).
#### func HasScxGonm(r rune) bool
HasScxGonm returns true if code point, r, has Unicode property, Script_Extensions (scx), of Masaram_Gondi (Gonm).
#### func HasScxGoth(r rune) bool
HasScxGoth returns true if code point, r, has Unicode property, Script_Extensions (scx), of Gothic (Goth).
#### func HasScxGran(r rune) bool
HasScxGran returns true if code point, r, has Unicode property, Script_Extensions (scx), of Grantha (Gran).
#### func HasScxGrek(r rune) bool
HasScxGrek returns true if code point, r, has Unicode property, Script_Extensions (scx), of Greek (Grek).
#### func HasScxGujr(r rune) bool
HasScxGujr returns true if code point, r, has Unicode property, Script_Extensions (scx), of Gujarati (Gujr).
#### func HasScxGuru(r rune) bool
HasScxGuru returns true if code point, r, has Unicode property, Script_Extensions (scx), of Gurmukhi (Guru).
#### func HasScxHang(r rune) bool
HasScxHang returns true if code point, r, has Unicode property, Script_Extensions (scx), of Hangul (Hang).
#### func HasScxHani(r rune) bool
HasScxHani returns true if code point, r, has Unicode property, Script_Extensions (scx), of Han (Hani).
#### func HasScxHano(r rune) bool
HasScxHano returns true if code point, r, has Unicode property, Script_Extensions (scx), of Hanunoo (Hano).
#### func HasScxHatr(r rune) bool
HasScxHatr returns true if code point, r, has Unicode property, Script_Extensions (scx), of Hatran (Hatr).
#### func HasScxHebr(r rune) bool
HasScxHebr returns true if code point, r, has Unicode property, Script_Extensions (scx), of Hebrew (Hebr).
#### func HasScxHira(r rune) bool
HasScxHira returns true if code point, r, has Unicode property, Script_Extensions (scx), of Hiragana (Hira).
#### func HasScxHluw(r rune) bool
HasScxHluw returns true if code point, r, has Unicode property, Script_Extensions (scx), of Anatolian_Hieroglyphs (Hluw).
#### func HasScxHmng(r rune) bool
HasScxHmng returns true if code point, r, has Unicode property, Script_Extensions (scx), of Pahawh_Hmong (Hmng).
#### func HasScxHmnp(r rune) bool
HasScxHmnp returns true if code point, r, has Unicode property, Script_Extensions (scx), of Nyiakeng_Puachue_Hmong (Hmnp).
#### func HasScxHung(r rune) bool
HasScxHung returns true if code point, r, has Unicode property, Script_Extensions (scx), of Old_Hungarian (Hung).
#### func HasScxItal(r rune) bool
HasScxItal returns true if code point, r, has Unicode property, Script_Extensions (scx), of Old_Italic (Ital).
#### func HasScxJava(r rune) bool
HasScxJava returns true if code point, r, has Unicode property, Script_Extensions (scx), of Javanese (Java).
#### func HasScxKali(r rune) bool
HasScxKali returns true if code point, r, has Unicode property, Script_Extensions (scx), of Kayah_Li (Kali).
#### func HasScxKana(r rune) bool
HasScxKana returns true if code point, r, has Unicode property, Script_Extensions (scx), of Katakana (Kana).
#### func HasScxKhar(r rune) bool
HasScxKhar returns true if code point, r, has Unicode property, Script_Extensions (scx), of Kharoshthi (Khar).
#### func HasScxKhmr(r rune) bool
HasScxKhmr returns true if code point, r, has Unicode property, Script_Extensions (scx), of Khmer (Khmr).
#### func HasScxKhoj(r rune) bool
HasScxKhoj returns true if code point, r, has Unicode property, Script_Extensions (scx), of Khojki (Khoj).
#### func HasScxKits(r rune) bool
HasScxKits returns true if code point, r, has Unicode property, Script_Extensions (scx), of Khitan_Small_Script (Kits).
#### func HasScxKnda(r rune) bool
HasScxKnda returns true if code point, r, has Unicode property, Script_Extensions (scx), of Kannada (Knda).
#### func HasScxKthi(r rune) bool
HasScxKthi returns true if code point, r, has Unicode property, Script_Extensions (scx), of Kaithi (Kthi).
#### func HasScxLana(r rune) bool
HasScxLana returns true if code point, r, has Unicode property, Script_Extensions (scx), of Tai_Tham (Lana).
#### func HasScxLaoo(r rune) bool
HasScxLaoo returns true if code point, r, has Unicode property, Script_Extensions (scx), of Lao (Laoo).
#### func HasScxLatn(r rune) bool
HasScxLatn returns true if code point, r, has Unicode property, Script_Extensions (scx), of Latin (Latn).
#### func HasScxLepc(r rune) bool
HasScxLepc returns true if code point, r, has Unicode property, Script_Extensions (scx), of Lepcha (Lepc).
#### func HasScxLimb(r rune) bool
HasScxLimb returns true if code point, r, has Unicode property, Script_Extensions (scx), of Limbu (Limb).
#### func HasScxLina(r rune) bool
HasScxLina returns true if code point, r, has Unicode property, Script_Extensions (scx), of Linear_A (Lina).
#### func HasScxLinb(r rune) bool
HasScxLinb returns true if code point, r, has Unicode property, Script_Extensions (scx), of Linear_B (Linb).
#### func HasScxLisu(r rune) bool
HasScxLisu returns true if code point, r, has Unicode property, Script_Extensions (scx), of Lisu.
#### func HasScxLyci(r rune) bool
HasScxLyci returns true if code point, r, has Unicode property, Script_Extensions (scx), of Lycian (Lyci).
#### func HasScxLydi(r rune) bool
HasScxLydi returns true if code point, r, has Unicode property, Script_Extensions (scx), of Lydian (Lydi).
#### func HasScxMahj(r rune) bool
HasScxMahj returns true if code point, r, has Unicode property, Script_Extensions (scx), of Mahajani (Mahj).
#### func HasScxMaka(r rune) bool
HasScxMaka returns true if code point, r, has Unicode property, Script_Extensions (scx), of Makasar (Maka).
#### func HasScxMand(r rune) bool
HasScxMand returns true if code point, r, has Unicode property, Script_Extensions (scx), of Mandaic (Mand).
#### func HasScxMani(r rune) bool
HasScxMani returns true if code point, r, has Unicode property, Script_Extensions (scx), of Manichaean (Mani).
#### func HasScxMarc(r rune) bool
HasScxMarc returns true if code point, r, has Unicode property, Script_Extensions (scx), of Marchen (Marc).
#### func HasScxMedf(r rune) bool
HasScxMedf returns true if code point, r, has Unicode property, Script_Extensions (scx), of Medefaidrin (Medf).
#### func HasScxMend(r rune) bool
HasScxMend returns true if code point, r, has Unicode property, Script_Extensions (scx), of Mende_Kikakui (Mend).
#### func HasScxMerc(r rune) bool
HasScxMerc returns true if code point, r, has Unicode property, Script_Extensions (scx), of Meroitic_Cursive (Merc).
#### func HasScxMero(r rune) bool
HasScxMero returns true if code point, r, has Unicode property, Script_Extensions (scx), of Meroitic_Hieroglyphs (Mero).
#### func HasScxMlym(r rune) bool
HasScxMlym returns true if code point, r, has Unicode property, Script_Extensions (scx), of Malayalam (Mlym).
#### func HasScxModi(r rune) bool
HasScxModi returns true if code point, r, has Unicode property, Script_Extensions (scx), of Modi.
#### func HasScxMong(r rune) bool
HasScxMong returns true if code point, r, has Unicode property, Script_Extensions (scx), of Mongolian (Mong).
#### func HasScxMroo(r rune) bool
HasScxMroo returns true if code point, r, has Unicode property, Script_Extensions (scx), of Mro (Mroo).
#### func HasScxMtei(r rune) bool
HasScxMtei returns true if code point, r, has Unicode property, Script_Extensions (scx), of Meetei_Mayek (Mtei).
#### func HasScxMult(r rune) bool
HasScxMult returns true if code point, r, has Unicode property, Script_Extensions (scx), of Multani (Mult).
#### func HasScxMymr(r rune) bool
HasScxMymr returns true if code point, r, has Unicode property, Script_Extensions (scx), of Myanmar (Mymr).
#### func HasScxNand(r rune) bool
HasScxNand returns true if code point, r, has Unicode property, Script_Extensions (scx), of Nandinagari (Nand).
#### func HasScxNarb(r rune) bool
HasScxNarb returns true if code point, r, has Unicode property, Script_Extensions (scx), of Old_North_Arabian (Narb).
#### func HasScxNbat(r rune) bool
HasScxNbat returns true if code point, r, has Unicode property, Script_Extensions (scx), of Nabataean (Nbat).
#### func HasScxNewa(r rune) bool
HasScxNewa returns true if code point, r, has Unicode property, Script_Extensions (scx), of Newa.
#### func HasScxNkoo(r rune) bool
HasScxNkoo returns true if code point, r, has Unicode property, Script_Extensions (scx), of Nko (Nkoo).
#### func HasScxNshu(r rune) bool
HasScxNshu returns true if code point, r, has Unicode property, Script_Extensions (scx), of Nushu (Nshu).
#### func HasScxOgam(r rune) bool
HasScxOgam returns true if code point, r, has Unicode property, Script_Extensions (scx), of Ogham (Ogam).
#### func HasScxOlck(r rune) bool
HasScxOlck returns true if code point, r, has Unicode property, Script_Extensions (scx), of Ol_Chiki (Olck).
#### func HasScxOrkh(r rune) bool
HasScxOrkh returns true if code point, r, has Unicode property, Script_Extensions (scx), of Old_Turkic (Orkh).
#### func HasScxOrya(r rune) bool
HasScxOrya returns true if code point, r, has Unicode property, Script_Extensions (scx), of Oriya (Orya).
#### func HasScxOsge(r rune) bool
HasScxOsge returns true if code point, r, has Unicode property, Script_Extensions (scx), of Osage (Osge).
#### func HasScxOsma(r rune) bool
HasScxOsma returns true if code point, r, has Unicode property, Script_Extensions (scx), of Osmanya (Osma).
#### func HasScxOugr(r rune) bool
HasScxOugr returns true if code point, r, has Unicode property, Script_Extensions (scx), of Old_Uyghur (Ougr).
#### func HasScxPalm(r rune) bool
HasScxPalm returns true if code point, r, has Unicode property, Script_Extensions (scx), of Palmyrene (Palm).
#### func HasScxPauc(r rune) bool
HasScxPauc returns true if code point, r, has Unicode property, Script_Extensions (scx), of Pau_Cin_Hau (Pauc).
#### func HasScxPerm(r rune) bool
HasScxPerm returns true if code point, r, has Unicode property, Script_Extensions (scx), of Old_Permic (Perm).
#### func HasScxPhag(r rune) bool
HasScxPhag returns true if code point, r, has Unicode property, Script_Extensions (scx), of Phags_Pa (Phag).
#### func HasScxPhli(r rune) bool
HasScxPhli returns true if code point, r, has Unicode property, Script_Extensions (scx), of Inscriptional_Pahlavi (Phli).
#### func HasScxPhlp(r rune) bool
HasScxPhlp returns true if code point, r, has Unicode property, Script_Extensions (scx), of Psalter_Pahlavi (Phlp).
#### func HasScxPhnx(r rune) bool
HasScxPhnx returns true if code point, r, has Unicode property, Script_Extensions (scx), of Phoenician (Phnx).
#### func HasScxPlrd(r rune) bool
HasScxPlrd returns true if code point, r, has Unicode property, Script_Extensions (scx), of Miao (Plrd).
#### func HasScxPrti(r rune) bool
HasScxPrti returns true if code point, r, has Unicode property, Script_Extensions (scx), of Inscriptional_Parthian (Prti).
#### func HasScxRjng(r rune) bool
HasScxRjng returns true if code point, r, has Unicode property, Script_Extensions (scx), of Rejang (Rjng).
#### func HasScxRohg(r rune) bool
HasScxRohg returns true if code point, r, has Unicode property, Script_Extensions (scx), of Hanifi_Rohingya (Rohg).
#### func HasScxRunr(r rune) bool
HasScxRunr returns true if code point, r, has Unicode property, Script_Extensions (scx), of Runic (Runr).
#### func HasScxSamr(r rune) bool
HasScxSamr returns true if code point, r, has Unicode property, Script_Extensions (scx), of Samaritan (Samr).
#### func HasScxSarb(r rune) bool
HasScxSarb returns true if code point, r, has Unicode property, Script_Extensions (scx), of Old_South_Arabian (Sarb).
#### func HasScxSaur(r rune) bool
HasScxSaur returns true if code point, r, has Unicode property, Script_Extensions (scx), of Saurashtra (Saur).
#### func HasScxSgnw(r rune) bool
HasScxSgnw returns true if code point, r, has Unicode property, Script_Extensions (scx), of SignWriting (Sgnw).
#### func HasScxShaw(r rune) bool
HasScxShaw returns true if code point, r, has Unicode property, Script_Extensions (scx), of Shavian (Shaw).
#### func HasScxShrd(r rune) bool
HasScxShrd returns true if code point, r, has Unicode property, Script_Extensions (scx), of Sharada (Shrd).
#### func HasScxSidd(r rune) bool
HasScxSidd returns true if code point, r, has Unicode property, Script_Extensions (scx), of Siddham (Sidd).
#### func HasScxSind(r rune) bool
HasScxSind returns true if code point, r, has Unicode property, Script_Extensions (scx), of Khudawadi (Sind).
#### func HasScxSinh(r rune) bool
HasScxSinh returns true if code point, r, has Unicode property, Script_Extensions (scx), of Sinhala (Sinh).
#### func HasScxSogd(r rune) bool
HasScxSogd returns true if code point, r, has Unicode property, Script_Extensions (scx), of Sogdian (Sogd).
#### func HasScxSogo(r rune) bool
HasScxSogo returns true if code point, r, has Unicode property, Script_Extensions (scx), of Old_Sogdian (Sogo).
#### func HasScxSora(r rune) bool
HasScxSora returns true if code point, r, has Unicode property, Script_Extensions (scx), of Sora_Sompeng (Sora).
#### func HasScxSoyo(r rune) bool
HasScxSoyo returns true if code point, r, has Unicode property, Script_Extensions (scx), of Soyombo (Soyo).
#### func HasScxSund(r rune) bool
HasScxSund returns true if code point, r, has Unicode property, Script_Extensions (scx), of Sundanese (Sund).
#### func HasScxSylo(r rune) bool
HasScxSylo returns true if code point, r, has Unicode property, Script_Extensions (scx), of Syloti_Nagri (Sylo).
#### func HasScxSyrc(r rune) bool
HasScxSyrc returns true if code point, r, has Unicode property, Script_Extensions (scx), of Syriac (Syrc).
#### func HasScxTagb(r rune) bool
HasScxTagb returns true if code point, r, has Unicode property, Script_Extensions (scx), of Tagbanwa (Tagb).
#### func HasScxTakr(r rune) bool
HasScxTakr returns true if code point, r, has Unicode property, Script_Extensions (scx), of Takri (Takr).
#### func HasScxTale(r rune) bool
HasScxTale returns true if code point, r, has Unicode property, Script_Extensions (scx), of Tai_Le (Tale).
#### func HasScxTalu(r rune) bool
HasScxTalu returns true if code point, r, has Unicode property, Script_Extensions (scx), of New_Tai_Lue (Talu).
#### func HasScxTaml(r rune) bool
HasScxTaml returns true if code point, r, has Unicode property, Script_Extensions (scx), of Tamil (Taml).
#### func HasScxTang(r rune) bool
HasScxTang returns true if code point, r, has Unicode property, Script_Extensions (scx), of Tangut (Tang).
#### func HasScxTavt(r rune) bool
HasScxTavt returns true if code point, r, has Unicode property, Script_Extensions (scx), of Tai_Viet (Tavt).
#### func HasScxTelu(r rune) bool
HasScxTelu returns true if code point, r, has Unicode property, Script_Extensions (scx), of Telugu (Telu).
#### func HasScxTfng(r rune) bool
HasScxTfng returns true if code point, r, has Unicode property, Script_Extensions (scx), of Tifinagh (Tfng).
#### func HasScxTglg(r rune) bool
HasScxTglg returns true if code point, r, has Unicode property, Script_Extensions (scx), of Tagalog (Tglg).
#### func HasScxThaa(r rune) bool
HasScxThaa returns true if code point, r, has Unicode property, Script_Extensions (scx), of Thaana (Thaa).
#### func HasScxThai(r rune) bool
HasScxThai returns true if code point, r, has Unicode property, Script_Extensions (scx), of Thai.
#### func HasScxTibt(r rune) bool
HasScxTibt returns true if code point, r, has Unicode property, Script_Extensions (scx), of Tibetan (Tibt).
#### func HasScxTirh(r rune) bool
HasScxTirh returns true if code point, r, has Unicode property, Script_Extensions (scx), of Tirhuta (Tirh).
#### func HasScxTnsa(r rune) bool
HasScxTnsa returns true if code point, r, has Unicode property, Script_Extensions (scx), of Tangsa (Tnsa).
#### func HasScxToto(r rune) bool
HasScxToto returns true if code point, r, has Unicode property, Script_Extensions (scx), of Toto.
#### func HasScxUgar(r rune) bool
HasScxUgar returns true if code point, r, has Unicode property, Script_Extensions (scx), of Ugaritic (Ugar).
#### func HasScxVaii(r rune) bool
HasScxVaii returns true if code point, r, has Unicode property, Script_Extensions (scx), of Vai (Vaii).
#### func HasScxVith(r rune) bool
HasScxVith returns true if code point, r, has Unicode property, Script_Extensions (scx), of Vithkuqi (Vith).
#### func HasScxWara(r rune) bool
HasScxWara returns true if code point, r, has Unicode property, Script_Extensions (scx), of Warang_Citi (Wara).
#### func HasScxWcho(r rune) bool
HasScxWcho returns true if code point, r, has Unicode property, Script_Extensions (scx), of Wancho (Wcho).
#### func HasScxXpeo(r rune) bool
HasScxXpeo returns true if code point, r, has Unicode property, Script_Extensions (scx), of Old_Persian (Xpeo).
#### func HasScxXsux(r rune) bool
HasScxXsux returns true if code point, r, has Unicode property, Script_Extensions (scx), of Cuneiform (Xsux).
#### func HasScxYezi(r rune) bool
HasScxYezi returns true if code point, r, has Unicode property, Script_Extensions (scx), of Yezidi (Yezi).
#### func HasScxYiii(r rune) bool
HasScxYiii returns true if code point, r, has Unicode property, Script_Extensions (scx), of Yi (Yiii).
#### func HasScxZanb(r rune) bool
HasScxZanb returns true if code point, r, has Unicode property, Script_Extensions (scx), of Zanabazar_Square (Zanb).
#### func HasScxZinh(r rune) bool
HasScxZinh returns true if code point, r, has Unicode property, Script_Extensions (scx), of Inherited (Zinh).
#### func HasScxZyyy(r rune) bool
HasScxZyyy returns true if code point, r, has Unicode property, Script_Extensions (scx), of Common (Zyyy).
#### func HasScxZzzz(r rune) bool
HasScxZzzz returns true if code point, r, has Unicode property, Script_Extensions (scx), of Unknown (Zzzz).
### Other Functions
## Supported Unicode Properties
<table><thead><tr><th>Property:</th><th>Value:</th><th>Function:</th></tr></thead>
<tbody>
<tr><td colspan="3">Script_Extensions (scx)</td></tr>
<tr><td>&nbsp;</td><td>Adlam (Adlm)</td><td>`HasScxAdlm(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Caucasian_Albanian (Aghb)</td><td>`HasScxAghb(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Ahom</td><td>`HasScxAhom(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Arabic (Arab)</td><td>`HasScxArab(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Imperial_Aramaic (Armi)</td><td>`HasScxArmi(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Armenian (Armn)</td><td>`HasScxArmn(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Avestan (Avst)</td><td>`HasScxAvst(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Balinese (Bali)</td><td>`HasScxBali(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Bamum (Bamu)</td><td>`HasScxBamu(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Bassa_Vah (Bass)</td><td>`HasScxBass(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Batak (Batk)</td><td>`HasScxBatk(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Bengali (Beng)</td><td>`HasScxBeng(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Bhaiksuki (Bhks)</td><td>`HasScxBhks(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Bopomofo (Bopo)</td><td>`HasScxBopo(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Brahmi (Brah)</td><td>`HasScxBrah(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Braille (Brai)</td><td>`HasScxBrai(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Buginese (Bugi)</td><td>`HasScxBugi(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Buhid (Buhd)</td><td>`HasScxBuhd(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Chakma (Cakm)</td><td>`HasScxCakm(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Canadian_Aboriginal (Cans)</td><td>`HasScxCans(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Carian (Cari)</td><td>`HasScxCari(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Cham</td><td>`HasScxCham(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Cherokee (Cher)</td><td>`HasScxCher(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Chorasmian (Chrs)</td><td>`HasScxChrs(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Coptic (Copt, Qaac)</td><td>`HasScxCopt(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Cypro_Minoan (Cpmn)</td><td>`HasScxCpmn(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Cypriot (Cprt)</td><td>`HasScxCprt(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Cyrillic (Cyrl)</td><td>`HasScxCyrl(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Devanagari (Deva)</td><td>`HasScxDeva(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Dives_Akuru (Diak)</td><td>`HasScxDiak(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Dogra (Dogr)</td><td>`HasScxDogr(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Deseret (Dsrt)</td><td>`HasScxDsrt(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Duployan (Dupl)</td><td>`HasScxDupl(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Egyptian_Hieroglyphs (Egyp)</td><td>`HasScxEgyp(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Elbasan (Elba)</td><td>`HasScxElba(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Elymaic (Elym)</td><td>`HasScxElym(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Ethiopic (Ethi)</td><td>`HasScxEthi(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Georgian (Geor)</td><td>`HasScxGeor(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Glagolitic (Glag)</td><td>`HasScxGlag(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Gunjala_Gondi (Gong)</td><td>`HasScxGong(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Masaram_Gondi (Gonm)</td><td>`HasScxGonm(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Gothic (Goth)</td><td>`HasScxGoth(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Grantha (Gran)</td><td>`HasScxGran(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Greek (Grek)</td><td>`HasScxGrek(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Gujarati (Gujr)</td><td>`HasScxGujr(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Gurmukhi (Guru)</td><td>`HasScxGuru(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Hangul (Hang)</td><td>`HasScxHang(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Han (Hani)</td><td>`HasScxHani(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Hanunoo (Hano)</td><td>`HasScxHano(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Hatran (Hatr)</td><td>`HasScxHatr(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Hebrew (Hebr)</td><td>`HasScxHebr(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Hiragana (Hira)</td><td>`HasScxHira(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Anatolian_Hieroglyphs (Hluw)</td><td>`HasScxHluw(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Pahawh_Hmong (Hmng)</td><td>`HasScxHmng(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Nyiakeng_Puachue_Hmong (Hmnp)</td><td>`HasScxHmnp(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Old_Hungarian (Hung)</td><td>`HasScxHung(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Old_Italic (Ital)</td><td>`HasScxItal(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Javanese (Java)</td><td>`HasScxJava(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Kayah_Li (Kali)</td><td>`HasScxKali(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Katakana (Kana)</td><td>`HasScxKana(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Kharoshthi (Khar)</td><td>`HasScxKhar(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Khmer (Khmr)</td><td>`HasScxKhmr(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Khojki (Khoj)</td><td>`HasScxKhoj(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Khitan_Small_Script (Kits)</td><td>`HasScxKits(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Kannada (Knda)</td><td>`HasScxKnda(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Kaithi (Kthi)</td><td>`HasScxKthi(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Tai_Tham (Lana)</td><td>`HasScxLana(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Lao (Laoo)</td><td>`HasScxLaoo(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Latin (Latn)</td><td>`HasScxLatn(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Lepcha (Lepc)</td><td>`HasScxLepc(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Limbu (Limb)</td><td>`HasScxLimb(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Linear_A (Lina)</td><td>`HasScxLina(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Linear_B (Linb)</td><td>`HasScxLinb(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Lisu</td><td>`HasScxLisu(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Lycian (Lyci)</td><td>`HasScxLyci(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Lydian (Lydi)</td><td>`HasScxLydi(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Mahajani (Mahj)</td><td>`HasScxMahj(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Makasar (Maka)</td><td>`HasScxMaka(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Mandaic (Mand)</td><td>`HasScxMand(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Manichaean (Mani)</td><td>`HasScxMani(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Marchen (Marc)</td><td>`HasScxMarc(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Medefaidrin (Medf)</td><td>`HasScxMedf(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Mende_Kikakui (Mend)</td><td>`HasScxMend(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Meroitic_Cursive (Merc)</td><td>`HasScxMerc(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Meroitic_Hieroglyphs (Mero)</td><td>`HasScxMero(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Malayalam (Mlym)</td><td>`HasScxMlym(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Modi</td><td>`HasScxModi(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Mongolian (Mong)</td><td>`HasScxMong(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Mro (Mroo)</td><td>`HasScxMroo(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Meetei_Mayek (Mtei)</td><td>`HasScxMtei(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Multani (Mult)</td><td>`HasScxMult(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Myanmar (Mymr)</td><td>`HasScxMymr(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Nandinagari (Nand)</td><td>`HasScxNand(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Old_North_Arabian (Narb)</td><td>`HasScxNarb(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Nabataean (Nbat)</td><td>`HasScxNbat(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Newa</td><td>`HasScxNewa(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Nko (Nkoo)</td><td>`HasScxNkoo(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Nushu (Nshu)</td><td>`HasScxNshu(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Ogham (Ogam)</td><td>`HasScxOgam(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Ol_Chiki (Olck)</td><td>`HasScxOlck(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Old_Turkic (Orkh)</td><td>`HasScxOrkh(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Oriya (Orya)</td><td>`HasScxOrya(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Osage (Osge)</td><td>`HasScxOsge(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Osmanya (Osma)</td><td>`HasScxOsma(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Old_Uyghur (Ougr)</td><td>`HasScxOugr(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Palmyrene (Palm)</td><td>`HasScxPalm(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Pau_Cin_Hau (Pauc)</td><td>`HasScxPauc(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Old_Permic (Perm)</td><td>`HasScxPerm(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Phags_Pa (Phag)</td><td>`HasScxPhag(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Inscriptional_Pahlavi (Phli)</td><td>`HasScxPhli(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Psalter_Pahlavi (Phlp)</td><td>`HasScxPhlp(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Phoenician (Phnx)</td><td>`HasScxPhnx(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Miao (Plrd)</td><td>`HasScxPlrd(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Inscriptional_Parthian (Prti)</td><td>`HasScxPrti(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Rejang (Rjng)</td><td>`HasScxRjng(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Hanifi_Rohingya (Rohg)</td><td>`HasScxRohg(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Runic (Runr)</td><td>`HasScxRunr(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Samaritan (Samr)</td><td>`HasScxSamr(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Old_South_Arabian (Sarb)</td><td>`HasScxSarb(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Saurashtra (Saur)</td><td>`HasScxSaur(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>SignWriting (Sgnw)</td><td>`HasScxSgnw(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Shavian (Shaw)</td><td>`HasScxShaw(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Sharada (Shrd)</td><td>`HasScxShrd(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Siddham (Sidd)</td><td>`HasScxSidd(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Khudawadi (Sind)</td><td>`HasScxSind(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Sinhala (Sinh)</td><td>`HasScxSinh(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Sogdian (Sogd)</td><td>`HasScxSogd(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Old_Sogdian (Sogo)</td><td>`HasScxSogo(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Sora_Sompeng (Sora)</td><td>`HasScxSora(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Soyombo (Soyo)</td><td>`HasScxSoyo(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Sundanese (Sund)</td><td>`HasScxSund(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Syloti_Nagri (Sylo)</td><td>`HasScxSylo(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Syriac (Syrc)</td><td>`HasScxSyrc(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Tagbanwa (Tagb)</td><td>`HasScxTagb(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Takri (Takr)</td><td>`HasScxTakr(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Tai_Le (Tale)</td><td>`HasScxTale(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>New_Tai_Lue (Talu)</td><td>`HasScxTalu(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Tamil (Taml)</td><td>`HasScxTaml(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Tangut (Tang)</td><td>`HasScxTang(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Tai_Viet (Tavt)</td><td>`HasScxTavt(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Telugu (Telu)</td><td>`HasScxTelu(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Tifinagh (Tfng)</td><td>`HasScxTfng(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Tagalog (Tglg)</td><td>`HasScxTglg(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Thaana (Thaa)</td><td>`HasScxThaa(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Thai</td><td>`HasScxThai(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Tibetan (Tibt)</td><td>`HasScxTibt(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Tirhuta (Tirh)</td><td>`HasScxTirh(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Tangsa (Tnsa)</td><td>`HasScxTnsa(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Toto</td><td>`HasScxToto(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Ugaritic (Ugar)</td><td>`HasScxUgar(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Vai (Vaii)</td><td>`HasScxVaii(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Vithkuqi (Vith)</td><td>`HasScxVith(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Warang_Citi (Wara)</td><td>`HasScxWara(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Wancho (Wcho)</td><td>`HasScxWcho(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Old_Persian (Xpeo)</td><td>`HasScxXpeo(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Cuneiform (Xsux)</td><td>`HasScxXsux(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Yezidi (Yezi)</td><td>`HasScxYezi(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Yi (Yiii)</td><td>`HasScxYiii(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Zanabazar_Square (Zanb)</td><td>`HasScxZanb(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Inherited (Zinh, Qaai)</td><td>`HasScxZinh(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Common (Zyyy)</td><td>`HasScxZyyy(r rune) bool`</td></tr>
<tr><td>&nbsp;</td><td>Unknown (Zzzz)</td><td>`HasScxZzzz(r rune) bool`</td></tr>
</tbody></table>
