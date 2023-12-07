class Wallet {
  constructor(balance) {
    this.balance = balance;
  }

  getBalance() {
    return this.balance;
  }

  deposit(amount) {
    this.balance += amount;
  }

  withdraw(amount) {
    if (amount > this.balance) {
      console.log("Not enough in your wallet, cannot withdraw");
    } else this.balance -= amount;
  }
}

const wallet = new Wallet(0);
module.exports = wallet;
