# ssh-find

Finds SSH public keys by username. Provide a username and the application will return the content of SSH keys (e.g: `id_rsa.pub`) used in online services.

![Demo](demo/ssh-find-demo.gif)

## Use case example

Copy your public keys from GitHub (Mac):

    ssh-find alombarte | pbcopy

Or Linux:

    ssh-find alombarte | xclip -selection clipboard

Append them in `authorized_keys` of a remote server.

## Installation

Download from the [releases page](https://github.com/alombarte/ssh-find/releases), decompress, and move to `/usr/local/bin` or similar directory.
