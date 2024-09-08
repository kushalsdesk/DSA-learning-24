#include <iostream>
#include <stack>
using namespace std;

int main() {
  stack<char> stck;
  string input = "(()(())()";

  for (char c : input) {
    if (!stck.empty() && c == ')' && stck.top() == '(') {
      stck.pop();
    } else {
      stck.push(c);
    }
  }

  if (stck.empty()) {
    cout << "Balanced" << endl;
  } else {
    cout << "Not Balanced" << endl;
  }
}
