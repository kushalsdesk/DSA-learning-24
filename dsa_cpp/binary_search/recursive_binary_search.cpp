// Find Element using binary search and recursion
// Complexity : O(logn) <BinarySearch>
#include <cmath>
#include <iostream>
using namespace std;
#define MAX 5

bool flag = false;

int b_search(int *arr, int f, int e, int k) {
  int mid;
  while (f <= e) {
    mid = ceil(f + e / 2);
    if (arr[mid] == k) {
      flag = true;
      break;
    } else if (arr[mid] > k) {
      e = mid - 1;
      b_search(arr, f, e, k);
    } else if (arr[mid] < k) {
      f = mid + 1;
      b_search(arr, f, e, k);
    }
  }
  return mid;
}

int main() {
  int key;
  int arr[MAX] = {12, 34, 56, 78, 99};
  int front = 0, end = MAX - 1;
  cout << "Enter Key To find : ";
  cin >> key;
  int pos = b_search(arr, front, end, key);
  if (flag) {
    cout << " Element Found at :" << pos << endl;
  } else {
    cout << "Element Not Found" << endl;
  }
  return 0;
}
