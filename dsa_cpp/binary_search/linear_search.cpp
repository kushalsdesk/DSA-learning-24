// Find a K in a given sorted array. using Loop.
// Complexity : O(n) <LinearSearch>
#include <iostream>
using namespace std;
#define MAX 5

int search_element(int arr[], int key) {
  for (int i = 0; i < MAX; i++) {
    if (arr[i] == key) {
      return i;
      break;
    }
  }
  return -1;
}

int main() {
  int k;
  int arr[MAX] = {1, 23, 45, 67, 89};
  cout << "enter key :: ";
  cin >> k;
  int pos = search_element(arr, k);
  if (pos == -1) {
    cout << "element not found" << endl;
  } else {
    cout << "element found at :" << pos << endl;
  }
  return 0;
}
