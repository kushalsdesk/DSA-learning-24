// find the first index where the element is found.
// Complexity : O(logn) <Due To Binary Search>

#include <cmath>
#include <iostream>
using namespace std;

int b_search(int arr[], int f, int e, int k) {
  int res = -1;
  while (f <= e) {
    int mid = ceil((f + e) / 2);
    if (arr[mid] == k) {
      res = mid;   // Holding the result!
      e = mid - 1; // rechecking if there is more of the same element
    } else if (arr[mid] > k) {
      e = mid - 1;
    } else if (arr[mid] < k) {
      f = mid + 1;
    }
  }
  return res;
}

int main() {
  int arr[] = {12, 20, 34, 34, 56, 78, 90};
  int key = 34;
  int n = sizeof(arr) / sizeof(arr[0]);

  int pos = b_search(arr, 0, n - 1, key);
  if (pos == -1) {
    cout << "element not present" << endl;
  } else {
    cout << " Element found at " << pos << endl;
  }
  return 0;
}
