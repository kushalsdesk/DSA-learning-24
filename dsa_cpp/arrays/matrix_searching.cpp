// Searching an element in a sorted matrix
#include <iostream>
using namespace std;

int main() {
  int matrix[3][3] = {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}};
  int x, found = 0;
  cout << "Enter the element this program needs to find: ";
  cin >> x;
  for (int i = 0; i < 3; i++) {
    for (int j = 0; j < 3; j++) {
      if (matrix[i][j] == x) {
        cout << "Element Found At Matrix[" << i << "][" << j << "]" << "\n";
        found++;
      }
    }
  }
  if (found == 0) {
    cout << "Element Not Found !!" << endl;
  }
  return 0;
}
