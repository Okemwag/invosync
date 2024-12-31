# 📚 Books API

A simple RESTful API for managing books, built with **Go** and **Gorilla Mux**. 🚀

## 🌟 Features
- 📝 **CRUD Operations**: Create, Read, Update, and Delete books.
- ⚡ Lightweight and fast.


4. The API will run on **http://localhost:8080**. 🎉

---

## 📖 API Endpoints

### Base URL: `http://localhost:8080`

| Method | Endpoint         | Description            |
|--------|------------------|------------------------|
| GET    | `/books`         | Get all books 📚       |
| GET    | `/books/{id}`    | Get a book by ID 🔍    |
| POST   | `/books`         | Create a new book ✍️   |
| PUT    | `/books/{id}`    | Update a book 🔄       |
| DELETE | `/books/{id}`    | Delete a book ❌       |

---



---

## 🐳 Example Requests
### Add a Book
```bash
curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d '{"title":"1984","author":"George Orwell","isbn":"111-222-333"}'
```

### Get All Books
```bash
curl -X GET http://localhost:8080/books
```

---

## 📜 License
Licensed under the **MIT License**. ✨

---

👾 **Contributions are welcome!** Feel free to open issues or submit pull requests. 💻

Happy coding! ❤️