# goexpert-cep2temp

Exercise description:
```
Objetivo: Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

Requisitos:

O sistema deve receber um CEP válido de 8 digitos
O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
O sistema deve responder adequadamente nos seguintes cenários:
Em caso de sucesso:
Código HTTP: 200
Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
Em caso de falha, caso o CEP não seja válido (com formato correto):
Código HTTP: 422
Mensagem: invalid zipcode
​​​Em caso de falha, caso o CEP não seja encontrado:
Código HTTP: 404
Mensagem: can not found zipcode
Deverá ser realizado o deploy no Google Cloud Run.
Dicas:

Utilize a API viaCEP (ou similar) para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
Utilize a API WeatherAPI (ou similar) para consultar as temperaturas desejadas: https://www.weatherapi.com/
Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula: F = C * 1,8 + 32
Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula: K = C + 273
Sendo F = Fahrenheit
Sendo C = Celsius
Sendo K = Kelvin
Entrega:

O código-fonte completo da implementação.
Documentação explicando como rodar o projeto em ambiente dev e production.
Testes automatizados demonstrando o funcionamento.
Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
Deploy realizado no Google Cloud Run (free tier) e endereço ativo para ser acessado.
```

### Requirements
- Valid API Key for WeatherAPI
- The API Key should be replaced in the config file (configs/config.env)

### How to run locally with Go
- Clone the repository
- Make sure you have Golang 1.21 or higher installed
- Open terminal in project folder
- Run `go run cmd/apid/main.go`

### How to run locally with docker-compose
- Clone the repository
- Make sure you have docker-compose installed
- Open terminal in project folder
- Make sure your OS and architecture matches with the ones in the Dockerfile (line 5)
- Run `docker-compose up --build`

### Tests
- You can locally test the http calls using the .http files in the http folder
- You can test the app running on Google Cloud Run in the url: https://goexpert-cep2temp-ec355himpa-rj.a.run.app/cep2temp/{CEP}

### Contact
- mateusmatinato@gmail.com