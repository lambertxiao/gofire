# Gofire

Gofire is a network framework that supports user-defined message formats and message transmission protocols. Currently, the ChannelGenerator of tcp and udp is built-in. Users can easily reuse the capabilities of the two to achieve message transmission. Gofire does not limit the underlying protocol of message transmission. In theory, users can transmit messages through any network transmission protocol, and only need to implement the IChannel interface of the program.
