Tips
=====

This chapter will be used to record tips during the usage of Prometheus.

Install Prometheus as a Systemd Service
----------------------------------------

::

  sudo useradd --no-create-home --shell /bin/false prometheus

  tar -zxvf prometheus-xxxxxx.linux-amd64.tar.gz
  sudo mv prometheus-xxxxxx.linux-amd64 /opt/prometheus
  sudo mkdir /opt/prometheus/data
  sudo chown -R prometheus:prometheus /opt/prometheus

  sudo mkdir /etc/prometheus
  sudo cp /opt/prometheus/prometheus.yml /etc/prometheus
  sudo chown -R prometheus:prometheus /etc/prometheus

  sudo cat > /etc/systemd/system/prometheus.service <<-EOF
  [Unit]
  Description=Prometheus
  Wants=network-online.target
  After=network-online.target

  [Service]
  User=prometheus
  Group=prometheus
  Type=simple
  Restart=on-failure
  ExecStart=/opt/prometheus/prometheus \
  --config.file /etc/prometheus/prometheus.yml \
  --storage.tsdb.path /opt/prometheus/data \
  --storage.tsdb.retention.time 15d

  [Install]
  WantedBy=multi-user.target
  EOF

  sudo systemctl daemon-reload
  sudo systemctl start prometheus

Import Community Defined Grafana Dashboard
--------------------------------------------

Grafana is powerful allowing end users define dashboards based on their requirements. However, this does not mean we have to define our dashboards from scratch all the time (or import dashboards we ourselves previously defined). Actually, we can import dashboards defined by the community easily.

Here is an example how to import a community defined dashboard:

1. Find available dashboards published by the community `here <https://grafana.com/grafana/dashboards>`_;
#. E.g., we want to create a dashboard for the wmi_exporter:

   .. image:: images/grafana_community_dashboardsearch.png

#. Click the first dashboard which is the most popular one and we can see its description together with a **dashboard ID**:

   .. image:: images/grafana_community_dashboardid.png

#. Copy the dashboard ID, then go to our dashboard GUI: Create->Import->Paste the dashboard ID->Wait for a while->Name it and select the Prometheus data source->Import;
#. We have a working dashboard now:

   .. image:: images/grafana_community_dashboardshow.png

#. If the dashboard does not meet your requirements, modify it!
#. Done!

Add New Labels
---------------

This tip shows 2 x methods to add labels.

- The original configuration: label "node_type" is added to both targets with the same value.

  ::

    - job_name: 'node_exporter'
      static_configs:
        - targets:
            - '192.168.10.10:9100'
            - '192.168.10.11:9100'
          labels:
            node_type: 'unity_node'

- Add different labels by splitting targets:

  ::

    - job_name: 'node_exporter'
      static_configs:
        - targets:
            - '192.168.10.10:9100'
          labels:
            node_type: 'unity_node'
            node: node1
        - targets:
            - '192.168.10.11:9100'
          labels:
            node_type: 'unity_node'
            node: node2

- Add different labels by using relabel_configs:

  ::

    - job_name: 'node_exporter'
      static_configs:
        - targets:
            - '192.168.10.10:9100'
            - '192.168.10.11:9100'
          labels:
            node_type: 'unity_node'
      relabel_configs:
        - source_labels: [__address__]
          regex: '.+?\.10:9100'
          target_label: 'node'
          replacement: 'node1'
        - source_labels: [__address__]
          regex: '.+?\.11:9100'
          target_label: 'node'
          replacement: 'node2'

Select Legends to Display on Grafana Panel
--------------------------------------------

- Click the color icon "-" of a legend on a panel:

  - Select the color to be used
  - Customize the color to be used
  - Align the legend to left/right Y axis

- Click the name of a legend

  - Only this legend will be displayed on the panel
  - Click again, all legends will be displayed as before

- Shift + Click legends: select multiple legends to display on the panel
- Ctrl + Click legends : select multiple legends to not display

Graph Top N in Grafana
------------------------

PromQL **topk** will show more than expected results on Grafana panels because of `this issue <https://github.com/prometheus/prometheus/issues/586>`_.

The problem can be worked around by defining a variable containing the top N results, then filter query results with this variable in Panel. The details can be found `here <https://www.robustperception.io/graph-top-n-time-series-in-grafana>`_.

Below is a straightforward example:

1. Metrics:

   - disk_read_average

     ::

       disk_read_average{instance="192.168.10.11:9272",job="vcenter",vm_name="vm1"}
       disk_read_average{instance="192.168.10.11:9272",job="vcenter",vm_name="vm2"}
       ...
       disk_read_average{instance="192.168.10.11:9272",job="vcenter",vm_name="vm100"}

   - disk_write_average

     ::

       disk_write_average{instance="192.168.10.11:9272",job="vcenter",vm_name="vm1"}
       disk_write_average{instance="192.168.10.11:9272",job="vcenter",vm_name="vm2"}
       ...
       disk_write_average{instance="192.168.10.11:9272",job="vcenter",vm_name="vm100"}

#. Goal: show disk I/O (read + write) for the top 5 x VMs
#. Define a variable (top_vm_io) which returns the top 5 x VMs

   ::

     # Query
     query_result(topk(5, avg_over_time((disk_read_average + disk_write_average)[${__range_s}s:])))
     # Regex
     /vm_name="(.*)"/
     # Enable "Multi-value" and "Include All option"

#. Panel query

   ::

     disk_read_average{vm_name=~"$top_vm_io"} + disk_write_average{vm_name=~"$top_vm_io"}

**Notes**:

- PromQL functions avg_over_time/min_over_time/max_over_time: should be selected based on the use case;
- __range_s is a builtin variable, refer `here <https://grafana.com/docs/grafana/latest/reference/templating/#the-range-variable>`_ for details;
- [${__range_s}s:] is a subquery, refer `here <https://prometheus.io/docs/prometheus/latest/querying/examples/#subquery>`_ for details.

Use Telegraf as Exporters
--------------------------

`Telegraf <https://github.com/influxdata/telegraf>`_ is a part of `the TICK Stack <https://www.influxdata.com/blog/introduction-to-influxdatas-influxdb-and-tick-stack/>`_ monitoring solution. Telegraf supports collecting metrics from different sources through input plugins and shipping metrics to different destinations through output plugins.

Prometheus is a supported output destination, in other words, Telegraf can be used as Prometheus exporters. It supports a large num. of input plugins, including OSs, databases, clouds, etc.

Usage:

- List supported input plugins:

  ::

    telegraf --input-list

- List supported output plugins:

  ::

    telegraf --output-list

- Generate a config with vSphere input plugin and Prometheus output plugin:

  ::

    telegraf --section-filter agent:inputs:outputs --input-filter vsphere --output-filter prometheus_client config | tee telegraf.conf

- Run Telegraf:

  ::

    # After tuning the config
    telegraf --config telegraf.conf
