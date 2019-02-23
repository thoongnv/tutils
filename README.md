# Tutils

Quickly generate commit message inside a Gitlab repository

## Installing

* Duplicate `tutils.yaml.default` to `tutils.yaml` and modify

* Install and build executable
```
go get github.com/thoongnv/tutils
go install github.com/thoongnv/tutils
```

* Install Linux, Unix clipboard command ('xclip' or 'xsel')
```
sudo apt install xclip xsel
```

* Run executable inside Gitlab repository
```
~/go/bin/tutils
```
