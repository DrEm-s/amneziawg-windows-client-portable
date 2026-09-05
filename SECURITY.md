# Security

## Portable configuration storage

This portable build intentionally stores tunnel configurations as unencrypted `.conf` files so the directory can be moved between Windows installations.

A tunnel configuration contains private key material. Anyone who can read a configuration file may be able to use those credentials.

Recommendations:

- keep the portable directory on trusted storage;
- do not place it in a public/shared cloud folder;
- do not commit `Data/` or `.conf` files to Git;
- securely remove old copies when rotating tunnel keys;
- verify release hashes before use.

For issues in the portable changes, report them in this repository. For upstream AmneziaWG issues, report them to the upstream project.
