package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
	"unicode"
)

// normalizePathSegment replaces numeric segments with a placeholder
func normalizePathSegment(segment string) string {
	// Check if the segment consists only of digits
	isNumeric := true
	for _, r := range segment {
		if !unicode.IsDigit(r) {
			isNumeric = false
			break
		}
	}
	if isNumeric {
		return "NUM"
	}
	return segment
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	seen := make(map[string]string) // key: domain+normalized_directory+extension, value: full URL

	for scanner.Scan() {
		urlStr := scanner.Text()
		parsedURL, err := url.Parse(urlStr)
		if err != nil {
			continue
		}

		// Get the domain and extension
		domain := parsedURL.Host
		ext := path.Ext(parsedURL.Path)

		// Normalize the directory path
		dir := path.Dir(parsedURL.Path)
		parts := strings.Split(dir, "/")
		for i, part := range parts {
			parts[i] = normalizePathSegment(part)
		}
		normalizedDir := strings.Join(parts, "/")

		key := domain + normalizedDir + ext

		// If we haven't seen this combination before, store and print it
		if _, exists := seen[key]; !exists {
			seen[key] = urlStr
			fmt.Println(urlStr)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
