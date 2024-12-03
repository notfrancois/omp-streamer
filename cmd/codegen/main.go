package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const srcTpl = `#include "include/%s"

#ifdef __cplusplus
extern "C" {
#endif

%s

#ifdef __cplusplus
}
#endif
`

const funcTpl = `%s%s %s(%s) {
%s%sreturn call<%s>(%s);
%s}`

// ^\s*([\w\s\*]+)\s+(\w+)\s*\(([^)]*)\)\s*;
var funcPrototypeRe = regexp.MustCompile(`(?m)` +
	`^` +
	`\s*` +
	`([\w\s\*]+)` + // type
	`\s+` +
	`(\w+)` + // name
	`\s*` +
	`\(` +
	`([^)]*)` + // arguments
	`\)` +
	`\s*` +
	`;`,
)

type headerFunction struct {
	returnType string
	name       string
	parameters string
}

func main() {
	if err := generateCppWrappers("include"); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating C++ wrappers: %v\n", err)
		os.Exit(1)
	}
}

func generateCppWrappers(headersPath string) error {
	headers, err := os.ReadDir(headersPath)
	if err != nil {
		return fmt.Errorf("error reading headers directory %s: %w", headersPath, err)
	}

	for _, header := range headers {
		if header.Name() == "omp.h" {
			continue
		}

		headerPath := filepath.Join(headersPath, header.Name())
		if err := createCppWrapper(headerPath); err != nil {
			return fmt.Errorf("error creating wrapper for %s: %w", headerPath, err)
		}
	}

	return nil
}

func createCppWrapper(headerPath string) error {
	contents, err := os.ReadFile(headerPath)
	if err != nil {
		return fmt.Errorf("error reading header file %s: %w", headerPath, err)
	}

	matches := funcPrototypeRe.FindAllStringSubmatch(string(contents), -1)
	wrapperFunctions := make([]string, 0, len(matches))

	for _, match := range matches {
		fn := headerFunction{
			returnType: match[1],
			name:       match[2],
			parameters: match[3],
		}
		wrapperFunctions = append(wrapperFunctions, createWrapperFunction(fn, "    "))
	}

	headerName := strings.TrimSuffix(filepath.Base(headerPath), filepath.Ext(headerPath))
	wrapperContent := fmt.Sprintf(srcTpl, filepath.Base(headerPath), strings.Join(wrapperFunctions, "\n\n"))

	wrapperPath := fmt.Sprintf("api/%s.cpp", headerName)
	if err := os.WriteFile(wrapperPath, []byte(wrapperContent), 0644); err != nil {
		return fmt.Errorf("error writing wrapper file %s: %w", wrapperPath, err)
	}

	return nil
}

func createWrapperFunction(fn headerFunction, indent string) string {
	callArgs := []string{fmt.Sprintf(`"%s"`, fn.name)}

	if fn.parameters != "" {
		callArgs = append(callArgs, extractParameterNames(fn.parameters)...)
	}

	return fmt.Sprintf(
		funcTpl,
		indent,
		fn.returnType,
		fn.name,
		fn.parameters,
		indent, indent,
		fn.returnType,
		strings.Join(callArgs, ", "),
		indent,
	)
}

func extractParameterNames(parameters string) []string {
	paramList := strings.Split(parameters, ",")
	names := make([]string, 0, len(paramList))

	for _, param := range paramList {
		words := strings.Fields(strings.TrimSpace(param))
		if len(words) > 0 {
			names = append(names, words[len(words)-1])
		}
	}

	return names
}
