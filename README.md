# Plank
## Description
Plank is a simple benchmarking set of tools. This was originally created for benchmarking ec2 instances to see if there is a noisy neighbor problem and terminate and launch another one if needed. 

What this is and what it's not. This set of tests is not to provide anything particularly telling about a system. What it is, is a set of tests that can be ran with the same parameters on systems. It is intended that you know what a _good_ system is and you keep the results of these tests as a _good_ result. You then run the same tests on future machines to make sure they are not worse than your … _benchmark_

Currently there are two tools a cpu and or memory test that runs iternations of a recursive fibonacci function and a disk / cpu test that writes a file of specified sizes filled with zeroes. The fib uses a lot of memory and is susceptible to OOM kills. For example on a 16GB machine without swap 33 iterations is enough to get it killed. The disk test is done via writing zeroes via a loop so it does stress the cpu as well while its writing. 

## Instalation
Intended to be ran as docker container. Rest of info provided later 

## Usage
cpu / memory test

`./cpu_fib -i 32`

cpu / disk test

`./disk -s 100`

## Documentation
Still to come
### options
Still to come

* option 1
* option 2

