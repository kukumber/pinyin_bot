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
git clone https://github.com/yourusername/pinyin_bot.git
cd pinyin_bot
go mod tidy
```

### Configuration Parameters
The application can be configured using the following parameters:
- `--api-key`: your telegram bot API key
- `--init-offset`: initial telegram message offset
- `--update-interval`: update interval in seconds, default 60

### Environment Variables
The application can be configured using the following environment variables:
- `API_KEY`: your telegram bot API key
- `INIT_OFFSET`: initial telegram message offset
- `UPDATE_INTERVAL`: update interval in seconds, default 60

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

### Contributing
Contributions are welcome. Please open an issue or pull request.