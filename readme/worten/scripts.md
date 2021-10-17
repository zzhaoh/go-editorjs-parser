# worten | scripts

## vpn up

```bash
sudo rm ~/vpn_openconnect.log
echo <VPN_PASSWORD> | sudo openconnect --background --user=rodrigocosta --passwd-on-stdin remote.sonaemc.com &> ~/vpn_openconnect.log
```

## vpn down

```bash
sudo kill -9 `ps -ef | grep openconnect | grep -v grep | awk '{print $2;}'`
```

## selenoid restart

```bash
docker container restart $(docker container ls -aq --filter "name=selenoid" | awk '{print $1;}')
```