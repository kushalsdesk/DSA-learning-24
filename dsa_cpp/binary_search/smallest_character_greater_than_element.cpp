// Find the Smallest character greater than element present in a sorted
// character array
// Complexity : O(n) <LinearSearch>

#include <iostream>
using namespace std;
bool found;

void find_scgtep(char arr[], int size, char key) {
  for (int i = 0; i < size; i++) {
    if (int(arr[i]) > int(key)) {
      found = true;
      cout << "Smallest character greater than element present is " << arr[i]
           << endl;
      break;
    }
  }
  if (found != true) {
    cout << "Not found" << endl;
  }
}

int main() {
  char ch, arr[] = {'a', 'b', 'c', 'f', 'k', 'l'};
  int size = sizeof(arr) / sizeof(arr[0]);
  cout << " Enter The element to find :";
  cin >> ch;
  find_scgtep(arr, size, ch);
  return 0;
}
