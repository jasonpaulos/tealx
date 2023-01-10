# Tealx

Tealx is an an intermediate-level smart contract language that compiles to Algorand's TEAL assembly.

For example, this Tealx code:

```xml
<program version="8">
    <main>
        <log>
            <bytes value="hello world" format="utf-8" />
        </log>
        <program-return>
            <int value="1" />
        </program-return>
    </main>
</program>

```

Compiles to this TEAL assembly:

```
#pragma version 8
byte 0x68656c6c6f20776f726c64
log
int 1
return
```

Check out [examples/example.xml](examples/example.xml) for a larger example.

This project was completed as part of Algorand's internal January 2023 hackathon. It serves as a proof of concept, but it is not production ready.

## Benefits

Tealx is not meant to be easy for developers to write (though it does try to be easy to understand). Instead, its goal is to be easy for other compilers to target.

As a result, an intermediate-level language like Tealx has the following benefits:

* Easy for other compilers (like [PyTeal](https://github.com/algorand/pyteal)) to target. By using Tealx as a target, other compilers can focus more on being easy to use and improving their frontend. As a bonus, if multiple other compilers target Tealx, improvements to Tealx can be made once but benefit multiple compilers.

* Easy to optimize. Inspiration was taken from [Ethereum's Yul](https://docs.soliditylang.org/en/v0.8.17/yul.html), which is a smart contract language that has an important property: optimization passes can be made to Yul programs directly without converting them to a lower level. Since optimization passes output Yul programs, multiple optimizations can be chained or repeated with ease. While no optimizations are implemented at the moment, Tealx lays the foundation for this same approach.

* (Potentially) being able to output to environments other than TEAL. Tealx programs are high-level enough to be potentially compiled to other assembly languages.

## Usage

To compile Tealx locally, follow these steps:

1. Clone this repo (requires Go 1.17).
2. Run `make build`. This will produce a CLI binary in `bin/tealx`
3. Use the `tealx` binary to compile Tealx programs to TEAL. For example:

```bash
$ ./bin/tealx -i examples/example.xml -o examples/example.teal
Successfully compiled examples/example.xml to examples/example.teal
```

Also, check out this block-based frontend for writing Algorand Smart Contracts, which can produce Tealx programs: https://github.com/algochoi/blockly
