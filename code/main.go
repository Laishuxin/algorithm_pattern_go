package main

import (
	"code/Practice/StackPractice"
	"fmt"
)

func ExecuteStackPractice() {
	files := make([]string, 0, 20)
	//files, _ = StackPractice.GetAllFile("D:\\markdown_notebook\\algorithm_pattern_go\\code", files);
	files, _ = StackPractice.GetAllFile("D:\\markdown_notebook\\algorithm_pattern_go\\code\\List", files);
	for _, file := range files {
		fmt.Println(file)
	}
}

func ExecuteStackPractice2() {
	files := make([]string, 0, 20)
	files, _ = StackPractice.GetAllFileUsingStack("D:\\markdown_notebook\\algorithm_pattern_go\\code\\List", files);
	//files, _ = StackPractice.GetAllFileUsingStack("D:\\markdown_notebook\\algorithm_pattern_go\\code\\List\\ArrayList", files);
	for _, file := range files {
		fmt.Println(file)
	}
}

func main() {
	ExecuteStackPractice()
	ExecuteStackPractice2()
}
