import model.Wallet;

import java.util.Scanner;

public class Main {

    public static void displayMenu(Wallet wallet) {
        Scanner scanner = new Scanner(System.in);
        int choice;
        do {
            System.out.print("""
                    What do you want to do ?
                    1) See balance
                    2) Deposit
                    3) Withdraw
                    4) Never mind
                    """);
            System.out.print("Enter your choice: ");
            choice = scanner.nextInt();

            switch (choice) {
                case 1 -> System.out.println("Your balance is: " + wallet.getBalance() + " Ariary. \n");
                case 2 -> {
                    System.out.print("Enter the amount to deposit: ");
                    Integer deposit = scanner.nextInt();
                    wallet.deposit(deposit);
                }
                case 3 -> {
                    System.out.print("Enter the amount to withdraw: ");
                    Integer withdrawal = scanner.nextInt();
                    wallet.withdraw(withdrawal);
                }
                case 4 -> {
                    System.out.println("Bye !");
                }
            }
        }
        while (choice != 4);
    }

    public static void main(String[] args) {
        Wallet wallet = Wallet.getInstance();
       displayMenu(wallet);
    }
}
