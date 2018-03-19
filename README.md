# crudgo
CRUD Golang


## Prerequisites

* [Golang](https://github.com/golang/go) 
* [GoVendor](https://github.com/kardianos/govendor)
* [Sqlite](https://www.sqlite.org/index.html) or [PostgreSQL](https://www.postgresql.org/)
* [Gitflow](https://github.com/nvie/gitflow) For Contributing


## Installing

### Get this project

```
go get github.com/julioc98/crudgo
```
### Change directory

```
cd $GOPATH/src/github.com/julioc98/crudgo
```

### Installing the dependencies

```
govendor sync
``` 

### Run project( Default Env vars)

```
go run app.go
```

### or

### Run project( Change Env vars) Ex.:

```
DB=postgres DATABASE_URL=postgres://user:password@host:5432/database go run app.go
```

## Contributing

Please read [CONTRIBUTING.md](https://github.com/julioc98/crudgo/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/julioc98/crudgo/tags).

## Authors

* **[Julio Cesar](https://julioc98.github.io)**
* **[Katarina Massako Inoue](https://github.com/katarinamai)**

See also the list of [contributors](https://github.com/julioc98/crudgo/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.