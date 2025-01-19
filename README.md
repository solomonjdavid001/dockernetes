# 🖥️ Docker & Minikube Manager Desktop App 🚀

Welcome to the **Docker & Minikube Manager**, a powerful desktop application that allows you to manage Docker containers, visualize running containers, and interact with Minikube—all in one place! 🌟

---

## 🛠️ Features

- **Docker Management & Visualization**: View and manage your Docker containers, images, and monitor resource usage. 📊
- **Minikube Control**: Interact with Minikube running locally, start/stop clusters, and manage resources. 🎛️
- **Resource Monitoring**: View real-time **GPU**, **CPU**, and **RAM** usage to keep track of your system's performance. 📈

---

## 💻 Tech Stack

Our application is built using the following tech stack:

- **Go** for the backend 🐹
- **Go Gin Gonic** for APIs ⚙️
- **TypeScript** for frontend logic 🔤
- **TailwindCSS** for styling 🌈
- **Next.js** for the frontend framework 🌐
- **Wails** for building cross-platform desktop apps 🖥️
- **Shadcn** for UI components ⚙️

---

## 🌍 Supported Platforms

- **Linux** 🐧
- **Darwin** (macOS) 🍏
- **Windows** 🪟

## Running the Application

1. Prerequisites:

   - [Go](https://go.dev/doc/install) v1.19+
   - [Node.js](https://nodejs.org/en/download) v18+
   - [Wails](https://wails.io/docs/gettingstarted/installation/) v2.9.2+
   - [Docker](https://www.docker.com/products/docker-desktop/) v20.10.17+
   - [Redis](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/)

2. Clone the repository:'

```bash
git clone https://github.com/solomonjdavid001/Dockernetes.git
```

3. Running the application:

- Start the Docker daemon
- Start the Redis server

- Then run the application:

```bash
wails dev
```

4. Build the application for required platforms:

```bash
wails build -platform darwin/arm64,linux/amd64,windows/amd64
```
