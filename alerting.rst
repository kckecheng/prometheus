Alerting
=========

Alerting is an important part of monitoring. For Prometheus, alerting is separated into two parts:

- Alerting rules:

  * Defined on Prometheus servers;
  * Get triggered when rule expressions are met;
  * Sent to Alertmanager once triggered.

- Alertmanager:

  * Manage alerts, such as silencing, inhibition, aggregation;
  * Notifications.

Alertmanager Deployment
------------------------

Alertmanager can be downloaded from Prometheus `official download page <https://prometheus.io/download/#alertmanager>`_. The deployment process is as easy as:

1. Decompress the tarball;
#. Start the service: ./alertmanager
#. The service should be accessible from http://<FQDN or IP>:9093
#. To enable notification (send emails in our example), change the configuration as below (refer to `the official example <https://github.com/prometheus/alertmanager/blob/master/doc/examples/simple.yml>`_):

   ::

     # alertmanager.yml
     global:
       resolve_timeout: 5m
       smtp_smarthost: '<smtp server>:<smtp server port>'
       smtp_from: '<the default sender email>'

     route:
       group_by: ['alertname']
       group_wait: 10s
       group_interval: 10s
       repeat_interval: 1h
       receiver: monitor-admin

     receivers:
     - name: 'monitor-admin'
       email_configs:
       - to: '<recevier email>'
         tls_config:
           insecure_skip_verify: true

#. Restart Alertmanager
#. Done

For more inforamtion, refer to `Alertmanger introduction <https://prometheus.io/docs/alerting/alertmanager/>`_ for details.

Alerting Rules
---------------

Alerting rules are configured on Prometheus servers:

1. Tune Prometheus configurations:

   ::

     global:
       evaluation_interval: 30s # How frequently to evaluate rules, 1m as default

     # Alertmanager configuration
     alerting:
       alertmanagers:
       - static_configs:
         - targets:
           - alertmanager:9093

     # Load rules once and periodically evaluate them according to the global 'evaluation_interval'.
     rule_files:
       - "rules/sample.yml" # Define rules in files

#. Create rule definition files (rules/samples in this example):

   ::

     groups:
     - name: UnityCapacity
       rules:
       - alert: HighCapacity
         expr: physical_used{job="unity-capacity"} / physical_total{job="unity-capacity"} > 0.8
         for: 1m
         labels:
           severity: p1
         annotations:
           summary: High Capacity Usage
           description: "High Capacity Usage {{ $labels.instance }} {{ $value }}"

#. Restart Prometheus or "kill -s SIGHUP `pgrep prometheus`"
#. Done

**Notes**: notification supports templating, A.K.A embeding variables into static paragraph. It is an important feature should be understood. Please refer to `examples <https://prometheus.io/docs/alerting/notification_examples/>`_ for the usage.


