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

Undetstand Prometheus Data
----------------------------

Exporter Data
~~~~~~~~~~~~~~~

Data scraped by exporters are as below:

.. image:: images/exporter_data.png

Let's explain the data based on the metric (counter, gauge, etc., refer to `Metrics Types <https://prometheus.io/docs/concepts/metric_types/>`_) **unity_basic_reads**:

- Each metric has a name, in this example, its name is unity_basic_reads;
- An metric may have some labels associated with it to distinguish its instances. In this example, unity_basic_reads has 2 x lables: sp, unity. Based on the label values, instances can be differentiated easily - this is important for data filter;
- Metrics will be scraped based on the interval configured for exporters, but they won't be saved.

Prometheus Data
~~~~~~~~~~~~~~~~~

Data scraped by Prometheus from exporters are as below:

.. images:: images/prometheus_data.png

Let's explain the data based on the same metric **unity_basic_reads**:

- Query/Filter can be executed for all metrics supported by exporters. In this example, unity_basic_reads is a metric scraped from an exporter, hence we can query it from Prometheus directly;
- Beside the labels provided by an exporter (as above), Prometheus will add several more lables. In this example, 2 x lables are added: instance, job:

  - instance: this label is added to all exporters. It is the same as the **targets** configured for a scrape job;
  - job: this lable is added to all exporters. It is the same as the job name as defined in prometheus.yml;
  - Additional labels can be added. Refer to `static_config and relabel_config <https://prometheus.io/docs/prometheus/latest/configuration/configuration/#static_config>`_

- Advanced queries/filters can be achieved through the use of `PromQL <https://prometheus.io/docs/prometheus/latest/querying/basics/>`_.

Add Data Source
----------------

Grafana is only responsible for displaying time series data as graphs, it does not store data but retrieve data from data sources. Before using Grafana, the first step is adding at least a data source.

Grafana can use quite a lot systems as data sources, including Prometheus (our focus), Graphite, InfluxDB, etc. It is easy to add a data source: **Configuaration->Data Sources->Add Data Srouce->Prometheus->Input Inforamtion->Save & Test->Done**

Create Dashabord
-----------------

Grafana organizes graphs as dashboards. In other words, a dashboard is the container for holding graphs - hence a dashabord need to be created before adding any graph. The creation of a dashboard is straightfoward: **Create->Dashboard**

**Notes:** Remember to save changes by clicking **Save dashboard** on the up right corner. Otherwise, your customization effort will be lost.

Varaiables
~~~~~~~~~~~

Dashboards have some special settings. The most important one is **Variables**. By defining variables, we can control the behavior of graphs within a dashboard flexsibly but not hard coded.

Add Panel
~~~~~~~~~~

TBD

Repeat
~~~~~~~

TBD

Reference
-----------

- `Query Prometheus <https://prometheus.io/docs/prometheus/latest/querying/basics/>`_
- `Grafana Templating Variables <https://grafana.com/docs/grafana/latest/reference/templating/>`_
- `Using Prometheus in Grafana <https://grafana.com/docs/grafana/latest/features/datasources/prometheus/>`_
