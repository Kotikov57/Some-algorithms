/* Задан массив a размера n. Необходимо посчитать количество уникальных элементов в данном массиве.
Элемент называется уникальным, если встречается в массиве ровно один раз. */

#include <iostream>
#include <map>
#include <vector>

using namespace std;

int main() {
    map <int, int> a;
    int n;
    cin >> n;
    vector<int> b(n);
    int count =  0;
    for (int i = 0; i < n; i++) {
        cin >> b[i];
        a[b[i]]++;
    }
    for (int i = 0; i < n; i++) {
        if (a[b[i]] == 1) {
            count++;
        }
    }
    cout << count;
}