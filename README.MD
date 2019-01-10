### NATS & NATS STREAMING


##NATS

extremely lightweight, scalable, performant messaging system.

does not guarantee sequence, but can be used to implement "at most once"-delivery (queueing)
offers no persistence in standalone/default configuration, so missed messages are lost.

can be used for message fanout, addressing/discovery, CNC / drain for real time data, load balancing, 

###NATS-STREAMING

sets up and manages a seperate cluster of nats-streaming nodes, that keep track of message sequence, channels("topics"), consumer/group offsets.
The nats-streaming cluster is managed via the underlying nats-cluster (node discovery, leader election, quorum - all done via NATS).

nats-streaming clients, communicate by sending "nats-streaming-protocol"-messages inside of "nats-protocol"-messages on the nats-cluster
(subscribing to a topic starting with an offset or joining a durable group (and using its offset) is done by posting and reading from special "nats-streaming" topics in NATS).

### serialization

like kafka, the size of messages and the overhead of parsing them is solved by compressing messages to e binary format.
what avro is to kafka, protocolbuffers is to NATS

##protobuf

developed by google, and while lacking some features of avro (schema registry, flexible and dynamic schemas), it outperforms acvro by 2-3 times the speed when encoding and decoding, as well as message size.


