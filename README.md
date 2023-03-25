# `snixr` 👨‍💻🔗

Welcome to `snixr`, the link shortener service that's quick, easy, and fun to use! 😎

## About `snixr`

`snixr` is a fast and lightweight link shortener service built with Go and the Fiber web framework. With `snixr`, you can easily create short links for all of your favorite websites, share them with your friends, and track clicks to see how many people are visiting your links. 🚀

## Getting Started

To get started with `snixr`, simply clone the repository and run the following command:

```shell
go run main.go
```

This will start the `snixr` server and allow you to access the `snixr` API at http://localhost:3000. 🚀

## API Documentation

`snixr` provides a simple and intuitive API for creating, retrieving, and tracking short links. The API endpoints are:

- `POST /api/links` - Create a new short link
- `GET /:code` - Redirect to the original URL for a given short link code
- `GET /api/links/:id` - Retrieve information about a specific short link
- `GET /api/links` - Retrieve a list of all short links

## Contributing

We welcome contributions to the `snixr` project! To get started, simply fork the repository, make your changes, and submit a pull request. 

## License

`snixr` is open source software released under the MIT License. For more information about the `snixr` license, please refer to the LICENSE file. 📜


We hope you enjoy using `snixr` as much as we enjoyed building it! Please don't hesitate to reach out if you have any questions or feedback. 😊

