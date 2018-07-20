Kasa CLI
===========================
An unoffical tool to communicate with the remote Kasa API and control your smart devices

## Usage

    Usage:
      kasa [flags]
      kasa [command]

    Available Commands:
      help        Help about any command
      list        list devices in kasa
      off         turn off a device
      on          turn on a device
      toggle      toggle device state

    Flags:
      -h, --help              help for kasa
      -p, --password string   your Kasa password (defaults to KASA_PASSWORD environment variable)
      -u, --username string   your Kasa username (defaults to KASA_USERNAME environment variable)
          --version           version for kasa

    Use "kasa [command] --help" for more information about a command.

## Roadmap
There's several commands available. As demanded, incorporate [these commands](https://github.com/softScheck/tplink-smartplug/blob/master/tplink-smarthome-commands.txt) into the CLI.
