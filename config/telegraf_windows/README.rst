About
======

Collect metrics for Windows with Telegraf and expose them to Prometheus.

Init a Configuration Skeleton
-------------------------------

**Notes**: the configuration skeleton generated can not be used directly since it is not encoded as UTF-8. To covert it:

- Open the configuration file with notepad++;
- Encoding -> UTF-8;
- Save.

::

  'C:\Program Files\Telegraf\telegraf.exe' --section-filter agent:inputs:outputs --input-filter win_perf_counters --output-filter prometheus_client config > telegraf.conf

Configure
----------

Refer to the sample named telegraf.conf under the same directory of this README.

Usage
-----

::

  'C:\Program Files\Telegraf\telegraf.exe' --config 'C:\Program Files\Telegraf\telegraf.conf'

Reference
-----------

- `Running Telegraf as a Windows Service <https://github.com/influxdata/telegraf/blob/master/docs/WINDOWS_SERVICE.md>`_
- `Monitoring Windows Services with Grafana, InfluxDB and Telegraf <https://www.influxdata.com/blog/monitoring-windows-services-with-grafana-influxdb-and-telegraf/>`_
