# rabbitMQ
 1. запустити ребіт локально docker run -d --hostname my-rabbit --name some-rabbit rabbitmq:3-management
 2. відкрити термінал в папці rabbitMQ/cmd
 3. відкрити 2 консолі
 4. на першій прописати go run receive.go
 5. на другій - go run send.go
 6. результат по рівню infо, error, debug буде записаний в файл logs/all.log та відправлено на ребіт, можна переконатись за посиланням http://localhost:15672, а логи рівня error - в консолі
