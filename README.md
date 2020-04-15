# ytdlx - youtube-dl wrapper command

## Features

Supported multi URL(ID) and file that is enumerated URL(ID).

## Installation

Do one or the other.

+ Download binary from [Release](https://github.com/kazuya0202/ytdlx/releases) (**Recommend**)

+ Execute command `$ go get`

  ```sh
  $ go get github.com/kazuya0202/ytdlx
  ```

## Usage (Download format)

> `URL` == `URL | ID | FILE`

### 1. Default

```sh
$ ytdlx [URL]
```

### 2. Audio

```sh
$ ytdlx -a [URL]
$ ytdlx --audio [URL]
```

### 3. Video

```sh
$ ytdlx -v [URL]
$ ytdlx --video [URL]
```

### ...