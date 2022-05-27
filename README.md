# packet-monitoring

Packet monitoring with Prometheus and Grafana and Golang socket programming.

This is my Computer Networks course final project at CEIT/AUT.

## What is this project?
In this project, we create some agents that push their system status
to our prometheus client. Then will get those data and create metrics for
prometheus, and then we will monitor them with grafana.

The communication between agents and prometheus client is based
on TCP socket programming.

## How to start?
You can start the project on docker with following command:
```shell
docker-compose up -d 
```

Now you can access the following endpoints:
- Metrics: 0.0.0.0:1224/metrics
- Prometheus: 0.0.0.0:9090
- Grafana: 0.0.0.0:3000
  - username: admin
  - password: admin

## Configs
You can set your own configs:
```yaml
number_of_agents: 5  # node app number of agents
agent:  # agents configs
  host: "127.0.0.1"
  port: "8080"
  type: "tcp"
client: # client configs
  telemetry: # telemetry configs
    address: "0.0.0.0:1224"
    enabled: true
  host: "127.0.0.1"
  port: "8080"
  type: "tcp"
```
