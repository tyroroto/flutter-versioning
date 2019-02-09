# Flutter Versioning
### How it work
```
copy flutter-versioning to flutter project
execute flutter-versioning arg is version(5 Digit number or more)
example 
./go-flutter-version 56789
VERSION 56789
1.0.56

Version minor is version_number / 100000
Version patch is ( version_minor * 100000 - version_number) / 10000

version view change pubspec.yaml file then run build your flutter app
```

### Use with Gitlab-CI
```
 - ./flutter-versioning $CI_PIPELINE_ID
```