/* Рассмотрим три числа и упорядочим их по возрастанию.

Какое число будет стоять между двумя другими? */

#include <iostream>
#include <algorithm>

using namespace std;

int main() {
    int a[3];
    for (int i = 0; i < 3; i++) cin >> a[i];
    sort(a, a + 3);
    cout << a[1];
    return 0;
}