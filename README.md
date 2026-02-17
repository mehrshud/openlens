# 📊 OpenLens
[![Build Status](https://img.shields.io/github/actions/workflow/status/mehrshud/OpenLens/build.yml?branch=main)](https://github.com/mehrshud/OpenLens/actions)
[![License](https://img.shields.io/github/license/mehrshud/OpenLens)](https://github.com/mehrshud/OpenLens/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/mehrshud/OpenLens)](https://github.com/mehrshud/OpenLens/blob/main/go.mod)
[![Stars](https://img.shields.io/github/stars/mehrshud/OpenLens)](https://github.com/mehrshud/OpenLens)
[![Issues](https://img.shields.io/github/issues/mehrshud/OpenLens)](https://github.com/mehrshud/OpenLens/issues)
[![Codecov](https://img.shields.io/codecov/c/github/mehrshud/OpenLens)](https://codecov.io/gh/mehrshud/OpenLens)
![Demo](docs/assets/demo.gif)
**Unlock the full potential of your data with OpenLens, the open-source alternative to proprietary data analytics platforms.**

## ✨ Features
* 📈 Data Analytics
* 📊 Data Visualization
* 🚀 Fast and Scalable
* 📁 Support for Multiple Data Sources

## 🚀 Quick Start
To get started with OpenLens, run the following commands:
git clone https://github.com/mehrshud/OpenLens.git
cd OpenLens
go build -o openlens main.go
./openlens
Open your web browser and navigate to `http://localhost:8080` to access the OpenLens dashboard.

## 📐 Architecture
graph TD
  A[Client] -->|HTTP| B[API Gateway]
  B --> C[Service Layer]
  C --> D[Database]
  D --> E[Cache]
  E --> F[GraphQL API]
  F --> G[Frontend]
  G --> H[User]
The OpenLens architecture consists of a client, API gateway, service layer, database, cache, GraphQL API, and frontend.

## 📦 Installation
To install OpenLens, you can use the following methods:
* **Docker**: `docker run -p 8080:8080 mehrshud/openlens`
* **Binary**: Download the latest binary from the [Releases](https://github.com/mehrshud/OpenLens/releases) page and run it on your system.

## 🔧 Configuration
The following configuration options are available:
| Environment Variable | Description | Default Value |
| --- | --- | --- |
| `DEBUG` | Enable debug mode | `false` |
| `DB_URL` | Database URL | `localhost:5432` |
| `CACHE_TTL` | Cache TTL | `1h` |

Create a `.env` file in the root of the repository and add your configuration options.

## 🤝 Contributing
To contribute to OpenLens, follow these steps:
1. Fork the repository
2. Create a new branch for your feature or bug fix
3. Submit a pull request

## 📊 GitHub Stats:
[![Stats](https://github-readme-stats.vercel.app/api?username=mehrshud&show_icons=true&theme=radical)]()

## 📄 License
OpenLens is licensed under the MIT license.

Made with ❤️ by [mehrshud](https://github.com/mehrshud)