# desafio-docker-nginx-node-mysql
Desafio do curso FullCycle, trabalhando com docker.

### Descrição do desafio
> A ideia principal é que quando um usuário acesse o nginx, o mesmo fará uma chamada em nossa aplicação node.js. Essa aplicação por sua vez adicionará um registro no banco de dados mysql, cadastrando um nome na tabela people.

__O retorno da aplicação node.js para o nginx deverá ser:__
```html
<h1>Full Cycle Rocks!</h1>

- Nome cadastrado no banco de dados.
```

### Requisitos
1. Toda a aplicação deve estar disponível na porta 8080.

  
### Para rodar
```
docker-compose up [-d]
```
<br/>
<br/>
