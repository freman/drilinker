# DRI Linker

DRI Linker is a tiny tool to create links /dev/dri/renderD* named after the driver for that card.

This is *my* solution to the problem I was having with the cards swapping between /dev/dri/renderD128 and /dev/dri/renderD129 after every boot
meaning I had to update the Jellyfin hardware accelerator configuration almost every time there was a reboot.

After running this tool I get:
- renderD128
- renderD129
- renderNi915
- renderNnouveau

There is one minor catch, your binary needs to be owned by root and have the suid bit set, or you need to run it through sudo/as root
I am running mine in an extension to the jellyfin/jellyfin docker container so I use the following

```dockerfile
ADD drilinker /usr/bin/drilinker
RUN chmod +s /usr/bin/drilinker
```
