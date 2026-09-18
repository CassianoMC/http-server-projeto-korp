# http-server-projeto-korp

Projeto desenvolvido para o desafio técnico da Korp: um serviço HTTP em Go, containerizado com Docker, exposto por um proxy reverso Nginx, monitorado com Prometheus e Grafana, e provisionado por Ansible.

## Visão geral

A aplicação expõe o endpoint `GET /projeto-korp`, que retorna um JSON com o nome do projeto e o horário atual em UTC. O ambiente completo roda com Docker Compose e inclui:

- Aplicação Go na porta interna `8080`
- Nginx recebendo requisições na porta `80`
- Prometheus coletando métricas na porta `9090`
- Grafana exibindo dashboard na porta `3000`
- Ansible automatizando instalação do Docker, build e subida dos containers

## Arquitetura

Fluxo principal:

```text
Cliente
  |
  | GET http://localhost/projeto-korp
  v
Nginx :80
  |
  | proxy_pass
  v
Aplicação Go :8080
  |
  | /metrics
  v
Prometheus :9090
  |
  v
Grafana :3000
```

## Tecnologias

- Go `1.25.0`
- Docker e Docker Compose
- Nginx
- Prometheus
- Grafana
- Ansible

## Estrutura do projeto

```text
.
├── README.md
├── Documentação/
│   └── prints/
│       ├── 01-docker-compose-ps.png
│       ├── 02-curl-endpoint.png
│       ├── 03-endpoint-navegador.png
│       ├── 04-grafana-query-disponibilidade.png
│       ├── 05-grafana-query-requisicoes.png
│       ├── 06-grafana-dashboard-disponivel.png
│       ├── 07-grafana-dashboard-indisponivel.png
│       └── 08-ansible-playbook-validacao.png
└── http-server-projeto-korp/
    ├── app/
    │   ├── Dockerfile
    │   ├── go.mod
    │   ├── go.sum
    │   └── main.go
    ├── ansible/
    │   ├── inventory.ini
    │   ├── playbook.yml
    │   └── requirements.yml
    ├── grafana/
    │   ├── dashboards-json/
    │   └── provisioning/
    ├── nginx/
    │   └── http-server-projeto-korp.conf
    ├── prometheus/
    │   └── prometheus.yml
    └── docker-compose.yml
```

## Como executar com Docker Compose

Entre na pasta do ambiente:

```bash
cd http-server-projeto-korp
```

Suba todos os serviços:

```bash
docker compose up --build -d
```

Verifique os containers:

```bash
docker compose ps
```

Teste o endpoint pela porta pública do Nginx:

```bash
curl http://localhost/projeto-korp
```

Resposta esperada:

```json
{
  "nome": "Projeto Korp",
  "horario": "2026-09-18T19:00:00Z"
}
```

## Endpoints

| Serviço | URL | Descrição |
|---|---|---|
| Aplicação via Nginx | `http://localhost/projeto-korp` | Retorna o JSON principal do desafio |
| Métricas da aplicação | `http://localhost/metrics` | Exibe métricas Prometheus via Nginx |
| Prometheus | `http://localhost:9090` | Interface para consultas PromQL |
| Grafana | `http://localhost:3000` | Dashboard de monitoramento |

## Métricas

A aplicação registra a métrica customizada:

```text
http_requests_total{method="GET",status="OK"}
```

Ela contabiliza requisições recebidas pelo endpoint `/projeto-korp`, separando por método HTTP e status textual.

O Prometheus coleta as métricas do container da aplicação usando o nome do serviço Docker:

```yaml
targets: ["http-server-projeto-korp:8080"]
```

## Dashboard Grafana

O Grafana é provisionado automaticamente com:

- Datasource Prometheus
- Dashboard `Projeto Korp - Monitoramento`
- Painel `Disponibilidade do Serviço`
- Painel `Requisições por Minuto`

Depois de subir o ambiente, acesse:

```text
http://localhost:3000
```

Login padrão do Grafana:

```text
usuário: admin
senha: admin
```

