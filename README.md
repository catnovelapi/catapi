# Catapi - Ciweimao Client Library for Go

Catapi is a Go client library for accessing the Ciweimao API, a popular Chinese web novel platform. It provides a simple and efficient way of interacting with the Ciweimao API, allowing users to fetch information, perform operations, and handle responses from the API in a Go-friendly way.

## Installation
Install Catapi by running:
```
go get github.com/catnovelapi/catapi
```

## Usage

First, import the library:
```go
import "github.com/catnovelapi/catapi/catapi"
```

To start using the library, you need to create a new client:
```go
client := catapi.NewCiweimaoClient()
```

You can then use this client to call various methods on the Ciweimao API. For example, to get account information:
```go
result, err := client.Ciweimao.AccountInfoApi()
if err != nil {
    log.Fatal(err)
} else {
    fmt.Println(result)
}
```

## Methods

The `CiweimaoClient` struct has the following methods:

- `NewCiweimaoClient()` - Creates a new `CiweimaoClient`.
- `SetVersion(version string)` - Sets the version for the client.
- `SetDebug()` - Sets the client in Debug mode. Logs will be written to `catapi.log`.
- `SetProxy(proxy string)` - Sets the proxy for the client.
- `SetLoginToken(loginToken string)` - Sets the login token for the client.
- `SetAccount(account string)` - Sets the account for the client.
- `SetAuth(account, loginToken string)` - Sets both the account and login token for the client.

## API Endpoints

Here are the main API endpoints that the `Ciweimao` struct provides:

### Account Management
- `LoginApi(username, password string) (LoginResult, error)`: Logs in to an account.
- `AccountInfoApi() (AccountInfoResult, error)`: Retrieves account information.

### Book Management
- `BookshelfApi() (BookshelfResult, error)`: Returns the list of books on the user's bookshelf.
- `BookInfoApiByBookId(bookId string) (BookInfoResult, error)`: Retrieves book information by book ID.
- `BookInfoApiByBookName(bookName string) (BookInfoResult, error)`: Retrieves book information by book name.
- `ChaptersCatalogApi(bookId string) (ChaptersCatalogResult, error)`: Retrieves the catalog of chapters for a given book.
- `DownloadCover(url string) (DownloadCoverResult, error)`: Downloads a cover image.

### Search
- `SearchByKeywordApi(keyword, page string) (SearchResult, error)`: Searches for a book by keyword.

### Chapter Management
- `ChapterContentApi(bookId, chapterId string) (ChapterContentResult, error)`: Retrieves the content of a specific chapter.

### Rankings
- `RankingsApi() (RankingsResult, error)`: Retrieves the platform's rankings.

Please note that each of these methods returns a specific result struct and an error. For example, `LoginApi` returns a `LoginResult` and an `error`. You should always check the error before using the result.

## Contributing

Contributions are welcome! Please submit a pull request or create an issue on the GitHub page for this project.

## License

The Catapi library is open source and available under the [MIT License](https://opensource.org/licenses/MIT).

## Disclaimer

This library is not officially affiliated with Ciweimao. Please use responsibly and in accordance with Ciweimao's terms of service.