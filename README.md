# relax-dd

Relax you and run dd.

## Install

```
$ go install github.com/buzztaiki/relax-dd@latest
```

## Usage

```console
$ sudo relax-dd ~/Downloads/archlinux-2022.03.01-x86_64.iso /dev/sda
src:
/home/you/Downloads/archlinux-2022.03.01-x86_64.iso: ISO 9660 CD-ROM filesystem data (DOS/MBR boot sector) 'ARCH_202203' (bootable)

dst:
NAME   FSTYPE  FSVER            LABEL       UUID                                 FSAVAIL FSUSE% MOUNTPOINTS
sda    iso9660 Joliet Extension ARCH_202203 2022-03-01-15-50-40-00
└─sda1 vfat    FAT32            USB         42B5-F2FC                               3.7G     0% /run/media/you/USB

execute the follwing command:
  /usr/bin/dd if=/home/you/Downloads/archlinux-2022.03.01-x86_64.iso of=/dev/sda bs=4M status=progress conv=fsync oflag=direct
ok? (yes/NO) yes
838860800 bytes (839 MB, 800 MiB) copied, 126 s, 6.6 MB/s
201+1 records in
201+1 records out
846540800 bytes (847 MB, 807 MiB) copied, 127.476 s, 6.6 MB/s
```

## License
MIT
