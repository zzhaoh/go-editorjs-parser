# tests | selenoid

## Install CM

```bash
sudo wget https://github.com/aerokube/cm/releases/download/1.7.2/cm_linux_amd64 -O /usr/local/bin/cm ; sudo chmod 777 /usr/local/bin/cm
```

## Cleanup selenoid config

```bash
cm selenoid cleanup
```

## Run Selenoid

```bash
cm selenoid start --browsers 'chrome;firefox;opera' --last-versions 2 --vnc
```

## Run Selenoid ui

```bash
cm selenoid-ui start
```