package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"

	"honnef.co/go/tools/analysis/lint"

	"github.com/sourcegraph/sourcegraph/dev/linters/staticcheck"
)

var ignoredAnalyzer = map[string]string{
	"SAXXXX": "I am an exmaple for a linter that should be ignored",
}

// if you add analyzers here make sure that staticcheck.go knows about it too!
var analyzers []*lint.Analyzer = sorted(staticcheck.AllAnalyzers)

var BazelBuildTemplate = `# GENERATED FILE - DO NOT EDIT
# This file was generated by running go generate on dev/linters/staticcheck
#
# If you want to ignore an analyzer add it to the ignore list in dev/linters/staticcheck/cmd/gen.go,
# and re-run go generate.
#
# IMPORTANT: Starlark is spaces/tabs sensitive, make sure the below is SPACE indented.

load("@io_bazel_rules_go//go:def.bzl", "go_library")

def staticcheck_targets():
{{ range .Analyzers }}    go_library(
        name = "{{.Analyzer.Name}}",
        srcs = ["staticcheck.go"],
        importpath = "github.com/sourcegraph/sourcegraph/dev/linters/staticcheck/{{.Analyzer.Name}}",
        visibility = ["//visibility:public"],
        x_defs = {"AnalyzerName": "{{.Analyzer.Name}}"},
        deps = [
            "//dev/linters/nolint",
            "@co_honnef_go_tools//analysis/lint",
            "@co_honnef_go_tools//simple",
            "@co_honnef_go_tools//staticcheck",
            "@org_golang_x_tools//go/analysis",
        ],
    )

{{ end -}}
`

var BazelDefTemplate = `# DO NOT EDIT - this file was generated by running go generate on dev/linters/staticcheck
#
# If you want to ignore an analyzer add it to the ignore list in dev/linters/staticcheck/cmd/gen.go,
# and re-run go generate

STATIC_CHECK_ANALYZERS = [
{{- range .Analyzers }}
    "//dev/linters/staticcheck:{{.Analyzer.Name}}",
{{- end }}
]
`

func unique(analyzers ...*lint.Analyzer) []*lint.Analyzer {
	set := make(map[string]bool)
	uniq := make([]*lint.Analyzer, 0)

	for _, a := range analyzers {
		if _, ok := set[a.Analyzer.Name]; !ok {
			// first time we see this analyzer!
			uniq = append(uniq, a)
			set[a.Analyzer.Name] = true
		}
	}

	return uniq
}

func filterIgnored(analyzers []*lint.Analyzer, ignored map[string]string) []*lint.Analyzer {
	result := make([]*lint.Analyzer, 0)
	for _, a := range analyzers {
		if _, shouldIgnore := ignored[a.Analyzer.Name]; !shouldIgnore {
			result = append(result, a)
		}
	}

	return result
}

func sorted(analyzers []*lint.Analyzer) []*lint.Analyzer {
	linters := filterIgnored(unique(analyzers...), ignoredAnalyzer)
	// remove ignored linters first
	// now sort them
	sort.SliceStable(linters, func(i, j int) bool {
		return strings.Compare(linters[i].Analyzer.Name, linters[j].Analyzer.Name) < 0
	})
	return linters
}

func writeTemplate(targetFile, templateDef string) error {
	name := targetFile
	tmpl := template.Must(template.New(name).Parse(templateDef))

	f, err := os.OpenFile(targetFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0o666)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, struct {
		Analyzers []*lint.Analyzer
	}{
		Analyzers: analyzers,
	})
	if err != nil {
		return err
	}

	return nil
}

// We support two position arguments:
// 1: buildfile path - file where the analyzer targets should be generated to
// 2: analyzer definition path - file where a convienience analyzer array is generated that contains all the targets
func main() {
	targetFile := "targets.bzl"
	if len(os.Args) > 1 {
		targetFile = os.Args[1]
	}

	// Generate targets for all the analyzers
	if err := writeTemplate(targetFile, BazelBuildTemplate); err != nil {
		fmt.Fprintln(os.Stderr, "failed to render Bazel buildfile template")
		panic(err)
	}

	// Generate a file where we can import the list of analyzers into our bazel scripts
	targetFile = "analyzers.bzl"
	if len(os.Args) > 2 {
		targetFile = os.Args[2]
	}
	if err := writeTemplate(targetFile, BazelDefTemplate); err != nil {
		fmt.Fprintln(os.Stderr, "failed to render Anazlyers definiton template")
		panic(err)
	}
}
