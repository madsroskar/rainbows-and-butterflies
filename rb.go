package main

import (
	"fmt"
	"errors"
	"io/ioutil"
	"bufio"
	"os"
)

var (
	fileName string
	opMvL    string = "\U0001f60f" // smirking-face
	opMvR    string = "\U0001f60e" // smiling-face-with-sunglasses
	opInc    string = "\U0001f308" // rainbow
	opDec    string = "\U0001f98b" // butterfly
	opOut    string = "\U0001f6bd" // toilet (heh, flush)
	opIn     string = "\U0001f442"
	opJmpF   string = "\U0001f91c" // right-facing-fist
	opJmpB   string = "\U0001f91b" // left-facing-fist
)

func init() {
	args := os.Args
	if len(args) != 2 {
		panic("A filename must be supplied")
	}

	fileName = args[1]
}

//func compile(input string) ()

func readFile() (string, error) {
	s, err := ioutil.ReadFile(fileName)
	return string(s), err
}

type operation struct {
	operator string
	operand  uint16
}

func compile(input string) (program []operation, err error) {
	var pc, jmpPc uint16 = 0, 0
	jmpStk := make([]uint16, 0)
	for _, op := range input {
		switch string(op) {
		case opMvR, opMvL, opInc, opDec, opIn, opOut:
			program = append(program, operation{string(op), 0})
		case opJmpF:
			program = append(program, operation{string(op), 0})
			jmpStk = append(jmpStk, pc)
		case opJmpB:
			if len(jmpStk) == 0 {
				return nil, errors.New("Compilation error.")
			}
			jmpPc = jmpStk[len(jmpStk)-1]
			jmpStk = jmpStk[:len(jmpStk)-1]
			program = append(program, operation{string(op), jmpPc})
			program[jmpPc].operand = pc
		default:
			pc--
		}
		pc++
	}
	return
}

func execute(program []operation) {
	data := make([]int16, 65535)
	var data_ptr uint16  = 0
	reader := bufio.NewReader(os.Stdin)
	for progCounter := 0; progCounter < len(program); progCounter++ {
		switch program[progCounter].operator {
		case opMvR:
			data_ptr++
		case opMvL:
			data_ptr--
		case opInc:
			data[data_ptr]++
		case opDec:
			data[data_ptr]--
		case opOut:
			fmt.Printf("%c", data[data_ptr])
		case opIn:
			read_val, _ := reader.ReadByte()
			data[data_ptr] = int16(read_val)
		case opJmpF:
			if data[data_ptr] == 0 {
				progCounter = int(program[progCounter].operand)
			}
		case opJmpB:
			if data[data_ptr] != 0 {
				progCounter = int(program[progCounter].operand) - 1
			}
		default:
			panic("Unknown operator.")
		}
	}
}

func main() {
	fileContent, err := readFile()
	if err != nil {
		panic(err)
	}
	program, err := compile(fileContent)
	execute(program)
}
