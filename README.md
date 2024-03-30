# KeyKeeper
KeyKeeper is a single-binary version of Pass.

## Why another Pass?
Pass by Jason A. Donenfeld is an excellent tool for managing passwords.
It is a simple bash script, and it uses GnuPG to encrypt and decrypt the password files and then puts these files into Git for version management.

As a result, the architecture is simple and uses real-world-tested software to achieve features many commercial products would spend a lot of effort on.

However, GnuPG and Git come with a lot of dependencies. Thus, it takes such an effort to use them across different command-line environments, such as installing them on resource-limited laptops or VPS.

This project tries to build a Go version of Pass within one binary to simplify the installation.

## Installation
KeeKeeper is currently devloped on macOS and OpenBSD. Following the release, prebulit binaries can be found at Release section.

## Build from source

1. Install Golang
2. Clone this repoistory to your desk
3. Run `go build keekeeper/cmd/kk` within repoistory
