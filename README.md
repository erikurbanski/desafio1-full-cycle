# desafio1-full-cycle
Desafio de número 1 para Imersão Full Stack &amp;&amp; Full Cycle

Basta executar o comando: go run main.go 

Depois executar as chamadas conforme proposto no desafio através do Postman ou Insomnia:

POST http://localhost:8000/bank-accounts/
{
	"account_number": "4444-44"
}

POST http://localhost:8000/bank-accounts/transfer
{
	"from": "1111-11",
	"to": "2222-22",
	"amount": 100
}

GET  http://localhost:8000/bank-accounts/

Banco de dados já esta incluso no projeto.
