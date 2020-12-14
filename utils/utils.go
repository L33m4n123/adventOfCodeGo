package utils

import (
	"math/big"
	"regexp"
	"strconv"
)

var one = big.NewInt(1)

// ConvertLinesToInt takes the line from the input and parses its values to integer.
func ConvertLinesToInt(lines []string) []int {
	result := []int{}

	for _, line := range lines {
		intLine, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		result = append(result, intLine)
	}

	return result
}

// UniqueStringSlice takes in a onedimensional slice and removes all duplicate entries
func UniqueStringSlice(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// UniqueMultipleStringSlice takes in a 2dimensional slice and makes sure each internal slice is unique. Not caring if a slice is same as another one
func UniqueMultipleStringSlice(slice [][]string) [][]string {
	list := make([][]string, len(slice))
	for key, internalSlice := range slice {
		list[key] = UniqueStringSlice(internalSlice)
	}

	return list
}

// FindStringSubmatchWithNamedMatches is a wrapper function to match a pattern in a string and return the named capture groups
func FindStringSubmatchWithNamedMatches(pattern *regexp.Regexp, input string) map[string]string {
	match := pattern.FindStringSubmatch(input)
	if len(match) == 0 {
		panic("No matches found!")
	}
	result := make(map[string]string)
	for i, name := range pattern.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result
}