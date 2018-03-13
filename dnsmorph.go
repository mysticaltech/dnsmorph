package main

import ("flag"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings")

// program version
const version = "1.0.0-dev3"

var (
	g = color.New(color.FgGreen)
	y = color.New(color.FgYellow)
	r = color.New(color.FgRed)
	b = color.New(color.FgBlue)
	domain = flag.String("d", "", "domain")
	verbose = flag.Bool("v", false, "enable verbosity")
	)

// sets up command-line arguments
func setup(){
	y.Printf("DNSMORPH")
	fmt.Printf(" v.%s\n\n", version)

	flag.Parse()

	if *domain == "" {
		r.Printf("please supply a domain\n\n")
		os.Exit(1)
	}
}

// returns a count of characters in a word
func countChar(word string) map[rune]int {
	count := make(map[rune]int)
	for _, r := range []rune(word){
		count[r]++
	}
	return count
}

// performs a hyphenation attack
func hyphenationAttack(domain string){
	
	tld := strings.Split(domain, ".")[1]
	dom := strings.Split(domain, ".")[0]

	for i := 1; i < len(dom); i++ {
		if (rune(dom[i]) != '-' || rune(dom[i]) != '.') && (rune(dom[i-1]) != '-' || rune(dom[i-1]) != '.') {
			if *verbose == false {
				fmt.Println(dom[:i] + "-" + dom[i:] + "." + tld)
			} else if *verbose == true {
				fmt.Println("hyphenation:  " + dom[:i] + "-" + dom[i:] + "." + tld)
			}
		}
	}
}

// performs a bitsquat permutation attack
func bitsquattingAttack(domain string) {

	tld := strings.Split(domain, ".")[1]
	dom := strings.Split(domain, ".")[0]
	masks := []int32{1, 2, 4, 8, 16, 32, 64, 128}

	for i, c := range dom {
		for m := range masks {
			b := rune(int(c) ^ m)
			o := int(b)
			if (o >= 48 && o <= 57) || (o >= 97 && o <= 122) || o == 45 {
				if *verbose == false {
					fmt.Println(dom[:i]+ string(b) + dom[i+1:] + "."+ tld)
				} else if *verbose == true {
					fmt.Println("bitsquatting: " + dom[:i]+ string(b) + dom[i+1:] + "."+ tld)
				}
			}
		}
	}
}

// performs a homograph permutation attack
func homographAttack(domain string){
	glyphs := map[rune][]rune{
		'a': []rune{'à', 'á', 'â', 'ã', 'ä', 'å', 'ɑ', 'а', 'ạ', 'ǎ', 'ă', 'ȧ','α','ａ'},
		'b': []rune{'d', 'ʙ', 'Ь', 'ɓ', 'Б', 'ß', 'β', 'ᛒ'}, // 'lb', 'ib', 'b̔'
		'c': []rune{'ϲ', 'с', 'ƈ', 'ċ', 'ć', 'ç', 'ｃ'},
		'd': []rune{'b', 'ԁ', 'ժ', 'ɗ', 'đ'}, // 'cl', 'dl', 'di'
		'e': []rune{'é', 'ê', 'ë', 'ē', 'ĕ', 'ě', 'ė', 'е', 'ẹ', 'ę', 'є', 'ϵ', 'ҽ'},
		'f': []rune{'Ϝ', 'ƒ', 'Ғ'},
		'g': []rune{'q', 'ɢ', 'ɡ', 'Ԍ', 'Ԍ', 'ġ', 'ğ', 'ց', 'ǵ', 'ģ'},
		'h': []rune{'һ', 'հ', 'Ꮒ', 'н'}, // 'lh', 'ih'
		'i': []rune{'1', 'l', 'Ꭵ', 'í', 'ï', 'ı', 'ɩ', 'ι', 'ꙇ', 'ǐ', 'ĭ'},
		'j': []rune{'ј', 'ʝ', 'ϳ', 'ɉ'},
		'k': []rune{'κ', 'ⲕ', 'κ'}, // 'lk', 'ik', 'lc'
		'l': []rune{'1', 'i', 'ɫ', 'ł'},
		'm': []rune{'n', 'ṃ', 'ᴍ', 'м', 'ɱ'}, // 'nn', 'rn', 'rr'
		'n': []rune{'m', 'r', 'ń'},
		'o': []rune{'0', 'Ο', 'ο', 'О', 'о', 'Օ', 'ȯ', 'ọ', 'ỏ', 'ơ', 'ó', 'ö', 'ӧ', 'ｏ'},
		'p': []rune{'ρ', 'р', 'ƿ', 'Ϸ', 'Þ'},
		'q': []rune{'g', 'զ', 'ԛ', 'գ', 'ʠ'},
		'r': []rune{'ʀ', 'Г', 'ᴦ', 'ɼ', 'ɽ'},
		's': []rune{'Ⴝ', 'Ꮪ', 'ʂ', 'ś', 'ѕ'},
		't': []rune{'τ', 'т', 'ţ'},
		'u': []rune{'μ', 'υ', 'Ս', 'ս', 'ц', 'ᴜ', 'ǔ', 'ŭ'},
		'v': []rune{'ѵ', 'ν'}, // 'v̇'
		'w': []rune{'ѡ', 'ա', 'ԝ'}, // 'vv'
		'x': []rune{'х', 'ҳ'}, // 'ẋ'
		'y': []rune{'ʏ', 'γ', 'у', 'Ү', 'ý'},
		'z': []rune{'ʐ', 'ż', 'ź', 'ʐ', 'ᴢ'},
	}
	// set local variables
	doneCount := make(map[rune]bool)
	tld := strings.Split(domain, ".")[1]
	dom := strings.Split(domain, ".")[0]
	runes := []rune(dom)
	count := countChar(dom)

	for i, char := range runes {
		index := i
		index++
		charGlyph := glyphs[char]
		// perform attack against single character
		for _, glyph := range charGlyph {
			if *verbose == false {
				fmt.Println(string(runes[:i]) + string(glyph) + string(runes[index:]) + "." + tld)
			} else if *verbose == true {
				fmt.Println("homograph:    " + string(runes[:i]) + string(glyph) + string(runes[index:]) + "." + tld)
			}
		}
		// determine if character is a duplicate
		// and if the attack has already been performed
		// against all characters at the same time
		if (count[char] > 1 && doneCount[char]!= true) {
			doneCount[char] = true
			for _, glyph := range charGlyph {
				str := strings.Replace(dom, string(char), string(glyph), -1)
				if *verbose == false {
					fmt.Println(str + "." + tld)
				} else if *verbose == true {
					fmt.Println("homograph:    " + str + "." + tld)
				}
			}
		}
	}
}

// main program entry point
func main(){
	setup()
	homographAttack(*domain)
	hyphenationAttack(*domain)
	bitsquattingAttack(*domain)
}