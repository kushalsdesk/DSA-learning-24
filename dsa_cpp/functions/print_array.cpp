#include <iostream>
using namespace std;
#define MAX 5
void print_array(int *arr) {
  int i;
  for (i = 0; i < MAX; i++) {
    cout << arr[i] << " ";
  }
}

int main() {
  int arr[MAX] = {1, 2, 3, 4, 5};
  print_array(arr);
  return 0;
}
