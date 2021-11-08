<h1 align="center">Unicode​® queries of common properties (e.g. General_Category, White_Space, etc.) and extended grapheme cluster parsing</h1>

**ubasic** is a Go package for querying common Unicode properties. Other Unicode properties are supported by other packages in this module.

For more information about the properties implemented herein, consult [UAX #44](https://www.unicode.org/reports/tr44/).

For example, the function, `HasGcLu(rune r) bool`, returns true only if the code point, r, has Unicode property, General_Category, with value, Uppercase_Letter (Lu), often referred to in regular expressions as "/p{Lu}".

Many of the Unicode property queries of this package are also available in standard Go packages, unicode and regexp. For example, `HasGcZ(rune r) bool` is analogous to `unicode.Is(unicode.Z, r) bool` or the regexp class `\pZ`. If you are already using regular expressions on the provided text, it may be preferrable to just use regexp. Otherwise, some benefits of using this class include the following. These benefits apply at the time of this writing, but may change by the time you read this.

- Some properties and values are available here that are not in the packages, unicode and regexp, such as General_Category values Cn and LC.
- This package appears to perform about twice as fast as the unicode package (slightly more than 2X for the ASCII character range while slightly less than 2X for the full Unicode character range) and about 55X-85X faster than using regexp. This performance comparison was measured by running `go test -bench` on a platform. Other platforms and software versions may produce different results. It also appears there may be room for additional improvements in subsequent revisions of this module.

```sh
BenchmarkUBasic-8            	     663	   1785613 ns/op
BenchmarkUBasicAscii-8       	  189930	      6264 ns/op
BenchmarkPkgUnicode-8        	     374	   3250854 ns/op
BenchmarkPkgUnicodeAscii-8   	   79880	     14949 ns/op
BenchmarkPkgRegexp-8         	      10	 100037250 ns/op
BenchmarkPkgRegexpAscii-8    	    2217	    530573 ns/op
```

For a summary of Unicode properties and which functions of which packages of this module provide the corresponding support, peruse the file, **SUMMARY.txt**, in the top level directory of this module.

This package also provides parsing of extended grapheme cluster boundaries. In many cases, especially dealing with various languages and emojis, what we think of as a "character" may span multiple code points. These code points provide attributes that contribute to the meaning of the character. Text should be processed at extended grapheme cluster boundaries to prevent corruption. For more details about extended grapheme cluster boundaries, consult [UAX #29](https://www.unicode.org/reports/tr29/).

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

import "github.com/orkvozku/go/uni/ubasic"
import "fmt"

func main() {
    text := "Hi!"
    for _, r := range text {
        if ubasic.HasGcLu(r) {
            fmt.Printf("'%c' is an uppercase letter\n", r)
        } else if ubasic.HasGcLl(r) {
            fmt.Printf("'%c' is a lowercase letter\n", r)
        } else if ubasic.HasGcP(r) {
            fmt.Printf("'%c' is a punctuation character\n", r)
        }
    }

    // output:
    // 'H' is an uppercase letter
    // 'i' is a lowercase letter
    // '!' is a punctuation character
}
```
## API
### General Constants
#### UnicodeVersion
UnicodeVersion is the current version of the Unicode specification that this package is intended to comply with.

UnicodeVersion = "14.0.0".
### Unicode Property Queries
#### func HasGcC(r rune) bool
HasGcC returns true if code point, r, has Unicode property, General_Category (gc), of Other (C).
#### func HasGcCc(r rune) bool
HasGcCc returns true if code point, r, has Unicode property, General_Category (gc), of Control (Cc).
#### func HasGcCf(r rune) bool
HasGcCf returns true if code point, r, has Unicode property, General_Category (gc), of Format (Cf).
#### func HasGcCn(r rune) bool
HasGcCn returns true if code point, r, has Unicode property, General_Category (gc), of Unassigned (Cn).
#### func HasGcCo(r rune) bool
HasGcCo returns true if code point, r, has Unicode property, General_Category (gc), of Private_Use (Co).
#### func HasGcCs(r rune) bool
HasGcCs returns true if code point, r, has Unicode property, General_Category (gc), of Surrogate (Cs).
#### func HasGcL(r rune) bool
HasGcL returns true if code point, r, has Unicode property, General_Category (gc), of Letter (L).
#### func HasGcLc(r rune) bool
HasGcLc returns true if code point, r, has Unicode property, General_Category (gc), of Cased_Letter (LC).
#### func HasGcLl(r rune) bool
HasGcLl returns true if code point, r, has Unicode property, General_Category (gc), of Lowercase_Letter (Ll).
#### func HasGcLm(r rune) bool
HasGcLm returns true if code point, r, has Unicode property, General_Category (gc), of Modifier_Letter (Lm).
#### func HasGcLo(r rune) bool
HasGcLo returns true if code point, r, has Unicode property, General_Category (gc), of Other_Letter (Lo).
#### func HasGcLt(r rune) bool
HasGcLt returns true if code point, r, has Unicode property, General_Category (gc), of Titlecase_Letter (Lt).
#### func HasGcLu(r rune) bool
HasGcLu returns true if code point, r, has Unicode property, General_Category (gc), of Uppercase_Letter (Lu).
#### func HasGcM(r rune) bool
HasGcM returns true if code point, r, has Unicode property, General_Category (gc), of Mark (M).
#### func HasGcMc(r rune) bool
HasGcMc returns true if code point, r, has Unicode property, General_Category (gc), of Spacing_Mark (Mc).
#### func HasGcMe(r rune) bool
HasGcMe returns true if code point, r, has Unicode property, General_Category (gc), of Enclosing_Mark (Me).
#### func HasGcMn(r rune) bool
HasGcMn returns true if code point, r, has Unicode property, General_Category (gc), of Nonspacing_Mark (Mn).
#### func HasGcN(r rune) bool
HasGcN returns true if code point, r, has Unicode property, General_Category (gc), of Number (N).
#### func HasGcNd(r rune) bool
HasGcNd returns true if code point, r, has Unicode property, General_Category (gc), of Decimal_Number (Nd).
#### func HasGcNl(r rune) bool
HasGcNl returns true if code point, r, has Unicode property, General_Category (gc), of Letter_Number (Nl).
#### func HasGcNo(r rune) bool
HasGcNo returns true if code point, r, has Unicode property, General_Category (gc), of Other_Number (No).
#### func HasGcP(r rune) bool
HasGcP returns true if code point, r, has Unicode property, General_Category (gc), of Punctuation (P).
#### func HasGcPc(r rune) bool
HasGcPc returns true if code point, r, has Unicode property, General_Category (gc), of Connector_Punctuation (Pc).
#### func HasGcPd(r rune) bool
HasGcPd returns true if code point, r, has Unicode property, General_Category (gc), of Dash_Punctuation (Pd).
#### func HasGcPe(r rune) bool
HasGcPe returns true if code point, r, has Unicode property, General_Category (gc), of Close_Punctuation (Pe).
#### func HasGcPf(r rune) bool
HasGcPf returns true if code point, r, has Unicode property, General_Category (gc), of Final_Punctuation (Pf).
#### func HasGcPi(r rune) bool
HasGcPi returns true if code point, r, has Unicode property, General_Category (gc), of Initial_Punctuation (Pi).
#### func HasGcPo(r rune) bool
HasGcPo returns true if code point, r, has Unicode property, General_Category (gc), of Other_Punctuation (Po).
#### func HasGcPs(r rune) bool
HasGcPs returns true if code point, r, has Unicode property, General_Category (gc), of Open_Punctuation (Ps).
#### func HasGcS(r rune) bool
HasGcS returns true if code point, r, has Unicode property, General_Category (gc), of Symbol (S).
#### func HasGcSc(r rune) bool
HasGcSc returns true if code point, r, has Unicode property, General_Category (gc), of Currency_Symbol (Sc).
#### func HasGcSk(r rune) bool
HasGcSk returns true if code point, r, has Unicode property, General_Category (gc), of Modifier_Symbol (Sk).
#### func HasGcSm(r rune) bool
HasGcSm returns true if code point, r, has Unicode property, General_Category (gc), of Math_Symbol (Sm).
#### func HasGcSo(r rune) bool
HasGcSo returns true if code point, r, has Unicode property, General_Category (gc), of Other_Symbol (So).
#### func HasGcZ(r rune) bool
HasGcZ returns true if code point, r, has Unicode property, General_Category (gc), of Separator (Z).
#### func HasGcZl(r rune) bool
HasGcZl returns true if code point, r, has Unicode property, General_Category (gc), of Line_Separator (Zl).
#### func HasGcZp(r rune) bool
HasGcZp returns true if code point, r, has Unicode property, General_Category (gc), of Paragraph_Separator (Zp).
#### func HasGcZs(r rune) bool
HasGcZs returns true if code point, r, has Unicode property, General_Category (gc), of Space_Separator (Zs).
#### func HasWspaceN(r rune) bool
HasWspaceN returns true if code point, r, has Unicode property, White_Space (WSpace), of No (N).
#### func HasWspaceY(r rune) bool
HasWspaceY returns true if code point, r, has Unicode property, White_Space (WSpace), of Yes (Y).
#### func HasGcbCn(r rune) bool
HasGcbCn returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of Control (CN).
#### func HasGcbCr(r rune) bool
HasGcbCr returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of CR.
#### func HasGcbEx(r rune) bool
HasGcbEx returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of Extend (EX).
#### func HasGcbL(r rune) bool
HasGcbL returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of L.
#### func HasGcbLf(r rune) bool
HasGcbLf returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of LF.
#### func HasGcbLv(r rune) bool
HasGcbLv returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of LV.
#### func HasGcbLvt(r rune) bool
HasGcbLvt returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of LVT.
#### func HasGcbPp(r rune) bool
HasGcbPp returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of Prepend (PP).
#### func HasGcbRi(r rune) bool
HasGcbRi returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of Regional_Indicator (RI).
#### func HasGcbSm(r rune) bool
HasGcbSm returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of SpacingMark (SM).
#### func HasGcbT(r rune) bool
HasGcbT returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of T.
#### func HasGcbV(r rune) bool
HasGcbV returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of V.
#### func HasGcbXx(r rune) bool
HasGcbXx returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of Other (XX).
#### func HasGcbZwj(r rune) bool
HasGcbZwj returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of ZWJ.
#### func HasExtpictN(r rune) bool
HasExtpictN returns true if code point, r, has Unicode property, Extended_Pictographic (ExtPict), of No (N).
#### func HasExtpictY(r rune) bool
HasExtpictY returns true if code point, r, has Unicode property, Extended_Pictographic (ExtPict), of Yes (Y).
### Other Functions
## Supported Unicode Properties
<table><thead><tr><th>Property:</th><th>Value:</th><th>Function:</th></tr></thead>
<tbody>
<tr><td colspan="3">General_Category (gc)</td></tr>
<tr><td>&nbsp;</td><td>Other (C)</td><td>HasGcC(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Control (Cc, cntrl)</td><td>HasGcCc(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Format (Cf)</td><td>HasGcCf(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Unassigned (Cn)</td><td>HasGcCn(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Private_Use (Co)</td><td>HasGcCo(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Surrogate (Cs)</td><td>HasGcCs(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Letter (L)</td><td>HasGcL(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Cased_Letter (LC)</td><td>HasGcLc(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Lowercase_Letter (Ll)</td><td>HasGcLl(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Modifier_Letter (Lm)</td><td>HasGcLm(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Other_Letter (Lo)</td><td>HasGcLo(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Titlecase_Letter (Lt)</td><td>HasGcLt(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Uppercase_Letter (Lu)</td><td>HasGcLu(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Mark (M, Combining_Mark)</td><td>HasGcM(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Spacing_Mark (Mc)</td><td>HasGcMc(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Enclosing_Mark (Me)</td><td>HasGcMe(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Nonspacing_Mark (Mn)</td><td>HasGcMn(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Number (N)</td><td>HasGcN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Decimal_Number (Nd, digit)</td><td>HasGcNd(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Letter_Number (Nl)</td><td>HasGcNl(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Other_Number (No)</td><td>HasGcNo(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Punctuation (P, punct)</td><td>HasGcP(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Connector_Punctuation (Pc)</td><td>HasGcPc(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Dash_Punctuation (Pd)</td><td>HasGcPd(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Close_Punctuation (Pe)</td><td>HasGcPe(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Final_Punctuation (Pf)</td><td>HasGcPf(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Initial_Punctuation (Pi)</td><td>HasGcPi(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Other_Punctuation (Po)</td><td>HasGcPo(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Open_Punctuation (Ps)</td><td>HasGcPs(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Symbol (S)</td><td>HasGcS(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Currency_Symbol (Sc)</td><td>HasGcSc(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Modifier_Symbol (Sk)</td><td>HasGcSk(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Math_Symbol (Sm)</td><td>HasGcSm(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Other_Symbol (So)</td><td>HasGcSo(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Separator (Z)</td><td>HasGcZ(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Line_Separator (Zl)</td><td>HasGcZl(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Paragraph_Separator (Zp)</td><td>HasGcZp(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Space_Separator (Zs)</td><td>HasGcZs(r rune) bool</td></tr>
<tr><td colspan="3">White_Space (WSpace, space)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasWspaceN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasWspaceY(r rune) bool</td></tr>
<tr><td colspan="3">Grapheme_Cluster_Break (GCB)</td></tr>
<tr><td>&nbsp;</td><td>Control (CN)</td><td>HasGcbCn(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>CR</td><td>HasGcbCr(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Extend (EX)</td><td>HasGcbEx(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>L</td><td>HasGcbL(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>LF</td><td>HasGcbLf(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>LV</td><td>HasGcbLv(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>LVT</td><td>HasGcbLvt(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Prepend (PP)</td><td>HasGcbPp(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Regional_Indicator (RI)</td><td>HasGcbRi(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>SpacingMark (SM)</td><td>HasGcbSm(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>T</td><td>HasGcbT(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>V</td><td>HasGcbV(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Other (XX)</td><td>HasGcbXx(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>ZWJ</td><td>HasGcbZwj(r rune) bool</td></tr>
<tr><td colspan="3">Extended_Pictographic (ExtPict)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasExtpictN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasExtpictY(r rune) bool</td></tr>
</tbody></table>
## Performance
Running the `go test -bench` command on this package produces performance benchmarks as follows.

- **BenchmarkUBasic** - Runs through a variety of General_Category tests on a large swath of Unicode code points.
- **BenchmarkPkgUnicode** - Runs the same tests as **BenchmarkUcdCommon** except using `unicode.Is` instead.
- **BenchmarkPkgRegexp** - Runs the same tests as **BenchmarkUcdCommon** except using pre-compiled regular expressions corresponding to the same tests, such as `\p{Lu}`.
- **BenchmarkUBasicAscii** - Same as **BenchmarkUBasic** except tests are limited to the ASCII character range, [0x00, 0x7F].
- **BenchmarkPkgUnicodeAscii** - Same as **BenchmarkPkgUnicode** except tests are limited to the ASCII character range, [0x00, 0x7F].
- **BenchmarkPkgRegexpAscii** - Same as **BenchmarkPkgRegexp** except tests are limited to the ASCII character range, [0x00, 0x7F].

The test names are as follows.

- Names containing **UBasic** - perform the tests using this package (e.g. `ubasic.HasGcLl(0x61)`.
- Names containing **PkgUnicode** - perform the tests using the standard unicode package (e.g. `unicode.Is(unicode.Ll, 0x61)`.
- Names containing **PkgRegexp** - perform the tests using the standard regexp package, providing a pre-compiled regexp (e.g. `\p{Ll}`) (e.g. `reLl.MatchString("/x61")`).
- Names not containing **Ascii** - perform 36 different General_Category queries on 14663 different code points scattered throughout the range of valid code points, for a total of 527868 tests.
- Names containing **Ascii** - perform 36 different General_Category queries on 127 code points within the ASCII range, for a total of 4572 tests.
