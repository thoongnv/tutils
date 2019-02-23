# Tzutils

Quickly generate commit message inside a Gitlab repository

## Installing

* Duplicate `tzutils.yaml.default` to `tzutils.yaml` and modify

* Install and build executable
```
go get github.com/thoongnv/tzutils
go install github.com/thoongnv/tzutils
```

* Install Linux, Unix clipboard command ('xclip' or 'xsel')
```
sudo apt install xclip xsel
```

* Run executable inside Gitlab repository
```
~/go/bin/tzutils
```
