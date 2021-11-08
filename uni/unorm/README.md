<h1 align="center">Unicode​® normalization (NFC, NFD, NFKC, NFKD) and related queries</h1>

**unorm** is a Go package for normalizaing Unicode text to canonical forms, NFC, NFD, NFKC, or NFKD.

For more information about these properties, consult [UAX 15](https://www.unicode.org/reports/tr15/) and chapter 3 of the Unicode Specification at [https://www.unicode.org/](https://www.unicode.org/).

Other Unicode properties are supported by other packages in this module. For a summary of Unicode properties and which functions of which packages of this module provide the corresponding support, peruse the file, **SUMMARY.txt**, in the top level directory of this module.

This functionality is also provided in golang.org/x/text/unicode/norm. Each package has its advantages and disadvantages. As of this writing, this package appears to perform better in most cases. The following is an example output of `go test -bench` on this author's platform.

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

The above benchmarks compare this package with golang.org/x/text/unicode/norm. It shows that for ASCII text, both process it very fast but the golang package performs better. Other tests generally show this package performing better. For example, for Hangul (i.e. the Korean alphabet), for NFC normalization, this package takes 0.0000747 ns/op but the golang/x/text/unicode/norm package takes 0.001203 ns/op, or about 16 times longer. The "QcSingles" is a test of code points that are already normalized to all four forms without any decomposition, composition, or non-starters, commonly used in a variety of languages and should be processed quickly, with this package performing about 2400/765 or a little more than 3 times faster for all four normalization forms.

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
    // convert textNfd to be in NFC form, like text Nfc:
    fmt.Printf("textA \"%s\", composed as: %U\n", textA, []rune(textA))
    fmt.Printf("textB \"%s\", composed as: %U\n", textA, []rune(textB))
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
# API
## General Constants
### UnicodeVersion
UnicodeVersion is the current version of the Unicode specification that this package is intended to comply with.

UnicodeVersion = "14.0.0".
## Normalization
### func NextNFC(data []rune, offset int) ([]rune, int)
NextNFC is equivalent to `NextNFCLen(data, offset, len(data), false)`.
### func NextNFCLen(data []rune, offset int, dataLength int, abandonAtEnd bool) ([]rune, int)
NextNFCLen performs Unicode NFC normalization on the sequence of runes started at the specified offset in data. The returned value consists of the normalized runes and the count of runes consumed from the input data to produce them. If abandonAtEnd is true and the next sequence of runes cannot be unambiguously determined without knowing subsequent data, then the normalization process will abort early and return 0 runes with the length indicated to the end of the buffer. Normally, abandonAtEnd is set to false except in cases where the caller is streaming input through a buffer, to allow the caller to refill the buffer without the overhead of completing the normalization on the last chunk of data in the buffer.

NextNFCLen may return a slice of runes from the input buffer. Typically, this is the case if it determines that the input data is already normalized. The caller should take care to not modify the contents of the input buffer until the output data has been consumed.
### func NFCReader(rr io.RuneReader) io.RuneReader
NextNFCReader is a wrapper over NextNFCLen that filters the input through an `io.RuneReader` and outputs to another `io.RuneReader`. This may be a more convenient interface in some cases and also avoids the potential implications of sharing an input buffer with the output.

The performance of this function is typically lower than calling `NextNFCLen` directly.

Due to internal caching, this function may fail to normalize some extreme cases where a character is composed of a large number of non-starter characters. [UAX#15](https://golang.org/doc/) defines a ##Stream-Safe Text Format## which limits sequences of non-starters to no more than 30. This function handles text in Stream-Safe Text Format.
### func NextNFD(data []rune, offset int) ([]rune, int)
NextNFD is equivalent to `NextNFDLen(data, offset, len(data), false)`.
### func NextNFDLen(data []rune, offset int, dataLength int, abandonAtEnd bool) ([]rune, int)
NextNFDLen is like NextNFCLen execpt it performs NFD normalization.
### func NFDReader(rr io.RuneReader) io.RuneReader
NextNFDReader is like NextNFCReader except it is a wrapper over NextNFDLen.
### func NextNFKC(data []rune, offset int) ([]rune, int)
NextNFKC is equivalent to `NextNFKCLen(data, offset, len(data), false)`.
### func NextNFKCLen(data []rune, offset int, dataLength int, abandonAtEnd bool) ([]rune, int)
NextNFKCLen is like NextNFCLen execpt it performs NFKC normalization.
### func NFKCReader(rr io.RuneReader) io.RuneReader
NextNFKCReader is like NextNFCReader except it is a wrapper over NextNFKCLen.
### func NextNFKD(data []rune, offset int) ([]rune, int)
NextNFKD is equivalent to `NextNFKDLen(data, offset, len(data), false)`.
### func NextNFKDLen(data []rune, offset int, dataLength int, abandonAtEnd bool) ([]rune, int)
NextNFKDLen is like NextNFCLen execpt it performs NFKD normalization.
### func NFKDReader(rr io.RuneReader) io.RuneReader
NextNFKDReader is like NextNFCReader except it is a wrapper over NextNFKDLen.
## Unicode Property Queries
### func ValueOfCcc(r rune, d int) int
ValueOfCcc returns the value of Unicode property, Canonical_Combining_Class (ccc), for code point, r. If r is not assigned a value for this property, then d is returned.
### func HasNfcqcM(r rune) bool
HasNfcqcM returns true if code point, r, has Unicode property, NFC_Quick_Check (NFC_QC), of Maybe (M).
### func HasNfcqcN(r rune) bool
HasNfcqcN returns true if code point, r, has Unicode property, NFC_Quick_Check (NFC_QC), of No (N).
### func HasNfcqcY(r rune) bool
HasNfcqcY returns true if code point, r, has Unicode property, NFC_Quick_Check (NFC_QC), of Yes (Y).
### func HasNfdqcN(r rune) bool
HasNfdqcN returns true if code point, r, has Unicode property, NFD_Quick_Check (NFD_QC), of No (N).
### func HasNfdqcY(r rune) bool
HasNfdqcY returns true if code point, r, has Unicode property, NFD_Quick_Check (NFD_QC), of Yes (Y).
### func HasNfkcqcM(r rune) bool
HasNfkcqcM returns true if code point, r, has Unicode property, NFKC_Quick_Check (NFKC_QC), of Maybe (M).
### func HasNfkcqcN(r rune) bool
HasNfkcqcN returns true if code point, r, has Unicode property, NFKC_Quick_Check (NFKC_QC), of No (N).
### func HasNfkcqcY(r rune) bool
HasNfkcqcY returns true if code point, r, has Unicode property, NFKC_Quick_Check (NFKC_QC), of Yes (Y).
### func HasNfkdqcN(r rune) bool
HasNfkdqcN returns true if code point, r, has Unicode property, NFKD_Quick_Check (NFKD_QC), of No (N).
### func HasNfkdqcY(r rune) bool
HasNfkdqcY returns true if code point, r, has Unicode property, NFKD_Quick_Check (NFKD_QC), of Yes (Y).
## Other Functions
# Supported Unicode Properties
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
