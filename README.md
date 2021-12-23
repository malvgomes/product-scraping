# Scraper de Produtos

Este projeto consiste na implmentação de uma API de busca de dados em uma URL passada por parâmetro.

## Como funciona

- Para realizar a busca, basta enviar uma requisição `POST` para o endpoint `http://localhost:3000/product`, enviando no corpo 
  da requisição a url do produto. Exemplo:

```sh
  curl --location --request POST 'http://localhost:3000/product' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "url": "https://www.amazon.com.br/dp/B07SG8F1QF?ref=MarsFS_Echo_show8"
  }'
```
- Ao enviar a URL, é feita uma busca no banco de dados dos dados da URL buscada. Caso existam os dados no banco e o horário de 
  inserção foi há menos de uma hora, retorna os dados. Caso os dados não existam no banco ou foram inseridos há mais de uma 
  hora, realiza o scrape dos dados, e atualiza a entrada no banco com os dados atualizados (se existirem) e o horário atual.
- Caso a requisição tenha sido feita com sucesso, é retornado código 200 acompanhado de um JSON com os campos `title`, 
  `imageURL`, `price` (em centavos), `description` (se existir) e por fim, a `url` do produto. Exemplo de resposta:

```json
{
  "title": "Echo Show 8 (1ª Geração): Smart Speaker com tela de 8\" e Alexa - Cor Preta",
  "imageURL": "https://m.media-amazon.com/images/I/61xX62L2hGL._AC_SY300_SX300_.jpg",
  "price": 66405,
  "description": "",
  "url": "https://www.amazon.com.br/dp/B07SG8F1QF"
}
```

## Tecnologias utilizadas
- Golang versão 1.17
- Mysql versão 5.7.36

## Bibliotecas utilizadas
###API
- https://github.com/gocolly/colly - Faz o scraping das URLs
- https://github.com/go-gorp/gorp - Faz o mapeamento das colunas do banco para `structs` do Go
- https://github.com/go-chi/chi - Proporciona uma maneira simples de construir APIs REST em Go
- https://github.com/go-sql-driver/mysql Driver MYSQL para Go
- https://github.com/nleof/goyesql - Faz o mapeamento de um arquivo `.sql` para um `map[string]string` em Go

###Testes
- https://github.com/stretchr/testify - Conjunto de ferramentas para testes, como o `assert`, que foi utilizado no projeto
- https://github.com/DATA-DOG/go-sqlmock - Simula a conexão com o banco de dados
- https://github.com/golang/mock - Permite a criação de mocks das interfaces utilizadas no projeto
- https://github.com/jarcoal/httpmock - Permite o mock de respostas a requisições HTTP

# Ferramentas necessárias
- make
- docker-compose
- docker

# Como executar
- `make up`: inicia os containers e a aplicação
- `make down`: remove os containers
- `make db`: acessa o container do mysql, onde é possível executar queries e verificar se os dados foram realmente inseridos 
- `make test` : executa os testes unitários do projeto
- Envie requisições `POST` para o endpoint `localhost:3000/product`, como exemplificado neste documento e faça testes :). Execute o comando `make db` e verifique se os dados foram inseridos corretamente
