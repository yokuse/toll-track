# Project brief - traffic toll calculator
Building different microservices to handle minute tasks.
- microservice 1: receive gps coordinates from vehicles's OBU (gps coords will be mocked)
- microservice 2: Calculate distance of vehicle to something like erp
- microservice 3: invoicer to create invoices, maybe with ui
- microservice 4: invoice calculator based on distance

## Tools
Kafka to send messages to topics for OBU receiver and distance calculator
websocket connection for obu
Some DB to store the data for invoicer
Prometheus, grafana (future)
