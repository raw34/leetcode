<?php

namespace book\design_pattern\command;

class Bulb
{
    public function turnOn(): string
    {
        return "Bulb has been lit!";
    }

    public function turnOff(): string
    {
        return "Darkness!";
    }
}