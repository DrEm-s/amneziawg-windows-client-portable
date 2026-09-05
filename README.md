# AmneziaWG Windows Portable

Unofficial portable build of **AmneziaWG for Windows**.

Based on the official [`amnezia-vpn/amneziawg-windows-client`](https://github.com/amnezia-vpn/amneziawg-windows-client).

## Download

Latest portable release:

- [Windows x64 (amd64)](https://github.com/DrEm-s/amneziawg-windows-client-portable/releases/latest/download/amneziawg-windows-portable-amd64.zip)
- [Windows x86](https://github.com/DrEm-s/amneziawg-windows-client-portable/releases/latest/download/amneziawg-windows-portable-x86.zip)
- [Windows ARM64](https://github.com/DrEm-s/amneziawg-windows-client-portable/releases/latest/download/amneziawg-windows-portable-arm64.zip)
- [SHA-256 checksums](https://github.com/DrEm-s/amneziawg-windows-client-portable/releases/latest/download/SHA256SUMS.txt)

## Portable mode

No installer is required. Extract the archive to a writable directory and run `amneziawg.exe`.

Persistent application data is stored next to the executable:

```text
AmneziaWG-Portable\
├── amneziawg.exe
├── awg.exe
├── wintun.dll
└── Data\
    ├── Configurations\
    │   └── <tunnel>.conf
    └── log.bin
```

### Security warning

**Tunnel configuration files are stored as unencrypted `.conf` files.**

They contain private keys and must be protected like passwords. Do not keep the portable directory on an untrusted/shared computer or storage device.

The portable build deliberately does not use Windows DPAPI because DPAPI binds encrypted configurations to a particular Windows account/machine and prevents the directory from being moved between systems.

## Differences from upstream

| Official client | Portable build |
| --- | --- |
| `%ProgramFiles%\AmneziaWG\Data` | `Data\` next to `amneziawg.exe` |
| Saved configs use DPAPI | Saved configs are plain `.conf` |
| `.conf` migration to DPAPI is enabled | DPAPI migration is disabled |
| Installer-oriented distribution | Ready-to-extract ZIP archives |

Windows administrator privileges can still be required for installing/running the AmneziaWG manager and tunnel services. Portable mode does not bypass Windows security requirements.

## Source layout

This repository contains the Windows client. The modified AmneziaWG Windows library is pinned as a Git submodule:

- client: [`DrEm-s/amneziawg-windows-client-portable`](https://github.com/DrEm-s/amneziawg-windows-client-portable)
- portable library: [`DrEm-s/amneziawg-windows-portable`](https://github.com/DrEm-s/amneziawg-windows-portable)

Clone recursively:

```powershell
git clone --recursive https://github.com/DrEm-s/amneziawg-windows-client-portable.git
cd amneziawg-windows-client-portable
.\build.bat
```

## Versioning

Portable releases follow the upstream client version:

```text
v3.1.1-portable.1
v3.1.1-portable.2
v3.1.2-portable.1
```

The final number is the revision of the portable patch for that upstream version.

## Upstream

- [AmneziaWG Windows client](https://github.com/amnezia-vpn/amneziawg-windows-client)
- [AmneziaWG Windows library](https://github.com/amnezia-vpn/amneziawg-windows)
- [Amnezia project](https://github.com/amnezia-vpn)

This repository is an **unofficial community build** and is not an official AmneziaVPN release.

## License

The upstream project is MIT-licensed. Original copyright and license notices are retained.
