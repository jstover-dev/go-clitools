// main_test.go
package main

import (
    "testing"
	"codeberg.org/jstover/clitools/pkg/cat"
)


func TestGetVal(t *testing.T) {

    var argv = []string{
        "-n",
        "/home/fenn/Pictures/Wallpapers/BlueSky-2650x1600.jpg", 
        "/home/fenn/Pictures/Wallpapers/ColdRipple-2560x1600.jpg", 
        "/home/fenn/Pictures/Wallpapers/DarkestHour-2560x1600.jpg", 
        "/home/fenn/Pictures/Wallpapers/Jupiter.jpg", 
        "/home/fenn/Pictures/Wallpapers/OneStandsOut-2560x1600.jpg", 
        "/home/fenn/Pictures/Wallpapers/Path-2560x1600.jpg", 
        "/home/fenn/Pictures/Wallpapers/Summer_1AM-2560x1600.jpg", 
        "/home/fenn/Pictures/Wallpapers/abstract-1779631-3840x2160.jpg", 
        "/home/fenn/Pictures/Wallpapers/abstract-1780241-3840x2160.jpg", 
        "/home/fenn/Pictures/Wallpapers/Sunset-wide.jpg", 
        "/home/fenn/Pictures/Wallpapers/b4c513e1d5f5ad7f3df7433867097c53.jpg", 
        "/home/fenn/Pictures/Wallpapers/antelope-canyon-984055-3240x2160.jpg", 
        "/home/fenn/Pictures/Wallpapers/city-sleeps-2560x1440.jpg",
        "/home/fenn/Pictures/Wallpapers/dreams-4674x2629.jpg",
        "/home/fenn/Pictures/Wallpapers/mountains-1412683-4000x2000.jpg",
        "/home/fenn/Pictures/Wallpapers/nature-3181144-3840x2160.jpg",
        "/home/fenn/Pictures/Wallpapers/nature-3058859-3840x2160.jpg", 
        "/home/fenn/Pictures/Wallpapers/sun-3845x2160.jpg",
        "/home/fenn/Pictures/Wallpapers/sunset-3840x2160.jpg",
        "/home/fenn/Pictures/Wallpapers/panorama-3629120.jpg",
    }

    cat.Cat(argv)
}
