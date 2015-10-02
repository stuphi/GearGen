#!/bin/bash
#     genanim.sh -- Simple utility to generate a sequence of images for
#     animation
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

function gencon {
  f=`printf "anim_%02d" $1`
  echo "Working on frame $1"
  GearGen -o $f -r $1 -n1 17 -n2 23
  inkscape --without-gui --file $f.svg --export-dpi=150 --export-png=$f.png &>/dev/null
  convert $f.png -flatten -resize 50% $f.png
  rm $f.svg
}

for ((i=0; i<100; i++))
do
  gencon $i &
done
echo "Waiting for jobs to complete."
wait
echo "Now make animation..."
convert -delay 5 -loop 0 anim*.png animation.gif
rm anim_??.png
echo "All done."
