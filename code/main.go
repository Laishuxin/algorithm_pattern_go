package main

import (
	"code/Practice/QueuePractice"
	"code/Practice/StackPractice"
	"fmt"
)

// ==== Stack Practice ====
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

// ==== Queue Practice ====
func ExecuteQueuePractice1() {
	files := make([]string, 0, 20)
	files, _ = QueuePractice.GetAllFiles("D:\\markdown_notebook\\algorithm_pattern_go\\code\\List", files);
	for _, file := range files {
		fmt.Println(file)
	}
}

func main() {
	// ==== Stack Practice ====
	//ExecuteStackPractice()
	//ExecuteStackPractice2()


	// ==== Queue Practice ====
	ExecuteQueuePractice1()
}
