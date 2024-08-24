#include <iostream>
using namespace std;
int main() {

  // Initialize Game Board
  char board[3][3] = {{' ', ' ', ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}};

  // Players
  const char playerX = 'X';
  const char playerO = 'O';
  char currentPlayer = playerX;
  char winner = ' ';

  int row = -1;
  int col = -1;

  for (int i = 0; i < 9; i++) {
    // Game board
    cout << "+---+---+---+" << endl;
    cout << "| " << board[0][0] << " | " << board[0][1] << " | " << board[0][2]
         << " |" << endl;
    cout << "+---+---+---+" << endl;
    cout << "| " << board[1][0] << " | " << board[1][1] << " | " << board[1][2]
         << " |" << endl;
    cout << "+---+---+---+" << endl;
    cout << "| " << board[2][0] << " | " << board[2][1] << " | " << board[2][2]
         << " |" << endl;
    cout << "+---+---+---+" << endl;

    // Winner Checking

    if (winner != ' ') {
      break;
    }

    // Get Player Input
    cout << "Current Player ::" << "Player " << currentPlayer << endl;
    while (true) {
      cout << "Enter Pos(Format Board[0-2][0-2]):: ";
      cin >> row >> col;
      if (row < 0 || row > 2 || col < 0 || col > 2) {
        cout << "Invalid Input" << endl;
      } else if (board[row][col] != ' ') {
        cout << "Already Placed !!" << endl;
      } else {
        break;
      }
      // Reset Values
      row = -1;
      col = -1;

      // Clear Cin Values
      cin.clear();             // Clear error flags
      cin.ignore(10000, '\n'); // Discard values
    }
    // Set Player Input
    board[row][col] = currentPlayer;

    // Change Player
    currentPlayer = (currentPlayer == playerX) ? playerO : playerX;

    // Check winning condition
    // Horizontal Validation
    for (int r = 0; r < 3; r++) {
      if (board[r][0] != ' ' && board[r][0] == board[r][1] &&
          board[r][1] == board[r][2]) {
        winner = board[r][0];
        break;
      }
    }
    // Vertical Validation
    for (int c = 0; c < 3; c++) {
      if (board[0][c] != ' ' && board[0][c] == board[1][c] &&
          board[1][c] == board[2][c]) {
        winner = board[0][c];
        break;
      }
    }
    // Diagonal Validation
    if (board[0][0] != ' ' && board[0][0] == board[1][1] &&
        board[1][1] == board[2][2]) {
      winner = board[0][0];
      break;
    }
    // Reverse Diagonal Validation
    else if (board[0][2] != ' ' && board[0][2] == board[1][1] &&
             board[1][1] == board[0][2]) {
      winner = board[0][2];
      break;
    }
  }
  if (winner != ' ') {
    cout << "================================================" << endl;
    cout << "Game Over" << "\nPlayer " << winner << " is the Winner" << endl;
    cout << "================================================" << endl;
  } else {
    cout << "================================================" << endl;
    cout << "Game Over" << "\n Game Tied" << endl;
    cout << "================================================" << endl;
  }
  return 0;
}
