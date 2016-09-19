### Install Vim Editor
* Clone this repository and change directory to `vim-editor`
* [On CoreOS] Run `./run install`.
* [On other Distros] 1) Copy `install` script to `/usr/local/bin/editor` 2) Replace `EDITOR_VERSION` with `latest` in `/usr/local/bin/editor`.
* Run `editor` from terminal to launch editor.

### Mounts
`editor` launches a container and mounts `/workspace` directory from host OS to inside the container.

### Usage
Place your files/repos under `/workspace` directory and edit them from inside the `editor` container.
