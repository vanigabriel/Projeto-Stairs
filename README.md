# Projeto-Stairs
 
Projeto desenvolvido como teste.
Sistema consistem em uma API de controle de uma adega, onde o usuário pode inserir, atualizar, listar e apagar vinhos de sua adega pessoal, em memória.
Implementado em Golang utilizando o framework GIN.

# Documentação API
https://documenter.getpostman.com/view/7995657/SVmtyKts?version=latest

# Features
 Identificador unico do vinho sendo o nome;
 Utilizado Mutex para evitar concorrencia nos maps;
 Govalidator para conferir campos obrigatórios;
 Log salvo em arquivo;
 TDD nas funções de controle dos vinhos;

# Rotas
 GET /adega : retorna todos os vinhos
	GET /ageda/vinho/:vinho : retorna um vinho específico
	POST /adega/vinho : adiciona um vinho 
	POST /adega/vinhos : adiciona vários vinhos de uma vez
	PUT /adega/vinho/:vinho : atualiza/insere um vinho específico
	PUT /adega/vinho/:vinho/garrafas/:quantidade : atualiza quantidade de garrafas de um vinho específico
	DELETE /adega/vinho/:vinho : apaga um vinho específico
 
# Backlog
  Coisas que eu implementaria a mais porém que infelizmente a disponibilidade não permitiu
- Busca dinâmica aos vinhos 
- Utilizar o mongo para persistir os dados
- TDD das rotas em si
- Melhor utilização de Clean Architecture 
