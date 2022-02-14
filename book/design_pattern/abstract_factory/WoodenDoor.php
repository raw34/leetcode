<?php

namespace book\design_pattern\abstract_factory;

class WoodenDoor implements Door
{
    public function getDescription(): string
    {
        return 'I am a wooden door';
    }
}