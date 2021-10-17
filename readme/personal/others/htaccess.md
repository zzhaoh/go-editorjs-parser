# others | htaccess

## Linux

### Force use of HTTPS

```
RewriteEngine On 
RewriteCond %{SERVER_PORT} 80 
RewriteRule ^(.*)$ https://www.dominio.com.br/$1 [R,L]
```

### Change domain

```
RewriteEngine on
RewriteCond %{HTTP_HOST} ^dominioantigo.com.br$
RewriteRule (.*) http://www.novodominio.com.br/$1 [R=301,L]
```

### Change publishing point

```
RewriteEngine On
RewriteCond %{HTTP_HOST} ^(www.)?seudominio.com.br$
RewriteRule ^(/)?$ diretorio [L]
```

### Change subdomain publishing point

```
RewriteEngine on
RewriteCond %{HTTP_HOST} ^subdominio.seudominio.com.br$ [NC,OR]
RewriteCond %{HTTP_HOST} ^www.seudominio.com.br$ -
RewriteCond %{REQUEST_URI} !diretorio/
RewriteRule (.*) /diretorio/$1 [L]
```

### Force use of “www”

```
RewriteEngine on
RewriteCond %{HTTP_HOST} ^dominio\.com\.br
RewriteRule ^(.*)$ http://www.dominio.com.br/$1 [R=permanent,L]
```