<h1 align="center">Unicode ® normalization (NFC, NFD, NFKC, NFKD) and related queries</h1>

**unorm** is a Go package for normalizaing Unicode text to canonical forms, NFC, NFD, NFKC, or NFKD.

Normalization is an important consideration especially when dealing with text in languages other than English. For more information about what normalization is and why it is important, see the section, **Why is Normalization Important**, below.

This functionality is also provided in the standard Go library at golang.org/x/text/unicode/norm. Each package has its advantages and disadvantages. As of this writing, this package appears to perform better in most cases, as shown by running `go test -bench`. Results may vary significantly with other platforms, software versions, and even just running the test in the same setup multiple times. The following summary is based on this author's experience of running `go text -bench` on this package.

- Both packages (this one and golang.org/x/text/unicode/norm) perform very fast on ASCII text (the most common characters used in English and similar European languages). This is expected because ASCII text is simple to process. However, because other characters take so much longer, even a small percentage of non-ASCII characters will overwhelm any differences in ASCII performance.
- For non-ASCII characters that do not need to be normalized in any form, this package runs between 3 and 4 times faster.
- For other tests of normalization, this package runs between the same to about 5 times faster, depending on the test.
- For Hangul characters (i.e. the Korean alphabet), this package runs about 17 to 18 times faster for the NFC and NFKC forms and roughly 3 times faster for the NFD and NFKD forms. Hangul characters are normalized differently than other Unicode characters.

More information about benchmarking can be found later in this document in the section, **Performance**.

_Unicode and the Unicode Logo are registered trademarks of Unicode, Inc. in the United States and other countries._

## Example

