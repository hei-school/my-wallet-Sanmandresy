class Wallet:
    def __init__(self):
       self.balance = 0

    def get_balance(self):
       return self.balance

    def deposit(self, amount):
       self.balance += amount

    def withdraw(self, amount):
       if amount > self.balance:
           print("Balance not enough for this withdrawal")
       else:
           self.balance -= amount