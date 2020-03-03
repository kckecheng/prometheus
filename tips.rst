Tips
=====

This chapter will be used to record tips during the usage of Prometheus.

Import Community Defined Grafana Dashboard
--------------------------------------------

Grafana is powerful allowing end users define dashboards based on their requirements. However, this does not mean we have to define our dashboards from scratch all the time (or import dashboards we ourselves previously defined). Actually, we can import dashboards defined by the community easily.

Here is an example how to import a community defined dashboard:

1. Find available dashabords published by the community `here <https://grafana.com/grafana/dashboards>`_;
#. E.g., we want to create a dashboard for the wmi_exporter:

   .. image:: images/grafana_community_dashboardsearch.png

#. Click the first dashboard which is the most popular one and we can see its description together with a **dashboard ID**:

   .. image:: images/grafana_community_dashboardid.png

#. Copy the dashboard ID, then go to our dashboard GUI: Create->Import->Paste the dashboard ID->Wait for a while->Name it and select the Prometheus data souce->Import;
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

- Add differnt labels by spliting targets:

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
