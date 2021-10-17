# javascript | nvm

## Installation of NVM - MacOs/Linux

### Using curl
```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash
```

### Using wget
```bash
wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash
```

### Verify Installation
```bash
command -v nvm
```

## Usage - MacOs/Linux

### Get a List of all available NodeJs versions
```bash
nvm ls-remote
```

### Install latest NodeJs version
```bash
nvm install node
```

### Install latest LTS Release
```bash
nvm install --lts
```

### Install particular NodeJs version
```bash
nvm install 12.13.1
```

### UnInstall the multiple NodeJs version
```bash
nvm uninstall 8.11.1
```

## Switching the NodeJs version - MacOs/Linux

### Get a List of installed NodeJs version
```bash
nvm ls
```

### Switching Between installed NodeJs Versions
```bash
nvm use 12.13.1
```

## Use Custom alias for installed NodeJs versions

### Create an alias
```bash
nvm alias awesome-project 12.13.1
```

### Switch to the alias
```bash
nvm use awesome-project
```

### Remove the Alias
```bash
nvm unalias awesome-project
```

## Other commands

### Run to specific NodeJs version, without switching
```bash
nvm run 8.11.1 app.js
```
or
```bash
nvm exec 8.11.1 node app.js
```

### To use installed nodeJs version in the system
```bash
nvm use system
```

### To completely remove nvm from the system
```bash
nvm unload 
```