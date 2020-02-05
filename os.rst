Operating System Exporters
===========================

Linux and Windows are major OSs used in data centers today. As a popular performance monitoring solution, Prometheus of course supports collecting metrics from them. This chapter will cover how to install exporters on Linux and Windows. The metrics collected by these exporters will be scraped by Prometheus server, and then leveraged by Grafana to create dashboards.

Linux node_exporter
--------------------

Since Linux is the most puplar OS used in data centres, the company of Prometheus develops an official exporter for collecting Linux metrics, which is called node_exporter.

The installation of node_exporter is as easy as decompressing a Linux tarball:

1. Download node_exporter from `here <https://prometheus.io/download/#node_exporter>`_;
#. On the target Linux, decompress it and change directory into the decompressed folder;
#. Kick started it as **./node_exporter**;
#. That is it, node_exporter is up and running. All collected metrics can be seen by accessing **http://<Linux IP>:9100/metrics**

Windows wmi_exporter
---------------------

Windows is also adopted widely in data centers. Unfortunately, the company of Prometheus does not develop an official export for Windows. Luckily enough, Prometheus has a really active community and the contributors develop an exporter for Windows, A.K.A wmi_exporter.

The installation of wmi_exporter won't take more effort than installing node_export on Linux:

1. Download the latest msi package from its `github repo release page <https://github.com/martinlindhe/wmi_exporter/releases>`_;
#. Double click the package to start the installation and follow the wizard to complete the process;
#. That is it. All collected metrics can be seen by accessing **http://<Windows IP>:9182/metrics**

Configure Prometheus
----------------------

Once an exporter (node_exporter, wmi_exporter and all others) is up and running, it starts to collect metrics from the monitored device(s). However, Prometheus has no idea about an exporter before it is enabled in the Prometheus configuration file (prometheus.yml).

To make Prometheus scrape metrics from node_exporter and wmi_exporter:

1. Modify the Promethus configuration file prometheus.yml and add below job definitions:

   ::

     # For node_exporter
     - job_name: 'node_exporter'
       static_configs:
         - targets:
             - '<Linux IP>:9100'

     # For wmi_exporter
     - job_name: 'wmi_exporter'
       static_configs:
         - targets:
             - '<Windows IP>:9182'

#. Restart Prometheus: since each Prometheus server will scrape multiple targets (exporters and pushgateway), it is not recommended to restar the whole Prometheus server process directly since it impacts all targets, instead, it is recommended to reload the configuration file only:

   ::

     # Find the Prometheus server process ID
     ps -ef | grep prometheus
     # Reload configuration file by sending SIGHUP
     kill -s SIGHUP <prometheus process ID>

#. If everything is fine, the newly added node_exporter and wmi_exporter should appear as a target under **http://<prometheus server IP>:9090/targets**;
#. You should able to see all metrics by clicking the corresponding targets.
