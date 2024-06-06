# 2clip

2clip is a CLI tool for managing your clipboard, that automatically adds the value of your created keys. Fast, simple and always available.

## Summary
* [Installation](#installation)
* [How to use](#how-to-use)
* [Technologies](#technologies)

<h2 id="installation">Installation</h2>

#### Linux and MacOS
   ```sh
   curl -O https://raw.githubusercontent.com/Paulooo0/2clip/main/install_2clip.sh
   chmod +x install.sh
   ./install.sh
   ```

#### Windows
* Coming soon

<h2 id="how-to-use">How to use</h2>

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

<h2 id="technologies">Technologies</h2>

* Golang
* BoltDB
* CobraCLI
