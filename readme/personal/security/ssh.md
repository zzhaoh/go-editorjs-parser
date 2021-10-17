# security | ssh

## Fix *"WARNING: UNPROTECTED PRIVATE KEY FILE!"*

```bash
sudo chmod 600 ~/.ssh/id_rsa
sudo chmod 600 ~/.ssh/id_rsa.pub
```

## Fix *Failed to add the host to the list of known hosts (/home/USER/.ssh/known_hosts)*

```bash
sudo chmod 644 ~/.ssh/known_hosts
sudo chmod 755 ~/.ssh
```

## Generate a new SSH key

```bash
ssh-keygen
```

## Add an existing SSH key

```bash
eval `ssh-agent`
ssh-add ~/.ssh/id_rsa
```