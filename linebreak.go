package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"golang.org/x/exp/slices"
)

var SpecialCase []rune = []rune{'ー', ',', '，', '、', '.', '。', '．', 'ぁ', 'ぃ', 'ぅ', 'ぇ', 'ぉ', 'っ', 'ゃ', 'ゅ', 'ょ', 'ァ', 'ィ', 'ゥ', 'ェ', 'ォ', 'ャ', 'ュ', 'ョ'}

func splitText(text string, size int) []string {
	lines := make([]string, 0)
	line := "　"
	lineLen := 0
	for _, char := range text {
		charWidth := utf8.RuneLen(char)
		if charWidth > 2 {
			charWidth = 2
		}
		match := slices.Index(SpecialCase, char)
		if lineLen+charWidth > size && match < 0 {
			lines = append(lines, line)
			line = "　"
			lineLen = 0
		}
		line += string(char)
		lineLen += charWidth
	}
	lines = append(lines, line)
	return lines
}

func main() {
	var n = flag.Int("n", 70, "number of letters per line")
	flag.Parse()
	numChars := n
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}
		text := scanner.Text()
		lines := splitText(text, *numChars)
		fmt.Println(strings.Join(lines, "\n"))
	}
}
