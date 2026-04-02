package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Tavershima Gbaa]
// Squad:  [The Interface]

func upper(tex string) string {
	t := strings.Fields(tex)
	e := strings.Join(t, " ")
	return strings.ToUpper(e)
}
func lower(tex string) string {
	t := strings.Fields(tex)
	e := strings.Join(t, " ")
	return strings.ToLower(e)
}
func cap(tex string) string {
	t := strings.Fields(tex)
	e := strings.Join(t, " ")
	x := strings.ToLower(e)
	tx := strings.Title(x)
	return tx
}
func title(tex string) string {
	tx := strings.Fields(tex)
	for i, w := range tx {
		if len(w) > 3 {
			tx[i] = strings.ToUpper(tx[i][:1]) + strings.ToLower(tx[i][1:])
		}
		tx[0] = strings.ToUpper(tx[0][:1]) + strings.ToLower(tx[0][1:])
	}
	return strings.Join(tx, " ")
}
func snake(tex string) string {
	t := strings.Fields(tex)
	e := strings.Join(t, " ")
	x := strings.ReplaceAll(e, " ", "_")
	tx := strings.NewReplacer("@", "", "'", "", "}", "", "{", "", "/", "", "?", "", ".", "", "!", "", "|", "", "&", "", "$", "", "£", "", "%", "", "*", "").Replace(x)
	return strings.ToLower(tx)
}
func reverse(tex string) string {
	t := strings.Split(tex, "")
	slices.Reverse(t)
	return strings.Join(t, "")
}
func main() {

	fmt.Println("________________________________")
	fmt.Println("___Vince Strings Transformer____")
	fmt.Println("________________________________")
	fmt.Print("Transformation Menu")
	fmt.Print(`
1.upper -- to uppercase string input
2.lower -- to lowercase string input
3.cap --  to Capitalise the first letter of every word
4.title -- ike cap, but small connector words stay lowercase 
5.snake -- all lowercase, spaces replaced with _.Remove any character that is not a letter, digit, underscore
6.reverse -- to reverse string input
7. exit -- to terminate the program
`)

	reader := bufio.NewReader(os.Stdin)

	var text string
start:
	for {
		fmt.Println()
		fmt.Print("ENTER INPUT : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(`
		1.upper -- to uppercase string input
		2.lower -- to lowercase string input
		3.cap --  to Capitalise the first letter of every word
		4.title -- ike cap, but small connector words stay lowercase 
		5.snake -- all lowercase, spaces replaced with _. Remove any character that is not a letter, digit, underscore
		6.reverse -- to reverse string input
		7.exit -- to terminate the program`)
			continue
		}

		if input == strings.ToLower("CAP") {
			fmt.Println("No text provided. Usage: cap <text>")
			continue
		}

		if input == "upper" {
			fmt.Println("No text provided. Usage: upper <text>")
			continue
		}

		if input == "snake" {
			fmt.Println("No text provided. Usage: snake <text>")
			continue
		}

		if input == "lower" {
			fmt.Println("No text provided. Usage: lower <text>")
			continue
		}
		if input == "title" {
			fmt.Println("No text provided. Usage: title <text>")
			continue
		}
		if input == "reverse" {
			fmt.Println("No text provided. Usage: reverse <text>")
			continue
		}
		if input == "exit" {
			return
		}

		text = input

		word := strings.Fields(text)

		first := word[0]
		marker := []string{"cap", "lower", "upper", "snake", "title", "reverse"}
		found := false

		for _, char := range marker {
			if first == char {
				found = true
				break
			}

		}
		if !found {
			fmt.Println("Unknown command:", word[0])
			fmt.Println(" Valid commands: upper, lower, cap, title, snake, reverse, exit")
			continue
		}

		break
	}

	if strings.HasPrefix(text, strings.ToLower("UPPER ")) {
		text = strings.TrimPrefix(text, "upper")
		result := upper(text)
		fmt.Println("OUTPUT :", result)
		goto start
	}
	if strings.HasPrefix(text, "lower ") {
		text = strings.TrimPrefix(text, "lower ")
		result := lower(text)
		fmt.Println("OUTPUT:", result)
		goto start
	}
	if strings.HasPrefix(text, "cap ") {
		text = strings.TrimPrefix(text, "cap ")
		result := cap(text)
		fmt.Println("OUTPUT:", result)
		goto start
	}
	if strings.HasPrefix(text, "title ") {
		text = strings.TrimPrefix(text, "title ")
		result := title(text)
		fmt.Println("OUTPUT:", result)
		goto start
	}
	if strings.HasPrefix(text, "snake ") {
		text = strings.TrimPrefix(text, "snake ")
		result := snake(text)
		fmt.Println("OUTPUT: ", result)
		goto start
	}
	if strings.HasPrefix(text, "reverse ") {
		text = strings.TrimPrefix(text, "reverse ")
		result := reverse(text)
		fmt.Println("OUTPUT: ", result)
		goto start
	}
}
