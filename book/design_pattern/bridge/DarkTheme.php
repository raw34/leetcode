<?php

namespace book\design_pattern\bridge;

class DarkTheme implements Theme
{
    public function getColor(): string
    {
        return 'Dark Black';
    }
}