<div align="center">
<h1>Router</h1>
<p>
      <img alt="Tessellated Logo" src="media/tessellated-logo.png" />
</p>
<small>Software by <a href="https://tessellated.io" target="_blank"> Tessellated // tessellated.io</a></small>
</div>

---

`Router`` is an abstraction around chains that allows users to look up and use chains across Tessellated's software. 

It is particularly useful when you need to abstract configurations around RPC addresses, for instance, to coordinate signing processes, delegation queries or manually configure RPCs.

## Installation

Simply run the following to update your `go.mod` file:

```
go get github.com/tessellated-io/pickaxe
```

## Usage

`Router` provides two useful abstractions:
- **Chain:** An abstraction around a blockchain, its metadata and RPC endpoints
- **Router:** A router is initialized with multiple chains and can query chain metadata and RPC endpoints by a chain name

Several errors are also exported for convenience from `errors.go`.

The most useful functionality for `Router` is to add a `replace` for a private go Module in your `go.mod`, in order to route to your own infrastructure. By design, `Router` is statically configured, but in dynamic systems, a replacement module could dynamically refresh routes and supported chains.
