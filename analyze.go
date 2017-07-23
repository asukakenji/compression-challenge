package main

import "fmt"

func analyze() {
	records := readRecords()
	iRecords := integerizeRecords(integerizeEnumFixed, records)
	analysis := analyzeRecords(iRecords)
	fmt.Printf("Minimum Difference of Field 0: %d\n", analysis.MinDiffIField0)
	fmt.Printf("Maximum Difference of Field 0: %d\n", analysis.MaxDiffIField0)
	fmt.Println()
	fmt.Printf("Minimum of Field 2: %d\n", analysis.MinIField2)
	fmt.Printf("Maximum of Field 2: %d\n", analysis.MaxIField2)
	fmt.Printf("Difference between Maximum and Minimum: %d\n", analysis.DiffMaxMinIField2)
	fmt.Println()
	fmt.Printf("Minimum of Field 3: %d\n", analysis.MinIField3)
	fmt.Printf("Maximum of Field 3: %d\n", analysis.MaxIField3)
	fmt.Printf("Difference between Maximum and Minimum: %d\n", analysis.DiffMaxMinIField3)
	fmt.Println()
	fmt.Printf("Minimum of Field 4: %d\n", analysis.MinIField4)
	fmt.Printf("Maximum of Field 4: %d\n", analysis.MaxIField4)
	fmt.Printf("Difference between Maximum and Minimum: %d\n", analysis.DiffMaxMinIField4)
	fmt.Println()
	fmt.Printf("Minimum of Field 2/3/4: %d\n", analysis.MinIField234)
	fmt.Printf("Maximum of Field 2/3/4: %d\n", analysis.MaxIField234)
	fmt.Printf("Difference between Maximum and Minimum: %d\n", analysis.DiffMaxMinIField234)
	fmt.Println()
}
