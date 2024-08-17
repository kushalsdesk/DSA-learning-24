#include <cmath>
#include <iostream>
#include <ostream>
using namespace std;
bool flag = false; // Flag variable to find if element found or not
#define size_of_array 8
int main() {
  int arr[size_of_array] = {1, 4, 6, 18, 34, 46, 65, 69};
  int key, mid;
  cout << "Enter Key to find:";
  cin >> key;
  int front = 0, end = size_of_array - 1;
  while (front <= end) {
    mid = ceil((front + end) / 2); // Ceiling Value
    if (arr[mid] == key) {
      flag = true;
      break;
    } else if (arr[mid] > key) {
      end = mid - 1;
    } else if (arr[mid] < key) {
      front = mid + 1;
    }
  }
  if (flag) {
    cout << "Element Fount at " << mid << endl;
  } else {
    cout << "Element Not Found" << endl;
  }
  return 0;
}
