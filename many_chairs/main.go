/* Остап узнал, что в городе все продавцы продают только по одному стулу, и каждый покупатель готов купить не более одного. Всего он нашел N предложений, стоимость стула у 
i-го из продавцов равна Ai рублей, конечно, цены могут отличаться. Кроме того, он нашёл M потенциальных покупателей, каждый из которых может купить стул не дороже 
Bi рублей. При этом сам Остап может купить и продать любое количество товара.
Остап хочет получить наибольшую прибыль, поэтому он обратился за помощью к вам. Определите максимальную прибыль, которую он может получить.
Формат ввода
Первая строка входных данных содержит два целых числа N,M (1≤N,M≤10^5) — количество предложений и количество потенциальных покупателей соответственно.
Вторая строка содержит N целых чисел Ai (0≤Ai≤10^9) — цены, по которым продавцы готовы продавать стулья.
Третья строка содержит M целых чисел Bi(0≤Bi≤10^9 ) — суммы, которые потенциальные покупатели готовы отдать при покупке стула. */

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.Fields(line)[0])
	m, _ := strconv.Atoi(strings.Fields(line)[1])
	demand := make([]int, n)
	supply := make([]int, m)
	linea, _ := reader.ReadString('\n')
	demands := strings.Fields(linea)
	for i := 0; i < n; i++ {
		demand[i], _ = strconv.Atoi(demands[i])
	}
	lineb, _ := reader.ReadString('\n')
	supplys := strings.Fields(lineb)
	for i := 0; i < m; i++ {
		supply[i], _ = strconv.Atoi(supplys[i])
	}
	maxsum := 0
	sort.Slice(demand, func(i, j int) bool {
		return demand[i] < demand[j]
	})
	sort.Slice(supply, func(i, j int) bool {
		return supply[i] > supply[j]
	})
	for i := 0; i < int(math.Min(float64(n), float64(m))); i++ {
		if supply[i]-demand[i] > 0 {
			maxsum += supply[i] - demand[i]
		}
	}
	fmt.Println(maxsum)
}
