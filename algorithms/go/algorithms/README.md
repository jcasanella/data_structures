When the package is local do the following steps:

* Import from another package using the module you want to import.
* Call the function using packageName.function
* Edit the module from where you're importing: `go mod edit -replace=github.com/jcasanella/algorithms=../algorithms`
* Run this command: `go mod tidy`
* Now you can run the app