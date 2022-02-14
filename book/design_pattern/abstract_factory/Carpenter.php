<?php

namespace book\design_pattern\abstract_factory;

class Carpenter implements DoorFittingExpert
{
    public function getDescription(): string
    {
        return 'I can only fit wooden doors';
    }
}