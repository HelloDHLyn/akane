package interpreter

import (
	"bufio"
	"os"

	"github.com/hellodhlyn/akane/internal/parser"
)

type Interpreter struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		reader: bufio.NewReader(os.Stdin),
		writer: bufio.NewWriter(os.Stdout),
	}
}

func (i *Interpreter) Run() error {
	i.writeString("*** Akane Programming Language Interpreter ***\n")
	for {
		i.writeString("\n>> ")
		source, err := i.readLine()
		if err != nil {
			return err
		}

		p := parser.NewParser(source)
		world, err := p.Parse()
		if err != nil {
			i.writeString(err.Error())
			i.write([]byte("\n"))
			continue
		}

		obj := world.Eval(nil)
		i.write(obj.Bytes())
		i.write([]byte("\n"))
	}
}

func (i *Interpreter) readLine() ([]byte, error) {
	return i.reader.ReadBytes('\n')
}

func (i *Interpreter) write(line []byte) error {
	_, err := i.writer.Write(line)
	if err != nil {
		return err
	}
	return i.writer.Flush()
}

func (i *Interpreter) writeString(line string) error {
	_, err := i.writer.WriteString(line)
	if err != nil {
		return err
	}
	return i.writer.Flush()
}
