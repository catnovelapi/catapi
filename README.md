# CatNovelAPI: Detailed Documentation

`CatNovelAPI` is a Go-based API client library that provides an easy-to-use interface for interacting with the `Ciweimao` API, a service that provides access to various book-related data and functionalities.

## Installation

To use `CatNovelAPI`, you are required to import the library in your Go project. Here is the import statement:

```go
import (
	"github.com/catnovelapi/catapi/catapi"
	"github.com/catnovelapi/catapi/options"
)
```

## Usage

### Creating a New Client

To create a new instance of the `Ciweimao` client, you can use the `NewCiweimaoClient` function. It accepts an arbitrary number of options that can be used to configure the client. This function returns a pointer to a new `Ciweimao` instance with the provided options applied.

```go
client := catapi.NewCiweimaoClient(
    options.Debug(),
    options.Version("2.9.290"),
    // other options...
)
```

### Client Options

The `options` package provides several functions that return `CiweimaoOption` instances. These can be used to configure the `Ciweimao` client. The available options are:

- `Debug()`: Enables debug mode. In debug mode, additional information might be logged (typically useful for development and debugging).
- `Version(string)`: Specifies the version of the API to use.
- `Proxy(string)`: Sets a proxy to use for API requests.
- `LoginToken(string)`: Sets the login token for the client. This is required for API methods that need authentication.
- `Account(string)`: Sets the account name for the client.
- `Auth(string, string)`: Sets both the account name and login token for the client.

### API Methods

The `Ciweimao` client provides several methods that correspond to the different API endpoints. These methods return a `gjson.Result` object and an error (if any occurred during the API request). Following are descriptions of some of these methods:

- `AccountInfoApi()`: Retrieves account information. Accepts no parameters.
- `CatalogByBookIdApi(string)`: Retrieves the catalog for a specific book by its `bookID`.
- `BookInfoApi(string)`: Retrieves information about a specific book by its `bookID`.
- `SearchByKeywordApi(string, string, string)`: Searches for books by keyword, page, and category index.
- `SignupApi(string, string)`: Signs up a new user with the provided `account` and `password`.

For a full list of API methods and their descriptions, please refer to the `catapi` package documentation.

## Contributing

Contributions to the `CatNovelAPI` are welcome. Whether you find a bug, have a suggestion for an improvement, or want to add a new feature, we appreciate your help. Please submit a pull request or create an issue to contribute.

## License

`CatNovelAPI` is released under the MIT license. This license is a short, permissive software license. Basically, you can do whatever you want with this software, as long as you include the original copyright and license notice in any copy of the software/source. For more information, see the LICENSE file in the repository.