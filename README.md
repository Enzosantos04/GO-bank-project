## Enzo's Bank - CLI Application

This repository contains a simple command-line interface (CLI) banking application written in Go. It allows users to check their balance, deposit, and withdraw money from a simulated account. The data is stored in a text file for persistence.

## Features

- Check balance
- Deposit money
- Withdraw money
- Persistent balance storage in `balance.txt`

## How to Run

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/your-repository.git
   cd your-repository
   ```

2. Run the code with:

   ```sh
   go run .
   ```

## Code Explanation

### Main Modules

- `writeBalancetoFile(balance float64)`: Writes the balance to the `balance.txt` file.
- `getBalanceFromFile() float64`: Retrieves the balance from the `balance.txt` file.
- `main()`: Runs the interactive bank loop, allowing users to choose between checking balance, depositing, or withdrawing money.

### Example Usage

```sh
Welcome to Enzo's Bank
What do you want to do?
1. Check balance
2. Deposit money
3. Withdraw money
4. Exit
Your Choice: 2
Deposit certain amount: 100
Your new balance is: 100
```

## License

This project is distributed under the MIT License. See the `LICENSE` file for more information.


