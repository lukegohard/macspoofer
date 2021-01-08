### macspoofer

****

**Build:**

```bash
make build
```

****

**Use:**

*Setting manually a mac address:*

```bash
sudo ifconfig <device> down
sudo macspoofer -c -w <device> -m <mac address>
sudo ifconfig <device> up
```

*Setting a random mac address:*

```bash
sudo ifconfig <device> down
sudo macspoofer -c -w <device> -m random
sudo ifconfig <device> up
```

*Printing current mac address:*

```bash
macspoofer -s -w <device>
```

*Printing random mac address:*

```bash
macspoofer -r 
```



****

**Example:**

```bash
sudo macspoofer -c -w wlan0 -m d9:51:fe:66:fa:ab
```

```bash
sudo macspoofer -c -w wlan0 -m random
```

```bash
macspoofer -s -w wlan0	
```

****

**Flag:**

- **-s** ---> print current mac address
- **-c** ---> change mac address
- **-w** ---> define device
- **-m** ---> define new mac address(xx:xx:xx:xx:xx:xx or random)
- **-r** ---> print a random mac address