# Docker Compose demo

#### Requirements

* Docker

#### Demo

1. Start containers: `docker-compose up -d`
2. Check the containers state: `docker ps`
3. Access the hi endpoint: `curl 0:8080/hi?id=123456`
4. Run a few more requests against the hi endpoint with different ids: `curl 0:8080/hi?id=98765`
5. Analyse the metrics endpoint: `curl 0:8080/metrics`
6. Open Prometheus UI: http://localhost:9090/graph
7. Select the `custom_requests_hi_total` metric and display it on console and as a graph
8. Run a few more requests to the hi endpoint to see how it reflects
9. Cleanup with: `docker-compose down`

#### Extra

* Explore the Prometheus Query Language: https://prometheus.io/docs/prometheus/latest/querying/examples

