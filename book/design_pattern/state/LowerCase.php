<?php

namespace book\design_pattern\state;

class LowerCase implements WritingState
{
    public function write(string $words): string
    {
        return strtolower($words);
    }
}