Exporter Deployment
====================

Based on our previous introduction on Prometheus, we know that exporters are the actual workers collect metrics from data souces like OSs, network devices, etc.

In this chapter, we will go through how to deploy node_exporter for Linux, how to deploy wmi_exporter for Windows, and how to deploy snmp_exporter for network equipment (switch, routers, firewall, etc.).

Actually, all exporters follow the similar deployment method and invovle similar Prometheus configuration changes. After reading this chapter, you should have a common idea on how to deploy and enable any exporter.

.. toctree::
   :maxdepth: 2

   os
   switch
