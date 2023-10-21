# Cat

Cat is a Go package that provides a client for interacting with the Ciweimao API. Ciweimao is an online platform for
reading Chinese novels. This package allows you to perform various operations such as retrieving account information,
searching for books, accessing book chapters, and more.

## Installation

To use Cat in your Go project, you can simply import it using the following command:

```go
import "github.com/catnovelapi/catapi"
```

## Usage

### Creating a Cat Client

To create a Cat client, you need to import the package and create an instance of the `Ciweimao` struct:

```go
catClient := cat.NewCiweimaoClient()
```

### Account Information

You can retrieve the account information using the `AccountInfoApi` method:

```go
accountInfo := catClient.AccountInfoApi()
```

### Book Catalog

To retrieve the catalog of a book by its ID, you can use the `CatalogByBookIDApi` method:

```go
catalog := catClient.CatalogByBookIDApi(bookID)
```

### New Catalog

To retrieve the new catalog of a book by its ID, you can use the `NewCatalogByBookIDApi` method:

```go
newCatalog := catClient.NewCatalogByBookIDApi(bookID)
```

### Book Information

To retrieve the information of a book by its ID, you can use the `BookInfoApi` method:

```go
bookInfo := catClient.BookInfoApi(bookID)
```

### Searching for Books

You can search for books using keywords and pagination using the `SearchByKeywordApi` method:

```go
searchResults := catClient.SearchByKeywordApi(keyword, page)
```

### Signing Up

To sign up with a new account, you can use the `SignupApi` method:

```go
account, loginToken := catClient.SignupApi(account, password)
```

### Chapter Commands

You can retrieve the commands for a specific chapter using the `ChapterCommandApi` method:

```go
chapterCommands := catClient.ChapterCommandApi(chapterID)
```

### Chapter Information

To retrieve the information of a chapter, including its content, you can use the `ChapterInfoApi` method:

```go
chapterInfo := catClient.ChapterInfoApi(chapterID, command)
```

### Geetest API

You can use the `useGeetestApi` method to check if the Geetest API needs to be used for login:

```go
needGeetest := catClient.useGeetestApi(loginName)
```

### Bookshelf ID List

To retrieve the list of bookshelf IDs, you can use the `BookShelfIdListApi` method:

```go
shelfIDs := catClient.BookShelfIdListApi()
```

### Bookshelf Book List

To retrieve the list of books in a bookshelf by its ID, you can use the `BookShelfListApi` method:

```go
shelfBooks := catClient.BookShelfListApi(shelfID)
```

### Bookmark List

To retrieve the list of bookmarks for a book, you can use the `BookmarkListApi` method:

```go
bookmarks := catClient.BookmarkListApi(bookID, page)
```

### Division List

To retrieve the list of divisions (chapters) for a book, you can use the `DivisionListApi` method:

```go
divisions := catClient.DivisionListApi(bookID)
```

### Tsukkomi Number

To retrieve the number of Tsukkomi (comments) for a chapter, you can use the `TsukkomiNumApi` method:

```go
tsukkomiNum := catClient.TsukkomiNumApi(chapterID)
```

### Bdaudio Information

To retrieve the Bdaudio information for a book, you can use the `BdaudioInfoApi` method:

```go
bdaudioInfo := catClient.BdaudioInfoApi(bookID)
```

### Add Readbook

To add a book to the readbook list, you can use the `AddReadbookApi` method:

```go
readbookResult := catClient.AddReadbookApi(bookID, readTimes, getTime)
```

### Set Last Read Chapter

To set the last read chapter for a book, you can use the `SetLastReadChapterApi` method:

```go
setLastReadResult := catClient.SetLastReadChapterApi(lastReadChapterID, bookID)
```

### Privacy Policy Version

To post the privacy policy version, you can use the `PostPrivacyPolicyVersionApi` method:

```go
privacyPolicyResult := catClient.PostPrivacyPolicyVersionApi()
```

### Property Information

To post the property information, you can use the `PostPropInfoApi` method:

```go
propInfo := catClient.PostPropInfoApi()
```

### Meta Data

To retrieve the meta data, you can use the `MetaDataApi` method:

```go
metaData := catClient.MetaDataApi()
```

### Version

To retrieve the version information, you can use the `VersionApi` method:

```go
versionInfo := catClient.VersionApi()
```

### Startpage URL List

To retrieve the list of startpage URLs, you can use the `StartpageUrlListApi` method:

```go
startpageURLs := catClient.StartpageUrlListApi()
```

### Third Party Switch

To retrieve the third-party switch information, you can use the `ThirdPartySwitchApi` method:

```go
thirdPartySwitch := catClient.ThirdPartySwitchApi()
```

### Cat Options

The Cat client provides several options that can be used to customize its behavior. These options can be passed as
arguments when creating the client instance. The available options are:

- `Debug()`: Enables debug mode.
- `NoDecode()`: Disables decoding of the API response.
- `MaxRetry(retry)`: Sets the maximum number of retries for API requests.
- `ApiBase(host)`: Sets the API base host.
- `Version(version)`: Sets the API version.
- `DecodeKey(decodeKey)`: Sets the decode key.
- `DeviceToken(deviceToken)`: Sets the device token.
- `AppVersion(appVersion)`: Sets the app version.
- `LoginToken(loginToken)`: Sets the login token.
- `Account(account)`: Sets the account.

These options can be combined as needed. Here's an example of how to create a Cat client with custom options:

```go
catClient := cat.NewCiweimaoClient(
cat.Debug(),
cat.MaxRetry(3),
cat.ApiBase("https://api.example.com"),
)
``` 