# pwhoiscache
A simple whois server that caches results from the original pwhois servers

# Installation
`go install github.com/bassem-essam/pwhoiscache@latest`

# Usage
1. Open a tmux session or any technique to run a command in the background
2. Run pwhoiscache

# Background
For information about pwhois (Prefix WhoIs) Visit: [https://pwhois.org/](https://pwhois.org/)

# Why
Querying information about an IP address is easy using any whois client, like the linux command line `whois`.

When querying information about a list of ip addresses it takes a long time to communicate with the whois server for all these queries.

Many of these queries are unnecessary because the list probably has many ips that share the same prefix. So caching a prefix along with information revolving around it (like organization, asn number ... etc) is sufficient to know information regarding an IP lying under the same prefix.
