#!/bin/bash

inkscape --without-gui --file TestFile.svg --export-png=TestFile.png
inkscape --without-gui --file TestFile.svg --export-pdf=TestFile.pdf
convert TestFile.png -flatten TestFile.jpg
