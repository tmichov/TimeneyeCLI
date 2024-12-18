# TimeneyeCLI

TimeneyeCLI is a command-line tool to interact with the Timeneye API. It allows you to authenticate, manage projects, create entities, and retrieve information from your Timeneye account.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/tmichov/TimeneyeCLI.git
   ```

2. Navigate to the directory:
   ```bash
   cd TimeneyeCLI
   ```

3. Build the binary:
   ```bash
   go build -o timeneye
   ```

4. Move the binary to your PATH:
   ```bash
   mv timeneye /usr/local/bin/
   ```

## Usage

Run the `timeneye` command followed by a specific subcommand and its arguments.

```bash
timeneye [command] [arguments]
```

### Commands

#### **auth**
Authenticate to the Timeneye API.

**Usage:**
```bash
timeneye auth -t <token>
```

**Arguments:**
- `-t, --token` (required): Token to authenticate.

---

#### **projects**
List all projects.

**Usage:**
```bash
timeneye projects
```

---

#### **create**
Create a new entity in Timeneye.

**Usage:**
```bash
timeneye create [options]
```

**Options:**
- `-t, --type` (required): Type of entity to create.
- `-p, --project`: Project associated with the entity.
- `-d, --date`: Date for the entity (format: YYYY-MM-DD).
- `-n, --name`: Name of the entity.
- `-l, --duration`: Duration of the entity.
- `-D, --description`: Description of the entity.

**Example:**
```bash
timeneye create -t task -p 1234 -d 2024-12-18 -n "Meeting" -l 60 -D "Weekly status meeting"
```

---

#### **help**
Display help information for available commands.

**Usage:**
```bash
timeneye help
```

---

#### **version**
Retrieve the current version of the CLI.

**Usage:**
```bash
timeneye version
```

---

## Contributing

Feel free to fork the repository, open issues, and submit pull requests to improve TimeneyeCLI.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

