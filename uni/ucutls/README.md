<h1 align="center">support utilities for github.com/orkvozku/go/uni.</h1>

**ucutls** is a Go package supporting other packages within this module.

## Example

If you are new to the Go language, consider perusing the documentation at [golang.org](https://golang.org/doc/) before proceeding.

Download and install it:

```sh
go get github.com/orkvozku/go/uni

```
Example: Wrapping ubasic.NextExtGraphemeClusterBreakLen in a io.RuneReader to ucutls.RunesReader filter for parsing extended grapheme cluster boundaries.

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
