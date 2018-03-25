### Table of Contents
1. [About](#About)
2. [Getting Started](#GettingStarted)
    1. [Installation](#Installation)
        1. [Windows](#WindowsInstallation)
        2. [Linux/BSD/MacOSX/POSIX](#PosixInstallation)
    2. [Configuration](#Configuration)
    3. [Controlling and Querying hxd via hxctl](#HxctlConfig)
    4. [Mining](#Mining)
3. [Help](#Help)
    1. [Startup](#Startup)
        1. [Using bootstrap.dat](#BootstrapDat)
    2. [Network Configuration](#NetworkConfig)
    3. [Wallet](#Wallet)
4. [Contact](#Contact)
    1. [IRC](#ContactIRC)
    2. [Mailing Lists](#MailingLists)
5. [Developer Resources](#DeveloperResources)
    1. [Code Contribution Guidelines](#ContributionGuidelines)
    2. [JSON-RPC Reference](#JSONRPCReference)
    3. [The Hx-related Go Packages](#GoPackages)

<a name="About" />

### 1. About
hxd is a full node hx implementation written in [Go](http://golang.org),
licensed under the [copyfree](http://www.copyfree.org) ISC License.

This project is currently under active development and is in a Beta state. It is
extremely stable and has been in production use since February 2016.

It also properly relays newly mined blocks, maintains a transaction pool, and
relays individual transactions that have not yet made it into a block. It
ensures all individual transactions admitted to the pool follow the rules
required into the block chain and also includes the vast majority of the more
strict checks which filter transactions based on miner requirements ("standard"
transactions).

<a name="GettingStarted" />

### 2. Getting Started

<a name="Installation" />

**2.1 Installation**<br />

The first step is to install hxd.  See one of the following sections for
details on how to install on the supported operating systems.

<a name="WindowsInstallation" />

**2.1.1 Windows Installation**<br />

* Install the MSI available at: https://github.com/hybridnetwork/hxd/releases
* Launch hxd from the Start Menu

<a name="PosixInstallation" />

**2.1.2 Linux/BSD/MacOSX/POSIX Installation**<br />

* Install Go according to the installation instructions here: http://golang.org/doc/install
* Run the following command to ensure your Go version is at least version 1.2: `$ go version`
* Run the following command to obtain hxd, its dependencies, and install it: `$ go get github.com/hybridnetwork/hxd/...`<br />
  * To upgrade, run the following command: `$ go get -u github.com/hybridnetwork/hxd/...`
* Run hxd: `$ hxd`

<a name="Configuration" />

**2.2 Configuration**<br />

hxd has a number of [configuration](http://godoc.org/github.com/hybridnetwork/hxd)
options, which can be viewed by running: `$ hxd --help`.

<a name="HxctlConfig" />

**2.3 Controlling and Querying hxd via hxctl**<br />

hxctl is a command line utility that can be used to both control and query hxd
via [RPC](http://www.wikipedia.org/wiki/Remote_procedure_call).  hxd does
**not** enable its RPC server by default;  You must configure at minimum both an
RPC username and password or both an RPC limited username and password:

* hxd.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
* hxctl.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
```
OR
```
[Application Options]
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
For a list of available options, run: `$ hxctl --help`

<a name="Mining" />

**2.4 Mining**<br />
hxd supports both the `getwork` and `getblocktemplate` RPCs although the
`getwork` RPC is deprecated and will likely be removed in a future release.
The limited user cannot access these RPCs.<br />

**1. Add the payment addresses with the `miningaddr` option.**<br />

```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
miningaddr=12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX
miningaddr=1M83ju3EChKYyysmM2FXtLNftbacagd8FR
```

**2. Add hxd's RPC TLS certificate to system Certificate Authority list.**<br />

`cgminer` uses [curl](http://curl.haxx.se/) to fetch data from the RPC server.
Since curl validates the certificate by default, we must install the `hxd` RPC
certificate into the default system Certificate Authority list.

**Ubuntu**<br />

1. Copy rpc.cert to /usr/share/ca-certificates: `# cp /home/user/.hxd/rpc.cert /usr/share/ca-certificates/hxd.crt`<br />
2. Add hxd.crt to /etc/ca-certificates.conf: `# echo hxd.crt >> /etc/ca-certificates.conf`<br />
3. Update the CA certificate list: `# update-ca-certificates`<br />

**3. Set your mining software url to use https.**<br />

`$ cgminer -o https://127.0.0.1:14009 -u rpcuser -p rpcpassword`

<a name="Help" />

### 3. Help

<a name="Startup" />

**3.1 Startup**<br />

Typically hxd will run and start downloading the block chain with no extra
configuration necessary, however, there is an optional method to use a
`bootstrap.dat` file that may speed up the initial block chain download process.

<a name="BootstrapDat" />

**3.1.1 bootstrap.dat**<br />
* [Using bootstrap.dat](https://github.com/hybridnetwork/hxd/tree/master/docs/using_bootstrap_dat.md)

<a name="NetworkConfig" />

**3.1.2 Network Configuration**<br />
* [What Ports Are Used by Default?](https://github.com/hybridnetwork/hxd/tree/master/docs/default_ports.md)
* [How To Listen on Specific Interfaces](https://github.com/hybridnetwork/hxd/tree/master/docs/configure_peer_server_listen_interfaces.md)
* [How To Configure RPC Server to Listen on Specific Interfaces](https://github.com/hybridnetwork/hxd/tree/master/docs/configure_rpc_server_listen_interfaces.md)
* [Configuring hxd with Tor](https://github.com/hybridnetwork/hxd/tree/master/docs/configuring_tor.md)

<a name="Wallet" />

**3.1 Wallet**<br />

hxd was intentionally developed without an integrated wallet for security
reasons.  Please see [hxwallet](https://github.com/hybridnetwork/hxwallet) for more
information.

<a name="DeveloperResources" />

### 4. Developer Resources

<a name="ContributionGuidelines" />

* [Code Contribution Guidelines](https://github.com/hybridnetwork/hxd/tree/master/docs/code_contribution_guidelines.md)
<a name="JSONRPCReference" />

* [JSON-RPC Reference](https://github.com/hybridnetwork/hxd/tree/master/docs/json_rpc_api.md)
    * [RPC Examples](https://github.com/hybridnetwork/hxd/tree/master/docs/json_rpc_api.md#ExampleCode)
<a name="GoPackages" />

* The Hx-related Go Packages:
    * [hxrpcclient](https://github.com/hybridnetwork/hxrpcclient) - Implements a
	  robust and easy to use Websocket-enabled Hx JSON-RPC client
    * [hxjson](https://github.com/hybridnetwork/hxjson) - Provides an extensive API
	  for the underlying JSON-RPC command and return values
    * [wire](https://github.com/hybridnetwork/hxd/tree/master/wire) - Implements the
	  Hx wire protocol
    * [peer](https://github.com/hybridnetwork/hxd/tree/master/peer) -
	  Provides a common base for creating and managing Hx network peers.
    * [blockchain](https://github.com/hybridnetwork/hxd/tree/master/blockchain) -
	  Implements Hx block handling and chain selection rules
    * [txscript](https://github.com/hybridnetwork/hxd/tree/master/txscript) -
	  Implements the Hx transaction scripting language
    * [dcrec](https://github.com/hybridnetwork/hxd/tree/master/dcrec) - Implements
	  support for the elliptic curve cryptographic functions needed for the
	  Hx scripts
    * [database](https://github.com/hybridnetwork/hxd/tree/master/database) -
	  Provides a database interface for the Hx block chain
    * [hxutil](https://github.com/hybridnetwork/hxd/hxutil) - Provides Hx-specific
	  convenience functions and types
    * [chainhash](https://github.com/hybridnetwork/hxd/tree/master/chaincfg/chainhash) -
	  Provides a generic hash type and associated functions that allows the
	  specific hash algorithm to be abstracted.
    * [connmgr](https://github.com/hybridnetwork/hxd/tree/master/connmgr) -
      Package connmgr implements a generic Hx network connection manager.
