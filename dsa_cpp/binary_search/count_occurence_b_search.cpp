// Done using Binary Search, for O(logn) complexity.

#include <iostream>
using namespace std;

int b_search_last(int arr[], int f, int e, int k) {
  int res = -1;
  while (f <= e) {
    int mid = ceil((f + e) / 2);
    if (arr[mid] == k) {
      res = mid;   // Holding the result!
      f = mid + 1; // rechecking if there is more of the same element
    } else if (arr[mid] > k) {
      e = mid - 1;
    } else if (arr[mid] < k) {
      f = mid + 1;
    }
  }
  return res;
}
int b_search_first(int arr[], int f, int e, int k) {
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
  int arr[] = {12, 20, 34, 34, 34, 56, 78, 90};
  int key = 34;
  int n = sizeof(arr) / sizeof(arr[0]);

  int index_first = b_search_first(arr, 0, n - 1, key);
  int index_last = b_search_last(arr, 0, n - 1, key);

  if (index_first == -1 && index_last == -1) {
    cout << " Element not Found " << endl;
  } else {
    int total_encounter =
        (index_last - index_first) + 1; // Add 1 as array starts from 0.
    cout << "Element found " << total_encounter << " times" << endl;
  }

  return 0;
}
