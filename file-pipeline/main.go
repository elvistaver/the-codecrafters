// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Tavershima Gbaa]
// Squad:  [The Interface]
// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: [The Interface]
// ───────────────────────────────────────────
// Input line types:
/*[Lines in ALL CAPS
  Lines in all lowercase
  Lines with extra leading/trailing spaces
  Lines that are only dashes or blanks]*/

// Transformation rules (in order):
//  1. [Convert ALL CAPS lines to Title Case]
//  2. [Convert all lowercase lines to uppercase]
//  3. [Trim all leading and trailing whitespace]
//  4. [Reverse the words in any line that contains the word REVERSE]
//  5. [Remove lines that are only dashes or blanks]
//
// Output format:
//
//	[Header: yes]
//	[Line numbering format: "1"]
//	[Summary block: yes]
//
// Terminal summary fields:
//
//	[Lines read    : [number]]
//	[Lines written : [number]]
//	[Lines removed : [number]]
//	[Rules applied : [list your 5 rules]]
//
// ═══════════════════════════════════════════
package main

import (
	//"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func Title(text string) string {
	tex := strings.Fields(text)
	if text == strings.ToUpper(text) && text != strings.ToLower(text) {
		for i := range tex {
			tex[i] = strings.ToUpper(tex[i][:1]) + strings.ToLower(tex[i][1:])
		}
	}
	return strings.Join(tex, " ")
}
func ToUpper(text string) string {
	tex := strings.Fields(text)
	texx := strings.Join(tex, " ")
	return strings.ToUpper(texx)
}
func Trimspace(text string) string {
	tex := strings.Fields(text)
	texx := strings.Join(tex, " ")
	return strings.TrimSpace(texx)
}
func Reverse(text string) string {
	tex := strings.Split(text, "")
	slices.Reverse(tex)
	return strings.Join(tex, "")
}
func Remove(text string) string {
	tex := strings.Fields(text)
	texx := strings.Join(tex, " ")
	tx := strings.TrimSpace(texx)
	return strings.Trim(tx, "-")

}
func main() {

	fmt.Println("SENTINEL FIELD REPORT — PROCESSED")

	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("file does not exist", err)
		return
	}
	text0 := string(data)
	text1 := Title(text0)
	text2 := ToUpper(text1)
	text3 := Trimspace(text2)
	text4 := Reverse(text3)
	text5 := Remove(text4)

	err = os.WriteFile("output.txt", []byte(text5), 0644)

	if err != nil {
		fmt.Println("ivalid file", err)
		return
	}
	fmt.Println(text1)
}
