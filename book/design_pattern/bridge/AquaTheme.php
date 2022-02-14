<?php

namespace book\design_pattern\bridge;

class AquaTheme implements Theme
{
    public function getColor(): string
    {
        return 'Light blue';
    }
}