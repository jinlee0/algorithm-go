package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	N := readInt()
	valueArr := make([]int, N)
	sum := 0
	freqMap := map[int]int{}
	for i := 0; i < N; i++ {
		valueArr[i] = readInt()
		sum += valueArr[i]
		freqMap[valueArr[i]]++
	}
	sort.Ints(valueArr)

	mean := int(math.Round(float64(sum) / float64(N)))
	median := valueArr[N/2]
	_range := valueArr[N-1] - valueArr[0]

	valueFreqArr := make([][]int, len(freqMap))
	i := 0
	for value, freq := range freqMap {
		valueFreqArr[i] = make([]int, 2)
		valueFreqArr[i][0], valueFreqArr[i][1] = value, freq
		i++
	}
	sort.SliceStable(valueFreqArr, func(i, j int) bool {
		return valueFreqArr[i][1] < valueFreqArr[j][1]
	})
	mostFreq := valueFreqArr[len(valueFreqArr)-1][1]
	var modes []int
	for _, valueFreq := range valueFreqArr {
		if valueFreq[1] == mostFreq {
			modes = append(modes, valueFreq[0])
		}
	}
	sort.Ints(modes)
	mode := modes[0]
	if len(modes) > 1 {
		mode = modes[1]
	}

	fmt.Printf("%d\n%d\n%d\n%d", mean, median, mode, _range)

}

func readInt() int {
	scanner.Scan()
	atoi, _ := strconv.Atoi(scanner.Text())
	return atoi
}
