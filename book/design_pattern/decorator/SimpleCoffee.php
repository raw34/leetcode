<?php

namespace book\design_pattern\decorator;

class SimpleCoffee implements Coffee
{
    public function getCost(): float
    {
        return 10;
    }

    public function getDescription(): string
    {
        return 'Simple Coffee';
    }
}