/* TODO
   - Swap elements
   - Delete elements
   - Edit elements
   - Add more math functions
   - Show history
   - Use tcell?
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func checkNumberOfParams(stack *[]float64, numberOfElements int) error {
	l := len(*stack)
	if l < numberOfElements {
		return errors.New("To few elements in the stack")
	}

	return nil
}

func add(stack *[]float64) error {
	if err := checkNumberOfParams(stack, 2); err != nil {
		return err
	}

	newStack := *stack
	l := len(newStack)
	newStack[l-2] = newStack[l-2] + newStack[l-1]
	*stack = newStack[:len(newStack)-1]

	return nil
}

func mul(stack *[]float64) error {
	if err := checkNumberOfParams(stack, 2); err != nil {
		return err
	}

	newStack := *stack
	l := len(newStack)
	newStack[l-2] = newStack[l-2] * newStack[l-1]
	*stack = newStack[:len(newStack)-1]

	return nil
}

func pow(stack *[]float64) error {
	if err := checkNumberOfParams(stack, 2); err != nil {
		return err
	}

	newStack := *stack
	l := len(newStack)
	newStack[l-2] = math.Pow(newStack[l-2], newStack[l-1])
	*stack = newStack[:len(newStack)-1]

	return nil
}

func div(stack *[]float64) error {
	if err := checkNumberOfParams(stack, 2); err != nil {
		return err
	}

	newStack := *stack
	l := len(newStack)
	newStack[l-2] = newStack[l-2] / newStack[l-1]
	*stack = newStack[:len(newStack)-1]

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	stack := make([]float64, 0)

	for {
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]

		switch input {
		case "+":
			if err := add(&stack); err != nil {
				fmt.Println(err)
				continue
			}
		case "-":
			if checkNumberOfParams(&stack, 2) != nil {
				continue
			}
			stack[len(stack)-1] = -stack[len(stack)-1]
			add(&stack)
		case "*":
			if err := mul(&stack); err != nil {
				fmt.Println(err)
				continue
			}
		case "/":
			if err := div(&stack); err != nil {
				fmt.Println(err)
				continue
			}
		case "^":
			if err := pow(&stack); err != nil {
				fmt.Println(err)
				continue
			}
		case "n": // negate the last element
			if err := checkNumberOfParams(&stack, 1); err != nil {
				fmt.Println(err)
				continue
			}
			stack[len(stack)-1] = -stack[len(stack)-1]
		case "\t": // TODO: swap the two last elements
			if err := checkNumberOfParams(&stack, 2); err != nil {
				fmt.Println(err)
				continue
			}
			l := len(stack)
			stack[l-1], stack[l-2] = stack[l-2], stack[l-1]
		case "p": // print the stack
			break
		default:
			// Duplicate the last element
			if len(input) == 0 && len(stack) >= 1 {
				stack = append(stack, stack[len(stack)-1])
				fmt.Println(stack)
				continue
			}
			value, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			stack = append(stack, value)
		}

		fmt.Println(stack)
	}
}
