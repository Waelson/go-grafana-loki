# Go + Uber Zap + Promtail + Loki + Grafana


### Solução

### Arquitetura
![Architecture](documentation/images/architecture.png)

### Visualização de logs
![Dashboard](documentation/images/dashboard.png)

### Executando
A aplicação está configurada para ser executada com Docker Compose. Para iniciá-la, entre no diretório raiz da aplicação e execute o seguinte comando:

`
docker-compose up --build
`

### Simulação

### Promtail em modo Dry Run
O Promtail permite que você faça experimentos, validações e alteração dos logs que são enviados para o Grafana Loki, sem a necessidade de inicializar o servidor.
Para isso, instale o Promtail localmente utilizando o comando abaixo:

`
brew install promtail
`

Para fins de testes, alimente o arquivo `config/log.txt` com os log e execute o comando abaixo para validar o formato do log.

`
cat log.txt | promtail --config.file ./promtail-config.yaml --stdin --dry-run --inspect
`

![Dry Run](documentation/images/dry-run.png)