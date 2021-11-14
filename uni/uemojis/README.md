<h1 align="center">Unicode ® queries of emoji-related properties</h1>

**uemojis** is a Go package for querying the Unicode emoji-related properties.

For more information about these properties, consult [UAX 51](https://www.unicode.org/reports/tr51/).

Other Unicode properties are supported by other packages in this module. For a summary of Unicode properties and which functions of which packages of this module provide the corresponding support, peruse the file, **SUMMARY.txt**, in the top level directory of this module.

For example, the function, `HasEcompN(r rune) bool`, returns true only if the code point, r, has Unicode property, Emoji_Component, value, No (N).

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

import "github.com/orkvozku/go/uni/uemojis"
import "fmt"

func main() {
    text := "\U0001F9D1\U0001F3FC\u200D\U0001F9B0 \U0001F478\U0001F3FD"
    fmt.Printf("text: \"%s\": %U\n", text, []rune(text))
    for _, r := range text {
        summary := fmt.Sprintf("%U:", r)
        if uemojis.HasEmojiY(r) {
            summary += " Emoji"
        }
        if uemojis.HasEcompY(r) {
            summary += " Emoji_Component"
        }
        if uemojis.HasEmodY(r) {
            summary += " Emoji_Modifier"
        }
        if uemojis.HasEbaseY(r) {
            summary += " Emoji_Modifier_Base"
        }
        if uemojis.HasEpresY(r) {
            summary += " Emoji_Presentation"
        }
        if uemojis.HasExtpictY(r) {
            summary += " Extended_Pictographic"
        }
        if uemojis.HasRiY(r) {
            summary += " Regional_Indicator"
        }
        if r == 0x200D {
            summary += " zero width joiner (ZWJ)"
        }
        if r == 0x20 {
            summary += " space"
        }
        fmt.Printf("%s\n", summary)
    }

    // output: (emojis on first line appear differently on different platforms)
    // text: ":) :)": [U+1F9D1 U+1F3FC U+200D U+1F9B0 U+0020 U+1F478 U+1F3FD]
    // U+1F9D1: Emoji Emoji_Modifier_Base Emoji_Presentation Extended_Pictographic
    // U+1F3FC: Emoji Emoji_Component Emoji_Modifier Emoji_Presentation
    // U+200D: Emoji_Component zero width joiner (ZWJ)
    // U+1F9B0: Emoji Emoji_Component Emoji_Presentation Extended_Pictographic
    // U+0020: space
    // U+1F478: Emoji Emoji_Modifier_Base Emoji_Presentation Extended_Pictographic
    // U+1F3FD: Emoji Emoji_Component Emoji_Modifier Emoji_Presentation
}
```
## API
### General Constants
#### UnicodeVersion
UnicodeVersion is the current version of the Unicode specification that this package is intended to comply with.

UnicodeVersion = "14.0.0".
### Unicode Property Queries
#### func HasEmojiN(r rune) bool
HasEmojiN returns true if code point, r, has Unicode property, Emoji, of No (N).
#### func HasEmojiY(r rune) bool
HasEmojiY returns true if code point, r, has Unicode property, Emoji, of Yes (Y).
#### func HasEpresN(r rune) bool
HasEpresN returns true if code point, r, has Unicode property, Emoji_Presentation (EPres), of No (N).
#### func HasEpresY(r rune) bool
HasEpresY returns true if code point, r, has Unicode property, Emoji_Presentation (EPres), of Yes (Y).
#### func HasEmodN(r rune) bool
HasEmodN returns true if code point, r, has Unicode property, Emoji_Modifier (EMod), of No (N).
#### func HasEmodY(r rune) bool
HasEmodY returns true if code point, r, has Unicode property, Emoji_Modifier (EMod), of Yes (Y).
#### func HasEbaseN(r rune) bool
HasEbaseN returns true if code point, r, has Unicode property, Emoji_Modifier_Base (EBase), of No (N).
#### func HasEbaseY(r rune) bool
HasEbaseY returns true if code point, r, has Unicode property, Emoji_Modifier_Base (EBase), of Yes (Y).
#### func HasEcompN(r rune) bool
HasEcompN returns true if code point, r, has Unicode property, Emoji_Component (EComp), of No (N).
#### func HasEcompY(r rune) bool
HasEcompY returns true if code point, r, has Unicode property, Emoji_Component (EComp), of Yes (Y).
#### func HasExtpictN(r rune) bool
HasExtpictN returns true if code point, r, has Unicode property, Extended_Pictographic (ExtPict), of No (N).
#### func HasExtpictY(r rune) bool
HasExtpictY returns true if code point, r, has Unicode property, Extended_Pictographic (ExtPict), of Yes (Y).
#### func HasRiN(r rune) bool
HasRiN returns true if code point, r, has Unicode property, Regional_Indicator (RI), of No (N).
#### func HasRiY(r rune) bool
HasRiY returns true if code point, r, has Unicode property, Regional_Indicator (RI), of Yes (Y).
### func HasXidsY(r rune) bool
HasXidsY returns true if code point, r, has Unicode property, XID_Start (XIDS), of Yes (Y).
### Other Functions
## Supported Unicode Properties
<table><thead><tr><th>Property:</th><th>Value:</th><th>Function:</th></tr></thead>
<tbody>
<tr><td colspan="3">Emoji</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasEmojiN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasEmojiY(r rune) bool</td></tr>
<tr><td colspan="3">Emoji_Presentation (EPres)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasEpresN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasEpresY(r rune) bool</td></tr>
<tr><td colspan="3">Emoji_Modifier (EMod)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasEmodN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasEmodY(r rune) bool</td></tr>
<tr><td colspan="3">Emoji_Modifier_Base (EBase)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasEbaseN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasEbaseY(r rune) bool</td></tr>
<tr><td colspan="3">Emoji_Component (EComp)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasEcompN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasEcompY(r rune) bool</td></tr>
<tr><td colspan="3">Extended_Pictographic (ExtPict)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasExtpictN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasExtpictY(r rune) bool</td></tr>
<tr><td colspan="3">Regional_Indicator (RI)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasRiN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasRiY(r rune) bool</td></tr>
</tbody></table>
