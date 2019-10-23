package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"golang.org/x/tools/go/packages"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

const (
	elementTypeSymbol    = "ElementType"
	sliceTypeSymbol      = "SliceType"
	sliceSliceTypeSymbol = "SliceSliceType" // todo: generalize for N slices
)

var (
	unitType    = flag.String("T", "", "Type for which to generate data")
	sliceType   = flag.String("S", "", "Corresponding Slice Type for T")
	packag      = flag.String("P", "types", "Package for which to generate")
	numberTypes = map[string]struct{}{
		"int":     {},
		"int32":   {},
		"int64":   {},
		"uint":    {},
		"uint32":  {},
		"uint64":  {},
		"float":   {},
		"float32": {},
		"float64": {},
	}
	stringTypes = map[string]struct{}{
		"string": {},
	}
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// the Type / sliceType for which we are generating data..
type symbols struct {
	T, ST string
}

type generator struct {
	buf     bytes.Buffer        // accumulate the output
	pkg     string              // package in which to place the generated code.
	imports map[string]struct{} // the imports that the generator needs to add
}

// currently hasgo works like this:
// once the generator is trigger, I scan all files in the named packages
// I generate one hasgo file with functions
// all the functions get  stored here
// I write out _once_, collecting all data
// Importantly, hasgo should trigger only once.. I should figure out how to do so

func main() {
	flag.Parse()
	fmt.Printf("type: %v - slice: %v\n", *unitType, *sliceType)
	sym := symbols{*unitType, *sliceType}
	g := generator{
		imports: map[string]struct{}{},
		pkg:     *packag,
	}
	g.parsePackage(os.Args, nil)
	// stringer prints everything in one file. This might be bad.
	// but let's roll with it for now :-)
	g.generate(sym)
	err := ioutil.WriteFile(fmt.Sprintf("%v_hasgo.go", *sliceType), g.format(), 0644)
	if err != nil {
		panic(err)
	}
}

// is the function valid for the type?
func validFunction(function, T string) bool {
	domains, ok := funcDomains[function]
	if !ok {
		return false
	}
	for _, d := range domains {
		switch d {
		case ForNumbers:
			if _, ok := numberTypes[T]; ok {
				return true
			}
			break
		case ForStrings:
			if _, ok := stringTypes[T]; ok {
				return true
			}
			break
		case ForStructs:
			// everything that is not a string or number we will consider a struct!
			if _, ok := numberTypes[T]; !ok {
				if _, ok := stringTypes[T]; !ok {
					return true
				}
			}
			break

		}
	}
	return false
}

// write the data for the generator
func (g *generator) generate(s symbols) {
	funcs := []string{}
	for function := range hasgoTemplates {
		funcs = append(funcs, function)
	}
	sort.Strings(funcs)
	for _, function := range funcs {
		template := hasgoTemplates[function]
		if !validFunction(function, s.T) {
			continue
		}
		template, imports := extractImports(template)
		g.addImports(imports)
		g.printf("// =============== %v =================\n", function)
		g.print(generify(template, s))
		g.newline()
	}
}

// extracts the imports and removes them from the template
// we assume that the go code is correctly formatted!
func extractImports(template string) (outtemp string, imports []string) {
	// split template into lines
	// check for import statements in line
	lines := strings.Split(template, "\n")
	nonImportLines := []string{}
	for i := 0; i < len(lines); {
		line := lines[i]
		if strings.Contains(line, "import") {
			// determine if it's a single import or multiple..
			if strings.Contains(line, "(") {
				// implement
				i++
				line := lines[i]
				// check end condition
				for !strings.Contains(line, ")") {
					imports = append(imports, line)
					i++
					line = lines[i]
				}
				i++
				continue
			} else {
				line = strings.TrimSpace(line)[len("import"):]
				imports = append(imports, line)
				i++
				continue
			}
		}
		nonImportLines = append(nonImportLines, line)
		i++
	}

	return strings.Join(nonImportLines, "\n"), imports
}

// replace the placeholders by the correct symbols
func generify(template string, sym symbols) (out string) {
	out = template
	out = strings.Replace(out, elementTypeSymbol, sym.T, -1)
	out = strings.Replace(out, sliceSliceTypeSymbol, "[][]"+sym.T, -1)
	out = strings.Replace(out, sliceTypeSymbol, sym.ST, -1)
	return
}

func (g *generator) printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *generator) print(s string) {
	fmt.Fprint(&g.buf, s)
}

func (g *generator) newline() {
	g.printf("\n")
}

func (g *generator) addImports(imports []string) {
	for _, imp := range imports {
		// sanitize:
		sane := strings.TrimSpace(imp)
		g.imports[sane] = struct{}{}
	}
}

// analyze the package
func (g *generator) parsePackage(patterns []string, tags []string) {
	cfg := &packages.Config{
		Mode:       packages.LoadSyntax,
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
	}
	pkgs, err := packages.Load(cfg, patterns...)
	check(err)
	if len(pkgs) == 0 {
		panic("not enough packages")
	}
	fmt.Printf("packages %v\n", pkgs)
	g.addPackage(pkgs[0])
}

// add a package to the generator
func (g *generator) addPackage(pkg *packages.Package) {
	fmt.Println("adding files..")
	for _, file := range pkg.Syntax {
		fmt.Printf("file %v\n", file)
	}
}

// return all unique imports
func (g *generator) listImports() (out []string) {
	for k := range g.imports {
		out = append(out, k)
	}
	return
}

// print the formatted source code
// adds the package declaration + imports
func (g *generator) format() []byte {
	code := "// Package " + g.pkg + " contains code generated by hasgo. [DO NOT EDIT!]" +
		"\n" +
		"package " + g.pkg + "\n"

	if imports := g.listImports(); len(imports) > 0 {
		code += "import (\n"
		code += strings.Join(imports, "\n") + "\n)\n"
	}
	// add our generated functions
	src := append([]byte(code), g.buf.Bytes()...)
	src, err := format.Source(src)
	if err != nil {
		// should never happen. But, during development, might.
		fmt.Println("Printing code without formatting")
		return append([]byte(code), g.buf.Bytes()...)
	}
	return src
}
