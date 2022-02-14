<?php

namespace book\design_pattern\abstract_factory;

class Welder implements DoorFittingExpert
{
    public function getDescription(): string
    {
        return 'I can only fit iron doors';
    }
}