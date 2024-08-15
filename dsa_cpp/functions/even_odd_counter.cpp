#include <iostream>
using namespace std;
#define MAX 7
void counter(int *arr) {
  int even = 0, odd = 0;
  for (int i = 0; i < MAX; i++) {
    if (arr[i] % 2 == 0) {
      even++;
    } else {
      odd++;
    }
  }
  cout << "Total Even = " << even << " Total Odd = " << odd << endl;
}

int main() {
  int arr[MAX] = {12, 34, 1, 56, 7, 99, 23};
  counter(arr);
  return 0;
}
