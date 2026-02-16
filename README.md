# holon-test

A test repository for [Holon](https://github.com/holon-run/holon) - an autonomous AI agent system for software development workflows.

## Overview

This repository serves as a testing ground for Holon's autonomous issue resolution and pull request generation capabilities. It is used to validate and demonstrate the Holon system's ability to:

- Automatically analyze GitHub issues
- Implement code changes based on issue requirements
- Generate pull requests with proper formatting
- Test autonomous development workflows

## Purpose

The `holon-test` repository is designed to:

1. **Test Holon Integration**: Validate that Holon triggers work correctly via GitHub Actions workflows
2. **Demonstrate Autonomy**: Showcase autonomous issue-to-PR workflows
3. **Debug and Improve**: Provide a controlled environment for testing and improving Holon's capabilities

## Installation

Since this is a test repository, there are no traditional dependencies to install. However, to use Holon in your own repository:

### Using GitHub Actions (Recommended)

1. Create a new GitHub repository
2. Add a Holon trigger workflow in `.github/workflows/`:

```yaml
name: Holon Trigger

on:
  issue_comment:
    types: [created]
  issues:
    types: [labeled, assigned]
  pull_request:
    types: [labeled]

permissions:
  contents: write
  issues: write
  pull-requests: write

jobs:
  holon:
    name: Run Holon
    uses: holon-run/holon/.github/workflows/holon-solve.yml@main
    secrets: inherit
```

3. Configure your Holon bot credentials in repository secrets
4. Trigger Holon by commenting `@holonbot please implement: [description]` on an issue

## Usage

### Triggering Holon

To trigger Holon to work on an issue:

1. Create a new GitHub issue with clear requirements
2. Add a comment: `@holonbot please implement: [task description]`
3. Holon will automatically:
   - Analyze the issue and comments
   - Explore the codebase
   - Implement the required changes
   - Create a pull request

### Example Issue Format

```markdown
## Feature: Add User Authentication

### Description
Implement user authentication using JWT tokens.

### Acceptance Criteria
- [ ] Add login endpoint
- [ ] Add JWT token validation
- [ ] Include unit tests
- [ ] Update documentation
```

Then trigger with:
```
@holonbot please implement: Add user authentication with JWT tokens
```

### Monitoring Progress

Holon provides real-time updates via GitHub Actions:
- Check the Actions tab for workflow runs
- Review logs for detailed progress
- Pull requests are created automatically upon completion

## Project Structure

```
holon-test/
├── .github/
│   └── workflows/
│       └── holon-trigger.yml    # Holon GitHub Actions workflow
├── .gitignore                   # Git ignore rules
├── LICENSE                      # MIT License
└── README.md                    # This file
```

## Contributing

This is a test repository used primarily for Holon development and testing. However, contributions are welcome!

### For Developers

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

### For Testing Holon

To test Holon's autonomous capabilities:

1. Create a new issue with a clear task description
2. Include specific acceptance criteria
3. Trigger Holon with `@holonbot please implement: [task]`
4. Monitor the GitHub Actions workflow
5. Review the generated pull request

### Contribution Guidelines

- **Be Specific**: Issues should have clear, actionable requirements
- **Include Context**: Provide relevant background and constraints
- **Define Success**: List acceptance criteria when possible
- **Test Thoroughly**: Verify Holon's output and provide feedback

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Related Projects

- [Holon](https://github.com/holon-run/holon) - The main Holon repository
- [holon-solve](https://github.com/holon-run/holon) - Holon's solve workflow

## Support

For questions, issues, or feedback related to Holon:

- Visit the [Holon GitHub repository](https://github.com/holon-run/holon)
- Check existing [GitHub issues](https://github.com/holon-run/holon/issues)
- Review the [Holon documentation](https://github.com/holon-run/holon#readme)

---

**Note**: This repository is maintained as part of the Holon project testing infrastructure. For production use, please refer to the main Holon repository documentation.
