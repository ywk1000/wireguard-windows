### WireGuard for Windows Attack Surface

_This is an evolving document, describing currently known attack surface, a few mitigations, and several open questions. This is a work in progress. We document our current understanding with the intent of improving both our understanding and our security posture over time._

WireGuard for Windows consists of four components: a kernel driver, and three separate interacting userspace parts.

#### Wintun

Wintun is a kernel driver. It exposes:

  - A miniport driver to the ndis stack, meaning any process on the system that can access the network stack in a reasonable way can send and receive packets, hitting those related ndis handlers.
  - There are also various ndis OID calls, accessible to certain users, which hit further code.
  - A virtual file in `\\Device\WINTUN%d`, whose permissions are set to `SDDL_DEVOBJ_SYS_ALL`. Presumably this means only the "Local System" user can open the file and do things, but it might be worth double checking that. It sends and receives layer 3 packets, and does minimal parsing of the IP header in order to determine packet family. It also does more complex struct alignment pointer arithmetic, as it can send and receive several packets at a time in a single bundle.

### Tunnel Service

The tunnel service is a userspace service running as Local System, responsible for creating UDP sockets, creating Wintun adapters, and speaking the WireGuard protocol between the two. It exposes:

  - A listening pipe in `\\.\pipe\WireGuard\%s`, where `%s` is some basename of an already valid filename. Its permissions are set to `O:SYD:(A;;GA;;;SY)`, which presumably means only the "Local System" user can access it and do things, but it might be worth double checking that. This pipe gives access to private keys and allows for reconfiguration of the interface, as well as rebinding to different ports (below 1024, even).
  - It handles data from its two UDP sockets, accessible to the public Internet.
  - It handles data from Wintun, accessible to all users who can do anything with the network stack.
  - It does not yet drop privileges.

### Manager Service

The manager service is a userspace service running as Local System, responsible for starting and stopping tunnel services, and ensuring a UI program with certain handles is available to Administrators. It exposes:

  - Extensive IPC using unnamed pipes, inherited by the unprivileged UI process.
  - A writable `CreateFileMapping` handle to a binary ringlog shared by all services, inherited by the unprivileged UI process. It's unclear if this brings with it surprising hidden attack surface in the mm system.
  - It listens for service changes in tunnel services according to the string prefix "WireGuardTunnel$".
  - It manages DPAPI-encrypted configuration files in Local System's local appdata directory, and makes some effort to enforce good configuration filenames.
  - It uses `wtsEnumerateSessions` and `WTSSESSION_NOTIFICATION` to walk through each available session. It then uses `wtfQueryUserToken`, and then calls `GetTokenInformation(TokenGroups)` on it. If one of the returned group's SIDs matches `CreateWellKnownSid(WinBuiltinAdministratorsSid)`, then it spawns the unprivileged UI process as that user token, passing it three unnamed pipe handles for IPC and the log mapping handle, as descried above.

### UI

The UI is an unprivileged process running as the ordinary user for each user who is in the Administrators group (per the above). It exposes:

  - There currently are no anti-debugging mechanisms, which means any process that can, debug, introspect, inject, send messages, or interact with the UI, may be able to take control of the above handles passed to it by the manager service. In theory that shouldn't be too horrible for the security model, but rather than risk it, we might consider importing all of the insane things VirtualBox does in this domain. Alternatively, since the anti-debugging stuff is pretty ugly and probably doesn't even work properly everywhere, we may be better off with starting the process at some heightened security level, such that only other processes at that level can introspect; however, we'll need to take special care that doing so doesn't give the UI process any special accesses it wouldn't otherwise have.
  - It renders highlighted config files to a msftedit.dll control, which typically is capable of all sorts of OLE and RTF nastiness that we make some attempt to avoid.

