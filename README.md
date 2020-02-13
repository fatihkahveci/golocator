# GoLocator

GoLocator is a Go library for dealing with simple Country data.

## Installation

```bash
go get -u github.com/fatihkahveci/golocator
```

After installation package you need to download "golocator.db" file to your project folder.

## Usage

```go
package main

import (
	"fmt"

	"github.com/fatihkahveci/golocator"
)

func main() {
	locator := golocator.NewGoLocatorService("golocator.db")
	country, _ := locator.FindCountryByIso2("US")
	fmt.Println(country)
	fmt.Println(country.PhoneCode)
	fmt.Println(locator.FindCities(country))
}

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)