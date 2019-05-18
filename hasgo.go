package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"golang.org/x/tools/go/packages"
	"io/ioutil"
	"os"
	"strings"
)

const (
	ElementTypeSymbol = "ElementType"
	SliceTypeSymbol   = "SliceType"
)

var (
	Type  = flag.String("T", "", "Type for which to generate data")
	SType = flag.String("S", "", "Corresponding Slice Type for T")
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
	buf    bytes.Buffer      // accumulate the output
	funpkg *packages.Package // package where the functions are
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
	g := Generator{}
	g.parsePackage(os.Args, nil)
	// stringer prints everything in one file. This might be bad.
	// but let's roll with it for now :-)
	g.Printf("// Code generated by hasgo.go [DO NOT EDIT]")
	g.Printf("\n")
	// todo: unhardcode
	g.Printf("package types\n")
	g.generate(sym)
	ioutil.WriteFile(fmt.Sprintf("%v_hasgo.go", *Type), g.format(), 0644)
}

// write the data for the generator
func (g *Generator) generate(s symbols) {
	for function, template := range hasgoTemplates {
		g.Printf("//===============%v=============\n", function)
		g.Printf(generify(template, s))
		g.Printf("//===============\n")
	}
}

// replace the placeholders by the correct symbols
func generify(template string, sym symbols) (out string) {
	out = template
	out = strings.Replace(out, ElementTypeSymbol, sym.T, -1)
	out = strings.Replace(out, SliceTypeSymbol, sym.ST, -1)
	return
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
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

// print the formatted source code
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// should never happen. During development, might.
		fmt.Println("Printing code without formatting")
		return g.buf.Bytes()
	}
	return src
}
