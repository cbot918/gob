package scanner

import (
	"fmt"
	"gob/src/pkg/ygo/yutil"
	"os"
	"strings"
)

type Scanner struct {
	data   string
	pos    int
	objArr []map[string]string
}

func NewScanner(fileName string) *Scanner {
	s := new(Scanner)
	// p, _ := filepath.Abs("/" + fileName)
	bdata, err := os.ReadFile(fileName)
	yutil.Checke(err)

	s.data = string(bdata)
	s.pos = 0

	s.Uglify()
	s.Parse()

	return s
}

func (s *Scanner) Uglify() {
	s.data = strings.Replace(s.data, " ", "", -1)
	s.data = strings.Replace(s.data, "\r", "", -1)
	s.data = strings.Replace(s.data, "\n", "", -1)
}
func (s *Scanner) Parse() []map[string]string {
	for true {
		if s.GetChar() == "[" {
		}
		if s.GetChar() == "{" {
			temp := s.GetStringUntil("}")
			s.SaveWord(temp)

		}
		if s.GetChar() == "," {
		}
		if s.GetChar() == "]" {
			yutil.Logg("finish")
			break
		}
		s.Skip(1)
	}
	return s.GetObjArr()
}

func (s *Scanner) ShowData() {
	yutil.Logg(s.data)
}

func (s *Scanner) Set(newData string) {
	s.data = newData
}

func (s *Scanner) Get() string {
	return s.data
}

func (s *Scanner) GetChar() string {
	return string(s.data[s.pos])

}

func (s *Scanner) Skip(incr int) {
	s.pos += incr
}

func (s *Scanner) ShowAllChar() {
	for pos, char := range s.data {
		fmt.Printf("%q : %d\n", char, pos)
	}
}

func (s *Scanner) ShowChar() {
	fmt.Printf("%q : %d\n", string(s.data[s.pos]), s.pos)
}

func (s *Scanner) GetStringUntil(endStr string) string {
	str := ""
	for string(s.data[s.pos]) != endStr {
		str += string(s.data[s.pos])
		s.Skip(1)
	}
	return str + "}"
}
func (s *Scanner) ShowPos() {
	yutil.Logg(s.pos)
}

func (s *Scanner) tok(word string, endSymbol string) (string, int) {
	c := 0
	temp := ""
	c += 1
	for string(word[c]) != endSymbol {
		temp += string(word[c])
		c += 1
	}
	return temp, c
}

func (s *Scanner) SaveWord(word string) {
	newWord := word
	tokArr := []string{}
	for true {
		if string(newWord[0]) == "\"" {
			tok, move := s.tok(newWord, "\"")
			newWord = yutil.TrimS(newWord, move)
			tokArr = append(tokArr, tok)
		}
		if string(newWord[0]) == "{" || string(newWord[0]) == "," || string(newWord[0]) == ":" {
		}
		if string(newWord[0]) == "}" {
			break
		}
		newWord = yutil.TrimS(newWord, 1)
	}

	obj := make(map[string]string)
	obj[tokArr[0]] = tokArr[1]
	obj[tokArr[2]] = tokArr[3]
	s.objArr = append(s.objArr, obj)
}

func (s *Scanner) GetObjArr() []map[string]string {
	return s.objArr
}
