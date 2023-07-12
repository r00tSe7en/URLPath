package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/edoardottt/golazy"
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")
	levelsPtr := flag.Int("l", -1, "Number of levels to print. -1 to print all levels.")

	flag.Parse()

	if *helpPtr {
		help()
	}

	input := ScanTargets()
	output := GetPaths(golazy.RemoveDuplicateValues(input), *levelsPtr)

	for _, elem := range output {
		fmt.Println(elem)
	}
}

// help shows the usage.
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout all the unique paths (at any level).
	$> cat input | URLPath -l 2 
	-l is number of levels to print. -1 to print all levels.`

	fmt.Println()
	fmt.Println(usage)
	fmt.Println()
	os.Exit(0)
}

// ScanTargets return the array of elements
// taken as input on stdin.
func ScanTargets() []string {
	var result []string

	// accept domains on stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := sc.Text()
		result = append(result, domain)
	}

	return result
}

// GetPaths.
func GetPaths(s []string, levels int) []string {
	var result []string

	for _, elem := range s {
		if len(elem) != 0 {
			var paths []string

			if golazy.HasProtocol(elem) {
				if GetPath(elem) != "" {
					paths = GetAllLevelsPaths(GetPath(elem), elem, levels)
				}
			}

			if len(paths) != 0 {
				result = append(result, paths...)
			}
		}
	}

	return golazy.RemoveDuplicateValues(result)
}

// GetPath.
func GetPath(input string) string {
	u, err := url.Parse(input)
	if err != nil {
		return ""
	}

	if len(u.Path) > 1 {
		return u.Path[1:]
	} else {
		return ""
	}
}

// GetAllLevelsPaths.
func GetAllLevelsPaths(input string, url string, levels int) []string {
	if input == "" {
		return []string{}
	}

	var result []string

	if input[len(input)-1] != '/' {
		input += "/"
	}

	var elems = strings.Split(input, "/")
	var tmpurl = strings.Split(url, "/")
	if len(elems) == 2 {
		if strings.Contains(elems[0], ".") || strings.Contains(elems[0], "?") {
			return []string{}
		}
		return []string{tmpurl[0] + "//" + tmpurl[2] + "/" + elems[0]}
	}

	for i := range elems {
		if elems[i] == "*" {
			break
		}

		if levels >= 0 && i >= levels {
			for j := 0; j < levels; j++ {
				if strings.Contains(elems[j], "*") || elems[j] == "*" {
					break
				}
				if strings.Contains(elems[j], ".") || strings.Contains(elems[j], "?") {
					break
				}
				resTemp := strings.Join(elems[:j+1], "/")
				resTemp = tmpurl[0] + "//" + tmpurl[2] + "/" + resTemp
				result = append(result, resTemp)
			}
		} else {	
			for j := 0; j < i; j++ {
				if strings.Contains(elems[j], "*") || elems[j] == "*" {
					break
				}
				if strings.Contains(elems[j], ".") || strings.Contains(elems[j], "?") {
					break
				}
				resTemp := strings.Join(elems[:j+1], "/")
				resTemp = tmpurl[0] + "//" + tmpurl[2] + "/" + resTemp
				result = append(result, resTemp)
			}
		}
	}

	return golazy.RemoveDuplicateValues(result)
}
