## Compiler Analyses

<!-- Please edit the `README.md.tmpl` file instead of the `README.md` -->

These are tables that analyse different compilers for different languages.

#### Build Time

|Language/Compiler|Command|Time|
|-----------------|-------|----|
|C (`tcc`)|`tcc code/hello_world.c`|`0m0.009s`|
|C (`gcc`)|`gcc code/hello_world.c`|`0m0.217s`|
|C (`cc`)|`cc code/hello_world.c`|`0m0.036s`|
|Python|`python3 -m py_compile code/hello_world.py`|`0m0.026s`|
|Go|`go build code/hello_world.go`|`0m0.222s`|
|Rust|`rustc code/hello_world.rs`|`0m4.948s`|

#### Run Time

|Language/Compiler|Command|Time|
|-----------------|-------|----|
|C (`tcc`)|`./a.out`|`0m0.001s`|
|C (`gcc`)|`./a.out`|`0m0.001s`|
|C (`cc`)|`./a.out`|`0m0.001s`|
|Python|`python3 code/hello_world.py`|`0m0.019s`|
|Go|`./hello_world`|`0m0.001s`|
|Rust|`./hello_world`|`0m0.001s`|

#### Binary Size

|Language/Compiler|Size|
|-----------------|----|
|C (`tcc`)|`20K`|
|C (`gcc`)|`20K`|
|C (`cc`)|`20K`|
|Python|`4.0K`|
|Go|`3.8M`|
|Rust|`3.8M`|
