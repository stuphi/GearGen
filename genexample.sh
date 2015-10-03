#!/bin/bash
#     genexample.sh -- Simple utility to generate alternative formats of example
#     file
#     Copyright (C) 2015  Philip Stubbs
#
#     This program is free software: you can redistribute it and/or modify
#     it under the terms of the GNU General Public License as published by
#     the Free Software Foundation, either version 3 of the License, or
#     (at your option) any later version.
#
#     This program is distributed in the hope that it will be useful,
#     but WITHOUT ANY WARRANTY; without even the implied warranty of
#     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#     GNU General Public License for more details.
#
#     You should have received a copy of the GNU General Public License
#     along with this program.  If not, see <http://www.gnu.org/licenses/>.

GearGen -o Example
inkscape --without-gui --file Example.svg --export-dpi=300 \
  --export-png=Example.png &>/dev/null
inkscape --without-gui --file Example.svg --export-pdf=Example.pdf &>/dev/null
convert Example.png -flatten Example.jpg
