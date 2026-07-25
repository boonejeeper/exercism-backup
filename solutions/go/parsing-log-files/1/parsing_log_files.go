package parsinglogfiles

import (
    "regexp"
    "fmt"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
 	if nil != err {
        fmt.Println(err)
        return false
    }  
    fmt.Println(re)
    results := re.MatchString(text)
    return results
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~\*=\-]*>`)
    fmt.Println(re)
    results := re.Split(text,-1)
    return results
}

func CountQuotedPasswords(lines []string) int {
    var count int
	re := regexp.MustCompile(`(?i)".*(password)+.*"`)
    for _, logLine := range lines {
        results := re.MatchString(logLine)
        if results {
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
    re := regexp.MustCompile(`end-of-line[\d]+`)
    results := re.ReplaceAllString(text,``)
    return results
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\W+([a-zA-Z0-9]+)`)
    for index, logLine := range lines {
    	results := re.FindStringSubmatch(logLine)
        if len(results) > 0 {
            modifiedLogLine := fmt.Sprintf("[USR] %s %s", results[1], logLine)
            lines[index] = modifiedLogLine
        }
    }
    return lines
}
