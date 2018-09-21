gomeboycolor
============================
**This project is a work in progress and is no way near complete!**

Nintendo Gameboy Color emulator, this is my first emulator so I'm learning as I go along...

This repository is the frontend for running the emulator in a browser, by compiling it to WebAssembly. The backend that implements the emulator can be found in: [djhworld/gomeboycolor](https://github.com/djhworld/gomeboycolor)


![mario tennis](https://github.com/djhworld/gomeboycolor-wasm/raw/master/mariotennis.png)

You are welcome to visit the github page for this project by [clicking here](http://djhworld.github.io/gomeboycolor)

FAQ
============================

### How do I build it?

#### Pre-requisites 

You will need 

* an installation of [Go](http://golang.org) (version >= 1.11.0)

#### Run

```
make
```

This will install the WASM binary under `static/gomeboycolor/wasm/`


### How do I run it?

Either:

* Host the contents of `static/` on a web server and navigate to `index.html` in a browser
* Run the small server I've provided and visit http://localhost:8080

```
$ make server
$ gomeboycolor-wasm-server
Open your web browser and navigate to: http://localhost:8080
```

### Features

See [here](https://github.com/djhworld/gomeboycolor#features)

Battery saves are stored on in the browser LocalStorage


License
-----------------------------

MIT License

Copyright (c) 2013-2018 Daniel James Harper

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

Progress
---------------------------

![mario tennis](https://github.com/djhworld/gomeboycolor-wasm/raw/master/mariotennis.png)
![pokemon](https://github.com/djhworld/gomeboycolor-wasm/raw/master/pokemon.png)
![tetrisdx](https://github.com/djhworld/gomeboycolor-wasm/raw/master/tetrisdx.png)
![zelda](https://github.com/djhworld/gomeboycolor-wasm/raw/master/zelda.png)
