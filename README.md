# About:

A sample Go program to connect two machines using devp2p.
A p2p protocol is defined which sends and receives a message between connected peers.

Program supports 2 optional arguments:

--port: This is used for changing default 30301 port to some other port, where the server listens.

--connect-to: This is used for connecting to a server for communicating. If this is not given, server runs forever and keeps accepting connection requests otherwise server stops after sending and receiving message from connected peer.

**How to build:**

* go build

**How to start peer-1:**

* ./devp2p-sample

* Get printed enode url, to be used for connecting with second peer. Replace '[::]' after @ with ip of the machine, which is reachable from second peer.

**How to start peer-2:**

* ./devp2p-sample --connect-to [copied-enode-url-with-ip]


This can be run on same or different machine. For running on same machine, both the servers needs to run on different ports.
