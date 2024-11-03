# Magical Crawler 🧙‍♂️🐾

![Magic Crawler](https://github.com/Group8Delta/magical-crawler/blob/main/assets/mc.webp)

## 📖 Introduction
**Magical Crawler** is a powerful web crawling tool designed to automate the process of collecting, filtering, and managing web data. Created with a blend of efficiency and versatility, this tool is ideal for applications requiring regular data collection, such as monitoring classified ads, gathering analytics data, or tracking updates across multiple sources. With a secure, high-performance codebase, Magical Crawler ensures reliable and consistent results.

## 🛠️ Technology Stack
- **Golang** - core language for development
- **PostgreSQL** - primary data storage for crawled content
- **Gorm** - ORM for database management
- **Telegram Bot** - for real-time notifications
- **Docker** - for containerized deployment
- **Testify** - for testing and assertions

## ⚙️ Setup
To get started with Magical Crawler:

1. **Clone the Repository**
   ```bash
   git clone https://github.com/username/magical-crawler.git
   ```

2. **Navigate to the Project Directory**
   ```bash
   cd magical-crawler
   ```

3. **Install Dependencies**
   Ensure you have Go installed, then install required packages:
   ```bash
   go mod tidy
   ```

4. **Configure Environment**
   Set up your `.env` file with necessary configurations such as database connection details, Telegram bot credentials, and email settings.

5. **Start the Project with Docker**
   ```bash
   sudo docker-compose up --build
   ```

6. **Access the API**
   The API is accessible at `localhost:8080` by default.

## 🧪 Testing
Run tests to verify functionality and stability:

```bash
go test ../main.go
```
