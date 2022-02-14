<?php

namespace book\design_pattern\bridge;

class LightTheme implements Theme
{
    public function getColor(): string
    {
        return 'Off White';
    }
}
