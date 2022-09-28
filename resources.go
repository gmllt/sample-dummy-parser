package main

import (
	"bufio"
	"encoding/json"
	"os"
	"regexp"
	"strings"
)

type Log struct {
	Original string
}

func FromStdin() Log {
	var stdin []byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdin = append(stdin, scanner.Bytes()...)
	}
	return Log{Original: string(stdin)}
}

func (l *Log) Correct() string {
	reg, err := regexp.Compile(`"message": "(.*)", "log_reqid"`)
	if err != nil {
		return ""
	}
	if reg.Match([]byte(l.Original)) {
		stringSubMatch := reg.FindStringSubmatch(l.Original)
		text := stringSubMatch[1]
		newText := strings.Replace(text, `"`, `\"`, -1)
		l.Original = strings.Replace(l.Original, text, newText, -1)
	}
	var unmarshalled map[string]interface{}
	err = json.Unmarshal([]byte(l.Original), &unmarshalled)
	if err != nil {
		return ""
	}
	marshalled, err := json.Marshal(unmarshalled)
	if err != nil {
		return ""
	}
	return string(marshalled)

}

func (l *Log) Unmarshall() map[string]interface{} {
	var result map[string]interface{}
	_ = json.Unmarshal([]byte(l.Original), &result)
	return result
}
