package main
import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strings"
)

func logError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func printMap(bill map[string]float32) {
	for key, val := range bill {
		fmt.Println(key, val)
	}
}

func printSlice(list []byte) {
	for i, v := range list {
		fmt.Println(i, ":", v)
	}
}

func AppendByte(slice []byte, data ...byte) []byte {
    m := len(slice)
    n := m + len(data)
    if n > cap(slice) {
        newSlice := make([]byte, (n+1)*2)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:n]
    copy(slice[m:n], data)
    return slice
}

func printBill(bill_array []byte) {
	bill := string(bill_array[:])
	fmt.Println(bill)
}

func is_lower_case(tmp byte) bool {
	return ('a' <= tmp && tmp <= 'z')
}

func is_upper_case(tmp byte) bool {
	return ('A' <= tmp && tmp <= 'Z')
}

func is_alpha(tmp byte) bool {
	return is_lower_case(tmp) || is_upper_case(tmp)
}

func saveBill(filename string, bill map[string]float32, tip []float32) {
	d := []byte("Bill breakdown:\n")
	var s string
	var t string
	var sum float32 = 0.0
	var total float32 = 0.0
	for key, v := range bill {
		s = fmt.Sprintf("%-25s ...$%0.2f", key+":", v)
		total += v
		for i := 0; i < len(s); i++ {
			if i == 0 {
				if is_alpha(s[i]) {
					if is_lower_case(s[i]) {
						d = AppendByte(d, s[i]-32)
					} else {
						d = AppendByte(d, s[i])
					}
				}
			} else {
				d = AppendByte(d, s[i])
			}
		}
		d = AppendByte(d, '\n')
	}

	for j := 0; j < len(tip); j++ {
		sum += tip[j]
	}

	total += sum
	t = fmt.Sprintf("%-25s ...$%0.2f", "Tip:", sum)
	for k := 0; k < len(t); k++ {
		d = AppendByte(d, t[k])
	}

	d = AppendByte(d, '\n')
	tmp := fmt.Sprintf("%-25s ...$%0.2f", "Total:", total)
	for l := 0; l < len(tmp); l++ {
		d = AppendByte(d, tmp[l])
	}

	printBill(d)

	if len(bill) != 0 || len(tip) != 0 {
		err := os.WriteFile(filename+".txt", d, 0666)
		logError(err)
	}
}

func main() {
	var choice byte = 'a'
	var item string = "temu"
	var price float32 = 3.14
	var tip float32 = 1.6
	var tips []float32
	var bill map[string]float32
	bill = make(map[string]float32)
	log.SetPrefix("Input: ")
	log.SetFlags(0)
	reader :=  bufio.NewReader(os.Stdin)
	read := bufio.NewReader(os.Stdin)
	fmt.Print("Create a new bill name: ")
	name, err := read.ReadString('\n')
	logError(err)
	name = strings.TrimSpace(name)
	fmt.Println("Created the bill -", name)
	for choice != 's' {
		fmt.Print("Choose option (a: add item, s: save bill, t: add tip): ")
		_, err2 := fmt.Scanf("%c", &choice)
		logError(err2)
		_, er := reader.ReadBytes('\n')
		logError(er)

		if choice == 'a' {
			fmt.Print("Item: ")
			_, i := fmt.Scanln(&item)
			logError(i)
			fmt.Print("Item price: ")
			_, p := fmt.Scanln(&price)
			logError(p)
			bill[item] = price
			fmt.Println("Item added -", item, price)
		} else if choice == 't' {
			fmt.Print("Enter tip amount ($): ")
			_, t := fmt.Scanln(&tip)
			logError(t)
			fmt.Print("Tip is now $", tip)
			fmt.Print("\n")
			tips = append(tips, tip)
		} 
	}//for
	saveBill(name, bill, tips)
}
