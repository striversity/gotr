# Go Resource Pool

This is a simple example of how to implement a resoruce pool in Go.

Since Go is a gargabe collected language, you can stress the GC if you allocating a large number of objects, which are short lived.

Imagine you application allocated a ClientReq object for each request from a client. For each request, a goroutine is spun up to handle the request. If there are thousands of connections per minute for example, each object may only be around for a few milliseconds to seconds, before they are no longer needed.

In this scenario, the GC will be called to create thoustands of objects per seconds, and then later have to free thousands of objects frequently.