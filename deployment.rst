Prometheus Server Deployment
===============================

Prometheus server is the center of the monitoring solution. It is responsible for:

- Scraping metrics from exporters and Pushgateway;
- Supporting queries through PromQL;
- Triggering alerts to Alertmanager;
- ...

Deployment
-----------

Let's deploy a Prometheus server:

1. Download the latest tarball (we use Linux) from `the download page <https://prometheus.io/download/>`_;
#. Decompress the pacakge and change into the directory;
#. The configuration file is **prometheus.yml**, there is no need to change it for now since no export has been deployed;
#. Kick started Prometheus **./prometheus**;
#. We should be able to access it through **http://<Prometheus Server IP>:9090**.

After getting a Prometheus server running, we can move forward deploying exporters and scraping metrics from them.

Resources
----------

Prometheus is easy and straightforward, but it stills need some reading work if you want to use it in a flexible way. Here are some documents you should go through when you have enough bandwidth:

- `Prometheus Getting Started <https://prometheus.io/docs/prometheus/latest/getting_started/>`_
- `Prometheus Configuration Reference <https://prometheus.io/docs/prometheus/latest/configuration/configuration/>`_
- `Querying Prometheus <https://prometheus.io/docs/prometheus/latest/querying/basics/>`_

