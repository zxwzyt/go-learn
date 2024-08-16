package designer

import "fmt"

type Computer1 interface {
	Print()
	SetPrinter(Printer)
}

type Mac1 struct {
	printer Printer
}

func (m *Mac1) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac1) SetPrinter(p Printer) {
	m.printer = p
}

type Windows1 struct {
	printer Printer
}

func (w *Windows1) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows1) SetPrinter(p Printer) {
	w.printer = p
}

type Printer interface {
	PrintFile()
}

type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a Epson Printer")
}

type Hp struct{}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

func RunBridge() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac1{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows1{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
