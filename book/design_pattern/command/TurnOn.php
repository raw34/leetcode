<?php

namespace book\design_pattern\command;

class TurnOn implements Command
{
    protected $bulb;

    public function __construct(Bulb $bulb)
    {
        $this->bulb = $bulb;
    }

    public function execute(): string
    {
        return $this->bulb->turnOn();
    }

    public function undo(): string
    {
        return $this->bulb->turnOff();
    }

    public function redo(): string
    {
        return $this->execute();
    }
}