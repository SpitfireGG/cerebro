cerebro
 ❌ detect mutants ✅ detect LLMs 

cerebro is a intuitive Terminal User Interface designed to make requests to & access Large Language Models functionalities. This document provides everything you need to know to get started, configure the application, and troubleshoot any issues.
✨ Features

    Interactive Interface: A rich, interactive TUI that's easy to navigate.

    Highly Configurable: Customize everything from appearance to behavior.

    Extensive Logging: Detailed logs for easy debugging.
---

💾 Installation

Getting started is simple. Just clone the repository and run the application.

# Clone the repository
git clone https://github.com/spitfireGG/cerebro.git

# Navigate to the project directory
cd cerebro

# Run the TUI
go run cmd/cerebro/main.go

---

🕹️ Usage

Once the application is running, you can navigate using the arrow keys. Here are some common keys:

    ↑ / ↓: Navigate up and down lists.

    ← / →: Switch between panels or tabs.

    Enter: Select an item.

    q or Ctrl+C: Quit the application.

⚙️ Configuration

You can customize the TUI's behavior via within the model's configuration option provided in the TUI itself

    💡 Tip: You do not need any specific YAML files or anything similar for configuration

---

# debuggings
You can use logging functionality to inspect the request & other debugging stuffs by running:
```bash

# to run debugging functionalities run: 
DEBUG=1 go run cmd/cerebro/main.go 

# inspect the log messages being logged followed by the above command
tail -f debug.log

#inspect the http requests
tail -f gemini_req_dump.log http_req_dump_gemini.log
```

---

Happy Hacking! 🎉
