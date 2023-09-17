<h1 align="center">Unicode ® queries of segmentation properties</h1>

**useg** is a Go package for querying Unicode segmentation properties.

For more information about these properties, consult [UAX 29](https://www.unicode.org/reports/tr29/) and [UAX 14](https://www.unicode.org/reports/tr14/).

Other Unicode properties are supported by other packages in this module. For a summary of Unicode properties and which functions of which packages of this module provide the corresponding support, peruse the file, **SUMMARY.txt**, in the top level directory of this module.

For example, the function, `HasGcLu(rune r) bool`, returns true only if the code point, r, has Unicode property, General_Category, with value, Uppercase_Letter (Lu), often referred to in regular expressions as "/p{Lu}".

This package also provides parsing of extended grapheme cluster boundaries. In many cases, especially dealing with various languages and emojis, what we think of as a "character" may span multiple code points. These code points provide attributes that contribute to the meaning of the character. Software that processes text in segments (i.e. most software that manipulates text) should divide that text only at extended grapheme cluster boundaries. For more details about extended grapheme cluster boundaries, consult [Unicode Standard Annex #29](https://www.unicode.org/reports/tr29/).

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

import "github.com/orkvozku/go/uni/useg"
import "fmt"

func main() {
    text := "Hi! \U0001f467\u200d\U0001f469 :)"
    runes := []rune(text)
    offset := 0
    for {
        delta, _ := useg.NextExtGraphemeClusterBreak(runes, offset)
        if delta == 0 {
            break
        }
        egc := runes[offset : offset+delta]
        fmt.Printf("Extended grapheme cluster '%s': %U\n", string(egc), egc)
        offset += delta
    }

    // output:
    // Extended grapheme cluster 'H': [U+0048]
    // Extended grapheme cluster 'i': [U+0069]
    // Extended grapheme cluster '!': [U+0021]
    // Extended grapheme cluster ' ': [U+0020]
    // Extended grapheme cluster '👧‍👩': [U+1F467 U+200D U+1F469]
    // Extended grapheme cluster ' ': [U+0020]
    // Extended grapheme cluster ':': [U+003A]
    // Extended grapheme cluster ')': [U+0029]
}
```
Example - Alternative based on ucutls.NewRunesReaderFromBreaker:

```go
package main

import "github.com/orkvozku/go/uni/useg"
import "github.com/orkvozku/go/uni/ucutls"
import "strings"
import "fmt"

func main() {
    txt := "Hi! \U0001f467\u200d\U0001f469 :)"
    r := strings.NewReader(txt)
    rrb := ucutls.NewRunesReaderFromBreaker( useg.NextExtGraphemeClusterBreakLen,
        64, // anything >= the largest extended grapheme cluster works here
        r)
    for {
        egc, _, err := rrb.ReadRunes()
        if ucutls.IsEOF(err) {
            break
        }
        fmt.Printf("Extended grapheme cluster '%s': %U\n", string(egc), egc)
    }

    // output:
    // Extended grapheme cluster 'H': [U+0048]
    // Extended grapheme cluster 'i': [U+0069]
    // Extended grapheme cluster '!': [U+0021]
    // Extended grapheme cluster ' ': [U+0020]
    // Extended grapheme cluster '👧‍👩': [U+1F467 U+200D U+1F469]
    // Extended grapheme cluster ' ': [U+0020]
    // Extended grapheme cluster ':': [U+003A]
    // Extended grapheme cluster ')': [U+0029]
}
```
## API
### General Constants
#### UnicodeVersion
UnicodeVersion is the current version of the Unicode specification that this package is intended to comply with.

UnicodeVersion = "15.1.0".
### Unicode Property Queries
#### func ValueOfGcb(r rune, d int) int
ValueOfGcb returns the value of Unicode property, Grapheme_Cluster_Break (GCB), for code point, r. If r is not assigned a value for this property, then d is returned.
#### func HasExtpictN(r rune) bool
HasExtpictN returns true if code point, r, has Unicode property, Extended_Pictographic (ExtPict), of No (N).
#### func HasExtpictY(r rune) bool
HasExtpictY returns true if code point, r, has Unicode property, Extended_Pictographic (ExtPict), of Yes (Y).
#### func ValueOfLb(r rune, d int) int
ValueOfLb returns the value of Unicode property, Line_Break (lb), for code point, r. If r is not assigned a value for this property, then d is returned.
#### func ValueOfWb(r rune, d int) int
ValueOfWb returns the value of Unicode property, Word_Break (WB), for code point, r. If r is not assigned a value for this property, then d is returned.
#### func ValueOfSb(r rune, d int) int
ValueOfSb returns the value of Unicode property, Sentence_Break (SB), for code point, r. If r is not assigned a value for this property, then d is returned.
#### func ValueOfIncb(r rune, d int) int
ValueOfIncb returns the value of Unicode property, Indic_Conjunct_Break (InCB), for code point, r. If r is not assigned a value for this property, then d is returned.
### Other Functions
## Supported Unicode Properties
<table><thead><tr><th>Property:</th><th>Value:</th><th>Function:</th></tr></thead>
<tbody>
<tr><td colspan="2">Grapheme_Cluster_Break (GCB)</td><td>ValueOfGcb(r rune, d int) int</td>
<tr><td colspan="3">Extended_Pictographic (ExtPict)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasExtpictN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasExtpictY(r rune) bool</td></tr>
<tr><td colspan="2">Line_Break (lb)</td><td>ValueOfLb(r rune, d int) int</td>
<tr><td colspan="2">Word_Break (WB)</td><td>ValueOfWb(r rune, d int) int</td>
<tr><td colspan="2">Sentence_Break (SB)</td><td>ValueOfSb(r rune, d int) int</td>
<tr><td colspan="2">Indic_Conjunct_Break (InCB)</td><td>ValueOfIncb(r rune, d int) int</td>
</tbody></table>
