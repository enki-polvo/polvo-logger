# `polvo-logger`

<p align="center">
  <img src="https://github.com/user-attachments/assets/d100f6b1-1465-49a3-aad4-21030f0e64f3" width="250" height="250" alt="The mascot of polvo-logger"/>
</p>



`polvo-logger` is a simple logging utility written in Go.

It provides a structured and efficient way to log messages for your applications.

## Features

- Lightweight and easy to use
- Structured logging for better readability
- Configurable log levels

## Note for Developers

Activate the project's **Git hook** to ensure your commit messages follow the recommended conventions.

After cloning this repository, run:

```sh
chmod +x .githooks/commit-msg
git config core.hooksPath .githooks
git config core.hooksPath   # To verify
```

If you made changes to the project, don't forget to create the PlantUML diagram by running:

```sh
goplantuml -recursive . > ./paper/classDiagram.puml
```
