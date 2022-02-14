<?php

namespace book\design_pattern\simple_factory;

use PHPUnit\Framework\TestCase;

class SimpleFactoryTest extends TestCase
{
    public function testCase1()
    {
        $door = DoorFactory::makeDoor(100, 200);
        self::assertEquals(100, $door->getWidth());
        self::assertEquals(200, $door->getHeight());

        $door2 = DoorFactory::makeDoor(50, 100);
        self::assertEquals(50, $door2->getWidth());
        self::assertEquals(100, $door2->getHeight());
    }
}
