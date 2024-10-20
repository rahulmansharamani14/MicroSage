# MicroSage
An AI-powered system for detecting anomalies in microservice metrics to proactively identify and resolve potential issues.


## **Table of Contents**

- [Overview](#overview)
- [Problem Statement](#problem-statement)
- [Features](#features)




## **Overview**

MicroSage is an ongoing project focused on integrating AI-driven anomaly detection into microservices monitoring. By collecting and analyzing metrics from a Go-based microservice, we aim to proactively identify unusual patterns using machine learning techniques. This approach helps prevent potential system failures and ensures a resilient microservices infrastructure.


## **Problem Statement**
Microservices architectures are complex, making it challenging to monitor the health and performance of each service effectively. Traditional monitoring tools often react to issues after they've occurred, leading to downtime or a degraded user experience. There's a need for a proactive solution that can detect anomalies in real-time, allowing teams to address potential problems before they impact users.

## **Features**

- **Microservice Monitoring:** A microservice (in Go) instrumented for metrics collection.
- **Metrics Collection:** Utilize Prometheus to collect and store metrics data from the microservice.
- **Visualization:** Set up Grafana dashboards for real-time monitoring of the microservice metrics.
- **Anomaly Detection:** Build a Python service that uses machine learning to detect anomalies in the collected metrics.
- **Alerting Mechanism:** Aim to implement notifications (e.g., email or Slack alerts) when anomalies are detected.


## Note
This project is a work in progress. Features and plans may evolve as we progress and learn more about the challenges involved.
