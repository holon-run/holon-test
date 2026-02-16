# holon-test

A test repository for [Holon](https://github.com/holon-run/holon) - an autonomous AI agent system for software development.

## About

This repository serves as a testing ground for Holon, an autonomous agent system designed to assist with software engineering tasks. It is used to validate and demonstrate Holon's capabilities in handling GitHub issues, pull requests, and code changes through autonomous AI agents.

## Project Setup

### Prerequisites

- A GitHub account
- Access to the Holon system
- Basic understanding of Git and GitHub workflows

### Repository Structure

```
.
├── .github/
│   └── workflows/
│       └── holon-trigger.yml    # GitHub Actions workflow for Holon integration
├── .gitignore                    # Git ignore patterns for Go projects
├── LICENSE                       # Apache License 2.0
└── README.md                     # This file
```

### How Holon Works

This repository is configured with a GitHub Actions workflow that triggers Holon agents in response to:

- **Issue comments**: When someone comments `@holonbot please implement...`
- **Issue labels/assignments**: When issues are labeled or assigned
- **Pull request labels**: When PRs are labeled

The workflow (`.github/workflows/holon-trigger.yml`) calls the Holon solve action, which processes the request and generates code changes, tests, or documentation updates.

## Usage

### Triggering Holon

To trigger Holon to work on an issue or PR:

1. **Via Comment**: Add a comment on an issue or PR:
   ```
   @holonbot please implement this issue: [task description]
   ```

2. **Via Label**: Apply a label that triggers the Holon workflow

3. **Via Assignment**: Assign the issue to the Holon bot

### Example Workflow

1. Create a new GitHub issue with a clear description of the task
2. Add a comment: `@holonbot please implement this issue: [your task]`
3. Holon will:
   - Analyze the issue and existing code
   - Create an implementation plan
   - Make necessary code changes
   - Run tests and verify changes
   - Create a summary of the work done

### Viewing Results

Holon generates output in the form of:
- Modified source files
- Test results
- Documentation updates
- A summary of changes

## Development

### Local Development

Since this is a test repository, local development primarily involves:

1. Forking and cloning the repository
2. Making changes as needed for testing
3. Submitting pull requests

### Testing Holon Functionality

To test Holon's capabilities:

1. Create an issue with a specific task
2. Use the trigger phrase to invoke Holon
3. Monitor the GitHub Actions logs
4. Review the generated changes

## Contributing

We welcome contributions! Here's how you can help:

### Getting Started

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add amazing feature'`)
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

### Contribution Guidelines

- **Clear Communication**: Provide clear descriptions in issues and PRs
- **Testing**: Ensure any changes are well-tested
- **Documentation**: Update documentation as needed
- **Respect**: Be respectful to all contributors
- **Holon Integration**: Feel free to test Holon by using `@holonbot` commands

### Using Holon for Contributions

You can also use Holon to help with contributions:

- Use `@holonbot please implement...` to have Holon help with feature implementation
- Use `@holonbot` for code review suggestions
- Provide clear requirements when requesting Holon's assistance

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Support

For questions about:
- **Holon system**: Visit [holon-run/holon](https://github.com/holon-run/holon)
- **This repository**: Open a GitHub issue

## Related Projects

- [Holon](https://github.com/holon-run/holon) - The main Holon autonomous agent system
- [Holon Documentation](https://github.com/holon-run/holon) - Documentation and guides

---

**Note**: This is a test repository used for validating and demonstrating Holon's autonomous software development capabilities.
