package main
 
import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)
 
func main() {
	log.SetFlags(0)
	log.SetPrefix("markov: ")
	input := flag.String("in", "data.txt", "input file")
	n := flag.Int("n", 2, "number of words to use as prefix")
	words := flag.Int("words", 8, "Minimum number of words.")
	flag.Parse()
 
	rand.Seed(time.Now().UnixNano())
 
	m, err := NewMarkovFromFile(*input, *n)
	if err != nil {
		log.Fatal(err)
	}
 
	err = m.Output(os.Stdout, *words)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n -- FaKe Software Quote \n")
}

// Markov is a Markov chain text generator.
type Markov struct {
	n           int	// Window length
	capitalized int // number of keys that start capitalized
	suffix      map[string][]string // Map of prefixString -> [poosible suffix words]
}
 
// NewMarkovFromFile initializes the Markov text generator with window 'n' from the contents of 'filename'.
func NewMarkovFromFile(filename string, n int) (*Markov, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewMarkov(f, n)
}
 
// NewMarkov initializes the Markov text generator with window `n` from the contents of `r`.
func NewMarkov(r io.Reader, n int) (*Markov, error) {
	m := &Markov{
		n:      n,
		suffix: make(map[string][]string),
	}
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	window := make([]string, 0, n)
	for sc.Scan() {
		word := sc.Text()
		if len(window) == n {
			prefix := strings.Join(window, " ")
			m.suffix[prefix] = append(m.suffix[prefix], word)
			if isCapitalized(prefix) {
				m.capitalized++
			}
		}
		window = appendMax(n, window, word)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return m, nil
}
 
// Output writes generated text of approximately `n` words to `w`.
func (m *Markov) Output(w io.Writer, n int) error {
	bw := bufio.NewWriter(w)
 
	i := rand.Intn(m.capitalized)

	var prefix string
	for prefix = range m.suffix {
		if !isCapitalized(prefix) {
			continue
		}
		if i == 0 {
			break
		}
		i--
	}
 
	bw.WriteString(prefix)
	prefixWords := strings.Split(prefix, " ")
	n -= len(prefixWords)
 
	for {
		suffixChoices := m.suffix[prefix]
		if len(suffixChoices) == 0 {
			break
		}
		i = rand.Intn(len(suffixChoices))
		suffix := suffixChoices[i]

		bw.WriteByte(' ')
		if _, err := bw.WriteString(suffix); err != nil {
			break
		}
		n--
		if n < 0 && isSentenceEnd(suffix) {
			break
		}
 
		prefixWords = appendMax(m.n, prefixWords, suffix)
		prefix = strings.Join(prefixWords, " ")
	}
	return bw.Flush()
}
 
func isCapitalized(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}
 
func isSentenceEnd(s string) bool {
	r, _ := utf8.DecodeLastRuneInString(s)
	return r == '.' || r == '?' || r == '!'
}
 
func appendMax(max int, slice []string, value string) []string {
	if len(slice)+1 > max {
		n := copy(slice, slice[1:])
		slice = slice[:n]
	}
	return append(slice, value)
}