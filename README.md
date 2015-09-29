# GearGen
A simple gear profile generator written in go and output to SVG.

## Install
Install the app source with the the go tool.

    go get github.com/stuphi/GearGen
    go install github.com/stuphi/GearGen

##Use
It will ask for the distance between centers, the number of teeth on drive and driven gear and the pressure angle.

The progrem then plots each gear on a 5mm grid background.

The output currently goes to the the file TestFile.svg A simple script convert.sh is provided to generate a PDF, PNG and JPG version of the plot for esy viewing.

![Example](/Example.png)
