package printer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/ionos-cloud/ionosctl/pkg/constants"

	"github.com/spf13/viper"
)

// Type defines an formatter format.
type Type string

func (p Type) String() string {
	return string(p)
}

const (
	// TypeJSON defines a JSON formatter.
	TypeJSON = Type("json")
	// TypeText defines a human-readable formatted formatter.
	TypeText = Type("text")
)

type Registry map[string]PrintService

var unknownTypeFormatErr = "unknown type format %s. Hint: use --output json|text"

func NewPrinterRegistry(out, outErr io.Writer, noHeaders bool) (Registry, error) {
	if viper.GetString(constants.ArgOutput) != TypeJSON.String() &&
		viper.GetString(constants.ArgOutput) != TypeText.String() {
		return nil, errors.New(fmt.Sprintf(unknownTypeFormatErr, viper.GetString(constants.ArgOutput)))
	}

	return Registry{
		TypeJSON.String(): &JSONPrinter{
			Stderr: outErr,
			Stdout: out,
		},
		TypeText.String(): &TextPrinter{
			Stderr:    outErr,
			Stdout:    out,
			NoHeaders: noHeaders,
		},
	}, nil
}

type PrintService interface {
	Print(interface{}) error
	Warn(interface{}) error
	Verbose(format string, a ...interface{})

	GetStdout() io.Writer
	SetStdout(io.Writer)
	GetStderr() io.Writer
	SetStderr(io.Writer)
}

type JSONPrinter struct {
	Stdout io.Writer
	Stderr io.Writer
}

func (p *JSONPrinter) write(out io.Writer, v interface{}) error {
	switch v.(type) {
	case Result:
		result := v.(Result)
		if err := result.PrintJSON(out); err != nil {
			return err
		}
	default:
		var msg DefaultMsgPrint
		msg.Message = v
		err := WriteJSON(&msg, out)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *JSONPrinter) Print(v interface{}) error {
	return p.write(p.Stdout, v)
}

func (p *JSONPrinter) Warn(v interface{}) error {
	return p.write(p.Stderr, v)
}

func (p *JSONPrinter) Verbose(format string, a ...interface{}) {
	flag := viper.GetBool(constants.ArgVerbose)
	var toPrint = ToPrint{}
	if flag {
		str := fmt.Sprintf("[INFO] "+format, a...)
		toPrint.Message = str
		err := WriteJSON(&toPrint, p.Stderr)
		if err != nil {
			return
		}
	}
}

func (p *JSONPrinter) GetStdout() io.Writer {
	return p.Stdout
}

func (p *JSONPrinter) SetStdout(writer io.Writer) {
	p.Stdout = writer
}

func (p *JSONPrinter) GetStderr() io.Writer {
	return p.Stderr
}

func (p *JSONPrinter) SetStderr(writer io.Writer) {
	p.Stderr = writer
}

type TextPrinter struct {
	Stdout    io.Writer
	Stderr    io.Writer
	NoHeaders bool
}

func (p *TextPrinter) write(out io.Writer, v interface{}) error {
	switch v.(type) {
	case Result:
		result := v.(Result)
		if err := result.PrintText(out, p.NoHeaders); err != nil {
			return err
		}
	case string:
		v = strings.TrimRight(v.(string), "\n")
		if _, err := fmt.Fprintf(out, "%v\n", v); err != nil {
			return err
		}
	default:
		_, err := fmt.Fprintf(out, "%v\n", v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *TextPrinter) Print(v interface{}) error {
	return p.write(p.Stdout, v)
}

func (p *TextPrinter) Warn(v interface{}) error {
	return p.write(p.Stderr, v)
}

func (p *TextPrinter) Verbose(format string, a ...interface{}) {
	flag := viper.GetBool(constants.ArgVerbose)
	if flag {
		fmt.Printf("[INFO] "+format+"\n", a...)
	} else {
		return
	}
}

func (p *TextPrinter) GetStdout() io.Writer {
	return p.Stdout
}

func (p *TextPrinter) SetStdout(writer io.Writer) {
	p.Stdout = writer
}

func (p *TextPrinter) GetStderr() io.Writer {
	return p.Stderr
}

func (p *TextPrinter) SetStderr(writer io.Writer) {
	p.Stderr = writer
}

type DefaultMsgPrint struct {
	Message interface{} `json:"Message,omitempty"`
}

type ToPrint struct {
	Message string
}

func WriteJSON(item interface{}, writer io.Writer) error {
	j, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(writer, "%s\n", string(j))
	if err != nil {
		return err
	}
	return nil
}
