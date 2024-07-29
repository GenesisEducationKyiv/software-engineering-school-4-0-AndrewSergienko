# Currency rate service

This is a simple currency rate service that provides the exchange rate between two currencies.


## Observability
There are many information criteria that we can use to monitor the quality of a system.
They are of two types: business criteria and technical criteria.
The first ones show information for the business that it needs to create/change its next development steps.
The latter show information about the program's operation, its workload, speed, capacity, etc.

Below is a list of metrics that can be collected in the project:

### Business criteria
- **Number of users**: The number of users that use the service.
- **Number of unique requests**: The number of unique requests that the service receives.
- **Customer retention rate**: The percentage of returning users.
- **Customer satisfaction score**: A measure of user satisfaction, often collected through surveys.
- **Churn rate**: The percentage of users who stop using the service over a given period.

### Technical criteria
- **Requests per second (RPS)**: The number of requests that the service receives per second.
- **Success rate (general and per endpoint)**: The percentage of successful requests.
- **Mean latency (general and per endpoint)**: The average time taken for a request to be processed.

- **Success rate per currency rate source**: The percentage of successful requests for each currency rate source.

- **CPU usage**: The percentage of CPU resources used by the service.
- **Memory usage**: The amount of memory used by the service.
- **Disk I/O**: The rate of read and write operations on the disk.
- **Network I/O**: The rate of data transmission over the network.

### Alerts
- **High error rate**: If the error rate exceeds a certain threshold, an alert is triggered.
- **High latency**: If the latency exceeds a certain threshold, an alert is triggered.
- **High CPU usage**: If the CPU usage exceeds a certain threshold, an alert is triggered.
- **High memory usage**: If the memory usage exceeds a certain threshold, an alert is triggered.