Na primeira entrada, o Grafana pode pedir a troca de senha.

## Automação com Ansible

O playbook em `http-server-projeto-korp/ansible/playbook.yml` automatiza:

- Atualização do cache `apt`
- Instalação de dependências para o repositório Docker
- Configuração da chave GPG e repositório oficial Docker
- Instalação do Docker Engine, CLI e plugin Compose
- Inicialização e habilitação do serviço Docker
- Inclusão do usuário no grupo `docker`
- Build e subida do ambiente via Docker Compose
- Validação do endpoint `http://localhost/projeto-korp`

Para executar:

```bash
cd http-server-projeto-korp/ansible
ansible-galaxy collection install -r requirements.yml
ansible-playbook -i inventory.ini playbook.yml --ask-become-pass
```

## Evidências de execução

Os prints abaixo documentam a execução do projeto de ponta a ponta.

### Containers ativos

[![Containers ativos no Docker Compose](https://github.com/CassianoMC/http-server-projeto-korp/raw/main/Documenta%C3%A7%C3%A3o/prints/01-docker-compose-ps.png)](/CassianoMC/http-server-projeto-korp/blob/main/Documenta%C3%A7%C3%A3o/prints/01-docker-compose-ps.png)

### Endpoint respondendo no terminal

[![Resposta do endpoint via curl](https://github.com/CassianoMC/http-server-projeto-korp/raw/main/Documenta%C3%A7%C3%A3o/prints/02-curl-endpoint.png)](/CassianoMC/http-server-projeto-korp/blob/main/Documenta%C3%A7%C3%A3o/prints/02-curl-endpoint.png)

### Queries dos painéis no Grafana

[![Query de disponibilidade no Grafana](https://github.com/CassianoMC/http-server-projeto-korp/raw/main/Documenta%C3%A7%C3%A3o/prints/04-grafana-query-disponibilidade.png)](/CassianoMC/http-server-projeto-korp/blob/main/Documenta%C3%A7%C3%A3o/prints/04-grafana-query-disponibilidade.png)

[![Query de requisições por minuto no Grafana](https://github.com/CassianoMC/http-server-projeto-korp/raw/main/Documenta%C3%A7%C3%A3o/prints/05-grafana-query-requisicoes.png)](/CassianoMC/http-server-projeto-korp/blob/main/Documenta%C3%A7%C3%A3o/prints/05-grafana-query-requisicoes.png)

### Dashboard Grafana

[![Dashboard Grafana com serviço disponível](https://github.com/CassianoMC/http-server-projeto-korp/raw/main/Documenta%C3%A7%C3%A3o/prints/06-grafana-dashboard-disponivel.png)](/CassianoMC/http-server-projeto-korp/blob/main/Documenta%C3%A7%C3%A3o/prints/06-grafana-dashboard-disponivel.png)

[![Dashboard Grafana com serviço indisponível](https://github.com/CassianoMC/http-server-projeto-korp/raw/main/Documenta%C3%A7%C3%A3o/prints/07-grafana-dashboard-indisponivel.png)](/CassianoMC/http-server-projeto-korp/blob/main/Documenta%C3%A7%C3%A3o/prints/07-grafana-dashboard-indisponivel.png)

### Validação com Ansible

[![Execução do Ansible validando o endpoint](https://github.com/CassianoMC/http-server-projeto-korp/raw/main/Documenta%C3%A7%C3%A3o/prints/08-ansible-playbook-validacao.png)](/CassianoMC/http-server-projeto-korp/blob/main/Documenta%C3%A7%C3%A3o/prints/08-ansible-playbook-validacao.png)

## Status do desafio

- [x] Serviço HTTP em Go
- [x] Endpoint `/projeto-korp` com JSON dinâmico em UTC
- [x] Dockerfile da aplicação
- [x] Docker Compose com rede bridge
- [x] Proxy reverso com Nginx
- [x] Métricas Prometheus
- [x] Dashboard Grafana provisionado
- [x] Automação com Ansible