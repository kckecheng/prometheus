.. contents:: Switch Performance Monitoring with snmp_exporter

Switch Performance Monitoring
==============================

This chapter will cover all ideas on how to monitor switch performance with the official snmp_exporter from Prometheus.

SNMP Introduction
-------------------

SNMP, the acronym for **Simple Network Mangement Protocl**,  is an application–layer protocol defined by the Internet Architecture Board (IAB) in RFC1157 for exchanging management information between network devices. It is a part of TCP/IP protocol suite.

Its main usage focus on collecting and organizing information about managed devices on IP networks and for modifying that information to change device behavior. Devices that typically support SNMP include cable modems, routers, switches, servers, workstations, printers, and more.

**Notes:** This document only shares the most basic knowledge on SNMP, for detailed in-depth information, please leverage google or read books.

Main Components
~~~~~~~~~~~~~~~~

SNMP consist of:

- SNMP Manager
- SNMP Agent
- MIB

SNMP Manager
+++++++++++++

A SNMP manager or management system is responsible to communicate with the SNMP agent implemented network devices. This is typically a computer that is used to run one or more network management systems (NMS). The key functions are as below:

- Query agents
- Set variables in agents
- Get notifications from agents

SNMP Agent
+++++++++++

A SNMP agent is a program that is packaged within the managed devices. Enabling agents allow agents collect the management information databases from the managed devices locally and make them available to the SNMP manager during query. These agents could be standard (e.g. Net-SNMP) or specific to a vendor (e.g. HP insight agent).

The key functions of a SNMP agent is:

- Collect management information locally on managed devices
- Stores and retrieves management information as defined in the MIB
- Signals an event to the manager

MIB
++++

MIB is short for Management Information Base, which describes the managed device parameters, such as port status, througput, etc. The SNMP manager uses this database to request the agent for specific information and further translates the information as needed for the Network Management System (NMS).

Typically these MIB contains standard set of statistical and control values defined for hardware nodes on a network. SNMP also allows the extension of these standard values with values specific to a particular agent through the use of **private MIBs**.

MIB is organized as a well defined tree structure with each leaf stands for a specific object of the managed device, which is referred as **Object Identifier (OID)**. Each OID is unique and can be located from the root of the MIB tree with an address like .1.1.2.3.x.x.x. For common frequently used OIDs, human friendly names will be assigned/mapped to the number dot address, e.g., "sysDescr" is the same as ".1.3.6.1.2.1.1.1", which is defined by RFC1213 .

.. image:: images/mib_tree.png

SNMP Versions
~~~~~~~~~~~~~~~

There are 3 versions of SNMP protocol:

- SNMPv1
- SNMPv2 (SNMPv2c, SNMPv2u)
- SNMPv3

The main differences are all about security, SNMPv1 is not secure enough, SNMPv3 is too strict, hence SNMPv2 are the most popular adopted deployment.

Community String
+++++++++++++++++

A SNMP manager needs to talk to a SNMP agent to work, so a mechanism to protect the connection is required. Community based or user based authentication can be used for the purpose.

Community string is the most straightfoward method for authentication if SNMPv2c is used. Its implementation is quite simple: a string is defined as a kind of password on SNMP agents, and SNMP manger queries agents by providing the correct string for authentication.

SNMP can be used to query agents, and also can be leveraged to set variables to change something (e.g., online/offline a port). The community string provides READ and WRITE capability accordingly:

- READ ONLY: also referred as "public community string", and the default value is "public" for most managed devices once SNMP agents are enabled. It can only be used to query MIB inforamtion;
- READ WRITE/WRITE: also referred as "private community string", and this is not enabled/set by default. It can be used to change object status, such as reboot, port online/offline, etc.

Poll and Trap
~~~~~~~~~~~~~~

SNMP supports 2 ways to get infromation from MIB:

- Poll: Poll is triggered from SNMP managers, which send queries to SNMP agents on managed devices, which listen at UDP port 161. Each poll is a synchronous opeartion, BTW.
- Trap: Instead of performing queries from SNMP managers, trap is a mechanism to let SNMP agents send asynchronous events to SNMP mangers directly. With this scenario, SNMP managers listen at UDP port 162 for agent connections, and may take actions following the events (ack, etc.).

Poll Commands
~~~~~~~~~~~~~~~

SNMP ships very simple commands to support queries to MIB. The most frequently used commands are as below:

- GET: retrieve information on one specified OID
- GET NEXT: retrieve information on the next OID
- GET BULK: retrive inforamtion for a group of OIDs which share similar features
- WALK: actully WALK is not a SNMP command, but just a wrapper of GET NEXT. It is used to get information from a tree of OIDs.

MIB Browser
~~~~~~~~~~~~~

Beside network management system (SNMP Manager), a lightweight tool called **MIB Browser** can be leveraged to explore SNMP MIB inforamtion. Below is an overview of a GUI based MIB browser from iReasoning (free to use).

.. image:: images/mib_browser_overview.png
