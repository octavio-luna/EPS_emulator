# EPS emulator

This is a simple emulator for the EPS (Electric Power System) of the CubeSat. It is written in Go and uses two named pipes to communicate with the EPS. The emulator is able to simulate the EPS behavior and to send responses.

## Usage

To use the emulator, it's recommended to have Go installed. Then, you can run the emulator with the following command:

```
go build -o eps_emulator main.go
./eps_emulator
```

### Fifos

The emulator uses two named pipes to communicate with the EPS. The pipes are created in the `/tmp` folder and are named `eps_write_fifo` and `eps_read_fifo`. The EPS will write to the `eps_write_fifo` and read from the `eps_read_fifo`.

### Binary

The binary is available in the `bin` folder. The binary is available for Linux and MacOS. If you want to use the binary, you can run the following command:

```
./bin/{OS}_{ARCH}/eps_emulator
```

Where `{OS}` is the operating system and `{ARCH}` is the architecture. For example, for Linux and AMD64, you can run the following command:

```
./bin/linux_amd64/eps_emulator
```


## Note

If go is not installed, you can use the pre-built binary in the `bin` folder. Also, if the binary for your OS is not available, you can build it using the build.sh script.   


### Disclaimer

The build script is only available for Linux and MacOS, it asumes bash is installed and the go compiler is available in the PATH. It'll try to install golang if it's not available, create the pipes in the /tmp folder and it'll try to build the binary for the current OS.