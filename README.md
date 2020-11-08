# tenki 天気

> 天気 (tenki), Japanese for _weather conditions_

tenki is a command-line client for [wttr.in](https://wttr.in/) written in Go, using the [Cobra](https://github.com/spf13/cobra) library.

## Usage Examples

```shell
$ ./tenki --help
Tenki is a CLI for querying weather information from wttr.in/ endpoints.

Users can define their location based on various query parameters including
location, nearest point-of-interest, nearest airport, or public domain name.

Usage:
  tenki [command]

Available Commands:
  get         Query weather information
  help        Help about any command

Flags:
  -h, --help              help for tenki
      --language string   Language to generate output in (default "en")

Use "tenki [command] --help" for more information about a command.
```

```shell
$ ./tenki get --help
Query weather information either for your default location or for a specific one.

Supported location types:

    /paris                  # city name
    /~Eiffel+tower          # any location (+ for spaces)
    /Москва                 # Unicode name of any location in any language
    /muc                    # airport code (3 letters)
    /@stackoverflow.com     # domain name
    /94107                  # area codes
    /-78.46,106.79          # GPS coordinates

Usage:
  tenki get [--location location] [--png path/to/png] [flags]

Flags:
  -h, --help                 help for get
      --location string      Query location
      --path-to-png string   Location to store PNG output

Global Flags:
      --language string   Language to generate output in (default "en")
```

```shell
./tenki get --location Kyoto --language fr
Prévisions météo pour: Kyoto

   _`/"".-.     Light rain shower, mist
    ,\_(   ).   18 °C          
     /(___(__)  ↓ 0 km/h       
       ‘ ‘ ‘ ‘  5 km           
      ‘ ‘ ‘ ‘   0.1 mm         
                                                       ┌─────────────┐                                                       
┌──────────────────────────────┬───────────────────────┤ sam. 07 nov.├───────────────────────┬──────────────────────────────┐
│             Matin            │          Après-midi   └──────┬──────┘       Soir            │             Nuit             │
├──────────────────────────────┼──────────────────────────────┼──────────────────────────────┼──────────────────────────────┤
│               Couvert        │  _`/"".-.     Pluies éparses │      .-.      Pluie légère   │      .-.      Pluie légère   │
│      .--.     16 °C          │   ,\_(   ).   16 °C          │     (   ).    15 °C          │     (   ).    9 °C           │
│   .-(    ).   ← 4-5 km/h     │    /(___(__)  ← 5-6 km/h     │    (___(__)   ↙ 3-4 km/h     │    (___(__)   ↗ 0 km/h       │
│  (___.__)__)  10 km          │      ‘ ‘ ‘ ‘  10 km          │     ‘ ‘ ‘ ‘   9 km           │     ‘ ‘ ‘ ‘   9 km           │
│               0.0 mm | 0%    │     ‘ ‘ ‘ ‘   0.1 mm | 63%   │    ‘ ‘ ‘ ‘    1.6 mm | 96%   │    ‘ ‘ ‘ ‘    2.0 mm | 77%   │
└──────────────────────────────┴──────────────────────────────┴──────────────────────────────┴──────────────────────────────┘
                                                       ┌─────────────┐                                                       
┌──────────────────────────────┬───────────────────────┤ dim. 08 nov.├───────────────────────┬──────────────────────────────┐
│             Matin            │          Après-midi   └──────┬──────┘       Soir            │             Nuit             │
├──────────────────────────────┼──────────────────────────────┼──────────────────────────────┼──────────────────────────────┤
│  _`/"".-.     Pluies éparses │    \  /       Partiellement …│    \  /       Partiellement …│     \   /     Temps clair    │
│   ,\_(   ).   16 °C          │  _ /"".-.     18 °C          │  _ /"".-.     10..12 °C      │      .-.      8..10 °C       │
│    /(___(__)  ↘ 15-18 km/h   │    \_(   ).   ↘ 20-23 km/h   │    \_(   ).   ↘ 15-19 km/h   │   ― (   ) ―   ↘ 16-21 km/h   │
│      ‘ ‘ ‘ ‘  9 km           │    /(___(__)  10 km          │    /(___(__)  10 km          │      `-’      10 km          │
│     ‘ ‘ ‘ ‘   0.4 mm | 62%   │               0.0 mm | 0%    │               0.0 mm | 0%    │     /   \     0.0 mm | 0%    │
└──────────────────────────────┴──────────────────────────────┴──────────────────────────────┴──────────────────────────────┘
                                                       ┌─────────────┐                                                       
┌──────────────────────────────┬───────────────────────┤ lun. 09 nov.├───────────────────────┬──────────────────────────────┐
│             Matin            │          Après-midi   └──────┬──────┘       Soir            │             Nuit             │
├──────────────────────────────┼──────────────────────────────┼──────────────────────────────┼──────────────────────────────┤
│    \  /       Partiellement …│    \  /       Partiellement …│    \  /       Partiellement …│    \  /       Partiellement …│
│  _ /"".-.     10..11 °C      │  _ /"".-.     13..14 °C      │  _ /"".-.     8..10 °C       │  _ /"".-.     7..9 °C        │
│    \_(   ).   → 9-10 km/h    │    \_(   ).   → 13-15 km/h   │    \_(   ).   ↘ 13-17 km/h   │    \_(   ).   ↘ 10-14 km/h   │
│    /(___(__)  10 km          │    /(___(__)  10 km          │    /(___(__)  10 km          │    /(___(__)  10 km          │
│               0.0 mm | 0%    │               0.0 mm | 0%    │               0.0 mm | 0%    │               0.0 mm | 0%    │
└──────────────────────────────┴──────────────────────────────┴──────────────────────────────┴──────────────────────────────┘
Emplacement: 京都市, 京都府, 日本 [35.0185804,135.763835]
```

tenki also lets you store colored output as PNG, using the `--path-to-png` flag. It can be combined with any language or location specification, e.g.

```shell
$ ./tenki get --location ~Tokyo+Skytree --language de --path-to-png ./examples/TokyoSkytree.png
```

produces the following image:

![Weather at Tokyo Skytree in German](./examples/TokyoSkytree.png)