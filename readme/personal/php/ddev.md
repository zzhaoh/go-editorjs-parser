# php | ddev

## Ddev

- *ddev is an open source tool that makes it simple to get local PHP development environments up and running in minutes*
- <https://github.com/drud/ddev>

### Basic Usage

#### Installation

```bash
brew install drud/ddev/ddev
```

#### Update

```bash
ddev poweroff && brew upgrade drud/ddev/ddev
```

#### Config

*Go to the project folder and run:*

```bash
ddev config
```

#### Start

```bash
ddev start
```

#### Stop

```bash
ddev stop
```

#### See Logs

```bash
ddev logs
```

#### Follow Logs

```bash
ddev logs -f
```

#### Follow Logs DB

```bash
ddev logs -s db
```

#### Delete project

```bash
ddev delete --omit-snapshot <PROJECT>>
```

### Remove ddev

```bash
ddev delete --omit-snapshot <PROJECT>
ddev hostname --remove-inactive
rm -rf ~/.ddev
rm -rf .ddev
ddev delete images
docker rm $(docker ps -a | awk '/ddev/ { print $1 }')
docker rmi $(docker images | awk '/ddev/ {print $3}')
docker rmi -f $(docker images -q)
docker volume rm $(docker volume ls | awk '/ddev|-mariadb/ { print $2 }')
brew uninstall ddev
sudo rm /usr/local/bin/ddev
```