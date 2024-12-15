# GitHub User Activity

A simple command line interface to fetch the recent activity of a GitHub user and display it in the terminal.

## Requirements

The application should run from the command line, accept the GitHub username as an argument, fetch the user's recent activity using the GitHub API, and display it in the terminal. The user should be able to:

-   Provide the GitHub username as an argument when running the CLI

```bash
github-activity <username>
```

-   Fetch the recent activity of the specified GitHub user using the GitHub API.

```bash
https://api.github.com/users/<username>/events
```

-   Display the fetched activity in the terminal.

```bash
Output:
    - Pushed 3 commits
    - Opened a new issue
    - Starred repository
```
