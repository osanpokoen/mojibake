# mojibake

## Usage

### Install

```console
$ go install github.com/kakitamama/mojibake
```

### Encode

```console
$ echo なにがですか | mojibake
縺ｪ縺ｫ縺後〒縺吶°
$ echo いいんすかそれで | mojibake
縺�＞繧薙☆縺九◎繧後〒
```

### Decode

If the garbled text in its complete form, it can be decoded.

```console
$ echo 縺ｪ縺ｫ縺後〒縺吶° | mojibake -d
なにがですか
```

If the garbled text is incomplete as shown below, the decoded will fail.

```console
$ echo 縺�＞繧薙☆縺九◎繧後〒 | mojibake -d
encoding: rune not supported by encoding.
```

### Docker

If this is too much work, you can use Docker to run it.  
Docker Hub Repository: [kakitamama/mojibake](https://hub.docker.com/r/kakitamama/mojibake)

```console
$ echo なにがですか | docker run --rm -i kakitamama/mojibake
縺ｪ縺ｫ縺後〒縺吶°
$ echo なにがですか | docker run --rm -i kakitamama/mojibake | docker run --rm -i kakitamama/mojibake -d
なにがですか
```