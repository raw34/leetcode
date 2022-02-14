<?php

namespace book\design_pattern\abstract_factory;

class IronDoor implements Door
{
    public function getDescription(): string
    {
        return 'I am an iron door';
    }
}