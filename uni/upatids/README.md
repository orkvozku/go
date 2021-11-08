<h1 align="center">Unicode​® queries of identifer and pattern syntax properties</h1>

**upatids** is a Go package for querying the Unicode identifier and pattern syntax properties.

For more information about these properties, consult [UAX 31](https://www.unicode.org/reports/tr31/).

Other Unicode properties are supported by other packages in this module. For a summary of Unicode properties and which functions of which packages of this module provide the corresponding support, peruse the file, **SUMMARY.txt**, in the top level directory of this module.

For example, the function, `HasXidcN(r rune) bool`, returns true only if the code point, r, has Unicode property, XID_Continue, value, No (N).

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

import "github.com/orkvozku/go/uni/upatids"
import "fmt"

func main() {
    text := "Hi! :)"
    for _, r := range text {
        if upatids.HasXidsY(r) {
            fmt.Printf("'%c' has XID_Start=Y\n", r)
        }
        if upatids.HasXidcY(r) {
            fmt.Printf("'%c' has XID_Continue=Y\n", r)
        }
        if upatids.HasPatwsY(r) {
            fmt.Printf("'%c' has Pattern_White_Space=Y\n", r)
        }
        if upatids.HasPatsynY(r) {
            fmt.Printf("'%c' has Pattern_Syntax=Y\n", r)
        }
    }

    // output:
    // 'H' has XID_Start=Y
    // 'H' has XID_Continue=Y
    // 'i' has XID_Start=Y
    // 'i' has XID_Continue=Y
    // '!' has Pattern_Syntax=Y
    // ' ' has Pattern_White_Space=Y
    // ':' has Pattern_Syntax=Y
    // ')' has Pattern_Syntax=Y
}
```
# API
## General Constants
### UnicodeVersion
UnicodeVersion is the current version of the Unicode specification that this package is intended to comply with.

UnicodeVersion = "14.0.0".
## Unicode Property Queries
### func HasPatwsN(r rune) bool
HasPatwsN returns true if code point, r, has Unicode property, Pattern_White_Space (Pat_WS), of No (N).
### func HasPatwsY(r rune) bool
HasPatwsY returns true if code point, r, has Unicode property, Pattern_White_Space (Pat_WS), of Yes (Y).
### func HasPatsynN(r rune) bool
HasPatsynN returns true if code point, r, has Unicode property, Pattern_Syntax (Pat_Syn), of No (N).
### func HasPatsynY(r rune) bool
HasPatsynY returns true if code point, r, has Unicode property, Pattern_Syntax (Pat_Syn), of Yes (Y).
### func HasIdcN(r rune) bool
HasIdcN returns true if code point, r, has Unicode property, ID_Continue (IDC), of No (N).
### func HasIdcY(r rune) bool
HasIdcY returns true if code point, r, has Unicode property, ID_Continue (IDC), of Yes (Y).
### func HasIdsN(r rune) bool
HasIdsN returns true if code point, r, has Unicode property, ID_Start (IDS), of No (N).
### func HasIdsY(r rune) bool
HasIdsY returns true if code point, r, has Unicode property, ID_Start (IDS), of Yes (Y).
### func HasXidcN(r rune) bool
HasXidcN returns true if code point, r, has Unicode property, XID_Continue (XIDC), of No (N).
### func HasXidcY(r rune) bool
HasXidcY returns true if code point, r, has Unicode property, XID_Continue (XIDC), of Yes (Y).
### func HasXidsN(r rune) bool
HasXidsN returns true if code point, r, has Unicode property, XID_Start (XIDS), of No (N).
### func HasXidsY(r rune) bool
HasXidsY returns true if code point, r, has Unicode property, XID_Start (XIDS), of Yes (Y).
## Other Functions
# Supported Unicode Properties
<table><thead><tr><th>Property:</th><th>Value:</th><th>Function:</th></tr></thead>
<tbody>
<tr><td colspan="3">Pattern_White_Space (Pat_WS)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasPatwsN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasPatwsY(r rune) bool</td></tr>
<tr><td colspan="3">Pattern_Syntax (Pat_Syn)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasPatsynN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasPatsynY(r rune) bool</td></tr>
<tr><td colspan="3">ID_Continue (IDC)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasIdcN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasIdcY(r rune) bool</td></tr>
<tr><td colspan="3">ID_Start (IDS)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasIdsN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasIdsY(r rune) bool</td></tr>
<tr><td colspan="3">XID_Continue (XIDC)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasXidcN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasXidcY(r rune) bool</td></tr>
<tr><td colspan="3">XID_Start (XIDS)</td></tr>
<tr><td>&nbsp;</td><td>No (N, F, False)</td><td>HasXidsN(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Yes (Y, T, True)</td><td>HasXidsY(r rune) bool</td></tr>
</tbody></table>
