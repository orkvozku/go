<h1 align="center">Unicode® queries of segmentation properties</h1>

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
# API
## General Constants
### UnicodeVersion
UnicodeVersion is the current version of the Unicode specification that this package is intended to comply with.

UnicodeVersion = "14.0.0".
## Unicode Property Queries
### func HasGcbCn(r rune) bool
HasGcbCn returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of Control (CN).
### func HasGcbCr(r rune) bool
HasGcbCr returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of CR.
### func HasGcbEx(r rune) bool
HasGcbEx returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of Extend (EX).
### func HasGcbL(r rune) bool
HasGcbL returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of L.
### func HasGcbLf(r rune) bool
HasGcbLf returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of LF.
### func HasGcbLv(r rune) bool
HasGcbLv returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of LV.
### func HasGcbLvt(r rune) bool
HasGcbLvt returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of LVT.
### func HasGcbPp(r rune) bool
HasGcbPp returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of Prepend (PP).
### func HasGcbRi(r rune) bool
HasGcbRi returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of Regional_Indicator (RI).
### func HasGcbSm(r rune) bool
HasGcbSm returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of SpacingMark (SM).
### func HasGcbT(r rune) bool
HasGcbT returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of T.
### func HasGcbV(r rune) bool
HasGcbV returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of V.
### func HasGcbXx(r rune) bool
HasGcbXx returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of Other (XX).
### func HasGcbZwj(r rune) bool
HasGcbZwj returns true if code point, r, has Unicode property, Grapheme_Cluster_Break (GCB), of ZWJ.
### func HasExtpictN(r rune) bool
HasExtpictN returns true if code point, r, has Unicode property, Extended_Pictographic (ExtPict), of No (N).
### func HasExtpictY(r rune) bool
HasExtpictY returns true if code point, r, has Unicode property, Extended_Pictographic (ExtPict), of Yes (Y).
### func HasLbAi(r rune) bool
HasLbAi returns true if code point, r, has Unicode property, Line_Break (lb), of Ambiguous (AI).
### func HasLbAl(r rune) bool
HasLbAl returns true if code point, r, has Unicode property, Line_Break (lb), of Alphabetic (AL).
### func HasLbB2(r rune) bool
HasLbB2 returns true if code point, r, has Unicode property, Line_Break (lb), of Break_Both (B2).
### func HasLbBa(r rune) bool
HasLbBa returns true if code point, r, has Unicode property, Line_Break (lb), of Break_After (BA).
### func HasLbBb(r rune) bool
HasLbBb returns true if code point, r, has Unicode property, Line_Break (lb), of Break_Before (BB).
### func HasLbBk(r rune) bool
HasLbBk returns true if code point, r, has Unicode property, Line_Break (lb), of Mandatory_Break (BK).
### func HasLbCb(r rune) bool
HasLbCb returns true if code point, r, has Unicode property, Line_Break (lb), of Contingent_Break (CB).
### func HasLbCj(r rune) bool
HasLbCj returns true if code point, r, has Unicode property, Line_Break (lb), of Conditional_Japanese_Starter (CJ).
### func HasLbCl(r rune) bool
HasLbCl returns true if code point, r, has Unicode property, Line_Break (lb), of Close_Punctuation (CL).
### func HasLbCm(r rune) bool
HasLbCm returns true if code point, r, has Unicode property, Line_Break (lb), of Combining_Mark (CM).
### func HasLbCp(r rune) bool
HasLbCp returns true if code point, r, has Unicode property, Line_Break (lb), of Close_Parenthesis (CP).
### func HasLbCr(r rune) bool
HasLbCr returns true if code point, r, has Unicode property, Line_Break (lb), of Carriage_Return (CR).
### func HasLbEb(r rune) bool
HasLbEb returns true if code point, r, has Unicode property, Line_Break (lb), of E_Base (EB).
### func HasLbEm(r rune) bool
HasLbEm returns true if code point, r, has Unicode property, Line_Break (lb), of E_Modifier (EM).
### func HasLbEx(r rune) bool
HasLbEx returns true if code point, r, has Unicode property, Line_Break (lb), of Exclamation (EX).
### func HasLbGl(r rune) bool
HasLbGl returns true if code point, r, has Unicode property, Line_Break (lb), of Glue (GL).
### func HasLbH2(r rune) bool
HasLbH2 returns true if code point, r, has Unicode property, Line_Break (lb), of H2.
### func HasLbH3(r rune) bool
HasLbH3 returns true if code point, r, has Unicode property, Line_Break (lb), of H3.
### func HasLbHl(r rune) bool
HasLbHl returns true if code point, r, has Unicode property, Line_Break (lb), of Hebrew_Letter (HL).
### func HasLbHy(r rune) bool
HasLbHy returns true if code point, r, has Unicode property, Line_Break (lb), of Hyphen (HY).
### func HasLbId(r rune) bool
HasLbId returns true if code point, r, has Unicode property, Line_Break (lb), of Ideographic (ID).
### func HasLbIn(r rune) bool
HasLbIn returns true if code point, r, has Unicode property, Line_Break (lb), of Inseparable (IN).
### func HasLbIs(r rune) bool
HasLbIs returns true if code point, r, has Unicode property, Line_Break (lb), of Infix_Numeric (IS).
### func HasLbJl(r rune) bool
HasLbJl returns true if code point, r, has Unicode property, Line_Break (lb), of JL.
### func HasLbJt(r rune) bool
HasLbJt returns true if code point, r, has Unicode property, Line_Break (lb), of JT.
### func HasLbJv(r rune) bool
HasLbJv returns true if code point, r, has Unicode property, Line_Break (lb), of JV.
### func HasLbLf(r rune) bool
HasLbLf returns true if code point, r, has Unicode property, Line_Break (lb), of Line_Feed (LF).
### func HasLbNl(r rune) bool
HasLbNl returns true if code point, r, has Unicode property, Line_Break (lb), of Next_Line (NL).
### func HasLbNs(r rune) bool
HasLbNs returns true if code point, r, has Unicode property, Line_Break (lb), of Nonstarter (NS).
### func HasLbNu(r rune) bool
HasLbNu returns true if code point, r, has Unicode property, Line_Break (lb), of Numeric (NU).
### func HasLbOp(r rune) bool
HasLbOp returns true if code point, r, has Unicode property, Line_Break (lb), of Open_Punctuation (OP).
### func HasLbPo(r rune) bool
HasLbPo returns true if code point, r, has Unicode property, Line_Break (lb), of Postfix_Numeric (PO).
### func HasLbPr(r rune) bool
HasLbPr returns true if code point, r, has Unicode property, Line_Break (lb), of Prefix_Numeric (PR).
### func HasLbQu(r rune) bool
HasLbQu returns true if code point, r, has Unicode property, Line_Break (lb), of Quotation (QU).
### func HasLbRi(r rune) bool
HasLbRi returns true if code point, r, has Unicode property, Line_Break (lb), of Regional_Indicator (RI).
### func HasLbSa(r rune) bool
HasLbSa returns true if code point, r, has Unicode property, Line_Break (lb), of Complex_Context (SA).
### func HasLbSg(r rune) bool
HasLbSg returns true if code point, r, has Unicode property, Line_Break (lb), of Surrogate (SG).
### func HasLbSp(r rune) bool
HasLbSp returns true if code point, r, has Unicode property, Line_Break (lb), of Space (SP).
### func HasLbSy(r rune) bool
HasLbSy returns true if code point, r, has Unicode property, Line_Break (lb), of Break_Symbols (SY).
### func HasLbWj(r rune) bool
HasLbWj returns true if code point, r, has Unicode property, Line_Break (lb), of Word_Joiner (WJ).
### func HasLbXx(r rune) bool
HasLbXx returns true if code point, r, has Unicode property, Line_Break (lb), of Unknown (XX).
### func HasLbZw(r rune) bool
HasLbZw returns true if code point, r, has Unicode property, Line_Break (lb), of ZWSpace (ZW).
### func HasLbZwj(r rune) bool
HasLbZwj returns true if code point, r, has Unicode property, Line_Break (lb), of ZWJ.
### func HasWbCr(r rune) bool
HasWbCr returns true if code point, r, has Unicode property, Word_Break (WB), of CR.
### func HasWbDq(r rune) bool
HasWbDq returns true if code point, r, has Unicode property, Word_Break (WB), of Double_Quote (DQ).
### func HasWbEx(r rune) bool
HasWbEx returns true if code point, r, has Unicode property, Word_Break (WB), of ExtendNumLet (EX).
### func HasWbExtend(r rune) bool
HasWbExtend returns true if code point, r, has Unicode property, Word_Break (WB), of Extend.
### func HasWbFo(r rune) bool
HasWbFo returns true if code point, r, has Unicode property, Word_Break (WB), of Format (FO).
### func HasWbHl(r rune) bool
HasWbHl returns true if code point, r, has Unicode property, Word_Break (WB), of Hebrew_Letter (HL).
### func HasWbKa(r rune) bool
HasWbKa returns true if code point, r, has Unicode property, Word_Break (WB), of Katakana (KA).
### func HasWbLe(r rune) bool
HasWbLe returns true if code point, r, has Unicode property, Word_Break (WB), of ALetter (LE).
### func HasWbLf(r rune) bool
HasWbLf returns true if code point, r, has Unicode property, Word_Break (WB), of LF.
### func HasWbMb(r rune) bool
HasWbMb returns true if code point, r, has Unicode property, Word_Break (WB), of MidNumLet (MB).
### func HasWbMl(r rune) bool
HasWbMl returns true if code point, r, has Unicode property, Word_Break (WB), of MidLetter (ML).
### func HasWbMn(r rune) bool
HasWbMn returns true if code point, r, has Unicode property, Word_Break (WB), of MidNum (MN).
### func HasWbNl(r rune) bool
HasWbNl returns true if code point, r, has Unicode property, Word_Break (WB), of Newline (NL).
### func HasWbNu(r rune) bool
HasWbNu returns true if code point, r, has Unicode property, Word_Break (WB), of Numeric (NU).
### func HasWbRi(r rune) bool
HasWbRi returns true if code point, r, has Unicode property, Word_Break (WB), of Regional_Indicator (RI).
### func HasWbSq(r rune) bool
HasWbSq returns true if code point, r, has Unicode property, Word_Break (WB), of Single_Quote (SQ).
### func HasWbWsegspace(r rune) bool
HasWbWsegspace returns true if code point, r, has Unicode property, Word_Break (WB), of WSegSpace.
### func HasWbXx(r rune) bool
HasWbXx returns true if code point, r, has Unicode property, Word_Break (WB), of Other (XX).
### func HasWbZwj(r rune) bool
HasWbZwj returns true if code point, r, has Unicode property, Word_Break (WB), of ZWJ.
### func HasSbAt(r rune) bool
HasSbAt returns true if code point, r, has Unicode property, Sentence_Break (SB), of ATerm (AT).
### func HasSbCl(r rune) bool
HasSbCl returns true if code point, r, has Unicode property, Sentence_Break (SB), of Close (CL).
### func HasSbCr(r rune) bool
HasSbCr returns true if code point, r, has Unicode property, Sentence_Break (SB), of CR.
### func HasSbEx(r rune) bool
HasSbEx returns true if code point, r, has Unicode property, Sentence_Break (SB), of Extend (EX).
### func HasSbFo(r rune) bool
HasSbFo returns true if code point, r, has Unicode property, Sentence_Break (SB), of Format (FO).
### func HasSbLe(r rune) bool
HasSbLe returns true if code point, r, has Unicode property, Sentence_Break (SB), of OLetter (LE).
### func HasSbLf(r rune) bool
HasSbLf returns true if code point, r, has Unicode property, Sentence_Break (SB), of LF.
### func HasSbLo(r rune) bool
HasSbLo returns true if code point, r, has Unicode property, Sentence_Break (SB), of Lower (LO).
### func HasSbNu(r rune) bool
HasSbNu returns true if code point, r, has Unicode property, Sentence_Break (SB), of Numeric (NU).
### func HasSbSc(r rune) bool
HasSbSc returns true if code point, r, has Unicode property, Sentence_Break (SB), of SContinue (SC).
### func HasSbSe(r rune) bool
HasSbSe returns true if code point, r, has Unicode property, Sentence_Break (SB), of Sep (SE).
### func HasSbSp(r rune) bool
HasSbSp returns true if code point, r, has Unicode property, Sentence_Break (SB), of Sp (SP).
### func HasSbSt(r rune) bool
HasSbSt returns true if code point, r, has Unicode property, Sentence_Break (SB), of STerm (ST).
### func HasSbUp(r rune) bool
HasSbUp returns true if code point, r, has Unicode property, Sentence_Break (SB), of Upper (UP).
### func HasSbXx(r rune) bool
HasSbXx returns true if code point, r, has Unicode property, Sentence_Break (SB), of Other (XX).
## Other Functions
# Supported Unicode Properties
<table><thead><tr><th>Property:</th><th>Value:</th><th>Function:</th></tr></thead>
<tbody>
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
<tr><td colspan="3">Line_Break (lb)</td></tr>
<tr><td>&nbsp;</td><td>Ambiguous (AI)</td><td>HasLbAi(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Alphabetic (AL)</td><td>HasLbAl(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Break_Both (B2)</td><td>HasLbB2(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Break_After (BA)</td><td>HasLbBa(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Break_Before (BB)</td><td>HasLbBb(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Mandatory_Break (BK)</td><td>HasLbBk(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Contingent_Break (CB)</td><td>HasLbCb(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Conditional_Japanese_Starter (CJ)</td><td>HasLbCj(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Close_Punctuation (CL)</td><td>HasLbCl(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Combining_Mark (CM)</td><td>HasLbCm(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Close_Parenthesis (CP)</td><td>HasLbCp(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Carriage_Return (CR)</td><td>HasLbCr(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>E_Base (EB)</td><td>HasLbEb(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>E_Modifier (EM)</td><td>HasLbEm(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Exclamation (EX)</td><td>HasLbEx(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Glue (GL)</td><td>HasLbGl(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>H2</td><td>HasLbH2(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>H3</td><td>HasLbH3(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Hebrew_Letter (HL)</td><td>HasLbHl(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Hyphen (HY)</td><td>HasLbHy(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Ideographic (ID)</td><td>HasLbId(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Inseparable (IN, Inseperable)</td><td>HasLbIn(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Infix_Numeric (IS)</td><td>HasLbIs(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>JL</td><td>HasLbJl(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>JT</td><td>HasLbJt(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>JV</td><td>HasLbJv(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Line_Feed (LF)</td><td>HasLbLf(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Next_Line (NL)</td><td>HasLbNl(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Nonstarter (NS)</td><td>HasLbNs(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Numeric (NU)</td><td>HasLbNu(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Open_Punctuation (OP)</td><td>HasLbOp(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Postfix_Numeric (PO)</td><td>HasLbPo(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Prefix_Numeric (PR)</td><td>HasLbPr(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Quotation (QU)</td><td>HasLbQu(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Regional_Indicator (RI)</td><td>HasLbRi(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Complex_Context (SA)</td><td>HasLbSa(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Surrogate (SG)</td><td>HasLbSg(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Space (SP)</td><td>HasLbSp(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Break_Symbols (SY)</td><td>HasLbSy(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Word_Joiner (WJ)</td><td>HasLbWj(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Unknown (XX)</td><td>HasLbXx(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>ZWSpace (ZW)</td><td>HasLbZw(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>ZWJ</td><td>HasLbZwj(r rune) bool</td></tr>
<tr><td colspan="3">Word_Break (WB)</td></tr>
<tr><td>&nbsp;</td><td>CR</td><td>HasWbCr(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Double_Quote (DQ)</td><td>HasWbDq(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>ExtendNumLet (EX)</td><td>HasWbEx(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Extend</td><td>HasWbExtend(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Format (FO)</td><td>HasWbFo(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Hebrew_Letter (HL)</td><td>HasWbHl(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Katakana (KA)</td><td>HasWbKa(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>ALetter (LE)</td><td>HasWbLe(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>LF</td><td>HasWbLf(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>MidNumLet (MB)</td><td>HasWbMb(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>MidLetter (ML)</td><td>HasWbMl(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>MidNum (MN)</td><td>HasWbMn(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Newline (NL)</td><td>HasWbNl(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Numeric (NU)</td><td>HasWbNu(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Regional_Indicator (RI)</td><td>HasWbRi(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Single_Quote (SQ)</td><td>HasWbSq(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>WSegSpace</td><td>HasWbWsegspace(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Other (XX)</td><td>HasWbXx(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>ZWJ</td><td>HasWbZwj(r rune) bool</td></tr>
<tr><td colspan="3">Sentence_Break (SB)</td></tr>
<tr><td>&nbsp;</td><td>ATerm (AT)</td><td>HasSbAt(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Close (CL)</td><td>HasSbCl(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>CR</td><td>HasSbCr(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Extend (EX)</td><td>HasSbEx(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Format (FO)</td><td>HasSbFo(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>OLetter (LE)</td><td>HasSbLe(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>LF</td><td>HasSbLf(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Lower (LO)</td><td>HasSbLo(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Numeric (NU)</td><td>HasSbNu(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>SContinue (SC)</td><td>HasSbSc(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Sep (SE)</td><td>HasSbSe(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Sp (SP)</td><td>HasSbSp(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>STerm (ST)</td><td>HasSbSt(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Upper (UP)</td><td>HasSbUp(r rune) bool</td></tr>
<tr><td>&nbsp;</td><td>Other (XX)</td><td>HasSbXx(r rune) bool</td></tr>
</tbody></table>
