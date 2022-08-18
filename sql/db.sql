CREATE TABLE account (id INTEGER PRIMARY KEY AUTOINCREMENT, account_number TEXT NOT NULL, amount REAL DEFAULT 0);
INSERT INTO account (account_number, amount) VALUES ('1111-11', 1000);
SELECT * FROM account;