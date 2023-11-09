# servercheck
A small program for comprehensive evaluation and verification of server systems

## Features
- Automatically add functions

## Quick start
```sh
make build
./_output/bin/servercheck
```
output
```text
register checker: CPUArch
register checker: CPUCore
check:  CPUArch
check:  CPUCore
| Ability | Details | Result | Passed | Recommendation |
| --- | --- | --- | --- | --- |
| CPUArch | check CPU arch | [arch] actual: [amd64], expect: [amd64 arm64] | pass |  |
| CPUCore | check CPU core | [number of cores] acutal: 8, expect: 4 | pass |  |
```

### add new memory checker
```sh
make gen_checker name=Memory
```

## todo
- Connect with AI to achieve true automation
