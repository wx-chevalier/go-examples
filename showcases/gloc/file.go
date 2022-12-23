package gocloc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ClocFile struct {
	Code     int32  `xml:"code,attr"`
	Comments int32  `xml:"comment,attr"`
	Blanks   int32  `xml:"blank,attr"`
	Name     string `xml:"name,attr"`
	Lang     string `xml:"language,attr"`
}

type ClocFiles []ClocFile

func (cf ClocFiles) Len() int {
	return len(cf)
}
func (cf ClocFiles) Swap(i, j int) {
	cf[i], cf[j] = cf[j], cf[i]
}
func (cf ClocFiles) Less(i, j int) bool {
	if cf[i].Code == cf[j].Code {
		return cf[i].Name < cf[j].Name
	}
	return cf[i].Code > cf[j].Code
}

func AnalyzeFile(filename string, language *Language, opts *ClocOptions) *ClocFile {
	if opts.Debug {
		fmt.Printf("filename=%v\n", filename)
	}

	clocFile := &ClocFile{
		Name: filename,
	}

	fp, err := os.Open(filename)
	if err != nil {
		return clocFile // ignore error
	}
	defer fp.Close()

	isFirstLine := true
	isInComments := false
	isInCommentsSame := false
	buf := getByteSlice()
	defer putByteSlice(buf)
	scanner := bufio.NewScanner(fp)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		lineOrg := scanner.Text()
		line := strings.TrimSpace(lineOrg)

		if len(strings.TrimSpace(line)) == 0 {
			clocFile.Blanks++
			if opts.Debug {
				fmt.Printf("[BLNK,cd:%d,cm:%d,bk:%d,iscm:%v] %s\n",
					clocFile.Code, clocFile.Comments, clocFile.Blanks, isInComments, lineOrg)
			}
			continue
		}

		// shebang line is 'code'
		if isFirstLine && strings.HasPrefix(line, "#!") {
			clocFile.Code++
			isFirstLine = false
			if opts.Debug {
				fmt.Printf("[CODE,cd:%d,cm:%d,bk:%d,iscm:%v] %s\n",
					clocFile.Code, clocFile.Comments, clocFile.Blanks, isInComments, lineOrg)
			}
			continue
		}

		if len(language.lineComments) > 0 {
			isSingleComment := false
			if isFirstLine {
				line = trimBOM(line)
			}
			for _, singleComment := range language.lineComments {
				if strings.HasPrefix(line, singleComment) {
					clocFile.Comments++
					isSingleComment = true
					break
				}
			}
			if isSingleComment {
				if opts.Debug {
					fmt.Printf("[COMM,cd:%d,cm:%d,bk:%d,iscm:%v] %s\n",
						clocFile.Code, clocFile.Comments, clocFile.Blanks, isInComments, lineOrg)
				}
				continue
			}
		}

		isCode := false
		multiLine := ""
		multiLineEnd := ""
		for i := range language.multiLines {
			multiLine = language.multiLines[i][0]
			multiLineEnd = language.multiLines[i][1]
			if multiLine != "" {
				if strings.HasPrefix(line, multiLine) {
					isInComments = true
				} else if strings.HasSuffix(line, multiLineEnd) {
					isInComments = true
				} else if containComments(line, multiLine, multiLineEnd) {
					isInComments = true
					if (multiLine != multiLineEnd) &&
						(strings.HasSuffix(line, multiLine) || strings.HasPrefix(line, multiLineEnd)) {
						clocFile.Code++
						isCode = true
						if opts.Debug {
							fmt.Printf("[CODE,cd:%d,cm:%d,bk:%d,iscm:%v] %s\n",
								clocFile.Code, clocFile.Comments, clocFile.Blanks, isInComments, lineOrg)
						}
						continue
					}
				}
				if isInComments {
					break
				}
			}
		}

		if isInComments && isCode {
			continue
		}

		if isInComments {
			if multiLine == multiLineEnd {
				if strings.Count(line, multiLineEnd) == 2 {
					isInComments = false
					isInCommentsSame = false
				} else if strings.HasPrefix(line, multiLineEnd) ||
					strings.HasSuffix(line, multiLineEnd) {
					if isInCommentsSame {
						isInComments = false
					}
					isInCommentsSame = !isInCommentsSame
				}
			} else {
				if strings.Contains(line, multiLineEnd) {
					isInComments = false
				}
			}
			clocFile.Comments++
			if opts.Debug {
				fmt.Printf("[COMM,cd:%d,cm:%d,bk:%d,iscm:%v,iscms:%v] %s\n",
					clocFile.Code, clocFile.Comments, clocFile.Blanks, isInComments, isInCommentsSame, lineOrg)
			}
			continue
		}

		clocFile.Code++
		if opts.Debug {
			fmt.Printf("[CODE,cd:%d,cm:%d,bk:%d,iscm:%v] %s\n",
				clocFile.Code, clocFile.Comments, clocFile.Blanks, isInComments, lineOrg)
		}
	}

	return clocFile
}
