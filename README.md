# `snixr` 👨‍💻🔗

`snixr` is a fun and easy-to-use link shortening service built with Go, Docker, and Redis. It's perfect for when you need to share long URLs on social media, in emails, or anywhere else.

## 🚀 Getting Started

To get started with `snixr`, simply clone the repository and follow these steps:

1. Make sure you have `docker` and `docker-compose` installed on your system.
1. Create a `.env` file based on the `.env.example` file in the root directory of the repository.
1. Run `docker-compose up` in the root directory of the repository.
1. Access snixr at `http://localhost`.

That's it! You're now ready to start shortening links with snixr.

## 🌩️ Cloud Native

`snixr` was built with cloud-native principles in mind, which means it's optimized for deployment in a cloud environment. It uses Docker containers to package the application and its dependencies, making it easy to deploy and scale on any cloud platform that supports Docker.

## 🎯 Features

Some of the features of snixr include:

- Fast and efficient link shortening using custom hash-based algorithm.
- Shortened links that are easy to share and remember.
- Customizable link expiration times. (coming soon)
- API endpoints for programmatic link shortening.

## 📝 Contributing

We welcome contributions to `snixr`! To contribute, simply fork the repository, make your changes, and submit a pull request.

## 📜 License

`snixr` is open-source software licensed under the [MIT license](https://github.com/sschneemelcher/snixr/blob/main/LICENSE).

