from wallet import Wallet

def display_menu(wallet):
   while True:
       print("""
               What do you want to do ?
               1) See balance
               2) Deposit
               3) Withdraw
               4) Never mind
               """)
       choice = int(input("Enter your choice: "))

       if choice == 1:
           print("Your balance is: {} Ariary. \n".format(wallet.get_balance()))
       elif choice == 2:
           deposit = int(input("How much: "))
           wallet.deposit(deposit)
       elif choice == 3:
           withdrawal = int(input("How much: "))
           wallet.withdraw(withdrawal)
       elif choice == 4:
           print("Bye !")
           break

def main():
   wallet = Wallet()
   display_menu(wallet)

if __name__ == "__main__":
  main()