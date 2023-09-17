<h1 align="center">a collection of Unicode ® packages Go/Golang</h1>

This module contains several packages to support use of Unicode properties. These include common operations such as property queries, as well as some other support such as parsing extended grapheme cluster breaks. For a full list, see the file, **SUMMARY.txt** in this directory, and the **README.md** files in the individual package directories. It also includes a normalization package, **unorm** that appears faster than the golang.org unicode/norm package.

Note with Module Version 0.0.11 update to correspond with Unicode 15.1.0, there were some changes to the API. For the following properties, the functions of the form, HasXY(r rune) bool, have been deprecated and replaced with one function for the entire
property returning the value. These deprecated functions will go away soon, probably with the next update.

- Grapheme_Cluster_Break (GCB) - use `ValueOfGcb(r rune, -1) int` (ubasic, useg)
- Line_Break (lb) - use `ValueOfLb(r rune, -1) int` (useg)
- Sentence_Break (SB) - use `ValueOfSb(r rune, -1) int` (useg)
- Word_Break (WB) - use `ValueOfWb(r rune, -1) int` (useg)

For example, deprecated `ubasic.HasGcbEx(r)` can be replaced with `ubasic.ValueOfGcb(r, -1) == ubasic.GcbValueEx`.

_Unicode and the Unicode Logo are registered trademarks of Unicode, Inc. in the United States and other countries._

## Example: Unicode General_Category

This property is one of the most popular and well-known properties. It has values such as **Ll** (Lowercase_Letter), **L** (Letter), **N** (Number), and so on. Perusing the file, **SUMMARY.txt** shows that the **ubasic** package supports this property through functions such as `HasGcLu`, `HasGcLl`, `HasGcP`, etc.

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

## Example: Extended Grapheme Clusters

What we think of as a "character" does not always correspond to one Unicode code point. Many of them do, such as the Latin alphabet capital A (U+0041), but there are numerous exceptions. For example, क्षि, is composed of four Unicode code points, U+0915, U+094D, U+0937, and U+093F. Many emojis are composed of multiple code points, or may include modifiers for such things and preferred skin tone. If software breaks text within these extended grapheme cluster boundaries, the results can seem unexpected. For more information, consult [UAX #29](https://www.unicode.org/reports/tr29/). This support is provided in both the **ubasic** and **useg** packages. These packages provide the Unicode property queries necessary to parse them, but also provide parsing functions so you do not have to do it. Another package, **ucutls**, provides a utility to wrap these functions into pipe that receives input from an io.RuneReader and outputs to an ucutls.RunesReader.

```go
package main

import "github.com/orkvozku/go/uni/ucutls"
import "github.com/orkvozku/go/uni/ubasic"
import "strings"
import "fmt"

func main() {
    txt := "คุณเป็นอย่างไร?"
    r := strings.NewReader(txt)
    rrb := ucutls.NewRunesReaderFromBreaker( ubasic.NextExtGraphemeClusterBreakLen,
        256, // anything >= the largest extended grapheme cluster works here
        r)
    fmt.Printf("The text is %d bytes.\n", len(txt))
    fmt.Printf("The text is %d runes.\n", len([]rune(txt)))
    tegcs := 0
    for {
        egc, _, err := rrb.ReadRunes()
        if ucutls.IsEOF(err) {
            break
        }
        tegcs++
        fmt.Printf("Extended grapheme cluster '%s': %U\n", string(egc), egc)
    }
    fmt.Printf("The text is %d extended grapheme clusters.\n", tegcs)

    // output:
    // The text is 43 bytes.
    // The text is 15 runes.
    // Extended grapheme cluster 'คุ': [U+0E04 U+0E38]
    // Extended grapheme cluster 'ณ': [U+0E13]
    // Extended grapheme cluster 'เ': [U+0E40]
    // Extended grapheme cluster 'ป็': [U+0E1B U+0E47]
    // Extended grapheme cluster 'น': [U+0E19]
    // Extended grapheme cluster 'อ': [U+0E2D]
    // Extended grapheme cluster 'ย่': [U+0E22 U+0E48]
    // Extended grapheme cluster 'า': [U+0E32]
    // Extended grapheme cluster 'ง': [U+0E07]
    // Extended grapheme cluster 'ไ': [U+0E44]
    // Extended grapheme cluster 'ร': [U+0E23]
    // Extended grapheme cluster '?': [U+003F]
    // The text is 12 extended grapheme clusters.
}
```
The above examples are a teaser introduction to this module. Peruse the **SUMMARY.txt** file to explore further.
