#include <iostream>
using namespace std;

void array_print(int arr[], int start, int size) {
  if (start == size) {
    return;
  } else {
    cout << arr[start] << " ";
    array_print(arr, start + 1, size);
  }
}

int main() {
  int arr[] = {1, 2, 3, 4, 5};
  int sz = sizeof(arr) / sizeof(arr[0]);
  array_print(arr, 0, sz);
  return 0;
}
