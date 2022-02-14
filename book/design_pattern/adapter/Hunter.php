<?php

namespace book\design_pattern\adapter;

class Hunter
{
    public function hunt(Lion $lion): string
    {
        return $lion->roar();
    }
}