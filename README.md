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

### Promtail
O Promtail permite realizar experimentos, validações e alterações nos logs enviados para o Grafana Loki sem a 
necessidade de inicializar o servidor. Para instalar o Promtail localmente, utilize o seguinte comando:

```bash
brew install promtail
```

Para fins de teste, insira logs no arquivo `config/log.txt e execute o comando abaixo para validar o formato dos logs:

```bash
cat log.txt | promtail --config.file ./promtail-config.yaml --stdin --dry-run --inspect
```

Este comando lê o conteúdo de `log.txt e executa o Promtail em modo de teste (--dry-run), exibindo informações detalhadas sobre o processamento dos logs sem enviá-los ao Loki. Assim, você pode inspecionar e ajustar o formato dos logs antes de enviá-los para produção.

#### Explicação
- **`--stdin`**: Lê a entrada diretamente do stdin em vez de um arquivo.
- **`--dry-run`**: Executa o Promtail em modo de teste, o que significa que os logs são processados, mas não enviados ao Loki.
- **`--inspect`**: Fornece uma saída detalhada sobre o processamento de cada entrada de log, ajudando a identificar problemas de formatação ou incompatibilidades de configuração.
  Essa configuração permite que você inspecione e ajuste o formato dos seus logs localmente antes de enviá-los para o ambiente de produção do Loki.

#### Exemplo de Fluxo de Trabalho
1. Edite `config/log.txt` para incluir logs de exemplo.
2. Ajuste `promtail-config.yaml` conforme necessário para análise e rotulagem dos logs.
3. Execute o comando acima para ver a saída detalhada e garantir que tudo esteja configurado corretamente.

Happy logging!

![Dry Run](documentation/images/dry-run.png)