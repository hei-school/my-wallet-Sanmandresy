package model;

public class Wallet {
    private static Wallet instance;
    private Integer balance;
   private Wallet() {
        this.balance = 0;
   }

    public static Wallet getInstance() {
        if (instance == null) {
            instance = new Wallet();
        }
        return instance;
    }

    public double getBalance() {
        return this.balance;
    }

    public void deposit(Integer amount) {
       this.balance += amount;
    }

    public void withdraw(Integer amount) {
       if( amount > this.balance) {
           System.out.println("Balance not enough for this withdrawal");
       }
        else this.balance -= amount;
    }
}
