<?php

namespace book\design_pattern\command;

class TurnOff implements Command
{
    private Bulb $bulb;

    public function __construct(Bulb $bulb)
    {
        $this->bulb = $bulb;
    }

    public function execute(): string
    {
        return $this->bulb->turnOff();
    }

    public function undo(): string
    {
        return $this->bulb->turnOn();
    }

    public function redo(): string
    {
        return $this->execute();
    }
}