If you are new to the Go language, consider perusing the documentation at [golang.org](https://golang.org/doc/) before proceeding.

Download and install it:

```sh
go get github.com/orkvozku/go/uni

```
Example:

```go
package main

import "github.com/orkvozku/go/uni/unorm"
import "fmt"

func main() {
    textA := "caf\u00E9" // NFC form: é as U+00E9
    textB := "cafe\u0301" // NFD form: é as e followed by U+0301
    // convert textB to be in NFC form, like textA:
    fmt.Printf("textA \"%s\", composed as: %U\n", textA, []rune(textA))
    fmt.Printf("textB \"%s\", composed as: %U\n", textB, []rune(textB))
    if textA == textB {
        fmt.Printf("textA matches textB.\n")
    } else {
        fmt.Printf("textA does not match textB.\n")
    }
    runesB := []rune(textB)
    runesC := []rune{}
    // Walk through our runesB (of textB) and normalize into runesC.
    ofs := 0
    for {
        rs, delta := unorm.NextNFC(runesB, ofs)
        if delta == 0 {
            break
        }
        runesC = append(runesC, rs...)
        ofs += delta
    }
    // Let's convert our runesC to a string.
    textC := string(runesC)
    fmt.Printf("textC \"%s\", composed as: %U\n", textC, []rune(textC))
    if textA == textC {
        fmt.Printf("textA matches textC.\n")
    } else {
        fmt.Printf("textA does not match textC.\n")
    }

    // output:
    // textA "café", composed as: [U+0063 U+0061 U+0066 U+00E9]
    // textB "café", composed as: [U+0063 U+0061 U+0066 U+0065 U+0301]
    // textA does not match textB.
    // textC "café", composed as: [U+0063 U+0061 U+0066 U+00E9]
    // textA matches textC.
}
```
Example: Various Usage Cases:

```go
package main

import "github.com/orkvozku/go/uni/unorm"
import "strings"
import "fmt"

func main() {
    txt       := "d'à côté"
    txtNFC1   := ""
    txtNFD1   := ""
    txtNFC2   := ""
    txtNFD2   := ""
    txtRunes  := []rune(txt)
    var rs    []rune
    var ofs   int
    var delta int
    //=====================================================
    // txt => txtNFC1 using unorm.NextNFC:
    ofs = 0
    for {
        rs, delta = unorm.NextNFC(txtRunes, ofs)
        if delta == 0 {
            break
        }
        txtNFC1 += string(rs)
        ofs += delta
    }
    //=====================================================
    // txt => txtNFD1 using unorm.NextNFDLen:
    ofs = 0
    for {
        rs, delta = unorm.NextNFDLen(txtRunes, ofs, len(txtRunes), false)
        if delta == 0 {
            break
        }
        txtNFD1 += string(rs)
        ofs += delta
    }
    //=====================================================
    // txt => txtNFC2 using unorm.NFCReader:
    for rr := unorm.NFCReader(strings.NewReader(txt)); ; {
        ch, _, err := rr.ReadRune()
        if err != nil {
            break
        }
        txtNFC2 += string(ch)
    }
    //=====================================================
    // txt => txtNFD2 using unorm.NFDReader:
    for rr := unorm.NFDReader(strings.NewReader(txt)); ; {
        ch, _, err := rr.ReadRune()
        if err != nil {
            break
        }
        txtNFD2 += string(ch)
    }
    //=====================================================
    // output:
    fmt.Printf("NFC1: \"%s\" %U\n", txtNFC1, []rune(txtNFC1))
    fmt.Printf("NFD1: \"%s\" %U\n", txtNFD1, []rune(txtNFD1))
    fmt.Printf("NFC2: \"%s\" %U\n", txtNFC2, []rune(txtNFC2))
    fmt.Printf("NFD2: \"%s\" %U\n", txtNFD2, []rune(txtNFD2))

    // output:
    // NFC1: "d'à côté" [U+0064 U+0027 U+00E0 U+0020 U+0063 U+00F4 U+0074 U+00E9]
    // NFD1: "d'à côté" [U+0064 U+0027 U+0061 U+0300 U+0020 U+0063 U+006F U+0302 U+0074 U+0065 U+0301]
    // NFC2: "d'à côté" [U+0064 U+0027 U+00E0 U+0020 U+0063 U+00F4 U+0074 U+00E9]
    // NFD2: "d'à côté" [U+0064 U+0027 U+0061 U+0300 U+0020 U+0063 U+006F U+0302 U+0074 U+0065 U+0301]
}
```
## Why is Normalization Important
The Unicode standard allows for multiple ways to encode the same text. This means, for example, that `"cliché" != "cliché"`, in spite of being the exact same word. Also, `len("cliché") != len("cliché")` because one of them uses two Unicode code points for the **é** while the other uses only one. The Unicode standard defines two different normalization forms that can be used to encode these to a common canonical form, called NFC and NFD. NFC is the more common and compact representation. NFD is an alternate form that is (usually) less compact. Additionally, there are two other analogous forms, NFKC and NFKD, that map different but related characters into a common form. For example, superscripts and subscripts map to their base forms in NFKC and NFKD (e.g. **E=mc²** becomes **E=mc2**). As a less familiar example, the Latin Small Letter Long S, **ſ**, is mapped by NFKC and NFKD to the more familiar lower case **s**. NFKC and NFKD are used less often, but may be useful in cases where it is desirable to combine related forms, such as key word searches.

In most cases, English language text is already normalized to all four forms. Normalization becomes more important when dealing with other languages.
### Example: Web Pages
Consider the following example web page consisting of two files.

**sample.css**:

```css
body {
    color: #0000FF;
    background: #FFFFFF;
    font-weight: bold;
}
.bronzé {
    color: #CD7F32;
}
```
**sample.html**:

```html
<html><head>
<link href="sample.css" rel="stylesheet">
</head><body>
<p class="bronzé">Is this blue or bronze?</p>
</body></html>
```

Is the text displayed in blue or bronze? That depends on whether or not the web browser matches the class in `<p class="bronzé">` to the `.bronzé` in the .css file. The **é** in **bronzé** may be encoded as either one code point, U+00E9, or two code points, U+0065 U+0301. If sample.css has it encoded in one way while sample.html has it in the other, they might not match. A correctly-implemented web browser will normalize both of these to either NFC or NFD form before comparing them to ensure the text shows up as bronze regardless of the underlying encoding. Older or buggy web browsers might show the text in blue instead. The W3C has an online tool that web developers can use to check for possible inconsistencies [W3C Internationalization Checker](https://validator.w3.org/i18n-checker/).

### Which Should I Use
If you are unsure which form to use, NFC is typically the best choice. Normalization is more efficient in cases where the incoming text is already encoded in that form and most real world text is already encoded in NFC form. Also, NFC is usually more compact than NFD. In most cases, either NFC or NFD will work fine as long as you are consistent. NFKC and NFKD should be avoided except in cases where you know you want to use them.

### Further Reading
There is a nicely written article about normalization here [Text normalization in Go](https://go.dev/blog/normalization). If you want to get into the details, consult [UAX#15](https://unicode.org/reports/tr15/) and [Chapter 3 of the Unicode Standard](https://www.unicode.org/versions/latest/).
## API
### General Constants
#### UnicodeVersion
UnicodeVersion is the current version of the Unicode specification that this package is intended to comply with.

UnicodeVersion = "14.0.0".
### Normalization
#### func NextNFC(data []rune, offset int) ([]rune, int)
NextNFC is equivalent to `NextNFCLen(data, offset, len(data), false)`.
#### func NextNFCLen(data []rune, offset int, dataLength int, abandonAtEnd bool) ([]rune, int)
NextNFCLen performs Unicode NFC normalization on the sequence of runes started at the specified offset in data. The returned value consists of the normalized runes and the count of runes consumed from the input data to produce them. If abandonAtEnd is true and the next sequence of runes cannot be unambiguously determined without knowing subsequent data, then the normalization process will abort early and return 0 runes with the length indicated to the end of the buffer. Normally, abandonAtEnd is set to false except in cases where the caller is streaming input through a buffer, to allow the caller to refill the buffer without the overhead of completing the normalization on the last chunk of data in the buffer.

NextNFCLen may return a slice of runes from the input buffer. Typically, this is the case if it determines that the input data is already normalized. The caller should take care to not modify the contents of the input buffer until the output data has been consumed.
#### func NFCReader(rr io.RuneReader) io.RuneReader
NextNFCReader is a wrapper over NextNFCLen that filters the input through an `io.RuneReader` and outputs to another `io.RuneReader`. This may be a more convenient interface in some cases and also avoids the potential implications of sharing an input buffer with the output.

The performance of this function is typically lower than calling `NextNFCLen` directly.

Due to internal caching, this function may fail to normalize some extreme cases where a character is composed of a large number of non-starters. These should only occur in contrived cases. [UAX#15](https://golang.org/doc/) defines a **Stream-Safe Text Format** which limits sequences of non-starters to no more than 30. This function handles text in Stream-Safe Text Format.
#### func NextNFD(data []rune, offset int) ([]rune, int)
NextNFD is equivalent to `NextNFDLen(data, offset, len(data), false)`.
#### func NextNFDLen(data []rune, offset int, dataLength int, abandonAtEnd bool) ([]rune, int)
NextNFDLen is like NextNFCLen execpt it performs NFD normalization.
#### func NFDReader(rr io.RuneReader) io.RuneReader
NextNFDReader is like NextNFCReader except it is a wrapper over NextNFDLen.
#### func NextNFKC(data []rune, offset int) ([]rune, int)
NextNFKC is equivalent to `NextNFKCLen(data, offset, len(data), false)`.
#### func NextNFKCLen(data []rune, offset int, dataLength int, abandonAtEnd bool) ([]rune, int)
NextNFKCLen is like NextNFCLen execpt it performs NFKC normalization.
#### func NFKCReader(rr io.RuneReader) io.RuneReader
NextNFKCReader is like NextNFCReader except it is a wrapper over NextNFKCLen.
#### func NextNFKD(data []rune, offset int) ([]rune, int)
NextNFKD is equivalent to `NextNFKDLen(data, offset, len(data), false)`.
#### func NextNFKDLen(data []rune, offset int, dataLength int, abandonAtEnd bool) ([]rune, int)
NextNFKDLen is like NextNFCLen execpt it performs NFKD normalization.
#### func NFKDReader(rr io.RuneReader) io.RuneReader
NextNFKDReader is like NextNFCReader except it is a wrapper over NextNFKDLen.
### Unicode Property Queries
#### func ValueOfCcc(r rune, d int) int
ValueOfCcc returns the value of Unicode property, Canonical_Combining_Class (ccc), for code point, r. If r is not assigned a value for this property, then d is returned.
#### func HasNfcqcM(r rune) bool
HasNfcqcM returns true if code point, r, has Unicode property, NFC_Quick_Check (NFC_QC), of Maybe (M).
#### func HasNfcqcN(r rune) bool
HasNfcqcN returns true if code point, r, has Unicode property, NFC_Quick_Check (NFC_QC), of No (N).
#### func HasNfcqcY(r rune) bool
HasNfcqcY returns true if code point, r, has Unicode property, NFC_Quick_Check (NFC_QC), of Yes (Y).
#### func HasNfdqcN(r rune) bool
HasNfdqcN returns true if code point, r, has Unicode property, NFD_Quick_Check (NFD_QC), of No (N).
#### func HasNfdqcY(r rune) bool
HasNfdqcY returns true if code point, r, has Unicode property, NFD_Quick_Check (NFD_QC), of Yes (Y).
#### func HasNfkcqcM(r rune) bool
HasNfkcqcM returns true if code point, r, has Unicode property, NFKC_Quick_Check (NFKC_QC), of Maybe (M).
#### func HasNfkcqcN(r rune) bool
HasNfkcqcN returns true if code point, r, has Unicode property, NFKC_Quick_Check (NFKC_QC), of No (N).
#### func HasNfkcqcY(r rune) bool
HasNfkcqcY returns true if code point, r, has Unicode property, NFKC_Quick_Check (NFKC_QC), of Yes (Y).
#### func HasNfkdqcN(r rune) bool
HasNfkdqcN returns true if code point, r, has Unicode property, NFKD_Quick_Check (NFKD_QC), of No (N).
#### func HasNfkdqcY(r rune) bool
HasNfkdqcY returns true if code point, r, has Unicode property, NFKD_Quick_Check (NFKD_QC), of Yes (Y).
### Other Functions
#### func CanonicalComposition(r1, r2 rune) (rune, bool)
CanonicalComposition returns the canonical composition of runes r1 and r2. Values which are included in the property, Full_Composition_Exclusion, are not included. The second parameter is returned true to indicate a composition was found. If not composition was found, then 0 and false are returned.
#### func CanonicalDecomposition(r rune) ([]rune, []int)
CanonicalDecomposition returns the canonical decomposition (i.e. for normalization forms NFC and NFD) of code point, r, along with the value of the Canonical_Combining_Class (ccc) for each returned code point. The two returned slices are of equal length. If there is no Canonical Decomposition for this code point, then []rune{} and []int{} are returned.
#### func CompatibleDecomposition(r rune) ([]rune, []int)
CompatibleDecomposition returns the compatible decomposition (i.e. for normalization forms NFKC and NFKD) of code point, r, along with the value of the Canonical_Combining_Class (ccc) for each returned code point. The two returned slices are of equal length. If there is no Compatible Decomposition for this code point, then []rune{} and []int{} are returned.
#### func IsStarterNfc(r rune) bool
IsStarterNfc returns the equivalent of `ValueOfCcc(r, -1)==0 && HasNfcqcY(r)`.
#### func IsStarterNfd(r rune) bool
IsStarterNfd returns the equivalent of `ValueOfCcc(r, -1)==0 && HasNfdqcY(r)`.
#### func IsStarterNfkc(r rune) bool
IsStarterNfkc returns the equivalent of `ValueOfCcc(r, -1)==0 && HasNfkcqcY(r)`.
#### func IsStarterNfkd(r rune) bool
IsStarterNfkd returns the equivalent of `ValueOfCcc(r, -1)==0 && HasNfkdqcY(r)`.
## Supported Unicode Properties
<table><thead><tr><th>Property:</th><th>Value:</th><th>Function:</th></tr></thead>
<tbody>
<tr><td colspan="2">Canonical_Combining_Class (ccc)</td><td>ValueOfCcc(r rune, d int) int</td>
<tr><td colspan="3">NFC_Quick_Check (NFC_QC)</td></tr>
<tr><td>&nbsp;</td><td>Maybe (M)</td><td>HasNfcqcM(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>No (N)</td><td>HasNfcqcN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y)</td><td>HasNfcqcY(r rune) bool</td></tr>
<tr><td colspan="3">NFD_Quick_Check (NFD_QC)</td></tr>
<tr><td>&nbsp;</td><td>No (N)</td><td>HasNfdqcN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y)</td><td>HasNfdqcY(r rune) bool</td></tr>
<tr><td colspan="3">NFKC_Quick_Check (NFKC_QC)</td></tr>
<tr><td>&nbsp;</td><td>Maybe (M)</td><td>HasNfkcqcM(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>No (N)</td><td>HasNfkcqcN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y)</td><td>HasNfkcqcY(r rune) bool</td></tr>
<tr><td colspan="3">NFKD_Quick_Check (NFKD_QC)</td></tr>
<tr><td>&nbsp;</td><td>No (N)</td><td>HasNfkdqcN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y)</td><td>HasNfkdqcY(r rune) bool</td></tr>
</tbody></table>

## Performance
Running the `go test -bench` command on this package produces performance benchmarks that compare this package's normalization functions to those of the golang.org/x/text/unicode/norm. Because of the nature of the normalization process, different patterns of text may yield significantly different results. Therefore, a variety of different test samples are needed to get an understanding of the performance comparison. The following benchmarks explore some specific patterns that provide a broad sense of the performance behavior, but may not include all potentially interesting kinds of patterns.

Each test is named with the pattern, **Benchmark** (**Orkvozku** or **Norm**) (**Ascii**, **QcSingles**, **General**, **DeCo**, or **Hangul**) (**NFC**, **NFD**, **NFKC**, or **NFKD**).

- **Orkvozku** - uses this package by walking through a slice of runes using `unorm.NextNFCLen`, `unorm.NexdNFDLen`, `unorm.NextNFKCLen`, or `unorm.NextNFKDLen`.
- **Norm** - uses the golang.org/x/text/unicode/norm package by passing a string to `norm.NFC.String`, `norm.NFD.String`, `norm.NFKC.String`, or `norm.NFKD.String`. This is not entirely analogous to the Orkvozku as it receives the entire set of data in one call as a string, while the Orkvozku case may make multiple calls as it walks through a slice of runes. The intention is to utilize the best method associated with each package.
- **Ascii** - 12424 code points in the ASCII range. Most text in English and most other European languages is in the ASCII range. Processing is extremely fast with both packages. The performance data below indicates processing them is roughly 10x to 30X faster over 28.5% more data compared to the QcSingles data, which is another set of commonly used characters that are already normalized in all four forms. Therefore, even a small percentage of non-ASCII characters in the input can dominate the performance profile. The performance data below suggests that both packages have similar performance for ASCII characters.
- **QcSingles** - 9662 code points consisting of non-ASCII and non-Hangul characters of which single code points are already normalized to all four forms without decomposition or composition forms. These are very common characters in most languages.
- **General** - 15364 code points consisting of characters that have either canonical or compatible decompositions, along with their corresponding decomposed forms. These are meant to exercise typical cases where normalization work is done.
- **DeCo** - 11239 code points that exercise common yet heavier than average normalization work, with expanding decomposition followed by reducing composition.
- **Hangul** - 14763 code points consisting of composed Hangul characters followed by their decomposed form. These are the characters of the Korean alphabet and are handled differently by the normalization algorithms.
- **NFC**, **NFD**, **NFKC**, or **NFKD** - indicate which of the four normalization forms is being tested.

For example, the test named **BenchmarkOrkvozkuQcSinglesNFD** tests this package on the QcSingles test pattern for normalization form NFD.

The following is an example from one test performed on one platform.

```sh
BenchmarkOrkvozkuAsciiNFC-8        	1000000000	         0.0000085 ns/op
BenchmarkNormAsciiNFC-8            	1000000000	         0.0000058 ns/op
BenchmarkOrkvozkuAsciiNFD-8        	1000000000	         0.0000059 ns/op
BenchmarkNormAsciiNFD-8            	1000000000	         0.0000099 ns/op
BenchmarkOrkvozkuAsciiNFKC-8       	1000000000	         0.0000086 ns/op
BenchmarkNormAsciiNFKC-8           	1000000000	         0.0000058 ns/op
BenchmarkOrkvozkuAsciiNFKD-8       	1000000000	         0.0000064 ns/op
BenchmarkNormAsciiNFKD-8           	1000000000	         0.0000058 ns/op
BenchmarkOrkvozkuQcSinglesNFC-8    	1000000000	         0.0000765 ns/op
BenchmarkNormQcSinglesNFC-8        	1000000000	         0.0002403 ns/op
BenchmarkOrkvozkuQcSinglesNFD-8    	1000000000	         0.0000770 ns/op
BenchmarkNormQcSinglesNFD-8        	1000000000	         0.0002428 ns/op
BenchmarkOrkvozkuQcSinglesNFKC-8   	1000000000	         0.0000764 ns/op
BenchmarkNormQcSinglesNFKC-8       	1000000000	         0.0002399 ns/op
BenchmarkOrkvozkuQcSinglesNFKD-8   	1000000000	         0.0000786 ns/op
BenchmarkNormQcSinglesNFKD-8       	1000000000	         0.0002433 ns/op
BenchmarkOrkvozkuGeneralNFC-8      	1000000000	         0.0006274 ns/op
BenchmarkNormGeneraloNFC-8         	1000000000	         0.0008353 ns/op
BenchmarkOrkvozkuGeneralNFD-8      	1000000000	         0.0005420 ns/op
BenchmarkNormGeneralNFD-8          	1000000000	         0.0007311 ns/op
BenchmarkOrkvozkuGeneralNFKC-8     	1000000000	         0.001221 ns/op
BenchmarkNormGeneralNFKC-8         	1000000000	         0.001240 ns/op
BenchmarkOrkvozkuGeneralNFKD-8     	1000000000	         0.001187 ns/op
BenchmarkNormGeneralNFKD-8         	1000000000	         0.001139 ns/op
BenchmarkOrkvozkuDeCoNFC-8         	1000000000	         0.0001411 ns/op
BenchmarkNormDeCoNFC-8             	1000000000	         0.0004357 ns/op
BenchmarkOrkvozkuDeCoNFD-8         	1000000000	         0.0003141 ns/op
BenchmarkNormDeCoNFD-8             	1000000000	         0.001586 ns/op
BenchmarkOrkvozkuDeCoNFKC-8        	1000000000	         0.0003737 ns/op
BenchmarkNormDeCoNFKC-8            	1000000000	         0.0005205 ns/op
BenchmarkOrkvozkuDeCoNFKD-8        	1000000000	         0.0005310 ns/op
BenchmarkNormDeCoNFKD-8            	1000000000	         0.001673 ns/op
BenchmarkOrkvozkuHangulNFC-8       	1000000000	         0.0000747 ns/op
BenchmarkNormHangulNFC-8           	1000000000	         0.001203 ns/op
BenchmarkOrkvozkuHangulNFD-8       	1000000000	         0.0003600 ns/op
BenchmarkNormHangulNFD-8           	1000000000	         0.0009393 ns/op
BenchmarkOrkvozkuHangulNFKC-8      	1000000000	         0.0000725 ns/op
BenchmarkNormHangulNFKC-8          	1000000000	         0.001232 ns/op
BenchmarkOrkvozkuHangulNFKD-8      	1000000000	         0.0003471 ns/op
BenchmarkNormHangulNFKD-8          	1000000000	         0.001021 ns/op
PASS
ok  	github.com/orkvozku/go/uni/unorm	0.286s
```
