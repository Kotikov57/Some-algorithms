/* Для отрисовки календаря в студенческом проекте было решено выделить функциональность форматирования в отдельный модуль.
Параметрами модуля (формально для функции, которую можно будет импортировать из модуля) будут количество дней в месяце и название дня недели, на который приходится первое число месяца, записанное на английском языке.
Выведите все дни месяца по неделям, дополнив первую неделю пустыми значениями, если это требуется.
Формат ввода
В единственной строке входных данных записаны две величины: nDays (28≤nDays≤31) - количество дней в месяце;
weekday∈[Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday] — день недели, на который приходится первое число месяца.
Формат вывода
Выведите k строк (4≤k≤6), в i-й строке выведите даты, которые попадают на i-ю неделю месяца.
При выводе следуйте следующим правилам:
все строки, кроме последней, должны иметь ровно 7 элементов (в последней строке также может оказаться 7 элементов);
при выводе дней с номерами от 1 до 9 следует добавить символ точки (.) перед цифрой;
при выводе дней первой недели перед первым числом используйте две точки (..)
 */
 
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
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.Fields(line)[0])
	count := 0
	weekday := strings.Fields(line)[1]
	switch weekday {
	case "Monday":
		count = 0
	case "Tuesday":
		count = 1
	case "Wednesday":
		count = 2
	case "Thursday":
		count = 3
	case "Friday":
		count = 4
	case "Saturday":
		count = 5
	case "Sunday":
		count = 6
	}
	calendar := strings.Repeat(".. ", count)
	array := []string{}
	for i := 1; i < 10; i++ {
		str := fmt.Sprintf(".%d ", i)
		array = append(array, str)
	}
	for i := 10; i <= n; i++ {
		str := fmt.Sprintf("%d ", i)
		array = append(array, str)
	}
	i := count + 1
	for _, val := range array {
		calendar += val
		if i%7 == 0 {
			calendar += "\n"
		}
		i++
	}
	fmt.Println(calendar)
}
