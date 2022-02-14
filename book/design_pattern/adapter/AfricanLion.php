<?php

namespace book\design_pattern\adapter;

class AfricanLion implements Lion
{
    public function roar(): string
    {
        return 'African Lion: Roar!';
    }
}