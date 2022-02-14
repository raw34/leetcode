<?php

namespace book\design_pattern\visitor;

class Dolphin implements Animal
{
    public function speak(): string
    {
        return 'Tuut tutt tuutt!';
    }

    public function accept(AnimalOperation $operation): string
    {
        return $operation->visitDolphin($this);
    }
}