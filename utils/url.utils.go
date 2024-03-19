package utils

import (
	"log"
	"regexp"
	"strings"
)

func GetDomainName(url string) string {
	m := regexp.MustCompile(`\.?([^.]*.com)`)
	matches := m.FindStringSubmatch(url)
	log.Println(matches)
	return strings.Title(strings.Split(matches[1], ".")[0])
}
