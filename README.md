# conpack

🚀 **Container Package Checking Utility**

`conpack` is a command-line tool to check for the presence of a specified package in running containers using various container runtimes like Docker, Podman, and Finch.

## 📋 Usage

```sh
conpack [-p|--package <package_name>] [-r|--runtime <runtime>]
```

### Options

- `-p, --package <package_name>`: Specify the package name to search for.
- `-r, --runtime <runtime>`: Specify the container runtime to use (e.g., docker, podman, finch). Default is docker.
- `--help`: Show this help message.

### Examples

```sh
conpack -p curl
conpack -p curl -r podman
```

## 🛠️ How It Works

1. **Initialization**: The tool initializes by parsing command-line flags for the package name and runtime.
2. **Container Check**: It lists all running containers using the specified runtime.
3. **Package Search**: For each container, it checks if the specified package is installed by executing a command inside the container.
4. **Output**: Displays the containers where the package is found, along with their names and IDs.

## ✨ Features

- Supports multiple container runtimes.
- Provides a colorful and animated progress indicator while checking containers.
- Displays results in a clear and organized manner.

## 📜 License

This project is licensed under the MIT License.

## 👥 Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## 📧 Contact

For any questions or suggestions, feel free to reach out.

Happy container package checking! 🎉
