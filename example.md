# A New Approach to Managing Environment Variables with Keeper

For most projects, environment variables are handled through `.env` files that are added to `.gitignore` to avoid accidental commits. However, studies have shown that sensitive information often finds its way into repositories despite these safeguards. For example, a 2019 report by North Carolina State University revealed that nearly 100,000 GitHub repositories contained accidentally committed secrets like API keys and database credentials. This can lead to serious security vulnerabilities, especially if credentials are not rotated or are inadvertently exposed in public repositories.

Keeper offers a new way to manage environment variables securely without the risk of them ending up in your version control. Here's how you can set up your project to securely handle sensitive data, streamline team onboarding, and simplify deployment workflows.

## Setting Up a Project with Keeper

### 1. Create a Group for Your Project's Secrets

First, use Keeper to create a group for your project. For example, if you're building an application named "MyApp," you could create a group called "myapp-devs" specifically for managing project secrets:

```bash
keeper group create --name myapp-devs
```

### 2. Store Project Secrets in Keeper

Next, store essential credentials as secrets and share them with the "myapp-devs" group. For example, you might want to store the database username and password, and an API key:

```bash
keeper secrets create --name db_user --username "my_db_user" --password "my_db_password"
keeper secrets create --name api_key --username "apiuser" --password "my_secure_api_key"
keeper share togroup --secret-id 40 --group myapp-devs
keeper share togroup --secret-id 42 --group myapp-devs
```

### 3. Load Secrets as Environment Variables

Instead of storing credentials directly in a `.env` file, you can add them to an `.envrc` file with dynamic commands that retrieve values from Keeper:

```bash
export DB_USER="$(keeper secrets get --id 40)"
export DB_PASSWORD="$(keeper secrets get --id 41)"
export API_KEY="$(keeper secrets get --id 42)"
```

By including these commands in `.envrc`, you ensure sensitive credentials are never stored in plaintext on disk. Each time you start a new session or deploy the project, you can simply load these environment variables into the session by running:

```bash
source .envrc
```

### 4. Automate Environment Variable Setup for Team Members

One of the key benefits of Keeper is that it simplifies adding new team members to your project. Once added to the "myapp-devs" group, a new team member can access all the project's environment variables without any manual configuration. They simply need to:

* Clone the project repository
* Log in to Keeper:

```bash
keeper auth login --email newdev@example.com --password developerpassword
```

* Run `source .envrc` to load all necessary environment variables into their session.

This way, every developer on the team has access to the correct configuration without risking accidental exposure of secrets.

### 5. Seamless Deployment to Remote Environments

Deploying to remote environments becomes equally simple. You can automate the setup of Keeper in deployment scripts, so credentials are only fetched when needed. For instance, a deployment script could contain:

```bash
export EMAIL="deploybot@example.com"
export PASSWORD="securepassword"
keeper auth login --email "$EMAIL" --password "$PASSWORD"
source .envrc
```

This approach also facilitates automation across CI/CD pipelines, where securely managing sensitive information is critical.

### 6. Easy Credentials Rotation and Updates

Since Keeper synchronizes with the cloud, credentials can be rotated or updated easily from the CLI, TUI, or even a web interface. When a credential like an API key is rotated, the updated value is instantly accessible to all team members and deployment environments using Keeper. There's no need for manual updates to `.env` files, and team members won't be stuck using outdated credentials.

## Benefits of Using Keeper for Environment Variable Management

Keeper offers numerous advantages for managing environment variables:

* **Enhanced Security**: Secrets are never stored in `.env` files, reducing the risk of exposure.
* **Smooth Onboarding**: New team members simply log in to Keeper to access project secrets, avoiding complex setup steps.
* **Automated Deployment**: Deployment scripts with Keeper ensure credentials are securely fetched during deployment without storing them on disk.
* **Efficient Credential Rotation**: All instances of credentials across teams and environments are updated instantly when changed in Keeper.

By integrating Keeper, teams can manage secrets effectively and prevent the common pitfalls of managing sensitive data in `.env` files.
