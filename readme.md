# 🚀 Pblog - A TUI Blogging App

Pblog is a **terminal-based user interface (TUI) blogging application** that allows users to **create, read, update, and delete posts** efficiently. It supports **similarity-based search** and **paginated fetching** to ensure smooth querying without overloading the database.

## ✨ Features
- ✅ **CRUD Operations**: Manage your blog posts seamlessly.
- 🔍 **Similarity-Based Search**: Find posts based on content relevance.
- 📌 **Paginated Fetching**: Limits the number of posts retrieved per query to optimize performance.
- ⏭️ **Next Button for Pagination**: Fetches the next set of posts starting after the last retrieved post.

## 🔮 Future Enhancements
- 🎯 **Search Filtering**: Filter search results through categories.
- 🔐 **Authentication**: Implement login and write permissions.
- 🏷️ **User Preferences**: Posts will be recommended based on user preference tags.
- 💰 **Payment Gateway**: Enable monetization options.
- 🏗️ **PostgreSQL Deployment**: Deploy a PostgreSQL instance on a server for public access.

## 🛠️ Installation & Usage
1. 📥 Clone the repository:
   ```bash
   git clone https://github.com/KAwasthi2889/PBlog
   ```
2. 📂 Navigate into the project directory and create a `db` directory:
   ```bash
   cd PBlog
   mkdir db
   ```
3. 📝 Inside the `db` directory, create a file named `database.env` and add the following environment variables:
   ```
   POSTGRES_USER=<your preferred user name>
   POSTGRES_PASSWORD=<your preferred password>
   POSTGRES_DB=<your preferred db name>
   ```
   ⚠️ **Note**: Ensure there are no spaces around the `=` sign.

4. 🚀 Start the application using Docker:
   ```bash
   docker compose up -d && docker attach tui-app && docker compose down
   ```

✅ And you are good to go! 😃

📢 **Stay tuned for updates!**