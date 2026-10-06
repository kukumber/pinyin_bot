# Pinyin Bot

## Overview
Pinyin Bot is a Go-based application designed to provide text-to-Pinyin and Pinyin-to-Pallady conversion.

## Features
- **Pinyin Conversion**: Convert text with numbers into the Pinyin format.
- **Pallady Conversion**: Convert Pinyin text into the Cyrillic Pallady format.


## Getting Started

### Installation
Clone the repository and install dependencies:
```bash
git clone https://github.com/kukumber/pinyin_bot.git
cd pinyin_bot
go mod tidy
```

### Configuration Parameters
The application can be configured using the following parameters:
- `--api-key`: your telegram bot API key
- `--init-offset`: initial telegram message offset
- `--update-interval`: update interval in seconds, default 60
- `--debug`: enable telegram API debug logging, default false

Flags take precedence over environment variables, which take precedence over defaults.

### Environment Variables
The application can be configured using the following environment variables:
- `API_KEY`: your telegram bot API key
- `INIT_OFFSET`: initial telegram message offset
- `UPDATE_INTERVAL`: update interval in seconds, default 60
- `DEBUG`: enable telegram API debug logging, default false

### Running the Application
To start the application, run:
```bash
go run cmd/main.go
```

### Building the Application
To build the application, run:
```bash
go build -o pinyin_bot cmd/main.go
```

### Running with Docker
The application ships with a `Dockerfile` and `docker-compose.yml`. Create a `.env` file with your configuration (at minimum `API_KEY`), then build and start:
```bash
docker compose build
docker compose up -d
```

The image is tagged `pinyin_bot:${IMAGE_TAG:-latest}`. To build or deploy a specific version, set `IMAGE_TAG` in the shell or in `.env`:
```bash
IMAGE_TAG=2.0.2 docker compose build
IMAGE_TAG=2.0.2 docker compose up -d
```

The container attaches to the `pinyin_net` bridge network created by compose.

### Contributing
Contributions are welcome. Please open an issue or pull request.
