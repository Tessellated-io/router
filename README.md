<div align="center">
<p style="font-size:24px; font-weight: bold;">Router</p>
<p>
      <img alt="Tessellated Logo" src="media/tessellated-logo.png" />
<small>Software by <a href="https://tessellated.io" target="_blank"> Tessellated // tessellated.io</a></small>
</p>
</div>

---

## Introduction

`Router` is an abstraction around chains that allows users to look up and use chains across Tessellated's software. 

It is particularly useful when you need to abstract configurations around RPC addresses, for instance, to coordinate signing processes, delegation queries or manually configure RPCs.

## Installation

Simply run the following to update your `go.mod` file:

```
go get github.com/tessellated-io/router
```

## Usage

`Router` provides two useful abstractions:
- **Chain:** An abstraction around a blockchain, its metadata and RPC endpoints
- **Router:** A router is initialized with multiple chains and can query chain metadata and RPC endpoints by a chain name

Several errors are also exported for convenience from `errors.go`.

We provide two default routing strategies out of the box: 
- `static`: Allows configuration of a router with preconfigured chains
- `file`: Allows configuration of a router with routes in a file.
  A file router assumes a config such as:
  
  ```yaml
  - chain-id: my-chain
    grpc: tcp://1.2.3.4:9090
  ```

The most useful functionality for `Router` is to add a `replace` for a private go Module in your `go.mod`, in order to route to your own infrastructure. 
