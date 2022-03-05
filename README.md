# relax-dd

Relax you and run dd.

## Usage

```console
$ relax-dd  ~/Downloads/archlinux-2022.03.01-x86_64.iso /dev/sda
NAME   FSTYPE  FSVER            LABEL       UUID                                 FSAVAIL FSUSE% MOUNTPOINTS
sda    iso9660 Joliet Extension ARCH_202203 2022-03-01-15-50-40-00
|-sda1 iso9660 Joliet Extension ARCH_202203 2022-03-01-15-50-40-00
`-sda2 vfat    FAT16            ARCHISO_EFI 1308-C0B0

execute the following command:
  sudo dd if=/home/you/Downloads/archlinux-2022.03.01-x86_64.iso of=/dev/sda bs=4M conv=fsync oflag=direct status=progress

ok? (yes/NO) yes
[sudo] password for you:
843055104 bytes (843 MB, 804 MiB) copied, 156 s, 5.4 MB/s
201+1 records in
201+1 records out
846540800 bytes (847 MB, 807 MiB) copied, 157.142 s, 5.4 MB/s
```

## License
MIT
