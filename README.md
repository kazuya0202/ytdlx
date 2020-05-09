# ytdlx - youtube-dl wrapper command

## Features

+ Supported multi URL(ID) and file that is enumerated URL(ID).
+ The command easier to use than using pure youtube-dl.

## Requirements

+ youtube-dl

  ```sh
  $ scoop install youtube-dl  # using Scoop
  $ choco install youtube-dl  # using Chocolatey
  $ pip install youtube-dl  # using pip
  ```

+ FFmpeg

  ```sh
  $ scoop install ffmpeg  # using Scoop
  $ choco install ffmpeg  # using Chocolatey
  ```

## Installation

Do one or the other.

+ Download binary from [Release](https://github.com/kazuya0202/ytdlx/releases) (**Recommend**)

+ Execute command `$ go get`

  ```sh
  $ go get github.com/kazuya0202/ytdlx
  ```

## Usage

## 1. Download option

> `URL` == `URL | ID | FILE`

### 1-1. Default (== Best)

Downloads default(best) format.

```sh
$ ytdlx [URL]
```

### 1-2. Audio (only)

Downloads only audio format.  
The saved file don't include video.

```sh
$ ytdlx -a [URL]
$ ytdlx --audio [URL]
```

### 1-3. Video (only)

Downloads only video format.  
The saved file don't include audio.

```sh
$ ytdlx -v [URL]
$ ytdlx --video [URL]
```

### 1-4. Select

Downloads format that user selected.

**Note:** When user specify multi URL, every URL is executed selected format.

```sh
$ ytdlx -s [URL]
$ ytdlx --select [URL]
```



### 1-5. Select each

Downloads format that user selected each URL.

**Note:** When user specify multi URL, need to select format each URL.

```sh
$ ytdlx -S [URL]
$ ytdlx --select-each [URL]
```

### 1-6. Find

Downloads format that user selected from `format-list`.

> `format-list`: Get list by `$ youtube-dl -F [URL]`

```sh
$ ytdlx -f [URL]
$ ytdlx --find [URL]
```

<br>

## 2. Options

### 2-1. Help

Displays help message of `ytdlx`.

```sh
$ ytdlx -h
$ ytdlx --help
```

### 2-2. Output name

Specify output name that save file.

**Note:** When user specify multi URL, the name is `[output_name]_00X` (`00X`: 3 digits is filled 0.)

```sh
$ ytdlx --[Any_Option] [URL] -o [Output_name]
$ ytdlx --[Any_Option] [URL] --output [Output_name]
```

### 2-3. Get format list

The option is same as `$ youtube-dl -F`.

```sh
$ ytdlx -F [URL]
$ ytdlx --format-list [URL]
```

