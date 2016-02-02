package main
import (
	"os"
	"strings"
	"strconv"
	"fmt"
	"bufio"
	"math"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	n, _ := strconv.Atoi(input)
	var lrsum, rlsum int
	arr := make([][]int, n)
	for i:=0; i<n; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		ar := strings.Split(input, " ")
		arr[i] = make([]int, n)
		for j:=0; j<n; j++ {
			arr[i][j],_ = strconv.Atoi(ar[j])
		}
		lrsum += arr[i][i]
		rlsum += arr[i][n-i-1]
	}

	fmt.Println(math.Abs(float64(lrsum)-float64(rlsum)))

}
