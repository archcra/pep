#!/bin/bash

. ~/.bashrc
go test github.com/archcra/pep/minimax -cpuprofile cpu.out
