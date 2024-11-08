/* Вы обслуживаете сайт и отслеживаете возникающие проблемы. Клиент получил ошибку после того, как попытался добавить свой пост в систему. Вы хотите понять, на каком из серверов эта ошибка произошла.
Есть n серверов, на i-й из них приходится a процентов запросов, из которых b процентов завершаются с ошибкой. Для каждого сервера найдите вероятность того, что ошибка произошла именно на нём. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(str))
	probArray := make([][2]float64, n)
	totalErrorProb := 0.0
	for i := 0; i < n; i++ {
		stringProbArray, _ := reader.ReadString('\n')
		nums := strings.Fields(stringProbArray)
		a, _ := strconv.ParseFloat(nums[0], 64)
		b, _ := strconv.ParseFloat(nums[1], 64)
		probArray[i][0] = a
		probArray[i][1] = b
		totalErrorProb += (a * b) / 100
	}
	for i := range probArray {
		fmt.Println(probArray[i][0] * probArray[i][1] / 100 / totalErrorProb)
	}
}
