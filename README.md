func-version
----
> Version Check Service

### Get Latest GitHub Release
```shell
# repo: github repo name
# prerelease: include prerelease, default false
curl "https://version.func.dongfg.com/github?repo=dani-garcia/vaultwarden&prerelease=true"
```

### Get Latest NPM Package
```shell
# pkg: npm package name
curl "https://version.func.dongfg."

{"data":"6.11.2","msg":"success","timestamp":1701490585}
```