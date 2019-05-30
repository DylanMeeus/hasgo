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
	ElementTypeSymbol    = "ElementType"
	SliceTypeSymbol      = "SliceType"
	SliceSliceTypeSymbol = "SliceSliceType" // todo: generalize for N slices
)

var (
	Type        = flag.String("T", "", "Type for which to generate data")
	SType       = flag.String("S", "", "Corresponding Slice Type for T")
	Pkg         = flag.String("P", "types", "Package for which to generate")
	numberTypes = map[string]struct{}{
		"int":     struct{}{},
		"int32":   struct{}{},
		"int64":   struct{}{},
		"uint":    struct{}{},
		"uint32":  struct{}{},
		"uint64":  struct{}{},
		"float":   struct{}{},
		"float32": struct{}{},
		"float64": struct{}{},
	}
	stringTypes = map[string]struct{}{
		"string": struct{}{},
	}
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// the Type / SType for which we are generating data..
type symbols struct {
	T, ST string
}

type Generator struct {
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
// todo: figure out how to run hasgo for user files. but should I even do so?

func main() {
	flag.Parse()
	fmt.Printf("type: %v - slice: %v\n", *Type, *SType)
	sym := symbols{*Type, *SType}
	g := Generator{
		imports: map[string]struct{}{},
		pkg:     *Pkg,
	}
	g.parsePackage(os.Args, nil)
	// stringer prints everything in one file. This might be bad.
	// but let's roll with it for now :-)
	g.generate(sym)
	err := ioutil.WriteFile(fmt.Sprintf("%v_hasgo.go", *SType), g.format(), 0644)
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
	// todo: figure out how to determine if it's a struct!
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
func (g *Generator) generate(s symbols) {
	funcs := []string{}
	for function, _ := range hasgoTemplates {
		funcs = append(funcs, function)
	}
	sort.Strings(funcs)
	for _, function := range funcs {
		template := hasgoTemplates[function]
		if !validFunction(function, s.T) {
			continue
		}
		template, imports := extractImports(template)
		g.AddImports(imports)
		g.Printf("// =============== %v =================\n", function)
		g.Printf(generify(template, s))
		g.NewLine()
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
	out = strings.Replace(out, ElementTypeSymbol, sym.T, -1)
	out = strings.Replace(out, SliceSliceTypeSymbol, "[][]"+sym.T, -1)
	out = strings.Replace(out, SliceTypeSymbol, sym.ST, -1)
	return
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) NewLine() {
	g.Printf("\n")
}

func (g *Generator) AddImports(imports []string) {
	for _, imp := range imports {
		// sanitize:
		sane := strings.TrimSpace(imp)
		g.imports[sane] = struct{}{}
	}
}

// analyze the package
func (g *Generator) parsePackage(patterns []string, tags []string) {
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
func (g *Generator) addPackage(pkg *packages.Package) {
	fmt.Println("adding files..")
	for _, file := range pkg.Syntax {
		fmt.Printf("file %v\n", file)
	}
}

// return all unique imports
func (g *Generator) Imports() (out []string) {
	for k, _ := range g.imports {
		out = append(out, k)
	}
	return
}

// print the formatted source code
// adds the package declaration + imports
func (g *Generator) format() []byte {
	code := "// code generated by hasgo. [DO NOT EDIT!]" +
		"\n" +
		"package " + g.pkg + "\n"

	if imports := g.Imports(); len(imports) > 0 {
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
