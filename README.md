# Cerebro

<div align="center">

**❌ detect mutants ✅ detect LLMs**

*A Terminal User Interface for seamless Large Language Model interactions*

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg)]()

</div>

---

## ✨ Features

### Progress Ongoing...

🎨 **Interactive Interface** - Rich, intuitive TUI that's a joy to navigate  
⚙️ **Highly Configurable** - Customize appearance, behavior, and more  
📊 **Extensive Logging** - Detailed debugging capabilities built-in  
🚀 **Multi-Model Support** - Access various LLMs from one interface  
🎯 **Keyboard-Driven** - Efficient navigation without touching your mouse  

## 🚀 Quick Start

### Prerequisites
- Terminal with Unicode support (for optimal experience)

### Installation

```bash
# Clone the repository
git clone https://github.com/spitfireGG/cerebro.git

# Navigate to project directory
cd cerebro

# Run cerebro
go run cmd/cerebro/main.go
```

## 🎮 Usage

Navigate cerebro with these intuitive keybindings:

| Key | Action |
|-----|--------|
| `↑` / `↓` | Navigate lists vertically |
| `←` / `→` | Switch panels/tabs |
| `Enter` | Select item |
| `Tab` | Move to next input field |
| `Esc` | Go back/cancel |
| `q` / `Ctrl+C` | Quit application |

## ⚙️ Configuration

Cerebro uses an in-app configuration system - no external files needed!

> 💡 **Tip**: All settings are accessible through the TUI's configuration panel. Navigate to the settings section and customize to your heart's content.

## 🐛 Debugging

Enable detailed logging to troubleshoot issues:

```bash
# Enable debug mode
DEBUG=1 go run cmd/cerebro/main.go
```

### Log Files

Monitor different aspects of cerebro in real-time:

```bash
# General application logs
tail -f debug.log

# HTTP request inspection
tail -f gemini_req_dump.log

# Additional HTTP logs
tail -f http_req_dump_gemini.log
```

## 🤝 Contributing

We welcome contributions! Here's how you can help:

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. **Commit** your changes (`git commit -m 'Add amazing feature'`)
4. **Push** to the branch (`git push origin feature/amazing-feature`)
5. **Open** a Pull Request

### Development Guidelines

- Write clear, concise commit messages
- Add tests for new features
- Update documentation as needed
- Follow Go conventions and best practices

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) 🫧
- Inspired by the X-Men's Cerebro (but for AI, not mutants)

## Upcoming Features
### 🤖 Multi-LLM Support

**OpenAI GPT - GPT-3.5, GPT-4 integration**
**Grok series**
**Claude-3, other models**

### ⚙️ Enhanced Configuration

**Parameter Control Panel - Temperature, top-p, max tokens**
**Model Selection UI - Switch between different LLMs seamlessly**
**Custom Prompts - Save and manage prompt templates**
**Theme Customization - Multiple color schemes and layouts**

### 📚 Chat Management

**Chat History Viewer - Browse previous conversations**
**Session Management - Save, load, and organize chat sessions**
**History Search - Find specific conversations quickly**
**Export Options - Save chats as markdown, JSON, or plain text**
**Bulk Operations - Delete multiple sessions at once**

### 🔧 Advanced Features

**Streaming Responses - Real-time response display**
**Multi-turn Conversations - Context-aware chat sessions**
**File Upload Support - Analyze documents and images**
**Plugin System - Extensible architecture for custom features**


## 🔗 Links

- [Issues](https://github.com/spitfireGG/cerebro/issues) - Report bugs or request features
- [Discussions](https://github.com/spitfireGG/cerebro/discussions) - Community chat
- [Wiki](https://github.com/spitfireGG/cerebro/wiki) - Extended documentation

---

<div align="center">

**Happy Hacking! 🎉**

*If you like the work, star the repo to keep my motivation going*

</div>
