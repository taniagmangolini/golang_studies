package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"bufio"
	"os"
	"math"
)

func splitNumbersIntoPartitions(numbers []int, n int) [][]int {
	partitionSize := n / 4
	partitions := [][]int{
		numbers[:partitionSize],
		numbers[partitionSize : 2*partitionSize],
		numbers[2*partitionSize : 3*partitionSize],
		numbers[3*partitionSize:],
	}
	return partitions

}

func sortNumbersIntoPartition(wg *sync.WaitGroup, index int, partitions [][]int){
	defer wg.Done()
	fmt.Println("Sorting subarray:", partitions[index])
	sort.Ints(partitions[index])
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter integers numbers separated by spaces:")
	input, _ := reader.ReadString('\n')
	inputStrings := strings.Fields(input) //split string based on whitespace

	var numbers []int
	for _, v := range inputStrings {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Invalid number.")
			return
		}
		numbers = append(numbers, num)
	}

	n := len(numbers)
	if n < 4 {
		fmt.Println("Please enter at least 4 numbers.")
		return
	}

	partitions := splitNumbersIntoPartitions(numbers, n)

	var wg sync.WaitGroup
	for i := range partitions {
		wg.Add(1)
		go sortNumbersIntoPartition(&wg, i, partitions)
	}
	wg.Wait()

	result := mergeAndSortPartitions(partitions)

	fmt.Println("Sorted numbers:", result)
}

func mergeAndSortPartitions(partitions [][]int) []int {
	mergedPartitions := []int{}
	indices := make([]int, len(partitions)) 
	fmt.Println("Sorting:", partitions)
	for {
		minIndex := -1
		minValue := math.MaxInt64
		for i, partition := range partitions {
			if indices[i] < len(partition) && partition[indices[i]] < minValue {
				minIndex = i
				minValue = partition[indices[i]]
			}
		}
		if minIndex == -1 {
			break
		}
		mergedPartitions = append(mergedPartitions, minValue)
		indices[minIndex]++
	}
	return mergedPartitions
}
