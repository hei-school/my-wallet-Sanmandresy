const wallet = require("./wallet");
const prompt = require("prompt-sync")();

const displayMenu = (wallet) => {
  let choice;
  do {
    console.log(`
      What do you want to do ?
      1) See your balance
      2) Deposit
      3) Withdraw
      4) Never mind
    `);
    choice = +prompt("Enter your choice: ");
    switch (choice) {
      case 1:
        console.log(`Your balance is ${wallet.getBalance()} ariary`);
        break;
      case 2:
        amount = +prompt("How much ? ");
        wallet.deposit(amount);
        break;
      case 3:
        amount = +prompt("How much ? ");
        wallet.withdraw(amount);
        break;
      case 4:
        console.log("Bye !");
        break;
    }
  } while (choice != 4);
};

displayMenu(wallet);
