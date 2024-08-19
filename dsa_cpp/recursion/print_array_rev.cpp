#include <iostream>
using namespace std;

void array_print(int arr[], int start, int size) {
  if (size == start) {
    return;
  } else {
    cout << arr[size - 1] << " ";
    array_print(arr, start, size - 1);
  }
}

int main() {
  int arr[] = {1, 2, 3, 4, 5};
  int sz = sizeof(arr) / sizeof(arr[0]);
  array_print(arr, 0, sz);
  return 0;
}
