About
======

Telegraf can be used to collect metrics from vSphere and expose them to Prometheus.


Init a Configuration Skeleton
-------------------------------

::

  ./usr/bin/telegraf --section-filter agent:inputs:outputs --input-filter vsphere --output-filter prometheus_client config > telegraf.conf

Configure
----------

Refer to the sample named telegraf.conf under the same directory of this README.

Usage
-----

::

  ./usr/bin/telegraf --config telegraf.conf

Reference
----------

- `Telegraf vSphere Input Plugin <https://github.com/influxdata/telegraf/tree/master/plugins/inputs/vsphere>`_
- `vSphere Performance Counter Reference <https://www.vmware.com/support/developer/converter-sdk/conv60_apireference/vim.PerformanceManager.html>`_
- `How to convert vSphere summation counters <https://kb.vmware.com/s/article/2002181>`_
