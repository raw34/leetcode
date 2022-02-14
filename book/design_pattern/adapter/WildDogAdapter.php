<?php

namespace book\design_pattern\adapter;

class WildDogAdapter implements Lion
{
    private WildDog $dog;

    public function __construct(WildDog $dog)
    {
        $this->dog = $dog;
    }

    public function roar(): string
    {
        return $this->dog->bark();
    }
}