# empreGo-web

Plataforma agregadora e gerenciadora de vagas de estágio em tecnologia, com notificações em tempo real via Discord.

**[Acessar plataforma](https://front-empre-go.vercel.app)**

---

## Screenshots

Home 
<img width="1366" height="768" alt="vagas" src="https://github.com/user-attachments/assets/a1d3ecad-5153-481e-84d0-561085adff7b" />

Candidaturas Exemplo
<img width="1366" height="768" alt="candidaturas" src="https://github.com/user-attachments/assets/d7a66701-206e-4f1f-8ed2-724ba2daeb94" />


---

## Como funciona

```
[Worker - scraping LinkedIn/Gupy]
          |
          ├── salva vaga ──────────────► [SQLite]
          |                                 |
          ▼                                 ▼
      [ NATS ]                          [API HTTP]
      vagas.fullstack                   consulta vagas
      vagas.backend        	          e candidaturas
      vagas.frontend ...
          |
          ▼
  [Subscriber - Discord Bot]
  envia para canais específicos
```

O sistema é composto por **3 executáveis independentes**:

- **`cmd/api`** — API HTTP, autenticação via OAuth + cookie de sessão, gerenciamento de candidaturas
- **`cmd/worker`** — scraping do LinkedIn e Gupy a cada 2h (via crontab). Publica a vaga no NATS e, se o envio for bem-sucedido, salva no SQLite
- **`cmd/subscriber`** — escuta os tópicos do NATS e distribui as vagas nos canais correspondentes do Discord

---

## Tecnologias

| Camada | Tecnologia |
|---|---|
| Backend | Go |
| Frontend | React + Vite |
| Banco de dados | SQLite |
| Mensageria | NATS |
| Cache | Redis *(em breve)* |
| Infraestrutura | GCP e2-micro |
| Deploy frontend | Vercel |

---

## Funcionalidades

- Agregação automática de vagas do LinkedIn e Gupy
- Notificações em tempo real por categoria no Discord
- Autenticação via OAuth (GitHub)
- Sistema de candidaturas (acompanhe em qual vaga você aplicou)
- Atualização automática a cada 2h (08h–20h)
- Paginação e cache com Redis *(em breve)*
- Bot do Telegram *(planejado)*

---

## Rodando localmente

### Pré-requisitos

- Go 1.25+
- Docker e Docker Compose
- Conta no GitHub 
- Bot do Discord configurado
- Goose para migrações
- OAuth app no GitHub

### Setup

```bash
git clone https://github.com/1-AkM-0/empreGo-web
cd empreGo-web

# Sobe NATS via Docker
docker-compose up -d

# Roda as migrations
goose sqlite3 -dir migrations/ vagas.db up

# As variáveis de ambiente são lidas via os.Getenv, ou seja, você tem que passar como argumento do executável
# Recomendado: crie scripts locais para cada executável (não versione esses arquivos)

# Exemplo: run-api.sh
export GITHUB_CLIENT_ID=xxx 
export GITHUB_CLIENT_SECRET=xxx
export SESSION_SECRET=xxx 
go run cmd/api/main.go


# Exemplo: run-worker.sh
go run cmd/worker/main.go

# Exemplo: run-subscriber.sh
BOT_TOKEN=xxx 
go run cmd/subscriber/main.go

# Adicione os scripts ao .gitignore para não vazar credenciais
echo "run-*.sh" >> .gitignore
```

---

## Arquitetura/Decisões técnicas

**Por que NATS?** Toda a infra roda numa VM e2-micro na GCP (2 vCPUs compartilhados, 1GB RAM). O NATS é extremamente leve e performático, Kafka ou RabbitMQ seriam overhead desnecessário aqui.

**Por que SQLite?** Semelhante ao motivo do NATS, o SQLite é um banco relacional leve e que não precisa de um servidor como o PostgreSQL, economizando recursos da VM.

**Por que cookie de sessão em vez de JWT?** JWT faz sentido em arquiteturas com múltiplos serviços precisando validar o token de forma independente. Aqui tenho um único serviço de autenticação, então cookie de sessão resolve com menos complexidade.

**Por que 3 executáveis separados?** Separação de responsabilidades: posso atualizar o subscriber sem derrubar a API, e no futuro adicionar um subscriber do Telegram sem tocar no worker ou na API.

---

## Roadmap

- [ ] Paginação nas listagens e candidaturas
- [ ] Cache com Redis
- [ ] Persistência de mensagens com NATS JetStream
- [ ] Filtro de vagas por tecnologia/palavra-chave
- [ ] Bot do Telegram
- [ ] Adicionar testes

---

## Licença

MIT
