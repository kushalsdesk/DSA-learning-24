#include <iostream>
using namespace std;
#define MAX 6

bool check_prime(int num) {
  int pf = 0;
  bool checker;
  for (int j = 1; j < num; j++) {
    if (num % j == 0) {
      pf++;
    }
  }
  if (pf > 2) {
    checker = false;
  } else {

    checker = true;
  }
  return checker;
}
int how_many_prime(int *arr) {
  int count = 0;
  for (int i = 1; i < MAX; i++) {
    if (check_prime(arr[i])) {
      count++;
    }
  }
  return count;
}
int main() {
  int arr[MAX] = {12, 34, 56, 3, 1, 2};
  int no = how_many_prime(arr);
  cout << " TOTAL Prime :" << no << endl;
  return 0;
}
