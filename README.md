Simple CLI to get emoji's into your clip
========================================

* Uses `github.com/atotto/clipboard` for clipboard 
* Built on `github.com/spf13/cobra`
* Tested `xclip` install on ubuntu - `sudo apt-get install -y xclip` read the `github.com/atotto/clipboard` docs for more details

USE
---
* **Build**
```
    #Will build the code in the bin/ folder in this repo.
    make build
```

* The `emojis` directory in your path or in the home directory `$HOME/emojis` will be used, and the `essentials.yaml/essentials.json/essentials.toml` will be picked up. Look at [viper](https://github.com/spf13/viper) for more details.

```
    emoji get grin => π
```
```
    emoji list
==============πTHE AVAILABLE EMOJISπ============
==== The name thumbsup => π 
==== The name triplefinger => πππ 
==== The name angry => π  
==== The name finger => π 
==== The name grin => π 
==== The name shrug => π€·
```

Useful
------

`set PATH=$PATH:$PWD/bin` for ease.

Resources
---------
[GopherCon 2019: Carolyn Van Slyck - Design Command-Line Tools People Love](https://www.youtube.com/watch?v=eMz0vni6PAw)