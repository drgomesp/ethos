# 🛠 Ethos
> An agency-increasing toolkit for Ethereum developers.

## Features

- Managed `geth` node capable of connecting to 
  most networks available
- Usage of Ethereum `devtools` in an easy way  

## Dependencies

To use Ethos, the following tools are **required**:

1. ### [Task][1]

If you have Go installed, installing Task is as easy as running: 
```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

2. ### [Solidity Compiler][2]

For now, to build and compile Solidity source code we rely on
an external compiler, which should ideally be either **solc** or **solcjs**.

For the purpose of these instructions, the choice will be **solcjs**, which
can be easily installed with `npm` by running: 

```bash
npm install -g solc
```

---

[1]: https://docs.soliditylang.org/en/v0.8.13/installing-solidity.html
[2]: (https://docs.soliditylang.org/en/v0.8.13/installing-solidity.html)