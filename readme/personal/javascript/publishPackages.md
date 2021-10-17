# javascript | publish packages

## Creating your first package

This section is for you if you haven’t published a package to npm before. Feel free to skip to the next section if you’ve published one before.

To publish your first package to npm, you need to go through these steps:

First, you need to have a npm account. Create one [here](https://www.npmjs.com/signup) if you don’t have one yet.

Second, you need to login to your npm account through the command line. (You need to have Node and npm installed on your system before you perform this step. Install them [here](https://nodejs.org/en/)).

To sign in, you use

```bash
npm login
```

You’ll be prompted to enter your username, password, and email address.

```bash
[~] npm login                                                                                                                                                                                   ──(Wed,Sep15)─┘
Username: rodrigoodhin
Password: 
Email: (this IS public) rodrigoodhin@gmail.com
Logged in as rodrigoodhin on https://registry.npmjs.org/.
```

Third, you need to create a package. To do so, create a folder somewhere on your computer and navigate to it. The command line version is:

```bash
# Creating a folder named how-to-publish-to-npm# Npm | Publishpackages# . | Npm | Publishpackages# . | Npm | Publishpackages# . | Npm | Publishpackages# Npm | Publishpackages# npm | publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages
mkdir how-to-publish-to-npm

# Navigating into the folder# Npm | Publishpackages# . | Npm | Publishpackages# . | Npm | Publishpackages# . | Npm | Publishpackages# Npm | Publishpackages# npm | publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages# Npm | Publish Packages
cd how-to-publish-to-npm
```

Next, you want to begin the project with the command:

```bash
npm init
```

This command runs you through a few questions and creates a `package.json` file for you at the end. This `package.json` file contains the bare necessities you need to publish your project. (Feel free to skip questions that don’t make sense).

```bash
[~] npm init                                                                                                                                                                                    ──(Wed,Sep15)─┘
This utility will walk you through creating a package.json file.
It only covers the most common items, and tries to guess sensible defaults.

See `npm help init` for definitive documentation on these fields
and exactly what they do.

Use `npm install <pkg>` afterwards to install a package and
save it as a dependency in the package.json file.

Press ^C at any time to quit.
package name: (how-to-publish-to-npm) 
version: (1.0.0) 
description: how to publish to npm
entry point: (index.js) 
test command: 
git repository: 
keywords: 
author: 
license: (ISC) MIT
About to write to /Users/rodrigo/VirusPandora/GolangProjects/readme/how-to-publish-to-npm/package.json:

{
  "name": "how-to-publish-to-npm",
  "version": "1.0.0",
  "description": "how to publish to npm",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "",
  "license": "MIT"
}

Is this OK? (yes) yes
```

The final step is to publish your package with the command:

```bash
npm publish --access=public
```

### Configure public access in `package.json`

```json
{
  "publishConfig": {
    "access": "public"
  }
}
```

### Link with more information

- <https://zellwk.com/blog/publish-to-npm/>