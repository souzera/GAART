# BackEnd

## Inicializando o Bando de dados

1 - Rodar Docker compose

```bash
sudo docker compose up
```

2 - Sugestões de ferramentas para fazer o gerenciamento do banco de dados
    - PGAdmin
    - Dbeaver

3 - Criar extenção do UUID no banco de dados executando o comando SQL a seguir

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```