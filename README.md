# Troy

The EVM foot soldier

## Usage

#### Scan an existing contract on a given network:

```
$ troy -a <CONTRACT_ADDRESS> -k <INFURA_ID> -n <NETWORK>
```

Example:

```
$ troy -a 0x6b175474e89094c44da98b954eedeac495271d0f -k <INFURA_ID> -n mainnet
```

#### Scan bytecode

```
$ troy -c <BYTE_CODE>
```

Example:

```
$ troy -c 60ff60ff
```

## Notes

Ideas of some things we can scan for initially

- Reentrency
  - detect CALL and DELEGATECALL before sstore
- Highlight notable opcodes CALL DELEGATECALL CHAINID
- Highlight notable function calls "approve", "mint", "transfer"
  - detect if they have a CALLER comparison with EQ and JUMPI
