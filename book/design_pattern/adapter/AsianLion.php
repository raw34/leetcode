<?php

namespace book\design_pattern\adapter;

class AsianLion implements Lion
{
    public function roar(): string
    {
        return 'Asian Lion: Roar!';
    }
}