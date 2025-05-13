
# Book Management System

Book Management System API developed as a hands-on exercise to learn and apply foundational concepts of Go (Golang) and RESTful API development. It serves as a practical implementation to understand how to build, structure using Go.


## API Reference

#### Get all books

```http
  GET /book/
```
#### Create a book

```http
  POST /book/
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `Name`      | `string` | Name of the Book |
| `Author`      | `string` | Author of the Book |
| `Publication`      | `string` | Publication of the Book |

#### Get a Book

```http
  GET book/{bookId}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bookId`      | `string` | **Required**. Id of book to fetch |

#### Update a book

```http
  PUT /book/{bookId}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bookId`      | `string` | **Required**. Id of book to which update |
| `Name`      | `string` | Name of the Book |
| `Author`      | `string` | Author of the Book |
| `Publication`      | `string` | Publication of the Book |

```http
  DELETE book/{bookId}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bookId`      | `string` | **Required**. Id of book to Delete. |