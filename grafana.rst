Grafana Dashboard
===================

This chapter will cover the knowledge on how to create Grafana dashboards.

How all the stuff works?
----------------------------

Before moving forward, we should understand the relationshipt of Prometheus exporters, Prometheus, and Grafana.

- Prometheus exporters are responsible for collecting performance stats from monitored devices. They work as standalone apps running close to or even on the targets. The Prometheus company or users can implement their own exporters based on their requirements. There exist quite a few exporters for use directly, refer `here <https://prometheus.io/docs/instrumenting/exporters/>`_
- Prometheus itself does not collect data from monitored devices directly, instead, it collects (called **scrape** by Prometheus) data from enabled exporters (based on configuration file prometheus.yml) and save the data locally into the **./data** directly or remotely to a database (such as InfluxDB).
- Grafana leverages Prometheus as a data source and showcases the data based on user defined dashboards.

**Notes:**

- Exporters only scrape data from monitored devices, they won't save data;
- Prometheus scraps data from exporters and save the data centraly. Advanced queries against the data are supported by Prometheus through PromQL;
- Grafana leverages PromQL and its builtin query functions to filter data and display them on dashboards;
- Grafana uses Prometheus as a kind of data source, and also support quite a lot other data sources such as ELK, InfluxDB, etc.

Reference
-----------

- `Query Prometheus <https://prometheus.io/docs/prometheus/latest/querying/basics/>`_
- `Grafana Templating Variables <https://grafana.com/docs/grafana/latest/reference/templating/>`_
- `Using Prometheus in Grafana <https://grafana.com/docs/grafana/latest/features/datasources/prometheus/>`_
