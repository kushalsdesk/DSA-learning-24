#include <iostream>
using namespace std;

int min(int a, int b) {
  if (a < b) {
    return a;
  } else {
    return b;
  }
}
int array_min(int arr[], int start, int size, int minimum) {
  if (start == size) {
    return minimum;
  }

  minimum = min(minimum, arr[start]);
  return array_min(arr, start + 1, size, minimum);
}

int main() {
  int arr[] = {10, 12, 3, 46, 51};
  int sz = sizeof(arr) / sizeof(arr[0]);
  int min = array_min(arr, 0, sz, arr[0]);
  cout << " minimum = " << min << endl;
  return 0;
}
