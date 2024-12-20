# Password manager CLI

## Overview

This CLI tool is designed to manage authentication, secrets, groups, and sharing functionalities. It is built using Go and the Cobra library.

## Command Structure

```
dv
├── auth
│   ├── login
│   ├── register
│   └── logout
├── secrets
│   ├── create
│   ├── get
│   ├── update
│   ├── delete
│   └── list
├── group
│   ├── create
│   ├── get
│   ├── update
│   ├── delete
│   ├── list
│   ├── adduser
│   └── removeuser
└── share
    ├── togroup
    └── touser
```

## Command Details

### Auth Commands

#### login

- Short description: Log in to the system
- Long description: Authenticate a user with their email and password. If no arguments are provided, an interactive TUI form will be presented for input.
- Flags:
  - `--email` (required): User's email address
  - `--password` (required): User's password
- Example: `dv auth login --email user@example.com --password mysecurepassword`

#### register

- Short description: Register a new user
- Long description: Create a new user account with email and password. If no arguments are provided, an interactive TUI form will be presented for input.
- Flags:
  - `--email` (required): New user's email address
  - `--password` (required): New user's password
- Example: `dv auth register --email newuser@example.com --password mysecurepassword`

#### logout

- Short description: Log out from the system
- Long description: End the current user session and clear authentication tokens
- Example: `dv auth logout`

### Secrets Commands

#### create

- Short description: Create a new secret
- Long description: Add a new secret with a name, username and password . If no arguments are provided, an interactive TUI form will be presented for input.
- Flags:
  - `--name` (required): Name of the secret
  - `--username` (required): Username for the secret
  - `--password` (required): Password for the secret
- Example: `dv secrets create --name myapi --username apiuser --password apisecret `

#### get

- Short description: Retrieve a secret
- Long description: Fetch and display a secret by its ID
- Flags:
  - `--id` (required): ID of the secret to retrieve
- Example: `dv secrets get --id 12345`

#### update

- Short description: Update an existing secret
- Long description: Modify one or more fields of an existing secret
- Flags:
  - `--id` (required): ID of the secret to update
  - `--name`: New name for the secret
  - `--username`: New username for the secret
  - `--password`: New password for the secret
- Example: `dv secrets update --id 12345 --password newsecretpassword`

#### delete

- Short description: Delete a secret
- Long description: Remove a secret from the system by its ID
- Flags:
  - `--id` (required): ID of the secret to delete
- Example: `dv secrets delete --id 12345`

#### list

- Short description: List all secrets
- Long description: Display a list of all secrets accessible to the current user
- Flags:
  - `--limit`: Maximum number of secrets to display (default: 10)
  - `--offset`: Number of secrets to skip (for pagination)
- Example: `dv secrets list --limit 20 --offset 40`

### Group Commands

#### create

- Short description: Create a new group
- Long description: Create a new group with a specified name
- Flags:
  - `--name` (required): Name of the new group
- Example: `dv group create --name developers`

#### get

- Short description: Retrieve group information
- Long description: Display details of a group by its name, by default list secrets shared with the group.
  Use `--users` flag to get list of user present in the group
- Flags:
  - `--name` (required): Name of the group to retrieve
  - `--users`: Name of the group to retrieve
- Example: `dv group get --name developers`

#### update

- Short description: Update a group's name
- Long description: Change the name of an existing group
- Flags:
  - `--oldname` (required): Current name of the group
  - `--newname` (required): New name for the group
- Example: `dv group update --oldname developers --newname engineering`

#### delete

- Short description: Delete a group
- Long description: Remove a group from the system by its name
- Flags:
  - `--name` (required): Name of the group to delete
- Example: `dv group delete --name obsolete-group`

#### list

- Short description: List all groups
- Long description: Display a list of all groups accessible to the current user
- Example: `dv group list`

#### adduser

- Short description: Add a user to a group
- Long description: Add a user to an existing group using their email address
- Flags:
  - `--group` (required): Name of the group
  - `--email` (required): Email of the user to add
- Example: `dv group adduser --group developers --email newdev@example.com`

#### rmuser

- Short description: Remove a user from a group
- Long description: Remove a user from an existing group using their email address
- Flags:
  - `--group` (required): Name of the group
  - `--email` (required): Email of the user to remove
- Example: `dv group rmuser --group developers --email formerdev@example.com`

### Share Commands

#### togroup

- Short description: Share a secret with a group
- Long description: Grant access to a secret for all members of a specified group
- Flags:
  - `--secret-id` (required): ID of the secret to share
  - `--group` (required): Name of the group to share with
- Example: `dv share togroup --secret-id 12345 --group developers`

#### touser

- Short description: Share a secret with a user
- Long description: Grant access to a secret for a specific user by their email address
- Flags:
  - `--secret-id` (required): ID of the secret to share
  - `--email` (required): Email of the user to share with
- Example: `dv share touser --secret-id 12345 --email collaborator@example.com`
