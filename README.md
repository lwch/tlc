# tlc

tiny linux container supports:

- all in one executable
- use image of package file tar, tar.gz, tar.bz2, zip, ...
- save data when restart container
- reuse host network, no tun/tap interface or bridge network

## deployment

1. register `tlcd` service

      sudo ./tlc service install
2. start `tlcd` service

      sudo systemctl start tlcd.service

## container

1. run container

      sudo ./tlc run --dir \<path-to-rootfs>
2. attach container

      sudo ./tlc exec \<container-id> /bin/bash
4. stop container

      sudo ./tlc stop \<container-id>
5. start container

      sudo ./tlc start \<container-id>
6. restart container

      sudo ./tlc restart \<container-id>
7. remove container

      sudo ./tlc rm [-f] \<container-id>