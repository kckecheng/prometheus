Exporter Deployment
====================

Based on our previous introduction to Prometheus, we know that exporters are the actual workers who collect metrics from data souces like OSs, network devices, etc.

In this chapter, we will go through how to deploy node_exporter for Linux, how to deploy wmi_exporter for Windows, and how to deploy snmp_exporter for network equipment (switches, routers, firewalls, etc.).

Actually, all exporters follow a similar deployment method and invovle similar Prometheus configuration changes. After reading this chapter, you should have an idea on how to deploy and enable any exporter.

.. toctree::
   :maxdepth: 2

   os
   switch
