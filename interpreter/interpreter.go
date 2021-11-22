package interpreter

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type interpreter struct {
}

func New() interpreter {
	return interpreter{}
}

func (it interpreter) Run(code string) {
	cells := []int{0}
	currentCell := 0

	bracketSkipLayer := 0
	bracketStack := []int{}

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(code); i++ {
		//fmt.Printf("i: %d, cellLen: %d, currentCell: %d, currentCellVal: %d \n", i, len(cells), currentCell, cells[currentCell])

		switch code[i] {

		// add to cell
		case '+':
			cells[currentCell]++

		// subtract from cell
		case '-':
			cells[currentCell]--

		// move cell left
		case '>':
			currentCell++
		// move cell right
		case '<':
			currentCell--

		// put input into current cell
		case ',':
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			cells[currentCell] = int(input[0])
		//output current cell
		case '.':
			fmt.Print(string(cells[currentCell]))

		// loop
		case '[':
			// skip if cell == 0
			if cells[currentCell] == 0 {
				bracketSkipLayer++
				i++
				// jump to end of loop
				for j := 0; j < len(code)-i; j++ {
					if code[i+j] == '[' {
						bracketSkipLayer++
					}
					if code[i+j] == ']' {
						bracketSkipLayer--
					}
					if bracketSkipLayer == 0 {
						i += j
						break
					}
				}
				break
			}
			// add bracket to bracketStack so that code can jump back
			bracketStack = append(bracketStack, i)

		case ']':
			// end loop if current cell == 0
			if cells[currentCell] == 0 {
				// remove loop from bracketStack
				bracketStack = bracketStack[:len(bracketStack)-1]
				break
			}
			// jump to start of loop
			i = bracketStack[len(bracketStack)-1]

		}

		// creates missing cells
		for len(cells)-1 < currentCell {
			cells = append(cells, 0)
		}

		// wrap around current cell
		if currentCell < 0 {
			currentCell = len(cells) - 1
		}

		// 8 bit
		if cells[currentCell] >= 256 {
			cells[currentCell] -= 256
		}
		if cells[currentCell] < 0 {
			cells[currentCell] += 256
		}
	}
}
