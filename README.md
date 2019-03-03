## ScreenD - Daemon with smaller daemons
Simple screen daemon written in Go.

### Installation
- `sudo su -`
- `cd /root`
- `git clone https://github.com/OkaeriPoland/screend`
- `cd screend`
- `go build -o screend`
- `chmod +x screend`
- `mv screend.sh /etc/init.d/screend`
- `chmod 755 /etc/init.d/screend`
- `chown root:root /etc/init.d/screend`
- `update.rc screend defaults`
- `/etc/init.d/screend start`

### Removal
- `update.rc screend remove`
- `sudo rm -rf /root/screend /etc/init.d/screend`

### Configuration
To add own screen to screend create new .ini file in `/root/screend/daemons`, then run `/etc/init.d/screend restart`

```ini
[Screen]
name = Example Daemon
user = example

[Execution]
run_directory = /home/example
command = bash

[Daemon]
enabled = true
```

### Protip
As current installation method sucks, you can edit `screend.sh` and install screend in way you like. If you know what you are doing, of course.