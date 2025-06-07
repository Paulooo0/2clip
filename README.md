# 2clip

2clip is a CLI tool for managing your clipboard, that automatically adds the value of your created keys. Fast, simple and always available.

## Summary
* [Installation](#installation)
* [How to use](#how-to-use)
* [Use cases](#use-cases)
* [Technologies](#technologies)

<h2 id="installation">Installation</h2>

#### Linux and MacOS
   ```sh
   curl -O https://raw.githubusercontent.com/Paulooo0/2clip/main/script/install_2clip.sh
   chmod +x install_2clip.sh
   ./install_2clip.sh
   ```

#### Windows
* Coming soon

<h2 id="how-to-use">How to use</h2>

### Commands
**Add** - adds a new key-value pair to the database
```bash
2clip add "My key"
```
* Shorthand: `a`
* Arguments:
  * Extended: `-e` or `--extended`. Activates multiline input mode
  * Protected: `-p` or `protected`. Protects throught authentication to get the value of the key
* You can use `'` or `"` to add values with spaces. The terminal will enter in input mode to receive your key value, and then store it in database.

**Get** - adds to your clipboard the value of your provided key
```bash
2clip get "My key"
```
* Shorthand: `g`
* Arguments:
  * Index: `-i` or `--index`. Searches by index in `list command`, example: `2clip g -i 5`. Will get the value of 5th element in `2clip ls`
* If the key is `protected`, then `get` command will require your password to access the value

**List** - lists all keys alphabetically sorted, and its own index
```bash
2clip list

➜  A
1 a-key

➜  B
2 b-key

➜  C
3 c-key 🔒
```
* Shorthand: `ls`

**Auth** - manages your authentication
```bash
2clip auth
```
* Arguments:
  * Update: `-u` or `--update` updates your current password

**Remove** - removes a key-pair value from the database
```bash
2clip remove "My key"
```
* Shorthand: `rm`
* Arguments:
  * Index: `-i` or `--index`. Searches by index in `list command`, example: `2clip rm -i 5`. Will remove the 5th element in `2clip ls`

<h2 id="use-cases">Use cases</h2>

* Store passwords securely
* Use as flash cards for studies
* Useful for store both important and disposable informations, since is easy and quick to add or remove then

<h2 id="technologies">Technologies</h2>

* Golang
* BoltDB
* CobraCLI
