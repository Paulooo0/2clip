# 2clip

2clip is a CLI tool for managing your clipboard, that automatically adds the value of your created keys. Fast, simple and always available.

### Installation

#### Unix

Run the following command, to make the script executable:
```bash
chmod +x install_2clip.sh
```

Run the bash script:
```bash
./install_2clip.sh
```

#### Windows

Open PowerShell as an `administrator`.

Run the PowerShell script:
```shell
.\install_2clip.ps1
```

## How to use
### Main commands
#### add - adds a new key-value pair to the database
```bash
2clip add "My key" "My value"
```
* use `add -p` to add protected keys, that can only be accessed using authentication
#### get - adds to your clipboard the value of your provided key
```bash
2clip get "My key"
```
Output:
```bash
My value
Value copied to clipboard
```
#### list - lists all keys alphabetically sorted
```bash
2clip list
```
Output:
```bash
A
akey

B
bkey
```

### Other commands
#### auth - allows you to create a password
```bash
2clip auth
```
* use `auth -u` to update your actual password
#### remove - removes a key-pair value from the database
```bash
2clip remove "My key"
```

## Technologies

* Golang
* BoltDB
* CobraCLI
