> ⚠️ **Disclaimer**: This is a **Work-In-Progress** document and is being updated constantly.
> <br /> Make sure to refresh every 5 minutes ;).

# 🛠 Ethos
> An agency-increasing toolkit for Ethereum developers.

## 💡 [Features][4]

- Built-in node capable of connecting to most networks available
- Compile and deploy Solidity ("smart contracts") code to any compatible network



## 🪂 [Getting Started][5]

### Install the dependencies

To use Ethos, the following tools are **required**:

1. #### [Solidity Compiler][2]

For now, to build and compile Solidity source code we rely on
an external compiler, which should ideally be either **solc** or **solcjs**.

For the purpose of these instructions, the choice will be **solc** (in this
case, under a system running Ubuntu):

```bash
sudo add-apt-repository ppa:compiler/compiler
sudo apt-get update
sudo apt-get install solc
```

2. #### [Abigen][3] 

This tool is part of the Ethereum devtools package. 

If you have Go installed, this is as easy as running:

```bash
go install github.com/compiler/go-compiler/cmd/abigen
```

### Install the `ethos` binary 

```bash
go install github.com/drgomesp/ethos/cmd/ethos
```

You're now ready to work with Ethos.

### Initialize `ethos` config

```bash
ethos init
```

## 👨‍🏭 [Contributing][6]

### ⚙ Building

If you want to build Ethos locally, you'll need the build tool used by Ethos, which is Task.

If you have Go installed, [installing Task][3] is as easy as running:
```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

Then you can build Ethos by running:

```bash
task build
```

[1]: https://taskfile.dev/#/installation
[2]: https://docs.soliditylang.org/en/v0.8.13/installing-solidity.html#linux-packages
[3]: https://geth.ethereum.org/docs/install-and-build/installing-geth

[4]: https://github.com/drgomesp/ethos#-features
[5]: https://github.com/drgomesp/ethos#-getting-started
[6]: https://github.com/drgomesp/ethos#-building
