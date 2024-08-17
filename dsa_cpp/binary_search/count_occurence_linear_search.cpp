// Done using LinearSearch
// Complexity : O(n) <LinearSearch>
#include <iostream>
using namespace std;

int main() {
  int sorted_arr[] = {1, 12, 12, 34, 34, 34, 56, 57, 58};
  int size = sizeof(sorted_arr) / sizeof(sorted_arr[0]);
  int key = 34;
  int occurence = 0;
  for (int i = 0; i < size; i++) {
    if (sorted_arr[i] == key) {
      occurence++;
    }
  }
  if (occurence == 0) {
    cout << "element not found" << endl;
  } else {
    cout << "element found " << occurence << " times" << endl;
  }
}
