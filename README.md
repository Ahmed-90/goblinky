GoBlinky
-

Small server to manage blinking patterns for anything linked to a Raspberry Pi GPIO pins

### Features

- Custom blinking patterns.
- Seconds and milliseconds blinking intervals.
- Small and lightweight webserver written in GO.
- GPIO control is based on the `wiringpi` library, so its compatible with all models.
- Control via simple `GET` requests API.

### Instillation

- Install `wiringpi` via your package manager `apt-get install
   wiringpi`.
- Download latest version from releases page above.

```bash
tar xvf goblinky-1.0.0.tar.gz
cd goblinky-1.0.0
mv goblinky /usr/local/bin/
mv init.sh /etc/init.d/goblinky
update-rc.d goblinky defaults
```

#### Options
In `init.sh` you can find the available options

|Description         |Option           |
|--------------------|----------------------|
|Executable location |`dir="/usr/local/bin"`|
|Executable name     |`cmd="goblinky"`|
|Server port         |`port="4500"`|
|System user         |`user=""`|



### Usage

All you need to do now is to make a GET requests from any local or remote application


#### Start blinking
**URL**: `/set/[pin number]/[time unit]/[pattern]`

Parameters:

 - **Pin Number:** GPIO pin number, following the [BCM Numbering](https://pinout.xyz/).
 - **Time unit:** the time unit to be used to execute the pattern, takes 2 options `s` for seconds and `ms` for milliseconds.
 - **Pattern:** this is the core of the functionality, you can create a custom pattern on the fly by sending a serialized list of intervals for example setting the `time unit` to `s` and pattern to `1,2` will mean blink on for 1 sec then off for 2 sec and repeat, technically the pattern can be as long as you wish like `1,4,1,5,8,1,2`. GoBlinky will walk though the pattern and repeat until its stopped.
*The first action in the pattern will always be `HIGH/ON/1/True` what ever you name it.*

Example, set the LED to blink 3 sec ON and 1 sec off:
```
http://127.0.0.1:4500/set/10/2/3,1
```
---
#### Stop blinking
**URL**: `/stop/[pin number]`

Pin number same as mentioned above.

Example:
```
http://127.0.0.1:4500/stop/10
```

