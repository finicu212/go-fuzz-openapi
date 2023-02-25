### Proxy-API tool

This is a tool made in Node, ExpressJS that uses http-proxy-middleware module to forward requests with a certain delay. The delay is specified using the `delay` query parameter.

In the context of our project, we will host multiple such Proxy API instances (using Docker for rapid deployment), at different ports (ideally IPs too), and try to synchronize many requests at the same time interval.
