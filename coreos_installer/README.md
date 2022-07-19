# coreos-installer

Create butane config
```bash
butane < config.bu > config.ign
```

Modify ISO image
```bash
coreos-installer download <some flags to get .iso for x86_64 metal>
coreos iso customize --live-ignition ./butane/config.ign fcos.iso
```
