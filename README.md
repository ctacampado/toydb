# toyDB
A simple implementation of a distributed in-memory key-value store using golang

## STATUS
- implemented a very basic logger, toylog

## TODO
- key-value store implementation using the following KV stores: RocksDB, memcached, redis
- dockerize
- add implementation for distributed capability via consensus (raft, tendermint)
- implement an abstraction layer to the kv store to support other kv stores
- add an SQL layer
