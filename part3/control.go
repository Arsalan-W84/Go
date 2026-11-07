//CONTROL STATEMENTS IN GO :
//CONDITIONAL : IF , IF ELSE , SWITCH
//LOOP : ONLY FOR KEYWORD IS AVAILABLE , BUT WE CAN WRITE 'WHILE' LOOPS WITH IT

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintSerial(n int) string{
	var op strings.Builder

	for i:=1;i<=n;i++ {
		if i<n{	
			op.WriteString(strconv.Itoa(i));
			op.WriteString(",");
		}else{
			op.WriteString(strconv.Itoa(i));
		}
		
	}
	return op.String()
}

func Printwhileloop(n int ) string{
	var op strings.Builder
	i:=1
	for i<=n {
		if i<n {	
			op.WriteString(strconv.Itoa(i));
			op.WriteString(",");
		}else{
			op.WriteString(strconv.Itoa(i));
		}
		i++
	}
	return op.String()
}
func main() {
	fmt.Println(PrintSerial(10))
}