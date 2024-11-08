/* На кольцевой трассе расположены n гоночных автомобилей. Наша цель оценить зрелищность заезда количество обгонов в будущей гонке!
Для простоты будем считать, что все автомобили стартуют в одно время из одной точки и движутся с постоянной скоростью: скорость i-го автомобиля 
Определите количество обгонов, которые совершит автомобиль с номером 1, за время t, если длина кольцевого трека равна s.
Обратите внимание, что в один момент времени автомобиль может совершить сразу несколько обгонов. Автомобили, находящиеся в одной точке в момент времени 0 и t, не совершают обгонов. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cus_floor (x float64) int{
	if float64(int(x)) == x{
		return int(x) - 1
	} else{
		return int(x)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	numStrings := strings.Fields(line)
	numbers := make([]int, len(numStrings))
	for i, numStr := range numStrings {
		num, _ := strconv.Atoi(numStr)
		numbers[i] = num
	}
	n, t, s := numbers[0], numbers[1], numbers[2]
	speedStr, _ := reader.ReadString('\n')
	speedStr = strings.TrimSpace(speedStr)
	numSpeeds := strings.Fields(speedStr)
	array := make([]int, len(numSpeeds))
	for i, numStr := range numSpeeds {
		num, _ := strconv.Atoi(numStr)
		array[i] = num
	}
	count := 0
	target := array[0]
	for i := 1; i < n; i++ {
		if target > array[i]{
			count += cus_floor(float64(t)*(float64(target) - float64(array[i]))/float64(s))
		}
	}
	fmt.Print(count)
}
