While hxd is highly configurable when it comes to the network configuration,
the following is intended to be a quick reference for the default ports used so
port forwarding can be configured as required.

hxd provides a `--upnp` flag which can be used to automatically map the Hx
peer-to-peer listening port if your router supports UPnP.  If your router does
not support UPnP, or you don't wish to use it, please note that only the Hx
peer-to-peer port should be forwarded unless you specifically want to allow RPC
access to your hxd from external sources such as in more advanced network
configurations.

|Name|Port|
|----|----|
|Default Hx peer-to-peer port|TCP 14008|
|Default RPC port|TCP 14009|
