<?php

namespace book\design_pattern\visitor;

interface Animal
{
    public function accept(AnimalOperation $operation): string;
}