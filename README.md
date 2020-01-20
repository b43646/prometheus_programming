# prometheus_programming
Prometheus programming(metrics exporter/remote storage)

### metrics exporter

Add the following config into prometheus.yml

```
    static_configs:
    - targets: ['192.168.28.131:9145']
```
### remote write

![image](https://prometheus.io/docs/prometheus/latest/images/remote_integrations.png)

add the following config into your prometheus.yml

```
remote_write:
  - url: "http://localhost:1234/receive"
```
