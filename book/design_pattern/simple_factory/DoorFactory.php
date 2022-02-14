<?php

namespace book\design_pattern\simple_factory;

class DoorFactory
{
    public static function makeDoor(float $width, float $height): Door
    {
        return new WoodenDoor($width, $height);
    }
}

//
//// Make me a door of 100x200
//$door = DoorFactory::makeDoor(100, 200);
//
//echo 'Width: ' . $door->getWidth();
//echo 'Height: ' . $door->getHeight();
//
//// Make me a door of 50x100
//$door2 = DoorFactory::makeDoor(50, 100);
