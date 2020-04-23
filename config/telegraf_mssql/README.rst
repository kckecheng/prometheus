About
======

Collect metrics for Microsfot SQL Servers with Telegraf and expose them to Prometheus.

Init a Configuration Skeleton
-------------------------------

**Notes**: If the configuration skeleton is generated on Windows, it can not be used directly since it is not encoded as UTF-8. To covert it:

- Open the configuration file with notepad++;
- Encoding -> UTF-8;
- Save.

::

  # The telegraf application can be run on both Linux and Windows
  # 'C:\Program Files\Telegraf\telegraf.exe' --section-filter agent:inputs:outputs --input-filter sqlserver --output-filter prometheus_client config > telegraf.conf
  ./usr/bin/telegraf --section-filter agent:inputs:outputs --input-filter sqlserver --output-filter prometheus_client config > telegraf.conf

Configure
----------

**SQL Server**

An login user needs to be created for the SQL Servers. Refer to `SQL Server Input Plugin <https://github.com/influxdata/telegraf/tree/master/plugins/inputs/sqlserver>`_ for steps on how to create a login.

**Telegraf**

Refer to the sample named telegraf.conf under the same directory of this README.

Usage
-----

::

  # 'C:\Program Files\Telegraf\telegraf.exe' --config 'C:\Program Files\Telegraf\telegraf.conf'
  ./usr/bin/telegraf --config telegraf.conf
